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

// ReturnSubmission return submission
// swagger:model ReturnSubmission
type ReturnSubmission struct {

	// attributes
	Attributes *ReturnSubmissionAttributes `json:"attributes,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ReturnSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this return submission
func (m *ReturnSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmission) UnmarshalBinary(b []byte) error {
	var res ReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ReturnSubmissionAttributes return submission attributes
// swagger:model ReturnSubmissionAttributes
type ReturnSubmissionAttributes struct {

	// limit breach end datetime
	// Read Only: true
	// Format: date-time
	LimitBreachEndDatetime strfmt.DateTime `json:"limit_breach_end_datetime,omitempty"`

	// limit breach start datetime
	// Read Only: true
	// Format: date-time
	LimitBreachStartDatetime strfmt.DateTime `json:"limit_breach_start_datetime,omitempty"`

	// redirected account number
	RedirectedAccountNumber string `json:"redirected_account_number,omitempty"`

	// redirected bank id
	RedirectedBankID string `json:"redirected_bank_id,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// settlement cycle
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// settlement date
	// Format: date
	SettlementDate strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status ReturnSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// transaction start datetime
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

// Validate validates this return submission attributes
func (m *ReturnSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimitBreachEndDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionAttributes) validateLimitBreachEndDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachEndDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_end_datetime", "body", "date-time", m.LimitBreachEndDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateLimitBreachStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_start_datetime", "body", "date-time", m.LimitBreachStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
