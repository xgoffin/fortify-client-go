// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIResultIssueActionResponse Api result issue action response
// swagger:model ApiResult«IssueActionResponse»
type APIResultIssueActionResponse struct {

	// count
	Count int64 `json:"count,omitempty"`

	// data
	Data *IssueActionResponse `json:"data,omitempty"`

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

// Validate validates this Api result issue action response
func (m *APIResultIssueActionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIResultIssueActionResponse) validateData(formats strfmt.Registry) error {

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
func (m *APIResultIssueActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResultIssueActionResponse) UnmarshalBinary(b []byte) error {
	var res APIResultIssueActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
