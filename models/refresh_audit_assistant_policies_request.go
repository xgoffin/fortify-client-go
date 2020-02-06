// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RefreshAuditAssistantPoliciesRequest Audit Assistant server configuration (mandatory) and policy replacement mapping (optional)
// swagger:model RefreshAuditAssistantPoliciesRequest
type RefreshAuditAssistantPoliciesRequest struct {

	// A list with current policy names from Audit Assistant server.
	CurrentPolicies []string `json:"currentPolicies"`

	// A list of policy objects containing existing obsolete AV policy names as a result of comparison with current server policies.
	ObsoletePolicies []*ObsoletePolicy `json:"obsoletePolicies"`

	// Mapping between old (obsolete) and new (existing) policy names; current AV policies and/or system default policy are to be replaced
	PolicyReplacements []*PolicyReplacement `json:"policyReplacements"`

	// Audit Assistant server configuration properties to be used policies retrieval from Fortify Scan Analytics server
	// Required: true
	Properties []*ConfigProperty `json:"properties"`
}

// Validate validates this refresh audit assistant policies request
func (m *RefreshAuditAssistantPoliciesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObsoletePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyReplacements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RefreshAuditAssistantPoliciesRequest) validateObsoletePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.ObsoletePolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.ObsoletePolicies); i++ {
		if swag.IsZero(m.ObsoletePolicies[i]) { // not required
			continue
		}

		if m.ObsoletePolicies[i] != nil {
			if err := m.ObsoletePolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("obsoletePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RefreshAuditAssistantPoliciesRequest) validatePolicyReplacements(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyReplacements) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyReplacements); i++ {
		if swag.IsZero(m.PolicyReplacements[i]) { // not required
			continue
		}

		if m.PolicyReplacements[i] != nil {
			if err := m.PolicyReplacements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyReplacements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RefreshAuditAssistantPoliciesRequest) validateProperties(formats strfmt.Registry) error {

	if err := validate.Required("properties", "body", m.Properties); err != nil {
		return err
	}

	for i := 0; i < len(m.Properties); i++ {
		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {
			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RefreshAuditAssistantPoliciesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefreshAuditAssistantPoliciesRequest) UnmarshalBinary(b []byte) error {
	var res RefreshAuditAssistantPoliciesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
