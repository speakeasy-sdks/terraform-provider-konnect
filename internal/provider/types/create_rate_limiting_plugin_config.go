// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateRateLimitingPluginConfig struct {
	Day               types.Number `tfsdk:"day"`
	ErrorCode         types.Number `tfsdk:"error_code"`
	ErrorMessage      types.String `tfsdk:"error_message"`
	FaultTolerant     types.Bool   `tfsdk:"fault_tolerant"`
	HeaderName        types.String `tfsdk:"header_name"`
	HideClientHeaders types.Bool   `tfsdk:"hide_client_headers"`
	Hour              types.Number `tfsdk:"hour"`
	LimitBy           types.String `tfsdk:"limit_by"`
	Minute            types.Number `tfsdk:"minute"`
	Month             types.Number `tfsdk:"month"`
	Path              types.String `tfsdk:"path"`
	Policy            types.String `tfsdk:"policy"`
	RedisDatabase     types.Int64  `tfsdk:"redis_database"`
	RedisHost         types.String `tfsdk:"redis_host"`
	RedisPassword     types.String `tfsdk:"redis_password"`
	RedisPort         types.Int64  `tfsdk:"redis_port"`
	RedisServerName   types.String `tfsdk:"redis_server_name"`
	RedisSsl          types.Bool   `tfsdk:"redis_ssl"`
	RedisSslVerify    types.Bool   `tfsdk:"redis_ssl_verify"`
	RedisTimeout      types.Number `tfsdk:"redis_timeout"`
	RedisUsername     types.String `tfsdk:"redis_username"`
	Second            types.Number `tfsdk:"second"`
	SyncRate          types.Number `tfsdk:"sync_rate"`
	Year              types.Number `tfsdk:"year"`
}
