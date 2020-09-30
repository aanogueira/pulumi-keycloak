# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['GoogleIdentityProvider']


class GoogleIdentityProvider(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepts_prompt_none_forward_from_client: Optional[pulumi.Input[bool]] = None,
                 add_read_token_role_on_create: Optional[pulumi.Input[bool]] = None,
                 authenticate_by_default: Optional[pulumi.Input[bool]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 default_scopes: Optional[pulumi.Input[str]] = None,
                 disable_user_info: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 first_broker_login_flow_alias: Optional[pulumi.Input[str]] = None,
                 hide_on_login_page: Optional[pulumi.Input[bool]] = None,
                 hosted_domain: Optional[pulumi.Input[str]] = None,
                 link_only: Optional[pulumi.Input[bool]] = None,
                 post_broker_login_flow_alias: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 request_refresh_token: Optional[pulumi.Input[bool]] = None,
                 store_token: Optional[pulumi.Input[bool]] = None,
                 trust_email: Optional[pulumi.Input[bool]] = None,
                 use_user_ip_param: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows for creating and managing OIDC Identity Providers within Keycloak.

        OIDC (OpenID Connect) identity providers allows users to authenticate through a third party system using the OIDC standard.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        google = keycloak.oidc.GoogleIdentityProvider("google",
            realm=realm.id,
            client_id=var["google_identity_provider_client_id"],
            client_secret=var["google_identity_provider_client_secret"],
            trust_email=True,
            hosted_domain="example.com",
            extra_config={
                "syncMode": "IMPORT",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accepts_prompt_none_forward_from_client: When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
        :param pulumi.Input[bool] add_read_token_role_on_create: When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        :param pulumi.Input[bool] authenticate_by_default: Enable/disable authenticate users by default.
        :param pulumi.Input[str] client_id: The client or client identifier registered within the identity provider.
        :param pulumi.Input[str] client_secret: The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
        :param pulumi.Input[str] default_scopes: The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
        :param pulumi.Input[bool] disable_user_info: When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        :param pulumi.Input[bool] enabled: When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
        :param pulumi.Input[str] first_broker_login_flow_alias: The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
        :param pulumi.Input[bool] hide_on_login_page: When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
        :param pulumi.Input[str] hosted_domain: Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
        :param pulumi.Input[bool] link_only: When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        :param pulumi.Input[str] post_broker_login_flow_alias: The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        :param pulumi.Input[str] provider_id: The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
        :param pulumi.Input[str] realm: The name of the realm. This is unique across Keycloak.
        :param pulumi.Input[bool] request_refresh_token: Sets the "access_type" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
        :param pulumi.Input[bool] store_token: When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        :param pulumi.Input[bool] trust_email: When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        :param pulumi.Input[bool] use_user_ip_param: Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['accepts_prompt_none_forward_from_client'] = accepts_prompt_none_forward_from_client
            __props__['add_read_token_role_on_create'] = add_read_token_role_on_create
            __props__['authenticate_by_default'] = authenticate_by_default
            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['default_scopes'] = default_scopes
            __props__['disable_user_info'] = disable_user_info
            __props__['enabled'] = enabled
            __props__['extra_config'] = extra_config
            __props__['first_broker_login_flow_alias'] = first_broker_login_flow_alias
            __props__['hide_on_login_page'] = hide_on_login_page
            __props__['hosted_domain'] = hosted_domain
            __props__['link_only'] = link_only
            __props__['post_broker_login_flow_alias'] = post_broker_login_flow_alias
            __props__['provider_id'] = provider_id
            if realm is None:
                raise TypeError("Missing required property 'realm'")
            __props__['realm'] = realm
            __props__['request_refresh_token'] = request_refresh_token
            __props__['store_token'] = store_token
            __props__['trust_email'] = trust_email
            __props__['use_user_ip_param'] = use_user_ip_param
            __props__['alias'] = None
            __props__['display_name'] = None
            __props__['internal_id'] = None
        super(GoogleIdentityProvider, __self__).__init__(
            'keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accepts_prompt_none_forward_from_client: Optional[pulumi.Input[bool]] = None,
            add_read_token_role_on_create: Optional[pulumi.Input[bool]] = None,
            alias: Optional[pulumi.Input[str]] = None,
            authenticate_by_default: Optional[pulumi.Input[bool]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            default_scopes: Optional[pulumi.Input[str]] = None,
            disable_user_info: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            first_broker_login_flow_alias: Optional[pulumi.Input[str]] = None,
            hide_on_login_page: Optional[pulumi.Input[bool]] = None,
            hosted_domain: Optional[pulumi.Input[str]] = None,
            internal_id: Optional[pulumi.Input[str]] = None,
            link_only: Optional[pulumi.Input[bool]] = None,
            post_broker_login_flow_alias: Optional[pulumi.Input[str]] = None,
            provider_id: Optional[pulumi.Input[str]] = None,
            realm: Optional[pulumi.Input[str]] = None,
            request_refresh_token: Optional[pulumi.Input[bool]] = None,
            store_token: Optional[pulumi.Input[bool]] = None,
            trust_email: Optional[pulumi.Input[bool]] = None,
            use_user_ip_param: Optional[pulumi.Input[bool]] = None) -> 'GoogleIdentityProvider':
        """
        Get an existing GoogleIdentityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accepts_prompt_none_forward_from_client: When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
        :param pulumi.Input[bool] add_read_token_role_on_create: When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        :param pulumi.Input[str] alias: (Computed) The alias for the Google identity provider.
        :param pulumi.Input[bool] authenticate_by_default: Enable/disable authenticate users by default.
        :param pulumi.Input[str] client_id: The client or client identifier registered within the identity provider.
        :param pulumi.Input[str] client_secret: The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
        :param pulumi.Input[str] default_scopes: The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
        :param pulumi.Input[bool] disable_user_info: When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        :param pulumi.Input[str] display_name: (Computed) Display name for the Google identity provider in the GUI.
        :param pulumi.Input[bool] enabled: When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
        :param pulumi.Input[str] first_broker_login_flow_alias: The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
        :param pulumi.Input[bool] hide_on_login_page: When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
        :param pulumi.Input[str] hosted_domain: Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
        :param pulumi.Input[str] internal_id: (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
        :param pulumi.Input[bool] link_only: When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        :param pulumi.Input[str] post_broker_login_flow_alias: The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        :param pulumi.Input[str] provider_id: The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
        :param pulumi.Input[str] realm: The name of the realm. This is unique across Keycloak.
        :param pulumi.Input[bool] request_refresh_token: Sets the "access_type" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
        :param pulumi.Input[bool] store_token: When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        :param pulumi.Input[bool] trust_email: When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        :param pulumi.Input[bool] use_user_ip_param: Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accepts_prompt_none_forward_from_client"] = accepts_prompt_none_forward_from_client
        __props__["add_read_token_role_on_create"] = add_read_token_role_on_create
        __props__["alias"] = alias
        __props__["authenticate_by_default"] = authenticate_by_default
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["default_scopes"] = default_scopes
        __props__["disable_user_info"] = disable_user_info
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["extra_config"] = extra_config
        __props__["first_broker_login_flow_alias"] = first_broker_login_flow_alias
        __props__["hide_on_login_page"] = hide_on_login_page
        __props__["hosted_domain"] = hosted_domain
        __props__["internal_id"] = internal_id
        __props__["link_only"] = link_only
        __props__["post_broker_login_flow_alias"] = post_broker_login_flow_alias
        __props__["provider_id"] = provider_id
        __props__["realm"] = realm
        __props__["request_refresh_token"] = request_refresh_token
        __props__["store_token"] = store_token
        __props__["trust_email"] = trust_email
        __props__["use_user_ip_param"] = use_user_ip_param
        return GoogleIdentityProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptsPromptNoneForwardFromClient")
    def accepts_prompt_none_forward_from_client(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
        """
        return pulumi.get(self, "accepts_prompt_none_forward_from_client")

    @property
    @pulumi.getter(name="addReadTokenRoleOnCreate")
    def add_read_token_role_on_create(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        """
        return pulumi.get(self, "add_read_token_role_on_create")

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        (Computed) The alias for the Google identity provider.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="authenticateByDefault")
    def authenticate_by_default(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable/disable authenticate users by default.
        """
        return pulumi.get(self, "authenticate_by_default")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The client or client identifier registered within the identity provider.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="defaultScopes")
    def default_scopes(self) -> pulumi.Output[Optional[str]]:
        """
        The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
        """
        return pulumi.get(self, "default_scopes")

    @property
    @pulumi.getter(name="disableUserInfo")
    def disable_user_info(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        """
        return pulumi.get(self, "disable_user_info")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        (Computed) Display name for the Google identity provider in the GUI.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "extra_config")

    @property
    @pulumi.getter(name="firstBrokerLoginFlowAlias")
    def first_broker_login_flow_alias(self) -> pulumi.Output[Optional[str]]:
        """
        The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
        """
        return pulumi.get(self, "first_broker_login_flow_alias")

    @property
    @pulumi.getter(name="hideOnLoginPage")
    def hide_on_login_page(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
        """
        return pulumi.get(self, "hide_on_login_page")

    @property
    @pulumi.getter(name="hostedDomain")
    def hosted_domain(self) -> pulumi.Output[Optional[str]]:
        """
        Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
        """
        return pulumi.get(self, "hosted_domain")

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> pulumi.Output[str]:
        """
        (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
        """
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter(name="linkOnly")
    def link_only(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        """
        return pulumi.get(self, "link_only")

    @property
    @pulumi.getter(name="postBrokerLoginFlowAlias")
    def post_broker_login_flow_alias(self) -> pulumi.Output[Optional[str]]:
        """
        The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        """
        return pulumi.get(self, "post_broker_login_flow_alias")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
        """
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        """
        The name of the realm. This is unique across Keycloak.
        """
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter(name="requestRefreshToken")
    def request_refresh_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Sets the "access_type" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
        """
        return pulumi.get(self, "request_refresh_token")

    @property
    @pulumi.getter(name="storeToken")
    def store_token(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        """
        return pulumi.get(self, "store_token")

    @property
    @pulumi.getter(name="trustEmail")
    def trust_email(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        """
        return pulumi.get(self, "trust_email")

    @property
    @pulumi.getter(name="useUserIpParam")
    def use_user_ip_param(self) -> pulumi.Output[Optional[bool]]:
        """
        Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
        """
        return pulumi.get(self, "use_user_ip_param")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

