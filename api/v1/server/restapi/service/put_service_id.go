// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutServiceIDHandlerFunc turns a function with the right signature into a put service ID handler
type PutServiceIDHandlerFunc func(PutServiceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutServiceIDHandlerFunc) Handle(params PutServiceIDParams) middleware.Responder {
	return fn(params)
}

// PutServiceIDHandler interface for that can handle valid put service ID params
type PutServiceIDHandler interface {
	Handle(PutServiceIDParams) middleware.Responder
}

// NewPutServiceID creates a new http.Handler for the put service ID operation
func NewPutServiceID(ctx *middleware.Context, handler PutServiceIDHandler) *PutServiceID {
	return &PutServiceID{Context: ctx, Handler: handler}
}

/*
	PutServiceID swagger:route PUT /service/{id} service putServiceId

Create or update service
*/
type PutServiceID struct {
	Context *middleware.Context
	Handler PutServiceIDHandler
}

func (o *PutServiceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutServiceIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
