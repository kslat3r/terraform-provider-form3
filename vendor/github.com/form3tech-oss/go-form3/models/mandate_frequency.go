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

// MandateFrequency mandate frequency
// swagger:model MandateFrequency
type MandateFrequency string

const (

	// MandateFrequencyDaily captures enum value "daily"
	MandateFrequencyDaily MandateFrequency = "daily"

	// MandateFrequencyWeekly captures enum value "weekly"
	MandateFrequencyWeekly MandateFrequency = "weekly"

	// MandateFrequencyFortnightly captures enum value "fortnightly"
	MandateFrequencyFortnightly MandateFrequency = "fortnightly"

	// MandateFrequencyMonthly captures enum value "monthly"
	MandateFrequencyMonthly MandateFrequency = "monthly"

	// MandateFrequencyBimonthly captures enum value "bimonthly"
	MandateFrequencyBimonthly MandateFrequency = "bimonthly"

	// MandateFrequencyQuarterly captures enum value "quarterly"
	MandateFrequencyQuarterly MandateFrequency = "quarterly"

	// MandateFrequencyYearly captures enum value "yearly"
	MandateFrequencyYearly MandateFrequency = "yearly"
)

// for schema
var mandateFrequencyEnum []interface{}

func init() {
	var res []MandateFrequency
	if err := json.Unmarshal([]byte(`["daily","weekly","fortnightly","monthly","bimonthly","quarterly","yearly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mandateFrequencyEnum = append(mandateFrequencyEnum, v)
	}
}

func (m MandateFrequency) validateMandateFrequencyEnum(path, location string, value MandateFrequency) error {
	if err := validate.Enum(path, location, value, mandateFrequencyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mandate frequency
func (m MandateFrequency) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMandateFrequencyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
