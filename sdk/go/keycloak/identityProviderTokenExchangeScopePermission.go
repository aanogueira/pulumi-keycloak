// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/oidc"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tokenExchangeRealm, err := keycloak.NewRealm(ctx, "tokenExchangeRealm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("token-exchange_destination_realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tokenExchangeMyOidcIdp, err := oidc.NewIdentityProvider(ctx, "tokenExchangeMyOidcIdp", &oidc.IdentityProviderArgs{
// 			Realm:            tokenExchangeRealm.ID(),
// 			Alias:            pulumi.String("myIdp"),
// 			AuthorizationUrl: pulumi.String("http://localhost:8080/auth/realms/someRealm/protocol/openid-connect/auth"),
// 			TokenUrl:         pulumi.String("http://localhost:8080/auth/realms/someRealm/protocol/openid-connect/token"),
// 			ClientId:         pulumi.String("clientId"),
// 			ClientSecret:     pulumi.String("secret"),
// 			DefaultScopes:    pulumi.String("openid"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewClient(ctx, "token-exchangeWebappClient", &openid.ClientArgs{
// 			RealmId:             tokenExchangeRealm.ID(),
// 			ClientId:            pulumi.String("webapp_client"),
// 			ClientSecret:        pulumi.String("secret"),
// 			Description:         pulumi.String("a webapp client on the destination realm"),
// 			AccessType:          pulumi.String("CONFIDENTIAL"),
// 			StandardFlowEnabled: pulumi.Bool(true),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/*"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewIdentityProviderTokenExchangeScopePermission(ctx, "oidcIdpPermission", &keycloak.IdentityProviderTokenExchangeScopePermissionArgs{
// 			RealmId:       tokenExchangeRealm.ID(),
// 			ProviderAlias: tokenExchangeMyOidcIdp.Alias,
// 			PolicyType:    pulumi.String("client"),
// 			Clients: pulumi.StringArray{
// 				token_exchangeWebappClient.ID(),
// 			},
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
// This resource can be imported using the format `{{realm_id}}/{{provider_alias}}`, where `provider_alias` is the alias that you assign to the identity provider upon creation. Examplebash
//
// ```sh
//  $ pulumi import keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission oidc_idp_permission my-realm/myIdp
// ```
type IdentityProviderTokenExchangeScopePermission struct {
	pulumi.CustomResourceState

	// (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
	AuthorizationIdpResourceId pulumi.StringOutput `pulumi:"authorizationIdpResourceId"`
	// (Computed) Resource server ID representing the realm management client on which this permission is managed.
	AuthorizationResourceServerId pulumi.StringOutput `pulumi:"authorizationResourceServerId"`
	// (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
	AuthorizationTokenExchangeScopePermissionId pulumi.StringOutput `pulumi:"authorizationTokenExchangeScopePermissionId"`
	// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
	Clients pulumi.StringArrayOutput `pulumi:"clients"`
	// (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// Defaults to "client" This is also the only value policy type supported by this provider.
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// Alias of the identity provider.
	ProviderAlias pulumi.StringOutput `pulumi:"providerAlias"`
	// The realm that the identity provider exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewIdentityProviderTokenExchangeScopePermission registers a new resource with the given unique name, arguments, and options.
func NewIdentityProviderTokenExchangeScopePermission(ctx *pulumi.Context,
	name string, args *IdentityProviderTokenExchangeScopePermissionArgs, opts ...pulumi.ResourceOption) (*IdentityProviderTokenExchangeScopePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Clients == nil {
		return nil, errors.New("invalid value for required argument 'Clients'")
	}
	if args.ProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'ProviderAlias'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource IdentityProviderTokenExchangeScopePermission
	err := ctx.RegisterResource("keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProviderTokenExchangeScopePermission gets an existing IdentityProviderTokenExchangeScopePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProviderTokenExchangeScopePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderTokenExchangeScopePermissionState, opts ...pulumi.ResourceOption) (*IdentityProviderTokenExchangeScopePermission, error) {
	var resource IdentityProviderTokenExchangeScopePermission
	err := ctx.ReadResource("keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProviderTokenExchangeScopePermission resources.
type identityProviderTokenExchangeScopePermissionState struct {
	// (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
	AuthorizationIdpResourceId *string `pulumi:"authorizationIdpResourceId"`
	// (Computed) Resource server ID representing the realm management client on which this permission is managed.
	AuthorizationResourceServerId *string `pulumi:"authorizationResourceServerId"`
	// (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
	AuthorizationTokenExchangeScopePermissionId *string `pulumi:"authorizationTokenExchangeScopePermissionId"`
	// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
	Clients []string `pulumi:"clients"`
	// (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
	PolicyId *string `pulumi:"policyId"`
	// Defaults to "client" This is also the only value policy type supported by this provider.
	PolicyType *string `pulumi:"policyType"`
	// Alias of the identity provider.
	ProviderAlias *string `pulumi:"providerAlias"`
	// The realm that the identity provider exists in.
	RealmId *string `pulumi:"realmId"`
}

type IdentityProviderTokenExchangeScopePermissionState struct {
	// (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
	AuthorizationIdpResourceId pulumi.StringPtrInput
	// (Computed) Resource server ID representing the realm management client on which this permission is managed.
	AuthorizationResourceServerId pulumi.StringPtrInput
	// (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
	AuthorizationTokenExchangeScopePermissionId pulumi.StringPtrInput
	// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
	Clients pulumi.StringArrayInput
	// (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
	PolicyId pulumi.StringPtrInput
	// Defaults to "client" This is also the only value policy type supported by this provider.
	PolicyType pulumi.StringPtrInput
	// Alias of the identity provider.
	ProviderAlias pulumi.StringPtrInput
	// The realm that the identity provider exists in.
	RealmId pulumi.StringPtrInput
}

func (IdentityProviderTokenExchangeScopePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderTokenExchangeScopePermissionState)(nil)).Elem()
}

type identityProviderTokenExchangeScopePermissionArgs struct {
	// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
	Clients []string `pulumi:"clients"`
	// Defaults to "client" This is also the only value policy type supported by this provider.
	PolicyType *string `pulumi:"policyType"`
	// Alias of the identity provider.
	ProviderAlias string `pulumi:"providerAlias"`
	// The realm that the identity provider exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a IdentityProviderTokenExchangeScopePermission resource.
type IdentityProviderTokenExchangeScopePermissionArgs struct {
	// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
	Clients pulumi.StringArrayInput
	// Defaults to "client" This is also the only value policy type supported by this provider.
	PolicyType pulumi.StringPtrInput
	// Alias of the identity provider.
	ProviderAlias pulumi.StringInput
	// The realm that the identity provider exists in.
	RealmId pulumi.StringInput
}

func (IdentityProviderTokenExchangeScopePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderTokenExchangeScopePermissionArgs)(nil)).Elem()
}

type IdentityProviderTokenExchangeScopePermissionInput interface {
	pulumi.Input

	ToIdentityProviderTokenExchangeScopePermissionOutput() IdentityProviderTokenExchangeScopePermissionOutput
	ToIdentityProviderTokenExchangeScopePermissionOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionOutput
}

func (*IdentityProviderTokenExchangeScopePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderTokenExchangeScopePermission)(nil)).Elem()
}

func (i *IdentityProviderTokenExchangeScopePermission) ToIdentityProviderTokenExchangeScopePermissionOutput() IdentityProviderTokenExchangeScopePermissionOutput {
	return i.ToIdentityProviderTokenExchangeScopePermissionOutputWithContext(context.Background())
}

func (i *IdentityProviderTokenExchangeScopePermission) ToIdentityProviderTokenExchangeScopePermissionOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTokenExchangeScopePermissionOutput)
}

// IdentityProviderTokenExchangeScopePermissionArrayInput is an input type that accepts IdentityProviderTokenExchangeScopePermissionArray and IdentityProviderTokenExchangeScopePermissionArrayOutput values.
// You can construct a concrete instance of `IdentityProviderTokenExchangeScopePermissionArrayInput` via:
//
//          IdentityProviderTokenExchangeScopePermissionArray{ IdentityProviderTokenExchangeScopePermissionArgs{...} }
type IdentityProviderTokenExchangeScopePermissionArrayInput interface {
	pulumi.Input

	ToIdentityProviderTokenExchangeScopePermissionArrayOutput() IdentityProviderTokenExchangeScopePermissionArrayOutput
	ToIdentityProviderTokenExchangeScopePermissionArrayOutputWithContext(context.Context) IdentityProviderTokenExchangeScopePermissionArrayOutput
}

type IdentityProviderTokenExchangeScopePermissionArray []IdentityProviderTokenExchangeScopePermissionInput

func (IdentityProviderTokenExchangeScopePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityProviderTokenExchangeScopePermission)(nil)).Elem()
}

func (i IdentityProviderTokenExchangeScopePermissionArray) ToIdentityProviderTokenExchangeScopePermissionArrayOutput() IdentityProviderTokenExchangeScopePermissionArrayOutput {
	return i.ToIdentityProviderTokenExchangeScopePermissionArrayOutputWithContext(context.Background())
}

func (i IdentityProviderTokenExchangeScopePermissionArray) ToIdentityProviderTokenExchangeScopePermissionArrayOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTokenExchangeScopePermissionArrayOutput)
}

// IdentityProviderTokenExchangeScopePermissionMapInput is an input type that accepts IdentityProviderTokenExchangeScopePermissionMap and IdentityProviderTokenExchangeScopePermissionMapOutput values.
// You can construct a concrete instance of `IdentityProviderTokenExchangeScopePermissionMapInput` via:
//
//          IdentityProviderTokenExchangeScopePermissionMap{ "key": IdentityProviderTokenExchangeScopePermissionArgs{...} }
type IdentityProviderTokenExchangeScopePermissionMapInput interface {
	pulumi.Input

	ToIdentityProviderTokenExchangeScopePermissionMapOutput() IdentityProviderTokenExchangeScopePermissionMapOutput
	ToIdentityProviderTokenExchangeScopePermissionMapOutputWithContext(context.Context) IdentityProviderTokenExchangeScopePermissionMapOutput
}

type IdentityProviderTokenExchangeScopePermissionMap map[string]IdentityProviderTokenExchangeScopePermissionInput

func (IdentityProviderTokenExchangeScopePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityProviderTokenExchangeScopePermission)(nil)).Elem()
}

func (i IdentityProviderTokenExchangeScopePermissionMap) ToIdentityProviderTokenExchangeScopePermissionMapOutput() IdentityProviderTokenExchangeScopePermissionMapOutput {
	return i.ToIdentityProviderTokenExchangeScopePermissionMapOutputWithContext(context.Background())
}

func (i IdentityProviderTokenExchangeScopePermissionMap) ToIdentityProviderTokenExchangeScopePermissionMapOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderTokenExchangeScopePermissionMapOutput)
}

type IdentityProviderTokenExchangeScopePermissionOutput struct{ *pulumi.OutputState }

func (IdentityProviderTokenExchangeScopePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderTokenExchangeScopePermission)(nil)).Elem()
}

func (o IdentityProviderTokenExchangeScopePermissionOutput) ToIdentityProviderTokenExchangeScopePermissionOutput() IdentityProviderTokenExchangeScopePermissionOutput {
	return o
}

func (o IdentityProviderTokenExchangeScopePermissionOutput) ToIdentityProviderTokenExchangeScopePermissionOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionOutput {
	return o
}

type IdentityProviderTokenExchangeScopePermissionArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderTokenExchangeScopePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityProviderTokenExchangeScopePermission)(nil)).Elem()
}

func (o IdentityProviderTokenExchangeScopePermissionArrayOutput) ToIdentityProviderTokenExchangeScopePermissionArrayOutput() IdentityProviderTokenExchangeScopePermissionArrayOutput {
	return o
}

func (o IdentityProviderTokenExchangeScopePermissionArrayOutput) ToIdentityProviderTokenExchangeScopePermissionArrayOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionArrayOutput {
	return o
}

func (o IdentityProviderTokenExchangeScopePermissionArrayOutput) Index(i pulumi.IntInput) IdentityProviderTokenExchangeScopePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityProviderTokenExchangeScopePermission {
		return vs[0].([]*IdentityProviderTokenExchangeScopePermission)[vs[1].(int)]
	}).(IdentityProviderTokenExchangeScopePermissionOutput)
}

type IdentityProviderTokenExchangeScopePermissionMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderTokenExchangeScopePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityProviderTokenExchangeScopePermission)(nil)).Elem()
}

func (o IdentityProviderTokenExchangeScopePermissionMapOutput) ToIdentityProviderTokenExchangeScopePermissionMapOutput() IdentityProviderTokenExchangeScopePermissionMapOutput {
	return o
}

func (o IdentityProviderTokenExchangeScopePermissionMapOutput) ToIdentityProviderTokenExchangeScopePermissionMapOutputWithContext(ctx context.Context) IdentityProviderTokenExchangeScopePermissionMapOutput {
	return o
}

func (o IdentityProviderTokenExchangeScopePermissionMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderTokenExchangeScopePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityProviderTokenExchangeScopePermission {
		return vs[0].(map[string]*IdentityProviderTokenExchangeScopePermission)[vs[1].(string)]
	}).(IdentityProviderTokenExchangeScopePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderTokenExchangeScopePermissionInput)(nil)).Elem(), &IdentityProviderTokenExchangeScopePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderTokenExchangeScopePermissionArrayInput)(nil)).Elem(), IdentityProviderTokenExchangeScopePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderTokenExchangeScopePermissionMapInput)(nil)).Elem(), IdentityProviderTokenExchangeScopePermissionMap{})
	pulumi.RegisterOutputType(IdentityProviderTokenExchangeScopePermissionOutput{})
	pulumi.RegisterOutputType(IdentityProviderTokenExchangeScopePermissionArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderTokenExchangeScopePermissionMapOutput{})
}
