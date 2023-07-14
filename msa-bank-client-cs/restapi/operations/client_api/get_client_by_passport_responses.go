// Code generated by go-swagger; DO NOT EDIT.

package client_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"msa-bank-client/models"
)

// GetClientByPassportOKCode is the HTTP code returned for type GetClientByPassportOK
const GetClientByPassportOKCode int = 200

/*
GetClientByPassportOK successful operation

swagger:response getClientByPassportOK
*/
type GetClientByPassportOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Client `json:"body,omitempty"`
}

// NewGetClientByPassportOK creates GetClientByPassportOK with default headers values
func NewGetClientByPassportOK() *GetClientByPassportOK {

	return &GetClientByPassportOK{}
}

// WithPayload adds the payload to the get client by passport o k response
func (o *GetClientByPassportOK) WithPayload(payload []*models.Client) *GetClientByPassportOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get client by passport o k response
func (o *GetClientByPassportOK) SetPayload(payload []*models.Client) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientByPassportOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Client, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetClientByPassportMethodNotAllowedCode is the HTTP code returned for type GetClientByPassportMethodNotAllowed
const GetClientByPassportMethodNotAllowedCode int = 405

/*
GetClientByPassportMethodNotAllowed Invalid input

swagger:response getClientByPassportMethodNotAllowed
*/
type GetClientByPassportMethodNotAllowed struct {
}

// NewGetClientByPassportMethodNotAllowed creates GetClientByPassportMethodNotAllowed with default headers values
func NewGetClientByPassportMethodNotAllowed() *GetClientByPassportMethodNotAllowed {

	return &GetClientByPassportMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetClientByPassportMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// GetClientByPassportInternalServerErrorCode is the HTTP code returned for type GetClientByPassportInternalServerError
const GetClientByPassportInternalServerErrorCode int = 500

/*
GetClientByPassportInternalServerError Internal server error

swagger:response getClientByPassportInternalServerError
*/
type GetClientByPassportInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetClientByPassportInternalServerError creates GetClientByPassportInternalServerError with default headers values
func NewGetClientByPassportInternalServerError() *GetClientByPassportInternalServerError {

	return &GetClientByPassportInternalServerError{}
}

// WithPayload adds the payload to the get client by passport internal server error response
func (o *GetClientByPassportInternalServerError) WithPayload(payload *models.Error) *GetClientByPassportInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get client by passport internal server error response
func (o *GetClientByPassportInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientByPassportInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}