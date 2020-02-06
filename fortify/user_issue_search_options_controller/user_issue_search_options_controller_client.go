// Code generated by go-swagger; DO NOT EDIT.

package user_issue_search_options_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user issue search options controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user issue search options controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserIssueSearchOptions(params *GetUserIssueSearchOptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserIssueSearchOptionsOK, error)

	UpdateUserIssueSearchOptions(params *UpdateUserIssueSearchOptionsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserIssueSearchOptionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetUserIssueSearchOptions gets
*/
func (a *Client) GetUserIssueSearchOptions(params *GetUserIssueSearchOptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserIssueSearchOptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserIssueSearchOptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserIssueSearchOptions",
		Method:             "GET",
		PathPattern:        "/userIssueSearchOptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserIssueSearchOptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserIssueSearchOptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserIssueSearchOptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUserIssueSearchOptions updates
*/
func (a *Client) UpdateUserIssueSearchOptions(params *UpdateUserIssueSearchOptionsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserIssueSearchOptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserIssueSearchOptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserIssueSearchOptions",
		Method:             "PUT",
		PathPattern:        "/userIssueSearchOptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserIssueSearchOptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserIssueSearchOptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserIssueSearchOptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
