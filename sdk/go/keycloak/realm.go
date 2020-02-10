// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/realm.html.markdown.
type Realm struct {
	s *pulumi.ResourceState
}

// NewRealm registers a new resource with the given unique name, arguments, and options.
func NewRealm(ctx *pulumi.Context,
	name string, args *RealmArgs, opts ...pulumi.ResourceOpt) (*Realm, error) {
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessCodeLifespan"] = nil
		inputs["accessCodeLifespanLogin"] = nil
		inputs["accessCodeLifespanUserAction"] = nil
		inputs["accessTokenLifespan"] = nil
		inputs["accessTokenLifespanForImplicitFlow"] = nil
		inputs["accountTheme"] = nil
		inputs["actionTokenGeneratedByAdminLifespan"] = nil
		inputs["actionTokenGeneratedByUserLifespan"] = nil
		inputs["adminTheme"] = nil
		inputs["attributes"] = nil
		inputs["browserFlow"] = nil
		inputs["clientAuthenticationFlow"] = nil
		inputs["directGrantFlow"] = nil
		inputs["displayName"] = nil
		inputs["dockerAuthenticationFlow"] = nil
		inputs["duplicateEmailsAllowed"] = nil
		inputs["editUsernameAllowed"] = nil
		inputs["emailTheme"] = nil
		inputs["enabled"] = nil
		inputs["internationalization"] = nil
		inputs["loginTheme"] = nil
		inputs["loginWithEmailAllowed"] = nil
		inputs["offlineSessionIdleTimeout"] = nil
		inputs["offlineSessionMaxLifespan"] = nil
		inputs["passwordPolicy"] = nil
		inputs["realm"] = nil
		inputs["refreshTokenMaxReuse"] = nil
		inputs["registrationAllowed"] = nil
		inputs["registrationEmailAsUsername"] = nil
		inputs["registrationFlow"] = nil
		inputs["rememberMe"] = nil
		inputs["resetCredentialsFlow"] = nil
		inputs["resetPasswordAllowed"] = nil
		inputs["revokeRefreshToken"] = nil
		inputs["securityDefenses"] = nil
		inputs["smtpServer"] = nil
		inputs["sslRequired"] = nil
		inputs["ssoSessionIdleTimeout"] = nil
		inputs["ssoSessionMaxLifespan"] = nil
		inputs["verifyEmail"] = nil
	} else {
		inputs["accessCodeLifespan"] = args.AccessCodeLifespan
		inputs["accessCodeLifespanLogin"] = args.AccessCodeLifespanLogin
		inputs["accessCodeLifespanUserAction"] = args.AccessCodeLifespanUserAction
		inputs["accessTokenLifespan"] = args.AccessTokenLifespan
		inputs["accessTokenLifespanForImplicitFlow"] = args.AccessTokenLifespanForImplicitFlow
		inputs["accountTheme"] = args.AccountTheme
		inputs["actionTokenGeneratedByAdminLifespan"] = args.ActionTokenGeneratedByAdminLifespan
		inputs["actionTokenGeneratedByUserLifespan"] = args.ActionTokenGeneratedByUserLifespan
		inputs["adminTheme"] = args.AdminTheme
		inputs["attributes"] = args.Attributes
		inputs["browserFlow"] = args.BrowserFlow
		inputs["clientAuthenticationFlow"] = args.ClientAuthenticationFlow
		inputs["directGrantFlow"] = args.DirectGrantFlow
		inputs["displayName"] = args.DisplayName
		inputs["dockerAuthenticationFlow"] = args.DockerAuthenticationFlow
		inputs["duplicateEmailsAllowed"] = args.DuplicateEmailsAllowed
		inputs["editUsernameAllowed"] = args.EditUsernameAllowed
		inputs["emailTheme"] = args.EmailTheme
		inputs["enabled"] = args.Enabled
		inputs["internationalization"] = args.Internationalization
		inputs["loginTheme"] = args.LoginTheme
		inputs["loginWithEmailAllowed"] = args.LoginWithEmailAllowed
		inputs["offlineSessionIdleTimeout"] = args.OfflineSessionIdleTimeout
		inputs["offlineSessionMaxLifespan"] = args.OfflineSessionMaxLifespan
		inputs["passwordPolicy"] = args.PasswordPolicy
		inputs["realm"] = args.Realm
		inputs["refreshTokenMaxReuse"] = args.RefreshTokenMaxReuse
		inputs["registrationAllowed"] = args.RegistrationAllowed
		inputs["registrationEmailAsUsername"] = args.RegistrationEmailAsUsername
		inputs["registrationFlow"] = args.RegistrationFlow
		inputs["rememberMe"] = args.RememberMe
		inputs["resetCredentialsFlow"] = args.ResetCredentialsFlow
		inputs["resetPasswordAllowed"] = args.ResetPasswordAllowed
		inputs["revokeRefreshToken"] = args.RevokeRefreshToken
		inputs["securityDefenses"] = args.SecurityDefenses
		inputs["smtpServer"] = args.SmtpServer
		inputs["sslRequired"] = args.SslRequired
		inputs["ssoSessionIdleTimeout"] = args.SsoSessionIdleTimeout
		inputs["ssoSessionMaxLifespan"] = args.SsoSessionMaxLifespan
		inputs["verifyEmail"] = args.VerifyEmail
	}
	s, err := ctx.RegisterResource("keycloak:index/realm:Realm", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Realm{s: s}, nil
}

// GetRealm gets an existing Realm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealm(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RealmState, opts ...pulumi.ResourceOpt) (*Realm, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessCodeLifespan"] = state.AccessCodeLifespan
		inputs["accessCodeLifespanLogin"] = state.AccessCodeLifespanLogin
		inputs["accessCodeLifespanUserAction"] = state.AccessCodeLifespanUserAction
		inputs["accessTokenLifespan"] = state.AccessTokenLifespan
		inputs["accessTokenLifespanForImplicitFlow"] = state.AccessTokenLifespanForImplicitFlow
		inputs["accountTheme"] = state.AccountTheme
		inputs["actionTokenGeneratedByAdminLifespan"] = state.ActionTokenGeneratedByAdminLifespan
		inputs["actionTokenGeneratedByUserLifespan"] = state.ActionTokenGeneratedByUserLifespan
		inputs["adminTheme"] = state.AdminTheme
		inputs["attributes"] = state.Attributes
		inputs["browserFlow"] = state.BrowserFlow
		inputs["clientAuthenticationFlow"] = state.ClientAuthenticationFlow
		inputs["directGrantFlow"] = state.DirectGrantFlow
		inputs["displayName"] = state.DisplayName
		inputs["dockerAuthenticationFlow"] = state.DockerAuthenticationFlow
		inputs["duplicateEmailsAllowed"] = state.DuplicateEmailsAllowed
		inputs["editUsernameAllowed"] = state.EditUsernameAllowed
		inputs["emailTheme"] = state.EmailTheme
		inputs["enabled"] = state.Enabled
		inputs["internationalization"] = state.Internationalization
		inputs["loginTheme"] = state.LoginTheme
		inputs["loginWithEmailAllowed"] = state.LoginWithEmailAllowed
		inputs["offlineSessionIdleTimeout"] = state.OfflineSessionIdleTimeout
		inputs["offlineSessionMaxLifespan"] = state.OfflineSessionMaxLifespan
		inputs["passwordPolicy"] = state.PasswordPolicy
		inputs["realm"] = state.Realm
		inputs["refreshTokenMaxReuse"] = state.RefreshTokenMaxReuse
		inputs["registrationAllowed"] = state.RegistrationAllowed
		inputs["registrationEmailAsUsername"] = state.RegistrationEmailAsUsername
		inputs["registrationFlow"] = state.RegistrationFlow
		inputs["rememberMe"] = state.RememberMe
		inputs["resetCredentialsFlow"] = state.ResetCredentialsFlow
		inputs["resetPasswordAllowed"] = state.ResetPasswordAllowed
		inputs["revokeRefreshToken"] = state.RevokeRefreshToken
		inputs["securityDefenses"] = state.SecurityDefenses
		inputs["smtpServer"] = state.SmtpServer
		inputs["sslRequired"] = state.SslRequired
		inputs["ssoSessionIdleTimeout"] = state.SsoSessionIdleTimeout
		inputs["ssoSessionMaxLifespan"] = state.SsoSessionMaxLifespan
		inputs["verifyEmail"] = state.VerifyEmail
	}
	s, err := ctx.ReadResource("keycloak:index/realm:Realm", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Realm{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Realm) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Realm) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Realm) AccessCodeLifespan() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessCodeLifespan"])
}

func (r *Realm) AccessCodeLifespanLogin() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessCodeLifespanLogin"])
}

func (r *Realm) AccessCodeLifespanUserAction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessCodeLifespanUserAction"])
}

func (r *Realm) AccessTokenLifespan() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessTokenLifespan"])
}

func (r *Realm) AccessTokenLifespanForImplicitFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessTokenLifespanForImplicitFlow"])
}

func (r *Realm) AccountTheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accountTheme"])
}

func (r *Realm) ActionTokenGeneratedByAdminLifespan() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["actionTokenGeneratedByAdminLifespan"])
}

func (r *Realm) ActionTokenGeneratedByUserLifespan() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["actionTokenGeneratedByUserLifespan"])
}

func (r *Realm) AdminTheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["adminTheme"])
}

func (r *Realm) Attributes() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["attributes"])
}

// Which flow should be used for BrowserFlow
func (r *Realm) BrowserFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["browserFlow"])
}

// Which flow should be used for ClientAuthenticationFlow
func (r *Realm) ClientAuthenticationFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientAuthenticationFlow"])
}

// Which flow should be used for DirectGrantFlow
func (r *Realm) DirectGrantFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["directGrantFlow"])
}

func (r *Realm) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// Which flow should be used for DockerAuthenticationFlow
func (r *Realm) DockerAuthenticationFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dockerAuthenticationFlow"])
}

func (r *Realm) DuplicateEmailsAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["duplicateEmailsAllowed"])
}

func (r *Realm) EditUsernameAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["editUsernameAllowed"])
}

func (r *Realm) EmailTheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["emailTheme"])
}

func (r *Realm) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

func (r *Realm) Internationalization() pulumi.Output {
	return r.s.State["internationalization"]
}

func (r *Realm) LoginTheme() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loginTheme"])
}

func (r *Realm) LoginWithEmailAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["loginWithEmailAllowed"])
}

func (r *Realm) OfflineSessionIdleTimeout() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["offlineSessionIdleTimeout"])
}

func (r *Realm) OfflineSessionMaxLifespan() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["offlineSessionMaxLifespan"])
}

// String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported policies
// can be found in the server-info providers page. example: "upperCase(1) and length(8) and forceExpiredPasswordChange(365)
// and notUsername(undefined)"
func (r *Realm) PasswordPolicy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["passwordPolicy"])
}

func (r *Realm) Realm() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realm"])
}

func (r *Realm) RefreshTokenMaxReuse() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["refreshTokenMaxReuse"])
}

func (r *Realm) RegistrationAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["registrationAllowed"])
}

func (r *Realm) RegistrationEmailAsUsername() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["registrationEmailAsUsername"])
}

// Which flow should be used for RegistrationFlow
func (r *Realm) RegistrationFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["registrationFlow"])
}

func (r *Realm) RememberMe() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["rememberMe"])
}

// Which flow should be used for ResetCredentialsFlow
func (r *Realm) ResetCredentialsFlow() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resetCredentialsFlow"])
}

func (r *Realm) ResetPasswordAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["resetPasswordAllowed"])
}

func (r *Realm) RevokeRefreshToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["revokeRefreshToken"])
}

func (r *Realm) SecurityDefenses() pulumi.Output {
	return r.s.State["securityDefenses"]
}

func (r *Realm) SmtpServer() pulumi.Output {
	return r.s.State["smtpServer"]
}

// SSL Required: Values can be 'none', 'external' or 'all'.
func (r *Realm) SslRequired() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sslRequired"])
}

func (r *Realm) SsoSessionIdleTimeout() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ssoSessionIdleTimeout"])
}

func (r *Realm) SsoSessionMaxLifespan() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ssoSessionMaxLifespan"])
}

func (r *Realm) VerifyEmail() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["verifyEmail"])
}

// Input properties used for looking up and filtering Realm resources.
type RealmState struct {
	AccessCodeLifespan interface{}
	AccessCodeLifespanLogin interface{}
	AccessCodeLifespanUserAction interface{}
	AccessTokenLifespan interface{}
	AccessTokenLifespanForImplicitFlow interface{}
	AccountTheme interface{}
	ActionTokenGeneratedByAdminLifespan interface{}
	ActionTokenGeneratedByUserLifespan interface{}
	AdminTheme interface{}
	Attributes interface{}
	// Which flow should be used for BrowserFlow
	BrowserFlow interface{}
	// Which flow should be used for ClientAuthenticationFlow
	ClientAuthenticationFlow interface{}
	// Which flow should be used for DirectGrantFlow
	DirectGrantFlow interface{}
	DisplayName interface{}
	// Which flow should be used for DockerAuthenticationFlow
	DockerAuthenticationFlow interface{}
	DuplicateEmailsAllowed interface{}
	EditUsernameAllowed interface{}
	EmailTheme interface{}
	Enabled interface{}
	Internationalization interface{}
	LoginTheme interface{}
	LoginWithEmailAllowed interface{}
	OfflineSessionIdleTimeout interface{}
	OfflineSessionMaxLifespan interface{}
	// String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported
	// policies can be found in the server-info providers page. example: "upperCase(1) and length(8) and
	// forceExpiredPasswordChange(365) and notUsername(undefined)"
	PasswordPolicy interface{}
	Realm interface{}
	RefreshTokenMaxReuse interface{}
	RegistrationAllowed interface{}
	RegistrationEmailAsUsername interface{}
	// Which flow should be used for RegistrationFlow
	RegistrationFlow interface{}
	RememberMe interface{}
	// Which flow should be used for ResetCredentialsFlow
	ResetCredentialsFlow interface{}
	ResetPasswordAllowed interface{}
	RevokeRefreshToken interface{}
	SecurityDefenses interface{}
	SmtpServer interface{}
	// SSL Required: Values can be 'none', 'external' or 'all'.
	SslRequired interface{}
	SsoSessionIdleTimeout interface{}
	SsoSessionMaxLifespan interface{}
	VerifyEmail interface{}
}

// The set of arguments for constructing a Realm resource.
type RealmArgs struct {
	AccessCodeLifespan interface{}
	AccessCodeLifespanLogin interface{}
	AccessCodeLifespanUserAction interface{}
	AccessTokenLifespan interface{}
	AccessTokenLifespanForImplicitFlow interface{}
	AccountTheme interface{}
	ActionTokenGeneratedByAdminLifespan interface{}
	ActionTokenGeneratedByUserLifespan interface{}
	AdminTheme interface{}
	Attributes interface{}
	// Which flow should be used for BrowserFlow
	BrowserFlow interface{}
	// Which flow should be used for ClientAuthenticationFlow
	ClientAuthenticationFlow interface{}
	// Which flow should be used for DirectGrantFlow
	DirectGrantFlow interface{}
	DisplayName interface{}
	// Which flow should be used for DockerAuthenticationFlow
	DockerAuthenticationFlow interface{}
	DuplicateEmailsAllowed interface{}
	EditUsernameAllowed interface{}
	EmailTheme interface{}
	Enabled interface{}
	Internationalization interface{}
	LoginTheme interface{}
	LoginWithEmailAllowed interface{}
	OfflineSessionIdleTimeout interface{}
	OfflineSessionMaxLifespan interface{}
	// String that represents the passwordPolicies that are in place. Each policy is separated with " and ". Supported
	// policies can be found in the server-info providers page. example: "upperCase(1) and length(8) and
	// forceExpiredPasswordChange(365) and notUsername(undefined)"
	PasswordPolicy interface{}
	Realm interface{}
	RefreshTokenMaxReuse interface{}
	RegistrationAllowed interface{}
	RegistrationEmailAsUsername interface{}
	// Which flow should be used for RegistrationFlow
	RegistrationFlow interface{}
	RememberMe interface{}
	// Which flow should be used for ResetCredentialsFlow
	ResetCredentialsFlow interface{}
	ResetPasswordAllowed interface{}
	RevokeRefreshToken interface{}
	SecurityDefenses interface{}
	SmtpServer interface{}
	// SSL Required: Values can be 'none', 'external' or 'all'.
	SslRequired interface{}
	SsoSessionIdleTimeout interface{}
	SsoSessionMaxLifespan interface{}
	VerifyEmail interface{}
}
