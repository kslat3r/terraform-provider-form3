// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PaymentSubmissionRelationships payment submission relationships
// swagger:model PaymentSubmissionRelationships
type PaymentSubmissionRelationships struct {

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// validations
	Validations *RelationshipLinks `json:"validations,omitempty"`
}

// Validate validates this payment submission relationships
func (m *PaymentSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentSubmissionRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentSubmissionRelationships) validateValidations(formats strfmt.Registry) error {

	if swag.IsZero(m.Validations) { // not required
		return nil
	}

	if m.Validations != nil {
		if err := m.Validations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
