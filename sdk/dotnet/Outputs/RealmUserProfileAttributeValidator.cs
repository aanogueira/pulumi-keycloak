// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Outputs
{

    [OutputType]
    public sealed class RealmUserProfileAttributeValidator
    {
        /// <summary>
        /// A map defining the configuration of the validator.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Config;
        /// <summary>
        /// The name of the group.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private RealmUserProfileAttributeValidator(
            ImmutableDictionary<string, string>? config,

            string name)
        {
            Config = config;
            Name = name;
        }
    }
}
