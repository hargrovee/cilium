// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutEndpointIDHandlerFunc turns a function with the right signature into a put endpoint ID handler
type PutEndpointIDHandlerFunc func(PutEndpointIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutEndpointIDHandlerFunc) Handle(params PutEndpointIDParams) middleware.Responder {
	return fn(params)
}

// PutEndpointIDHandler interface for that can handle valid put endpoint ID params
type PutEndpointIDHandler interface {
	Handle(PutEndpointIDParams) middleware.Responder
}

// NewPutEndpointID creates a new http.Handler for the put endpoint ID operation
func NewPutEndpointID(ctx *middleware.Context, handler PutEndpointIDHandler) *PutEndpointID {
	return &PutEndpointID{Context: ctx, Handler: handler}
}

/*
	PutEndpointID swagger:route PUT /endpoint/{id} endpoint putEndpointId

# Create endpoint

Creates a new endpoint
*/
type PutEndpointID struct {
	Context *middleware.Context
	Handler PutEndpointIDHandler
}

func (o *PutEndpointID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutEndpointIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
