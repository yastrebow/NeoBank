package main

import (
	"context"
	"encoding/json"
	"fmt"
	"msa-bank-report-ss/models"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"

	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:29092"},
		GroupID:     "consumer-group-id",
		GroupTopics: []string{"dev.msa_bank_product_cs_schema.product", "dev.msa_bank_credit_cs_schema.credit", "dev.msa_bank_client_cs_schema.client"},
		MinBytes:    10e2, // 1KB
		MaxBytes:    10e6, // 10MB
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		// fmt.Printf("kafka message at topic/partition/offset %v/%v/%v: %s = %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		var m models.Message
		err = json.Unmarshal(msg.Value, &m)
		if err != nil {
			panic("could not parse message from kafka " + err.Error())
		}

		table := m.Payload.Source.Table
		after, err := json.Marshal(m.Payload.After)
		if err != nil {
			panic(err)
		}

		var IDS models.IDS
		err = json.Unmarshal(m.Payload.After, &IDS)
		if err != nil {
			panic("could not parse clientID " + err.Error())
		}

		switch table {

		case "credit":
			cmd := redis.NewStringCmd(ctx, "select", 0)
			err = rdb.Process(ctx, cmd)
			if err != nil {
				panic("could not change DB redis " + err.Error())
			}
			err = rdb.HSet(ctx, IDS.ClientId, IDS.Id, after).Err()
			if err != nil {
				panic("could not write to redis " + err.Error())
			}

			val, err := rdb.HGetAll(ctx, IDS.ClientId).Result()
			if err != nil {
				panic(err)
			}
			fmt.Printf("credit redis key: %s val = %s\n", IDS.ClientId, val)
			

		case "client":

			cmd := redis.NewStringCmd(ctx, "select", 1)
			err = rdb.Process(ctx, cmd)
			if err != nil {
				panic("could not change DB redis " + err.Error())
			}
			err = rdb.Set(ctx, IDS.Id, after, 0).Err()
			if err != nil {
				panic("could not write to redis " + err.Error())
			}

			val, err := rdb.Get(ctx, IDS.Id).Result()
			if err != nil {
				panic(err)
			}
			fmt.Println("client redis key -", IDS.Id)
			fmt.Println("client redis val -", val)

		case "product":
			cmd := redis.NewStringCmd(ctx, "select", 2)
			err = rdb.Process(ctx, cmd)
			if err != nil {
				panic("could not change DB redis " + err.Error())
			}
			err = rdb.HSet(ctx, IDS.ClientId, IDS.Id, after).Err()
			if err != nil {
				panic("could not write to redis " + err.Error())
			}

			val, err := rdb.HGetAll(ctx, IDS.ClientId).Result()
			if err != nil {
				panic(err)
			}
			fmt.Printf("product redis key: %s val = %s\n", IDS.ClientId, val)

		}


	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
