// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type AIPromptTemplatePluginProtocols string

const (
	AIPromptTemplatePluginProtocolsGrpc           AIPromptTemplatePluginProtocols = "grpc"
	AIPromptTemplatePluginProtocolsGrpcs          AIPromptTemplatePluginProtocols = "grpcs"
	AIPromptTemplatePluginProtocolsHTTP           AIPromptTemplatePluginProtocols = "http"
	AIPromptTemplatePluginProtocolsHTTPS          AIPromptTemplatePluginProtocols = "https"
	AIPromptTemplatePluginProtocolsTCP            AIPromptTemplatePluginProtocols = "tcp"
	AIPromptTemplatePluginProtocolsTLS            AIPromptTemplatePluginProtocols = "tls"
	AIPromptTemplatePluginProtocolsTLSPassthrough AIPromptTemplatePluginProtocols = "tls_passthrough"
	AIPromptTemplatePluginProtocolsUDP            AIPromptTemplatePluginProtocols = "udp"
	AIPromptTemplatePluginProtocolsWs             AIPromptTemplatePluginProtocols = "ws"
	AIPromptTemplatePluginProtocolsWss            AIPromptTemplatePluginProtocols = "wss"
)

func (e AIPromptTemplatePluginProtocols) ToPointer() *AIPromptTemplatePluginProtocols {
	return &e
}

func (e *AIPromptTemplatePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AIPromptTemplatePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AIPromptTemplatePluginProtocols: %v", v)
	}
}

// AIPromptTemplatePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AIPromptTemplatePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptTemplatePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AIPromptTemplatePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AIPromptTemplatePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptTemplatePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AIPromptTemplatePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AIPromptTemplatePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptTemplatePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AIPromptTemplatePluginTemplates struct {
	// Unique name for the template, can be called with `{template://NAME}`
	Name *string `json:"name,omitempty"`
	// Template string for this request, supports mustache-style `{{"{{"}}placeholders}}`
	Template *string `json:"template,omitempty"`
}

func (o *AIPromptTemplatePluginTemplates) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AIPromptTemplatePluginTemplates) GetTemplate() *string {
	if o == nil {
		return nil
	}
	return o.Template
}

type AIPromptTemplatePluginConfig struct {
	// Array of templates available to the request context.
	Templates []AIPromptTemplatePluginTemplates `json:"templates,omitempty"`
	// Set true to allow requests that don't call or match any template.
	AllowUntemplatedRequests *bool `default:"true" json:"allow_untemplated_requests"`
	// Set true to add the original request to the Kong log plugin(s) output.
	LogOriginalRequest *bool `default:"false" json:"log_original_request"`
}

func (a AIPromptTemplatePluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AIPromptTemplatePluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AIPromptTemplatePluginConfig) GetTemplates() []AIPromptTemplatePluginTemplates {
	if o == nil {
		return nil
	}
	return o.Templates
}

func (o *AIPromptTemplatePluginConfig) GetAllowUntemplatedRequests() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUntemplatedRequests
}

func (o *AIPromptTemplatePluginConfig) GetLogOriginalRequest() *bool {
	if o == nil {
		return nil
	}
	return o.LogOriginalRequest
}

// AIPromptTemplatePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AIPromptTemplatePlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"ai-prompt-template" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AIPromptTemplatePluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AIPromptTemplatePluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AIPromptTemplatePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AIPromptTemplatePluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64                       `json:"created_at,omitempty"`
	ID        *string                      `json:"id,omitempty"`
	Config    AIPromptTemplatePluginConfig `json:"config"`
}

func (a AIPromptTemplatePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AIPromptTemplatePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AIPromptTemplatePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AIPromptTemplatePlugin) GetName() string {
	return "ai-prompt-template"
}

func (o *AIPromptTemplatePlugin) GetProtocols() []AIPromptTemplatePluginProtocols {
	if o == nil {
		return []AIPromptTemplatePluginProtocols{}
	}
	return o.Protocols
}

func (o *AIPromptTemplatePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AIPromptTemplatePlugin) GetConsumer() *AIPromptTemplatePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AIPromptTemplatePlugin) GetRoute() *AIPromptTemplatePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AIPromptTemplatePlugin) GetService() *AIPromptTemplatePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AIPromptTemplatePlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AIPromptTemplatePlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AIPromptTemplatePlugin) GetConfig() AIPromptTemplatePluginConfig {
	if o == nil {
		return AIPromptTemplatePluginConfig{}
	}
	return o.Config
}
