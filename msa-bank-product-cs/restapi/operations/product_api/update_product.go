// Code generated by go-swagger; DO NOT EDIT.

package product_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateProductHandlerFunc turns a function with the right signature into a update product handler
type UpdateProductHandlerFunc func(UpdateProductParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateProductHandlerFunc) Handle(params UpdateProductParams) middleware.Responder {
	return fn(params)
}

// UpdateProductHandler interface for that can handle valid update product params
type UpdateProductHandler interface {
	Handle(UpdateProductParams) middleware.Responder
}

// NewUpdateProduct creates a new http.Handler for the update product operation
func NewUpdateProduct(ctx *middleware.Context, handler UpdateProductHandler) *UpdateProduct {
	return &UpdateProduct{Context: ctx, Handler: handler}
}

/*
	UpdateProduct swagger:route PUT /product product-api updateProduct

Update an existing product
*/
type UpdateProduct struct {
	Context *middleware.Context
	Handler UpdateProductHandler
}

func (o *UpdateProduct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateProductParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}