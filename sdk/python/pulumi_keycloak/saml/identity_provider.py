# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IdentityProvider(pulumi.CustomResource):
    add_read_token_role_on_create: pulumi.Output[bool]
    """
    Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
    """
    alias: pulumi.Output[str]
    """
    The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
    """
    authenticate_by_default: pulumi.Output[bool]
    """
    Enable/disable authenticate users by default.
    """
    backchannel_supported: pulumi.Output[bool]
    """
    Does the external IDP support backchannel logout?
    """
    display_name: pulumi.Output[str]
    """
    Friendly name for Identity Providers.
    """
    enabled: pulumi.Output[bool]
    """
    Enable/disable this identity provider.
    """
    first_broker_login_flow_alias: pulumi.Output[str]
    """
    Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
    that there is not yet existing Keycloak account linked with the authenticated identity provider account.
    """
    force_authn: pulumi.Output[bool]
    """
    Require Force Authn.
    """
    hide_on_login_page: pulumi.Output[bool]
    """
    Hide On Login Page.
    """
    internal_id: pulumi.Output[str]
    """
    Internal Identity Provider Id
    """
    link_only: pulumi.Output[bool]
    """
    If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
    want to allow login from the provider, but want to integrate with a provider
    """
    name_id_policy_format: pulumi.Output[str]
    """
    Name ID Policy Format.
    """
    post_binding_authn_request: pulumi.Output[bool]
    """
    Post Binding Authn Request.
    """
    post_binding_logout: pulumi.Output[bool]
    """
    Post Binding Logout.
    """
    post_binding_response: pulumi.Output[bool]
    """
    Post Binding Response.
    """
    post_broker_login_flow_alias: pulumi.Output[str]
    """
    Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
    additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
    you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
    authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
    """
    realm: pulumi.Output[str]
    """
    Realm Name
    """
    signature_algorithm: pulumi.Output[str]
    """
    Signing Algorithm.
    """
    signing_certificate: pulumi.Output[str]
    """
    Signing Certificate.
    """
    single_logout_service_url: pulumi.Output[str]
    """
    Logout URL.
    """
    single_sign_on_service_url: pulumi.Output[str]
    """
    SSO Logout URL.
    """
    store_token: pulumi.Output[bool]
    """
    Enable/disable if tokens must be stored after authenticating users.
    """
    trust_email: pulumi.Output[bool]
    """
    If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
    """
    validate_signature: pulumi.Output[bool]
    """
    Enable/disable signature validation of SAML responses.
    """
    want_assertions_encrypted: pulumi.Output[bool]
    """
    Want Assertions Encrypted.
    """
    want_assertions_signed: pulumi.Output[bool]
    """
    Want Assertions Signed.
    """
    xml_sign_key_info_key_name_transformer: pulumi.Output[str]
    """
    Sign Key Transformer.
    """
    def __init__(__self__, resource_name, opts=None, add_read_token_role_on_create=None, alias=None, authenticate_by_default=None, backchannel_supported=None, display_name=None, enabled=None, first_broker_login_flow_alias=None, force_authn=None, hide_on_login_page=None, link_only=None, name_id_policy_format=None, post_binding_authn_request=None, post_binding_logout=None, post_binding_response=None, post_broker_login_flow_alias=None, realm=None, signature_algorithm=None, signing_certificate=None, single_logout_service_url=None, single_sign_on_service_url=None, store_token=None, trust_email=None, validate_signature=None, want_assertions_encrypted=None, want_assertions_signed=None, xml_sign_key_info_key_name_transformer=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # saml.IdentityProvider

        Allows to create and manage SAML Identity Providers within Keycloak.

        SAML (Security Assertion Markup Language) identity providers allows to authenticate through a third-party system, using SAML standard.

        ### Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm_identity_provider = keycloak.saml.IdentityProvider("realmIdentityProvider",
            alias="my-idp",
            backchannel_supported=True,
            force_authn=True,
            post_binding_authn_request=True,
            post_binding_logout=True,
            post_binding_response=True,
            realm="my-realm",
            single_logout_service_url="https://domain.com/adfs/ls/?wa=wsignout1.0",
            single_sign_on_service_url="https://domain.com/adfs/ls/",
            store_token=False,
            trust_email=True)
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm` - (Required) The name of the realm. This is unique across Keycloak.
        - `alias` - (Optional) The uniq name of identity provider.
        - `enabled` - (Optional) When false, users and clients will not be able to access this realm. Defaults to `true`.
        - `display_name` - (Optional) The display name for the realm that is shown when logging in to the admin console.
        - `store_token` - (Optional) Enable/disable if tokens must be stored after authenticating users. Defaults to `true`.
        - `add_read_token_role_on_create` - (Optional) Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role. Defaults to `false`.
        - `trust_email` - (Optional) If enabled then email provided by this provider is not verified even if verification is enabled for the realm. Defaults to `false`.
        - `link_only` - (Optional) If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't want to allow login from the provider, but want to integrate with a provider. Defaults to `false`.
        - `hide_on_login_page` - (Optional) If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
        - `first_broker_login_flow_alias` - (Optional) Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
        - `post_broker_login_flow_alias` - (Optional) Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
        - `authenticate_by_default` - (Optional) Authenticate users by default. Defaults to `false`.

        #### SAML Configuration

        - `single_sign_on_service_url` - (Optional) The Url that must be used to send authentication requests (SAML AuthnRequest).
        - `single_logout_service_url` - (Optional) The Url that must be used to send logout requests.
        - `backchannel_supported` - (Optional) Does the external IDP support back-channel logout ?.
        - `name_id_policy_format` - (Optional) Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
        - `post_binding_response` - (Optional) Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
        - `post_binding_authn_request` - (Optional) Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        - `post_binding_logout` - (Optional) Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        - `want_assertions_signed` - (Optional) Indicates whether this service provider expects a signed Assertion.
        - `want_assertions_encrypted` - (Optional) Indicates whether this service provider expects an encrypted Assertion.
        - `force_authn` - (Optional) Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
        - `validate_signature` - (Optional) Enable/disable signature validation of SAML responses.
        - `signing_certificate` - (Optional) Signing Certificate.
        - `signature_algorithm` - (Optional) Signing Algorithm. Defaults to empty.
        - `xml_sign_key_info_key_name_transformer` - (Optional) Sign Key Transformer. Defaults to empty.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_read_token_role_on_create: Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        :param pulumi.Input[str] alias: The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        :param pulumi.Input[bool] authenticate_by_default: Enable/disable authenticate users by default.
        :param pulumi.Input[bool] backchannel_supported: Does the external IDP support backchannel logout?
        :param pulumi.Input[str] display_name: Friendly name for Identity Providers.
        :param pulumi.Input[bool] enabled: Enable/disable this identity provider.
        :param pulumi.Input[str] first_broker_login_flow_alias: Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
               that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        :param pulumi.Input[bool] force_authn: Require Force Authn.
        :param pulumi.Input[bool] hide_on_login_page: Hide On Login Page.
        :param pulumi.Input[bool] link_only: If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
               want to allow login from the provider, but want to integrate with a provider
        :param pulumi.Input[str] name_id_policy_format: Name ID Policy Format.
        :param pulumi.Input[bool] post_binding_authn_request: Post Binding Authn Request.
        :param pulumi.Input[bool] post_binding_logout: Post Binding Logout.
        :param pulumi.Input[bool] post_binding_response: Post Binding Response.
        :param pulumi.Input[str] post_broker_login_flow_alias: Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
               additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
               you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
               authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[str] signature_algorithm: Signing Algorithm.
        :param pulumi.Input[str] signing_certificate: Signing Certificate.
        :param pulumi.Input[str] single_logout_service_url: Logout URL.
        :param pulumi.Input[str] single_sign_on_service_url: SSO Logout URL.
        :param pulumi.Input[bool] store_token: Enable/disable if tokens must be stored after authenticating users.
        :param pulumi.Input[bool] trust_email: If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        :param pulumi.Input[bool] validate_signature: Enable/disable signature validation of SAML responses.
        :param pulumi.Input[bool] want_assertions_encrypted: Want Assertions Encrypted.
        :param pulumi.Input[bool] want_assertions_signed: Want Assertions Signed.
        :param pulumi.Input[str] xml_sign_key_info_key_name_transformer: Sign Key Transformer.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['add_read_token_role_on_create'] = add_read_token_role_on_create
            if alias is None:
                raise TypeError("Missing required property 'alias'")
            __props__['alias'] = alias
            __props__['authenticate_by_default'] = authenticate_by_default
            __props__['backchannel_supported'] = backchannel_supported
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            __props__['first_broker_login_flow_alias'] = first_broker_login_flow_alias
            __props__['force_authn'] = force_authn
            __props__['hide_on_login_page'] = hide_on_login_page
            __props__['link_only'] = link_only
            __props__['name_id_policy_format'] = name_id_policy_format
            __props__['post_binding_authn_request'] = post_binding_authn_request
            __props__['post_binding_logout'] = post_binding_logout
            __props__['post_binding_response'] = post_binding_response
            __props__['post_broker_login_flow_alias'] = post_broker_login_flow_alias
            if realm is None:
                raise TypeError("Missing required property 'realm'")
            __props__['realm'] = realm
            __props__['signature_algorithm'] = signature_algorithm
            __props__['signing_certificate'] = signing_certificate
            __props__['single_logout_service_url'] = single_logout_service_url
            if single_sign_on_service_url is None:
                raise TypeError("Missing required property 'single_sign_on_service_url'")
            __props__['single_sign_on_service_url'] = single_sign_on_service_url
            __props__['store_token'] = store_token
            __props__['trust_email'] = trust_email
            __props__['validate_signature'] = validate_signature
            __props__['want_assertions_encrypted'] = want_assertions_encrypted
            __props__['want_assertions_signed'] = want_assertions_signed
            __props__['xml_sign_key_info_key_name_transformer'] = xml_sign_key_info_key_name_transformer
            __props__['internal_id'] = None
        super(IdentityProvider, __self__).__init__(
            'keycloak:saml/identityProvider:IdentityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_read_token_role_on_create=None, alias=None, authenticate_by_default=None, backchannel_supported=None, display_name=None, enabled=None, first_broker_login_flow_alias=None, force_authn=None, hide_on_login_page=None, internal_id=None, link_only=None, name_id_policy_format=None, post_binding_authn_request=None, post_binding_logout=None, post_binding_response=None, post_broker_login_flow_alias=None, realm=None, signature_algorithm=None, signing_certificate=None, single_logout_service_url=None, single_sign_on_service_url=None, store_token=None, trust_email=None, validate_signature=None, want_assertions_encrypted=None, want_assertions_signed=None, xml_sign_key_info_key_name_transformer=None):
        """
        Get an existing IdentityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_read_token_role_on_create: Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        :param pulumi.Input[str] alias: The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        :param pulumi.Input[bool] authenticate_by_default: Enable/disable authenticate users by default.
        :param pulumi.Input[bool] backchannel_supported: Does the external IDP support backchannel logout?
        :param pulumi.Input[str] display_name: Friendly name for Identity Providers.
        :param pulumi.Input[bool] enabled: Enable/disable this identity provider.
        :param pulumi.Input[str] first_broker_login_flow_alias: Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
               that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        :param pulumi.Input[bool] force_authn: Require Force Authn.
        :param pulumi.Input[bool] hide_on_login_page: Hide On Login Page.
        :param pulumi.Input[str] internal_id: Internal Identity Provider Id
        :param pulumi.Input[bool] link_only: If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
               want to allow login from the provider, but want to integrate with a provider
        :param pulumi.Input[str] name_id_policy_format: Name ID Policy Format.
        :param pulumi.Input[bool] post_binding_authn_request: Post Binding Authn Request.
        :param pulumi.Input[bool] post_binding_logout: Post Binding Logout.
        :param pulumi.Input[bool] post_binding_response: Post Binding Response.
        :param pulumi.Input[str] post_broker_login_flow_alias: Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
               additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
               you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
               authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[str] signature_algorithm: Signing Algorithm.
        :param pulumi.Input[str] signing_certificate: Signing Certificate.
        :param pulumi.Input[str] single_logout_service_url: Logout URL.
        :param pulumi.Input[str] single_sign_on_service_url: SSO Logout URL.
        :param pulumi.Input[bool] store_token: Enable/disable if tokens must be stored after authenticating users.
        :param pulumi.Input[bool] trust_email: If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        :param pulumi.Input[bool] validate_signature: Enable/disable signature validation of SAML responses.
        :param pulumi.Input[bool] want_assertions_encrypted: Want Assertions Encrypted.
        :param pulumi.Input[bool] want_assertions_signed: Want Assertions Signed.
        :param pulumi.Input[str] xml_sign_key_info_key_name_transformer: Sign Key Transformer.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_read_token_role_on_create"] = add_read_token_role_on_create
        __props__["alias"] = alias
        __props__["authenticate_by_default"] = authenticate_by_default
        __props__["backchannel_supported"] = backchannel_supported
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["first_broker_login_flow_alias"] = first_broker_login_flow_alias
        __props__["force_authn"] = force_authn
        __props__["hide_on_login_page"] = hide_on_login_page
        __props__["internal_id"] = internal_id
        __props__["link_only"] = link_only
        __props__["name_id_policy_format"] = name_id_policy_format
        __props__["post_binding_authn_request"] = post_binding_authn_request
        __props__["post_binding_logout"] = post_binding_logout
        __props__["post_binding_response"] = post_binding_response
        __props__["post_broker_login_flow_alias"] = post_broker_login_flow_alias
        __props__["realm"] = realm
        __props__["signature_algorithm"] = signature_algorithm
        __props__["signing_certificate"] = signing_certificate
        __props__["single_logout_service_url"] = single_logout_service_url
        __props__["single_sign_on_service_url"] = single_sign_on_service_url
        __props__["store_token"] = store_token
        __props__["trust_email"] = trust_email
        __props__["validate_signature"] = validate_signature
        __props__["want_assertions_encrypted"] = want_assertions_encrypted
        __props__["want_assertions_signed"] = want_assertions_signed
        __props__["xml_sign_key_info_key_name_transformer"] = xml_sign_key_info_key_name_transformer
        return IdentityProvider(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

