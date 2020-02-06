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

// IssueAttachment attachment (such as screenshot) associated with issue
// swagger:model Issue Attachment
type IssueAttachment struct {

	// description
	Description string `json:"description,omitempty"`

	// document ID referencing the attachment
	// Required: true
	// Read Only: true
	FileDocID int64 `json:"fileDocId"`

	// id
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// image
	// Format: byte
	Image strfmt.Base64 `json:"image,omitempty"`

	// original file name
	OriginalFileName string `json:"originalFileName,omitempty"`

	// update time
	// Required: true
	// Read Only: true
	// Format: date-time
	UpdateTime Iso8601MilliDateTime `json:"updateTime"`
}

// Validate validates this issue attachment
func (m *IssueAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileDocID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAttachment) validateFileDocID(formats strfmt.Registry) error {

	if err := validate.Required("fileDocId", "body", int64(m.FileDocID)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAttachment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAttachment) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("updateTime", "body", Iso8601MilliDateTime(m.UpdateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("updateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAttachment) UnmarshalBinary(b []byte) error {
	var res IssueAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
