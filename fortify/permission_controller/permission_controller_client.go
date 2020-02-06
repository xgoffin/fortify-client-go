// Code generated by go-swagger; DO NOT EDIT.

package permission_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new permission controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for permission controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListPermission(params *ListPermissionParams, authInfo runtime.ClientAuthInfoWriter) (*ListPermissionOK, error)

	ReadPermission(params *ReadPermissionParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPermissionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListPermission lists
*/
func (a *Client) ListPermission(params *ListPermissionParams, authInfo runtime.ClientAuthInfoWriter) (*ListPermissionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPermissionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPermission",
		Method:             "GET",
		PathPattern:        "/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPermissionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPermissionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPermission: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadPermission reads
*/
func (a *Client) ReadPermission(params *ReadPermissionParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPermissionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadPermissionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readPermission",
		Method:             "GET",
		PathPattern:        "/permissions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadPermissionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadPermissionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readPermission: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
