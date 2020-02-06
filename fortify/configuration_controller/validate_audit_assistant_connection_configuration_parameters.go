// Code generated by go-swagger; DO NOT EDIT.

package configuration_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// NewValidateAuditAssistantConnectionConfigurationParams creates a new ValidateAuditAssistantConnectionConfigurationParams object
// with the default values initialized.
func NewValidateAuditAssistantConnectionConfigurationParams() *ValidateAuditAssistantConnectionConfigurationParams {
	var ()
	return &ValidateAuditAssistantConnectionConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateAuditAssistantConnectionConfigurationParamsWithTimeout creates a new ValidateAuditAssistantConnectionConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateAuditAssistantConnectionConfigurationParamsWithTimeout(timeout time.Duration) *ValidateAuditAssistantConnectionConfigurationParams {
	var ()
	return &ValidateAuditAssistantConnectionConfigurationParams{

		timeout: timeout,
	}
}

// NewValidateAuditAssistantConnectionConfigurationParamsWithContext creates a new ValidateAuditAssistantConnectionConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateAuditAssistantConnectionConfigurationParamsWithContext(ctx context.Context) *ValidateAuditAssistantConnectionConfigurationParams {
	var ()
	return &ValidateAuditAssistantConnectionConfigurationParams{

		Context: ctx,
	}
}

// NewValidateAuditAssistantConnectionConfigurationParamsWithHTTPClient creates a new ValidateAuditAssistantConnectionConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateAuditAssistantConnectionConfigurationParamsWithHTTPClient(client *http.Client) *ValidateAuditAssistantConnectionConfigurationParams {
	var ()
	return &ValidateAuditAssistantConnectionConfigurationParams{
		HTTPClient: client,
	}
}

/*ValidateAuditAssistantConnectionConfigurationParams contains all the parameters to send to the API endpoint
for the validate audit assistant connection configuration operation typically these are written to a http.Request
*/
type ValidateAuditAssistantConnectionConfigurationParams struct {

	/*Resource
	  resource

	*/
	Resource *models.ConfigProperties

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) WithTimeout(timeout time.Duration) *ValidateAuditAssistantConnectionConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) WithContext(ctx context.Context) *ValidateAuditAssistantConnectionConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) WithHTTPClient(client *http.Client) *ValidateAuditAssistantConnectionConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) WithResource(resource *models.ConfigProperties) *ValidateAuditAssistantConnectionConfigurationParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the validate audit assistant connection configuration params
func (o *ValidateAuditAssistantConnectionConfigurationParams) SetResource(resource *models.ConfigProperties) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateAuditAssistantConnectionConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Resource != nil {
		if err := r.SetBodyParam(o.Resource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
