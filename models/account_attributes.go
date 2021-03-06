// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAttributes account attributes
// swagger:model AccountAttributes
type AccountAttributes struct {

	// account number
	// Pattern: ^[A-Z0-9]{0,64}$
	AccountNumber string `json:"account_number,omitempty"`

	// bank id
	// Pattern: ^[A-Z0-9]{0,16}$
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	// Pattern: ^[A-Z]{0,16}$
	BankIDCode string `json:"bank_id_code,omitempty"`

	// base currency
	// Pattern: ^[A-Z]{3}$
	BaseCurrency string `json:"base_currency,omitempty"`

	// bic
	// Pattern: ^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$
	Bic string `json:"bic,omitempty"`

	// country
	// Required: true
	// Pattern: ^[A-Z]{2}$
	Country *string `json:"country"`

	// customer id
	// Pattern: ^[a-zA-Z0-9-$@., ]{0,256}$
	CustomerID string `json:"customer_id,omitempty"`

	// iban
	// Pattern: ^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$
	Iban string `json:"iban,omitempty"`
}

// Validate validates this account attributes
func (m *AccountAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIban(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAttributes) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBankID(formats strfmt.Registry) error {

	if swag.IsZero(m.BankID) { // not required
		return nil
	}

	if err := validate.Pattern("bank_id", "body", string(m.BankID), `^[A-Z0-9]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := validate.Pattern("bank_id_code", "body", string(m.BankIDCode), `^[A-Z]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBaseCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseCurrency) { // not required
		return nil
	}

	if err := validate.Pattern("base_currency", "body", string(m.BaseCurrency), `^[A-Z]{3}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBic(formats strfmt.Registry) error {

	if swag.IsZero(m.Bic) { // not required
		return nil
	}

	if err := validate.Pattern("bic", "body", string(m.Bic), `^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.Pattern("country", "body", string(*m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.Pattern("customer_id", "body", string(m.CustomerID), `^[a-zA-Z0-9-$@., ]{0,256}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateIban(formats strfmt.Registry) error {

	if swag.IsZero(m.Iban) { // not required
		return nil
	}

	if err := validate.Pattern("iban", "body", string(m.Iban), `^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributes) UnmarshalBinary(b []byte) error {
	var res AccountAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
