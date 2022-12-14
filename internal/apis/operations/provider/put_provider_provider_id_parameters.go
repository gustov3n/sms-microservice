// Code generated by go-swagger; DO NOT EDIT.

package provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPutProviderProviderIDParams creates a new PutProviderProviderIDParams object
//
// There are no default values defined in the spec.
func NewPutProviderProviderIDParams() PutProviderProviderIDParams {

	return PutProviderProviderIDParams{}
}

// PutProviderProviderIDParams contains all the bound params for the put provider provider ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutProviderProviderID
type PutProviderProviderIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body PutProviderProviderIDBody
	/*
	  Required: true
	  In: path
	*/
	ProviderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutProviderProviderIDParams() beforehand.
func (o *PutProviderProviderIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PutProviderProviderIDBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	}

	rProviderID, rhkProviderID, _ := route.Params.GetOK("provider_id")
	if err := o.bindProviderID(rProviderID, rhkProviderID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProviderID binds and validates parameter ProviderID from path.
func (o *PutProviderProviderIDParams) bindProviderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ProviderID = raw

	return nil
}
