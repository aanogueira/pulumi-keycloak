// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # OpenId.HardcodedClaimProtocolMapper
// 
// Allows for creating and managing hardcoded claim protocol mappers within
// Keycloak.
// 
// Hardcoded claim protocol mappers allow you to define a claim with a hardcoded
// value. Protocol mappers can be defined for a single client, or they can
// be defined for a client scope which can be shared between multiple different
// clients.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `claimName` - (Required) The name of the claim to insert into a token.
// - `claimValue` - (Required) The hardcoded value of the claim.
// - `claimValueType` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
// - `addToIdToken` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_hardcoded_claim_protocol_mapper.html.markdown.
type HardcodedClaimProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewHardcodedClaimProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedClaimProtocolMapper(ctx *pulumi.Context,
	name string, args *HardcodedClaimProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*HardcodedClaimProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.ClaimValue == nil {
		return nil, errors.New("missing required argument 'ClaimValue'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addToAccessToken"] = nil
		inputs["addToIdToken"] = nil
		inputs["addToUserinfo"] = nil
		inputs["claimName"] = nil
		inputs["claimValue"] = nil
		inputs["claimValueType"] = nil
		inputs["clientId"] = nil
		inputs["clientScopeId"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["addToAccessToken"] = args.AddToAccessToken
		inputs["addToIdToken"] = args.AddToIdToken
		inputs["addToUserinfo"] = args.AddToUserinfo
		inputs["claimName"] = args.ClaimName
		inputs["claimValue"] = args.ClaimValue
		inputs["claimValueType"] = args.ClaimValueType
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:OpenId/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HardcodedClaimProtocolMapper{s: s}, nil
}

// GetHardcodedClaimProtocolMapper gets an existing HardcodedClaimProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedClaimProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HardcodedClaimProtocolMapperState, opts ...pulumi.ResourceOpt) (*HardcodedClaimProtocolMapper, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addToAccessToken"] = state.AddToAccessToken
		inputs["addToIdToken"] = state.AddToIdToken
		inputs["addToUserinfo"] = state.AddToUserinfo
		inputs["claimName"] = state.ClaimName
		inputs["claimValue"] = state.ClaimValue
		inputs["claimValueType"] = state.ClaimValueType
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:OpenId/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HardcodedClaimProtocolMapper{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HardcodedClaimProtocolMapper) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HardcodedClaimProtocolMapper) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Indicates if the attribute should be a claim in the access token.
func (r *HardcodedClaimProtocolMapper) AddToAccessToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToAccessToken"])
}

// Indicates if the attribute should be a claim in the id token.
func (r *HardcodedClaimProtocolMapper) AddToIdToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToIdToken"])
}

// Indicates if the attribute should appear in the userinfo response body.
func (r *HardcodedClaimProtocolMapper) AddToUserinfo() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToUserinfo"])
}

func (r *HardcodedClaimProtocolMapper) ClaimName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimName"])
}

func (r *HardcodedClaimProtocolMapper) ClaimValue() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimValue"])
}

// Claim type used when serializing tokens.
func (r *HardcodedClaimProtocolMapper) ClaimValueType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimValueType"])
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (r *HardcodedClaimProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (r *HardcodedClaimProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

// A human-friendly name that will appear in the Keycloak console.
func (r *HardcodedClaimProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The realm id where the associated client or client scope exists.
func (r *HardcodedClaimProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering HardcodedClaimProtocolMapper resources.
type HardcodedClaimProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	ClaimValue interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
}

// The set of arguments for constructing a HardcodedClaimProtocolMapper resource.
type HardcodedClaimProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	ClaimValue interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
}
