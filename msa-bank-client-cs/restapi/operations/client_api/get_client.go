// Code generated by go-swagger; DO NOT EDIT.

package client_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClientHandlerFunc turns a function with the right signature into a get client handler
type GetClientHandlerFunc func(GetClientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClientHandlerFunc) Handle(params GetClientParams) middleware.Responder {
	return fn(params)
}

// GetClientHandler interface for that can handle valid get client params
type GetClientHandler interface {
	Handle(GetClientParams) middleware.Responder
}

// NewGetClient creates a new http.Handler for the get client operation
func NewGetClient(ctx *middleware.Context, handler GetClientHandler) *GetClient {
	return &GetClient{Context: ctx, Handler: handler}
}

/*
	GetClient swagger:route GET /client/{id} client-api getClient

# Get client from the store by id

Get client
*/
type GetClient struct {
	Context *middleware.Context
	Handler GetClientHandler
}

func (o *GetClient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClientParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
