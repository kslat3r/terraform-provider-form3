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

// ApprovalRequestAttributes approval request attributes
// swagger:model approvalRequestAttributes
type ApprovalRequestAttributes struct {

	// action
	// Pattern: ^[A-Za-z]*$
	Action string `json:"action,omitempty"`

	// action time
	ActionTime strfmt.DateTime `json:"action_time,omitempty"`

	// actioned by
	ActionedBy strfmt.UUID `json:"actioned_by,omitempty"`

	// after data
	AfterData interface{} `json:"after_data,omitempty"`

	// before data
	BeforeData interface{} `json:"before_data,omitempty"`

	// record id
	RecordID strfmt.UUID `json:"record_id,omitempty"`

	// record orgid
	RecordOrgid strfmt.UUID `json:"record_orgid,omitempty"`

	// record type
	// Pattern: ^[A-Za-z]*$
	RecordType string `json:"record_type,omitempty"`

	// record version
	// Minimum: 0
	RecordVersion *int64 `json:"record_version,omitempty"`

	// status
	// Pattern: ^[A-Za-z]*$
	Status string `json:"status,omitempty"`
}

// Validate validates this approval request attributes
func (m *ApprovalRequestAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActionTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActionedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordOrgid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApprovalRequestAttributes) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := validate.Pattern("action", "body", string(m.Action), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateActionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("action_time", "body", "date-time", m.ActionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateActionedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("actioned_by", "body", "uuid", m.ActionedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordID(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordID) { // not required
		return nil
	}

	if err := validate.FormatOf("record_id", "body", "uuid", m.RecordID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordOrgid(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordOrgid) { // not required
		return nil
	}

	if err := validate.FormatOf("record_orgid", "body", "uuid", m.RecordOrgid.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordType) { // not required
		return nil
	}

	if err := validate.Pattern("record_type", "body", string(m.RecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateRecordVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordVersion) { // not required
		return nil
	}

	if err := validate.MinimumInt("record_version", "body", int64(*m.RecordVersion), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalRequestAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.Pattern("status", "body", string(m.Status), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalRequestAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalRequestAttributes) UnmarshalBinary(b []byte) error {
	var res ApprovalRequestAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
