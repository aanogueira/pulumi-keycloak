// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # .GenericClientProtocolMapper
//
// Allows for creating and managing protocol mapper for both types of clients (openid-connect and saml) within Keycloak.
//
// There are two uses cases for using this resource:
// * If you implemented a custom protocol mapper, this resource can be used to configure it
// * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
//
// Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
// Therefore, if possible, a specific mapper should be used.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required) The client this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `protocol` - (Required) The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
// - `protocolMapper` - (Required) The name of the protocol mapper. The protocol mapper must be
//    compatible with the specified client.
// - `config` - (Required) A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_generic_client_protocol_mapper.html.markdown.
type GenericClientProtocolMapper struct {
	pulumi.CustomResourceState

	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	Config        pulumi.MapOutput       `pulumi:"config"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol of the client (openid-connect / saml).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The type of the protocol mapper.
	ProtocolMapper pulumi.StringOutput `pulumi:"protocolMapper"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGenericClientProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewGenericClientProtocolMapper(ctx *pulumi.Context,
	name string, args *GenericClientProtocolMapperArgs, opts ...pulumi.ResourceOption) (*GenericClientProtocolMapper, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ProtocolMapper == nil {
		return nil, errors.New("missing required argument 'ProtocolMapper'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GenericClientProtocolMapperArgs{}
	}
	var resource GenericClientProtocolMapper
	err := ctx.RegisterResource("keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenericClientProtocolMapper gets an existing GenericClientProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenericClientProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GenericClientProtocolMapperState, opts ...pulumi.ResourceOption) (*GenericClientProtocolMapper, error) {
	var resource GenericClientProtocolMapper
	err := ctx.ReadResource("keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GenericClientProtocolMapper resources.
type genericClientProtocolMapperState struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string                `pulumi:"clientScopeId"`
	Config        map[string]interface{} `pulumi:"config"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The protocol of the client (openid-connect / saml).
	Protocol *string `pulumi:"protocol"`
	// The type of the protocol mapper.
	ProtocolMapper *string `pulumi:"protocolMapper"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
}

type GenericClientProtocolMapperState struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	Config        pulumi.MapInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The protocol of the client (openid-connect / saml).
	Protocol pulumi.StringPtrInput
	// The type of the protocol mapper.
	ProtocolMapper pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
}

func (GenericClientProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*genericClientProtocolMapperState)(nil)).Elem()
}

type genericClientProtocolMapperArgs struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string                `pulumi:"clientScopeId"`
	Config        map[string]interface{} `pulumi:"config"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The protocol of the client (openid-connect / saml).
	Protocol string `pulumi:"protocol"`
	// The type of the protocol mapper.
	ProtocolMapper string `pulumi:"protocolMapper"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GenericClientProtocolMapper resource.
type GenericClientProtocolMapperArgs struct {
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	Config        pulumi.MapInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The protocol of the client (openid-connect / saml).
	Protocol pulumi.StringInput
	// The type of the protocol mapper.
	ProtocolMapper pulumi.StringInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
}

func (GenericClientProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*genericClientProtocolMapperArgs)(nil)).Elem()
}
