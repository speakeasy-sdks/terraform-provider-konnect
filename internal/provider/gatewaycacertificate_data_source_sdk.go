// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayCACertificateDataSourceModel) RefreshFromSharedCACertificate(resp *shared.CACertificate) {
	if resp != nil {
		r.Cert = types.StringValue(resp.Cert)
		r.CertDigest = types.StringPointerValue(resp.CertDigest)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}