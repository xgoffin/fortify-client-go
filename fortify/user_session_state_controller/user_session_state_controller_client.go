// Code generated by go-swagger; DO NOT EDIT.

package user_session_state_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user session state controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user session state controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListUserSessionState(params *ListUserSessionStateParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserSessionStateOK, error)

	PutUserSessionState(params *PutUserSessionStateParams, authInfo runtime.ClientAuthInfoWriter) (*PutUserSessionStateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListUserSessionState lists
*/
func (a *Client) ListUserSessionState(params *ListUserSessionStateParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserSessionStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserSessionStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserSessionState",
		Method:             "GET",
		PathPattern:        "/userSession/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListUserSessionStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUserSessionStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserSessionState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutUserSessionState puts
*/
func (a *Client) PutUserSessionState(params *PutUserSessionStateParams, authInfo runtime.ClientAuthInfoWriter) (*PutUserSessionStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUserSessionStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserSessionState",
		Method:             "PUT",
		PathPattern:        "/userSession/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutUserSessionStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutUserSessionStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putUserSessionState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
