// Code generated by go-swagger; DO NOT EDIT.

package responsibility_assignment_of_project_version_controller

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

// NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParams creates a new UpdateCollectionResponsibilityAssignmentOfProjectVersionParams object
// with the default values initialized.
func NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParams() *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	var ()
	return &UpdateCollectionResponsibilityAssignmentOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParamsWithTimeout creates a new UpdateCollectionResponsibilityAssignmentOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParamsWithTimeout(timeout time.Duration) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	var ()
	return &UpdateCollectionResponsibilityAssignmentOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParamsWithContext creates a new UpdateCollectionResponsibilityAssignmentOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParamsWithContext(ctx context.Context) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	var ()
	return &UpdateCollectionResponsibilityAssignmentOfProjectVersionParams{

		Context: ctx,
	}
}

// NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParamsWithHTTPClient creates a new UpdateCollectionResponsibilityAssignmentOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCollectionResponsibilityAssignmentOfProjectVersionParamsWithHTTPClient(client *http.Client) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	var ()
	return &UpdateCollectionResponsibilityAssignmentOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*UpdateCollectionResponsibilityAssignmentOfProjectVersionParams contains all the parameters to send to the API endpoint
for the update collection responsibility assignment of project version operation typically these are written to a http.Request
*/
type UpdateCollectionResponsibilityAssignmentOfProjectVersionParams struct {

	/*Data
	  data

	*/
	Data []*models.ResponsibilityAssignment
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) WithTimeout(timeout time.Duration) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) WithContext(ctx context.Context) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) WithHTTPClient(client *http.Client) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) WithData(data []*models.ResponsibilityAssignment) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) SetData(data []*models.ResponsibilityAssignment) {
	o.Data = data
}

// WithParentID adds the parentID to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) WithParentID(parentID int64) *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the update collection responsibility assignment of project version params
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCollectionResponsibilityAssignmentOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
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
