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

// PatchEndpointIDLabelsHandlerFunc turns a function with the right signature into a patch endpoint ID labels handler
type PatchEndpointIDLabelsHandlerFunc func(PatchEndpointIDLabelsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchEndpointIDLabelsHandlerFunc) Handle(params PatchEndpointIDLabelsParams) middleware.Responder {
	return fn(params)
}

// PatchEndpointIDLabelsHandler interface for that can handle valid patch endpoint ID labels params
type PatchEndpointIDLabelsHandler interface {
	Handle(PatchEndpointIDLabelsParams) middleware.Responder
}

// NewPatchEndpointIDLabels creates a new http.Handler for the patch endpoint ID labels operation
func NewPatchEndpointIDLabels(ctx *middleware.Context, handler PatchEndpointIDLabelsHandler) *PatchEndpointIDLabels {
	return &PatchEndpointIDLabels{Context: ctx, Handler: handler}
}

/*
	PatchEndpointIDLabels swagger:route PATCH /endpoint/{id}/labels endpoint patchEndpointIdLabels

# Set label configuration of endpoint

Sets labels associated with an endpoint. These can be user provided or
derived from the orchestration system.
*/
type PatchEndpointIDLabels struct {
	Context *middleware.Context
	Handler PatchEndpointIDLabelsHandler
}

func (o *PatchEndpointIDLabels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchEndpointIDLabelsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
