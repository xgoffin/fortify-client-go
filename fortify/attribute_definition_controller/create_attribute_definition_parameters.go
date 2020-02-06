// Code generated by go-swagger; DO NOT EDIT.

package attribute_definition_controller

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

// NewCreateAttributeDefinitionParams creates a new CreateAttributeDefinitionParams object
// with the default values initialized.
func NewCreateAttributeDefinitionParams() *CreateAttributeDefinitionParams {
	var ()
	return &CreateAttributeDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAttributeDefinitionParamsWithTimeout creates a new CreateAttributeDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAttributeDefinitionParamsWithTimeout(timeout time.Duration) *CreateAttributeDefinitionParams {
	var ()
	return &CreateAttributeDefinitionParams{

		timeout: timeout,
	}
}

// NewCreateAttributeDefinitionParamsWithContext creates a new CreateAttributeDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAttributeDefinitionParamsWithContext(ctx context.Context) *CreateAttributeDefinitionParams {
	var ()
	return &CreateAttributeDefinitionParams{

		Context: ctx,
	}
}

// NewCreateAttributeDefinitionParamsWithHTTPClient creates a new CreateAttributeDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAttributeDefinitionParamsWithHTTPClient(client *http.Client) *CreateAttributeDefinitionParams {
	var ()
	return &CreateAttributeDefinitionParams{
		HTTPClient: client,
	}
}

/*CreateAttributeDefinitionParams contains all the parameters to send to the API endpoint
for the create attribute definition operation typically these are written to a http.Request
*/
type CreateAttributeDefinitionParams struct {

	/*Resource
	  resource

	*/
	Resource *models.AttributeDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create attribute definition params
func (o *CreateAttributeDefinitionParams) WithTimeout(timeout time.Duration) *CreateAttributeDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create attribute definition params
func (o *CreateAttributeDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create attribute definition params
func (o *CreateAttributeDefinitionParams) WithContext(ctx context.Context) *CreateAttributeDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create attribute definition params
func (o *CreateAttributeDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create attribute definition params
func (o *CreateAttributeDefinitionParams) WithHTTPClient(client *http.Client) *CreateAttributeDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create attribute definition params
func (o *CreateAttributeDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the create attribute definition params
func (o *CreateAttributeDefinitionParams) WithResource(resource *models.AttributeDefinition) *CreateAttributeDefinitionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the create attribute definition params
func (o *CreateAttributeDefinitionParams) SetResource(resource *models.AttributeDefinition) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAttributeDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
