// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPListEntryMetadata Additional metadata assigned to an IP list entry
//
// swagger:model IPListEntryMetadata
type IPListEntryMetadata struct {

	// Name assigned to the IP (e.g. Kubernetes pod name)
	Name string `json:"name,omitempty"`

	// Namespace of the IP (e.g. Kubernetes namespace)
	Namespace string `json:"namespace,omitempty"`

	// Source of the IP entry and its metadata
	// Example: k8s
	Source string `json:"source,omitempty"`
}

// Validate validates this IP list entry metadata
func (m *IPListEntryMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP list entry metadata based on context it is used
func (m *IPListEntryMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPListEntryMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPListEntryMetadata) UnmarshalBinary(b []byte) error {
	var res IPListEntryMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
