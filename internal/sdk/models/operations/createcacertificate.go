// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateCaCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the new CA Certificate for creation
	CreateCACertificate shared.CreateCACertificate `request:"mediaType=application/json"`
}

func (o *CreateCaCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateCaCertificateRequest) GetCreateCACertificate() shared.CreateCACertificate {
	if o == nil {
		return shared.CreateCACertificate{}
	}
	return o.CreateCACertificate
}

type CreateCaCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created CA Certificate
	CACertificate *shared.CACertificate
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateCaCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCaCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCaCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCaCertificateResponse) GetCACertificate() *shared.CACertificate {
	if o == nil {
		return nil
	}
	return o.CACertificate
}

func (o *CreateCaCertificateResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}