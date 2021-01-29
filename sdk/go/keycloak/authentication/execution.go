// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing an authentication execution within Keycloak.
//
// An authentication execution is an action that the user or service may or may not take when authenticating through an authentication
// flow.
//
// > Due to limitations in the Keycloak API, the ordering of authentication executions within a flow must be specified using `dependsOn`. Authentication executions that are created first will appear first within the flow.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/authentication"
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
// 		executionOne, err := authentication.NewExecution(ctx, "executionOne", &authentication.ExecutionArgs{
// 			RealmId:         realm.ID(),
// 			ParentFlowAlias: flow.Alias,
// 			Authenticator:   pulumi.String("auth-cookie"),
// 			Requirement:     pulumi.String("ALTERNATIVE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authentication.NewExecution(ctx, "executionTwo", &authentication.ExecutionArgs{
// 			RealmId:         realm.ID(),
// 			ParentFlowAlias: flow.Alias,
// 			Authenticator:   pulumi.String("identity-provider-redirector"),
// 			Requirement:     pulumi.String("ALTERNATIVE"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			executionOne,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Authentication executions can be imported using the formats`{{realmId}}/{{parentFlowAlias}}/{{authenticationExecutionId}}`. Examplebash
//
// ```sh
//  $ pulumi import keycloak:authentication/execution:Execution keycloak_authentication_execution my-realm/my-flow/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
// ```
type Execution struct {
	pulumi.CustomResourceState

	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator pulumi.StringOutput `pulumi:"authenticator"`
	// The alias of the flow this execution is attached to.
	ParentFlowAlias pulumi.StringOutput `pulumi:"parentFlowAlias"`
	// The realm the authentication execution exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement pulumi.StringPtrOutput `pulumi:"requirement"`
}

// NewExecution registers a new resource with the given unique name, arguments, and options.
func NewExecution(ctx *pulumi.Context,
	name string, args *ExecutionArgs, opts ...pulumi.ResourceOption) (*Execution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authenticator == nil {
		return nil, errors.New("invalid value for required argument 'Authenticator'")
	}
	if args.ParentFlowAlias == nil {
		return nil, errors.New("invalid value for required argument 'ParentFlowAlias'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource Execution
	err := ctx.RegisterResource("keycloak:authentication/execution:Execution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExecution gets an existing Execution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExecutionState, opts ...pulumi.ResourceOption) (*Execution, error) {
	var resource Execution
	err := ctx.ReadResource("keycloak:authentication/execution:Execution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Execution resources.
type executionState struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator *string `pulumi:"authenticator"`
	// The alias of the flow this execution is attached to.
	ParentFlowAlias *string `pulumi:"parentFlowAlias"`
	// The realm the authentication execution exists in.
	RealmId *string `pulumi:"realmId"`
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement *string `pulumi:"requirement"`
}

type ExecutionState struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator pulumi.StringPtrInput
	// The alias of the flow this execution is attached to.
	ParentFlowAlias pulumi.StringPtrInput
	// The realm the authentication execution exists in.
	RealmId pulumi.StringPtrInput
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement pulumi.StringPtrInput
}

func (ExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionState)(nil)).Elem()
}

type executionArgs struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator string `pulumi:"authenticator"`
	// The alias of the flow this execution is attached to.
	ParentFlowAlias string `pulumi:"parentFlowAlias"`
	// The realm the authentication execution exists in.
	RealmId string `pulumi:"realmId"`
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement *string `pulumi:"requirement"`
}

// The set of arguments for constructing a Execution resource.
type ExecutionArgs struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator pulumi.StringInput
	// The alias of the flow this execution is attached to.
	ParentFlowAlias pulumi.StringInput
	// The realm the authentication execution exists in.
	RealmId pulumi.StringInput
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement pulumi.StringPtrInput
}

func (ExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionArgs)(nil)).Elem()
}

type ExecutionInput interface {
	pulumi.Input

	ToExecutionOutput() ExecutionOutput
	ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput
}

func (*Execution) ElementType() reflect.Type {
	return reflect.TypeOf((*Execution)(nil))
}

func (i *Execution) ToExecutionOutput() ExecutionOutput {
	return i.ToExecutionOutputWithContext(context.Background())
}

func (i *Execution) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionOutput)
}

type ExecutionOutput struct {
	*pulumi.OutputState
}

func (ExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Execution)(nil))
}

func (o ExecutionOutput) ToExecutionOutput() ExecutionOutput {
	return o
}

func (o ExecutionOutput) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExecutionOutput{})
}
