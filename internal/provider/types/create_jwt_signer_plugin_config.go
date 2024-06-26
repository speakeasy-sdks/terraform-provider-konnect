// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateJWTSignerPluginConfig struct {
	AccessTokenConsumerBy                   []types.String          `tfsdk:"access_token_consumer_by"`
	AccessTokenConsumerClaim                []types.String          `tfsdk:"access_token_consumer_claim"`
	AccessTokenIntrospectionAuthorization   types.String            `tfsdk:"access_token_introspection_authorization"`
	AccessTokenIntrospectionBodyArgs        types.String            `tfsdk:"access_token_introspection_body_args"`
	AccessTokenIntrospectionConsumerBy      []types.String          `tfsdk:"access_token_introspection_consumer_by"`
	AccessTokenIntrospectionConsumerClaim   []types.String          `tfsdk:"access_token_introspection_consumer_claim"`
	AccessTokenIntrospectionEndpoint        types.String            `tfsdk:"access_token_introspection_endpoint"`
	AccessTokenIntrospectionHint            types.String            `tfsdk:"access_token_introspection_hint"`
	AccessTokenIntrospectionJwtClaim        []types.String          `tfsdk:"access_token_introspection_jwt_claim"`
	AccessTokenIntrospectionLeeway          types.Number            `tfsdk:"access_token_introspection_leeway"`
	AccessTokenIntrospectionScopesClaim     []types.String          `tfsdk:"access_token_introspection_scopes_claim"`
	AccessTokenIntrospectionScopesRequired  []types.String          `tfsdk:"access_token_introspection_scopes_required"`
	AccessTokenIntrospectionTimeout         types.Number            `tfsdk:"access_token_introspection_timeout"`
	AccessTokenIssuer                       types.String            `tfsdk:"access_token_issuer"`
	AccessTokenJwksURI                      types.String            `tfsdk:"access_token_jwks_uri"`
	AccessTokenKeyset                       types.String            `tfsdk:"access_token_keyset"`
	AccessTokenLeeway                       types.Number            `tfsdk:"access_token_leeway"`
	AccessTokenOptional                     types.Bool              `tfsdk:"access_token_optional"`
	AccessTokenRequestHeader                types.String            `tfsdk:"access_token_request_header"`
	AccessTokenScopesClaim                  []types.String          `tfsdk:"access_token_scopes_claim"`
	AccessTokenScopesRequired               []types.String          `tfsdk:"access_token_scopes_required"`
	AccessTokenSigningAlgorithm             types.String            `tfsdk:"access_token_signing_algorithm"`
	AccessTokenUpstreamHeader               types.String            `tfsdk:"access_token_upstream_header"`
	AccessTokenUpstreamLeeway               types.Number            `tfsdk:"access_token_upstream_leeway"`
	AddClaims                               map[string]types.String `tfsdk:"add_claims"`
	CacheAccessTokenIntrospection           types.Bool              `tfsdk:"cache_access_token_introspection"`
	CacheChannelTokenIntrospection          types.Bool              `tfsdk:"cache_channel_token_introspection"`
	ChannelTokenConsumerBy                  []types.String          `tfsdk:"channel_token_consumer_by"`
	ChannelTokenConsumerClaim               []types.String          `tfsdk:"channel_token_consumer_claim"`
	ChannelTokenIntrospectionAuthorization  types.String            `tfsdk:"channel_token_introspection_authorization"`
	ChannelTokenIntrospectionBodyArgs       types.String            `tfsdk:"channel_token_introspection_body_args"`
	ChannelTokenIntrospectionConsumerBy     []types.String          `tfsdk:"channel_token_introspection_consumer_by"`
	ChannelTokenIntrospectionConsumerClaim  []types.String          `tfsdk:"channel_token_introspection_consumer_claim"`
	ChannelTokenIntrospectionEndpoint       types.String            `tfsdk:"channel_token_introspection_endpoint"`
	ChannelTokenIntrospectionHint           types.String            `tfsdk:"channel_token_introspection_hint"`
	ChannelTokenIntrospectionJwtClaim       []types.String          `tfsdk:"channel_token_introspection_jwt_claim"`
	ChannelTokenIntrospectionLeeway         types.Number            `tfsdk:"channel_token_introspection_leeway"`
	ChannelTokenIntrospectionScopesClaim    []types.String          `tfsdk:"channel_token_introspection_scopes_claim"`
	ChannelTokenIntrospectionScopesRequired []types.String          `tfsdk:"channel_token_introspection_scopes_required"`
	ChannelTokenIntrospectionTimeout        types.Number            `tfsdk:"channel_token_introspection_timeout"`
	ChannelTokenIssuer                      types.String            `tfsdk:"channel_token_issuer"`
	ChannelTokenJwksURI                     types.String            `tfsdk:"channel_token_jwks_uri"`
	ChannelTokenKeyset                      types.String            `tfsdk:"channel_token_keyset"`
	ChannelTokenLeeway                      types.Number            `tfsdk:"channel_token_leeway"`
	ChannelTokenOptional                    types.Bool              `tfsdk:"channel_token_optional"`
	ChannelTokenRequestHeader               types.String            `tfsdk:"channel_token_request_header"`
	ChannelTokenScopesClaim                 []types.String          `tfsdk:"channel_token_scopes_claim"`
	ChannelTokenScopesRequired              []types.String          `tfsdk:"channel_token_scopes_required"`
	ChannelTokenSigningAlgorithm            types.String            `tfsdk:"channel_token_signing_algorithm"`
	ChannelTokenUpstreamHeader              types.String            `tfsdk:"channel_token_upstream_header"`
	ChannelTokenUpstreamLeeway              types.Number            `tfsdk:"channel_token_upstream_leeway"`
	EnableAccessTokenIntrospection          types.Bool              `tfsdk:"enable_access_token_introspection"`
	EnableChannelTokenIntrospection         types.Bool              `tfsdk:"enable_channel_token_introspection"`
	EnableHsSignatures                      types.Bool              `tfsdk:"enable_hs_signatures"`
	EnableInstrumentation                   types.Bool              `tfsdk:"enable_instrumentation"`
	Realm                                   types.String            `tfsdk:"realm"`
	SetClaims                               map[string]types.String `tfsdk:"set_claims"`
	TrustAccessTokenIntrospection           types.Bool              `tfsdk:"trust_access_token_introspection"`
	TrustChannelTokenIntrospection          types.Bool              `tfsdk:"trust_channel_token_introspection"`
	VerifyAccessTokenExpiry                 types.Bool              `tfsdk:"verify_access_token_expiry"`
	VerifyAccessTokenIntrospectionExpiry    types.Bool              `tfsdk:"verify_access_token_introspection_expiry"`
	VerifyAccessTokenIntrospectionScopes    types.Bool              `tfsdk:"verify_access_token_introspection_scopes"`
	VerifyAccessTokenScopes                 types.Bool              `tfsdk:"verify_access_token_scopes"`
	VerifyAccessTokenSignature              types.Bool              `tfsdk:"verify_access_token_signature"`
	VerifyChannelTokenExpiry                types.Bool              `tfsdk:"verify_channel_token_expiry"`
	VerifyChannelTokenIntrospectionExpiry   types.Bool              `tfsdk:"verify_channel_token_introspection_expiry"`
	VerifyChannelTokenIntrospectionScopes   types.Bool              `tfsdk:"verify_channel_token_introspection_scopes"`
	VerifyChannelTokenScopes                types.Bool              `tfsdk:"verify_channel_token_scopes"`
	VerifyChannelTokenSignature             types.Bool              `tfsdk:"verify_channel_token_signature"`
}
