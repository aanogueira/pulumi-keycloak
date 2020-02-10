// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package Saml

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # Saml.UserAttributeProtocolMapper
// 
// Allows for creating and managing user attribute protocol mappers for
// SAML clients within Keycloak.
// 
// SAML user attribute protocol mappers allow you to map custom attributes defined
// for a user within Keycloak to an attribute in a SAML assertion. Protocol mappers
// can be defined for a single client, or they can be defined for a client scope which
// can be shared between multiple different clients.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The SAML client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The SAML client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `userAttribute` - (Required) The custom user attribute to map.
// - `friendlyName` - (Optional) An optional human-friendly name for this attribute.
// - `samlAttributeName` - (Required) The name of the SAML attribute.
// - `samlAttributeNameFormat` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/saml_user_attribute_protocol_mapper.html.markdown.
type UserAttributeProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewUserAttributeProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, args *UserAttributeProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*UserAttributeProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.SamlAttributeName == nil {
		return nil, errors.New("missing required argument 'SamlAttributeName'")
	}
	if args == nil || args.SamlAttributeNameFormat == nil {
		return nil, errors.New("missing required argument 'SamlAttributeNameFormat'")
	}
	if args == nil || args.UserAttribute == nil {
		return nil, errors.New("missing required argument 'UserAttribute'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clientId"] = nil
		inputs["clientScopeId"] = nil
		inputs["friendlyName"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
		inputs["samlAttributeName"] = nil
		inputs["samlAttributeNameFormat"] = nil
		inputs["userAttribute"] = nil
	} else {
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["friendlyName"] = args.FriendlyName
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
		inputs["samlAttributeName"] = args.SamlAttributeName
		inputs["samlAttributeNameFormat"] = args.SamlAttributeNameFormat
		inputs["userAttribute"] = args.UserAttribute
	}
	s, err := ctx.RegisterResource("keycloak:Saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, true, inputs, opts...)
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
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["friendlyName"] = state.FriendlyName
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
		inputs["samlAttributeName"] = state.SamlAttributeName
		inputs["samlAttributeNameFormat"] = state.SamlAttributeNameFormat
		inputs["userAttribute"] = state.UserAttribute
	}
	s, err := ctx.ReadResource("keycloak:Saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, id, inputs, opts...)
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

func (r *UserAttributeProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

func (r *UserAttributeProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

func (r *UserAttributeProtocolMapper) FriendlyName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["friendlyName"])
}

func (r *UserAttributeProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *UserAttributeProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

func (r *UserAttributeProtocolMapper) SamlAttributeName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["samlAttributeName"])
}

func (r *UserAttributeProtocolMapper) SamlAttributeNameFormat() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["samlAttributeNameFormat"])
}

func (r *UserAttributeProtocolMapper) UserAttribute() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userAttribute"])
}

// Input properties used for looking up and filtering UserAttributeProtocolMapper resources.
type UserAttributeProtocolMapperState struct {
	ClientId interface{}
	ClientScopeId interface{}
	FriendlyName interface{}
	Name interface{}
	RealmId interface{}
	SamlAttributeName interface{}
	SamlAttributeNameFormat interface{}
	UserAttribute interface{}
}

// The set of arguments for constructing a UserAttributeProtocolMapper resource.
type UserAttributeProtocolMapperArgs struct {
	ClientId interface{}
	ClientScopeId interface{}
	FriendlyName interface{}
	Name interface{}
	RealmId interface{}
	SamlAttributeName interface{}
	SamlAttributeNameFormat interface{}
	UserAttribute interface{}
}
