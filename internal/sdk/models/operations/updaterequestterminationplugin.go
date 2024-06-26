// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateRequestterminationPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// Description of the Plugin
	CreateRequestTerminationPlugin shared.CreateRequestTerminationPlugin `request:"mediaType=application/json"`
}

func (o *UpdateRequestterminationPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateRequestterminationPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRequestterminationPluginRequest) GetCreateRequestTerminationPlugin() shared.CreateRequestTerminationPlugin {
	if o == nil {
		return shared.CreateRequestTerminationPlugin{}
	}
	return o.CreateRequestTerminationPlugin
}

type UpdateRequestterminationPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Plugin
	RequestTerminationPlugin *shared.RequestTerminationPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateRequestterminationPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRequestterminationPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRequestterminationPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRequestterminationPluginResponse) GetRequestTerminationPlugin() *shared.RequestTerminationPlugin {
	if o == nil {
		return nil
	}
	return o.RequestTerminationPlugin
}

func (o *UpdateRequestterminationPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
