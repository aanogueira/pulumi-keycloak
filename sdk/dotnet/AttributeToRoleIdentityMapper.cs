// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public partial class AttributeToRoleIdentityMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// Attribute Friendly Name
        /// </summary>
        [Output("attributeFriendlyName")]
        public Output<string?> AttributeFriendlyName { get; private set; } = null!;

        /// <summary>
        /// Attribute Name
        /// </summary>
        [Output("attributeName")]
        public Output<string?> AttributeName { get; private set; } = null!;

        /// <summary>
        /// Attribute Value
        /// </summary>
        [Output("attributeValue")]
        public Output<string?> AttributeValue { get; private set; } = null!;

        /// <summary>
        /// OIDC Claim Name
        /// </summary>
        [Output("claimName")]
        public Output<string?> ClaimName { get; private set; } = null!;

        /// <summary>
        /// OIDC Claim Value
        /// </summary>
        [Output("claimValue")]
        public Output<string?> ClaimValue { get; private set; } = null!;

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Output("identityProviderAlias")]
        public Output<string> IdentityProviderAlias { get; private set; } = null!;

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Realm Name
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// Role Name
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a AttributeToRoleIdentityMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttributeToRoleIdentityMapper(string name, AttributeToRoleIdentityMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, args ?? new AttributeToRoleIdentityMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttributeToRoleIdentityMapper(string name, Input<string> id, AttributeToRoleIdentityMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AttributeToRoleIdentityMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttributeToRoleIdentityMapper Get(string name, Input<string> id, AttributeToRoleIdentityMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new AttributeToRoleIdentityMapper(name, id, state, options);
        }
    }

    public sealed class AttributeToRoleIdentityMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attribute Friendly Name
        /// </summary>
        [Input("attributeFriendlyName")]
        public Input<string>? AttributeFriendlyName { get; set; }

        /// <summary>
        /// Attribute Name
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// Attribute Value
        /// </summary>
        [Input("attributeValue")]
        public Input<string>? AttributeValue { get; set; }

        /// <summary>
        /// OIDC Claim Name
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// OIDC Claim Value
        /// </summary>
        [Input("claimValue")]
        public Input<string>? ClaimValue { get; set; }

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Input("identityProviderAlias", required: true)]
        public Input<string> IdentityProviderAlias { get; set; } = null!;

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// Role Name
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public AttributeToRoleIdentityMapperArgs()
        {
        }
    }

    public sealed class AttributeToRoleIdentityMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attribute Friendly Name
        /// </summary>
        [Input("attributeFriendlyName")]
        public Input<string>? AttributeFriendlyName { get; set; }

        /// <summary>
        /// Attribute Name
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// Attribute Value
        /// </summary>
        [Input("attributeValue")]
        public Input<string>? AttributeValue { get; set; }

        /// <summary>
        /// OIDC Claim Name
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// OIDC Claim Value
        /// </summary>
        [Input("claimValue")]
        public Input<string>? ClaimValue { get; set; }

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Input("identityProviderAlias")]
        public Input<string>? IdentityProviderAlias { get; set; }

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Role Name
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public AttributeToRoleIdentityMapperState()
        {
        }
    }
}
