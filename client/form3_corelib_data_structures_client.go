// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/client/accounts"
	"github.com/ewilde/go-form3/client/ace"
	"github.com/ewilde/go-form3/client/associations"
	"github.com/ewilde/go-form3/client/limits"
	"github.com/ewilde/go-form3/client/organisations"
	"github.com/ewilde/go-form3/client/payments"
	"github.com/ewilde/go-form3/client/roles"
	"github.com/ewilde/go-form3/client/subscriptions"
	"github.com/ewilde/go-form3/client/users"
)

// Default form3 corelib data structures HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.form3.tech"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new form3 corelib data structures HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Form3CorelibDataStructures {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new form3 corelib data structures HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Form3CorelibDataStructures {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new form3 corelib data structures client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Form3CorelibDataStructures {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Form3CorelibDataStructures)
	cli.Transport = transport

	cli.Accounts = accounts.New(transport, formats)

	cli.Ace = ace.New(transport, formats)

	cli.Associations = associations.New(transport, formats)

	cli.Limits = limits.New(transport, formats)

	cli.Organisations = organisations.New(transport, formats)

	cli.Payments = payments.New(transport, formats)

	cli.Roles = roles.New(transport, formats)

	cli.Subscriptions = subscriptions.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Form3CorelibDataStructures is a client for form3 corelib data structures
type Form3CorelibDataStructures struct {
	Accounts *accounts.Client

	Ace *ace.Client

	Associations *associations.Client

	Limits *limits.Client

	Organisations *organisations.Client

	Payments *payments.Client

	Roles *roles.Client

	Subscriptions *subscriptions.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Form3CorelibDataStructures) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Accounts.SetTransport(transport)

	c.Ace.SetTransport(transport)

	c.Associations.SetTransport(transport)

	c.Limits.SetTransport(transport)

	c.Organisations.SetTransport(transport)

	c.Payments.SetTransport(transport)

	c.Roles.SetTransport(transport)

	c.Subscriptions.SetTransport(transport)

	c.Users.SetTransport(transport)

}
