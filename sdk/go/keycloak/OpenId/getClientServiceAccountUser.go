// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func LookupClientServiceAccountUser(ctx *pulumi.Context, args *GetClientServiceAccountUserArgs) (*GetClientServiceAccountUserResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["clientId"] = args.ClientId
		inputs["realmId"] = args.RealmId
	}
	outputs, err := ctx.Invoke("keycloak:OpenId/getClientServiceAccountUser:getClientServiceAccountUser", inputs)
	if err != nil {
		return nil, err
	}
	return &GetClientServiceAccountUserResult{
		Attributes: outputs["attributes"],
		ClientId: outputs["clientId"],
		Email: outputs["email"],
		Enabled: outputs["enabled"],
		FederatedIdentities: outputs["federatedIdentities"],
		FirstName: outputs["firstName"],
		LastName: outputs["lastName"],
		RealmId: outputs["realmId"],
		Username: outputs["username"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getClientServiceAccountUser.
type GetClientServiceAccountUserArgs struct {
	ClientId interface{}
	RealmId interface{}
}

// A collection of values returned by getClientServiceAccountUser.
type GetClientServiceAccountUserResult struct {
	Attributes interface{}
	ClientId interface{}
	Email interface{}
	Enabled interface{}
	FederatedIdentities interface{}
	FirstName interface{}
	LastName interface{}
	RealmId interface{}
	Username interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}