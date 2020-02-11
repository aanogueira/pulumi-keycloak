// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # .User
// 
// Allows for creating and managing Users within Keycloak.
// 
// This resource was created primarily to enable the acceptance tests for the `.Group` resource.
// Creating users within Keycloak is not recommended. Instead, users should be federated from external sources
// by configuring user federation providers or identity providers.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this user belongs to.
// - `username` - (Required) The unique username of this user.
// - `initialPassword` (Optional) When given, the user's initial password will be set.
//    This attribute is only respected during initial user creation.
//     - `value` (Required) The initial password.
//     - `temporary` (Optional) If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
// - `enabled` - (Optional) When false, this user cannot log in. Defaults to `true`.
// - `email` - (Optional) The user's email.
// - `firstName` - (Optional) The user's first name.
// - `lastName` - (Optional) The user's last name.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/user.html.markdown.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attributes"] = nil
		inputs["email"] = nil
		inputs["enabled"] = nil
		inputs["federatedIdentities"] = nil
		inputs["firstName"] = nil
		inputs["initialPassword"] = nil
		inputs["lastName"] = nil
		inputs["realmId"] = nil
		inputs["username"] = nil
	} else {
		inputs["attributes"] = args.Attributes
		inputs["email"] = args.Email
		inputs["enabled"] = args.Enabled
		inputs["federatedIdentities"] = args.FederatedIdentities
		inputs["firstName"] = args.FirstName
		inputs["initialPassword"] = args.InitialPassword
		inputs["lastName"] = args.LastName
		inputs["realmId"] = args.RealmId
		inputs["username"] = args.Username
	}
	s, err := ctx.RegisterResource("keycloak:index/user:User", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserState, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attributes"] = state.Attributes
		inputs["email"] = state.Email
		inputs["enabled"] = state.Enabled
		inputs["federatedIdentities"] = state.FederatedIdentities
		inputs["firstName"] = state.FirstName
		inputs["initialPassword"] = state.InitialPassword
		inputs["lastName"] = state.LastName
		inputs["realmId"] = state.RealmId
		inputs["username"] = state.Username
	}
	s, err := ctx.ReadResource("keycloak:index/user:User", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *User) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *User) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *User) Attributes() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["attributes"])
}

func (r *User) Email() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["email"])
}

func (r *User) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

func (r *User) FederatedIdentities() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["federatedIdentities"])
}

func (r *User) FirstName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["firstName"])
}

func (r *User) InitialPassword() pulumi.Output {
	return r.s.State["initialPassword"]
}

func (r *User) LastName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastName"])
}

func (r *User) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

func (r *User) Username() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["username"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	Attributes interface{}
	Email interface{}
	Enabled interface{}
	FederatedIdentities interface{}
	FirstName interface{}
	InitialPassword interface{}
	LastName interface{}
	RealmId interface{}
	Username interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	Attributes interface{}
	Email interface{}
	Enabled interface{}
	FederatedIdentities interface{}
	FirstName interface{}
	InitialPassword interface{}
	LastName interface{}
	RealmId interface{}
	Username interface{}
}