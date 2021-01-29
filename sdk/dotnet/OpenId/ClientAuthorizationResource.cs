// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    [KeycloakResourceType("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource")]
    public partial class ClientAuthorizationResource : Pulumi.CustomResource
    {
        [Output("attributes")]
        public Output<ImmutableDictionary<string, object>?> Attributes { get; private set; } = null!;

        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        [Output("iconUri")]
        public Output<string?> IconUri { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("ownerManagedAccess")]
        public Output<bool?> OwnerManagedAccess { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;

        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("uris")]
        public Output<ImmutableArray<string>> Uris { get; private set; } = null!;


        /// <summary>
        /// Create a ClientAuthorizationResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientAuthorizationResource(string name, ClientAuthorizationResourceArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource", name, args ?? new ClientAuthorizationResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientAuthorizationResource(string name, Input<string> id, ClientAuthorizationResourceState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientAuthorizationResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientAuthorizationResource Get(string name, Input<string> id, ClientAuthorizationResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientAuthorizationResource(name, id, state, options);
        }
    }

    public sealed class ClientAuthorizationResourceArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("iconUri")]
        public Input<string>? IconUri { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerManagedAccess")]
        public Input<bool>? OwnerManagedAccess { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("resourceServerId", required: true)]
        public Input<string> ResourceServerId { get; set; } = null!;

        [Input("scopes")]
        private InputList<string>? _scopes;
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        public ClientAuthorizationResourceArgs()
        {
        }
    }

    public sealed class ClientAuthorizationResourceState : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("iconUri")]
        public Input<string>? IconUri { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerManagedAccess")]
        public Input<bool>? OwnerManagedAccess { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        public ClientAuthorizationResourceState()
        {
        }
    }
}
