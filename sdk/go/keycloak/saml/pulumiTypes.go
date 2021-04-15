// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientAuthenticationFlowBindingOverrides struct {
	// Browser flow id, (flow needs to exist)
	BrowserId *string `pulumi:"browserId"`
	// Direct grant flow id (flow needs to exist)
	DirectGrantId *string `pulumi:"directGrantId"`
}

// ClientAuthenticationFlowBindingOverridesInput is an input type that accepts ClientAuthenticationFlowBindingOverridesArgs and ClientAuthenticationFlowBindingOverridesOutput values.
// You can construct a concrete instance of `ClientAuthenticationFlowBindingOverridesInput` via:
//
//          ClientAuthenticationFlowBindingOverridesArgs{...}
type ClientAuthenticationFlowBindingOverridesInput interface {
	pulumi.Input

	ToClientAuthenticationFlowBindingOverridesOutput() ClientAuthenticationFlowBindingOverridesOutput
	ToClientAuthenticationFlowBindingOverridesOutputWithContext(context.Context) ClientAuthenticationFlowBindingOverridesOutput
}

type ClientAuthenticationFlowBindingOverridesArgs struct {
	// Browser flow id, (flow needs to exist)
	BrowserId pulumi.StringPtrInput `pulumi:"browserId"`
	// Direct grant flow id (flow needs to exist)
	DirectGrantId pulumi.StringPtrInput `pulumi:"directGrantId"`
}

func (ClientAuthenticationFlowBindingOverridesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesOutput() ClientAuthenticationFlowBindingOverridesOutput {
	return i.ToClientAuthenticationFlowBindingOverridesOutputWithContext(context.Background())
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthenticationFlowBindingOverridesOutput)
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return i.ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Background())
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthenticationFlowBindingOverridesOutput).ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx)
}

// ClientAuthenticationFlowBindingOverridesPtrInput is an input type that accepts ClientAuthenticationFlowBindingOverridesArgs, ClientAuthenticationFlowBindingOverridesPtr and ClientAuthenticationFlowBindingOverridesPtrOutput values.
// You can construct a concrete instance of `ClientAuthenticationFlowBindingOverridesPtrInput` via:
//
//          ClientAuthenticationFlowBindingOverridesArgs{...}
//
//  or:
//
//          nil
type ClientAuthenticationFlowBindingOverridesPtrInput interface {
	pulumi.Input

	ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput
	ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput
}

type clientAuthenticationFlowBindingOverridesPtrType ClientAuthenticationFlowBindingOverridesArgs

func ClientAuthenticationFlowBindingOverridesPtr(v *ClientAuthenticationFlowBindingOverridesArgs) ClientAuthenticationFlowBindingOverridesPtrInput {
	return (*clientAuthenticationFlowBindingOverridesPtrType)(v)
}

func (*clientAuthenticationFlowBindingOverridesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (i *clientAuthenticationFlowBindingOverridesPtrType) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return i.ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Background())
}

func (i *clientAuthenticationFlowBindingOverridesPtrType) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthenticationFlowBindingOverridesPtrOutput)
}

type ClientAuthenticationFlowBindingOverridesOutput struct{ *pulumi.OutputState }

func (ClientAuthenticationFlowBindingOverridesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesOutput() ClientAuthenticationFlowBindingOverridesOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o.ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Background())
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o.ApplyT(func(v ClientAuthenticationFlowBindingOverrides) *ClientAuthenticationFlowBindingOverrides {
		return &v
	}).(ClientAuthenticationFlowBindingOverridesPtrOutput)
}

// Browser flow id, (flow needs to exist)
func (o ClientAuthenticationFlowBindingOverridesOutput) BrowserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientAuthenticationFlowBindingOverrides) *string { return v.BrowserId }).(pulumi.StringPtrOutput)
}

// Direct grant flow id (flow needs to exist)
func (o ClientAuthenticationFlowBindingOverridesOutput) DirectGrantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientAuthenticationFlowBindingOverrides) *string { return v.DirectGrantId }).(pulumi.StringPtrOutput)
}

type ClientAuthenticationFlowBindingOverridesPtrOutput struct{ *pulumi.OutputState }

func (ClientAuthenticationFlowBindingOverridesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) Elem() ClientAuthenticationFlowBindingOverridesOutput {
	return o.ApplyT(func(v *ClientAuthenticationFlowBindingOverrides) ClientAuthenticationFlowBindingOverrides { return *v }).(ClientAuthenticationFlowBindingOverridesOutput)
}

// Browser flow id, (flow needs to exist)
func (o ClientAuthenticationFlowBindingOverridesPtrOutput) BrowserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthenticationFlowBindingOverrides) *string {
		if v == nil {
			return nil
		}
		return v.BrowserId
	}).(pulumi.StringPtrOutput)
}

// Direct grant flow id (flow needs to exist)
func (o ClientAuthenticationFlowBindingOverridesPtrOutput) DirectGrantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthenticationFlowBindingOverrides) *string {
		if v == nil {
			return nil
		}
		return v.DirectGrantId
	}).(pulumi.StringPtrOutput)
}

type GetClientAuthenticationFlowBindingOverride struct {
	BrowserId     string `pulumi:"browserId"`
	DirectGrantId string `pulumi:"directGrantId"`
}

// GetClientAuthenticationFlowBindingOverrideInput is an input type that accepts GetClientAuthenticationFlowBindingOverrideArgs and GetClientAuthenticationFlowBindingOverrideOutput values.
// You can construct a concrete instance of `GetClientAuthenticationFlowBindingOverrideInput` via:
//
//          GetClientAuthenticationFlowBindingOverrideArgs{...}
type GetClientAuthenticationFlowBindingOverrideInput interface {
	pulumi.Input

	ToGetClientAuthenticationFlowBindingOverrideOutput() GetClientAuthenticationFlowBindingOverrideOutput
	ToGetClientAuthenticationFlowBindingOverrideOutputWithContext(context.Context) GetClientAuthenticationFlowBindingOverrideOutput
}

type GetClientAuthenticationFlowBindingOverrideArgs struct {
	BrowserId     pulumi.StringInput `pulumi:"browserId"`
	DirectGrantId pulumi.StringInput `pulumi:"directGrantId"`
}

func (GetClientAuthenticationFlowBindingOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthenticationFlowBindingOverride)(nil)).Elem()
}

func (i GetClientAuthenticationFlowBindingOverrideArgs) ToGetClientAuthenticationFlowBindingOverrideOutput() GetClientAuthenticationFlowBindingOverrideOutput {
	return i.ToGetClientAuthenticationFlowBindingOverrideOutputWithContext(context.Background())
}

func (i GetClientAuthenticationFlowBindingOverrideArgs) ToGetClientAuthenticationFlowBindingOverrideOutputWithContext(ctx context.Context) GetClientAuthenticationFlowBindingOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClientAuthenticationFlowBindingOverrideOutput)
}

// GetClientAuthenticationFlowBindingOverrideArrayInput is an input type that accepts GetClientAuthenticationFlowBindingOverrideArray and GetClientAuthenticationFlowBindingOverrideArrayOutput values.
// You can construct a concrete instance of `GetClientAuthenticationFlowBindingOverrideArrayInput` via:
//
//          GetClientAuthenticationFlowBindingOverrideArray{ GetClientAuthenticationFlowBindingOverrideArgs{...} }
type GetClientAuthenticationFlowBindingOverrideArrayInput interface {
	pulumi.Input

	ToGetClientAuthenticationFlowBindingOverrideArrayOutput() GetClientAuthenticationFlowBindingOverrideArrayOutput
	ToGetClientAuthenticationFlowBindingOverrideArrayOutputWithContext(context.Context) GetClientAuthenticationFlowBindingOverrideArrayOutput
}

type GetClientAuthenticationFlowBindingOverrideArray []GetClientAuthenticationFlowBindingOverrideInput

func (GetClientAuthenticationFlowBindingOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClientAuthenticationFlowBindingOverride)(nil)).Elem()
}

func (i GetClientAuthenticationFlowBindingOverrideArray) ToGetClientAuthenticationFlowBindingOverrideArrayOutput() GetClientAuthenticationFlowBindingOverrideArrayOutput {
	return i.ToGetClientAuthenticationFlowBindingOverrideArrayOutputWithContext(context.Background())
}

func (i GetClientAuthenticationFlowBindingOverrideArray) ToGetClientAuthenticationFlowBindingOverrideArrayOutputWithContext(ctx context.Context) GetClientAuthenticationFlowBindingOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClientAuthenticationFlowBindingOverrideArrayOutput)
}

type GetClientAuthenticationFlowBindingOverrideOutput struct{ *pulumi.OutputState }

func (GetClientAuthenticationFlowBindingOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthenticationFlowBindingOverride)(nil)).Elem()
}

func (o GetClientAuthenticationFlowBindingOverrideOutput) ToGetClientAuthenticationFlowBindingOverrideOutput() GetClientAuthenticationFlowBindingOverrideOutput {
	return o
}

func (o GetClientAuthenticationFlowBindingOverrideOutput) ToGetClientAuthenticationFlowBindingOverrideOutputWithContext(ctx context.Context) GetClientAuthenticationFlowBindingOverrideOutput {
	return o
}

func (o GetClientAuthenticationFlowBindingOverrideOutput) BrowserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthenticationFlowBindingOverride) string { return v.BrowserId }).(pulumi.StringOutput)
}

func (o GetClientAuthenticationFlowBindingOverrideOutput) DirectGrantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthenticationFlowBindingOverride) string { return v.DirectGrantId }).(pulumi.StringOutput)
}

type GetClientAuthenticationFlowBindingOverrideArrayOutput struct{ *pulumi.OutputState }

func (GetClientAuthenticationFlowBindingOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClientAuthenticationFlowBindingOverride)(nil)).Elem()
}

func (o GetClientAuthenticationFlowBindingOverrideArrayOutput) ToGetClientAuthenticationFlowBindingOverrideArrayOutput() GetClientAuthenticationFlowBindingOverrideArrayOutput {
	return o
}

func (o GetClientAuthenticationFlowBindingOverrideArrayOutput) ToGetClientAuthenticationFlowBindingOverrideArrayOutputWithContext(ctx context.Context) GetClientAuthenticationFlowBindingOverrideArrayOutput {
	return o
}

func (o GetClientAuthenticationFlowBindingOverrideArrayOutput) Index(i pulumi.IntInput) GetClientAuthenticationFlowBindingOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetClientAuthenticationFlowBindingOverride {
		return vs[0].([]GetClientAuthenticationFlowBindingOverride)[vs[1].(int)]
	}).(GetClientAuthenticationFlowBindingOverrideOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientAuthenticationFlowBindingOverridesOutput{})
	pulumi.RegisterOutputType(ClientAuthenticationFlowBindingOverridesPtrOutput{})
	pulumi.RegisterOutputType(GetClientAuthenticationFlowBindingOverrideOutput{})
	pulumi.RegisterOutputType(GetClientAuthenticationFlowBindingOverrideArrayOutput{})
}
