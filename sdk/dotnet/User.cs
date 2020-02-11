// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// ## # keycloak..User
    /// 
    /// Allows for creating and managing Users within Keycloak.
    /// 
    /// This resource was created primarily to enable the acceptance tests for the `keycloak..Group` resource.
    /// Creating users within Keycloak is not recommended. Instead, users should be federated from external sources
    /// by configuring user federation providers or identity providers.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this user belongs to.
    /// - `username` - (Required) The unique username of this user.
    /// - `initial_password` (Optional) When given, the user's initial password will be set.
    ///    This attribute is only respected during initial user creation.
    ///     - `value` (Required) The initial password.
    ///     - `temporary` (Optional) If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
    /// - `enabled` - (Optional) When false, this user cannot log in. Defaults to `true`.
    /// - `email` - (Optional) The user's email.
    /// - `first_name` - (Optional) The user's first name.
    /// - `last_name` - (Optional) The user's last name.
    /// 
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/user.html.markdown.
    /// </summary>
    public partial class User : Pulumi.CustomResource
    {
        [Output("attributes")]
        public Output<ImmutableDictionary<string, object>?> Attributes { get; private set; } = null!;

        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("federatedIdentities")]
        public Output<ImmutableArray<Outputs.UserFederatedIdentities>> FederatedIdentities { get; private set; } = null!;

        [Output("firstName")]
        public Output<string?> FirstName { get; private set; } = null!;

        [Output("initialPassword")]
        public Output<Outputs.UserInitialPassword?> InitialPassword { get; private set; } = null!;

        [Output("lastName")]
        public Output<string?> LastName { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/user:User", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("federatedIdentities")]
        private InputList<Inputs.UserFederatedIdentitiesArgs>? _federatedIdentities;
        public InputList<Inputs.UserFederatedIdentitiesArgs> FederatedIdentities
        {
            get => _federatedIdentities ?? (_federatedIdentities = new InputList<Inputs.UserFederatedIdentitiesArgs>());
            set => _federatedIdentities = value;
        }

        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        [Input("initialPassword")]
        public Input<Inputs.UserInitialPasswordArgs>? InitialPassword { get; set; }

        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("federatedIdentities")]
        private InputList<Inputs.UserFederatedIdentitiesGetArgs>? _federatedIdentities;
        public InputList<Inputs.UserFederatedIdentitiesGetArgs> FederatedIdentities
        {
            get => _federatedIdentities ?? (_federatedIdentities = new InputList<Inputs.UserFederatedIdentitiesGetArgs>());
            set => _federatedIdentities = value;
        }

        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        [Input("initialPassword")]
        public Input<Inputs.UserInitialPasswordGetArgs>? InitialPassword { get; set; }

        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public UserState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class UserFederatedIdentitiesArgs : Pulumi.ResourceArgs
    {
        [Input("identityProvider", required: true)]
        public Input<string> IdentityProvider { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserFederatedIdentitiesArgs()
        {
        }
    }

    public sealed class UserFederatedIdentitiesGetArgs : Pulumi.ResourceArgs
    {
        [Input("identityProvider", required: true)]
        public Input<string> IdentityProvider { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserFederatedIdentitiesGetArgs()
        {
        }
    }

    public sealed class UserInitialPasswordArgs : Pulumi.ResourceArgs
    {
        [Input("temporary")]
        public Input<bool>? Temporary { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public UserInitialPasswordArgs()
        {
        }
    }

    public sealed class UserInitialPasswordGetArgs : Pulumi.ResourceArgs
    {
        [Input("temporary")]
        public Input<bool>? Temporary { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public UserInitialPasswordGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class UserFederatedIdentities
    {
        public readonly string IdentityProvider;
        public readonly string UserId;
        public readonly string UserName;

        [OutputConstructor]
        private UserFederatedIdentities(
            string identityProvider,
            string userId,
            string userName)
        {
            IdentityProvider = identityProvider;
            UserId = userId;
            UserName = userName;
        }
    }

    [OutputType]
    public sealed class UserInitialPassword
    {
        public readonly bool? Temporary;
        public readonly string Value;

        [OutputConstructor]
        private UserInitialPassword(
            bool? temporary,
            string value)
        {
            Temporary = temporary;
            Value = value;
        }
    }
    }
}