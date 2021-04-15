// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing MSAD user account control mappers for Keycloak
// users federated via LDAP.
//
// The MSAD (Microsoft Active Directory) user account control mapper is specific
// to LDAP user federation providers that are pulling from AD, and it can propagate
// AD user state to Keycloak in order to enforce settings like expired passwords
// or disabled accounts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/ldap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		ldapUserFederation, err := ldap.NewUserFederation(ctx, "ldapUserFederation", &ldap.UserFederationArgs{
// 			RealmId:               realm.ID(),
// 			UsernameLdapAttribute: pulumi.String("cn"),
// 			RdnLdapAttribute:      pulumi.String("cn"),
// 			UuidLdapAttribute:     pulumi.String("objectGUID"),
// 			UserObjectClasses: pulumi.StringArray{
// 				pulumi.String("person"),
// 				pulumi.String("organizationalPerson"),
// 				pulumi.String("user"),
// 			},
// 			ConnectionUrl:  pulumi.String("ldap://my-ad-server"),
// 			UsersDn:        pulumi.String("dc=example,dc=org"),
// 			BindDn:         pulumi.String("cn=admin,dc=example,dc=org"),
// 			BindCredential: pulumi.String("admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ldap.NewMsadUserAccountControlMapper(ctx, "msadUserAccountControlMapper", &ldap.MsadUserAccountControlMapperArgs{
// 			RealmId:              realm.ID(),
// 			LdapUserFederationId: ldapUserFederation.ID(),
// 		})
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
// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
//
// ```sh
//  $ pulumi import keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper msad_user_account_control_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
// ```
type MsadUserAccountControlMapper struct {
	pulumi.CustomResourceState

	// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
	LdapPasswordPolicyHintsEnabled pulumi.BoolPtrOutput `pulumi:"ldapPasswordPolicyHintsEnabled"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewMsadUserAccountControlMapper registers a new resource with the given unique name, arguments, and options.
func NewMsadUserAccountControlMapper(ctx *pulumi.Context,
	name string, args *MsadUserAccountControlMapperArgs, opts ...pulumi.ResourceOption) (*MsadUserAccountControlMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource MsadUserAccountControlMapper
	err := ctx.RegisterResource("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMsadUserAccountControlMapper gets an existing MsadUserAccountControlMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMsadUserAccountControlMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MsadUserAccountControlMapperState, opts ...pulumi.ResourceOption) (*MsadUserAccountControlMapper, error) {
	var resource MsadUserAccountControlMapper
	err := ctx.ReadResource("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MsadUserAccountControlMapper resources.
type msadUserAccountControlMapperState struct {
	// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
	LdapPasswordPolicyHintsEnabled *bool `pulumi:"ldapPasswordPolicyHintsEnabled"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
}

type MsadUserAccountControlMapperState struct {
	// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
	LdapPasswordPolicyHintsEnabled pulumi.BoolPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
}

func (MsadUserAccountControlMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*msadUserAccountControlMapperState)(nil)).Elem()
}

type msadUserAccountControlMapperArgs struct {
	// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
	LdapPasswordPolicyHintsEnabled *bool `pulumi:"ldapPasswordPolicyHintsEnabled"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a MsadUserAccountControlMapper resource.
type MsadUserAccountControlMapperArgs struct {
	// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
	LdapPasswordPolicyHintsEnabled pulumi.BoolPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
}

func (MsadUserAccountControlMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*msadUserAccountControlMapperArgs)(nil)).Elem()
}

type MsadUserAccountControlMapperInput interface {
	pulumi.Input

	ToMsadUserAccountControlMapperOutput() MsadUserAccountControlMapperOutput
	ToMsadUserAccountControlMapperOutputWithContext(ctx context.Context) MsadUserAccountControlMapperOutput
}

func (*MsadUserAccountControlMapper) ElementType() reflect.Type {
	return reflect.TypeOf((*MsadUserAccountControlMapper)(nil))
}

func (i *MsadUserAccountControlMapper) ToMsadUserAccountControlMapperOutput() MsadUserAccountControlMapperOutput {
	return i.ToMsadUserAccountControlMapperOutputWithContext(context.Background())
}

func (i *MsadUserAccountControlMapper) ToMsadUserAccountControlMapperOutputWithContext(ctx context.Context) MsadUserAccountControlMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadUserAccountControlMapperOutput)
}

func (i *MsadUserAccountControlMapper) ToMsadUserAccountControlMapperPtrOutput() MsadUserAccountControlMapperPtrOutput {
	return i.ToMsadUserAccountControlMapperPtrOutputWithContext(context.Background())
}

func (i *MsadUserAccountControlMapper) ToMsadUserAccountControlMapperPtrOutputWithContext(ctx context.Context) MsadUserAccountControlMapperPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadUserAccountControlMapperPtrOutput)
}

type MsadUserAccountControlMapperPtrInput interface {
	pulumi.Input

	ToMsadUserAccountControlMapperPtrOutput() MsadUserAccountControlMapperPtrOutput
	ToMsadUserAccountControlMapperPtrOutputWithContext(ctx context.Context) MsadUserAccountControlMapperPtrOutput
}

type msadUserAccountControlMapperPtrType MsadUserAccountControlMapperArgs

func (*msadUserAccountControlMapperPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MsadUserAccountControlMapper)(nil))
}

func (i *msadUserAccountControlMapperPtrType) ToMsadUserAccountControlMapperPtrOutput() MsadUserAccountControlMapperPtrOutput {
	return i.ToMsadUserAccountControlMapperPtrOutputWithContext(context.Background())
}

func (i *msadUserAccountControlMapperPtrType) ToMsadUserAccountControlMapperPtrOutputWithContext(ctx context.Context) MsadUserAccountControlMapperPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadUserAccountControlMapperPtrOutput)
}

// MsadUserAccountControlMapperArrayInput is an input type that accepts MsadUserAccountControlMapperArray and MsadUserAccountControlMapperArrayOutput values.
// You can construct a concrete instance of `MsadUserAccountControlMapperArrayInput` via:
//
//          MsadUserAccountControlMapperArray{ MsadUserAccountControlMapperArgs{...} }
type MsadUserAccountControlMapperArrayInput interface {
	pulumi.Input

	ToMsadUserAccountControlMapperArrayOutput() MsadUserAccountControlMapperArrayOutput
	ToMsadUserAccountControlMapperArrayOutputWithContext(context.Context) MsadUserAccountControlMapperArrayOutput
}

type MsadUserAccountControlMapperArray []MsadUserAccountControlMapperInput

func (MsadUserAccountControlMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MsadUserAccountControlMapper)(nil))
}

func (i MsadUserAccountControlMapperArray) ToMsadUserAccountControlMapperArrayOutput() MsadUserAccountControlMapperArrayOutput {
	return i.ToMsadUserAccountControlMapperArrayOutputWithContext(context.Background())
}

func (i MsadUserAccountControlMapperArray) ToMsadUserAccountControlMapperArrayOutputWithContext(ctx context.Context) MsadUserAccountControlMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadUserAccountControlMapperArrayOutput)
}

// MsadUserAccountControlMapperMapInput is an input type that accepts MsadUserAccountControlMapperMap and MsadUserAccountControlMapperMapOutput values.
// You can construct a concrete instance of `MsadUserAccountControlMapperMapInput` via:
//
//          MsadUserAccountControlMapperMap{ "key": MsadUserAccountControlMapperArgs{...} }
type MsadUserAccountControlMapperMapInput interface {
	pulumi.Input

	ToMsadUserAccountControlMapperMapOutput() MsadUserAccountControlMapperMapOutput
	ToMsadUserAccountControlMapperMapOutputWithContext(context.Context) MsadUserAccountControlMapperMapOutput
}

type MsadUserAccountControlMapperMap map[string]MsadUserAccountControlMapperInput

func (MsadUserAccountControlMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MsadUserAccountControlMapper)(nil))
}

func (i MsadUserAccountControlMapperMap) ToMsadUserAccountControlMapperMapOutput() MsadUserAccountControlMapperMapOutput {
	return i.ToMsadUserAccountControlMapperMapOutputWithContext(context.Background())
}

func (i MsadUserAccountControlMapperMap) ToMsadUserAccountControlMapperMapOutputWithContext(ctx context.Context) MsadUserAccountControlMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadUserAccountControlMapperMapOutput)
}

type MsadUserAccountControlMapperOutput struct {
	*pulumi.OutputState
}

func (MsadUserAccountControlMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsadUserAccountControlMapper)(nil))
}

func (o MsadUserAccountControlMapperOutput) ToMsadUserAccountControlMapperOutput() MsadUserAccountControlMapperOutput {
	return o
}

func (o MsadUserAccountControlMapperOutput) ToMsadUserAccountControlMapperOutputWithContext(ctx context.Context) MsadUserAccountControlMapperOutput {
	return o
}

func (o MsadUserAccountControlMapperOutput) ToMsadUserAccountControlMapperPtrOutput() MsadUserAccountControlMapperPtrOutput {
	return o.ToMsadUserAccountControlMapperPtrOutputWithContext(context.Background())
}

func (o MsadUserAccountControlMapperOutput) ToMsadUserAccountControlMapperPtrOutputWithContext(ctx context.Context) MsadUserAccountControlMapperPtrOutput {
	return o.ApplyT(func(v MsadUserAccountControlMapper) *MsadUserAccountControlMapper {
		return &v
	}).(MsadUserAccountControlMapperPtrOutput)
}

type MsadUserAccountControlMapperPtrOutput struct {
	*pulumi.OutputState
}

func (MsadUserAccountControlMapperPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MsadUserAccountControlMapper)(nil))
}

func (o MsadUserAccountControlMapperPtrOutput) ToMsadUserAccountControlMapperPtrOutput() MsadUserAccountControlMapperPtrOutput {
	return o
}

func (o MsadUserAccountControlMapperPtrOutput) ToMsadUserAccountControlMapperPtrOutputWithContext(ctx context.Context) MsadUserAccountControlMapperPtrOutput {
	return o
}

type MsadUserAccountControlMapperArrayOutput struct{ *pulumi.OutputState }

func (MsadUserAccountControlMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsadUserAccountControlMapper)(nil))
}

func (o MsadUserAccountControlMapperArrayOutput) ToMsadUserAccountControlMapperArrayOutput() MsadUserAccountControlMapperArrayOutput {
	return o
}

func (o MsadUserAccountControlMapperArrayOutput) ToMsadUserAccountControlMapperArrayOutputWithContext(ctx context.Context) MsadUserAccountControlMapperArrayOutput {
	return o
}

func (o MsadUserAccountControlMapperArrayOutput) Index(i pulumi.IntInput) MsadUserAccountControlMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsadUserAccountControlMapper {
		return vs[0].([]MsadUserAccountControlMapper)[vs[1].(int)]
	}).(MsadUserAccountControlMapperOutput)
}

type MsadUserAccountControlMapperMapOutput struct{ *pulumi.OutputState }

func (MsadUserAccountControlMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MsadUserAccountControlMapper)(nil))
}

func (o MsadUserAccountControlMapperMapOutput) ToMsadUserAccountControlMapperMapOutput() MsadUserAccountControlMapperMapOutput {
	return o
}

func (o MsadUserAccountControlMapperMapOutput) ToMsadUserAccountControlMapperMapOutputWithContext(ctx context.Context) MsadUserAccountControlMapperMapOutput {
	return o
}

func (o MsadUserAccountControlMapperMapOutput) MapIndex(k pulumi.StringInput) MsadUserAccountControlMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MsadUserAccountControlMapper {
		return vs[0].(map[string]MsadUserAccountControlMapper)[vs[1].(string)]
	}).(MsadUserAccountControlMapperOutput)
}

func init() {
	pulumi.RegisterOutputType(MsadUserAccountControlMapperOutput{})
	pulumi.RegisterOutputType(MsadUserAccountControlMapperPtrOutput{})
	pulumi.RegisterOutputType(MsadUserAccountControlMapperArrayOutput{})
	pulumi.RegisterOutputType(MsadUserAccountControlMapperMapOutput{})
}
