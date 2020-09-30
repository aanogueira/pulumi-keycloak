// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// Allows for creating and managing Keycloak clients that use the OpenID Connect protocol.
    /// 
    /// Clients are entities that can use Keycloak for user authentication. Typically,
    /// clients are applications that redirect users to Keycloak for authentication
    /// in order to take advantage of Keycloak's user sessions for SSO.
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
    ///         var openidClient = new Keycloak.OpenId.Client("openidClient", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = "test-client",
    ///             Enabled = true,
    ///             AccessType = "CONFIDENTIAL",
    ///             ValidRedirectUris = 
    ///             {
    ///                 "http://localhost:8080/openid-callback",
    ///             },
    ///             LoginTheme = "keycloak",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Client : Pulumi.CustomResource
    {
        /// <summary>
        /// The amount of time in seconds before an access token expires. This will override the default for the realm.
        /// </summary>
        [Output("accessTokenLifespan")]
        public Output<string?> AccessTokenLifespan { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of client, which can be one of the following:
        /// </summary>
        [Output("accessType")]
        public Output<string> AccessType { get; private set; } = null!;

        /// <summary>
        /// URL to the admin interface of the client.
        /// </summary>
        [Output("adminUrl")]
        public Output<string?> AdminUrl { get; private set; } = null!;

        /// <summary>
        /// Override realm authentication flow bindings
        /// </summary>
        [Output("authenticationFlowBindingOverrides")]
        public Output<Outputs.ClientAuthenticationFlowBindingOverrides?> AuthenticationFlowBindingOverrides { get; private set; } = null!;

        /// <summary>
        /// When this block is present, fine-grained authorization will be enabled for this client. The client's `access_type` must be `CONFIDENTIAL`, and `service_accounts_enabled` must be `true`. This block has the following arguments:
        /// </summary>
        [Output("authorization")]
        public Output<Outputs.ClientAuthorization?> Authorization { get; private set; } = null!;

        /// <summary>
        /// Default URL to use when the auth server needs to redirect or link back to the client.
        /// </summary>
        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The secret for clients with an `access_type` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// When `true`, users have to consent to client access.
        /// </summary>
        [Output("consentRequired")]
        public Output<bool?> ConsentRequired { get; private set; } = null!;

        /// <summary>
        /// The description of this client in the GUI.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Output("directAccessGrantsEnabled")]
        public Output<bool?> DirectAccessGrantsEnabled { get; private set; } = null!;

        /// <summary>
        /// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// When `true`, the parameter `session_state` will not be included in OpenID Connect Authentication Response.
        /// </summary>
        [Output("excludeSessionStateFromAuthResponse")]
        public Output<bool?> ExcludeSessionStateFromAuthResponse { get; private set; } = null!;

        /// <summary>
        /// Allow to include all roles mappings in the access token.
        /// </summary>
        [Output("fullScopeAllowed")]
        public Output<bool?> FullScopeAllowed { get; private set; } = null!;

        /// <summary>
        /// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Output("implicitFlowEnabled")]
        public Output<bool?> ImplicitFlowEnabled { get; private set; } = null!;

        /// <summary>
        /// The client login theme. This will override the default theme for the realm.
        /// </summary>
        [Output("loginTheme")]
        public Output<string?> LoginTheme { get; private set; } = null!;

        /// <summary>
        /// The display name of this client in the GUI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
        /// </summary>
        [Output("pkceCodeChallengeMethod")]
        public Output<string?> PkceCodeChallengeMethod { get; private set; } = null!;

        /// <summary>
        /// The realm this client is attached to.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// (Computed) When authorization is enabled for this client, this attribute is the unique ID for the client (the same value as the `.id` attribute).
        /// </summary>
        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;

        /// <summary>
        /// When specified, this URL is prepended to any relative URLs found within `valid_redirect_uris`, `web_origins`, and `admin_url`. NOTE: Due to limitations in the Keycloak API, when the `root_url` attribute is used, the `valid_redirect_uris`, `web_origins`, and `admin_url` attributes will be required.
        /// </summary>
        [Output("rootUrl")]
        public Output<string?> RootUrl { get; private set; } = null!;

        /// <summary>
        /// (Computed) When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
        /// </summary>
        [Output("serviceAccountUserId")]
        public Output<string> ServiceAccountUserId { get; private set; } = null!;

        /// <summary>
        /// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Output("serviceAccountsEnabled")]
        public Output<bool?> ServiceAccountsEnabled { get; private set; } = null!;

        /// <summary>
        /// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Output("standardFlowEnabled")]
        public Output<bool?> StandardFlowEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
        /// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standard_flow_enabled` or `implicit_flow_enabled`
        /// is set to `true`.
        /// </summary>
        [Output("validRedirectUris")]
        public Output<ImmutableArray<string>> ValidRedirectUris { get; private set; } = null!;

        /// <summary>
        /// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
        /// </summary>
        [Output("webOrigins")]
        public Output<ImmutableArray<string>> WebOrigins { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/client:Client", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time in seconds before an access token expires. This will override the default for the realm.
        /// </summary>
        [Input("accessTokenLifespan")]
        public Input<string>? AccessTokenLifespan { get; set; }

        /// <summary>
        /// Specifies the type of client, which can be one of the following:
        /// </summary>
        [Input("accessType", required: true)]
        public Input<string> AccessType { get; set; } = null!;

        /// <summary>
        /// URL to the admin interface of the client.
        /// </summary>
        [Input("adminUrl")]
        public Input<string>? AdminUrl { get; set; }

        /// <summary>
        /// Override realm authentication flow bindings
        /// </summary>
        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesArgs>? AuthenticationFlowBindingOverrides { get; set; }

        /// <summary>
        /// When this block is present, fine-grained authorization will be enabled for this client. The client's `access_type` must be `CONFIDENTIAL`, and `service_accounts_enabled` must be `true`. This block has the following arguments:
        /// </summary>
        [Input("authorization")]
        public Input<Inputs.ClientAuthorizationArgs>? Authorization { get; set; }

        /// <summary>
        /// Default URL to use when the auth server needs to redirect or link back to the client.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The secret for clients with an `access_type` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// When `true`, users have to consent to client access.
        /// </summary>
        [Input("consentRequired")]
        public Input<bool>? ConsentRequired { get; set; }

        /// <summary>
        /// The description of this client in the GUI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("directAccessGrantsEnabled")]
        public Input<bool>? DirectAccessGrantsEnabled { get; set; }

        /// <summary>
        /// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When `true`, the parameter `session_state` will not be included in OpenID Connect Authentication Response.
        /// </summary>
        [Input("excludeSessionStateFromAuthResponse")]
        public Input<bool>? ExcludeSessionStateFromAuthResponse { get; set; }

        /// <summary>
        /// Allow to include all roles mappings in the access token.
        /// </summary>
        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("implicitFlowEnabled")]
        public Input<bool>? ImplicitFlowEnabled { get; set; }

        /// <summary>
        /// The client login theme. This will override the default theme for the realm.
        /// </summary>
        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        /// <summary>
        /// The display name of this client in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
        /// </summary>
        [Input("pkceCodeChallengeMethod")]
        public Input<string>? PkceCodeChallengeMethod { get; set; }

        /// <summary>
        /// The realm this client is attached to.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// When specified, this URL is prepended to any relative URLs found within `valid_redirect_uris`, `web_origins`, and `admin_url`. NOTE: Due to limitations in the Keycloak API, when the `root_url` attribute is used, the `valid_redirect_uris`, `web_origins`, and `admin_url` attributes will be required.
        /// </summary>
        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("serviceAccountsEnabled")]
        public Input<bool>? ServiceAccountsEnabled { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("standardFlowEnabled")]
        public Input<bool>? StandardFlowEnabled { get; set; }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;

        /// <summary>
        /// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
        /// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standard_flow_enabled` or `implicit_flow_enabled`
        /// is set to `true`.
        /// </summary>
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        [Input("webOrigins")]
        private InputList<string>? _webOrigins;

        /// <summary>
        /// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
        /// </summary>
        public InputList<string> WebOrigins
        {
            get => _webOrigins ?? (_webOrigins = new InputList<string>());
            set => _webOrigins = value;
        }

        public ClientArgs()
        {
        }
    }

    public sealed class ClientState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time in seconds before an access token expires. This will override the default for the realm.
        /// </summary>
        [Input("accessTokenLifespan")]
        public Input<string>? AccessTokenLifespan { get; set; }

        /// <summary>
        /// Specifies the type of client, which can be one of the following:
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// URL to the admin interface of the client.
        /// </summary>
        [Input("adminUrl")]
        public Input<string>? AdminUrl { get; set; }

        /// <summary>
        /// Override realm authentication flow bindings
        /// </summary>
        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesGetArgs>? AuthenticationFlowBindingOverrides { get; set; }

        /// <summary>
        /// When this block is present, fine-grained authorization will be enabled for this client. The client's `access_type` must be `CONFIDENTIAL`, and `service_accounts_enabled` must be `true`. This block has the following arguments:
        /// </summary>
        [Input("authorization")]
        public Input<Inputs.ClientAuthorizationGetArgs>? Authorization { get; set; }

        /// <summary>
        /// Default URL to use when the auth server needs to redirect or link back to the client.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// The Client ID for this client, referenced in the URI during authentication and in issued tokens.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The secret for clients with an `access_type` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and should be treated with the same care as a password. If omitted, this will be generated by Keycloak.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// When `true`, users have to consent to client access.
        /// </summary>
        [Input("consentRequired")]
        public Input<bool>? ConsentRequired { get; set; }

        /// <summary>
        /// The description of this client in the GUI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("directAccessGrantsEnabled")]
        public Input<bool>? DirectAccessGrantsEnabled { get; set; }

        /// <summary>
        /// When `false`, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When `true`, the parameter `session_state` will not be included in OpenID Connect Authentication Response.
        /// </summary>
        [Input("excludeSessionStateFromAuthResponse")]
        public Input<bool>? ExcludeSessionStateFromAuthResponse { get; set; }

        /// <summary>
        /// Allow to include all roles mappings in the access token.
        /// </summary>
        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("implicitFlowEnabled")]
        public Input<bool>? ImplicitFlowEnabled { get; set; }

        /// <summary>
        /// The client login theme. This will override the default theme for the realm.
        /// </summary>
        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        /// <summary>
        /// The display name of this client in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
        /// </summary>
        [Input("pkceCodeChallengeMethod")]
        public Input<string>? PkceCodeChallengeMethod { get; set; }

        /// <summary>
        /// The realm this client is attached to.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// (Computed) When authorization is enabled for this client, this attribute is the unique ID for the client (the same value as the `.id` attribute).
        /// </summary>
        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        /// <summary>
        /// When specified, this URL is prepended to any relative URLs found within `valid_redirect_uris`, `web_origins`, and `admin_url`. NOTE: Due to limitations in the Keycloak API, when the `root_url` attribute is used, the `valid_redirect_uris`, `web_origins`, and `admin_url` attributes will be required.
        /// </summary>
        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        /// <summary>
        /// (Computed) When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
        /// </summary>
        [Input("serviceAccountUserId")]
        public Input<string>? ServiceAccountUserId { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("serviceAccountsEnabled")]
        public Input<bool>? ServiceAccountsEnabled { get; set; }

        /// <summary>
        /// When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
        /// </summary>
        [Input("standardFlowEnabled")]
        public Input<bool>? StandardFlowEnabled { get; set; }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;

        /// <summary>
        /// A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
        /// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standard_flow_enabled` or `implicit_flow_enabled`
        /// is set to `true`.
        /// </summary>
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        [Input("webOrigins")]
        private InputList<string>? _webOrigins;

        /// <summary>
        /// A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
        /// </summary>
        public InputList<string> WebOrigins
        {
            get => _webOrigins ?? (_webOrigins = new InputList<string>());
            set => _webOrigins = value;
        }

        public ClientState()
        {
        }
    }
}
