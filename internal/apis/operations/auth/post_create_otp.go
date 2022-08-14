// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostCreateOtpHandlerFunc turns a function with the right signature into a post create otp handler
type PostCreateOtpHandlerFunc func(PostCreateOtpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCreateOtpHandlerFunc) Handle(params PostCreateOtpParams) middleware.Responder {
	return fn(params)
}

// PostCreateOtpHandler interface for that can handle valid post create otp params
type PostCreateOtpHandler interface {
	Handle(PostCreateOtpParams) middleware.Responder
}

// NewPostCreateOtp creates a new http.Handler for the post create otp operation
func NewPostCreateOtp(ctx *middleware.Context, handler PostCreateOtpHandler) *PostCreateOtp {
	return &PostCreateOtp{Context: ctx, Handler: handler}
}

/* PostCreateOtp swagger:route POST /create-otp Auth postCreateOtp

Add a new pet to the store

Health check endpoint

*/
type PostCreateOtp struct {
	Context *middleware.Context
	Handler PostCreateOtpHandler
}

func (o *PostCreateOtp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCreateOtpParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCreateOtpBody post create otp body
//
// swagger:model PostCreateOtpBody
type PostCreateOtpBody struct {

	// user phone_number
	// Example: 081334702936
	PhoneNumber string `json:"phone_number,omitempty"`
}

// Validate validates this post create otp body
func (o *PostCreateOtpBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post create otp body based on context it is used
func (o *PostCreateOtpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCreateOtpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCreateOtpBody) UnmarshalBinary(b []byte) error {
	var res PostCreateOtpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCreateOtpOKBody post create otp o k body
//
// swagger:model PostCreateOtpOKBody
type PostCreateOtpOKBody struct {

	// data
	Data *PostCreateOtpOKBodyData `json:"data,omitempty"`

	// success retrieve data
	// Example: Success
	Message string `json:"message,omitempty"`
}

// Validate validates this post create otp o k body
func (o *PostCreateOtpOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCreateOtpOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCreateOtpOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCreateOtpOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post create otp o k body based on the context it is used
func (o *PostCreateOtpOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCreateOtpOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCreateOtpOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCreateOtpOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCreateOtpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCreateOtpOKBody) UnmarshalBinary(b []byte) error {
	var res PostCreateOtpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCreateOtpOKBodyData post create otp o k body data
//
// swagger:model PostCreateOtpOKBodyData
type PostCreateOtpOKBodyData struct {

	// user otp
	// Example: 123456
	Otp string `json:"otp,omitempty"`
}

// Validate validates this post create otp o k body data
func (o *PostCreateOtpOKBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post create otp o k body data based on context it is used
func (o *PostCreateOtpOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCreateOtpOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCreateOtpOKBodyData) UnmarshalBinary(b []byte) error {
	var res PostCreateOtpOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}