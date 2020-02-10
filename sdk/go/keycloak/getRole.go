// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # .Role data source
// 
// This data source can be used to fetch properties of a Keycloak role for
// usage with other resources, such as `.GroupRoles`.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this role exists within.
// - `clientId` - (Optional) When specified, this role is assumed to be a
//   client role belonging to the client with the provided ID
// - `name` - (Required) The name of the role
//   
// ### Attributes Reference
// 
// In addition to the arguments listed above, the following computed attributes are exported:
// 
// - `id` - The unique ID of the role, which can be used as an argument to
//   other resources supported by this provider.
// - `description` - The description of the role.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/d/role.html.markdown.
func LookupRole(ctx *pulumi.Context, args *GetRoleArgs) (*GetRoleResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["clientId"] = args.ClientId
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
	}
	outputs, err := ctx.Invoke("keycloak:index/getRole:getRole", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRoleResult{
		ClientId: outputs["clientId"],
		Description: outputs["description"],
		Name: outputs["name"],
		RealmId: outputs["realmId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRole.
type GetRoleArgs struct {
	ClientId interface{}
	Name interface{}
	RealmId interface{}
}

// A collection of values returned by getRole.
type GetRoleResult struct {
	ClientId interface{}
	Description interface{}
	Name interface{}
	RealmId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
