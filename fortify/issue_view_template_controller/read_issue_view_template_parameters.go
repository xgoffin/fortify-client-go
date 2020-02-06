// Code generated by go-swagger; DO NOT EDIT.

package issue_view_template_controller

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
)

// NewReadIssueViewTemplateParams creates a new ReadIssueViewTemplateParams object
// with the default values initialized.
func NewReadIssueViewTemplateParams() *ReadIssueViewTemplateParams {
	var ()
	return &ReadIssueViewTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadIssueViewTemplateParamsWithTimeout creates a new ReadIssueViewTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadIssueViewTemplateParamsWithTimeout(timeout time.Duration) *ReadIssueViewTemplateParams {
	var ()
	return &ReadIssueViewTemplateParams{

		timeout: timeout,
	}
}

// NewReadIssueViewTemplateParamsWithContext creates a new ReadIssueViewTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadIssueViewTemplateParamsWithContext(ctx context.Context) *ReadIssueViewTemplateParams {
	var ()
	return &ReadIssueViewTemplateParams{

		Context: ctx,
	}
}

// NewReadIssueViewTemplateParamsWithHTTPClient creates a new ReadIssueViewTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadIssueViewTemplateParamsWithHTTPClient(client *http.Client) *ReadIssueViewTemplateParams {
	var ()
	return &ReadIssueViewTemplateParams{
		HTTPClient: client,
	}
}

/*ReadIssueViewTemplateParams contains all the parameters to send to the API endpoint
for the read issue view template operation typically these are written to a http.Request
*/
type ReadIssueViewTemplateParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read issue view template params
func (o *ReadIssueViewTemplateParams) WithTimeout(timeout time.Duration) *ReadIssueViewTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read issue view template params
func (o *ReadIssueViewTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read issue view template params
func (o *ReadIssueViewTemplateParams) WithContext(ctx context.Context) *ReadIssueViewTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read issue view template params
func (o *ReadIssueViewTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read issue view template params
func (o *ReadIssueViewTemplateParams) WithHTTPClient(client *http.Client) *ReadIssueViewTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read issue view template params
func (o *ReadIssueViewTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read issue view template params
func (o *ReadIssueViewTemplateParams) WithFields(fields *string) *ReadIssueViewTemplateParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read issue view template params
func (o *ReadIssueViewTemplateParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read issue view template params
func (o *ReadIssueViewTemplateParams) WithID(id int64) *ReadIssueViewTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read issue view template params
func (o *ReadIssueViewTemplateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadIssueViewTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
