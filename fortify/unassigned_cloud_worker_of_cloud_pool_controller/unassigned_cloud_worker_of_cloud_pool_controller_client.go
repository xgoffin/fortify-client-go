// Code generated by go-swagger; DO NOT EDIT.

package unassigned_cloud_worker_of_cloud_pool_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new unassigned cloud worker of cloud pool controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for unassigned cloud worker of cloud pool controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListUnassignedCloudWorkerOfCloudPool(params *ListUnassignedCloudWorkerOfCloudPoolParams, authInfo runtime.ClientAuthInfoWriter) (*ListUnassignedCloudWorkerOfCloudPoolOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListUnassignedCloudWorkerOfCloudPool lists
*/
func (a *Client) ListUnassignedCloudWorkerOfCloudPool(params *ListUnassignedCloudWorkerOfCloudPoolParams, authInfo runtime.ClientAuthInfoWriter) (*ListUnassignedCloudWorkerOfCloudPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUnassignedCloudWorkerOfCloudPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUnassignedCloudWorkerOfCloudPool",
		Method:             "GET",
		PathPattern:        "/cloudpools/disabledWorkers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListUnassignedCloudWorkerOfCloudPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUnassignedCloudWorkerOfCloudPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUnassignedCloudWorkerOfCloudPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
