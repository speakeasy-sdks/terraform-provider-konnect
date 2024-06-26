// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateJQPluginProtocols string

const (
	CreateJQPluginProtocolsGrpc           CreateJQPluginProtocols = "grpc"
	CreateJQPluginProtocolsGrpcs          CreateJQPluginProtocols = "grpcs"
	CreateJQPluginProtocolsHTTP           CreateJQPluginProtocols = "http"
	CreateJQPluginProtocolsHTTPS          CreateJQPluginProtocols = "https"
	CreateJQPluginProtocolsTCP            CreateJQPluginProtocols = "tcp"
	CreateJQPluginProtocolsTLS            CreateJQPluginProtocols = "tls"
	CreateJQPluginProtocolsTLSPassthrough CreateJQPluginProtocols = "tls_passthrough"
	CreateJQPluginProtocolsUDP            CreateJQPluginProtocols = "udp"
	CreateJQPluginProtocolsWs             CreateJQPluginProtocols = "ws"
	CreateJQPluginProtocolsWss            CreateJQPluginProtocols = "wss"
)

func (e CreateJQPluginProtocols) ToPointer() *CreateJQPluginProtocols {
	return &e
}

func (e *CreateJQPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateJQPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateJQPluginProtocols: %v", v)
	}
}

// CreateJQPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateJQPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJQPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateJQPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJQPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateJQPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJQPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestJqProgramOptions struct {
	CompactOutput *bool `default:"true" json:"compact_output"`
	RawOutput     *bool `default:"false" json:"raw_output"`
	JoinOutput    *bool `default:"false" json:"join_output"`
	ASCIIOutput   *bool `default:"false" json:"ascii_output"`
	SortKeys      *bool `default:"false" json:"sort_keys"`
}

func (r RequestJqProgramOptions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestJqProgramOptions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestJqProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *RequestJqProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *RequestJqProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *RequestJqProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *RequestJqProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type ResponseJqProgramOptions struct {
	CompactOutput *bool `default:"true" json:"compact_output"`
	RawOutput     *bool `default:"false" json:"raw_output"`
	JoinOutput    *bool `default:"false" json:"join_output"`
	ASCIIOutput   *bool `default:"false" json:"ascii_output"`
	SortKeys      *bool `default:"false" json:"sort_keys"`
}

func (r ResponseJqProgramOptions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResponseJqProgramOptions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResponseJqProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *ResponseJqProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *ResponseJqProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *ResponseJqProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *ResponseJqProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type CreateJQPluginConfig struct {
	RequestJqProgram         *string                   `json:"request_jq_program,omitempty"`
	RequestJqProgramOptions  *RequestJqProgramOptions  `json:"request_jq_program_options,omitempty"`
	RequestIfMediaType       []string                  `json:"request_if_media_type,omitempty"`
	ResponseJqProgram        *string                   `json:"response_jq_program,omitempty"`
	ResponseJqProgramOptions *ResponseJqProgramOptions `json:"response_jq_program_options,omitempty"`
	ResponseIfMediaType      []string                  `json:"response_if_media_type,omitempty"`
	ResponseIfStatusCode     []int64                   `json:"response_if_status_code,omitempty"`
}

func (o *CreateJQPluginConfig) GetRequestJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.RequestJqProgram
}

func (o *CreateJQPluginConfig) GetRequestJqProgramOptions() *RequestJqProgramOptions {
	if o == nil {
		return nil
	}
	return o.RequestJqProgramOptions
}

func (o *CreateJQPluginConfig) GetRequestIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.RequestIfMediaType
}

func (o *CreateJQPluginConfig) GetResponseJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgram
}

func (o *CreateJQPluginConfig) GetResponseJqProgramOptions() *ResponseJqProgramOptions {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgramOptions
}

func (o *CreateJQPluginConfig) GetResponseIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.ResponseIfMediaType
}

func (o *CreateJQPluginConfig) GetResponseIfStatusCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseIfStatusCode
}

// CreateJQPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateJQPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"jq" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateJQPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateJQPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateJQPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateJQPluginService `json:"service,omitempty"`
	Config  CreateJQPluginConfig   `json:"config"`
}

func (c CreateJQPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateJQPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateJQPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateJQPlugin) GetName() string {
	return "jq"
}

func (o *CreateJQPlugin) GetProtocols() []CreateJQPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateJQPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateJQPlugin) GetConsumer() *CreateJQPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateJQPlugin) GetRoute() *CreateJQPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateJQPlugin) GetService() *CreateJQPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateJQPlugin) GetConfig() CreateJQPluginConfig {
	if o == nil {
		return CreateJQPluginConfig{}
	}
	return o.Config
}
