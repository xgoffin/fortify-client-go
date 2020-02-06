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

// AATrainingStatus Audit Assistant Training (AATrainingStatus) DTO object.
// swagger:model AATrainingStatus
type AATrainingStatus struct {

	// last training time
	// Format: date-time
	LastTrainingTime strfmt.DateTime `json:"lastTrainingTime,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// project version Id
	ProjectVersionID int64 `json:"projectVersionId,omitempty"`

	// status
	// Enum: [NONE TRAINING_SUBMITTED TRAINING_FAILED TRAINING_COMPLETE]
	Status string `json:"status,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this a a training status
func (m *AATrainingStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTrainingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AATrainingStatus) validateLastTrainingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTrainingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastTrainingTime", "body", "date-time", m.LastTrainingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var aATrainingStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","TRAINING_SUBMITTED","TRAINING_FAILED","TRAINING_COMPLETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aATrainingStatusTypeStatusPropEnum = append(aATrainingStatusTypeStatusPropEnum, v)
	}
}

const (

	// AATrainingStatusStatusNONE captures enum value "NONE"
	AATrainingStatusStatusNONE string = "NONE"

	// AATrainingStatusStatusTRAININGSUBMITTED captures enum value "TRAINING_SUBMITTED"
	AATrainingStatusStatusTRAININGSUBMITTED string = "TRAINING_SUBMITTED"

	// AATrainingStatusStatusTRAININGFAILED captures enum value "TRAINING_FAILED"
	AATrainingStatusStatusTRAININGFAILED string = "TRAINING_FAILED"

	// AATrainingStatusStatusTRAININGCOMPLETE captures enum value "TRAINING_COMPLETE"
	AATrainingStatusStatusTRAININGCOMPLETE string = "TRAINING_COMPLETE"
)

// prop value enum
func (m *AATrainingStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, aATrainingStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AATrainingStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AATrainingStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AATrainingStatus) UnmarshalBinary(b []byte) error {
	var res AATrainingStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
