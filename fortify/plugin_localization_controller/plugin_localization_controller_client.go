// Code generated by go-swagger; DO NOT EDIT.

package plugin_localization_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plugin localization controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugin localization controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ReadPluginLocalization(params *ReadPluginLocalizationParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPluginLocalizationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ReadPluginLocalization reads
*/
func (a *Client) ReadPluginLocalization(params *ReadPluginLocalizationParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPluginLocalizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadPluginLocalizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readPluginLocalization",
		Method:             "GET",
		PathPattern:        "/pluginlocalization",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadPluginLocalizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadPluginLocalizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readPluginLocalization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
