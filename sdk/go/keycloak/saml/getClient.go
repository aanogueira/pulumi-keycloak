// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/saml"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realmManagement, err := saml.LookupClient(ctx, &saml.LookupClientArgs{
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
	err := ctx.Invoke("keycloak:saml/getClient:getClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClient.
type LookupClientArgs struct {
	// The client id (not its unique ID).
	ClientId string `pulumi:"clientId"`
	// The realm id.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	AssertionConsumerPostUrl           string                                       `pulumi:"assertionConsumerPostUrl"`
	AssertionConsumerRedirectUrl       string                                       `pulumi:"assertionConsumerRedirectUrl"`
	AuthenticationFlowBindingOverrides []GetClientAuthenticationFlowBindingOverride `pulumi:"authenticationFlowBindingOverrides"`
	BaseUrl                            string                                       `pulumi:"baseUrl"`
	ClientId                           string                                       `pulumi:"clientId"`
	ClientSignatureRequired            bool                                         `pulumi:"clientSignatureRequired"`
	Description                        string                                       `pulumi:"description"`
	Enabled                            bool                                         `pulumi:"enabled"`
	EncryptAssertions                  bool                                         `pulumi:"encryptAssertions"`
	EncryptionCertificate              string                                       `pulumi:"encryptionCertificate"`
	ForceNameIdFormat                  bool                                         `pulumi:"forceNameIdFormat"`
	ForcePostBinding                   bool                                         `pulumi:"forcePostBinding"`
	FrontChannelLogout                 bool                                         `pulumi:"frontChannelLogout"`
	FullScopeAllowed                   bool                                         `pulumi:"fullScopeAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                              string   `pulumi:"id"`
	IdpInitiatedSsoRelayState       string   `pulumi:"idpInitiatedSsoRelayState"`
	IdpInitiatedSsoUrlName          string   `pulumi:"idpInitiatedSsoUrlName"`
	IncludeAuthnStatement           bool     `pulumi:"includeAuthnStatement"`
	LogoutServicePostBindingUrl     string   `pulumi:"logoutServicePostBindingUrl"`
	LogoutServiceRedirectBindingUrl string   `pulumi:"logoutServiceRedirectBindingUrl"`
	MasterSamlProcessingUrl         string   `pulumi:"masterSamlProcessingUrl"`
	Name                            string   `pulumi:"name"`
	NameIdFormat                    string   `pulumi:"nameIdFormat"`
	RealmId                         string   `pulumi:"realmId"`
	RootUrl                         string   `pulumi:"rootUrl"`
	SignAssertions                  bool     `pulumi:"signAssertions"`
	SignDocuments                   bool     `pulumi:"signDocuments"`
	SignatureAlgorithm              string   `pulumi:"signatureAlgorithm"`
	SigningCertificate              string   `pulumi:"signingCertificate"`
	SigningPrivateKey               string   `pulumi:"signingPrivateKey"`
	ValidRedirectUris               []string `pulumi:"validRedirectUris"`
}
