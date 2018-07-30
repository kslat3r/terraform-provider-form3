// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Links links
// swagger:model Links
type Links struct {

	// first
	First string `json:"first,omitempty"`

	// last
	Last string `json:"last,omitempty"`

	// next
	Next string `json:"next,omitempty"`

	// prev
	Prev string `json:"prev,omitempty"`

	// self
	Self string `json:"self,omitempty"`
}

// Validate validates this links
func (m *Links) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Links) UnmarshalBinary(b []byte) error {
	var res Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
