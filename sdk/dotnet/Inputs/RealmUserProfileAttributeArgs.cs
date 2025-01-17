// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmUserProfileAttributeArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// A map of annotations for the group.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The display name of the attribute.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("enabledWhenScopes")]
        private InputList<string>? _enabledWhenScopes;

        /// <summary>
        /// A list of scopes. The attribute will only be enabled when these scopes are requested by clients.
        /// </summary>
        public InputList<string> EnabledWhenScopes
        {
            get => _enabledWhenScopes ?? (_enabledWhenScopes = new InputList<string>());
            set => _enabledWhenScopes = value;
        }

        /// <summary>
        /// The group that the attribute belong to.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The permissions configuration information.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.RealmUserProfileAttributePermissionsArgs>? Permissions { get; set; }

        [Input("requiredForRoles")]
        private InputList<string>? _requiredForRoles;

        /// <summary>
        /// A list of roles for which the attribute will be required.
        /// </summary>
        public InputList<string> RequiredForRoles
        {
            get => _requiredForRoles ?? (_requiredForRoles = new InputList<string>());
            set => _requiredForRoles = value;
        }

        [Input("requiredForScopes")]
        private InputList<string>? _requiredForScopes;

        /// <summary>
        /// A list of scopes for which the attribute will be required.
        /// </summary>
        public InputList<string> RequiredForScopes
        {
            get => _requiredForScopes ?? (_requiredForScopes = new InputList<string>());
            set => _requiredForScopes = value;
        }

        [Input("validators")]
        private InputList<Inputs.RealmUserProfileAttributeValidatorArgs>? _validators;

        /// <summary>
        /// A list of validators for the attribute.
        /// </summary>
        public InputList<Inputs.RealmUserProfileAttributeValidatorArgs> Validators
        {
            get => _validators ?? (_validators = new InputList<Inputs.RealmUserProfileAttributeValidatorArgs>());
            set => _validators = value;
        }

        public RealmUserProfileAttributeArgs()
        {
        }
    }
}
