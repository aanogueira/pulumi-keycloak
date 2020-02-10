// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # OpenId.HardcodedRoleProtocolMapper
// 
// Allows for creating and managing hardcoded role protocol mappers within
// Keycloak.
// 
// Hardcoded role protocol mappers allow you to specify a single role to
// always map to an access token for a client. Protocol mappers can be
// defined for a single client, or they can be defined for a client scope
// which can be shared between multiple different clients.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the
//   GUI.
// - `roleId` - (Required) The ID of the role to map to an access token.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_hardcoded_role_protocol_mapper.html.markdown.
type HardcodedRoleProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewHardcodedRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *HardcodedRoleProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*HardcodedRoleProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clientId"] = nil
		inputs["clientScopeId"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
		inputs["roleId"] = nil
	} else {
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
		inputs["roleId"] = args.RoleId
	}
	s, err := ctx.RegisterResource("keycloak:OpenId/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HardcodedRoleProtocolMapper{s: s}, nil
}

// GetHardcodedRoleProtocolMapper gets an existing HardcodedRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HardcodedRoleProtocolMapperState, opts ...pulumi.ResourceOpt) (*HardcodedRoleProtocolMapper, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
		inputs["roleId"] = state.RoleId
	}
	s, err := ctx.ReadResource("keycloak:OpenId/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HardcodedRoleProtocolMapper{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HardcodedRoleProtocolMapper) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HardcodedRoleProtocolMapper) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (r *HardcodedRoleProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (r *HardcodedRoleProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

// A human-friendly name that will appear in the Keycloak console.
func (r *HardcodedRoleProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The realm id where the associated client or client scope exists.
func (r *HardcodedRoleProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

func (r *HardcodedRoleProtocolMapper) RoleId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleId"])
}

// Input properties used for looking up and filtering HardcodedRoleProtocolMapper resources.
type HardcodedRoleProtocolMapperState struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	RoleId interface{}
}

// The set of arguments for constructing a HardcodedRoleProtocolMapper resource.
type HardcodedRoleProtocolMapperArgs struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	RoleId interface{}
}
