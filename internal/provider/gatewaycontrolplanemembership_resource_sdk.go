// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayControlPlaneMembershipResourceModel) ToSharedGroupMembership() *shared.GroupMembership {
	var members []shared.Members = []shared.Members{}
	for _, membersItem := range r.Members {
		id := new(string)
		if !membersItem.ID.IsUnknown() && !membersItem.ID.IsNull() {
			*id = membersItem.ID.ValueString()
		} else {
			id = nil
		}
		members = append(members, shared.Members{
			ID: id,
		})
	}
	out := shared.GroupMembership{
		Members: members,
	}
	return &out
}