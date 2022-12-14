// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hanoman-id/xendit-payment/internal/models"
)

// GetRedirectOKCode is the HTTP code returned for type GetRedirectOK
const GetRedirectOKCode int = 200

/*GetRedirectOK Success

swagger:response getRedirectOK
*/
type GetRedirectOK struct {

	/*
	  In: Body
	*/
	Payload *GetRedirectOKBody `json:"body,omitempty"`
}

// NewGetRedirectOK creates GetRedirectOK with default headers values
func NewGetRedirectOK() *GetRedirectOK {

	return &GetRedirectOK{}
}

// WithPayload adds the payload to the get redirect o k response
func (o *GetRedirectOK) WithPayload(payload *GetRedirectOKBody) *GetRedirectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get redirect o k response
func (o *GetRedirectOK) SetPayload(payload *GetRedirectOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRedirectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRedirectBadRequestCode is the HTTP code returned for type GetRedirectBadRequest
const GetRedirectBadRequestCode int = 400

/*GetRedirectBadRequest Unauthorized

swagger:response getRedirectBadRequest
*/
type GetRedirectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRedirectBadRequest creates GetRedirectBadRequest with default headers values
func NewGetRedirectBadRequest() *GetRedirectBadRequest {

	return &GetRedirectBadRequest{}
}

// WithPayload adds the payload to the get redirect bad request response
func (o *GetRedirectBadRequest) WithPayload(payload *models.Error) *GetRedirectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get redirect bad request response
func (o *GetRedirectBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRedirectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRedirectUnauthorizedCode is the HTTP code returned for type GetRedirectUnauthorized
const GetRedirectUnauthorizedCode int = 401

/*GetRedirectUnauthorized Unauthorized

swagger:response getRedirectUnauthorized
*/
type GetRedirectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRedirectUnauthorized creates GetRedirectUnauthorized with default headers values
func NewGetRedirectUnauthorized() *GetRedirectUnauthorized {

	return &GetRedirectUnauthorized{}
}

// WithPayload adds the payload to the get redirect unauthorized response
func (o *GetRedirectUnauthorized) WithPayload(payload *models.Error) *GetRedirectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get redirect unauthorized response
func (o *GetRedirectUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRedirectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRedirectNotFoundCode is the HTTP code returned for type GetRedirectNotFound
const GetRedirectNotFoundCode int = 404

/*GetRedirectNotFound The specified resource was not found

swagger:response getRedirectNotFound
*/
type GetRedirectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRedirectNotFound creates GetRedirectNotFound with default headers values
func NewGetRedirectNotFound() *GetRedirectNotFound {

	return &GetRedirectNotFound{}
}

// WithPayload adds the payload to the get redirect not found response
func (o *GetRedirectNotFound) WithPayload(payload *models.Error) *GetRedirectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get redirect not found response
func (o *GetRedirectNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRedirectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
