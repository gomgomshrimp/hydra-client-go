// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/client/metadata"
	"github.com/ory/hydra-client-go/client/public"
)

// Default ory hydra HTTP client.
var Default = NewHTTPClient(nil, logrus.New())

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new ory hydra HTTP client.
func NewHTTPClient(formats strfmt.Registry, logger *logrus.Logger) *OryHydra {
	return NewHTTPClientWithConfig(formats, nil, logger)
}

// NewHTTPClientWithConfig creates a new ory hydra HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig, logger *logrus.Logger) *OryHydra {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats, logger)
}

// New creates a new ory hydra client
func New(transport runtime.ClientTransport, formats strfmt.Registry, logger *logrus.Logger) *OryHydra {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(OryHydra)
	cli.Transport = transport
	cli.Admin = admin.New(transport, formats, logger)
	cli.Metadata = metadata.New(transport, formats)
	cli.Public = public.New(transport, formats)
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

// OryHydra is a client for ory hydra
type OryHydra struct {
	Admin admin.ClientService

	Metadata metadata.ClientService

	Public public.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OryHydra) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Admin.SetTransport(transport)
	c.Metadata.SetTransport(transport)
	c.Public.SetTransport(transport)
}
