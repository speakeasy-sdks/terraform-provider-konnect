// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type CreateResponseTransformerPluginConfig struct {
	Add     *CreateResponseTransformerPluginAdd    `tfsdk:"add"`
	Append  *CreateResponseTransformerPluginAdd    `tfsdk:"append"`
	Remove  *CreateResponseTransformerPluginRemove `tfsdk:"remove"`
	Rename  *CreateResponseTransformerPluginRename `tfsdk:"rename"`
	Replace *CreateResponseTransformerPluginAdd    `tfsdk:"replace"`
}