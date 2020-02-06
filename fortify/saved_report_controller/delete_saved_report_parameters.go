// Code generated by go-swagger; DO NOT EDIT.

package saved_report_controller

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

// NewDeleteSavedReportParams creates a new DeleteSavedReportParams object
// with the default values initialized.
func NewDeleteSavedReportParams() *DeleteSavedReportParams {
	var ()
	return &DeleteSavedReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSavedReportParamsWithTimeout creates a new DeleteSavedReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSavedReportParamsWithTimeout(timeout time.Duration) *DeleteSavedReportParams {
	var ()
	return &DeleteSavedReportParams{

		timeout: timeout,
	}
}

// NewDeleteSavedReportParamsWithContext creates a new DeleteSavedReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSavedReportParamsWithContext(ctx context.Context) *DeleteSavedReportParams {
	var ()
	return &DeleteSavedReportParams{

		Context: ctx,
	}
}

// NewDeleteSavedReportParamsWithHTTPClient creates a new DeleteSavedReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSavedReportParamsWithHTTPClient(client *http.Client) *DeleteSavedReportParams {
	var ()
	return &DeleteSavedReportParams{
		HTTPClient: client,
	}
}

/*DeleteSavedReportParams contains all the parameters to send to the API endpoint
for the delete saved report operation typically these are written to a http.Request
*/
type DeleteSavedReportParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete saved report params
func (o *DeleteSavedReportParams) WithTimeout(timeout time.Duration) *DeleteSavedReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete saved report params
func (o *DeleteSavedReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete saved report params
func (o *DeleteSavedReportParams) WithContext(ctx context.Context) *DeleteSavedReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete saved report params
func (o *DeleteSavedReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete saved report params
func (o *DeleteSavedReportParams) WithHTTPClient(client *http.Client) *DeleteSavedReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete saved report params
func (o *DeleteSavedReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete saved report params
func (o *DeleteSavedReportParams) WithID(id int64) *DeleteSavedReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete saved report params
func (o *DeleteSavedReportParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSavedReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
