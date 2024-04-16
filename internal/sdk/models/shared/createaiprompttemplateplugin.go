// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateAIPromptTemplatePluginProtocols string

const (
	CreateAIPromptTemplatePluginProtocolsGrpc           CreateAIPromptTemplatePluginProtocols = "grpc"
	CreateAIPromptTemplatePluginProtocolsGrpcs          CreateAIPromptTemplatePluginProtocols = "grpcs"
	CreateAIPromptTemplatePluginProtocolsHTTP           CreateAIPromptTemplatePluginProtocols = "http"
	CreateAIPromptTemplatePluginProtocolsHTTPS          CreateAIPromptTemplatePluginProtocols = "https"
	CreateAIPromptTemplatePluginProtocolsTCP            CreateAIPromptTemplatePluginProtocols = "tcp"
	CreateAIPromptTemplatePluginProtocolsTLS            CreateAIPromptTemplatePluginProtocols = "tls"
	CreateAIPromptTemplatePluginProtocolsTLSPassthrough CreateAIPromptTemplatePluginProtocols = "tls_passthrough"
	CreateAIPromptTemplatePluginProtocolsUDP            CreateAIPromptTemplatePluginProtocols = "udp"
	CreateAIPromptTemplatePluginProtocolsWs             CreateAIPromptTemplatePluginProtocols = "ws"
	CreateAIPromptTemplatePluginProtocolsWss            CreateAIPromptTemplatePluginProtocols = "wss"
)

func (e CreateAIPromptTemplatePluginProtocols) ToPointer() *CreateAIPromptTemplatePluginProtocols {
	return &e
}

func (e *CreateAIPromptTemplatePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateAIPromptTemplatePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAIPromptTemplatePluginProtocols: %v", v)
	}
}

// CreateAIPromptTemplatePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateAIPromptTemplatePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptTemplatePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAIPromptTemplatePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateAIPromptTemplatePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptTemplatePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAIPromptTemplatePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateAIPromptTemplatePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAIPromptTemplatePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Templates struct {
	// Unique name for the template, can be called with `{template://NAME}`
	Name *string `json:"name,omitempty"`
	// Template string for this request, supports mustache-style `{{"{{"}}placeholders}}`
	Template *string `json:"template,omitempty"`
}

func (o *Templates) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Templates) GetTemplate() *string {
	if o == nil {
		return nil
	}
	return o.Template
}

type CreateAIPromptTemplatePluginConfig struct {
	// Array of templates available to the request context.
	Templates []Templates `json:"templates,omitempty"`
	// Set true to allow requests that don't call or match any template.
	AllowUntemplatedRequests *bool `default:"true" json:"allow_untemplated_requests"`
	// Set true to add the original request to the Kong log plugin(s) output.
	LogOriginalRequest *bool `default:"false" json:"log_original_request"`
}

func (c CreateAIPromptTemplatePluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAIPromptTemplatePluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAIPromptTemplatePluginConfig) GetTemplates() []Templates {
	if o == nil {
		return nil
	}
	return o.Templates
}

func (o *CreateAIPromptTemplatePluginConfig) GetAllowUntemplatedRequests() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUntemplatedRequests
}

func (o *CreateAIPromptTemplatePluginConfig) GetLogOriginalRequest() *bool {
	if o == nil {
		return nil
	}
	return o.LogOriginalRequest
}

// CreateAIPromptTemplatePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateAIPromptTemplatePlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"ai-prompt-template" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateAIPromptTemplatePluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateAIPromptTemplatePluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateAIPromptTemplatePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateAIPromptTemplatePluginService `json:"service,omitempty"`
	Config  CreateAIPromptTemplatePluginConfig   `json:"config"`
}

func (c CreateAIPromptTemplatePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAIPromptTemplatePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAIPromptTemplatePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateAIPromptTemplatePlugin) GetName() string {
	return "ai-prompt-template"
}

func (o *CreateAIPromptTemplatePlugin) GetProtocols() []CreateAIPromptTemplatePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateAIPromptTemplatePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateAIPromptTemplatePlugin) GetConsumer() *CreateAIPromptTemplatePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateAIPromptTemplatePlugin) GetRoute() *CreateAIPromptTemplatePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateAIPromptTemplatePlugin) GetService() *CreateAIPromptTemplatePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateAIPromptTemplatePlugin) GetConfig() CreateAIPromptTemplatePluginConfig {
	if o == nil {
		return CreateAIPromptTemplatePluginConfig{}
	}
	return o.Config
}