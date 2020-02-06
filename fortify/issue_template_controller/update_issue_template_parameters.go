// Code generated by go-swagger; DO NOT EDIT.

package issue_template_controller

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

// NewUpdateIssueTemplateParams creates a new UpdateIssueTemplateParams object
// with the default values initialized.
func NewUpdateIssueTemplateParams() *UpdateIssueTemplateParams {
	var ()
	return &UpdateIssueTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIssueTemplateParamsWithTimeout creates a new UpdateIssueTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateIssueTemplateParamsWithTimeout(timeout time.Duration) *UpdateIssueTemplateParams {
	var ()
	return &UpdateIssueTemplateParams{

		timeout: timeout,
	}
}

// NewUpdateIssueTemplateParamsWithContext creates a new UpdateIssueTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateIssueTemplateParamsWithContext(ctx context.Context) *UpdateIssueTemplateParams {
	var ()
	return &UpdateIssueTemplateParams{

		Context: ctx,
	}
}

// NewUpdateIssueTemplateParamsWithHTTPClient creates a new UpdateIssueTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateIssueTemplateParamsWithHTTPClient(client *http.Client) *UpdateIssueTemplateParams {
	var ()
	return &UpdateIssueTemplateParams{
		HTTPClient: client,
	}
}

/*UpdateIssueTemplateParams contains all the parameters to send to the API endpoint
for the update issue template operation typically these are written to a http.Request
*/
type UpdateIssueTemplateParams struct {

	/*ID
	  id

	*/
	ID string
	/*Resource
	  resource

	*/
	Resource *models.IssueTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update issue template params
func (o *UpdateIssueTemplateParams) WithTimeout(timeout time.Duration) *UpdateIssueTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update issue template params
func (o *UpdateIssueTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update issue template params
func (o *UpdateIssueTemplateParams) WithContext(ctx context.Context) *UpdateIssueTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update issue template params
func (o *UpdateIssueTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update issue template params
func (o *UpdateIssueTemplateParams) WithHTTPClient(client *http.Client) *UpdateIssueTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update issue template params
func (o *UpdateIssueTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update issue template params
func (o *UpdateIssueTemplateParams) WithID(id string) *UpdateIssueTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update issue template params
func (o *UpdateIssueTemplateParams) SetID(id string) {
	o.ID = id
}

// WithResource adds the resource to the update issue template params
func (o *UpdateIssueTemplateParams) WithResource(resource *models.IssueTemplate) *UpdateIssueTemplateParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update issue template params
func (o *UpdateIssueTemplateParams) SetResource(resource *models.IssueTemplate) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIssueTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
