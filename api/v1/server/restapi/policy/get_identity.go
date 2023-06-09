// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetIdentityHandlerFunc turns a function with the right signature into a get identity handler
type GetIdentityHandlerFunc func(GetIdentityParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIdentityHandlerFunc) Handle(params GetIdentityParams) middleware.Responder {
	return fn(params)
}

// GetIdentityHandler interface for that can handle valid get identity params
type GetIdentityHandler interface {
	Handle(GetIdentityParams) middleware.Responder
}

// NewGetIdentity creates a new http.Handler for the get identity operation
func NewGetIdentity(ctx *middleware.Context, handler GetIdentityHandler) *GetIdentity {
	return &GetIdentity{Context: ctx, Handler: handler}
}

/*
	GetIdentity swagger:route GET /identity policy getIdentity

Retrieves a list of identities that have metadata matching the provided parameters.

Retrieves a list of identities that have metadata matching the provided parameters, or all identities if no parameters are provided.
*/
type GetIdentity struct {
	Context *middleware.Context
	Handler GetIdentityHandler
}

func (o *GetIdentity) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetIdentityParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
