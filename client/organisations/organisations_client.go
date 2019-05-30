// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organisations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organisations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteUnitsID deletes organisation
*/
func (a *Client) DeleteUnitsID(params *DeleteUnitsIDParams) (*DeleteUnitsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUnitsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUnitsID",
		Method:             "DELETE",
		PathPattern:        "/units/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUnitsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUnitsIDNoContent), nil

}

/*
GetUnits lists all organisations
*/
func (a *Client) GetUnits(params *GetUnitsParams) (*GetUnitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUnits",
		Method:             "GET",
		PathPattern:        "/units",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUnitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUnitsOK), nil

}

/*
GetUnitsID fetches organisation
*/
func (a *Client) GetUnitsID(params *GetUnitsIDParams) (*GetUnitsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUnitsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUnitsID",
		Method:             "GET",
		PathPattern:        "/units/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUnitsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUnitsIDOK), nil

}

/*
PatchUnitsID updates organisation
*/
func (a *Client) PatchUnitsID(params *PatchUnitsIDParams) (*PatchUnitsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchUnitsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchUnitsID",
		Method:             "PATCH",
		PathPattern:        "/units/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUnitsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchUnitsIDOK), nil

}

/*
PostUnits creates organisation
*/
func (a *Client) PostUnits(params *PostUnitsParams) (*PostUnitsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUnits",
		Method:             "POST",
		PathPattern:        "/units",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUnitsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostUnitsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
