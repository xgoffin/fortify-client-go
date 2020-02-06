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

// CloudJob Cloudscan job
// swagger:model CloudJob
type CloudJob struct {

	// cloud pool
	CloudPool *CloudPool `json:"cloudPool,omitempty"`

	// cloud worker
	CloudWorker *CloudWorker `json:"cloudWorker,omitempty"`

	// job cancellable
	JobCancellable bool `json:"jobCancellable,omitempty"`

	// job duration
	JobDuration int64 `json:"jobDuration,omitempty"`

	// job expiry time
	// Format: date-time
	JobExpiryTime strfmt.DateTime `json:"jobExpiryTime,omitempty"`

	// job finished time
	// Format: date-time
	JobFinishedTime strfmt.DateTime `json:"jobFinishedTime,omitempty"`

	// job has fpr
	JobHasFpr bool `json:"jobHasFpr,omitempty"`

	// job has log
	JobHasLog bool `json:"jobHasLog,omitempty"`

	// job queued time
	// Format: date-time
	JobQueuedTime strfmt.DateTime `json:"jobQueuedTime,omitempty"`

	// job started time
	// Format: date-time
	JobStartedTime strfmt.DateTime `json:"jobStartedTime,omitempty"`

	// job state
	JobState string `json:"jobState,omitempty"`

	// job token
	// Read Only: true
	JobToken string `json:"jobToken,omitempty"`

	// project Id
	ProjectID int64 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// pv Id
	PvID int64 `json:"pvId,omitempty"`

	// pv name
	PvName string `json:"pvName,omitempty"`

	// queued duration
	QueuedDuration int64 `json:"queuedDuration,omitempty"`

	// sca args
	ScaArgs string `json:"scaArgs,omitempty"`

	// sca build Id
	ScaBuildID string `json:"scaBuildId,omitempty"`

	// sca version
	ScaVersion string `json:"scaVersion,omitempty"`

	// scan duration
	ScanDuration int64 `json:"scanDuration,omitempty"`

	// submitter email
	SubmitterEmail string `json:"submitterEmail,omitempty"`

	// submitter Ip address
	SubmitterIPAddress string `json:"submitterIpAddress,omitempty"`

	// submitter user name
	SubmitterUserName string `json:"submitterUserName,omitempty"`
}

// Validate validates this cloud job
func (m *CloudJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudWorker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobFinishedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobQueuedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobStartedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudJob) validateCloudPool(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPool) { // not required
		return nil
	}

	if m.CloudPool != nil {
		if err := m.CloudPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudPool")
			}
			return err
		}
	}

	return nil
}

func (m *CloudJob) validateCloudWorker(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudWorker) { // not required
		return nil
	}

	if m.CloudWorker != nil {
		if err := m.CloudWorker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudWorker")
			}
			return err
		}
	}

	return nil
}

func (m *CloudJob) validateJobExpiryTime(formats strfmt.Registry) error {

	if swag.IsZero(m.JobExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("jobExpiryTime", "body", "date-time", m.JobExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudJob) validateJobFinishedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.JobFinishedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("jobFinishedTime", "body", "date-time", m.JobFinishedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudJob) validateJobQueuedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.JobQueuedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("jobQueuedTime", "body", "date-time", m.JobQueuedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudJob) validateJobStartedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.JobStartedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("jobStartedTime", "body", "date-time", m.JobStartedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudJob) UnmarshalBinary(b []byte) error {
	var res CloudJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
