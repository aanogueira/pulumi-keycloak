// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # openid.Client data source
//
// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm id.
// - `clientId` - (Required) The client id.
//
// ### Attributes Reference
//
// See the docs for the `openid.Client` resource for details on the exported attributes.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/d/keycloak_openid_client.html.markdown.
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
	ClientId string `pulumi:"clientId"`
	RealmId  string `pulumi:"realmId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	AccessType                string                 `pulumi:"accessType"`
	Authorization             GetClientAuthorization `pulumi:"authorization"`
	ClientId                  string                 `pulumi:"clientId"`
	ClientSecret              string                 `pulumi:"clientSecret"`
	Description               string                 `pulumi:"description"`
	DirectAccessGrantsEnabled bool                   `pulumi:"directAccessGrantsEnabled"`
	Enabled                   bool                   `pulumi:"enabled"`
	FullScopeAllowed          bool                   `pulumi:"fullScopeAllowed"`
	// id is the provider-assigned unique ID for this managed resource.
	Id                     string   `pulumi:"id"`
	ImplicitFlowEnabled    bool     `pulumi:"implicitFlowEnabled"`
	Name                   string   `pulumi:"name"`
	RealmId                string   `pulumi:"realmId"`
	ResourceServerId       string   `pulumi:"resourceServerId"`
	ServiceAccountUserId   string   `pulumi:"serviceAccountUserId"`
	ServiceAccountsEnabled bool     `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled    bool     `pulumi:"standardFlowEnabled"`
	ValidRedirectUris      []string `pulumi:"validRedirectUris"`
	WebOrigins             []string `pulumi:"webOrigins"`
}
