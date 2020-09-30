// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    public partial class ClientDefaultScope : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// An array of client scope names to attach to this client.
        /// </summary>
        [Output("defaultScopes")]
        public Output<ImmutableArray<string>> DefaultScopes { get; private set; } = null!;

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientDefaultScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientDefaultScope(string name, ClientDefaultScopeArgs args, CustomResourceOptions? options = null)
            : base("keycloak:saml/clientDefaultScope:ClientDefaultScope", name, args ?? new ClientDefaultScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientDefaultScope(string name, Input<string> id, ClientDefaultScopeState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:saml/clientDefaultScope:ClientDefaultScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientDefaultScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientDefaultScope Get(string name, Input<string> id, ClientDefaultScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientDefaultScope(name, id, state, options);
        }
    }

    public sealed class ClientDefaultScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("defaultScopes", required: true)]
        private InputList<string>? _defaultScopes;

        /// <summary>
        /// An array of client scope names to attach to this client.
        /// </summary>
        public InputList<string> DefaultScopes
        {
            get => _defaultScopes ?? (_defaultScopes = new InputList<string>());
            set => _defaultScopes = value;
        }

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public ClientDefaultScopeArgs()
        {
        }
    }

    public sealed class ClientDefaultScopeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("defaultScopes")]
        private InputList<string>? _defaultScopes;

        /// <summary>
        /// An array of client scope names to attach to this client.
        /// </summary>
        public InputList<string> DefaultScopes
        {
            get => _defaultScopes ?? (_defaultScopes = new InputList<string>());
            set => _defaultScopes = value;
        }

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public ClientDefaultScopeState()
        {
        }
    }
}
