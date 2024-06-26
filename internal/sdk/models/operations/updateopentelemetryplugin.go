// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateOpentelemetryPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// Description of the Plugin
	CreateOpentelemetryPlugin shared.CreateOpentelemetryPlugin `request:"mediaType=application/json"`
}

func (o *UpdateOpentelemetryPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateOpentelemetryPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateOpentelemetryPluginRequest) GetCreateOpentelemetryPlugin() shared.CreateOpentelemetryPlugin {
	if o == nil {
		return shared.CreateOpentelemetryPlugin{}
	}
	return o.CreateOpentelemetryPlugin
}

type UpdateOpentelemetryPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Plugin
	OpentelemetryPlugin *shared.OpentelemetryPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateOpentelemetryPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateOpentelemetryPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateOpentelemetryPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateOpentelemetryPluginResponse) GetOpentelemetryPlugin() *shared.OpentelemetryPlugin {
	if o == nil {
		return nil
	}
	return o.OpentelemetryPlugin
}

func (o *UpdateOpentelemetryPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
