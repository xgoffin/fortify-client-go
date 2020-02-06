// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIResultCloudPoolWorkerActionResponse Api result cloud pool worker action response
// swagger:model ApiResult«CloudPoolWorkerActionResponse»
type APIResultCloudPoolWorkerActionResponse struct {

	// count
	Count int64 `json:"count,omitempty"`

	// data
	Data *CloudPoolWorkerActionResponse `json:"data,omitempty"`

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// links
	Links interface{} `json:"links,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// response code
	ResponseCode int32 `json:"responseCode,omitempty"`

	// stack trace
	StackTrace string `json:"stackTrace,omitempty"`

	// success count
	SuccessCount int64 `json:"successCount,omitempty"`
}

// Validate validates this Api result cloud pool worker action response
func (m *APIResultCloudPoolWorkerActionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIResultCloudPoolWorkerActionResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIResultCloudPoolWorkerActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResultCloudPoolWorkerActionResponse) UnmarshalBinary(b []byte) error {
	var res APIResultCloudPoolWorkerActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
