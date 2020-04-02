// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # .RealmEvents
//
// Allows for managing Realm Events settings within Keycloak.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The name of the realm the event settings apply to.
// - `adminEventsEnabled` - (Optional) When true, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
// - `adminEventsDetailsEnabled` - (Optional) When true, saved admin events will included detailed information for create/update requests. Defaults to `false`.
// - `eventsEnabled` - (Optional) When true, events from `enabledEventTypes` are saved to the database, making them available through the admin console. Defaults to `false`.
// - `eventsExpiration` - (Optional) The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
// - `enabledEventTypes` - (Optional) The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
// - `eventsListeners` - (Optional) The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_realm_events.html.markdown.
type RealmEvents struct {
	pulumi.CustomResourceState

	AdminEventsDetailsEnabled pulumi.BoolPtrOutput     `pulumi:"adminEventsDetailsEnabled"`
	AdminEventsEnabled        pulumi.BoolPtrOutput     `pulumi:"adminEventsEnabled"`
	EnabledEventTypes         pulumi.StringArrayOutput `pulumi:"enabledEventTypes"`
	EventsEnabled             pulumi.BoolPtrOutput     `pulumi:"eventsEnabled"`
	EventsExpiration          pulumi.IntPtrOutput      `pulumi:"eventsExpiration"`
	EventsListeners           pulumi.StringArrayOutput `pulumi:"eventsListeners"`
	RealmId                   pulumi.StringOutput      `pulumi:"realmId"`
}

// NewRealmEvents registers a new resource with the given unique name, arguments, and options.
func NewRealmEvents(ctx *pulumi.Context,
	name string, args *RealmEventsArgs, opts ...pulumi.ResourceOption) (*RealmEvents, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &RealmEventsArgs{}
	}
	var resource RealmEvents
	err := ctx.RegisterResource("keycloak:index/realmEvents:RealmEvents", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmEvents gets an existing RealmEvents resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmEvents(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmEventsState, opts ...pulumi.ResourceOption) (*RealmEvents, error) {
	var resource RealmEvents
	err := ctx.ReadResource("keycloak:index/realmEvents:RealmEvents", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmEvents resources.
type realmEventsState struct {
	AdminEventsDetailsEnabled *bool    `pulumi:"adminEventsDetailsEnabled"`
	AdminEventsEnabled        *bool    `pulumi:"adminEventsEnabled"`
	EnabledEventTypes         []string `pulumi:"enabledEventTypes"`
	EventsEnabled             *bool    `pulumi:"eventsEnabled"`
	EventsExpiration          *int     `pulumi:"eventsExpiration"`
	EventsListeners           []string `pulumi:"eventsListeners"`
	RealmId                   *string  `pulumi:"realmId"`
}

type RealmEventsState struct {
	AdminEventsDetailsEnabled pulumi.BoolPtrInput
	AdminEventsEnabled        pulumi.BoolPtrInput
	EnabledEventTypes         pulumi.StringArrayInput
	EventsEnabled             pulumi.BoolPtrInput
	EventsExpiration          pulumi.IntPtrInput
	EventsListeners           pulumi.StringArrayInput
	RealmId                   pulumi.StringPtrInput
}

func (RealmEventsState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmEventsState)(nil)).Elem()
}

type realmEventsArgs struct {
	AdminEventsDetailsEnabled *bool    `pulumi:"adminEventsDetailsEnabled"`
	AdminEventsEnabled        *bool    `pulumi:"adminEventsEnabled"`
	EnabledEventTypes         []string `pulumi:"enabledEventTypes"`
	EventsEnabled             *bool    `pulumi:"eventsEnabled"`
	EventsExpiration          *int     `pulumi:"eventsExpiration"`
	EventsListeners           []string `pulumi:"eventsListeners"`
	RealmId                   string   `pulumi:"realmId"`
}

// The set of arguments for constructing a RealmEvents resource.
type RealmEventsArgs struct {
	AdminEventsDetailsEnabled pulumi.BoolPtrInput
	AdminEventsEnabled        pulumi.BoolPtrInput
	EnabledEventTypes         pulumi.StringArrayInput
	EventsEnabled             pulumi.BoolPtrInput
	EventsExpiration          pulumi.IntPtrInput
	EventsListeners           pulumi.StringArrayInput
	RealmId                   pulumi.StringInput
}

func (RealmEventsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmEventsArgs)(nil)).Elem()
}