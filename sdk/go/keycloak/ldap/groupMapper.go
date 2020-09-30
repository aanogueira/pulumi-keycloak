// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing group mappers for Keycloak users federated via LDAP.
//
// The LDAP group mapper can be used to map an LDAP user's groups from some DN to Keycloak groups. This group mapper will also
// create the groups within Keycloak if they do not already exist.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/ldap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ldapUserFederation, err := ldap.NewUserFederation(ctx, "ldapUserFederation", &ldap.UserFederationArgs{
// 			RealmId:               realm.ID(),
// 			UsernameLdapAttribute: pulumi.String("cn"),
// 			RdnLdapAttribute:      pulumi.String("cn"),
// 			UuidLdapAttribute:     pulumi.String("entryDN"),
// 			UserObjectClasses: pulumi.StringArray{
// 				pulumi.String("simpleSecurityObject"),
// 				pulumi.String("organizationalRole"),
// 			},
// 			ConnectionUrl:  pulumi.String("ldap://openldap"),
// 			UsersDn:        pulumi.String("dc=example,dc=org"),
// 			BindDn:         pulumi.String("cn=admin,dc=example,dc=org"),
// 			BindCredential: pulumi.String("admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ldap.NewGroupMapper(ctx, "ldapGroupMapper", &ldap.GroupMapperArgs{
// 			RealmId:                realm.ID(),
// 			LdapUserFederationId:   ldapUserFederation.ID(),
// 			LdapGroupsDn:           pulumi.String("dc=example,dc=org"),
// 			GroupNameLdapAttribute: pulumi.String("cn"),
// 			GroupObjectClasses: pulumi.StringArray{
// 				pulumi.String("groupOfNames"),
// 			},
// 			MembershipAttributeType:     pulumi.String("DN"),
// 			MembershipLdapAttribute:     pulumi.String("member"),
// 			MembershipUserLdapAttribute: pulumi.String("cn"),
// 			MemberofLdapAttribute:       pulumi.String("memberOf"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GroupMapper struct {
	pulumi.CustomResourceState

	// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
	DropNonExistingGroupsDuringSync pulumi.BoolPtrOutput `pulumi:"dropNonExistingGroupsDuringSync"`
	// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
	GroupNameLdapAttribute pulumi.StringOutput `pulumi:"groupNameLdapAttribute"`
	// List of strings representing the object classes for the group. Must contain at least one.
	GroupObjectClasses pulumi.StringArrayOutput `pulumi:"groupObjectClasses"`
	// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
	GroupsLdapFilter pulumi.StringPtrOutput `pulumi:"groupsLdapFilter"`
	// When `true`, missing groups in the hierarchy will be ignored.
	IgnoreMissingGroups pulumi.BoolPtrOutput `pulumi:"ignoreMissingGroups"`
	// The LDAP DN where groups can be found.
	LdapGroupsDn pulumi.StringOutput `pulumi:"ldapGroupsDn"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
	MappedGroupAttributes pulumi.StringArrayOutput `pulumi:"mappedGroupAttributes"`
	// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
	MemberofLdapAttribute pulumi.StringPtrOutput `pulumi:"memberofLdapAttribute"`
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType pulumi.StringPtrOutput `pulumi:"membershipAttributeType"`
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute pulumi.StringOutput `pulumi:"membershipLdapAttribute"`
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute pulumi.StringOutput `pulumi:"membershipUserLdapAttribute"`
	// Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
	PreserveGroupInheritance pulumi.BoolPtrOutput `pulumi:"preserveGroupInheritance"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy pulumi.StringPtrOutput `pulumi:"userRolesRetrieveStrategy"`
}

// NewGroupMapper registers a new resource with the given unique name, arguments, and options.
func NewGroupMapper(ctx *pulumi.Context,
	name string, args *GroupMapperArgs, opts ...pulumi.ResourceOption) (*GroupMapper, error) {
	if args == nil || args.GroupNameLdapAttribute == nil {
		return nil, errors.New("missing required argument 'GroupNameLdapAttribute'")
	}
	if args == nil || args.GroupObjectClasses == nil {
		return nil, errors.New("missing required argument 'GroupObjectClasses'")
	}
	if args == nil || args.LdapGroupsDn == nil {
		return nil, errors.New("missing required argument 'LdapGroupsDn'")
	}
	if args == nil || args.LdapUserFederationId == nil {
		return nil, errors.New("missing required argument 'LdapUserFederationId'")
	}
	if args == nil || args.MembershipLdapAttribute == nil {
		return nil, errors.New("missing required argument 'MembershipLdapAttribute'")
	}
	if args == nil || args.MembershipUserLdapAttribute == nil {
		return nil, errors.New("missing required argument 'MembershipUserLdapAttribute'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GroupMapperArgs{}
	}
	var resource GroupMapper
	err := ctx.RegisterResource("keycloak:ldap/groupMapper:GroupMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMapper gets an existing GroupMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMapperState, opts ...pulumi.ResourceOption) (*GroupMapper, error) {
	var resource GroupMapper
	err := ctx.ReadResource("keycloak:ldap/groupMapper:GroupMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMapper resources.
type groupMapperState struct {
	// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
	DropNonExistingGroupsDuringSync *bool `pulumi:"dropNonExistingGroupsDuringSync"`
	// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
	GroupNameLdapAttribute *string `pulumi:"groupNameLdapAttribute"`
	// List of strings representing the object classes for the group. Must contain at least one.
	GroupObjectClasses []string `pulumi:"groupObjectClasses"`
	// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
	GroupsLdapFilter *string `pulumi:"groupsLdapFilter"`
	// When `true`, missing groups in the hierarchy will be ignored.
	IgnoreMissingGroups *bool `pulumi:"ignoreMissingGroups"`
	// The LDAP DN where groups can be found.
	LdapGroupsDn *string `pulumi:"ldapGroupsDn"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
	MappedGroupAttributes []string `pulumi:"mappedGroupAttributes"`
	// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
	MemberofLdapAttribute *string `pulumi:"memberofLdapAttribute"`
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType *string `pulumi:"membershipAttributeType"`
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute *string `pulumi:"membershipLdapAttribute"`
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute *string `pulumi:"membershipUserLdapAttribute"`
	// Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
	Mode *string `pulumi:"mode"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
	PreserveGroupInheritance *bool `pulumi:"preserveGroupInheritance"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
	// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy *string `pulumi:"userRolesRetrieveStrategy"`
}

type GroupMapperState struct {
	// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
	DropNonExistingGroupsDuringSync pulumi.BoolPtrInput
	// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
	GroupNameLdapAttribute pulumi.StringPtrInput
	// List of strings representing the object classes for the group. Must contain at least one.
	GroupObjectClasses pulumi.StringArrayInput
	// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
	GroupsLdapFilter pulumi.StringPtrInput
	// When `true`, missing groups in the hierarchy will be ignored.
	IgnoreMissingGroups pulumi.BoolPtrInput
	// The LDAP DN where groups can be found.
	LdapGroupsDn pulumi.StringPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
	MappedGroupAttributes pulumi.StringArrayInput
	// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
	MemberofLdapAttribute pulumi.StringPtrInput
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType pulumi.StringPtrInput
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute pulumi.StringPtrInput
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute pulumi.StringPtrInput
	// Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
	Mode pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
	PreserveGroupInheritance pulumi.BoolPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
	// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy pulumi.StringPtrInput
}

func (GroupMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMapperState)(nil)).Elem()
}

type groupMapperArgs struct {
	// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
	DropNonExistingGroupsDuringSync *bool `pulumi:"dropNonExistingGroupsDuringSync"`
	// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
	GroupNameLdapAttribute string `pulumi:"groupNameLdapAttribute"`
	// List of strings representing the object classes for the group. Must contain at least one.
	GroupObjectClasses []string `pulumi:"groupObjectClasses"`
	// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
	GroupsLdapFilter *string `pulumi:"groupsLdapFilter"`
	// When `true`, missing groups in the hierarchy will be ignored.
	IgnoreMissingGroups *bool `pulumi:"ignoreMissingGroups"`
	// The LDAP DN where groups can be found.
	LdapGroupsDn string `pulumi:"ldapGroupsDn"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
	MappedGroupAttributes []string `pulumi:"mappedGroupAttributes"`
	// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
	MemberofLdapAttribute *string `pulumi:"memberofLdapAttribute"`
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType *string `pulumi:"membershipAttributeType"`
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute string `pulumi:"membershipLdapAttribute"`
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute string `pulumi:"membershipUserLdapAttribute"`
	// Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
	Mode *string `pulumi:"mode"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
	PreserveGroupInheritance *bool `pulumi:"preserveGroupInheritance"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
	// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy *string `pulumi:"userRolesRetrieveStrategy"`
}

// The set of arguments for constructing a GroupMapper resource.
type GroupMapperArgs struct {
	// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
	DropNonExistingGroupsDuringSync pulumi.BoolPtrInput
	// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
	GroupNameLdapAttribute pulumi.StringInput
	// List of strings representing the object classes for the group. Must contain at least one.
	GroupObjectClasses pulumi.StringArrayInput
	// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
	GroupsLdapFilter pulumi.StringPtrInput
	// When `true`, missing groups in the hierarchy will be ignored.
	IgnoreMissingGroups pulumi.BoolPtrInput
	// The LDAP DN where groups can be found.
	LdapGroupsDn pulumi.StringInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
	MappedGroupAttributes pulumi.StringArrayInput
	// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
	MemberofLdapAttribute pulumi.StringPtrInput
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType pulumi.StringPtrInput
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute pulumi.StringInput
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute pulumi.StringInput
	// Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
	Mode pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
	PreserveGroupInheritance pulumi.BoolPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
	// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy pulumi.StringPtrInput
}

func (GroupMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMapperArgs)(nil)).Elem()
}
