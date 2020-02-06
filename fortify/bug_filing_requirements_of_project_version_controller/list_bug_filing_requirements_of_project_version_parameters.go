// Code generated by go-swagger; DO NOT EDIT.

package bug_filing_requirements_of_project_version_controller

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

// NewListBugFilingRequirementsOfProjectVersionParams creates a new ListBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized.
func NewListBugFilingRequirementsOfProjectVersionParams() *ListBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &ListBugFilingRequirementsOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBugFilingRequirementsOfProjectVersionParamsWithTimeout creates a new ListBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBugFilingRequirementsOfProjectVersionParamsWithTimeout(timeout time.Duration) *ListBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &ListBugFilingRequirementsOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewListBugFilingRequirementsOfProjectVersionParamsWithContext creates a new ListBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBugFilingRequirementsOfProjectVersionParamsWithContext(ctx context.Context) *ListBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &ListBugFilingRequirementsOfProjectVersionParams{

		Context: ctx,
	}
}

// NewListBugFilingRequirementsOfProjectVersionParamsWithHTTPClient creates a new ListBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBugFilingRequirementsOfProjectVersionParamsWithHTTPClient(client *http.Client) *ListBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &ListBugFilingRequirementsOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ListBugFilingRequirementsOfProjectVersionParams contains all the parameters to send to the API endpoint
for the list bug filing requirements of project version operation typically these are written to a http.Request
*/
type ListBugFilingRequirementsOfProjectVersionParams struct {

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

// WithTimeout adds the timeout to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) WithTimeout(timeout time.Duration) *ListBugFilingRequirementsOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) WithContext(ctx context.Context) *ListBugFilingRequirementsOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) WithHTTPClient(client *http.Client) *ListBugFilingRequirementsOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) WithFields(fields *string) *ListBugFilingRequirementsOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithParentID adds the parentID to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) WithParentID(parentID int64) *ListBugFilingRequirementsOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list bug filing requirements of project version params
func (o *ListBugFilingRequirementsOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ListBugFilingRequirementsOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
