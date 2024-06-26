// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type AIPromptGuardPluginProtocols string

const (
	AIPromptGuardPluginProtocolsGrpc           AIPromptGuardPluginProtocols = "grpc"
	AIPromptGuardPluginProtocolsGrpcs          AIPromptGuardPluginProtocols = "grpcs"
	AIPromptGuardPluginProtocolsHTTP           AIPromptGuardPluginProtocols = "http"
	AIPromptGuardPluginProtocolsHTTPS          AIPromptGuardPluginProtocols = "https"
	AIPromptGuardPluginProtocolsTCP            AIPromptGuardPluginProtocols = "tcp"
	AIPromptGuardPluginProtocolsTLS            AIPromptGuardPluginProtocols = "tls"
	AIPromptGuardPluginProtocolsTLSPassthrough AIPromptGuardPluginProtocols = "tls_passthrough"
	AIPromptGuardPluginProtocolsUDP            AIPromptGuardPluginProtocols = "udp"
	AIPromptGuardPluginProtocolsWs             AIPromptGuardPluginProtocols = "ws"
	AIPromptGuardPluginProtocolsWss            AIPromptGuardPluginProtocols = "wss"
)

func (e AIPromptGuardPluginProtocols) ToPointer() *AIPromptGuardPluginProtocols {
	return &e
}

func (e *AIPromptGuardPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AIPromptGuardPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AIPromptGuardPluginProtocols: %v", v)
	}
}

// AIPromptGuardPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AIPromptGuardPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptGuardPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AIPromptGuardPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AIPromptGuardPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptGuardPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AIPromptGuardPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AIPromptGuardPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AIPromptGuardPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AIPromptGuardPluginConfig struct {
	// Array of valid patterns, or valid questions from the 'user' role in chat.
	AllowPatterns []string `json:"allow_patterns,omitempty"`
	// Array of invalid patterns, or invalid questions from the 'user' role in chat.
	DenyPatterns []string `json:"deny_patterns,omitempty"`
	// If true, will ignore all previous chat prompts from the conversation history.
	AllowAllConversationHistory *bool `default:"false" json:"allow_all_conversation_history"`
}

func (a AIPromptGuardPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AIPromptGuardPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AIPromptGuardPluginConfig) GetAllowPatterns() []string {
	if o == nil {
		return nil
	}
	return o.AllowPatterns
}

func (o *AIPromptGuardPluginConfig) GetDenyPatterns() []string {
	if o == nil {
		return nil
	}
	return o.DenyPatterns
}

func (o *AIPromptGuardPluginConfig) GetAllowAllConversationHistory() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAllConversationHistory
}

// AIPromptGuardPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AIPromptGuardPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"ai-prompt-guard" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AIPromptGuardPluginProtocols `json:"protocols"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *AIPromptGuardPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AIPromptGuardPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AIPromptGuardPluginService `json:"service,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64                    `json:"created_at,omitempty"`
	ID        *string                   `json:"id,omitempty"`
	Config    AIPromptGuardPluginConfig `json:"config"`
}

func (a AIPromptGuardPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AIPromptGuardPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AIPromptGuardPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AIPromptGuardPlugin) GetName() string {
	return "ai-prompt-guard"
}

func (o *AIPromptGuardPlugin) GetProtocols() []AIPromptGuardPluginProtocols {
	if o == nil {
		return []AIPromptGuardPluginProtocols{}
	}
	return o.Protocols
}

func (o *AIPromptGuardPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AIPromptGuardPlugin) GetConsumer() *AIPromptGuardPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AIPromptGuardPlugin) GetRoute() *AIPromptGuardPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AIPromptGuardPlugin) GetService() *AIPromptGuardPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *AIPromptGuardPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AIPromptGuardPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AIPromptGuardPlugin) GetConfig() AIPromptGuardPluginConfig {
	if o == nil {
		return AIPromptGuardPluginConfig{}
	}
	return o.Config
}
