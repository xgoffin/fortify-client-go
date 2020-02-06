// Code generated by go-swagger; DO NOT EDIT.

package ldap_object_controller

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

// NewUpdateLdapObjectParams creates a new UpdateLdapObjectParams object
// with the default values initialized.
func NewUpdateLdapObjectParams() *UpdateLdapObjectParams {
	var ()
	return &UpdateLdapObjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapObjectParamsWithTimeout creates a new UpdateLdapObjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLdapObjectParamsWithTimeout(timeout time.Duration) *UpdateLdapObjectParams {
	var ()
	return &UpdateLdapObjectParams{

		timeout: timeout,
	}
}

// NewUpdateLdapObjectParamsWithContext creates a new UpdateLdapObjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLdapObjectParamsWithContext(ctx context.Context) *UpdateLdapObjectParams {
	var ()
	return &UpdateLdapObjectParams{

		Context: ctx,
	}
}

// NewUpdateLdapObjectParamsWithHTTPClient creates a new UpdateLdapObjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLdapObjectParamsWithHTTPClient(client *http.Client) *UpdateLdapObjectParams {
	var ()
	return &UpdateLdapObjectParams{
		HTTPClient: client,
	}
}

/*UpdateLdapObjectParams contains all the parameters to send to the API endpoint
for the update ldap object operation typically these are written to a http.Request
*/
type UpdateLdapObjectParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.LDAPEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ldap object params
func (o *UpdateLdapObjectParams) WithTimeout(timeout time.Duration) *UpdateLdapObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap object params
func (o *UpdateLdapObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap object params
func (o *UpdateLdapObjectParams) WithContext(ctx context.Context) *UpdateLdapObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap object params
func (o *UpdateLdapObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap object params
func (o *UpdateLdapObjectParams) WithHTTPClient(client *http.Client) *UpdateLdapObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap object params
func (o *UpdateLdapObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update ldap object params
func (o *UpdateLdapObjectParams) WithID(id int64) *UpdateLdapObjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update ldap object params
func (o *UpdateLdapObjectParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update ldap object params
func (o *UpdateLdapObjectParams) WithResource(resource *models.LDAPEntity) *UpdateLdapObjectParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update ldap object params
func (o *UpdateLdapObjectParams) SetResource(resource *models.LDAPEntity) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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
