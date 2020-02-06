// Code generated by go-swagger; DO NOT EDIT.

package report_definition_controller

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

// NewDeleteReportDefinitionParams creates a new DeleteReportDefinitionParams object
// with the default values initialized.
func NewDeleteReportDefinitionParams() *DeleteReportDefinitionParams {
	var ()
	return &DeleteReportDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReportDefinitionParamsWithTimeout creates a new DeleteReportDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteReportDefinitionParamsWithTimeout(timeout time.Duration) *DeleteReportDefinitionParams {
	var ()
	return &DeleteReportDefinitionParams{

		timeout: timeout,
	}
}

// NewDeleteReportDefinitionParamsWithContext creates a new DeleteReportDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteReportDefinitionParamsWithContext(ctx context.Context) *DeleteReportDefinitionParams {
	var ()
	return &DeleteReportDefinitionParams{

		Context: ctx,
	}
}

// NewDeleteReportDefinitionParamsWithHTTPClient creates a new DeleteReportDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteReportDefinitionParamsWithHTTPClient(client *http.Client) *DeleteReportDefinitionParams {
	var ()
	return &DeleteReportDefinitionParams{
		HTTPClient: client,
	}
}

/*DeleteReportDefinitionParams contains all the parameters to send to the API endpoint
for the delete report definition operation typically these are written to a http.Request
*/
type DeleteReportDefinitionParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete report definition params
func (o *DeleteReportDefinitionParams) WithTimeout(timeout time.Duration) *DeleteReportDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete report definition params
func (o *DeleteReportDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete report definition params
func (o *DeleteReportDefinitionParams) WithContext(ctx context.Context) *DeleteReportDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete report definition params
func (o *DeleteReportDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete report definition params
func (o *DeleteReportDefinitionParams) WithHTTPClient(client *http.Client) *DeleteReportDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete report definition params
func (o *DeleteReportDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete report definition params
func (o *DeleteReportDefinitionParams) WithID(id int64) *DeleteReportDefinitionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete report definition params
func (o *DeleteReportDefinitionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReportDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
