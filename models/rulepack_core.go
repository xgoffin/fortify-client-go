// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RulepackCore Rulepack Core DTO object
// swagger:model Rulepack Core
type RulepackCore struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// language
	// Required: true
	Language *string `json:"language"`

	// locale
	Locale string `json:"locale,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// rulepack GUID
	// Required: true
	RulepackGUID *string `json:"rulepackGUID"`

	// rulepack type
	// Required: true
	// Enum: [SCA RTA CATPACK]
	RulepackType *string `json:"rulepackType"`

	// sku
	// Required: true
	Sku *string `json:"sku"`

	// version
	// Required: true
	Version *string `json:"version"`

	// versions
	Versions []*RulepackCore `json:"versions"`
}

// Validate validates this rulepack core
func (m *RulepackCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulepackGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulepackType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RulepackCore) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RulepackCore) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *RulepackCore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RulepackCore) validateRulepackGUID(formats strfmt.Registry) error {

	if err := validate.Required("rulepackGUID", "body", m.RulepackGUID); err != nil {
		return err
	}

	return nil
}

var rulepackCoreTypeRulepackTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SCA","RTA","CATPACK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rulepackCoreTypeRulepackTypePropEnum = append(rulepackCoreTypeRulepackTypePropEnum, v)
	}
}

const (

	// RulepackCoreRulepackTypeSCA captures enum value "SCA"
	RulepackCoreRulepackTypeSCA string = "SCA"

	// RulepackCoreRulepackTypeRTA captures enum value "RTA"
	RulepackCoreRulepackTypeRTA string = "RTA"

	// RulepackCoreRulepackTypeCATPACK captures enum value "CATPACK"
	RulepackCoreRulepackTypeCATPACK string = "CATPACK"
)

// prop value enum
func (m *RulepackCore) validateRulepackTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rulepackCoreTypeRulepackTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RulepackCore) validateRulepackType(formats strfmt.Registry) error {

	if err := validate.Required("rulepackType", "body", m.RulepackType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRulepackTypeEnum("rulepackType", "body", *m.RulepackType); err != nil {
		return err
	}

	return nil
}

func (m *RulepackCore) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}

func (m *RulepackCore) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *RulepackCore) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RulepackCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RulepackCore) UnmarshalBinary(b []byte) error {
	var res RulepackCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
