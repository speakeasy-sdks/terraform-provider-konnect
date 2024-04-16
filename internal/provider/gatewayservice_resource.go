// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayServiceResource{}
var _ resource.ResourceWithImportState = &GatewayServiceResource{}

func NewGatewayServiceResource() resource.Resource {
	return &GatewayServiceResource{}
}

// GatewayServiceResource defines the resource implementation.
type GatewayServiceResource struct {
	client *sdk.Konnect
}

// GatewayServiceResourceModel describes the resource data model.
type GatewayServiceResourceModel struct {
	CaCertificates    []types.String       `tfsdk:"ca_certificates"`
	ClientCertificate *tfTypes.ACLConsumer `tfsdk:"client_certificate"`
	ConnectTimeout    types.Int64          `tfsdk:"connect_timeout"`
	ControlPlaneID    types.String         `tfsdk:"control_plane_id"`
	CreatedAt         types.Int64          `tfsdk:"created_at"`
	Enabled           types.Bool           `tfsdk:"enabled"`
	Host              types.String         `tfsdk:"host"`
	ID                types.String         `tfsdk:"id"`
	Name              types.String         `tfsdk:"name"`
	Path              types.String         `tfsdk:"path"`
	Port              types.Int64          `tfsdk:"port"`
	Protocol          types.String         `tfsdk:"protocol"`
	ReadTimeout       types.Int64          `tfsdk:"read_timeout"`
	Retries           types.Int64          `tfsdk:"retries"`
	Tags              []types.String       `tfsdk:"tags"`
	TLSVerify         types.Bool           `tfsdk:"tls_verify"`
	TLSVerifyDepth    types.Int64          `tfsdk:"tls_verify_depth"`
	UpdatedAt         types.Int64          `tfsdk:"updated_at"`
	URL               types.String         `tfsdk:"url"`
	WriteTimeout      types.Int64          `tfsdk:"write_timeout"`
}

func (r *GatewayServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_service"
}

func (r *GatewayServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayService Resource",
		Attributes: map[string]schema.Attribute{
			"ca_certificates": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `Array of ` + "`" + `CA Certificate` + "`" + ` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to ` + "`" + `null` + "`" + ` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted). Requires replacement if changed. `,
			},
			"client_certificate": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `Requires replacement if changed. `,
					},
				},
				Description: `Certificate to be used as client certificate while TLS handshaking to the upstream server. Requires replacement if changed. `,
			},
			"connect_timeout": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(60000),
				Description: `The timeout in milliseconds for establishing a connection to the upstream server. Requires replacement if changed. ; Default: 60000`,
			},
			"control_plane_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed. `,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Description: `Whether the Service is active. If set to ` + "`" + `false` + "`" + `, the proxy behavior will be as if any routes attached to it do not exist (404). Default: ` + "`" + `true` + "`" + `. Requires replacement if changed. ; Default: true`,
			},
			"host": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The host of the upstream server. Note that the host value is case sensitive. Requires replacement if changed. `,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Service to lookup`,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The Service name. Requires replacement if changed. `,
			},
			"path": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The path to be used in requests to the upstream server. Requires replacement if changed. `,
			},
			"port": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(80),
				Description: `The upstream server port. Requires replacement if changed. ; Default: 80`,
			},
			"protocol": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("http"),
				Description: `The protocol used to communicate with the upstream. Requires replacement if changed. ; must be one of ["grpc", "grpcs", "http", "https", "tcp", "tls", "tls_passthrough", "udp", "ws", "wss"]; Default: "http"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"grpc",
						"grpcs",
						"http",
						"https",
						"tcp",
						"tls",
						"tls_passthrough",
						"udp",
						"ws",
						"wss",
					),
				},
			},
			"read_timeout": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(60000),
				Description: `The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server. Requires replacement if changed. ; Default: 60000`,
			},
			"retries": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: `The number of retries to execute upon failure to proxy. Requires replacement if changed. ; Default: 5`,
			},
			"tags": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Service for grouping and filtering. Requires replacement if changed. `,
			},
			"tls_verify": schema.BoolAttribute{
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Whether to enable verification of upstream server TLS certificate. If set to ` + "`" + `null` + "`" + `, then the Nginx default is respected. Requires replacement if changed. `,
			},
			"tls_verify_depth": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Maximum depth of chain while verifying Upstream server's TLS certificate. If set to ` + "`" + `null` + "`" + `, then the Nginx default is respected. Requires replacement if changed. `,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
			"url": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Helper field to set ` + "`" + `protocol` + "`" + `, ` + "`" + `host` + "`" + `, ` + "`" + `port` + "`" + ` and ` + "`" + `path` + "`" + ` using a URL. This field is write-only and is not returned in responses. Requires replacement if changed. `,
			},
			"write_timeout": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(60000),
				Description: `The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server. Requires replacement if changed. ; Default: 60000`,
			},
		},
	}
}

func (r *GatewayServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayServiceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	createService := *data.ToSharedCreateService()
	request := operations.CreateServiceRequest{
		ControlPlaneID: controlPlaneID,
		CreateService:  createService,
	}
	res, err := r.client.Services.CreateService(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Service == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedService(res.Service)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	controlPlaneId1 := data.ControlPlaneID.ValueString()
	serviceID := data.ID.ValueString()
	request1 := operations.GetServiceRequest{
		ControlPlaneID: controlPlaneId1,
		ServiceID:      serviceID,
	}
	res1, err := r.client.Services.GetService(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.Service == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedService(res1.Service)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayServiceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	serviceID := data.ID.ValueString()
	request := operations.GetServiceRequest{
		ControlPlaneID: controlPlaneID,
		ServiceID:      serviceID,
	}
	res, err := r.client.Services.GetService(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Service == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedService(res.Service)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayServiceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayServiceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	serviceID := data.ID.ValueString()
	request := operations.DeleteServiceRequest{
		ControlPlaneID: controlPlaneID,
		ServiceID:      serviceID,
	}
	res, err := r.client.Services.DeleteService(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GatewayServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_service. Reason: composite imports strings not supported.")
}