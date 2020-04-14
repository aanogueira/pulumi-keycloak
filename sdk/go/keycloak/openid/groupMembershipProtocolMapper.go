// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # openid.GroupMembershipProtocolMapper
//
// Allows for creating and managing group membership protocol mappers within
// Keycloak.
//
// Group membership protocol mappers allow you to map a user's group memberships
// to a claim in a token. Protocol mappers can be defined for a single client,
// or they can be defined for a client scope which can be shared between multiple
// different clients.
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
// - `fullPath` - (Optional) Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
// - `addToIdToken` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
type GroupMembershipProtocolMapper struct {
	pulumi.CustomResourceState

	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	AddToIdToken     pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	AddToUserinfo    pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	ClaimName        pulumi.StringOutput  `pulumi:"claimName"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	FullPath      pulumi.BoolPtrOutput   `pulumi:"fullPath"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGroupMembershipProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewGroupMembershipProtocolMapper(ctx *pulumi.Context,
	name string, args *GroupMembershipProtocolMapperArgs, opts ...pulumi.ResourceOption) (*GroupMembershipProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GroupMembershipProtocolMapperArgs{}
	}
	var resource GroupMembershipProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembershipProtocolMapper gets an existing GroupMembershipProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembershipProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipProtocolMapperState, opts ...pulumi.ResourceOption) (*GroupMembershipProtocolMapper, error) {
	var resource GroupMembershipProtocolMapper
	err := ctx.ReadResource("keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembershipProtocolMapper resources.
type groupMembershipProtocolMapperState struct {
	AddToAccessToken *bool   `pulumi:"addToAccessToken"`
	AddToIdToken     *bool   `pulumi:"addToIdToken"`
	AddToUserinfo    *bool   `pulumi:"addToUserinfo"`
	ClaimName        *string `pulumi:"claimName"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	FullPath      *bool   `pulumi:"fullPath"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
}

type GroupMembershipProtocolMapperState struct {
	AddToAccessToken pulumi.BoolPtrInput
	AddToIdToken     pulumi.BoolPtrInput
	AddToUserinfo    pulumi.BoolPtrInput
	ClaimName        pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	FullPath      pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
}

func (GroupMembershipProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipProtocolMapperState)(nil)).Elem()
}

type groupMembershipProtocolMapperArgs struct {
	AddToAccessToken *bool  `pulumi:"addToAccessToken"`
	AddToIdToken     *bool  `pulumi:"addToIdToken"`
	AddToUserinfo    *bool  `pulumi:"addToUserinfo"`
	ClaimName        string `pulumi:"claimName"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	FullPath      *bool   `pulumi:"fullPath"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GroupMembershipProtocolMapper resource.
type GroupMembershipProtocolMapperArgs struct {
	AddToAccessToken pulumi.BoolPtrInput
	AddToIdToken     pulumi.BoolPtrInput
	AddToUserinfo    pulumi.BoolPtrInput
	ClaimName        pulumi.StringInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	FullPath      pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
}

func (GroupMembershipProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipProtocolMapperArgs)(nil)).Elem()
}
