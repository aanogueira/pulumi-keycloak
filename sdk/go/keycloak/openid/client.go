// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing Keycloak clients that use the OpenID Connect protocol.
//
// Clients are entities that can use Keycloak for user authentication. Typically,
// clients are applications that redirect users to Keycloak for authentication
// in order to take advantage of Keycloak's user sessions for SSO.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/openid"
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
// 		_, err = openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
// 			RealmId:    realm.ID(),
// 			ClientId:   pulumi.String("test-client"),
// 			Enabled:    pulumi.Bool(true),
// 			AccessType: pulumi.String("CONFIDENTIAL"),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/openid-callback"),
// 			},
// 			LoginTheme: pulumi.String("keycloak"),
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
// Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID. Examplebash
//
// ```sh
//  $ pulumi import keycloak:openid/client:Client openid_client my-realm/dcbc4c73-e478-4928-ae2e-d5e420223352
// ```
type Client struct {
	pulumi.CustomResourceState

	// The amount of time in seconds before an access token expires. This will override the default for the realm.
	AccessTokenLifespan pulumi.StringPtrOutput `pulumi:"accessTokenLifespan"`
	// Specifies the type of client, which can be one of the following:
	AccessType pulumi.StringOutput `pulumi:"accessType"`
	// URL to the admin interface of the client.
	AdminUrl pulumi.StringPtrOutput `pulumi:"adminUrl"`
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides ClientAuthenticationFlowBindingOverridesPtrOutput `pulumi:"authenticationFlowBindingOverrides"`
	// When this block is present, fine-grained authorization will be enabled for this client. The client's `accessType` must be `CONFIDENTIAL`, and `serviceAccountsEnabled` must be `true`. This block has the following arguments:
	Authorization ClientAuthorizationPtrOutput `pulumi:"authorization"`
	// Default URL to use when the auth server needs to redirect or link back to the client.
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Time a client session is allowed to be idle before it expires. Tokens are invalidated when a client session is expired. If not set it uses the standard SSO Session Idle value.
	ClientOfflineSessionIdleTimeout pulumi.StringPtrOutput `pulumi:"clientOfflineSessionIdleTimeout"`
	// Max time before a client session is expired. Tokens are invalidated when a client session is expired. If not set, it uses the standard SSO Session Max value.
	ClientOfflineSessionMaxLifespan pulumi.StringPtrOutput `pulumi:"clientOfflineSessionMaxLifespan"`
	// The secret for clients with an `accessType` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// Time a client offline session is allowed to be idle before it expires. Offline tokens are invalidated when a client offline session is expired. If not set it uses the Offline Session Idle value.
	ClientSessionIdleTimeout pulumi.StringPtrOutput `pulumi:"clientSessionIdleTimeout"`
	// Max time before a client offline session is expired. Offline tokens are invalidated when a client offline session is expired. If not set, it uses the Offline Session Max value.
	ClientSessionMaxLifespan pulumi.StringPtrOutput `pulumi:"clientSessionMaxLifespan"`
	// When `true`, users have to consent to client access.
	ConsentRequired pulumi.BoolPtrOutput `pulumi:"consentRequired"`
	// The description of this client in the GUI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
	DirectAccessGrantsEnabled pulumi.BoolPtrOutput `pulumi:"directAccessGrantsEnabled"`
	// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When `true`, the parameter `sessionState` will not be included in OpenID Connect Authentication Response.
	ExcludeSessionStateFromAuthResponse pulumi.BoolPtrOutput `pulumi:"excludeSessionStateFromAuthResponse"`
	// Allow to include all roles mappings in the access token.
	FullScopeAllowed pulumi.BoolPtrOutput `pulumi:"fullScopeAllowed"`
	// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
	ImplicitFlowEnabled pulumi.BoolPtrOutput `pulumi:"implicitFlowEnabled"`
	// The client login theme. This will override the default theme for the realm.
	LoginTheme pulumi.StringPtrOutput `pulumi:"loginTheme"`
	// The display name of this client in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
	PkceCodeChallengeMethod pulumi.StringPtrOutput `pulumi:"pkceCodeChallengeMethod"`
	// The realm this client is attached to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// (Computed) When authorization is enabled for this client, this attribute is the unique ID for the client (the same value as the `.id` attribute).
	ResourceServerId pulumi.StringOutput `pulumi:"resourceServerId"`
	// When specified, this URL is prepended to any relative URLs found within `validRedirectUris`, `webOrigins`, and `adminUrl`. NOTE: Due to limitations in the Keycloak API, when the `rootUrl` attribute is used, the `validRedirectUris`, `webOrigins`, and `adminUrl` attributes will be required.
	RootUrl pulumi.StringPtrOutput `pulumi:"rootUrl"`
	// (Computed) When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
	ServiceAccountUserId pulumi.StringOutput `pulumi:"serviceAccountUserId"`
	// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
	ServiceAccountsEnabled pulumi.BoolPtrOutput `pulumi:"serviceAccountsEnabled"`
	// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
	StandardFlowEnabled pulumi.BoolPtrOutput `pulumi:"standardFlowEnabled"`
	// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
	// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standardFlowEnabled` or `implicitFlowEnabled`
	// is set to `true`.
	ValidRedirectUris pulumi.StringArrayOutput `pulumi:"validRedirectUris"`
	// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
	WebOrigins pulumi.StringArrayOutput `pulumi:"webOrigins"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessType == nil {
		return nil, errors.New("invalid value for required argument 'AccessType'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource Client
	err := ctx.RegisterResource("keycloak:openid/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("keycloak:openid/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// The amount of time in seconds before an access token expires. This will override the default for the realm.
	AccessTokenLifespan *string `pulumi:"accessTokenLifespan"`
	// Specifies the type of client, which can be one of the following:
	AccessType *string `pulumi:"accessType"`
	// URL to the admin interface of the client.
	AdminUrl *string `pulumi:"adminUrl"`
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides *ClientAuthenticationFlowBindingOverrides `pulumi:"authenticationFlowBindingOverrides"`
	// When this block is present, fine-grained authorization will be enabled for this client. The client's `accessType` must be `CONFIDENTIAL`, and `serviceAccountsEnabled` must be `true`. This block has the following arguments:
	Authorization *ClientAuthorization `pulumi:"authorization"`
	// Default URL to use when the auth server needs to redirect or link back to the client.
	BaseUrl *string `pulumi:"baseUrl"`
	// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
	ClientId *string `pulumi:"clientId"`
	// Time a client session is allowed to be idle before it expires. Tokens are invalidated when a client session is expired. If not set it uses the standard SSO Session Idle value.
	ClientOfflineSessionIdleTimeout *string `pulumi:"clientOfflineSessionIdleTimeout"`
	// Max time before a client session is expired. Tokens are invalidated when a client session is expired. If not set, it uses the standard SSO Session Max value.
	ClientOfflineSessionMaxLifespan *string `pulumi:"clientOfflineSessionMaxLifespan"`
	// The secret for clients with an `accessType` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
	ClientSecret *string `pulumi:"clientSecret"`
	// Time a client offline session is allowed to be idle before it expires. Offline tokens are invalidated when a client offline session is expired. If not set it uses the Offline Session Idle value.
	ClientSessionIdleTimeout *string `pulumi:"clientSessionIdleTimeout"`
	// Max time before a client offline session is expired. Offline tokens are invalidated when a client offline session is expired. If not set, it uses the Offline Session Max value.
	ClientSessionMaxLifespan *string `pulumi:"clientSessionMaxLifespan"`
	// When `true`, users have to consent to client access.
	ConsentRequired *bool `pulumi:"consentRequired"`
	// The description of this client in the GUI.
	Description *string `pulumi:"description"`
	// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
	DirectAccessGrantsEnabled *bool `pulumi:"directAccessGrantsEnabled"`
	// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// When `true`, the parameter `sessionState` will not be included in OpenID Connect Authentication Response.
	ExcludeSessionStateFromAuthResponse *bool `pulumi:"excludeSessionStateFromAuthResponse"`
	// Allow to include all roles mappings in the access token.
	FullScopeAllowed *bool `pulumi:"fullScopeAllowed"`
	// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
	ImplicitFlowEnabled *bool `pulumi:"implicitFlowEnabled"`
	// The client login theme. This will override the default theme for the realm.
	LoginTheme *string `pulumi:"loginTheme"`
	// The display name of this client in the GUI.
	Name *string `pulumi:"name"`
	// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
	PkceCodeChallengeMethod *string `pulumi:"pkceCodeChallengeMethod"`
	// The realm this client is attached to.
	RealmId *string `pulumi:"realmId"`
	// (Computed) When authorization is enabled for this client, this attribute is the unique ID for the client (the same value as the `.id` attribute).
	ResourceServerId *string `pulumi:"resourceServerId"`
	// When specified, this URL is prepended to any relative URLs found within `validRedirectUris`, `webOrigins`, and `adminUrl`. NOTE: Due to limitations in the Keycloak API, when the `rootUrl` attribute is used, the `validRedirectUris`, `webOrigins`, and `adminUrl` attributes will be required.
	RootUrl *string `pulumi:"rootUrl"`
	// (Computed) When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
	ServiceAccountUserId *string `pulumi:"serviceAccountUserId"`
	// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
	ServiceAccountsEnabled *bool `pulumi:"serviceAccountsEnabled"`
	// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
	StandardFlowEnabled *bool `pulumi:"standardFlowEnabled"`
	// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
	// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standardFlowEnabled` or `implicitFlowEnabled`
	// is set to `true`.
	ValidRedirectUris []string `pulumi:"validRedirectUris"`
	// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
	WebOrigins []string `pulumi:"webOrigins"`
}

type ClientState struct {
	// The amount of time in seconds before an access token expires. This will override the default for the realm.
	AccessTokenLifespan pulumi.StringPtrInput
	// Specifies the type of client, which can be one of the following:
	AccessType pulumi.StringPtrInput
	// URL to the admin interface of the client.
	AdminUrl pulumi.StringPtrInput
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides ClientAuthenticationFlowBindingOverridesPtrInput
	// When this block is present, fine-grained authorization will be enabled for this client. The client's `accessType` must be `CONFIDENTIAL`, and `serviceAccountsEnabled` must be `true`. This block has the following arguments:
	Authorization ClientAuthorizationPtrInput
	// Default URL to use when the auth server needs to redirect or link back to the client.
	BaseUrl pulumi.StringPtrInput
	// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
	ClientId pulumi.StringPtrInput
	// Time a client session is allowed to be idle before it expires. Tokens are invalidated when a client session is expired. If not set it uses the standard SSO Session Idle value.
	ClientOfflineSessionIdleTimeout pulumi.StringPtrInput
	// Max time before a client session is expired. Tokens are invalidated when a client session is expired. If not set, it uses the standard SSO Session Max value.
	ClientOfflineSessionMaxLifespan pulumi.StringPtrInput
	// The secret for clients with an `accessType` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
	ClientSecret pulumi.StringPtrInput
	// Time a client offline session is allowed to be idle before it expires. Offline tokens are invalidated when a client offline session is expired. If not set it uses the Offline Session Idle value.
	ClientSessionIdleTimeout pulumi.StringPtrInput
	// Max time before a client offline session is expired. Offline tokens are invalidated when a client offline session is expired. If not set, it uses the Offline Session Max value.
	ClientSessionMaxLifespan pulumi.StringPtrInput
	// When `true`, users have to consent to client access.
	ConsentRequired pulumi.BoolPtrInput
	// The description of this client in the GUI.
	Description pulumi.StringPtrInput
	// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
	DirectAccessGrantsEnabled pulumi.BoolPtrInput
	// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// When `true`, the parameter `sessionState` will not be included in OpenID Connect Authentication Response.
	ExcludeSessionStateFromAuthResponse pulumi.BoolPtrInput
	// Allow to include all roles mappings in the access token.
	FullScopeAllowed pulumi.BoolPtrInput
	// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
	ImplicitFlowEnabled pulumi.BoolPtrInput
	// The client login theme. This will override the default theme for the realm.
	LoginTheme pulumi.StringPtrInput
	// The display name of this client in the GUI.
	Name pulumi.StringPtrInput
	// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
	PkceCodeChallengeMethod pulumi.StringPtrInput
	// The realm this client is attached to.
	RealmId pulumi.StringPtrInput
	// (Computed) When authorization is enabled for this client, this attribute is the unique ID for the client (the same value as the `.id` attribute).
	ResourceServerId pulumi.StringPtrInput
	// When specified, this URL is prepended to any relative URLs found within `validRedirectUris`, `webOrigins`, and `adminUrl`. NOTE: Due to limitations in the Keycloak API, when the `rootUrl` attribute is used, the `validRedirectUris`, `webOrigins`, and `adminUrl` attributes will be required.
	RootUrl pulumi.StringPtrInput
	// (Computed) When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
	ServiceAccountUserId pulumi.StringPtrInput
	// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
	ServiceAccountsEnabled pulumi.BoolPtrInput
	// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
	StandardFlowEnabled pulumi.BoolPtrInput
	// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
	// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standardFlowEnabled` or `implicitFlowEnabled`
	// is set to `true`.
	ValidRedirectUris pulumi.StringArrayInput
	// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
	WebOrigins pulumi.StringArrayInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// The amount of time in seconds before an access token expires. This will override the default for the realm.
	AccessTokenLifespan *string `pulumi:"accessTokenLifespan"`
	// Specifies the type of client, which can be one of the following:
	AccessType string `pulumi:"accessType"`
	// URL to the admin interface of the client.
	AdminUrl *string `pulumi:"adminUrl"`
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides *ClientAuthenticationFlowBindingOverrides `pulumi:"authenticationFlowBindingOverrides"`
	// When this block is present, fine-grained authorization will be enabled for this client. The client's `accessType` must be `CONFIDENTIAL`, and `serviceAccountsEnabled` must be `true`. This block has the following arguments:
	Authorization *ClientAuthorization `pulumi:"authorization"`
	// Default URL to use when the auth server needs to redirect or link back to the client.
	BaseUrl *string `pulumi:"baseUrl"`
	// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
	ClientId string `pulumi:"clientId"`
	// Time a client session is allowed to be idle before it expires. Tokens are invalidated when a client session is expired. If not set it uses the standard SSO Session Idle value.
	ClientOfflineSessionIdleTimeout *string `pulumi:"clientOfflineSessionIdleTimeout"`
	// Max time before a client session is expired. Tokens are invalidated when a client session is expired. If not set, it uses the standard SSO Session Max value.
	ClientOfflineSessionMaxLifespan *string `pulumi:"clientOfflineSessionMaxLifespan"`
	// The secret for clients with an `accessType` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
	ClientSecret *string `pulumi:"clientSecret"`
	// Time a client offline session is allowed to be idle before it expires. Offline tokens are invalidated when a client offline session is expired. If not set it uses the Offline Session Idle value.
	ClientSessionIdleTimeout *string `pulumi:"clientSessionIdleTimeout"`
	// Max time before a client offline session is expired. Offline tokens are invalidated when a client offline session is expired. If not set, it uses the Offline Session Max value.
	ClientSessionMaxLifespan *string `pulumi:"clientSessionMaxLifespan"`
	// When `true`, users have to consent to client access.
	ConsentRequired *bool `pulumi:"consentRequired"`
	// The description of this client in the GUI.
	Description *string `pulumi:"description"`
	// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
	DirectAccessGrantsEnabled *bool `pulumi:"directAccessGrantsEnabled"`
	// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// When `true`, the parameter `sessionState` will not be included in OpenID Connect Authentication Response.
	ExcludeSessionStateFromAuthResponse *bool `pulumi:"excludeSessionStateFromAuthResponse"`
	// Allow to include all roles mappings in the access token.
	FullScopeAllowed *bool `pulumi:"fullScopeAllowed"`
	// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
	ImplicitFlowEnabled *bool `pulumi:"implicitFlowEnabled"`
	// The client login theme. This will override the default theme for the realm.
	LoginTheme *string `pulumi:"loginTheme"`
	// The display name of this client in the GUI.
	Name *string `pulumi:"name"`
	// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
	PkceCodeChallengeMethod *string `pulumi:"pkceCodeChallengeMethod"`
	// The realm this client is attached to.
	RealmId string `pulumi:"realmId"`
	// When specified, this URL is prepended to any relative URLs found within `validRedirectUris`, `webOrigins`, and `adminUrl`. NOTE: Due to limitations in the Keycloak API, when the `rootUrl` attribute is used, the `validRedirectUris`, `webOrigins`, and `adminUrl` attributes will be required.
	RootUrl *string `pulumi:"rootUrl"`
	// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
	ServiceAccountsEnabled *bool `pulumi:"serviceAccountsEnabled"`
	// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
	StandardFlowEnabled *bool `pulumi:"standardFlowEnabled"`
	// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
	// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standardFlowEnabled` or `implicitFlowEnabled`
	// is set to `true`.
	ValidRedirectUris []string `pulumi:"validRedirectUris"`
	// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
	WebOrigins []string `pulumi:"webOrigins"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// The amount of time in seconds before an access token expires. This will override the default for the realm.
	AccessTokenLifespan pulumi.StringPtrInput
	// Specifies the type of client, which can be one of the following:
	AccessType pulumi.StringInput
	// URL to the admin interface of the client.
	AdminUrl pulumi.StringPtrInput
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides ClientAuthenticationFlowBindingOverridesPtrInput
	// When this block is present, fine-grained authorization will be enabled for this client. The client's `accessType` must be `CONFIDENTIAL`, and `serviceAccountsEnabled` must be `true`. This block has the following arguments:
	Authorization ClientAuthorizationPtrInput
	// Default URL to use when the auth server needs to redirect or link back to the client.
	BaseUrl pulumi.StringPtrInput
	// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
	ClientId pulumi.StringInput
	// Time a client session is allowed to be idle before it expires. Tokens are invalidated when a client session is expired. If not set it uses the standard SSO Session Idle value.
	ClientOfflineSessionIdleTimeout pulumi.StringPtrInput
	// Max time before a client session is expired. Tokens are invalidated when a client session is expired. If not set, it uses the standard SSO Session Max value.
	ClientOfflineSessionMaxLifespan pulumi.StringPtrInput
	// The secret for clients with an `accessType` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
	ClientSecret pulumi.StringPtrInput
	// Time a client offline session is allowed to be idle before it expires. Offline tokens are invalidated when a client offline session is expired. If not set it uses the Offline Session Idle value.
	ClientSessionIdleTimeout pulumi.StringPtrInput
	// Max time before a client offline session is expired. Offline tokens are invalidated when a client offline session is expired. If not set, it uses the Offline Session Max value.
	ClientSessionMaxLifespan pulumi.StringPtrInput
	// When `true`, users have to consent to client access.
	ConsentRequired pulumi.BoolPtrInput
	// The description of this client in the GUI.
	Description pulumi.StringPtrInput
	// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
	DirectAccessGrantsEnabled pulumi.BoolPtrInput
	// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// When `true`, the parameter `sessionState` will not be included in OpenID Connect Authentication Response.
	ExcludeSessionStateFromAuthResponse pulumi.BoolPtrInput
	// Allow to include all roles mappings in the access token.
	FullScopeAllowed pulumi.BoolPtrInput
	// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
	ImplicitFlowEnabled pulumi.BoolPtrInput
	// The client login theme. This will override the default theme for the realm.
	LoginTheme pulumi.StringPtrInput
	// The display name of this client in the GUI.
	Name pulumi.StringPtrInput
	// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
	PkceCodeChallengeMethod pulumi.StringPtrInput
	// The realm this client is attached to.
	RealmId pulumi.StringInput
	// When specified, this URL is prepended to any relative URLs found within `validRedirectUris`, `webOrigins`, and `adminUrl`. NOTE: Due to limitations in the Keycloak API, when the `rootUrl` attribute is used, the `validRedirectUris`, `webOrigins`, and `adminUrl` attributes will be required.
	RootUrl pulumi.StringPtrInput
	// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
	ServiceAccountsEnabled pulumi.BoolPtrInput
	// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
	StandardFlowEnabled pulumi.BoolPtrInput
	// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
	// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standardFlowEnabled` or `implicitFlowEnabled`
	// is set to `true`.
	ValidRedirectUris pulumi.StringArrayInput
	// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
	WebOrigins pulumi.StringArrayInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}

type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(ctx context.Context) ClientOutput
}

func (*Client) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil))
}

func (i *Client) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i *Client) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

type ClientOutput struct {
	*pulumi.OutputState
}

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil))
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClientOutput{})
}
