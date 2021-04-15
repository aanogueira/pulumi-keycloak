// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientAuthorizationScope struct {
	pulumi.CustomResourceState

	DisplayName      pulumi.StringPtrOutput `pulumi:"displayName"`
	IconUri          pulumi.StringPtrOutput `pulumi:"iconUri"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	RealmId          pulumi.StringOutput    `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput    `pulumi:"resourceServerId"`
}

// NewClientAuthorizationScope registers a new resource with the given unique name, arguments, and options.
func NewClientAuthorizationScope(ctx *pulumi.Context,
	name string, args *ClientAuthorizationScopeArgs, opts ...pulumi.ResourceOption) (*ClientAuthorizationScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	var resource ClientAuthorizationScope
	err := ctx.RegisterResource("keycloak:openid/clientAuthorizationScope:ClientAuthorizationScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientAuthorizationScope gets an existing ClientAuthorizationScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientAuthorizationScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientAuthorizationScopeState, opts ...pulumi.ResourceOption) (*ClientAuthorizationScope, error) {
	var resource ClientAuthorizationScope
	err := ctx.ReadResource("keycloak:openid/clientAuthorizationScope:ClientAuthorizationScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientAuthorizationScope resources.
type clientAuthorizationScopeState struct {
	DisplayName      *string `pulumi:"displayName"`
	IconUri          *string `pulumi:"iconUri"`
	Name             *string `pulumi:"name"`
	RealmId          *string `pulumi:"realmId"`
	ResourceServerId *string `pulumi:"resourceServerId"`
}

type ClientAuthorizationScopeState struct {
	DisplayName      pulumi.StringPtrInput
	IconUri          pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
}

func (ClientAuthorizationScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationScopeState)(nil)).Elem()
}

type clientAuthorizationScopeArgs struct {
	DisplayName      *string `pulumi:"displayName"`
	IconUri          *string `pulumi:"iconUri"`
	Name             *string `pulumi:"name"`
	RealmId          string  `pulumi:"realmId"`
	ResourceServerId string  `pulumi:"resourceServerId"`
}

// The set of arguments for constructing a ClientAuthorizationScope resource.
type ClientAuthorizationScopeArgs struct {
	DisplayName      pulumi.StringPtrInput
	IconUri          pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
}

func (ClientAuthorizationScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationScopeArgs)(nil)).Elem()
}

type ClientAuthorizationScopeInput interface {
	pulumi.Input

	ToClientAuthorizationScopeOutput() ClientAuthorizationScopeOutput
	ToClientAuthorizationScopeOutputWithContext(ctx context.Context) ClientAuthorizationScopeOutput
}

func (*ClientAuthorizationScope) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthorizationScope)(nil))
}

func (i *ClientAuthorizationScope) ToClientAuthorizationScopeOutput() ClientAuthorizationScopeOutput {
	return i.ToClientAuthorizationScopeOutputWithContext(context.Background())
}

func (i *ClientAuthorizationScope) ToClientAuthorizationScopeOutputWithContext(ctx context.Context) ClientAuthorizationScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationScopeOutput)
}

func (i *ClientAuthorizationScope) ToClientAuthorizationScopePtrOutput() ClientAuthorizationScopePtrOutput {
	return i.ToClientAuthorizationScopePtrOutputWithContext(context.Background())
}

func (i *ClientAuthorizationScope) ToClientAuthorizationScopePtrOutputWithContext(ctx context.Context) ClientAuthorizationScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationScopePtrOutput)
}

type ClientAuthorizationScopePtrInput interface {
	pulumi.Input

	ToClientAuthorizationScopePtrOutput() ClientAuthorizationScopePtrOutput
	ToClientAuthorizationScopePtrOutputWithContext(ctx context.Context) ClientAuthorizationScopePtrOutput
}

type clientAuthorizationScopePtrType ClientAuthorizationScopeArgs

func (*clientAuthorizationScopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationScope)(nil))
}

func (i *clientAuthorizationScopePtrType) ToClientAuthorizationScopePtrOutput() ClientAuthorizationScopePtrOutput {
	return i.ToClientAuthorizationScopePtrOutputWithContext(context.Background())
}

func (i *clientAuthorizationScopePtrType) ToClientAuthorizationScopePtrOutputWithContext(ctx context.Context) ClientAuthorizationScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationScopePtrOutput)
}

// ClientAuthorizationScopeArrayInput is an input type that accepts ClientAuthorizationScopeArray and ClientAuthorizationScopeArrayOutput values.
// You can construct a concrete instance of `ClientAuthorizationScopeArrayInput` via:
//
//          ClientAuthorizationScopeArray{ ClientAuthorizationScopeArgs{...} }
type ClientAuthorizationScopeArrayInput interface {
	pulumi.Input

	ToClientAuthorizationScopeArrayOutput() ClientAuthorizationScopeArrayOutput
	ToClientAuthorizationScopeArrayOutputWithContext(context.Context) ClientAuthorizationScopeArrayOutput
}

type ClientAuthorizationScopeArray []ClientAuthorizationScopeInput

func (ClientAuthorizationScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClientAuthorizationScope)(nil))
}

func (i ClientAuthorizationScopeArray) ToClientAuthorizationScopeArrayOutput() ClientAuthorizationScopeArrayOutput {
	return i.ToClientAuthorizationScopeArrayOutputWithContext(context.Background())
}

func (i ClientAuthorizationScopeArray) ToClientAuthorizationScopeArrayOutputWithContext(ctx context.Context) ClientAuthorizationScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationScopeArrayOutput)
}

// ClientAuthorizationScopeMapInput is an input type that accepts ClientAuthorizationScopeMap and ClientAuthorizationScopeMapOutput values.
// You can construct a concrete instance of `ClientAuthorizationScopeMapInput` via:
//
//          ClientAuthorizationScopeMap{ "key": ClientAuthorizationScopeArgs{...} }
type ClientAuthorizationScopeMapInput interface {
	pulumi.Input

	ToClientAuthorizationScopeMapOutput() ClientAuthorizationScopeMapOutput
	ToClientAuthorizationScopeMapOutputWithContext(context.Context) ClientAuthorizationScopeMapOutput
}

type ClientAuthorizationScopeMap map[string]ClientAuthorizationScopeInput

func (ClientAuthorizationScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClientAuthorizationScope)(nil))
}

func (i ClientAuthorizationScopeMap) ToClientAuthorizationScopeMapOutput() ClientAuthorizationScopeMapOutput {
	return i.ToClientAuthorizationScopeMapOutputWithContext(context.Background())
}

func (i ClientAuthorizationScopeMap) ToClientAuthorizationScopeMapOutputWithContext(ctx context.Context) ClientAuthorizationScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationScopeMapOutput)
}

type ClientAuthorizationScopeOutput struct {
	*pulumi.OutputState
}

func (ClientAuthorizationScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthorizationScope)(nil))
}

func (o ClientAuthorizationScopeOutput) ToClientAuthorizationScopeOutput() ClientAuthorizationScopeOutput {
	return o
}

func (o ClientAuthorizationScopeOutput) ToClientAuthorizationScopeOutputWithContext(ctx context.Context) ClientAuthorizationScopeOutput {
	return o
}

func (o ClientAuthorizationScopeOutput) ToClientAuthorizationScopePtrOutput() ClientAuthorizationScopePtrOutput {
	return o.ToClientAuthorizationScopePtrOutputWithContext(context.Background())
}

func (o ClientAuthorizationScopeOutput) ToClientAuthorizationScopePtrOutputWithContext(ctx context.Context) ClientAuthorizationScopePtrOutput {
	return o.ApplyT(func(v ClientAuthorizationScope) *ClientAuthorizationScope {
		return &v
	}).(ClientAuthorizationScopePtrOutput)
}

type ClientAuthorizationScopePtrOutput struct {
	*pulumi.OutputState
}

func (ClientAuthorizationScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationScope)(nil))
}

func (o ClientAuthorizationScopePtrOutput) ToClientAuthorizationScopePtrOutput() ClientAuthorizationScopePtrOutput {
	return o
}

func (o ClientAuthorizationScopePtrOutput) ToClientAuthorizationScopePtrOutputWithContext(ctx context.Context) ClientAuthorizationScopePtrOutput {
	return o
}

type ClientAuthorizationScopeArrayOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAuthorizationScope)(nil))
}

func (o ClientAuthorizationScopeArrayOutput) ToClientAuthorizationScopeArrayOutput() ClientAuthorizationScopeArrayOutput {
	return o
}

func (o ClientAuthorizationScopeArrayOutput) ToClientAuthorizationScopeArrayOutputWithContext(ctx context.Context) ClientAuthorizationScopeArrayOutput {
	return o
}

func (o ClientAuthorizationScopeArrayOutput) Index(i pulumi.IntInput) ClientAuthorizationScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientAuthorizationScope {
		return vs[0].([]ClientAuthorizationScope)[vs[1].(int)]
	}).(ClientAuthorizationScopeOutput)
}

type ClientAuthorizationScopeMapOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClientAuthorizationScope)(nil))
}

func (o ClientAuthorizationScopeMapOutput) ToClientAuthorizationScopeMapOutput() ClientAuthorizationScopeMapOutput {
	return o
}

func (o ClientAuthorizationScopeMapOutput) ToClientAuthorizationScopeMapOutputWithContext(ctx context.Context) ClientAuthorizationScopeMapOutput {
	return o
}

func (o ClientAuthorizationScopeMapOutput) MapIndex(k pulumi.StringInput) ClientAuthorizationScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClientAuthorizationScope {
		return vs[0].(map[string]ClientAuthorizationScope)[vs[1].(string)]
	}).(ClientAuthorizationScopeOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientAuthorizationScopeOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationScopePtrOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationScopeArrayOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationScopeMapOutput{})
}
