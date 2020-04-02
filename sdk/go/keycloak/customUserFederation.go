// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CustomUserFederation struct {
	pulumi.CustomResourceState

	CachePolicy pulumi.StringPtrOutput `pulumi:"cachePolicy"`
	Config      pulumi.MapOutput       `pulumi:"config"`
	// When false, this provider will not be used when performing queries for users.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Display name of the provider when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority of this provider when looking up users. Lower values are first.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId pulumi.StringOutput `pulumi:"providerId"`
	// The realm this provider will provide user federation for.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewCustomUserFederation registers a new resource with the given unique name, arguments, and options.
func NewCustomUserFederation(ctx *pulumi.Context,
	name string, args *CustomUserFederationArgs, opts ...pulumi.ResourceOption) (*CustomUserFederation, error) {
	if args == nil || args.ProviderId == nil {
		return nil, errors.New("missing required argument 'ProviderId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &CustomUserFederationArgs{}
	}
	var resource CustomUserFederation
	err := ctx.RegisterResource("keycloak:index/customUserFederation:CustomUserFederation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomUserFederation gets an existing CustomUserFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomUserFederation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomUserFederationState, opts ...pulumi.ResourceOption) (*CustomUserFederation, error) {
	var resource CustomUserFederation
	err := ctx.ReadResource("keycloak:index/customUserFederation:CustomUserFederation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomUserFederation resources.
type customUserFederationState struct {
	CachePolicy *string                `pulumi:"cachePolicy"`
	Config      map[string]interface{} `pulumi:"config"`
	// When false, this provider will not be used when performing queries for users.
	Enabled *bool `pulumi:"enabled"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// Priority of this provider when looking up users. Lower values are first.
	Priority *int `pulumi:"priority"`
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId *string `pulumi:"providerId"`
	// The realm this provider will provide user federation for.
	RealmId *string `pulumi:"realmId"`
}

type CustomUserFederationState struct {
	CachePolicy pulumi.StringPtrInput
	Config      pulumi.MapInput
	// When false, this provider will not be used when performing queries for users.
	Enabled pulumi.BoolPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// Priority of this provider when looking up users. Lower values are first.
	Priority pulumi.IntPtrInput
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId pulumi.StringPtrInput
	// The realm this provider will provide user federation for.
	RealmId pulumi.StringPtrInput
}

func (CustomUserFederationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customUserFederationState)(nil)).Elem()
}

type customUserFederationArgs struct {
	CachePolicy *string                `pulumi:"cachePolicy"`
	Config      map[string]interface{} `pulumi:"config"`
	// When false, this provider will not be used when performing queries for users.
	Enabled *bool `pulumi:"enabled"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// Priority of this provider when looking up users. Lower values are first.
	Priority *int `pulumi:"priority"`
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId string `pulumi:"providerId"`
	// The realm this provider will provide user federation for.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a CustomUserFederation resource.
type CustomUserFederationArgs struct {
	CachePolicy pulumi.StringPtrInput
	Config      pulumi.MapInput
	// When false, this provider will not be used when performing queries for users.
	Enabled pulumi.BoolPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// Priority of this provider when looking up users. Lower values are first.
	Priority pulumi.IntPtrInput
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId pulumi.StringInput
	// The realm this provider will provide user federation for.
	RealmId pulumi.StringInput
}

func (CustomUserFederationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customUserFederationArgs)(nil)).Elem()
}
