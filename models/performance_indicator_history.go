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

// PerformanceIndicatorHistory Performance Indicator History DTO object.
// swagger:model Performance Indicator History
type PerformanceIndicatorHistory struct {

	// description
	Description string `json:"description,omitempty"`

	// ID required when referencing an existing Performance Indicator History object
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Range of values
	Range string `json:"range,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp Iso8601MilliDateTime `json:"timestamp,omitempty"`

	// Required if referencing an existing Performance Indicator History  object
	// Required: true
	Value *float32 `json:"value"`
}

// Validate validates this performance indicator history
func (m *PerformanceIndicatorHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceIndicatorHistory) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceIndicatorHistory) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceIndicatorHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceIndicatorHistory) UnmarshalBinary(b []byte) error {
	var res PerformanceIndicatorHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
