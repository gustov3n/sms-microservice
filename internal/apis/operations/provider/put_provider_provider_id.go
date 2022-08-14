// Code generated by go-swagger; DO NOT EDIT.

package provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutProviderProviderIDHandlerFunc turns a function with the right signature into a put provider provider ID handler
type PutProviderProviderIDHandlerFunc func(PutProviderProviderIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutProviderProviderIDHandlerFunc) Handle(params PutProviderProviderIDParams) middleware.Responder {
	return fn(params)
}

// PutProviderProviderIDHandler interface for that can handle valid put provider provider ID params
type PutProviderProviderIDHandler interface {
	Handle(PutProviderProviderIDParams) middleware.Responder
}

// NewPutProviderProviderID creates a new http.Handler for the put provider provider ID operation
func NewPutProviderProviderID(ctx *middleware.Context, handler PutProviderProviderIDHandler) *PutProviderProviderID {
	return &PutProviderProviderID{Context: ctx, Handler: handler}
}

/* PutProviderProviderID swagger:route PUT /provider/{provider_id} Provider putProviderProviderId

Health check

Health check endpoint

*/
type PutProviderProviderID struct {
	Context *middleware.Context
	Handler PutProviderProviderIDHandler
}

func (o *PutProviderProviderID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutProviderProviderIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutProviderProviderIDBody put provider provider ID body
//
// swagger:model PutProviderProviderIDBody
type PutProviderProviderIDBody struct {

	// password
	Password string `json:"password,omitempty"`

	// provider name
	ProviderName string `json:"provider_name,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this put provider provider ID body
func (o *PutProviderProviderIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put provider provider ID body based on context it is used
func (o *PutProviderProviderIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutProviderProviderIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutProviderProviderIDBody) UnmarshalBinary(b []byte) error {
	var res PutProviderProviderIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutProviderProviderIDOKBody put provider provider ID o k body
//
// swagger:model PutProviderProviderIDOKBody
type PutProviderProviderIDOKBody struct {

	// response code
	ResponseCode string `json:"response_code,omitempty"`

	// response data
	ResponseData string `json:"response_data,omitempty"`

	// response messege
	ResponseMessege string `json:"response_messege,omitempty"`
}

// Validate validates this put provider provider ID o k body
func (o *PutProviderProviderIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put provider provider ID o k body based on context it is used
func (o *PutProviderProviderIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutProviderProviderIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutProviderProviderIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutProviderProviderIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
