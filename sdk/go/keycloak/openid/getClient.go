// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realmManagement, err := openid.LookupClient(ctx, &openid.LookupClientArgs{
// 			RealmId:  "my-realm",
// 			ClientId: "realm-management",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := realmManagement.Id
// 		_, err = keycloak.LookupRole(ctx, &keycloak.LookupRoleArgs{
// 			RealmId:  "my-realm",
// 			ClientId: &opt0,
// 			Name:     "realm-admin",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupClient(ctx *pulumi.Context, args *LookupClientArgs, opts ...pulumi.InvokeOption) (*LookupClientResult, error) {
	var rv LookupClientResult
	err := ctx.Invoke("keycloak:openid/getClient:getClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClient.
type LookupClientArgs struct {
	// The client id (not its unique ID).
	ClientId    string                 `pulumi:"clientId"`
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The realm id.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	AccessTokenLifespan                    string                                       `pulumi:"accessTokenLifespan"`
	AccessType                             string                                       `pulumi:"accessType"`
	AdminUrl                               string                                       `pulumi:"adminUrl"`
	AuthenticationFlowBindingOverrides     []GetClientAuthenticationFlowBindingOverride `pulumi:"authenticationFlowBindingOverrides"`
	Authorizations                         []GetClientAuthorization                     `pulumi:"authorizations"`
	BackchannelLogoutRevokeOfflineSessions bool                                         `pulumi:"backchannelLogoutRevokeOfflineSessions"`
	BackchannelLogoutSessionRequired       bool                                         `pulumi:"backchannelLogoutSessionRequired"`
	BackchannelLogoutUrl                   string                                       `pulumi:"backchannelLogoutUrl"`
	BaseUrl                                string                                       `pulumi:"baseUrl"`
	ClientId                               string                                       `pulumi:"clientId"`
	ClientOfflineSessionIdleTimeout        string                                       `pulumi:"clientOfflineSessionIdleTimeout"`
	ClientOfflineSessionMaxLifespan        string                                       `pulumi:"clientOfflineSessionMaxLifespan"`
	ClientSecret                           string                                       `pulumi:"clientSecret"`
	ClientSessionIdleTimeout               string                                       `pulumi:"clientSessionIdleTimeout"`
	ClientSessionMaxLifespan               string                                       `pulumi:"clientSessionMaxLifespan"`
	ConsentRequired                        bool                                         `pulumi:"consentRequired"`
	Description                            string                                       `pulumi:"description"`
	DirectAccessGrantsEnabled              bool                                         `pulumi:"directAccessGrantsEnabled"`
	Enabled                                bool                                         `pulumi:"enabled"`
	ExcludeSessionStateFromAuthResponse    bool                                         `pulumi:"excludeSessionStateFromAuthResponse"`
	ExtraConfig                            map[string]interface{}                       `pulumi:"extraConfig"`
	FullScopeAllowed                       bool                                         `pulumi:"fullScopeAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string   `pulumi:"id"`
	ImplicitFlowEnabled     bool     `pulumi:"implicitFlowEnabled"`
	LoginTheme              string   `pulumi:"loginTheme"`
	Name                    string   `pulumi:"name"`
	PkceCodeChallengeMethod string   `pulumi:"pkceCodeChallengeMethod"`
	RealmId                 string   `pulumi:"realmId"`
	ResourceServerId        string   `pulumi:"resourceServerId"`
	RootUrl                 string   `pulumi:"rootUrl"`
	ServiceAccountUserId    string   `pulumi:"serviceAccountUserId"`
	ServiceAccountsEnabled  bool     `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled     bool     `pulumi:"standardFlowEnabled"`
	UseRefreshTokens        bool     `pulumi:"useRefreshTokens"`
	ValidRedirectUris       []string `pulumi:"validRedirectUris"`
	WebOrigins              []string `pulumi:"webOrigins"`
}
