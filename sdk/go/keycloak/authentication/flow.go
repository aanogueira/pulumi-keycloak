// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing an authentication flow within Keycloak.
//
// [Authentication flows](https://www.keycloak.org/docs/11.0/server_admin/index.html#_authentication-flows) describe a sequence
// of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
// is a container for these actions, which are otherwise known as executions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/authentication"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		flow, err := authentication.NewFlow(ctx, "flow", &authentication.FlowArgs{
// 			RealmId: realm.ID(),
// 			Alias:   pulumi.String("my-flow-alias"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authentication.NewExecution(ctx, "execution", &authentication.ExecutionArgs{
// 			RealmId:         realm.ID(),
// 			ParentFlowAlias: flow.Alias,
// 			Authenticator:   pulumi.String("identity-provider-redirector"),
// 			Requirement:     pulumi.String("REQUIRED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Flow struct {
	pulumi.CustomResourceState

	// The alias for this authentication flow.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// A description for the authentication flow.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId pulumi.StringPtrOutput `pulumi:"providerId"`
	// The realm that the authentication flow exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil || args.Alias == nil {
		return nil, errors.New("missing required argument 'Alias'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &FlowArgs{}
	}
	var resource Flow
	err := ctx.RegisterResource("keycloak:authentication/flow:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("keycloak:authentication/flow:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
	// The alias for this authentication flow.
	Alias *string `pulumi:"alias"`
	// A description for the authentication flow.
	Description *string `pulumi:"description"`
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId *string `pulumi:"providerId"`
	// The realm that the authentication flow exists in.
	RealmId *string `pulumi:"realmId"`
}

type FlowState struct {
	// The alias for this authentication flow.
	Alias pulumi.StringPtrInput
	// A description for the authentication flow.
	Description pulumi.StringPtrInput
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId pulumi.StringPtrInput
	// The realm that the authentication flow exists in.
	RealmId pulumi.StringPtrInput
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	// The alias for this authentication flow.
	Alias string `pulumi:"alias"`
	// A description for the authentication flow.
	Description *string `pulumi:"description"`
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId *string `pulumi:"providerId"`
	// The realm that the authentication flow exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	// The alias for this authentication flow.
	Alias pulumi.StringInput
	// A description for the authentication flow.
	Description pulumi.StringPtrInput
	// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
	ProviderId pulumi.StringPtrInput
	// The realm that the authentication flow exists in.
	RealmId pulumi.StringInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}
