// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudPoolMapping Application version to CloudScan pool mapping
// swagger:model CloudPoolMapping
type CloudPoolMapping struct {

	// cloud pool
	CloudPool *CloudPool `json:"cloudPool,omitempty"`

	// project version Id
	ProjectVersionID int64 `json:"projectVersionId,omitempty"`
}

// Validate validates this cloud pool mapping
func (m *CloudPoolMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudPoolMapping) validateCloudPool(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPool) { // not required
		return nil
	}

	if m.CloudPool != nil {
		if err := m.CloudPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudPool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudPoolMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudPoolMapping) UnmarshalBinary(b []byte) error {
	var res CloudPoolMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
