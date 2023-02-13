// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"msa-bank-product/models"
	"msa-bank-product/pkg/db"
	"msa-bank-product/pkg/handlers"
	"msa-bank-product/restapi/operations"
	"msa-bank-product/restapi/operations/product_api"

	log "github.com/sirupsen/logrus"
)

//go:generate swagger generate server --target ../../src --name MsaBankProductCs --spec ../api/msa-bank-product-cs.yml --principal interface{}

func configureFlags(api *operations.MsaBankProductCsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MsaBankProductCsAPI) http.Handler {
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

	// Create product ...
	api.ProductAPIAddProductHandler = product_api.AddProductHandlerFunc(func(params product_api.AddProductParams) middleware.Responder {
		// Append to the Product table
		var product models.Product = *params.Body
		if result := h.DB.Create(&product); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return product_api.NewAddProductInternalServerError().WithPayload(&addError)
		}
		return product_api.NewAddProductCreated()
	})

	// Delete product ...
	api.ProductAPIDeleteProductHandler = product_api.DeleteProductHandlerFunc(func(params product_api.DeleteProductParams) middleware.Responder {
		var id = params.ID
		var product models.Product
		if result := h.DB.First(&product, "id = ?", id); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return product_api.NewDeleteProductInternalServerError().WithPayload(&addError)
		}
		// Delete that product
		h.DB.Delete(&product)
		return product_api.NewDeleteProductOK()
	})

	// Get product ...
	api.ProductAPIGetProductHandler = product_api.GetProductHandlerFunc(func(params product_api.GetProductParams) middleware.Responder {
		var id = params.ID
		var product models.Product
		if result := h.DB.First(&product, "id = ?", id); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return product_api.NewGetProductInternalServerError().WithPayload(&addError)
		}
		return product_api.NewGetProductOK().WithPayload(&product)
	})

	// Get all product ...
	api.ProductAPIGetProductsHandler = product_api.GetProductsHandlerFunc(func(params product_api.GetProductsParams) middleware.Responder {
		var product []*models.Product
		if result := h.DB.Find(&product); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return product_api.NewGetProductInternalServerError().WithPayload(&addError)
		}
		return product_api.NewGetProductsOK().WithPayload(product)
	})

	// Update product ...
	api.ProductAPIUpdateProductHandler = product_api.UpdateProductHandlerFunc(func(params product_api.UpdateProductParams) middleware.Responder {
		var id = params.Body.ID
		var updatedProduct models.Product = *params.Body
		var product models.Product
		if result := h.DB.First(&product, "id = ?", id); result.Error != nil {
			log.Error(result.Error)
			var addError models.Error
			addError.ErrorMessage = result.Error.Error()
			return product_api.NewUpdateProductInternalServerError().WithPayload(&addError)
		}
		product.ID = updatedProduct.ID
		product.Name = updatedProduct.Name
		product.Description = updatedProduct.Description

		h.DB.Save(&product)
		return product_api.NewUpdateProductOK().WithPayload(&product)
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
