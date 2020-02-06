// Code generated by go-swagger; DO NOT EDIT.

package aa_training_status_of_project_version_controller

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

// NewListAaTrainingStatusOfProjectVersionParams creates a new ListAaTrainingStatusOfProjectVersionParams object
// with the default values initialized.
func NewListAaTrainingStatusOfProjectVersionParams() *ListAaTrainingStatusOfProjectVersionParams {
	var ()
	return &ListAaTrainingStatusOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAaTrainingStatusOfProjectVersionParamsWithTimeout creates a new ListAaTrainingStatusOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAaTrainingStatusOfProjectVersionParamsWithTimeout(timeout time.Duration) *ListAaTrainingStatusOfProjectVersionParams {
	var ()
	return &ListAaTrainingStatusOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewListAaTrainingStatusOfProjectVersionParamsWithContext creates a new ListAaTrainingStatusOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAaTrainingStatusOfProjectVersionParamsWithContext(ctx context.Context) *ListAaTrainingStatusOfProjectVersionParams {
	var ()
	return &ListAaTrainingStatusOfProjectVersionParams{

		Context: ctx,
	}
}

// NewListAaTrainingStatusOfProjectVersionParamsWithHTTPClient creates a new ListAaTrainingStatusOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAaTrainingStatusOfProjectVersionParamsWithHTTPClient(client *http.Client) *ListAaTrainingStatusOfProjectVersionParams {
	var ()
	return &ListAaTrainingStatusOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ListAaTrainingStatusOfProjectVersionParams contains all the parameters to send to the API endpoint
for the list aa training status of project version operation typically these are written to a http.Request
*/
type ListAaTrainingStatusOfProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) WithTimeout(timeout time.Duration) *ListAaTrainingStatusOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) WithContext(ctx context.Context) *ListAaTrainingStatusOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) WithHTTPClient(client *http.Client) *ListAaTrainingStatusOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) WithFields(fields *string) *ListAaTrainingStatusOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithParentID adds the parentID to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) WithParentID(parentID int64) *ListAaTrainingStatusOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list aa training status of project version params
func (o *ListAaTrainingStatusOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAaTrainingStatusOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
