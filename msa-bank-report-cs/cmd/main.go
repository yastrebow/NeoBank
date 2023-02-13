package main

import (
	"context"
	"msa-bank-report-cs/models"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client

var ctx = context.Background()

const portNumber = 8084
func main() {
	
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	r := mux.NewRouter()
	r.HandleFunc("/report/{clientId}", handler)
	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", portNumber),
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	clientId := vars["clientId"]
	log.Info("ClientId: ", clientId)
	response := &models.Report{}
	cmd := redis.NewStringCmd(ctx, "select", 0)
	err := rdb.Process(ctx, cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	credits := []models.Credit{}
	val, err := rdb.HGetAll(ctx, clientId).Result()
	if err != nil {
		panic(err)
	}
	for _, item := range val {
		creditModel := &models.Credit{}
		err := json.Unmarshal([]byte(item), creditModel)
		if err != nil {
			panic(err)
		}
		
		log.Info("credits - ", creditModel)
		credits = append(credits, *creditModel)
	}
	response.Credit = credits

	cmd = redis.NewStringCmd(ctx, "select", 1)
	err = rdb.Process(ctx, cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client, err := rdb.Get(ctx, clientId).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case client == "":
		fmt.Println("value is empty")
	}

	clientModel := &models.Client{}
	err = json.Unmarshal([]byte(client), clientModel)
	if err != nil {
    	return
	}
	
	response.Client = *clientModel
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}