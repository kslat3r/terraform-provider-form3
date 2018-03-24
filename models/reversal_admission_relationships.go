// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReversalAdmissionRelationships reversal admission relationships
// swagger:model reversalAdmissionRelationships
type ReversalAdmissionRelationships struct {

	// reversal
	Reversal *RelationshipLinks `json:"reversal,omitempty"`
}

// Validate validates this reversal admission relationships
func (m *ReversalAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReversal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionRelationships) validateReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.Reversal) { // not required
		return nil
	}

	if m.Reversal != nil {

		if err := m.Reversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reversal")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
