// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Toha22BSK/UrlShortener/gen/models"
)

// GetAnalyticsOKCode is the HTTP code returned for type GetAnalyticsOK
const GetAnalyticsOKCode int = 200

/*
GetAnalyticsOK OK

swagger:response getAnalyticsOK
*/
type GetAnalyticsOK struct {

	/*
	  In: Body
	*/
	Payload *GetAnalyticsOKBody `json:"body,omitempty"`
}

// NewGetAnalyticsOK creates GetAnalyticsOK with default headers values
func NewGetAnalyticsOK() *GetAnalyticsOK {

	return &GetAnalyticsOK{}
}

// WithPayload adds the payload to the get analytics o k response
func (o *GetAnalyticsOK) WithPayload(payload *GetAnalyticsOKBody) *GetAnalyticsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics o k response
func (o *GetAnalyticsOK) SetPayload(payload *GetAnalyticsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnalyticsUnauthorizedCode is the HTTP code returned for type GetAnalyticsUnauthorized
const GetAnalyticsUnauthorizedCode int = 401

/*
GetAnalyticsUnauthorized Unauthorized

swagger:response getAnalyticsUnauthorized
*/
type GetAnalyticsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorV1 `json:"body,omitempty"`
}

// NewGetAnalyticsUnauthorized creates GetAnalyticsUnauthorized with default headers values
func NewGetAnalyticsUnauthorized() *GetAnalyticsUnauthorized {

	return &GetAnalyticsUnauthorized{}
}

// WithPayload adds the payload to the get analytics unauthorized response
func (o *GetAnalyticsUnauthorized) WithPayload(payload *models.ErrorV1) *GetAnalyticsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics unauthorized response
func (o *GetAnalyticsUnauthorized) SetPayload(payload *models.ErrorV1) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnalyticsNotFoundCode is the HTTP code returned for type GetAnalyticsNotFound
const GetAnalyticsNotFoundCode int = 404

/*
GetAnalyticsNotFound Not found.

swagger:response getAnalyticsNotFound
*/
type GetAnalyticsNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAnalyticsNotFound creates GetAnalyticsNotFound with default headers values
func NewGetAnalyticsNotFound() *GetAnalyticsNotFound {

	return &GetAnalyticsNotFound{}
}

// WithPayload adds the payload to the get analytics not found response
func (o *GetAnalyticsNotFound) WithPayload(payload string) *GetAnalyticsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics not found response
func (o *GetAnalyticsNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAnalyticsInternalServerErrorCode is the HTTP code returned for type GetAnalyticsInternalServerError
const GetAnalyticsInternalServerErrorCode int = 500

/*
GetAnalyticsInternalServerError Internal Server Error

swagger:response getAnalyticsInternalServerError
*/
type GetAnalyticsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorV1 `json:"body,omitempty"`
}

// NewGetAnalyticsInternalServerError creates GetAnalyticsInternalServerError with default headers values
func NewGetAnalyticsInternalServerError() *GetAnalyticsInternalServerError {

	return &GetAnalyticsInternalServerError{}
}

// WithPayload adds the payload to the get analytics internal server error response
func (o *GetAnalyticsInternalServerError) WithPayload(payload *models.ErrorV1) *GetAnalyticsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get analytics internal server error response
func (o *GetAnalyticsInternalServerError) SetPayload(payload *models.ErrorV1) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnalyticsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}