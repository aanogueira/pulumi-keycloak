// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
//
// Remarks:
//
// - A key must meet all filter criteria
// - This data source may return more than one value.
// - If no key matches the filter criteria, then an error will be returned.
func GetRealmKeys(ctx *pulumi.Context, args *GetRealmKeysArgs, opts ...pulumi.InvokeOption) (*GetRealmKeysResult, error) {
	var rv GetRealmKeysResult
	err := ctx.Invoke("keycloak:index/getRealmKeys:getRealmKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealmKeys.
type GetRealmKeysArgs struct {
	// When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
	Algorithms []string `pulumi:"algorithms"`
	// The realm from which the keys will be retrieved.
	RealmId string `pulumi:"realmId"`
	// When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
	Statuses []string `pulumi:"statuses"`
}

// A collection of values returned by getRealmKeys.
type GetRealmKeysResult struct {
	Algorithms []string `pulumi:"algorithms"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
	Keys    []GetRealmKeysKey `pulumi:"keys"`
	RealmId string            `pulumi:"realmId"`
	// Key status (string)
	Statuses []string `pulumi:"statuses"`
}
