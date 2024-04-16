// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateRatelimitingadvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// Description of the Plugin
	CreateRateLimitingAdvancedPlugin shared.CreateRateLimitingAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateRatelimitingadvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateRatelimitingadvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRatelimitingadvancedPluginRequest) GetCreateRateLimitingAdvancedPlugin() shared.CreateRateLimitingAdvancedPlugin {
	if o == nil {
		return shared.CreateRateLimitingAdvancedPlugin{}
	}
	return o.CreateRateLimitingAdvancedPlugin
}

type UpdateRatelimitingadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Plugin
	RateLimitingAdvancedPlugin *shared.RateLimitingAdvancedPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetRateLimitingAdvancedPlugin() *shared.RateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.RateLimitingAdvancedPlugin
}

func (o *UpdateRatelimitingadvancedPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}