// Code generated by go-swagger; DO NOT EDIT.

package alert_definition_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// NewUpdateAlertDefinitionParams creates a new UpdateAlertDefinitionParams object
// with the default values initialized.
func NewUpdateAlertDefinitionParams() *UpdateAlertDefinitionParams {
	var ()
	return &UpdateAlertDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAlertDefinitionParamsWithTimeout creates a new UpdateAlertDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAlertDefinitionParamsWithTimeout(timeout time.Duration) *UpdateAlertDefinitionParams {
	var ()
	return &UpdateAlertDefinitionParams{

		timeout: timeout,
	}
}

// NewUpdateAlertDefinitionParamsWithContext creates a new UpdateAlertDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAlertDefinitionParamsWithContext(ctx context.Context) *UpdateAlertDefinitionParams {
	var ()
	return &UpdateAlertDefinitionParams{

		Context: ctx,
	}
}

// NewUpdateAlertDefinitionParamsWithHTTPClient creates a new UpdateAlertDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAlertDefinitionParamsWithHTTPClient(client *http.Client) *UpdateAlertDefinitionParams {
	var ()
	return &UpdateAlertDefinitionParams{
		HTTPClient: client,
	}
}

/*UpdateAlertDefinitionParams contains all the parameters to send to the API endpoint
for the update alert definition operation typically these are written to a http.Request
*/
type UpdateAlertDefinitionParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.AlertDefinitionDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update alert definition params
func (o *UpdateAlertDefinitionParams) WithTimeout(timeout time.Duration) *UpdateAlertDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update alert definition params
func (o *UpdateAlertDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update alert definition params
func (o *UpdateAlertDefinitionParams) WithContext(ctx context.Context) *UpdateAlertDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update alert definition params
func (o *UpdateAlertDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update alert definition params
func (o *UpdateAlertDefinitionParams) WithHTTPClient(client *http.Client) *UpdateAlertDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update alert definition params
func (o *UpdateAlertDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update alert definition params
func (o *UpdateAlertDefinitionParams) WithID(id int64) *UpdateAlertDefinitionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update alert definition params
func (o *UpdateAlertDefinitionParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update alert definition params
func (o *UpdateAlertDefinitionParams) WithResource(resource *models.AlertDefinitionDto) *UpdateAlertDefinitionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update alert definition params
func (o *UpdateAlertDefinitionParams) SetResource(resource *models.AlertDefinitionDto) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAlertDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

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
