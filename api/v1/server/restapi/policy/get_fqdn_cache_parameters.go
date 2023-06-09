// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetFqdnCacheParams creates a new GetFqdnCacheParams object
//
// There are no default values defined in the spec.
func NewGetFqdnCacheParams() GetFqdnCacheParams {

	return GetFqdnCacheParams{}
}

// GetFqdnCacheParams contains all the bound params for the get fqdn cache operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFqdnCache
type GetFqdnCacheParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A CIDR range of IPs
	  In: query
	*/
	Cidr *string
	/*A toFQDNs compatible matchPattern expression
	  In: query
	*/
	Matchpattern *string
	/*Source from which FQDN entries come from
	  In: query
	*/
	Source *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFqdnCacheParams() beforehand.
func (o *GetFqdnCacheParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCidr, qhkCidr, _ := qs.GetOK("cidr")
	if err := o.bindCidr(qCidr, qhkCidr, route.Formats); err != nil {
		res = append(res, err)
	}

	qMatchpattern, qhkMatchpattern, _ := qs.GetOK("matchpattern")
	if err := o.bindMatchpattern(qMatchpattern, qhkMatchpattern, route.Formats); err != nil {
		res = append(res, err)
	}

	qSource, qhkSource, _ := qs.GetOK("source")
	if err := o.bindSource(qSource, qhkSource, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCidr binds and validates parameter Cidr from query.
func (o *GetFqdnCacheParams) bindCidr(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Cidr = &raw

	return nil
}

// bindMatchpattern binds and validates parameter Matchpattern from query.
func (o *GetFqdnCacheParams) bindMatchpattern(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Matchpattern = &raw

	return nil
}

// bindSource binds and validates parameter Source from query.
func (o *GetFqdnCacheParams) bindSource(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Source = &raw

	return nil
}
