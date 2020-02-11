// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// ## # keycloak..Role
    /// 
    /// Allows for creating and managing roles within Keycloak.
    /// 
    /// Roles allow you define privileges within Keycloak and map them to users
    /// and groups.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this role exists within.
    /// - `client_id` - (Optional) When specified, this role will be created as
    ///   a client role attached to the client with the provided ID
    /// - `name` - (Required) The name of the role
    /// - `description` - (Optional) The description of the role
    /// - `composite_roles` - (Optional) When specified, this role will be a
    ///   composite role, composed of all roles that have an ID present within
    ///   this list.
    /// 
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/role.html.markdown.
    /// </summary>
    public partial class Role : Pulumi.CustomResource
    {
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        [Output("compositeRoles")]
        public Output<ImmutableArray<string>> CompositeRoles { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/role:Role", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/role:Role", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("compositeRoles")]
        private InputList<string>? _compositeRoles;
        public InputList<string> CompositeRoles
        {
            get => _compositeRoles ?? (_compositeRoles = new InputList<string>());
            set => _compositeRoles = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public RoleArgs()
        {
        }
    }

    public sealed class RoleState : Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("compositeRoles")]
        private InputList<string>? _compositeRoles;
        public InputList<string> CompositeRoles
        {
            get => _compositeRoles ?? (_compositeRoles = new InputList<string>());
            set => _compositeRoles = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public RoleState()
        {
        }
    }
}