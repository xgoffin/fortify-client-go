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
)

// NewDeleteIssueTemplateParams creates a new DeleteIssueTemplateParams object
// with the default values initialized.
func NewDeleteIssueTemplateParams() *DeleteIssueTemplateParams {
	var ()
	return &DeleteIssueTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIssueTemplateParamsWithTimeout creates a new DeleteIssueTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIssueTemplateParamsWithTimeout(timeout time.Duration) *DeleteIssueTemplateParams {
	var ()
	return &DeleteIssueTemplateParams{

		timeout: timeout,
	}
}

// NewDeleteIssueTemplateParamsWithContext creates a new DeleteIssueTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIssueTemplateParamsWithContext(ctx context.Context) *DeleteIssueTemplateParams {
	var ()
	return &DeleteIssueTemplateParams{

		Context: ctx,
	}
}

// NewDeleteIssueTemplateParamsWithHTTPClient creates a new DeleteIssueTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIssueTemplateParamsWithHTTPClient(client *http.Client) *DeleteIssueTemplateParams {
	var ()
	return &DeleteIssueTemplateParams{
		HTTPClient: client,
	}
}

/*DeleteIssueTemplateParams contains all the parameters to send to the API endpoint
for the delete issue template operation typically these are written to a http.Request
*/
type DeleteIssueTemplateParams struct {

	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete issue template params
func (o *DeleteIssueTemplateParams) WithTimeout(timeout time.Duration) *DeleteIssueTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete issue template params
func (o *DeleteIssueTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete issue template params
func (o *DeleteIssueTemplateParams) WithContext(ctx context.Context) *DeleteIssueTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete issue template params
func (o *DeleteIssueTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete issue template params
func (o *DeleteIssueTemplateParams) WithHTTPClient(client *http.Client) *DeleteIssueTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete issue template params
func (o *DeleteIssueTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete issue template params
func (o *DeleteIssueTemplateParams) WithID(id string) *DeleteIssueTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete issue template params
func (o *DeleteIssueTemplateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIssueTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
