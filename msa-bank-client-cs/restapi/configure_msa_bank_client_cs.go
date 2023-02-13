// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"msa-bank-client/models"
	"msa-bank-client/pkg/db"
	"msa-bank-client/pkg/handlers"
	"msa-bank-client/restapi/operations"
	"msa-bank-client/restapi/operations/client_api"

	log "github.com/sirupsen/logrus"
)

//go:generate swagger generate server --target ../../src --name MsaBankClientCs --spec ../swagger-api/msa-bank-client-cs.yml --principal interface{}

func configureFlags(api *operations.MsaBankClientCsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MsaBankClientCsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	DB := db.Init()
	h := handlers.New(DB)
	db.Migration()

	// Create client ...
	api.ClientAPIAddClientHandler = client_api.AddClientHandlerFunc(func(params client_api.AddClientParams) middleware.Responder {
		// Append to the Client table
		var client models.Client = *params.Body
		if result := h.DB.Create(&client); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return client_api.NewAddClientInternalServerError().WithPayload(&addError)
		}
		return client_api.NewAddClientCreated()
	})

	// Delete client ...
	api.ClientAPIDeleteClientHandler = client_api.DeleteClientHandlerFunc(func(params client_api.DeleteClientParams) middleware.Responder {
		var id = params.ID
		var client models.Client
		if result := h.DB.First(&client, "id = ?", id); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return client_api.NewDeleteClientInternalServerError().WithPayload(&addError)
		}
		// Delete that client
		h.DB.Delete(&client)
		return client_api.NewDeleteClientOK()
	})

	// Get client ...
	api.ClientAPIGetClientHandler = client_api.GetClientHandlerFunc(func(params client_api.GetClientParams) middleware.Responder {
		var id = params.ID
		var client models.Client
		if result := h.DB.First(&client, "id = ?", id); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return client_api.NewGetClientInternalServerError().WithPayload(&addError)
		}
		return client_api.NewGetClientOK().WithPayload(&client)
	})

	// Get all clients ...
	api.ClientAPIGetClientsHandler = client_api.GetClientsHandlerFunc(func(params client_api.GetClientsParams) middleware.Responder {
		var clients []*models.Client
		if result := h.DB.Find(&clients); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return client_api.NewGetClientInternalServerError().WithPayload(&addError)
		}
		return client_api.NewGetClientsOK().WithPayload(clients)
	})

	// Update client ...
	api.ClientAPIUpdateClientHandler = client_api.UpdateClientHandlerFunc(func(params client_api.UpdateClientParams) middleware.Responder {
		var id = params.Body.ID
		var updatedClient models.Client = *params.Body
		var client models.Client
		if result := h.DB.First(&client, "id = ?", id); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return client_api.NewUpdateClientInternalServerError().WithPayload(&addError)
		}
		client.ID = updatedClient.ID
		client.FirstName = updatedClient.FirstName
		client.LastName = updatedClient.LastName
		client.BirthDate = updatedClient.BirthDate
		h.DB.Save(&client)
		return client_api.NewUpdateClientOK().WithPayload(&client)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
