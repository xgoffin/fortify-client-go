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

// CloudJobCancelRequest Request to cancel cloudscan jobs
// swagger:model CloudJobCancelRequest
type CloudJobCancelRequest struct {

	// List containing single cloud job token to cancel
	// Required: true
	JobTokens []string `json:"jobTokens"`
}

// Validate validates this cloud job cancel request
func (m *CloudJobCancelRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudJobCancelRequest) validateJobTokens(formats strfmt.Registry) error {

	if err := validate.Required("jobTokens", "body", m.JobTokens); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudJobCancelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudJobCancelRequest) UnmarshalBinary(b []byte) error {
	var res CloudJobCancelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
