// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public partial class ClientTimePolicy : Pulumi.CustomResource
    {
        [Output("dayMonth")]
        public Output<string?> DayMonth { get; private set; } = null!;

        [Output("dayMonthEnd")]
        public Output<string?> DayMonthEnd { get; private set; } = null!;

        [Output("decisionStrategy")]
        public Output<string> DecisionStrategy { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("hour")]
        public Output<string?> Hour { get; private set; } = null!;

        [Output("hourEnd")]
        public Output<string?> HourEnd { get; private set; } = null!;

        [Output("logic")]
        public Output<string?> Logic { get; private set; } = null!;

        [Output("minute")]
        public Output<string?> Minute { get; private set; } = null!;

        [Output("minuteEnd")]
        public Output<string?> MinuteEnd { get; private set; } = null!;

        [Output("month")]
        public Output<string?> Month { get; private set; } = null!;

        [Output("monthEnd")]
        public Output<string?> MonthEnd { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notBefore")]
        public Output<string?> NotBefore { get; private set; } = null!;

        [Output("notOnOrAfter")]
        public Output<string?> NotOnOrAfter { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;

        [Output("year")]
        public Output<string?> Year { get; private set; } = null!;

        [Output("yearEnd")]
        public Output<string?> YearEnd { get; private set; } = null!;


        /// <summary>
        /// Create a ClientTimePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientTimePolicy(string name, ClientTimePolicyArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientTimePolicy:ClientTimePolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ClientTimePolicy(string name, Input<string> id, ClientTimePolicyState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientTimePolicy:ClientTimePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientTimePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientTimePolicy Get(string name, Input<string> id, ClientTimePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientTimePolicy(name, id, state, options);
        }
    }

    public sealed class ClientTimePolicyArgs : Pulumi.ResourceArgs
    {
        [Input("dayMonth")]
        public Input<string>? DayMonth { get; set; }

        [Input("dayMonthEnd")]
        public Input<string>? DayMonthEnd { get; set; }

        [Input("decisionStrategy", required: true)]
        public Input<string> DecisionStrategy { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hour")]
        public Input<string>? Hour { get; set; }

        [Input("hourEnd")]
        public Input<string>? HourEnd { get; set; }

        [Input("logic")]
        public Input<string>? Logic { get; set; }

        [Input("minute")]
        public Input<string>? Minute { get; set; }

        [Input("minuteEnd")]
        public Input<string>? MinuteEnd { get; set; }

        [Input("month")]
        public Input<string>? Month { get; set; }

        [Input("monthEnd")]
        public Input<string>? MonthEnd { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notBefore")]
        public Input<string>? NotBefore { get; set; }

        [Input("notOnOrAfter")]
        public Input<string>? NotOnOrAfter { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("resourceServerId", required: true)]
        public Input<string> ResourceServerId { get; set; } = null!;

        [Input("year")]
        public Input<string>? Year { get; set; }

        [Input("yearEnd")]
        public Input<string>? YearEnd { get; set; }

        public ClientTimePolicyArgs()
        {
        }
    }

    public sealed class ClientTimePolicyState : Pulumi.ResourceArgs
    {
        [Input("dayMonth")]
        public Input<string>? DayMonth { get; set; }

        [Input("dayMonthEnd")]
        public Input<string>? DayMonthEnd { get; set; }

        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hour")]
        public Input<string>? Hour { get; set; }

        [Input("hourEnd")]
        public Input<string>? HourEnd { get; set; }

        [Input("logic")]
        public Input<string>? Logic { get; set; }

        [Input("minute")]
        public Input<string>? Minute { get; set; }

        [Input("minuteEnd")]
        public Input<string>? MinuteEnd { get; set; }

        [Input("month")]
        public Input<string>? Month { get; set; }

        [Input("monthEnd")]
        public Input<string>? MonthEnd { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notBefore")]
        public Input<string>? NotBefore { get; set; }

        [Input("notOnOrAfter")]
        public Input<string>? NotOnOrAfter { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        [Input("year")]
        public Input<string>? Year { get; set; }

        [Input("yearEnd")]
        public Input<string>? YearEnd { get; set; }

        public ClientTimePolicyState()
        {
        }
    }
}