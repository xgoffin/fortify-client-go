// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LdapServerTestRequest Request to test ldap server
// swagger:model LdapServerTestRequest
type LdapServerTestRequest struct {

	// LDAP server information
	// Required: true
	LdapConfig *LdapServerDto `json:"ldapConfig"`
}

// Validate validates this ldap server test request
func (m *LdapServerTestRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdapConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServerTestRequest) validateLdapConfig(formats strfmt.Registry) error {

	if err := validate.Required("ldapConfig", "body", m.LdapConfig); err != nil {
		return err
	}

	if m.LdapConfig != nil {
		if err := m.LdapConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerTestRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerTestRequest) UnmarshalBinary(b []byte) error {
	var res LdapServerTestRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
