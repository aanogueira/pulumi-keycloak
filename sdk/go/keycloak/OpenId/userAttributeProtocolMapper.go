// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # OpenId.UserAttributeProtocolMapper
// 
// Allows for creating and managing user attribute protocol mappers within
// Keycloak.
// 
// User attribute protocol mappers allow you to map custom attributes defined
// for a user within Keycloak to a claim in a token. Protocol mappers can be
// defined for a single client, or they can be defined for a client scope which
// can be shared between multiple different clients.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `userAttribute` - (Required) The custom user attribute to map a claim for.
// - `claimName` - (Required) The name of the claim to insert into a token.
// - `claimValueType` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
// - `multivalued` - (Optional) Indicates whether this attribute is a single value or an array of values. Defaults to `false`.
// - `addToIdToken` - (Optional) Indicates if the attribute should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the attribute should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the attribute should be added as a claim to the UserInfo response body. Defaults to `true`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_user_attribute_protocol_mapper.html.markdown.
type UserAttributeProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewUserAttributeProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, args *UserAttributeProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*UserAttributeProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.UserAttribute == nil {
		return nil, errors.New("missing required argument 'UserAttribute'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addToAccessToken"] = nil
		inputs["addToIdToken"] = nil
		inputs["addToUserinfo"] = nil
		inputs["claimName"] = nil
		inputs["claimValueType"] = nil
		inputs["clientId"] = nil
		inputs["clientScopeId"] = nil
		inputs["multivalued"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
		inputs["userAttribute"] = nil
	} else {
		inputs["addToAccessToken"] = args.AddToAccessToken
		inputs["addToIdToken"] = args.AddToIdToken
		inputs["addToUserinfo"] = args.AddToUserinfo
		inputs["claimName"] = args.ClaimName
		inputs["claimValueType"] = args.ClaimValueType
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["multivalued"] = args.Multivalued
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
		inputs["userAttribute"] = args.UserAttribute
	}
	s, err := ctx.RegisterResource("keycloak:OpenId/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserAttributeProtocolMapper{s: s}, nil
}

// GetUserAttributeProtocolMapper gets an existing UserAttributeProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserAttributeProtocolMapperState, opts ...pulumi.ResourceOpt) (*UserAttributeProtocolMapper, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addToAccessToken"] = state.AddToAccessToken
		inputs["addToIdToken"] = state.AddToIdToken
		inputs["addToUserinfo"] = state.AddToUserinfo
		inputs["claimName"] = state.ClaimName
		inputs["claimValueType"] = state.ClaimValueType
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["multivalued"] = state.Multivalued
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
		inputs["userAttribute"] = state.UserAttribute
	}
	s, err := ctx.ReadResource("keycloak:OpenId/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserAttributeProtocolMapper{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UserAttributeProtocolMapper) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UserAttributeProtocolMapper) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Indicates if the attribute should be a claim in the access token.
func (r *UserAttributeProtocolMapper) AddToAccessToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToAccessToken"])
}

// Indicates if the attribute should be a claim in the id token.
func (r *UserAttributeProtocolMapper) AddToIdToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToIdToken"])
}

// Indicates if the attribute should appear in the userinfo response body.
func (r *UserAttributeProtocolMapper) AddToUserinfo() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToUserinfo"])
}

func (r *UserAttributeProtocolMapper) ClaimName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimName"])
}

// Claim type used when serializing tokens.
func (r *UserAttributeProtocolMapper) ClaimValueType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimValueType"])
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (r *UserAttributeProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (r *UserAttributeProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

// Indicates whether this attribute is a single value or an array of values.
func (r *UserAttributeProtocolMapper) Multivalued() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["multivalued"])
}

// A human-friendly name that will appear in the Keycloak console.
func (r *UserAttributeProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The realm id where the associated client or client scope exists.
func (r *UserAttributeProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

func (r *UserAttributeProtocolMapper) UserAttribute() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userAttribute"])
}

// Input properties used for looking up and filtering UserAttributeProtocolMapper resources.
type UserAttributeProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	UserAttribute interface{}
}

// The set of arguments for constructing a UserAttributeProtocolMapper resource.
type UserAttributeProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	UserAttribute interface{}
}
