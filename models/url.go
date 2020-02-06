// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// URL URL
// swagger:model URL
type URL struct {

	// authority
	Authority string `json:"authority,omitempty"`

	// content
	Content interface{} `json:"content,omitempty"`

	// default port
	DefaultPort int32 `json:"defaultPort,omitempty"`

	// deserialized fields
	DeserializedFields URLStreamHandler `json:"deserializedFields,omitempty"`

	// file
	File string `json:"file,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// serialized hash code
	SerializedHashCode int32 `json:"serializedHashCode,omitempty"`

	// user info
	UserInfo string `json:"userInfo,omitempty"`
}

// Validate validates this URL
func (m *URL) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *URL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *URL) UnmarshalBinary(b []byte) error {
	var res URL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
