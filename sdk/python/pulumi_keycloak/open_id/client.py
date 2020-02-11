# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Client(pulumi.CustomResource):
    access_type: pulumi.Output[str]
    authorization: pulumi.Output[dict]
    client_id: pulumi.Output[str]
    client_secret: pulumi.Output[str]
    description: pulumi.Output[str]
    direct_access_grants_enabled: pulumi.Output[bool]
    enabled: pulumi.Output[bool]
    exclude_session_state_from_auth_response: pulumi.Output[bool]
    full_scope_allowed: pulumi.Output[bool]
    implicit_flow_enabled: pulumi.Output[bool]
    name: pulumi.Output[str]
    pkce_code_challenge_method: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    resource_server_id: pulumi.Output[str]
    service_account_user_id: pulumi.Output[str]
    service_accounts_enabled: pulumi.Output[bool]
    standard_flow_enabled: pulumi.Output[bool]
    valid_redirect_uris: pulumi.Output[list]
    web_origins: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, access_type=None, authorization=None, client_id=None, client_secret=None, description=None, direct_access_grants_enabled=None, enabled=None, exclude_session_state_from_auth_response=None, full_scope_allowed=None, implicit_flow_enabled=None, name=None, pkce_code_challenge_method=None, realm_id=None, service_accounts_enabled=None, standard_flow_enabled=None, valid_redirect_uris=None, web_origins=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # OpenId.Client
        
        Allows for creating and managing Keycloak clients that use the OpenID Connect protocol.
        
        Clients are entities that can use Keycloak for user authentication. Typically,
        clients are applications that redirect users to Keycloak for authentication
        in order to take advantage of Keycloak's user sessions for SSO.
        
        ### Argument Reference
        
        The following arguments are supported:
        
        - `realm_id` - (Required) The realm this client is attached to.
        - `client_id` - (Required) The unique ID of this client, referenced in the URI during authentication and in issued tokens.
        - `name` - (Optional) The display name of this client in the GUI.
        - `enabled` - (Optional) When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        - `description` - (Optional) The description of this client in the GUI.
        - `access_type` - (Required) Specifies the type of client, which can be one of the following:
            - `CONFIDENTIAL` - Used for server-side clients that require both client ID and secret when authenticating.
              This client should be used for applications using the Authorization Code or Client Credentials grant flows.
            - `PUBLIC` - Used for browser-only applications that do not require a client secret, and instead rely only on authorized redirect
              URIs for security. This client should be used for applications using the Implicit grant flow.
            - `BEARER-ONLY` - Used for services that never initiate a login. This client will only allow bearer token requests.
        - `client_secret` - (Optional) The secret for clients with an `access_type` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and
        should be treated with the same care as a password. If omitted, Keycloak will generate a GUID for this attribute.
        - `standard_flow_enabled` - (Optional) When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
        - `implicit_flow_enabled` - (Optional) When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
        - `direct_access_grants_enabled` - (Optional) When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
        - `service_accounts_enabled` - (Optional) When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
        - `valid_redirect_uris` - (Optional) A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
        wildcards in the form of an asterisk can be used here. This attribute must be set if either `standard_flow_enabled` or `implicit_flow_enabled`
        is set to `true`.
        - `web_origins` - (Optional) A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
        - `admin_url` - (Optional) URL to the admin interface of the client.
        - `base_url` - (Optional) Default URL to use when the auth server needs to redirect or link back to the client.
        - `pkce_code_challenge_method` - (Optional) The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
        - `full_scope_allowed` - (Optional) - Allow to include all roles mappings in the access token.
        
        ### Attributes Reference
        
        In addition to the arguments listed above, the following computed attributes are exported:
        
        - `service_account_user_id` - When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **authorization** object supports the following:
        
          * `allowRemoteResourceManagement` (`pulumi.Input[bool]`)
          * `keepDefaults` (`pulumi.Input[bool]`)
          * `policyEnforcementMode` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_client.html.markdown.
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

            if access_type is None:
                raise TypeError("Missing required property 'access_type'")
            __props__['access_type'] = access_type
            __props__['authorization'] = authorization
            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['description'] = description
            __props__['direct_access_grants_enabled'] = direct_access_grants_enabled
            __props__['enabled'] = enabled
            __props__['exclude_session_state_from_auth_response'] = exclude_session_state_from_auth_response
            __props__['full_scope_allowed'] = full_scope_allowed
            __props__['implicit_flow_enabled'] = implicit_flow_enabled
            __props__['name'] = name
            __props__['pkce_code_challenge_method'] = pkce_code_challenge_method
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            __props__['service_accounts_enabled'] = service_accounts_enabled
            __props__['standard_flow_enabled'] = standard_flow_enabled
            __props__['valid_redirect_uris'] = valid_redirect_uris
            __props__['web_origins'] = web_origins
            __props__['resource_server_id'] = None
            __props__['service_account_user_id'] = None
        super(Client, __self__).__init__(
            'keycloak:OpenId/client:Client',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_type=None, authorization=None, client_id=None, client_secret=None, description=None, direct_access_grants_enabled=None, enabled=None, exclude_session_state_from_auth_response=None, full_scope_allowed=None, implicit_flow_enabled=None, name=None, pkce_code_challenge_method=None, realm_id=None, resource_server_id=None, service_account_user_id=None, service_accounts_enabled=None, standard_flow_enabled=None, valid_redirect_uris=None, web_origins=None):
        """
        Get an existing Client resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **authorization** object supports the following:
        
          * `allowRemoteResourceManagement` (`pulumi.Input[bool]`)
          * `keepDefaults` (`pulumi.Input[bool]`)
          * `policyEnforcementMode` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_client.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access_type"] = access_type
        __props__["authorization"] = authorization
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["description"] = description
        __props__["direct_access_grants_enabled"] = direct_access_grants_enabled
        __props__["enabled"] = enabled
        __props__["exclude_session_state_from_auth_response"] = exclude_session_state_from_auth_response
        __props__["full_scope_allowed"] = full_scope_allowed
        __props__["implicit_flow_enabled"] = implicit_flow_enabled
        __props__["name"] = name
        __props__["pkce_code_challenge_method"] = pkce_code_challenge_method
        __props__["realm_id"] = realm_id
        __props__["resource_server_id"] = resource_server_id
        __props__["service_account_user_id"] = service_account_user_id
        __props__["service_accounts_enabled"] = service_accounts_enabled
        __props__["standard_flow_enabled"] = standard_flow_enabled
        __props__["valid_redirect_uris"] = valid_redirect_uris
        __props__["web_origins"] = web_origins
        return Client(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
