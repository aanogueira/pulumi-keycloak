// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Oidc
{
    public partial class GoogleIdentityProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In
        /// case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly
        /// returned to client, but the request with prompt=none will be forwarded to this identity provider.
        /// </summary>
        [Output("acceptsPromptNoneForwardFromClient")]
        public Output<bool?> AcceptsPromptNoneForwardFromClient { get; private set; } = null!;

        /// <summary>
        /// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        /// </summary>
        [Output("addReadTokenRoleOnCreate")]
        public Output<bool?> AddReadTokenRoleOnCreate { get; private set; } = null!;

        /// <summary>
        /// The alias uniquely identifies an identity provider and it is also used to build the redirect uri. In case of google this
        /// is computed and always google
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Output("authenticateByDefault")]
        public Output<bool?> AuthenticateByDefault { get; private set; } = null!;

        /// <summary>
        /// Client ID.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client Secret.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The scopes to be sent when asking for authorization. See the documentation for possible values, separator and default
        /// value'. Default: 'openid profile email'
        /// </summary>
        [Output("defaultScopes")]
        public Output<string?> DefaultScopes { get; private set; } = null!;

        /// <summary>
        /// Disable usage of User Info service to obtain additional user information? Default is to use this OIDC service.
        /// </summary>
        [Output("disableUserInfo")]
        public Output<bool?> DisableUserInfo { get; private set; } = null!;

        /// <summary>
        /// Not used by this provider, Will be implicitly Google
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this identity provider.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        /// <summary>
        /// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
        /// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        /// </summary>
        [Output("firstBrokerLoginFlowAlias")]
        public Output<string?> FirstBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// Hide On Login Page.
        /// </summary>
        [Output("hideOnLoginPage")]
        public Output<bool?> HideOnLoginPage { get; private set; } = null!;

        /// <summary>
        /// Set 'hd' query parameter when logging in with Google. Google will list accounts only for this domain. Keycloak validates
        /// that the returned identity token has a claim for this domain. When '*' is entered, any hosted account can be used.
        /// </summary>
        [Output("hostedDomain")]
        public Output<string?> HostedDomain { get; private set; } = null!;

        /// <summary>
        /// Internal Identity Provider Id
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
        /// want to allow login from the provider, but want to integrate with a provider
        /// </summary>
        [Output("linkOnly")]
        public Output<bool?> LinkOnly { get; private set; } = null!;

        /// <summary>
        /// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
        /// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
        /// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
        /// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        /// </summary>
        [Output("postBrokerLoginFlowAlias")]
        public Output<string?> PostBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// provider id, is always google, unless you have a extended custom implementation
        /// </summary>
        [Output("providerId")]
        public Output<string?> ProviderId { get; private set; } = null!;

        /// <summary>
        /// Realm Name
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// Set 'access_type' query parameter to 'offline' when redirecting to google authorization endpoint, to get a refresh token
        /// back. Useful if planning to use Token Exchange to retrieve Google token to access Google APIs when the user is not at
        /// the browser.
        /// </summary>
        [Output("requestRefreshToken")]
        public Output<bool?> RequestRefreshToken { get; private set; } = null!;

        /// <summary>
        /// Enable/disable if tokens must be stored after authenticating users.
        /// </summary>
        [Output("storeToken")]
        public Output<bool?> StoreToken { get; private set; } = null!;

        /// <summary>
        /// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        /// </summary>
        [Output("trustEmail")]
        public Output<bool?> TrustEmail { get; private set; } = null!;

        /// <summary>
        /// Set 'userIp' query parameter when invoking on Google's User Info service. This will use the user's ip address. Useful if
        /// Google is throttling access to the User Info service.
        /// </summary>
        [Output("useUserIpParam")]
        public Output<bool?> UseUserIpParam { get; private set; } = null!;


        /// <summary>
        /// Create a GoogleIdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GoogleIdentityProvider(string name, GoogleIdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, args ?? new GoogleIdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GoogleIdentityProvider(string name, Input<string> id, GoogleIdentityProviderState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GoogleIdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GoogleIdentityProvider Get(string name, Input<string> id, GoogleIdentityProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new GoogleIdentityProvider(name, id, state, options);
        }
    }

    public sealed class GoogleIdentityProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In
        /// case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly
        /// returned to client, but the request with prompt=none will be forwarded to this identity provider.
        /// </summary>
        [Input("acceptsPromptNoneForwardFromClient")]
        public Input<bool>? AcceptsPromptNoneForwardFromClient { get; set; }

        /// <summary>
        /// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

        /// <summary>
        /// Client ID.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Client Secret.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The scopes to be sent when asking for authorization. See the documentation for possible values, separator and default
        /// value'. Default: 'openid profile email'
        /// </summary>
        [Input("defaultScopes")]
        public Input<string>? DefaultScopes { get; set; }

        /// <summary>
        /// Disable usage of User Info service to obtain additional user information? Default is to use this OIDC service.
        /// </summary>
        [Input("disableUserInfo")]
        public Input<bool>? DisableUserInfo { get; set; }

        /// <summary>
        /// Enable/disable this identity provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
        /// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        /// </summary>
        [Input("firstBrokerLoginFlowAlias")]
        public Input<string>? FirstBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// Hide On Login Page.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// Set 'hd' query parameter when logging in with Google. Google will list accounts only for this domain. Keycloak validates
        /// that the returned identity token has a claim for this domain. When '*' is entered, any hosted account can be used.
        /// </summary>
        [Input("hostedDomain")]
        public Input<string>? HostedDomain { get; set; }

        /// <summary>
        /// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
        /// want to allow login from the provider, but want to integrate with a provider
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
        /// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
        /// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
        /// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// provider id, is always google, unless you have a extended custom implementation
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// Set 'access_type' query parameter to 'offline' when redirecting to google authorization endpoint, to get a refresh token
        /// back. Useful if planning to use Token Exchange to retrieve Google token to access Google APIs when the user is not at
        /// the browser.
        /// </summary>
        [Input("requestRefreshToken")]
        public Input<bool>? RequestRefreshToken { get; set; }

        /// <summary>
        /// Enable/disable if tokens must be stored after authenticating users.
        /// </summary>
        [Input("storeToken")]
        public Input<bool>? StoreToken { get; set; }

        /// <summary>
        /// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Set 'userIp' query parameter when invoking on Google's User Info service. This will use the user's ip address. Useful if
        /// Google is throttling access to the User Info service.
        /// </summary>
        [Input("useUserIpParam")]
        public Input<bool>? UseUserIpParam { get; set; }

        public GoogleIdentityProviderArgs()
        {
        }
    }

    public sealed class GoogleIdentityProviderState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In
        /// case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly
        /// returned to client, but the request with prompt=none will be forwarded to this identity provider.
        /// </summary>
        [Input("acceptsPromptNoneForwardFromClient")]
        public Input<bool>? AcceptsPromptNoneForwardFromClient { get; set; }

        /// <summary>
        /// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// The alias uniquely identifies an identity provider and it is also used to build the redirect uri. In case of google this
        /// is computed and always google
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

        /// <summary>
        /// Client ID.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client Secret.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The scopes to be sent when asking for authorization. See the documentation for possible values, separator and default
        /// value'. Default: 'openid profile email'
        /// </summary>
        [Input("defaultScopes")]
        public Input<string>? DefaultScopes { get; set; }

        /// <summary>
        /// Disable usage of User Info service to obtain additional user information? Default is to use this OIDC service.
        /// </summary>
        [Input("disableUserInfo")]
        public Input<bool>? DisableUserInfo { get; set; }

        /// <summary>
        /// Not used by this provider, Will be implicitly Google
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Enable/disable this identity provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
        /// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        /// </summary>
        [Input("firstBrokerLoginFlowAlias")]
        public Input<string>? FirstBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// Hide On Login Page.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// Set 'hd' query parameter when logging in with Google. Google will list accounts only for this domain. Keycloak validates
        /// that the returned identity token has a claim for this domain. When '*' is entered, any hosted account can be used.
        /// </summary>
        [Input("hostedDomain")]
        public Input<string>? HostedDomain { get; set; }

        /// <summary>
        /// Internal Identity Provider Id
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
        /// want to allow login from the provider, but want to integrate with a provider
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
        /// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
        /// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
        /// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// provider id, is always google, unless you have a extended custom implementation
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Set 'access_type' query parameter to 'offline' when redirecting to google authorization endpoint, to get a refresh token
        /// back. Useful if planning to use Token Exchange to retrieve Google token to access Google APIs when the user is not at
        /// the browser.
        /// </summary>
        [Input("requestRefreshToken")]
        public Input<bool>? RequestRefreshToken { get; set; }

        /// <summary>
        /// Enable/disable if tokens must be stored after authenticating users.
        /// </summary>
        [Input("storeToken")]
        public Input<bool>? StoreToken { get; set; }

        /// <summary>
        /// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Set 'userIp' query parameter when invoking on Google's User Info service. This will use the user's ip address. Useful if
        /// Google is throttling access to the User Info service.
        /// </summary>
        [Input("useUserIpParam")]
        public Input<bool>? UseUserIpParam { get; set; }

        public GoogleIdentityProviderState()
        {
        }
    }
}
