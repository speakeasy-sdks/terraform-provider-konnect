// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Healthchecks struct {
	Active    *Active      `tfsdk:"active"`
	Passive   *Passive     `tfsdk:"passive"`
	Threshold types.Number `tfsdk:"threshold"`
}
