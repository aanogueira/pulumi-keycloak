// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # openid.HardcodedRoleProtocolMapper
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
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_openid_hardcoded_role_protocol_mapper.html.markdown.
type HardcodedRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	RoleId  pulumi.StringOutput `pulumi:"roleId"`
}

// NewHardcodedRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *HardcodedRoleProtocolMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedRoleProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	if args == nil {
		args = &HardcodedRoleProtocolMapperArgs{}
	}
	var resource HardcodedRoleProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedRoleProtocolMapper gets an existing HardcodedRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedRoleProtocolMapperState, opts ...pulumi.ResourceOption) (*HardcodedRoleProtocolMapper, error) {
	var resource HardcodedRoleProtocolMapper
	err := ctx.ReadResource("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedRoleProtocolMapper resources.
type hardcodedRoleProtocolMapperState struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
	RoleId  *string `pulumi:"roleId"`
}

type HardcodedRoleProtocolMapperState struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
	RoleId  pulumi.StringPtrInput
}

func (HardcodedRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleProtocolMapperState)(nil)).Elem()
}

type hardcodedRoleProtocolMapperArgs struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
	RoleId  string `pulumi:"roleId"`
}

// The set of arguments for constructing a HardcodedRoleProtocolMapper resource.
type HardcodedRoleProtocolMapperArgs struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
	RoleId  pulumi.StringInput
}

func (HardcodedRoleProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleProtocolMapperArgs)(nil)).Elem()
}