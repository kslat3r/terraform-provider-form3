// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new direct debits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for direct debits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDirectdebits lists direct debits
*/
func (a *Client) GetDirectdebits(params *GetDirectdebitsParams) (*GetDirectdebitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDirectdebitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectdebits",
		Method:             "GET",
		PathPattern:        "/directdebits",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDirectdebitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectdebitsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
