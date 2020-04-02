// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package authentication

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Execution struct {
	pulumi.CustomResourceState

	Authenticator   pulumi.StringOutput    `pulumi:"authenticator"`
	ParentFlowAlias pulumi.StringOutput    `pulumi:"parentFlowAlias"`
	RealmId         pulumi.StringOutput    `pulumi:"realmId"`
	Requirement     pulumi.StringPtrOutput `pulumi:"requirement"`
}

// NewExecution registers a new resource with the given unique name, arguments, and options.
func NewExecution(ctx *pulumi.Context,
	name string, args *ExecutionArgs, opts ...pulumi.ResourceOption) (*Execution, error) {
	if args == nil || args.Authenticator == nil {
		return nil, errors.New("missing required argument 'Authenticator'")
	}
	if args == nil || args.ParentFlowAlias == nil {
		return nil, errors.New("missing required argument 'ParentFlowAlias'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &ExecutionArgs{}
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
	Authenticator   *string `pulumi:"authenticator"`
	ParentFlowAlias *string `pulumi:"parentFlowAlias"`
	RealmId         *string `pulumi:"realmId"`
	Requirement     *string `pulumi:"requirement"`
}

type ExecutionState struct {
	Authenticator   pulumi.StringPtrInput
	ParentFlowAlias pulumi.StringPtrInput
	RealmId         pulumi.StringPtrInput
	Requirement     pulumi.StringPtrInput
}

func (ExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionState)(nil)).Elem()
}

type executionArgs struct {
	Authenticator   string  `pulumi:"authenticator"`
	ParentFlowAlias string  `pulumi:"parentFlowAlias"`
	RealmId         string  `pulumi:"realmId"`
	Requirement     *string `pulumi:"requirement"`
}

// The set of arguments for constructing a Execution resource.
type ExecutionArgs struct {
	Authenticator   pulumi.StringInput
	ParentFlowAlias pulumi.StringInput
	RealmId         pulumi.StringInput
	Requirement     pulumi.StringPtrInput
}

func (ExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionArgs)(nil)).Elem()
}