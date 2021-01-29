// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Authentication
{
    /// <summary>
    /// Allows for creating and managing an authentication execution within Keycloak.
    /// 
    /// An authentication execution is an action that the user or service may or may not take when authenticating through an authentication
    /// flow.
    /// 
    /// &gt; Due to limitations in the Keycloak API, the ordering of authentication executions within a flow must be specified using `depends_on`. Authentication executions that are created first will appear first within the flow.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Realm = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var flow = new Keycloak.Authentication.Flow("flow", new Keycloak.Authentication.FlowArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             Alias = "my-flow-alias",
    ///         });
    ///         // first execution
    ///         var executionOne = new Keycloak.Authentication.Execution("executionOne", new Keycloak.Authentication.ExecutionArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ParentFlowAlias = flow.Alias,
    ///             Authenticator = "auth-cookie",
    ///             Requirement = "ALTERNATIVE",
    ///         });
    ///         // second execution
    ///         var executionTwo = new Keycloak.Authentication.Execution("executionTwo", new Keycloak.Authentication.ExecutionArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ParentFlowAlias = flow.Alias,
    ///             Authenticator = "identity-provider-redirector",
    ///             Requirement = "ALTERNATIVE",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 executionOne,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Authentication executions can be imported using the formats`{{realmId}}/{{parentFlowAlias}}/{{authenticationExecutionId}}`. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:authentication/execution:Execution keycloak_authentication_execution my-realm/my-flow/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:authentication/execution:Execution")]
    public partial class Execution : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        /// </summary>
        [Output("authenticator")]
        public Output<string> Authenticator { get; private set; } = null!;

        /// <summary>
        /// The alias of the flow this execution is attached to.
        /// </summary>
        [Output("parentFlowAlias")]
        public Output<string> ParentFlowAlias { get; private set; } = null!;

        /// <summary>
        /// The realm the authentication execution exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        /// </summary>
        [Output("requirement")]
        public Output<string?> Requirement { get; private set; } = null!;


        /// <summary>
        /// Create a Execution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Execution(string name, ExecutionArgs args, CustomResourceOptions? options = null)
            : base("keycloak:authentication/execution:Execution", name, args ?? new ExecutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Execution(string name, Input<string> id, ExecutionState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:authentication/execution:Execution", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Execution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Execution Get(string name, Input<string> id, ExecutionState? state = null, CustomResourceOptions? options = null)
        {
            return new Execution(name, id, state, options);
        }
    }

    public sealed class ExecutionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        /// </summary>
        [Input("authenticator", required: true)]
        public Input<string> Authenticator { get; set; } = null!;

        /// <summary>
        /// The alias of the flow this execution is attached to.
        /// </summary>
        [Input("parentFlowAlias", required: true)]
        public Input<string> ParentFlowAlias { get; set; } = null!;

        /// <summary>
        /// The realm the authentication execution exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        /// </summary>
        [Input("requirement")]
        public Input<string>? Requirement { get; set; }

        public ExecutionArgs()
        {
        }
    }

    public sealed class ExecutionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        /// </summary>
        [Input("authenticator")]
        public Input<string>? Authenticator { get; set; }

        /// <summary>
        /// The alias of the flow this execution is attached to.
        /// </summary>
        [Input("parentFlowAlias")]
        public Input<string>? ParentFlowAlias { get; set; }

        /// <summary>
        /// The realm the authentication execution exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        /// </summary>
        [Input("requirement")]
        public Input<string>? Requirement { get; set; }

        public ExecutionState()
        {
        }
    }
}
