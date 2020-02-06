// Code generated by go-swagger; DO NOT EDIT.

package folder_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new folder of project version controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for folder of project version controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListFolderOfProjectVersion(params *ListFolderOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListFolderOfProjectVersionOK, error)

	ReadFolderOfProjectVersion(params *ReadFolderOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ReadFolderOfProjectVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListFolderOfProjectVersion lists
*/
func (a *Client) ListFolderOfProjectVersion(params *ListFolderOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListFolderOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFolderOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFolderOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListFolderOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListFolderOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listFolderOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadFolderOfProjectVersion reads
*/
func (a *Client) ReadFolderOfProjectVersion(params *ReadFolderOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ReadFolderOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadFolderOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readFolderOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/folders/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadFolderOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadFolderOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readFolderOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
