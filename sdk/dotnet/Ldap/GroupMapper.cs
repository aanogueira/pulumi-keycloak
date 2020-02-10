// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap
{
    /// <summary>
    /// ## # keycloak.Ldap.GroupMapper
    /// 
    /// Allows for creating and managing group mappers for Keycloak users federated
    /// via LDAP.
    /// 
    /// The LDAP group mapper can be used to map an LDAP user's groups from some DN
    /// to Keycloak groups. This group mapper will also create the groups within Keycloak
    /// if they do not already exist.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
    /// - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
    /// - `name` - (Required) Display name of this mapper when displayed in the console.
    /// - `ldap_groups_dn` - (Required) The LDAP DN where groups can be found.
    /// - `group_name_ldap_attribute` - (Required) The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
    /// - `group_object_classes` - (Required) Array of strings representing the object classes for the group. Must contain at least one.
    /// - `preserve_group_inheritance` - (Optional) When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
    /// - `ignore_missing_groups` - (Optional) When `true`, missing groups in the hierarchy will be ignored.
    /// - `membership_ldap_attribute` - (Required) The name of the LDAP attribute that is used for membership mappings.
    /// - `membership_attribute_type` - (Optional) Can be one of `DN` or `UID`. Defaults to `DN`.
    /// - `membership_user_ldap_attribute` - (Required) The name of the LDAP attribute on a user that is used for membership mappings.
    /// - `groups_ldap_filter` - (Optional) When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
    /// - `mode` - (Optional) Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
    /// - `user_roles_retrieve_strategy` - (Optional) Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
    /// - `memberof_ldap_attribute` - (Optional) Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
    /// - `mapped_group_attributes` - (Optional) Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
    /// - `drop_non_existing_groups_during_sync` - (Optional) When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
    /// 
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/ldap_group_mapper.html.markdown.
    /// </summary>
    public partial class GroupMapper : Pulumi.CustomResource
    {
        [Output("dropNonExistingGroupsDuringSync")]
        public Output<bool?> DropNonExistingGroupsDuringSync { get; private set; } = null!;

        [Output("groupNameLdapAttribute")]
        public Output<string> GroupNameLdapAttribute { get; private set; } = null!;

        [Output("groupObjectClasses")]
        public Output<ImmutableArray<string>> GroupObjectClasses { get; private set; } = null!;

        [Output("groupsLdapFilter")]
        public Output<string?> GroupsLdapFilter { get; private set; } = null!;

        [Output("ignoreMissingGroups")]
        public Output<bool?> IgnoreMissingGroups { get; private set; } = null!;

        [Output("ldapGroupsDn")]
        public Output<string> LdapGroupsDn { get; private set; } = null!;

        /// <summary>
        /// The ldap user federation provider to attach this mapper to.
        /// </summary>
        [Output("ldapUserFederationId")]
        public Output<string> LdapUserFederationId { get; private set; } = null!;

        [Output("mappedGroupAttributes")]
        public Output<ImmutableArray<string>> MappedGroupAttributes { get; private set; } = null!;

        [Output("memberofLdapAttribute")]
        public Output<string?> MemberofLdapAttribute { get; private set; } = null!;

        [Output("membershipAttributeType")]
        public Output<string?> MembershipAttributeType { get; private set; } = null!;

        [Output("membershipLdapAttribute")]
        public Output<string> MembershipLdapAttribute { get; private set; } = null!;

        [Output("membershipUserLdapAttribute")]
        public Output<string> MembershipUserLdapAttribute { get; private set; } = null!;

        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// Display name of the mapper when displayed in the console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("preserveGroupInheritance")]
        public Output<bool?> PreserveGroupInheritance { get; private set; } = null!;

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("userRolesRetrieveStrategy")]
        public Output<string?> UserRolesRetrieveStrategy { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMapper(string name, GroupMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:Ldap/groupMapper:GroupMapper", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GroupMapper(string name, Input<string> id, GroupMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:Ldap/groupMapper:GroupMapper", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GroupMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMapper Get(string name, Input<string> id, GroupMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMapper(name, id, state, options);
        }
    }

    public sealed class GroupMapperArgs : Pulumi.ResourceArgs
    {
        [Input("dropNonExistingGroupsDuringSync")]
        public Input<bool>? DropNonExistingGroupsDuringSync { get; set; }

        [Input("groupNameLdapAttribute", required: true)]
        public Input<string> GroupNameLdapAttribute { get; set; } = null!;

        [Input("groupObjectClasses", required: true)]
        private InputList<string>? _groupObjectClasses;
        public InputList<string> GroupObjectClasses
        {
            get => _groupObjectClasses ?? (_groupObjectClasses = new InputList<string>());
            set => _groupObjectClasses = value;
        }

        [Input("groupsLdapFilter")]
        public Input<string>? GroupsLdapFilter { get; set; }

        [Input("ignoreMissingGroups")]
        public Input<bool>? IgnoreMissingGroups { get; set; }

        [Input("ldapGroupsDn", required: true)]
        public Input<string> LdapGroupsDn { get; set; } = null!;

        /// <summary>
        /// The ldap user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId", required: true)]
        public Input<string> LdapUserFederationId { get; set; } = null!;

        [Input("mappedGroupAttributes")]
        private InputList<string>? _mappedGroupAttributes;
        public InputList<string> MappedGroupAttributes
        {
            get => _mappedGroupAttributes ?? (_mappedGroupAttributes = new InputList<string>());
            set => _mappedGroupAttributes = value;
        }

        [Input("memberofLdapAttribute")]
        public Input<string>? MemberofLdapAttribute { get; set; }

        [Input("membershipAttributeType")]
        public Input<string>? MembershipAttributeType { get; set; }

        [Input("membershipLdapAttribute", required: true)]
        public Input<string> MembershipLdapAttribute { get; set; } = null!;

        [Input("membershipUserLdapAttribute", required: true)]
        public Input<string> MembershipUserLdapAttribute { get; set; } = null!;

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Display name of the mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("preserveGroupInheritance")]
        public Input<bool>? PreserveGroupInheritance { get; set; }

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("userRolesRetrieveStrategy")]
        public Input<string>? UserRolesRetrieveStrategy { get; set; }

        public GroupMapperArgs()
        {
        }
    }

    public sealed class GroupMapperState : Pulumi.ResourceArgs
    {
        [Input("dropNonExistingGroupsDuringSync")]
        public Input<bool>? DropNonExistingGroupsDuringSync { get; set; }

        [Input("groupNameLdapAttribute")]
        public Input<string>? GroupNameLdapAttribute { get; set; }

        [Input("groupObjectClasses")]
        private InputList<string>? _groupObjectClasses;
        public InputList<string> GroupObjectClasses
        {
            get => _groupObjectClasses ?? (_groupObjectClasses = new InputList<string>());
            set => _groupObjectClasses = value;
        }

        [Input("groupsLdapFilter")]
        public Input<string>? GroupsLdapFilter { get; set; }

        [Input("ignoreMissingGroups")]
        public Input<bool>? IgnoreMissingGroups { get; set; }

        [Input("ldapGroupsDn")]
        public Input<string>? LdapGroupsDn { get; set; }

        /// <summary>
        /// The ldap user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId")]
        public Input<string>? LdapUserFederationId { get; set; }

        [Input("mappedGroupAttributes")]
        private InputList<string>? _mappedGroupAttributes;
        public InputList<string> MappedGroupAttributes
        {
            get => _mappedGroupAttributes ?? (_mappedGroupAttributes = new InputList<string>());
            set => _mappedGroupAttributes = value;
        }

        [Input("memberofLdapAttribute")]
        public Input<string>? MemberofLdapAttribute { get; set; }

        [Input("membershipAttributeType")]
        public Input<string>? MembershipAttributeType { get; set; }

        [Input("membershipLdapAttribute")]
        public Input<string>? MembershipLdapAttribute { get; set; }

        [Input("membershipUserLdapAttribute")]
        public Input<string>? MembershipUserLdapAttribute { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Display name of the mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("preserveGroupInheritance")]
        public Input<bool>? PreserveGroupInheritance { get; set; }

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("userRolesRetrieveStrategy")]
        public Input<string>? UserRolesRetrieveStrategy { get; set; }

        public GroupMapperState()
        {
        }
    }
}
