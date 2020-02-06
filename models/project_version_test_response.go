// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectVersionTestResponse Response of testing whether application version with name and exists in specified application
// swagger:model ProjectVersionTestResponse
type ProjectVersionTestResponse struct {

	// Whether the application version was found
	// Read Only: true
	Found *bool `json:"found,omitempty"`
}

// Validate validates this project version test response
func (m *ProjectVersionTestResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVersionTestResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVersionTestResponse) UnmarshalBinary(b []byte) error {
	var res ProjectVersionTestResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
