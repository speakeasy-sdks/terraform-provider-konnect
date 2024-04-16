// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type GetSniRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the SNI to lookup
	SNIID string `pathParam:"style=simple,explode=false,name=SNIId"`
}

func (o *GetSniRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetSniRequest) GetSNIID() string {
	if o == nil {
		return ""
	}
	return o.SNIID
}

type GetSniResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched SNI
	Sni *shared.Sni
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *GetSniResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSniResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSniResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSniResponse) GetSni() *shared.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}

func (o *GetSniResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}