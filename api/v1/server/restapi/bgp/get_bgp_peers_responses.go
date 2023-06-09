// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetBgpPeersOKCode is the HTTP code returned for type GetBgpPeersOK
const GetBgpPeersOKCode int = 200

/*
GetBgpPeersOK Success

swagger:response getBgpPeersOK
*/
type GetBgpPeersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.BgpPeer `json:"body,omitempty"`
}

// NewGetBgpPeersOK creates GetBgpPeersOK with default headers values
func NewGetBgpPeersOK() *GetBgpPeersOK {

	return &GetBgpPeersOK{}
}

// WithPayload adds the payload to the get bgp peers o k response
func (o *GetBgpPeersOK) WithPayload(payload []*models.BgpPeer) *GetBgpPeersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bgp peers o k response
func (o *GetBgpPeersOK) SetPayload(payload []*models.BgpPeer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBgpPeersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BgpPeer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
