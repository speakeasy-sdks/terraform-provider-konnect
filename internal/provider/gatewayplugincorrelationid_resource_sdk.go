// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginCorrelationIDResourceModel) ToSharedCreateCorrelationIDPlugin() *shared.CreateCorrelationIDPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateCorrelationIDPluginProtocols = []shared.CreateCorrelationIDPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateCorrelationIDPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateCorrelationIDPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateCorrelationIDPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateCorrelationIDPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateCorrelationIDPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateCorrelationIDPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateCorrelationIDPluginService{
			ID: id2,
		}
	}
	headerName := new(string)
	if !r.Config.HeaderName.IsUnknown() && !r.Config.HeaderName.IsNull() {
		*headerName = r.Config.HeaderName.ValueString()
	} else {
		headerName = nil
	}
	generator := new(shared.Generator)
	if !r.Config.Generator.IsUnknown() && !r.Config.Generator.IsNull() {
		*generator = shared.Generator(r.Config.Generator.ValueString())
	} else {
		generator = nil
	}
	echoDownstream := new(bool)
	if !r.Config.EchoDownstream.IsUnknown() && !r.Config.EchoDownstream.IsNull() {
		*echoDownstream = r.Config.EchoDownstream.ValueBool()
	} else {
		echoDownstream = nil
	}
	config := shared.CreateCorrelationIDPluginConfig{
		HeaderName:     headerName,
		Generator:      generator,
		EchoDownstream: echoDownstream,
	}
	out := shared.CreateCorrelationIDPlugin{
		Enabled:   enabled,
		Protocols: protocols,
		Tags:      tags,
		Consumer:  consumer,
		Route:     route,
		Service:   service,
		Config:    config,
	}
	return &out
}

func (r *GatewayPluginCorrelationIDResourceModel) RefreshFromSharedCorrelationIDPlugin(resp *shared.CorrelationIDPlugin) {
	if resp != nil {
		r.Config.EchoDownstream = types.BoolPointerValue(resp.Config.EchoDownstream)
		if resp.Config.Generator != nil {
			r.Config.Generator = types.StringValue(string(*resp.Config.Generator))
		} else {
			r.Config.Generator = types.StringNull()
		}
		r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}