// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SepaReconciliationNewAssociation sepa reconciliation new association
// swagger:model SepaReconciliationNewAssociation
type SepaReconciliationNewAssociation struct {

	// attributes
	// Required: true
	Attributes *SepaReconciliationAssociationAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *SepaReconciliationAssociationRelationships `json:"relationships,omitempty"`

	// type
	// Required: true
	// Enum: [separeconciliation_associations]
	Type string `json:"type"`
}

// Validate validates this sepa reconciliation new association
func (m *SepaReconciliationNewAssociation) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaReconciliationNewAssociation) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
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

func (m *SepaReconciliationNewAssociation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationNewAssociation) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SepaReconciliationNewAssociation) validateRelationships(formats strfmt.Registry) error {

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

var sepaReconciliationNewAssociationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["separeconciliation_associations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sepaReconciliationNewAssociationTypeTypePropEnum = append(sepaReconciliationNewAssociationTypeTypePropEnum, v)
	}
}

const (

	// SepaReconciliationNewAssociationTypeSepareconciliationAssociations captures enum value "separeconciliation_associations"
	SepaReconciliationNewAssociationTypeSepareconciliationAssociations string = "separeconciliation_associations"
)

// prop value enum
func (m *SepaReconciliationNewAssociation) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sepaReconciliationNewAssociationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SepaReconciliationNewAssociation) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaReconciliationNewAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaReconciliationNewAssociation) UnmarshalBinary(b []byte) error {
	var res SepaReconciliationNewAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
