// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hanoman-id/xendit-payment/internal/models"
)

// PostCreateOtpOKCode is the HTTP code returned for type PostCreateOtpOK
const PostCreateOtpOKCode int = 200

/*PostCreateOtpOK Success

swagger:response postCreateOtpOK
*/
type PostCreateOtpOK struct {

	/*
	  In: Body
	*/
	Payload *PostCreateOtpOKBody `json:"body,omitempty"`
}

// NewPostCreateOtpOK creates PostCreateOtpOK with default headers values
func NewPostCreateOtpOK() *PostCreateOtpOK {

	return &PostCreateOtpOK{}
}

// WithPayload adds the payload to the post create otp o k response
func (o *PostCreateOtpOK) WithPayload(payload *PostCreateOtpOKBody) *PostCreateOtpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create otp o k response
func (o *PostCreateOtpOK) SetPayload(payload *PostCreateOtpOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateOtpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCreateOtpBadRequestCode is the HTTP code returned for type PostCreateOtpBadRequest
const PostCreateOtpBadRequestCode int = 400

/*PostCreateOtpBadRequest Unauthorized

swagger:response postCreateOtpBadRequest
*/
type PostCreateOtpBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostCreateOtpBadRequest creates PostCreateOtpBadRequest with default headers values
func NewPostCreateOtpBadRequest() *PostCreateOtpBadRequest {

	return &PostCreateOtpBadRequest{}
}

// WithPayload adds the payload to the post create otp bad request response
func (o *PostCreateOtpBadRequest) WithPayload(payload *models.Error) *PostCreateOtpBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create otp bad request response
func (o *PostCreateOtpBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateOtpBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCreateOtpUnauthorizedCode is the HTTP code returned for type PostCreateOtpUnauthorized
const PostCreateOtpUnauthorizedCode int = 401

/*PostCreateOtpUnauthorized Unauthorized

swagger:response postCreateOtpUnauthorized
*/
type PostCreateOtpUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostCreateOtpUnauthorized creates PostCreateOtpUnauthorized with default headers values
func NewPostCreateOtpUnauthorized() *PostCreateOtpUnauthorized {

	return &PostCreateOtpUnauthorized{}
}

// WithPayload adds the payload to the post create otp unauthorized response
func (o *PostCreateOtpUnauthorized) WithPayload(payload *models.Error) *PostCreateOtpUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create otp unauthorized response
func (o *PostCreateOtpUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateOtpUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCreateOtpNotFoundCode is the HTTP code returned for type PostCreateOtpNotFound
const PostCreateOtpNotFoundCode int = 404

/*PostCreateOtpNotFound The specified resource was not found

swagger:response postCreateOtpNotFound
*/
type PostCreateOtpNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostCreateOtpNotFound creates PostCreateOtpNotFound with default headers values
func NewPostCreateOtpNotFound() *PostCreateOtpNotFound {

	return &PostCreateOtpNotFound{}
}

// WithPayload adds the payload to the post create otp not found response
func (o *PostCreateOtpNotFound) WithPayload(payload *models.Error) *PostCreateOtpNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create otp not found response
func (o *PostCreateOtpNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateOtpNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
