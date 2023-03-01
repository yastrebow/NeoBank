package main

import (
	// "flag"
	"fmt"
	"log"
	"net/http"
	"os"

	 middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/gorilla/mux"
	"msa-bank-account-cs/restapi"
    "msa-bank-account-cs/pkg/db"
	"msa-bank-account-cs/pkg/repository"
	"msa-bank-account-cs/pkg/services"
)

func main() {

	swagger, err := restapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// // Clear out the servers array in the swagger spec, that skips validating
	// // that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

    DB := db.Init()
	db.Migration()
	accountRepo := repository.NewAccount(DB)
	service := services.AccountService(accountRepo)
	server := restapi.NewAccountApi(service)

	// // This is how you set up a basic Gorilla router
	r := mux.NewRouter()

	// // Use our validation middleware to check all requests against the
	// // OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// // We now register our server above as the handler for the interface
	restapi.HandlerFromMux(server, r)
	

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8080),
	}
	log.Println("start")
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())

}

