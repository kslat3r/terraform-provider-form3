// Code generated by go-swagger; DO NOT EDIT.

package account_routings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account routings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account routings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAccountRoutingsID deletes account routing
*/
func (a *Client) DeleteAccountRoutingsID(params *DeleteAccountRoutingsIDParams) (*DeleteAccountRoutingsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountRoutingsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAccountRoutingsID",
		Method:             "DELETE",
		PathPattern:        "/account_routings/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccountRoutingsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccountRoutingsIDNoContent), nil

}

/*
GetAccountRoutings lists account routings
*/
func (a *Client) GetAccountRoutings(params *GetAccountRoutingsParams) (*GetAccountRoutingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountRoutingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountRoutings",
		Method:             "GET",
		PathPattern:        "/account_routings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountRoutingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountRoutingsOK), nil

}

/*
GetAccountRoutingsID fetches account routing
*/
func (a *Client) GetAccountRoutingsID(params *GetAccountRoutingsIDParams) (*GetAccountRoutingsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountRoutingsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountRoutingsID",
		Method:             "GET",
		PathPattern:        "/account_routings/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountRoutingsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountRoutingsIDOK), nil

}

/*
PostAccountRoutings creates an account routing entry
*/
func (a *Client) PostAccountRoutings(params *PostAccountRoutingsParams) (*PostAccountRoutingsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccountRoutingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAccountRoutings",
		Method:             "POST",
		PathPattern:        "/account_routings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccountRoutingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAccountRoutingsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
