// Code generated by go-swagger; DO NOT EDIT.

package artifact_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new artifact controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for artifact controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ApproveArtifact(params *ApproveArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*ApproveArtifactOK, error)

	DeleteArtifact(params *DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteArtifactOK, error)

	PurgeArtifact(params *PurgeArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*PurgeArtifactOK, error)

	ReadArtifact(params *ReadArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*ReadArtifactOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ApproveArtifact approves the artifact for processing in spite of failing
*/
func (a *Client) ApproveArtifact(params *ApproveArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*ApproveArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApproveArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "approveArtifact",
		Method:             "POST",
		PathPattern:        "/artifacts/action/approve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApproveArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApproveArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for approveArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteArtifact deletes
*/
func (a *Client) DeleteArtifact(params *DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteArtifact",
		Method:             "DELETE",
		PathPattern:        "/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurgeArtifact purges the specified artifact from the system to recover space without affecting issue metrics use the delete operation instead if you want to completely revert all traces of an artifact
*/
func (a *Client) PurgeArtifact(params *PurgeArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*PurgeArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurgeArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "purgeArtifact",
		Method:             "POST",
		PathPattern:        "/artifacts/action/purge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PurgeArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PurgeArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for purgeArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadArtifact reads
*/
func (a *Client) ReadArtifact(params *ReadArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*ReadArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readArtifact",
		Method:             "GET",
		PathPattern:        "/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadArtifactOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readArtifact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
