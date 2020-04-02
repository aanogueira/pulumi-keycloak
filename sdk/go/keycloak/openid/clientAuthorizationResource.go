// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClientAuthorizationResource struct {
	pulumi.CustomResourceState

	Attributes         pulumi.MapOutput         `pulumi:"attributes"`
	DisplayName        pulumi.StringPtrOutput   `pulumi:"displayName"`
	IconUri            pulumi.StringPtrOutput   `pulumi:"iconUri"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	OwnerManagedAccess pulumi.BoolPtrOutput     `pulumi:"ownerManagedAccess"`
	RealmId            pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId   pulumi.StringOutput      `pulumi:"resourceServerId"`
	Scopes             pulumi.StringArrayOutput `pulumi:"scopes"`
	Type               pulumi.StringPtrOutput   `pulumi:"type"`
	Uris               pulumi.StringArrayOutput `pulumi:"uris"`
}

// NewClientAuthorizationResource registers a new resource with the given unique name, arguments, and options.
func NewClientAuthorizationResource(ctx *pulumi.Context,
	name string, args *ClientAuthorizationResourceArgs, opts ...pulumi.ResourceOption) (*ClientAuthorizationResource, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil {
		args = &ClientAuthorizationResourceArgs{}
	}
	var resource ClientAuthorizationResource
	err := ctx.RegisterResource("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientAuthorizationResource gets an existing ClientAuthorizationResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientAuthorizationResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientAuthorizationResourceState, opts ...pulumi.ResourceOption) (*ClientAuthorizationResource, error) {
	var resource ClientAuthorizationResource
	err := ctx.ReadResource("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientAuthorizationResource resources.
type clientAuthorizationResourceState struct {
	Attributes         map[string]interface{} `pulumi:"attributes"`
	DisplayName        *string                `pulumi:"displayName"`
	IconUri            *string                `pulumi:"iconUri"`
	Name               *string                `pulumi:"name"`
	OwnerManagedAccess *bool                  `pulumi:"ownerManagedAccess"`
	RealmId            *string                `pulumi:"realmId"`
	ResourceServerId   *string                `pulumi:"resourceServerId"`
	Scopes             []string               `pulumi:"scopes"`
	Type               *string                `pulumi:"type"`
	Uris               []string               `pulumi:"uris"`
}

type ClientAuthorizationResourceState struct {
	Attributes         pulumi.MapInput
	DisplayName        pulumi.StringPtrInput
	IconUri            pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	OwnerManagedAccess pulumi.BoolPtrInput
	RealmId            pulumi.StringPtrInput
	ResourceServerId   pulumi.StringPtrInput
	Scopes             pulumi.StringArrayInput
	Type               pulumi.StringPtrInput
	Uris               pulumi.StringArrayInput
}

func (ClientAuthorizationResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationResourceState)(nil)).Elem()
}

type clientAuthorizationResourceArgs struct {
	Attributes         map[string]interface{} `pulumi:"attributes"`
	DisplayName        *string                `pulumi:"displayName"`
	IconUri            *string                `pulumi:"iconUri"`
	Name               *string                `pulumi:"name"`
	OwnerManagedAccess *bool                  `pulumi:"ownerManagedAccess"`
	RealmId            string                 `pulumi:"realmId"`
	ResourceServerId   string                 `pulumi:"resourceServerId"`
	Scopes             []string               `pulumi:"scopes"`
	Type               *string                `pulumi:"type"`
	Uris               []string               `pulumi:"uris"`
}

// The set of arguments for constructing a ClientAuthorizationResource resource.
type ClientAuthorizationResourceArgs struct {
	Attributes         pulumi.MapInput
	DisplayName        pulumi.StringPtrInput
	IconUri            pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	OwnerManagedAccess pulumi.BoolPtrInput
	RealmId            pulumi.StringInput
	ResourceServerId   pulumi.StringInput
	Scopes             pulumi.StringArrayInput
	Type               pulumi.StringPtrInput
	Uris               pulumi.StringArrayInput
}

func (ClientAuthorizationResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationResourceArgs)(nil)).Elem()
}