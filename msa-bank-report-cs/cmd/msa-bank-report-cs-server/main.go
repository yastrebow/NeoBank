package main

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"msa-bank-report-cs/models"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

var rdb *redis.Client

var ctx = context.Background()

const portNumber = 8084

type ConfigDatabase struct {
	Port string `yaml:"port" env:"REDIS-PORT" env-default:"6379"`
	Host string `yaml:"host" env:"REDIS-HOST" env-default:"localhost"`
}

func main() {

	var cfg ConfigDatabase

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln("Cannot read config", err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	r := mux.NewRouter()
	r.HandleFunc("/report/{clientId}", handler)
	r.HandleFunc("/download-report/{clientId}", downloadExcel)
	fmt.Printf("Starting msa-bank-report-cs server on port %v\n", portNumber)
	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", portNumber),
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["clientId"]
	response := ReadRedis(clientId)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func ReadRedis(clientId string) *models.Report {
	response := &models.Report{}
	cmd := redis.NewStringCmd(ctx, "select", 0)
	err := rdb.Process(ctx, cmd)
	if err != nil {
		panic(err)
	}
	credits := []models.Credit{}
	credit, err := rdb.HGetAll(ctx, clientId).Result()
	if err != nil {
		panic(err)
	}
	for _, item := range credit {
		creditModel := &models.Credit{}
		err := json.Unmarshal([]byte(item), creditModel)
		if err != nil {
			panic(err)
		}

		log.Info("credits - ", creditModel)
		credits = append(credits, *creditModel)
	}
	response.Credits = credits

	cmd = redis.NewStringCmd(ctx, "select", 1)
	err = rdb.Process(ctx, cmd)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	response.Client = *clientModel

	cmd = redis.NewStringCmd(ctx, "select", 2)
	err = rdb.Process(ctx, cmd)
	if err != nil {
		panic(err)
	}
	accounts := []models.Account{}
	account, err := rdb.HGetAll(ctx, clientId).Result()
	if err != nil {
		panic(err)
	}
	for _, item := range account {
		accountModel := &models.Account{}
		err := json.Unmarshal([]byte(item), accountModel)
		if err != nil {
			panic(err)
		}

		log.Info("accounts - ", accountModel)
		accounts = append(accounts, *accountModel)
	}
	response.Accounts = accounts
	return response
}

func PrepareAndReturnExcel(userInputData *models.Report) *excelize.File {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Отчет по клиенту:")
	f.SetCellValue("Sheet1", "A2", "Фамилия - ")
	f.SetCellValue("Sheet1", "B2", userInputData.Client.LastName)
	f.SetCellValue("Sheet1", "A3", "Имя - ")
	f.SetCellValue("Sheet1", "B3", userInputData.Client.FirstName)
	f.SetCellValue("Sheet1", "A4", "Дата рождения - ")
	f.SetCellValue("Sheet1", "B4", userInputData.Client.BirthDate)
	f.SetCellValue("Sheet1", "A5", "ИД клиета - ")
	f.SetCellValue("Sheet1", "B5", userInputData.Client.ID)
	f.SetCellValue("Sheet1", "C6", "Счета клиента:")
	startRow := 7
	for i := startRow; i < (len(userInputData.Accounts) + startRow); i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i), "ИД счёта: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i), userInputData.Accounts[i-startRow].Id)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i), "Дата начала действия счёта: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i), userInputData.Accounts[i-startRow].StartDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i), "Дата окончания действия счёта: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i), userInputData.Accounts[i-startRow].EndDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i), "Сумма: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", i), userInputData.Accounts[i-startRow].Amount)
	}

	startRow += len(userInputData.Accounts)
	f.SetCellValue("Sheet1", fmt.Sprintf("C%d", startRow), "Кредиты клиента:")
	startRow++
	for i := startRow; i < (len(userInputData.Credits) + startRow); i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i), "ИД кредита: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i), userInputData.Credits[i-startRow].Id)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i), "Количество месяцев: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i), userInputData.Credits[i-startRow].Months)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i), "Ставка: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i), userInputData.Credits[i-startRow].Rate)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i), "Задолженность по кредиту: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i), userInputData.Credits[i-startRow].Amount)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i), "Сумма кредита: ")
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", i), userInputData.Credits[i-startRow].TotalAmount)
	}

	return f
}

func downloadExcel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["clientId"]
	userInputData := ReadRedis(clientId)
	// Get the Excel file with the user input data
	file := PrepareAndReturnExcel(userInputData)

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("File-Name", "report.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err := file.Write(w)
	if err != nil {
		panic(err)
	}
}
