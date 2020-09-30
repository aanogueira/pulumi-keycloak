// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing audience protocol mappers within Keycloak.
//
// Audience protocol mappers allow you add audiences to the `aud` claim within issued tokens. The audience can be a custom
// string, or it can be mapped to the ID of a pre-existing client.
//
// ## Example Usage
// ### Client)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/openid"
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
// 		openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
// 			RealmId:    realm.ID(),
// 			ClientId:   pulumi.String("client"),
// 			Enabled:    pulumi.Bool(true),
// 			AccessType: pulumi.String("CONFIDENTIAL"),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/openid-callback"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewAudienceProtocolMapper(ctx, "audienceMapper", &openid.AudienceProtocolMapperArgs{
// 			RealmId:                realm.ID(),
// 			ClientId:               openidClient.ID(),
// 			IncludedCustomAudience: pulumi.String("foo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Client Scope)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/openid"
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
// 		clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewAudienceProtocolMapper(ctx, "audienceMapper", &openid.AudienceProtocolMapperArgs{
// 			RealmId:                realm.ID(),
// 			ClientScopeId:          clientScope.ID(),
// 			IncludedCustomAudience: pulumi.String("foo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AudienceProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedClientAudience pulumi.StringPtrOutput `pulumi:"includedClientAudience"`
	// A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedCustomAudience pulumi.StringPtrOutput `pulumi:"includedCustomAudience"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewAudienceProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewAudienceProtocolMapper(ctx *pulumi.Context,
	name string, args *AudienceProtocolMapperArgs, opts ...pulumi.ResourceOption) (*AudienceProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &AudienceProtocolMapperArgs{}
	}
	var resource AudienceProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAudienceProtocolMapper gets an existing AudienceProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAudienceProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AudienceProtocolMapperState, opts ...pulumi.ResourceOption) (*AudienceProtocolMapper, error) {
	var resource AudienceProtocolMapper
	err := ctx.ReadResource("keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AudienceProtocolMapper resources.
type audienceProtocolMapperState struct {
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedClientAudience *string `pulumi:"includedClientAudience"`
	// A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedCustomAudience *string `pulumi:"includedCustomAudience"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type AudienceProtocolMapperState struct {
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedClientAudience pulumi.StringPtrInput
	// A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedCustomAudience pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (AudienceProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*audienceProtocolMapperState)(nil)).Elem()
}

type audienceProtocolMapperArgs struct {
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedClientAudience *string `pulumi:"includedClientAudience"`
	// A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedCustomAudience *string `pulumi:"includedCustomAudience"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a AudienceProtocolMapper resource.
type AudienceProtocolMapperArgs struct {
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedClientAudience pulumi.StringPtrInput
	// A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
	IncludedCustomAudience pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
}

func (AudienceProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*audienceProtocolMapperArgs)(nil)).Elem()
}
