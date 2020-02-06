// Code generated by go-swagger; DO NOT EDIT.

package cloud_job_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// NewCancelCloudJobParams creates a new CancelCloudJobParams object
// with the default values initialized.
func NewCancelCloudJobParams() *CancelCloudJobParams {
	var ()
	return &CancelCloudJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelCloudJobParamsWithTimeout creates a new CancelCloudJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelCloudJobParamsWithTimeout(timeout time.Duration) *CancelCloudJobParams {
	var ()
	return &CancelCloudJobParams{

		timeout: timeout,
	}
}

// NewCancelCloudJobParamsWithContext creates a new CancelCloudJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelCloudJobParamsWithContext(ctx context.Context) *CancelCloudJobParams {
	var ()
	return &CancelCloudJobParams{

		Context: ctx,
	}
}

// NewCancelCloudJobParamsWithHTTPClient creates a new CancelCloudJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelCloudJobParamsWithHTTPClient(client *http.Client) *CancelCloudJobParams {
	var ()
	return &CancelCloudJobParams{
		HTTPClient: client,
	}
}

/*CancelCloudJobParams contains all the parameters to send to the API endpoint
for the cancel cloud job operation typically these are written to a http.Request
*/
type CancelCloudJobParams struct {

	/*Resource
	  resource

	*/
	Resource *models.CloudJobCancelRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel cloud job params
func (o *CancelCloudJobParams) WithTimeout(timeout time.Duration) *CancelCloudJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel cloud job params
func (o *CancelCloudJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel cloud job params
func (o *CancelCloudJobParams) WithContext(ctx context.Context) *CancelCloudJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel cloud job params
func (o *CancelCloudJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel cloud job params
func (o *CancelCloudJobParams) WithHTTPClient(client *http.Client) *CancelCloudJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel cloud job params
func (o *CancelCloudJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the cancel cloud job params
func (o *CancelCloudJobParams) WithResource(resource *models.CloudJobCancelRequest) *CancelCloudJobParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the cancel cloud job params
func (o *CancelCloudJobParams) SetResource(resource *models.CloudJobCancelRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CancelCloudJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
