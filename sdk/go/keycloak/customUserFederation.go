// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/custom_user_federation.html.markdown.
type CustomUserFederation struct {
	s *pulumi.ResourceState
}

// NewCustomUserFederation registers a new resource with the given unique name, arguments, and options.
func NewCustomUserFederation(ctx *pulumi.Context,
	name string, args *CustomUserFederationArgs, opts ...pulumi.ResourceOpt) (*CustomUserFederation, error) {
	if args == nil || args.ProviderId == nil {
		return nil, errors.New("missing required argument 'ProviderId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cachePolicy"] = nil
		inputs["config"] = nil
		inputs["enabled"] = nil
		inputs["name"] = nil
		inputs["priority"] = nil
		inputs["providerId"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["cachePolicy"] = args.CachePolicy
		inputs["config"] = args.Config
		inputs["enabled"] = args.Enabled
		inputs["name"] = args.Name
		inputs["priority"] = args.Priority
		inputs["providerId"] = args.ProviderId
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:index/customUserFederation:CustomUserFederation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CustomUserFederation{s: s}, nil
}

// GetCustomUserFederation gets an existing CustomUserFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomUserFederation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CustomUserFederationState, opts ...pulumi.ResourceOpt) (*CustomUserFederation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cachePolicy"] = state.CachePolicy
		inputs["config"] = state.Config
		inputs["enabled"] = state.Enabled
		inputs["name"] = state.Name
		inputs["priority"] = state.Priority
		inputs["providerId"] = state.ProviderId
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:index/customUserFederation:CustomUserFederation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CustomUserFederation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CustomUserFederation) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CustomUserFederation) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *CustomUserFederation) CachePolicy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cachePolicy"])
}

func (r *CustomUserFederation) Config() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["config"])
}

// When false, this provider will not be used when performing queries for users.
func (r *CustomUserFederation) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// Display name of the provider when displayed in the console.
func (r *CustomUserFederation) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Priority of this provider when looking up users. Lower values are first.
func (r *CustomUserFederation) Priority() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["priority"])
}

// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
// interface
func (r *CustomUserFederation) ProviderId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["providerId"])
}

// The realm this provider will provide user federation for.
func (r *CustomUserFederation) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering CustomUserFederation resources.
type CustomUserFederationState struct {
	CachePolicy interface{}
	Config interface{}
	// When false, this provider will not be used when performing queries for users.
	Enabled interface{}
	// Display name of the provider when displayed in the console.
	Name interface{}
	// Priority of this provider when looking up users. Lower values are first.
	Priority interface{}
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId interface{}
	// The realm this provider will provide user federation for.
	RealmId interface{}
}

// The set of arguments for constructing a CustomUserFederation resource.
type CustomUserFederationArgs struct {
	CachePolicy interface{}
	Config interface{}
	// When false, this provider will not be used when performing queries for users.
	Enabled interface{}
	// Display name of the provider when displayed in the console.
	Name interface{}
	// Priority of this provider when looking up users. Lower values are first.
	Priority interface{}
	// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
	// interface
	ProviderId interface{}
	// The realm this provider will provide user federation for.
	RealmId interface{}
}