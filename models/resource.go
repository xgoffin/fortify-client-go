// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Resource resource
// swagger:model Resource
type Resource struct {

	// description
	Description string `json:"description,omitempty"`

	// file
	File *File `json:"file,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// input stream
	InputStream InputStream `json:"inputStream,omitempty"`

	// open
	Open bool `json:"open,omitempty"`

	// readable
	Readable bool `json:"readable,omitempty"`

	// uri
	URI *URI `json:"uri,omitempty"`

	// url
	URL *URL `json:"url,omitempty"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateFile(formats strfmt.Registry) error {

	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateURI(formats strfmt.Registry) error {

	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if m.URI != nil {
		if err := m.URI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uri")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if m.URL != nil {
		if err := m.URL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
