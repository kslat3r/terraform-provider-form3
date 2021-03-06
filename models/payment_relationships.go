// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PaymentRelationships payment relationships
// swagger:model PaymentRelationships
type PaymentRelationships struct {

	// payment admission
	PaymentAdmission *PaymentRelationshipsPaymentAdmission `json:"payment_admission,omitempty"`

	// payment return
	PaymentReturn *PaymentRelationshipsPaymentReturn `json:"payment_return,omitempty"`

	// payment reversal
	PaymentReversal *PaymentRelationshipsPaymentReversal `json:"payment_reversal,omitempty"`

	// payment submission
	PaymentSubmission *PaymentRelationshipsPaymentSubmission `json:"payment_submission,omitempty"`
}

// Validate validates this payment relationships
func (m *PaymentRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationships) validatePaymentAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdmission) { // not required
		return nil
	}

	if m.PaymentAdmission != nil {
		if err := m.PaymentAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_admission")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturn) { // not required
		return nil
	}

	if m.PaymentReturn != nil {
		if err := m.PaymentReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReversal) { // not required
		return nil
	}

	if m.PaymentReversal != nil {
		if err := m.PaymentReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentSubmission) { // not required
		return nil
	}

	if m.PaymentSubmission != nil {
		if err := m.PaymentSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentRelationshipsPaymentAdmission The payment admission resource related to the payment
// swagger:model PaymentRelationshipsPaymentAdmission
type PaymentRelationshipsPaymentAdmission struct {

	// data
	Data []*PaymentAdmission `json:"data"`
}

// Validate validates this payment relationships payment admission
func (m *PaymentRelationshipsPaymentAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentAdmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentAdmission) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentRelationshipsPaymentReturn The payment return resource related to the payment
// swagger:model PaymentRelationshipsPaymentReturn
type PaymentRelationshipsPaymentReturn struct {

	// data
	Data []*ReturnPayment `json:"data"`
}

// Validate validates this payment relationships payment return
func (m *PaymentRelationshipsPaymentReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentReturn) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReturn) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentRelationshipsPaymentReversal The payment reversal resource related to the payment
// swagger:model PaymentRelationshipsPaymentReversal
type PaymentRelationshipsPaymentReversal struct {

	// data
	Data []*ReversalPayment `json:"data"`
}

// Validate validates this payment relationships payment reversal
func (m *PaymentRelationshipsPaymentReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentReversal) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReversal) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentRelationshipsPaymentSubmission The payment submission resource related to the payment
// swagger:model PaymentRelationshipsPaymentSubmission
type PaymentRelationshipsPaymentSubmission struct {

	// data
	Data []*PaymentSubmission `json:"data"`
}

// Validate validates this payment relationships payment submission
func (m *PaymentRelationshipsPaymentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentSubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentSubmission) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
