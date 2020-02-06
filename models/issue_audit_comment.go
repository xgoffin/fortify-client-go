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

// IssueAuditComment An audit comment associated with an issue
// swagger:model Issue Audit Comment
type IssueAuditComment struct {

	// audit time
	// Required: true
	// Read Only: true
	// Format: date-time
	AuditTime Iso8601MilliDateTime `json:"auditTime"`

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// issue engine type
	// Read Only: true
	IssueEngineType string `json:"issueEngineType,omitempty"`

	// issue Id
	// Required: true
	// Read Only: true
	IssueID int64 `json:"issueId"`

	// issue instance Id
	// Read Only: true
	IssueInstanceID string `json:"issueInstanceId,omitempty"`

	// issue name
	// Read Only: true
	IssueName string `json:"issueName,omitempty"`

	// project name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// project version Id
	// Read Only: true
	ProjectVersionID int64 `json:"projectVersionId,omitempty"`

	// project version name
	// Read Only: true
	ProjectVersionName string `json:"projectVersionName,omitempty"`

	// seq number
	// Required: true
	// Read Only: true
	SeqNumber int32 `json:"seqNumber"`

	// user name
	// Required: true
	// Read Only: true
	UserName string `json:"userName"`
}

// Validate validates this issue audit comment
func (m *IssueAuditComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeqNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAuditComment) validateAuditTime(formats strfmt.Registry) error {

	if err := validate.Required("auditTime", "body", Iso8601MilliDateTime(m.AuditTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("auditTime", "body", "date-time", m.AuditTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditComment) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditComment) validateIssueID(formats strfmt.Registry) error {

	if err := validate.Required("issueId", "body", int64(m.IssueID)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditComment) validateSeqNumber(formats strfmt.Registry) error {

	if err := validate.Required("seqNumber", "body", int32(m.SeqNumber)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditComment) validateUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("userName", "body", string(m.UserName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueAuditComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAuditComment) UnmarshalBinary(b []byte) error {
	var res IssueAuditComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
