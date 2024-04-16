// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateKeyAuthPluginProtocols string

const (
	CreateKeyAuthPluginProtocolsGrpc           CreateKeyAuthPluginProtocols = "grpc"
	CreateKeyAuthPluginProtocolsGrpcs          CreateKeyAuthPluginProtocols = "grpcs"
	CreateKeyAuthPluginProtocolsHTTP           CreateKeyAuthPluginProtocols = "http"
	CreateKeyAuthPluginProtocolsHTTPS          CreateKeyAuthPluginProtocols = "https"
	CreateKeyAuthPluginProtocolsTCP            CreateKeyAuthPluginProtocols = "tcp"
	CreateKeyAuthPluginProtocolsTLS            CreateKeyAuthPluginProtocols = "tls"
	CreateKeyAuthPluginProtocolsTLSPassthrough CreateKeyAuthPluginProtocols = "tls_passthrough"
	CreateKeyAuthPluginProtocolsUDP            CreateKeyAuthPluginProtocols = "udp"
	CreateKeyAuthPluginProtocolsWs             CreateKeyAuthPluginProtocols = "ws"
	CreateKeyAuthPluginProtocolsWss            CreateKeyAuthPluginProtocols = "wss"
)

func (e CreateKeyAuthPluginProtocols) ToPointer() *CreateKeyAuthPluginProtocols {
	return &e
}

func (e *CreateKeyAuthPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateKeyAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateKeyAuthPluginProtocols: %v", v)
	}
}

// CreateKeyAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateKeyAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateKeyAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateKeyAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateKeyAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateKeyAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateKeyAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateKeyAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateKeyAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateKeyAuthPluginConfig struct {
	// Describes an array of parameter names where the plugin will look for a key. The key names may only contain [a-z], [A-Z], [0-9], [_] underscore, and [-] hyphen.
	KeyNames []string `json:"key_names,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin strips the credential from the request.
	HideCredentials *bool `default:"false" json:"hide_credentials"`
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure `4xx`.
	Anonymous *string `json:"anonymous,omitempty"`
	// If enabled (default), the plugin reads the request header and tries to find the key in it.
	KeyInHeader *bool `default:"true" json:"key_in_header"`
	// If enabled (default), the plugin reads the query parameter in the request and tries to find the key in it.
	KeyInQuery *bool `default:"true" json:"key_in_query"`
	// If enabled, the plugin reads the request body. Supported MIME types: `application/www-form-urlencoded`, `application/json`, and `multipart/form-data`.
	KeyInBody *bool `default:"false" json:"key_in_body"`
	// A boolean value that indicates whether the plugin should run (and try to authenticate) on `OPTIONS` preflight requests. If set to `false`, then `OPTIONS` requests are always allowed.
	RunOnPreflight *bool `default:"true" json:"run_on_preflight"`
}

func (c CreateKeyAuthPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateKeyAuthPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateKeyAuthPluginConfig) GetKeyNames() []string {
	if o == nil {
		return nil
	}
	return o.KeyNames
}

func (o *CreateKeyAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *CreateKeyAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *CreateKeyAuthPluginConfig) GetKeyInHeader() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInHeader
}

func (o *CreateKeyAuthPluginConfig) GetKeyInQuery() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInQuery
}

func (o *CreateKeyAuthPluginConfig) GetKeyInBody() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInBody
}

func (o *CreateKeyAuthPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

// CreateKeyAuthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateKeyAuthPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"key-auth" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateKeyAuthPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateKeyAuthPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateKeyAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateKeyAuthPluginService `json:"service,omitempty"`
	Config  CreateKeyAuthPluginConfig   `json:"config"`
}

func (c CreateKeyAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateKeyAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateKeyAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateKeyAuthPlugin) GetName() string {
	return "key-auth"
}

func (o *CreateKeyAuthPlugin) GetProtocols() []CreateKeyAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateKeyAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateKeyAuthPlugin) GetConsumer() *CreateKeyAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateKeyAuthPlugin) GetRoute() *CreateKeyAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateKeyAuthPlugin) GetService() *CreateKeyAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateKeyAuthPlugin) GetConfig() CreateKeyAuthPluginConfig {
	if o == nil {
		return CreateKeyAuthPluginConfig{}
	}
	return o.Config
}