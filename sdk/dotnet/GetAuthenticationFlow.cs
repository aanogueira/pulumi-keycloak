// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static class GetAuthenticationFlow
    {
        /// <summary>
        /// This data source can be used to fetch the ID of an authentication flow within Keycloak.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var browserAuthCookie = Keycloak.GetAuthenticationFlow.Invoke(new Keycloak.GetAuthenticationFlowInvokeArgs
        ///         {
        ///             RealmId = realm.Id,
        ///             Alias = "browser",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthenticationFlowResult> InvokeAsync(GetAuthenticationFlowArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthenticationFlowResult>("keycloak:index/getAuthenticationFlow:getAuthenticationFlow", args ?? new GetAuthenticationFlowArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch the ID of an authentication flow within Keycloak.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var browserAuthCookie = Keycloak.GetAuthenticationFlow.Invoke(new Keycloak.GetAuthenticationFlowInvokeArgs
        ///         {
        ///             RealmId = realm.Id,
        ///             Alias = "browser",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAuthenticationFlowResult> Invoke(GetAuthenticationFlowInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthenticationFlowResult>("keycloak:index/getAuthenticationFlow:getAuthenticationFlow", args ?? new GetAuthenticationFlowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthenticationFlowArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alias of the flow.
        /// </summary>
        [Input("alias", required: true)]
        public string Alias { get; set; } = null!;

        /// <summary>
        /// The realm the authentication flow exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetAuthenticationFlowArgs()
        {
        }
    }

    public sealed class GetAuthenticationFlowInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alias of the flow.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// The realm the authentication flow exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GetAuthenticationFlowInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthenticationFlowResult
    {
        public readonly string Alias;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RealmId;

        [OutputConstructor]
        private GetAuthenticationFlowResult(
            string alias,

            string id,

            string realmId)
        {
            Alias = alias;
            Id = id;
            RealmId = realmId;
        }
    }
}
