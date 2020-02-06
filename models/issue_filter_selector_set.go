// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IssueFilterSelectorSet Composite object that holds information about issue attributes that can be used for issue list filtering and grouping.
// swagger:model IssueFilterSelectorSet
type IssueFilterSelectorSet struct {

	// List of all possible issues filterring attributes.
	// Required: true
	// Read Only: true
	FilterBySet []*IssueFilterSelector `json:"filterBySet"`

	// List of all possible issues grouping attributes.
	// Required: true
	// Read Only: true
	GroupBySet []*IssueSelector `json:"groupBySet"`
}

// Validate validates this issue filter selector set
func (m *IssueFilterSelectorSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterBySet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupBySet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueFilterSelectorSet) validateFilterBySet(formats strfmt.Registry) error {

	if err := validate.Required("filterBySet", "body", m.FilterBySet); err != nil {
		return err
	}

	for i := 0; i < len(m.FilterBySet); i++ {
		if swag.IsZero(m.FilterBySet[i]) { // not required
			continue
		}

		if m.FilterBySet[i] != nil {
			if err := m.FilterBySet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filterBySet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IssueFilterSelectorSet) validateGroupBySet(formats strfmt.Registry) error {

	if err := validate.Required("groupBySet", "body", m.GroupBySet); err != nil {
		return err
	}

	for i := 0; i < len(m.GroupBySet); i++ {
		if swag.IsZero(m.GroupBySet[i]) { // not required
			continue
		}

		if m.GroupBySet[i] != nil {
			if err := m.GroupBySet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groupBySet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueFilterSelectorSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueFilterSelectorSet) UnmarshalBinary(b []byte) error {
	var res IssueFilterSelectorSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
