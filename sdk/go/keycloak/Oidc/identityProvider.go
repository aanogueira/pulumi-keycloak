// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package Oidc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type IdentityProvider struct {
	s *pulumi.ResourceState
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOpt) (*IdentityProvider, error) {
	if args == nil || args.Alias == nil {
		return nil, errors.New("missing required argument 'Alias'")
	}
	if args == nil || args.AuthorizationUrl == nil {
		return nil, errors.New("missing required argument 'AuthorizationUrl'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	if args == nil || args.TokenUrl == nil {
		return nil, errors.New("missing required argument 'TokenUrl'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addReadTokenRoleOnCreate"] = nil
		inputs["alias"] = nil
		inputs["authenticateByDefault"] = nil
		inputs["authorizationUrl"] = nil
		inputs["backchannelSupported"] = nil
		inputs["clientId"] = nil
		inputs["clientSecret"] = nil
		inputs["displayName"] = nil
		inputs["enabled"] = nil
		inputs["extraConfig"] = nil
		inputs["firstBrokerLoginFlowAlias"] = nil
		inputs["hideOnLoginPage"] = nil
		inputs["jwksUrl"] = nil
		inputs["linkOnly"] = nil
		inputs["loginHint"] = nil
		inputs["logoutUrl"] = nil
		inputs["postBrokerLoginFlowAlias"] = nil
		inputs["providerId"] = nil
		inputs["realm"] = nil
		inputs["storeToken"] = nil
		inputs["tokenUrl"] = nil
		inputs["trustEmail"] = nil
		inputs["uiLocales"] = nil
		inputs["userInfoUrl"] = nil
		inputs["validateSignature"] = nil
	} else {
		inputs["addReadTokenRoleOnCreate"] = args.AddReadTokenRoleOnCreate
		inputs["alias"] = args.Alias
		inputs["authenticateByDefault"] = args.AuthenticateByDefault
		inputs["authorizationUrl"] = args.AuthorizationUrl
		inputs["backchannelSupported"] = args.BackchannelSupported
		inputs["clientId"] = args.ClientId
		inputs["clientSecret"] = args.ClientSecret
		inputs["displayName"] = args.DisplayName
		inputs["enabled"] = args.Enabled
		inputs["extraConfig"] = args.ExtraConfig
		inputs["firstBrokerLoginFlowAlias"] = args.FirstBrokerLoginFlowAlias
		inputs["hideOnLoginPage"] = args.HideOnLoginPage
		inputs["jwksUrl"] = args.JwksUrl
		inputs["linkOnly"] = args.LinkOnly
		inputs["loginHint"] = args.LoginHint
		inputs["logoutUrl"] = args.LogoutUrl
		inputs["postBrokerLoginFlowAlias"] = args.PostBrokerLoginFlowAlias
		inputs["providerId"] = args.ProviderId
		inputs["realm"] = args.Realm
		inputs["storeToken"] = args.StoreToken
		inputs["tokenUrl"] = args.TokenUrl
		inputs["trustEmail"] = args.TrustEmail
		inputs["uiLocales"] = args.UiLocales
		inputs["userInfoUrl"] = args.UserInfoUrl
		inputs["validateSignature"] = args.ValidateSignature
	}
	inputs["internalId"] = nil
	s, err := ctx.RegisterResource("keycloak:Oidc/identityProvider:IdentityProvider", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IdentityProvider{s: s}, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IdentityProviderState, opts ...pulumi.ResourceOpt) (*IdentityProvider, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addReadTokenRoleOnCreate"] = state.AddReadTokenRoleOnCreate
		inputs["alias"] = state.Alias
		inputs["authenticateByDefault"] = state.AuthenticateByDefault
		inputs["authorizationUrl"] = state.AuthorizationUrl
		inputs["backchannelSupported"] = state.BackchannelSupported
		inputs["clientId"] = state.ClientId
		inputs["clientSecret"] = state.ClientSecret
		inputs["displayName"] = state.DisplayName
		inputs["enabled"] = state.Enabled
		inputs["extraConfig"] = state.ExtraConfig
		inputs["firstBrokerLoginFlowAlias"] = state.FirstBrokerLoginFlowAlias
		inputs["hideOnLoginPage"] = state.HideOnLoginPage
		inputs["internalId"] = state.InternalId
		inputs["jwksUrl"] = state.JwksUrl
		inputs["linkOnly"] = state.LinkOnly
		inputs["loginHint"] = state.LoginHint
		inputs["logoutUrl"] = state.LogoutUrl
		inputs["postBrokerLoginFlowAlias"] = state.PostBrokerLoginFlowAlias
		inputs["providerId"] = state.ProviderId
		inputs["realm"] = state.Realm
		inputs["storeToken"] = state.StoreToken
		inputs["tokenUrl"] = state.TokenUrl
		inputs["trustEmail"] = state.TrustEmail
		inputs["uiLocales"] = state.UiLocales
		inputs["userInfoUrl"] = state.UserInfoUrl
		inputs["validateSignature"] = state.ValidateSignature
	}
	s, err := ctx.ReadResource("keycloak:Oidc/identityProvider:IdentityProvider", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IdentityProvider{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IdentityProvider) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IdentityProvider) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
func (r *IdentityProvider) AddReadTokenRoleOnCreate() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addReadTokenRoleOnCreate"])
}

// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
func (r *IdentityProvider) Alias() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["alias"])
}

// Enable/disable authenticate users by default.
func (r *IdentityProvider) AuthenticateByDefault() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["authenticateByDefault"])
}

// OIDC authorization URL.
func (r *IdentityProvider) AuthorizationUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["authorizationUrl"])
}

// Does the external IDP support backchannel logout?
func (r *IdentityProvider) BackchannelSupported() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["backchannelSupported"])
}

// Client ID.
func (r *IdentityProvider) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// Client Secret.
func (r *IdentityProvider) ClientSecret() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientSecret"])
}

// Friendly name for Identity Providers.
func (r *IdentityProvider) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// Enable/disable this identity provider.
func (r *IdentityProvider) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

func (r *IdentityProvider) ExtraConfig() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["extraConfig"])
}

// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
func (r *IdentityProvider) FirstBrokerLoginFlowAlias() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["firstBrokerLoginFlowAlias"])
}

// Hide On Login Page.
func (r *IdentityProvider) HideOnLoginPage() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["hideOnLoginPage"])
}

// Internal Identity Provider Id
func (r *IdentityProvider) InternalId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internalId"])
}

// JSON Web Key Set URL
func (r *IdentityProvider) JwksUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["jwksUrl"])
}

// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
// want to allow login from the provider, but want to integrate with a provider
func (r *IdentityProvider) LinkOnly() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["linkOnly"])
}

// Login Hint.
func (r *IdentityProvider) LoginHint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loginHint"])
}

// Logout URL
func (r *IdentityProvider) LogoutUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["logoutUrl"])
}

// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
func (r *IdentityProvider) PostBrokerLoginFlowAlias() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["postBrokerLoginFlowAlias"])
}

// provider id, is always oidc, unless you have a custom implementation
func (r *IdentityProvider) ProviderId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["providerId"])
}

// Realm Name
func (r *IdentityProvider) Realm() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realm"])
}

// Enable/disable if tokens must be stored after authenticating users.
func (r *IdentityProvider) StoreToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["storeToken"])
}

// Token URL.
func (r *IdentityProvider) TokenUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tokenUrl"])
}

// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
func (r *IdentityProvider) TrustEmail() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["trustEmail"])
}

// Pass current locale to identity provider
func (r *IdentityProvider) UiLocales() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["uiLocales"])
}

// User Info URL
func (r *IdentityProvider) UserInfoUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userInfoUrl"])
}

// Enable/disable signature validation of SAML responses.
func (r *IdentityProvider) ValidateSignature() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["validateSignature"])
}

// Input properties used for looking up and filtering IdentityProvider resources.
type IdentityProviderState struct {
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate interface{}
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias interface{}
	// Enable/disable authenticate users by default.
	AuthenticateByDefault interface{}
	// OIDC authorization URL.
	AuthorizationUrl interface{}
	// Does the external IDP support backchannel logout?
	BackchannelSupported interface{}
	// Client ID.
	ClientId interface{}
	// Client Secret.
	ClientSecret interface{}
	// Friendly name for Identity Providers.
	DisplayName interface{}
	// Enable/disable this identity provider.
	Enabled interface{}
	ExtraConfig interface{}
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login'
	// means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias interface{}
	// Hide On Login Page.
	HideOnLoginPage interface{}
	// Internal Identity Provider Id
	InternalId interface{}
	// JSON Web Key Set URL
	JwksUrl interface{}
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly interface{}
	// Login Hint.
	LoginHint interface{}
	// Logout URL
	LogoutUrl interface{}
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set
	// it.
	PostBrokerLoginFlowAlias interface{}
	// provider id, is always oidc, unless you have a custom implementation
	ProviderId interface{}
	// Realm Name
	Realm interface{}
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken interface{}
	// Token URL.
	TokenUrl interface{}
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail interface{}
	// Pass current locale to identity provider
	UiLocales interface{}
	// User Info URL
	UserInfoUrl interface{}
	// Enable/disable signature validation of SAML responses.
	ValidateSignature interface{}
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate interface{}
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias interface{}
	// Enable/disable authenticate users by default.
	AuthenticateByDefault interface{}
	// OIDC authorization URL.
	AuthorizationUrl interface{}
	// Does the external IDP support backchannel logout?
	BackchannelSupported interface{}
	// Client ID.
	ClientId interface{}
	// Client Secret.
	ClientSecret interface{}
	// Friendly name for Identity Providers.
	DisplayName interface{}
	// Enable/disable this identity provider.
	Enabled interface{}
	ExtraConfig interface{}
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login'
	// means that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias interface{}
	// Hide On Login Page.
	HideOnLoginPage interface{}
	// JSON Web Key Set URL
	JwksUrl interface{}
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly interface{}
	// Login Hint.
	LoginHint interface{}
	// Logout URL
	LogoutUrl interface{}
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set
	// it.
	PostBrokerLoginFlowAlias interface{}
	// provider id, is always oidc, unless you have a custom implementation
	ProviderId interface{}
	// Realm Name
	Realm interface{}
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken interface{}
	// Token URL.
	TokenUrl interface{}
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail interface{}
	// Pass current locale to identity provider
	UiLocales interface{}
	// User Info URL
	UserInfoUrl interface{}
	// Enable/disable signature validation of SAML responses.
	ValidateSignature interface{}
}