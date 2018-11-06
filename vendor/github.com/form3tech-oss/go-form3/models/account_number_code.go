// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AccountNumberCode account number code
// swagger:model AccountNumberCode
type AccountNumberCode string

const (
	// AccountNumberCodeIBAN captures enum value "IBAN"
	AccountNumberCodeIBAN AccountNumberCode = "IBAN"
	// AccountNumberCodeBBAN captures enum value "BBAN"
	AccountNumberCodeBBAN AccountNumberCode = "BBAN"
)

// for schema
var accountNumberCodeEnum []interface{}

func init() {
	var res []AccountNumberCode
	if err := json.Unmarshal([]byte(`["IBAN","BBAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountNumberCodeEnum = append(accountNumberCodeEnum, v)
	}
}

func (m AccountNumberCode) validateAccountNumberCodeEnum(path, location string, value AccountNumberCode) error {
	if err := validate.Enum(path, location, value, accountNumberCodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account number code
func (m AccountNumberCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountNumberCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
