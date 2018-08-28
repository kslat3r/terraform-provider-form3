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

// Payment payment
// swagger:model Payment
type Payment struct {

	// attributes
	// Required: true
	Attributes *PaymentAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *PaymentRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this payment
func (m *Payment) Validate(formats strfmt.Registry) error {
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

func (m *Payment) validateAttributes(formats strfmt.Registry) error {

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

func (m *Payment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateRelationships(formats strfmt.Registry) error {

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

func (m *Payment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Payment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Payment) UnmarshalBinary(b []byte) error {
	var res Payment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributes payment attributes
// swagger:model PaymentAttributes
type PaymentAttributes struct {

	// amount
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// batch booking indicator
	BatchBookingIndicator string `json:"batch_booking_indicator,omitempty"`

	// batch id
	BatchID string `json:"batch_id,omitempty"`

	// batch type
	BatchType string `json:"batch_type,omitempty"`

	// beneficiary party
	BeneficiaryParty *PaymentAttributesBeneficiaryParty `json:"beneficiary_party,omitempty"`

	// category purpose
	CategoryPurpose string `json:"category_purpose,omitempty"`

	// category purpose coded
	CategoryPurposeCoded string `json:"category_purpose_coded,omitempty"`

	// charges information
	ChargesInformation *ChargesInformation `json:"charges_information,omitempty"`

	// clearing id
	ClearingID string `json:"clearing_id,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// debtor party
	DebtorParty *PaymentAttributesDebtorParty `json:"debtor_party,omitempty"`

	// end to end reference
	EndToEndReference string `json:"end_to_end_reference,omitempty"`

	// file number
	FileNumber string `json:"file_number,omitempty"`

	// fx
	Fx *PaymentAttributesFx `json:"fx,omitempty"`

	// intermediary bank
	IntermediaryBank *AccountHoldingEntity `json:"intermediary_bank,omitempty"`

	// numeric reference
	NumericReference string `json:"numeric_reference,omitempty"`

	// payment acceptance datetime
	// Format: date-time
	PaymentAcceptanceDatetime *strfmt.DateTime `json:"payment_acceptance_datetime,omitempty"`

	// payment instructing id
	PaymentInstructingID string `json:"payment_instructing_id,omitempty"`

	// payment purpose
	PaymentPurpose string `json:"payment_purpose,omitempty"`

	// payment purpose coded
	PaymentPurposeCoded string `json:"payment_purpose_coded,omitempty"`

	// payment scheme
	PaymentScheme string `json:"payment_scheme,omitempty"`

	// payment type
	PaymentType string `json:"payment_type,omitempty"`

	// processing date
	// Format: date
	ProcessingDate strfmt.Date `json:"processing_date,omitempty"`

	// receivers correspondent
	ReceiversCorrespondent *AccountHoldingEntity `json:"receivers_correspondent,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`

	// regulatory reporting
	RegulatoryReporting string `json:"regulatory_reporting,omitempty"`

	// reimbursement
	Reimbursement *AccountHoldingEntity `json:"reimbursement,omitempty"`

	// remittance information
	RemittanceInformation string `json:"remittance_information,omitempty"`

	// scheme payment sub type
	SchemePaymentSubType string `json:"scheme_payment_sub_type,omitempty"`

	// scheme payment type
	SchemePaymentType string `json:"scheme_payment_type,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme transaction id
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`

	// senders correspondent
	SendersCorrespondent *AccountHoldingEntity `json:"senders_correspondent,omitempty"`

	// structured reference
	StructuredReference *PaymentAttributesStructuredReference `json:"structured_reference,omitempty"`

	// swift
	Swift *PaymentAttributesSwift `json:"swift,omitempty"`

	// ultimate beneficiary
	UltimateBeneficiary *UltimateEntity `json:"ultimate_beneficiary,omitempty"`

	// ultimate debtor
	UltimateDebtor *UltimateEntity `json:"ultimate_debtor,omitempty"`

	// unique scheme id
	UniqueSchemeID string `json:"unique_scheme_id,omitempty"`
}

// Validate validates this payment attributes
func (m *PaymentAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargesInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntermediaryBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAcceptanceDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiversCorrespondent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReimbursement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendersCorrespondent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStructuredReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwift(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUltimateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUltimateDebtor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributes) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryParty) { // not required
		return nil
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateChargesInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargesInformation) { // not required
		return nil
	}

	if m.ChargesInformation != nil {
		if err := m.ChargesInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "charges_information")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateDebtorParty(formats strfmt.Registry) error {

	if swag.IsZero(m.DebtorParty) { // not required
		return nil
	}

	if m.DebtorParty != nil {
		if err := m.DebtorParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "debtor_party")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateFx(formats strfmt.Registry) error {

	if swag.IsZero(m.Fx) { // not required
		return nil
	}

	if m.Fx != nil {
		if err := m.Fx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "fx")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateIntermediaryBank(formats strfmt.Registry) error {

	if swag.IsZero(m.IntermediaryBank) { // not required
		return nil
	}

	if m.IntermediaryBank != nil {
		if err := m.IntermediaryBank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "intermediary_bank")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validatePaymentAcceptanceDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAcceptanceDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"payment_acceptance_datetime", "body", "date-time", m.PaymentAcceptanceDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"processing_date", "body", "date", m.ProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateReceiversCorrespondent(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceiversCorrespondent) { // not required
		return nil
	}

	if m.ReceiversCorrespondent != nil {
		if err := m.ReceiversCorrespondent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "receivers_correspondent")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateReimbursement(formats strfmt.Registry) error {

	if swag.IsZero(m.Reimbursement) { // not required
		return nil
	}

	if m.Reimbursement != nil {
		if err := m.Reimbursement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "reimbursement")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateSendersCorrespondent(formats strfmt.Registry) error {

	if swag.IsZero(m.SendersCorrespondent) { // not required
		return nil
	}

	if m.SendersCorrespondent != nil {
		if err := m.SendersCorrespondent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "senders_correspondent")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateStructuredReference(formats strfmt.Registry) error {

	if swag.IsZero(m.StructuredReference) { // not required
		return nil
	}

	if m.StructuredReference != nil {
		if err := m.StructuredReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "structured_reference")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateSwift(formats strfmt.Registry) error {

	if swag.IsZero(m.Swift) { // not required
		return nil
	}

	if m.Swift != nil {
		if err := m.Swift.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "swift")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateUltimateBeneficiary(formats strfmt.Registry) error {

	if swag.IsZero(m.UltimateBeneficiary) { // not required
		return nil
	}

	if m.UltimateBeneficiary != nil {
		if err := m.UltimateBeneficiary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "ultimate_beneficiary")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateUltimateDebtor(formats strfmt.Registry) error {

	if swag.IsZero(m.UltimateDebtor) { // not required
		return nil
	}

	if m.UltimateDebtor != nil {
		if err := m.UltimateDebtor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "ultimate_debtor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesBeneficiaryParty payment attributes beneficiary party
// swagger:model PaymentAttributesBeneficiaryParty
type PaymentAttributesBeneficiaryParty struct {

	// account name
	AccountName string `json:"account_name,omitempty"`

	// account number
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account type
	AccountType int64 `json:"account_type,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// address
	Address []string `json:"address"`

	// birth city
	BirthCity string `json:"birth_city,omitempty"`

	// birth country
	BirthCountry string `json:"birth_country,omitempty"`

	// birth date
	// Format: date
	BirthDate strfmt.Date `json:"birth_date,omitempty"`

	// birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organisation identification
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// organisation identification code
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// organisation identification issuer
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	// telephone number
	TelephoneNumber string `json:"telephone_number,omitempty"`
}

// Validate validates this payment attributes beneficiary party
func (m *PaymentAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"beneficiary_party"+"."+"account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributesBeneficiaryParty) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"beneficiary_party"+"."+"birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesDebtorParty payment attributes debtor party
// swagger:model PaymentAttributesDebtorParty
type PaymentAttributesDebtorParty struct {

	// account name
	AccountName string `json:"account_name,omitempty"`

	// account number
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// address
	Address []string `json:"address"`

	// birth city
	BirthCity string `json:"birth_city,omitempty"`

	// birth country
	BirthCountry string `json:"birth_country,omitempty"`

	// birth date
	// Format: date
	BirthDate strfmt.Date `json:"birth_date,omitempty"`

	// birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// customer id code
	CustomerIDCode string `json:"customer_id_code,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organisation identification
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// organisation identification code
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// organisation identification issuer
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`
}

// Validate validates this payment attributes debtor party
func (m *PaymentAttributesDebtorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributesDebtorParty) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"debtor_party"+"."+"account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributesDebtorParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "debtor_party" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *PaymentAttributesDebtorParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "debtor_party" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributesDebtorParty) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"debtor_party"+"."+"birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesDebtorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesDebtorParty) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesDebtorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesFx payment attributes fx
// swagger:model PaymentAttributesFx
type PaymentAttributesFx struct {

	// contract reference
	ContractReference string `json:"contract_reference,omitempty"`

	// exchange rate
	ExchangeRate string `json:"exchange_rate,omitempty"`

	// original amount
	OriginalAmount string `json:"original_amount,omitempty"`

	// original currency
	OriginalCurrency string `json:"original_currency,omitempty"`
}

// Validate validates this payment attributes fx
func (m *PaymentAttributesFx) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesFx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesFx) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesFx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesStructuredReference payment attributes structured reference
// swagger:model PaymentAttributesStructuredReference
type PaymentAttributesStructuredReference struct {

	// issuer
	Issuer string `json:"issuer,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`
}

// Validate validates this payment attributes structured reference
func (m *PaymentAttributesStructuredReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesStructuredReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesStructuredReference) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesStructuredReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesSwift payment attributes swift
// swagger:model PaymentAttributesSwift
type PaymentAttributesSwift struct {

	// bank operation code
	BankOperationCode string `json:"bank_operation_code,omitempty"`

	// header
	Header *PaymentAttributesSwiftHeader `json:"header,omitempty"`

	// instruction code
	InstructionCode string `json:"instruction_code,omitempty"`

	// sender receiver information
	SenderReceiverInformation string `json:"sender_receiver_information,omitempty"`

	// time indication
	TimeIndication string `json:"time_indication,omitempty"`
}

// Validate validates this payment attributes swift
func (m *PaymentAttributesSwift) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributesSwift) validateHeader(formats strfmt.Registry) error {

	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "swift" + "." + "header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesSwift) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesSwift) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesSwift
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaymentAttributesSwiftHeader payment attributes swift header
// swagger:model PaymentAttributesSwiftHeader
type PaymentAttributesSwiftHeader struct {

	// destination
	Destination string `json:"destination,omitempty"`

	// message type
	MessageType string `json:"message_type,omitempty"`

	// priority
	Priority string `json:"priority,omitempty"`

	// recipient
	Recipient string `json:"recipient,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// user reference
	UserReference string `json:"user_reference,omitempty"`
}

// Validate validates this payment attributes swift header
func (m *PaymentAttributesSwiftHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributesSwiftHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributesSwiftHeader) UnmarshalBinary(b []byte) error {
	var res PaymentAttributesSwiftHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}