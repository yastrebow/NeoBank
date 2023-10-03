package main

import (
	"fmt"

	"msa-bank-credit-cs/pkg/db"
	"msa-bank-credit-cs/pkg/handlers"
	"msa-bank-credit-cs/pkg/server"
	"msa-bank-credit-cs/pkg/services"
	"msa-bank-credit-cs/repository"
	"msa-bank-credit-cs/restapi"
)

func main() {
	e := server.NewEchoServer()
	database := db.Init()
	db.Migration()
	creditRepo := repository.NewCredit(database)
	repaymentRepo := repository.NewRepayment(database)

	creditService := services.NewCredit(creditRepo)
	repaymentService := services.NewRepayment(repaymentRepo)

	h := handlers.NewHandler(creditService, repaymentService)
	restapi.RegisterHandlers(e, h)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", "8081")))

}
