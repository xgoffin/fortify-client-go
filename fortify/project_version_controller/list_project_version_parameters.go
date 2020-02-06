// Code generated by go-swagger; DO NOT EDIT.

package project_version_controller

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

// NewListProjectVersionParams creates a new ListProjectVersionParams object
// with the default values initialized.
func NewListProjectVersionParams() *ListProjectVersionParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		onlyIfHasIssuesDefault  = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		OnlyIfHasIssues:  &onlyIfHasIssuesDefault,
		Start:            &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectVersionParamsWithTimeout creates a new ListProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProjectVersionParamsWithTimeout(timeout time.Duration) *ListProjectVersionParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		onlyIfHasIssuesDefault  = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		OnlyIfHasIssues:  &onlyIfHasIssuesDefault,
		Start:            &startDefault,

		timeout: timeout,
	}
}

// NewListProjectVersionParamsWithContext creates a new ListProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProjectVersionParamsWithContext(ctx context.Context) *ListProjectVersionParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		onlyIfHasIssuesDefault  = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		OnlyIfHasIssues:  &onlyIfHasIssuesDefault,
		Start:            &startDefault,

		Context: ctx,
	}
}

// NewListProjectVersionParamsWithHTTPClient creates a new ListProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProjectVersionParamsWithHTTPClient(client *http.Client) *ListProjectVersionParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		onlyIfHasIssuesDefault  = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		OnlyIfHasIssues:  &onlyIfHasIssuesDefault,
		Start:            &startDefault,
		HTTPClient:       client,
	}
}

/*ListProjectVersionParams contains all the parameters to send to the API endpoint
for the list project version operation typically these are written to a http.Request
*/
type ListProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*Fulltextsearch
	  If 'true', interpret 'q' parameter as full text search query, defaults to 'false'

	*/
	Fulltextsearch *bool
	/*IncludeInactive
	  includeInactive

	*/
	IncludeInactive *bool
	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*MyAssignedIssues
	  myAssignedIssues

	*/
	MyAssignedIssues *bool
	/*OnlyIfHasIssues
	  onlyIfHasIssues

	*/
	OnlyIfHasIssues *bool
	/*Orderby
	  Fields to order by

	*/
	Orderby *string
	/*Q
	  A search-spec of full text search query (see fulltextsearch parameter)

	*/
	Q *string
	/*Start
	  A start offset in object listing

	*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list project version params
func (o *ListProjectVersionParams) WithTimeout(timeout time.Duration) *ListProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project version params
func (o *ListProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project version params
func (o *ListProjectVersionParams) WithContext(ctx context.Context) *ListProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project version params
func (o *ListProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project version params
func (o *ListProjectVersionParams) WithHTTPClient(client *http.Client) *ListProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project version params
func (o *ListProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list project version params
func (o *ListProjectVersionParams) WithFields(fields *string) *ListProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list project version params
func (o *ListProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFulltextsearch adds the fulltextsearch to the list project version params
func (o *ListProjectVersionParams) WithFulltextsearch(fulltextsearch *bool) *ListProjectVersionParams {
	o.SetFulltextsearch(fulltextsearch)
	return o
}

// SetFulltextsearch adds the fulltextsearch to the list project version params
func (o *ListProjectVersionParams) SetFulltextsearch(fulltextsearch *bool) {
	o.Fulltextsearch = fulltextsearch
}

// WithIncludeInactive adds the includeInactive to the list project version params
func (o *ListProjectVersionParams) WithIncludeInactive(includeInactive *bool) *ListProjectVersionParams {
	o.SetIncludeInactive(includeInactive)
	return o
}

// SetIncludeInactive adds the includeInactive to the list project version params
func (o *ListProjectVersionParams) SetIncludeInactive(includeInactive *bool) {
	o.IncludeInactive = includeInactive
}

// WithLimit adds the limit to the list project version params
func (o *ListProjectVersionParams) WithLimit(limit *int32) *ListProjectVersionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list project version params
func (o *ListProjectVersionParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithMyAssignedIssues adds the myAssignedIssues to the list project version params
func (o *ListProjectVersionParams) WithMyAssignedIssues(myAssignedIssues *bool) *ListProjectVersionParams {
	o.SetMyAssignedIssues(myAssignedIssues)
	return o
}

// SetMyAssignedIssues adds the myAssignedIssues to the list project version params
func (o *ListProjectVersionParams) SetMyAssignedIssues(myAssignedIssues *bool) {
	o.MyAssignedIssues = myAssignedIssues
}

// WithOnlyIfHasIssues adds the onlyIfHasIssues to the list project version params
func (o *ListProjectVersionParams) WithOnlyIfHasIssues(onlyIfHasIssues *bool) *ListProjectVersionParams {
	o.SetOnlyIfHasIssues(onlyIfHasIssues)
	return o
}

// SetOnlyIfHasIssues adds the onlyIfHasIssues to the list project version params
func (o *ListProjectVersionParams) SetOnlyIfHasIssues(onlyIfHasIssues *bool) {
	o.OnlyIfHasIssues = onlyIfHasIssues
}

// WithOrderby adds the orderby to the list project version params
func (o *ListProjectVersionParams) WithOrderby(orderby *string) *ListProjectVersionParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the list project version params
func (o *ListProjectVersionParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithQ adds the q to the list project version params
func (o *ListProjectVersionParams) WithQ(q *string) *ListProjectVersionParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list project version params
func (o *ListProjectVersionParams) SetQ(q *string) {
	o.Q = q
}

// WithStart adds the start to the list project version params
func (o *ListProjectVersionParams) WithStart(start *int32) *ListProjectVersionParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list project version params
func (o *ListProjectVersionParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fulltextsearch != nil {

		// query param fulltextsearch
		var qrFulltextsearch bool
		if o.Fulltextsearch != nil {
			qrFulltextsearch = *o.Fulltextsearch
		}
		qFulltextsearch := swag.FormatBool(qrFulltextsearch)
		if qFulltextsearch != "" {
			if err := r.SetQueryParam("fulltextsearch", qFulltextsearch); err != nil {
				return err
			}
		}

	}

	if o.IncludeInactive != nil {

		// query param includeInactive
		var qrIncludeInactive bool
		if o.IncludeInactive != nil {
			qrIncludeInactive = *o.IncludeInactive
		}
		qIncludeInactive := swag.FormatBool(qrIncludeInactive)
		if qIncludeInactive != "" {
			if err := r.SetQueryParam("includeInactive", qIncludeInactive); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MyAssignedIssues != nil {

		// query param myAssignedIssues
		var qrMyAssignedIssues bool
		if o.MyAssignedIssues != nil {
			qrMyAssignedIssues = *o.MyAssignedIssues
		}
		qMyAssignedIssues := swag.FormatBool(qrMyAssignedIssues)
		if qMyAssignedIssues != "" {
			if err := r.SetQueryParam("myAssignedIssues", qMyAssignedIssues); err != nil {
				return err
			}
		}

	}

	if o.OnlyIfHasIssues != nil {

		// query param onlyIfHasIssues
		var qrOnlyIfHasIssues bool
		if o.OnlyIfHasIssues != nil {
			qrOnlyIfHasIssues = *o.OnlyIfHasIssues
		}
		qOnlyIfHasIssues := swag.FormatBool(qrOnlyIfHasIssues)
		if qOnlyIfHasIssues != "" {
			if err := r.SetQueryParam("onlyIfHasIssues", qOnlyIfHasIssues); err != nil {
				return err
			}
		}

	}

	if o.Orderby != nil {

		// query param orderby
		var qrOrderby string
		if o.Orderby != nil {
			qrOrderby = *o.Orderby
		}
		qOrderby := qrOrderby
		if qOrderby != "" {
			if err := r.SetQueryParam("orderby", qOrderby); err != nil {
				return err
			}
		}

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

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
