// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIBulkRequestItem Api bulk request item
// swagger:model ApiBulkRequestItem
type APIBulkRequestItem struct {

	// http verb
	// Enum: [GET PUT POST DELETE]
	HTTPVerb string `json:"httpVerb,omitempty"`

	// post data
	PostData interface{} `json:"postData,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`
}

// Validate validates this Api bulk request item
func (m *APIBulkRequestItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiBulkRequestItemTypeHTTPVerbPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET","PUT","POST","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiBulkRequestItemTypeHTTPVerbPropEnum = append(apiBulkRequestItemTypeHTTPVerbPropEnum, v)
	}
}

const (

	// APIBulkRequestItemHTTPVerbGET captures enum value "GET"
	APIBulkRequestItemHTTPVerbGET string = "GET"

	// APIBulkRequestItemHTTPVerbPUT captures enum value "PUT"
	APIBulkRequestItemHTTPVerbPUT string = "PUT"

	// APIBulkRequestItemHTTPVerbPOST captures enum value "POST"
	APIBulkRequestItemHTTPVerbPOST string = "POST"

	// APIBulkRequestItemHTTPVerbDELETE captures enum value "DELETE"
	APIBulkRequestItemHTTPVerbDELETE string = "DELETE"
)

// prop value enum
func (m *APIBulkRequestItem) validateHTTPVerbEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, apiBulkRequestItemTypeHTTPVerbPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *APIBulkRequestItem) validateHTTPVerb(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPVerb) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPVerbEnum("httpVerb", "body", m.HTTPVerb); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIBulkRequestItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIBulkRequestItem) UnmarshalBinary(b []byte) error {
	var res APIBulkRequestItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
