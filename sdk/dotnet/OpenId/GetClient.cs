// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static class GetClient
    {
        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
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
        ///         var realmManagement = Output.Create(Keycloak.OpenId.GetClient.InvokeAsync(new Keycloak.OpenId.GetClientArgs
        ///         {
        ///             RealmId = "my-realm",
        ///             ClientId = "realm-management",
        ///         }));
        ///         var admin = realmManagement.Apply(realmManagement =&gt; Output.Create(Keycloak.GetRole.InvokeAsync(new Keycloak.GetRoleArgs
        ///         {
        ///             RealmId = "my-realm",
        ///             ClientId = realmManagement.Id,
        ///             Name = "realm-admin",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClientResult> InvokeAsync(GetClientArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientResult>("keycloak:openid/getClient:getClient", args ?? new GetClientArgs(), options.WithVersion());
    }


    public sealed class GetClientArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The client id (not its unique ID).
        /// </summary>
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        [Input("clientOfflineSessionIdleTimeout")]
        public string? ClientOfflineSessionIdleTimeout { get; set; }

        [Input("clientOfflineSessionMaxLifespan")]
        public string? ClientOfflineSessionMaxLifespan { get; set; }

        [Input("clientSessionIdleTimeout")]
        public string? ClientSessionIdleTimeout { get; set; }

        [Input("clientSessionMaxLifespan")]
        public string? ClientSessionMaxLifespan { get; set; }

        /// <summary>
        /// The realm id.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        [Input("useRefreshTokens")]
        public bool? UseRefreshTokens { get; set; }

        public GetClientArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClientResult
    {
        public readonly string AccessTokenLifespan;
        public readonly string AccessType;
        public readonly string AdminUrl;
        public readonly ImmutableArray<Outputs.GetClientAuthenticationFlowBindingOverrideResult> AuthenticationFlowBindingOverrides;
        public readonly ImmutableArray<Outputs.GetClientAuthorizationResult> Authorizations;
        public readonly string BaseUrl;
        public readonly string ClientId;
        public readonly string? ClientOfflineSessionIdleTimeout;
        public readonly string? ClientOfflineSessionMaxLifespan;
        public readonly string ClientSecret;
        public readonly string? ClientSessionIdleTimeout;
        public readonly string? ClientSessionMaxLifespan;
        public readonly bool ConsentRequired;
        public readonly string Description;
        public readonly bool DirectAccessGrantsEnabled;
        public readonly bool Enabled;
        public readonly bool ExcludeSessionStateFromAuthResponse;
        public readonly bool FullScopeAllowed;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool ImplicitFlowEnabled;
        public readonly string LoginTheme;
        public readonly string Name;
        public readonly string PkceCodeChallengeMethod;
        public readonly string RealmId;
        public readonly string ResourceServerId;
        public readonly string RootUrl;
        public readonly string ServiceAccountUserId;
        public readonly bool ServiceAccountsEnabled;
        public readonly bool StandardFlowEnabled;
        public readonly bool? UseRefreshTokens;
        public readonly ImmutableArray<string> ValidRedirectUris;
        public readonly ImmutableArray<string> WebOrigins;

        [OutputConstructor]
        private GetClientResult(
            string accessTokenLifespan,

            string accessType,

            string adminUrl,

            ImmutableArray<Outputs.GetClientAuthenticationFlowBindingOverrideResult> authenticationFlowBindingOverrides,

            ImmutableArray<Outputs.GetClientAuthorizationResult> authorizations,

            string baseUrl,

            string clientId,

            string? clientOfflineSessionIdleTimeout,

            string? clientOfflineSessionMaxLifespan,

            string clientSecret,

            string? clientSessionIdleTimeout,

            string? clientSessionMaxLifespan,

            bool consentRequired,

            string description,

            bool directAccessGrantsEnabled,

            bool enabled,

            bool excludeSessionStateFromAuthResponse,

            bool fullScopeAllowed,

            string id,

            bool implicitFlowEnabled,

            string loginTheme,

            string name,

            string pkceCodeChallengeMethod,

            string realmId,

            string resourceServerId,

            string rootUrl,

            string serviceAccountUserId,

            bool serviceAccountsEnabled,

            bool standardFlowEnabled,

            bool? useRefreshTokens,

            ImmutableArray<string> validRedirectUris,

            ImmutableArray<string> webOrigins)
        {
            AccessTokenLifespan = accessTokenLifespan;
            AccessType = accessType;
            AdminUrl = adminUrl;
            AuthenticationFlowBindingOverrides = authenticationFlowBindingOverrides;
            Authorizations = authorizations;
            BaseUrl = baseUrl;
            ClientId = clientId;
            ClientOfflineSessionIdleTimeout = clientOfflineSessionIdleTimeout;
            ClientOfflineSessionMaxLifespan = clientOfflineSessionMaxLifespan;
            ClientSecret = clientSecret;
            ClientSessionIdleTimeout = clientSessionIdleTimeout;
            ClientSessionMaxLifespan = clientSessionMaxLifespan;
            ConsentRequired = consentRequired;
            Description = description;
            DirectAccessGrantsEnabled = directAccessGrantsEnabled;
            Enabled = enabled;
            ExcludeSessionStateFromAuthResponse = excludeSessionStateFromAuthResponse;
            FullScopeAllowed = fullScopeAllowed;
            Id = id;
            ImplicitFlowEnabled = implicitFlowEnabled;
            LoginTheme = loginTheme;
            Name = name;
            PkceCodeChallengeMethod = pkceCodeChallengeMethod;
            RealmId = realmId;
            ResourceServerId = resourceServerId;
            RootUrl = rootUrl;
            ServiceAccountUserId = serviceAccountUserId;
            ServiceAccountsEnabled = serviceAccountsEnabled;
            StandardFlowEnabled = standardFlowEnabled;
            UseRefreshTokens = useRefreshTokens;
            ValidRedirectUris = validRedirectUris;
            WebOrigins = webOrigins;
        }
    }
}
