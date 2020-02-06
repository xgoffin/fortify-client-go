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

// Project An application on the server
// swagger:model Project
type Project struct {

	// User that created this application
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// Date created
	// Required: true
	// Read Only: true
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Description
	Description string `json:"description,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Id of the Issue Template used
	// Required: true
	IssueTemplateID *string `json:"issueTemplateId"`

	// Name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateIssueTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("issueTemplateId", "body", m.IssueTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
