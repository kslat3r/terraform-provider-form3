// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new mandates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mandates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMandates lists mandates
*/
func (a *Client) GetMandates(params *GetMandatesParams) (*GetMandatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMandatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMandates",
		Method:             "GET",
		PathPattern:        "/mandates",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMandatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMandatesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
