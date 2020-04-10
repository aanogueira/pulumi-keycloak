// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # .Role
//
// Allows for creating and managing roles within Keycloak.
//
// Roles allow you define privileges within Keycloak and map them to users
// and groups.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this role exists within.
// - `clientId` - (Optional) When specified, this role will be created as
//   a client role attached to the client with the provided ID
// - `name` - (Required) The name of the role
// - `description` - (Optional) The description of the role
// - `compositeRoles` - (Optional) When specified, this role will be a
//   composite role, composed of all roles that have an ID present within
//   this list.
type Role struct {
	pulumi.CustomResourceState

	ClientId       pulumi.StringPtrOutput   `pulumi:"clientId"`
	CompositeRoles pulumi.StringArrayOutput `pulumi:"compositeRoles"`
	Description    pulumi.StringPtrOutput   `pulumi:"description"`
	Name           pulumi.StringOutput      `pulumi:"name"`
	RealmId        pulumi.StringOutput      `pulumi:"realmId"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &RoleArgs{}
	}
	var resource Role
	err := ctx.RegisterResource("keycloak:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("keycloak:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	ClientId       *string  `pulumi:"clientId"`
	CompositeRoles []string `pulumi:"compositeRoles"`
	Description    *string  `pulumi:"description"`
	Name           *string  `pulumi:"name"`
	RealmId        *string  `pulumi:"realmId"`
}

type RoleState struct {
	ClientId       pulumi.StringPtrInput
	CompositeRoles pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	RealmId        pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	ClientId       *string  `pulumi:"clientId"`
	CompositeRoles []string `pulumi:"compositeRoles"`
	Description    *string  `pulumi:"description"`
	Name           *string  `pulumi:"name"`
	RealmId        string   `pulumi:"realmId"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	ClientId       pulumi.StringPtrInput
	CompositeRoles pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	RealmId        pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}
