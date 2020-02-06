// Code generated by go-swagger; DO NOT EDIT.

package issue_of_project_version_controller

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

// NewUpdateTagForIssueOfProjectVersionParams creates a new UpdateTagForIssueOfProjectVersionParams object
// with the default values initialized.
func NewUpdateTagForIssueOfProjectVersionParams() *UpdateTagForIssueOfProjectVersionParams {
	var ()
	return &UpdateTagForIssueOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTagForIssueOfProjectVersionParamsWithTimeout creates a new UpdateTagForIssueOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTagForIssueOfProjectVersionParamsWithTimeout(timeout time.Duration) *UpdateTagForIssueOfProjectVersionParams {
	var ()
	return &UpdateTagForIssueOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewUpdateTagForIssueOfProjectVersionParamsWithContext creates a new UpdateTagForIssueOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTagForIssueOfProjectVersionParamsWithContext(ctx context.Context) *UpdateTagForIssueOfProjectVersionParams {
	var ()
	return &UpdateTagForIssueOfProjectVersionParams{

		Context: ctx,
	}
}

// NewUpdateTagForIssueOfProjectVersionParamsWithHTTPClient creates a new UpdateTagForIssueOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTagForIssueOfProjectVersionParamsWithHTTPClient(client *http.Client) *UpdateTagForIssueOfProjectVersionParams {
	var ()
	return &UpdateTagForIssueOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*UpdateTagForIssueOfProjectVersionParams contains all the parameters to send to the API endpoint
for the update tag for issue of project version operation typically these are written to a http.Request
*/
type UpdateTagForIssueOfProjectVersionParams struct {

	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Resource
	  resource

	*/
	Resource *models.IssueUpdateTagRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) WithTimeout(timeout time.Duration) *UpdateTagForIssueOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) WithContext(ctx context.Context) *UpdateTagForIssueOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) WithHTTPClient(client *http.Client) *UpdateTagForIssueOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) WithParentID(parentID int64) *UpdateTagForIssueOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithResource adds the resource to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) WithResource(resource *models.IssueUpdateTagRequest) *UpdateTagForIssueOfProjectVersionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update tag for issue of project version params
func (o *UpdateTagForIssueOfProjectVersionParams) SetResource(resource *models.IssueUpdateTagRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTagForIssueOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
