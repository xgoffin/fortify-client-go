// Code generated by go-swagger; DO NOT EDIT.

package project_version_of_auth_entity_controller

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

// NewAssignProjectVersionOfAuthEntityParams creates a new AssignProjectVersionOfAuthEntityParams object
// with the default values initialized.
func NewAssignProjectVersionOfAuthEntityParams() *AssignProjectVersionOfAuthEntityParams {
	var ()
	return &AssignProjectVersionOfAuthEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAssignProjectVersionOfAuthEntityParamsWithTimeout creates a new AssignProjectVersionOfAuthEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAssignProjectVersionOfAuthEntityParamsWithTimeout(timeout time.Duration) *AssignProjectVersionOfAuthEntityParams {
	var ()
	return &AssignProjectVersionOfAuthEntityParams{

		timeout: timeout,
	}
}

// NewAssignProjectVersionOfAuthEntityParamsWithContext creates a new AssignProjectVersionOfAuthEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewAssignProjectVersionOfAuthEntityParamsWithContext(ctx context.Context) *AssignProjectVersionOfAuthEntityParams {
	var ()
	return &AssignProjectVersionOfAuthEntityParams{

		Context: ctx,
	}
}

// NewAssignProjectVersionOfAuthEntityParamsWithHTTPClient creates a new AssignProjectVersionOfAuthEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAssignProjectVersionOfAuthEntityParamsWithHTTPClient(client *http.Client) *AssignProjectVersionOfAuthEntityParams {
	var ()
	return &AssignProjectVersionOfAuthEntityParams{
		HTTPClient: client,
	}
}

/*AssignProjectVersionOfAuthEntityParams contains all the parameters to send to the API endpoint
for the assign project version of auth entity operation typically these are written to a http.Request
*/
type AssignProjectVersionOfAuthEntityParams struct {

	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Resource
	  resource

	*/
	Resource *models.ProjectVersionAuthEntityAssignRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) WithTimeout(timeout time.Duration) *AssignProjectVersionOfAuthEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) WithContext(ctx context.Context) *AssignProjectVersionOfAuthEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) WithHTTPClient(client *http.Client) *AssignProjectVersionOfAuthEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) WithParentID(parentID int64) *AssignProjectVersionOfAuthEntityParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithResource adds the resource to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) WithResource(resource *models.ProjectVersionAuthEntityAssignRequest) *AssignProjectVersionOfAuthEntityParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the assign project version of auth entity params
func (o *AssignProjectVersionOfAuthEntityParams) SetResource(resource *models.ProjectVersionAuthEntityAssignRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *AssignProjectVersionOfAuthEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
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
