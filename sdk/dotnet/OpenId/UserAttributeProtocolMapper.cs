// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// ## # keycloak.OpenId.UserAttributeProtocolMapper
    /// 
    /// Allows for creating and managing user attribute protocol mappers within
    /// Keycloak.
    /// 
    /// User attribute protocol mappers allow you to map custom attributes defined
    /// for a user within Keycloak to a claim in a token. Protocol mappers can be
    /// defined for a single client, or they can be defined for a client scope which
    /// can be shared between multiple different clients.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this protocol mapper exists within.
    /// - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
    /// - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
    /// - `name` - (Required) The display name of this protocol mapper in the GUI.
    /// - `user_attribute` - (Required) The custom user attribute to map a claim for.
    /// - `claim_name` - (Required) The name of the claim to insert into a token.
    /// - `claim_value_type` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
    /// - `multivalued` - (Optional) Indicates whether this attribute is a single value or an array of values. Defaults to `false`.
    /// - `add_to_id_token` - (Optional) Indicates if the attribute should be added as a claim to the id token. Defaults to `true`.
    /// - `add_to_access_token` - (Optional) Indicates if the attribute should be added as a claim to the access token. Defaults to `true`.
    /// - `add_to_userinfo` - (Optional) Indicates if the attribute should be added as a claim to the UserInfo response body. Defaults to `true`.
    /// 
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_user_attribute_protocol_mapper.html.markdown.
    /// </summary>
    public partial class UserAttributeProtocolMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the attribute should be a claim in the access token.
        /// </summary>
        [Output("addToAccessToken")]
        public Output<bool?> AddToAccessToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if the attribute should be a claim in the id token.
        /// </summary>
        [Output("addToIdToken")]
        public Output<bool?> AddToIdToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if the attribute should appear in the userinfo response body.
        /// </summary>
        [Output("addToUserinfo")]
        public Output<bool?> AddToUserinfo { get; private set; } = null!;

        [Output("claimName")]
        public Output<string> ClaimName { get; private set; } = null!;

        /// <summary>
        /// Claim type used when serializing tokens.
        /// </summary>
        [Output("claimValueType")]
        public Output<string?> ClaimValueType { get; private set; } = null!;

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this attribute is a single value or an array of values.
        /// </summary>
        [Output("multivalued")]
        public Output<bool?> Multivalued { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("userAttribute")]
        public Output<string> UserAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a UserAttributeProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserAttributeProtocolMapper(string name, UserAttributeProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:OpenId/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private UserAttributeProtocolMapper(string name, Input<string> id, UserAttributeProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:OpenId/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserAttributeProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserAttributeProtocolMapper Get(string name, Input<string> id, UserAttributeProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new UserAttributeProtocolMapper(name, id, state, options);
        }
    }

    public sealed class UserAttributeProtocolMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the attribute should be a claim in the access token.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should be a claim in the id token.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should appear in the userinfo response body.
        /// </summary>
        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        [Input("claimName", required: true)]
        public Input<string> ClaimName { get; set; } = null!;

        /// <summary>
        /// Claim type used when serializing tokens.
        /// </summary>
        [Input("claimValueType")]
        public Input<string>? ClaimValueType { get; set; }

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// Indicates whether this attribute is a single value or an array of values.
        /// </summary>
        [Input("multivalued")]
        public Input<bool>? Multivalued { get; set; }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("userAttribute", required: true)]
        public Input<string> UserAttribute { get; set; } = null!;

        public UserAttributeProtocolMapperArgs()
        {
        }
    }

    public sealed class UserAttributeProtocolMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the attribute should be a claim in the access token.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should be a claim in the id token.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should appear in the userinfo response body.
        /// </summary>
        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// Claim type used when serializing tokens.
        /// </summary>
        [Input("claimValueType")]
        public Input<string>? ClaimValueType { get; set; }

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// Indicates whether this attribute is a single value or an array of values.
        /// </summary>
        [Input("multivalued")]
        public Input<bool>? Multivalued { get; set; }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("userAttribute")]
        public Input<string>? UserAttribute { get; set; }

        public UserAttributeProtocolMapperState()
        {
        }
    }
}