// Code generated by go-swagger; DO NOT EDIT.

package source_file_of_project_version_controller

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

// NewListSourceFileOfProjectVersionParams creates a new ListSourceFileOfProjectVersionParams object
// with the default values initialized.
func NewListSourceFileOfProjectVersionParams() *ListSourceFileOfProjectVersionParams {
	var ()
	return &ListSourceFileOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSourceFileOfProjectVersionParamsWithTimeout creates a new ListSourceFileOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSourceFileOfProjectVersionParamsWithTimeout(timeout time.Duration) *ListSourceFileOfProjectVersionParams {
	var ()
	return &ListSourceFileOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewListSourceFileOfProjectVersionParamsWithContext creates a new ListSourceFileOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSourceFileOfProjectVersionParamsWithContext(ctx context.Context) *ListSourceFileOfProjectVersionParams {
	var ()
	return &ListSourceFileOfProjectVersionParams{

		Context: ctx,
	}
}

// NewListSourceFileOfProjectVersionParamsWithHTTPClient creates a new ListSourceFileOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSourceFileOfProjectVersionParamsWithHTTPClient(client *http.Client) *ListSourceFileOfProjectVersionParams {
	var ()
	return &ListSourceFileOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ListSourceFileOfProjectVersionParams contains all the parameters to send to the API endpoint
for the list source file of project version operation typically these are written to a http.Request
*/
type ListSourceFileOfProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Q
	  The url-encoded value of a source file query expression "path:&lt;path_to_file_in_quotes&gt;+AND+scan_id:&lt;s_id&gt;". For example, "q=path:%22JavaSource%2Forg%2Fowasp%2Fwebgoat%2Flessons%2FCSRF.java%22%2Band%2Bscan_id:1". If 'scan_id' is not provided, the search defaults to the latest scan of the application version.

	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) WithTimeout(timeout time.Duration) *ListSourceFileOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) WithContext(ctx context.Context) *ListSourceFileOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) WithHTTPClient(client *http.Client) *ListSourceFileOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) WithFields(fields *string) *ListSourceFileOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithParentID adds the parentID to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) WithParentID(parentID int64) *ListSourceFileOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithQ adds the q to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) WithQ(q *string) *ListSourceFileOfProjectVersionParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list source file of project version params
func (o *ListSourceFileOfProjectVersionParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *ListSourceFileOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
