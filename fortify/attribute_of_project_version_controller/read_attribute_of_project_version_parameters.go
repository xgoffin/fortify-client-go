// Code generated by go-swagger; DO NOT EDIT.

package attribute_of_project_version_controller

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

// NewReadAttributeOfProjectVersionParams creates a new ReadAttributeOfProjectVersionParams object
// with the default values initialized.
func NewReadAttributeOfProjectVersionParams() *ReadAttributeOfProjectVersionParams {
	var ()
	return &ReadAttributeOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadAttributeOfProjectVersionParamsWithTimeout creates a new ReadAttributeOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadAttributeOfProjectVersionParamsWithTimeout(timeout time.Duration) *ReadAttributeOfProjectVersionParams {
	var ()
	return &ReadAttributeOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewReadAttributeOfProjectVersionParamsWithContext creates a new ReadAttributeOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadAttributeOfProjectVersionParamsWithContext(ctx context.Context) *ReadAttributeOfProjectVersionParams {
	var ()
	return &ReadAttributeOfProjectVersionParams{

		Context: ctx,
	}
}

// NewReadAttributeOfProjectVersionParamsWithHTTPClient creates a new ReadAttributeOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadAttributeOfProjectVersionParamsWithHTTPClient(client *http.Client) *ReadAttributeOfProjectVersionParams {
	var ()
	return &ReadAttributeOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ReadAttributeOfProjectVersionParams contains all the parameters to send to the API endpoint
for the read attribute of project version operation typically these are written to a http.Request
*/
type ReadAttributeOfProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ID
	  id

	*/
	ID int64
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) WithTimeout(timeout time.Duration) *ReadAttributeOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) WithContext(ctx context.Context) *ReadAttributeOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) WithHTTPClient(client *http.Client) *ReadAttributeOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) WithFields(fields *string) *ReadAttributeOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) WithID(id int64) *ReadAttributeOfProjectVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) SetID(id int64) {
	o.ID = id
}

// WithParentID adds the parentID to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) WithParentID(parentID int64) *ReadAttributeOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the read attribute of project version params
func (o *ReadAttributeOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAttributeOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
