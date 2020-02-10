// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # OpenId.AudienceProtocolMapper
// 
// Allows for creating and managing audience protocol mappers within
// Keycloak. This mapper was added in Keycloak v4.6.0.Final.
// 
// Audience protocol mappers allow you add audiences to the `aud` claim
// within issued tokens. The audience can be a custom string, or it can be
// mapped to the ID of a pre-existing client.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `includedClientAudience` - (Required if `includedCustomAudience` is not specified) A client ID to include within the token's `aud` claim.
// - `includedCustomAudience` - (Required if `includedClientAudience` is not specified) A custom audience to include within the token's `aud` claim.
// - `addToIdToken` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_audience_protocol_mapper.html.markdown.
type AudienceProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewAudienceProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewAudienceProtocolMapper(ctx *pulumi.Context,
	name string, args *AudienceProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*AudienceProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addToAccessToken"] = nil
		inputs["addToIdToken"] = nil
		inputs["clientId"] = nil
		inputs["clientScopeId"] = nil
		inputs["includedClientAudience"] = nil
		inputs["includedCustomAudience"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["addToAccessToken"] = args.AddToAccessToken
		inputs["addToIdToken"] = args.AddToIdToken
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["includedClientAudience"] = args.IncludedClientAudience
		inputs["includedCustomAudience"] = args.IncludedCustomAudience
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:OpenId/audienceProtocolMapper:AudienceProtocolMapper", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AudienceProtocolMapper{s: s}, nil
}

// GetAudienceProtocolMapper gets an existing AudienceProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAudienceProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AudienceProtocolMapperState, opts ...pulumi.ResourceOpt) (*AudienceProtocolMapper, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addToAccessToken"] = state.AddToAccessToken
		inputs["addToIdToken"] = state.AddToIdToken
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["includedClientAudience"] = state.IncludedClientAudience
		inputs["includedCustomAudience"] = state.IncludedCustomAudience
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:OpenId/audienceProtocolMapper:AudienceProtocolMapper", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AudienceProtocolMapper{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AudienceProtocolMapper) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AudienceProtocolMapper) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Indicates if this claim should be added to the access token.
func (r *AudienceProtocolMapper) AddToAccessToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToAccessToken"])
}

// Indicates if this claim should be added to the id token.
func (r *AudienceProtocolMapper) AddToIdToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToIdToken"])
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (r *AudienceProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (r *AudienceProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

// A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
func (r *AudienceProtocolMapper) IncludedClientAudience() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["includedClientAudience"])
}

// A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
func (r *AudienceProtocolMapper) IncludedCustomAudience() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["includedCustomAudience"])
}

// A human-friendly name that will appear in the Keycloak console.
func (r *AudienceProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The realm id where the associated client or client scope exists.
func (r *AudienceProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering AudienceProtocolMapper resources.
type AudienceProtocolMapperState struct {
	// Indicates if this claim should be added to the access token.
	AddToAccessToken interface{}
	// Indicates if this claim should be added to the id token.
	AddToIdToken interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
	IncludedClientAudience interface{}
	// A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
	IncludedCustomAudience interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
}

// The set of arguments for constructing a AudienceProtocolMapper resource.
type AudienceProtocolMapperArgs struct {
	// Indicates if this claim should be added to the access token.
	AddToAccessToken interface{}
	// Indicates if this claim should be added to the id token.
	AddToIdToken interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
	IncludedClientAudience interface{}
	// A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
	IncludedCustomAudience interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
}
