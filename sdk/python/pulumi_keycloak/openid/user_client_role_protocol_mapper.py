# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['UserClientRoleProtocolMapper']


class UserClientRoleProtocolMapper(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 add_to_userinfo: Optional[pulumi.Input[bool]] = None,
                 claim_name: Optional[pulumi.Input[str]] = None,
                 claim_value_type: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_id_for_role_mappings: Optional[pulumi.Input[str]] = None,
                 client_role_prefix: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 multivalued: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows for creating and managing user client role protocol mappers within Keycloak.

        User client role protocol mappers allow you to define a claim containing the list of a client roles.

        Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
        multiple different clients.

        ## Example Usage
        ### Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client = keycloak.openid.Client("openidClient",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="CONFIDENTIAL",
            valid_redirect_uris=["http://localhost:8080/openid-callback"])
        user_client_role_mapper = keycloak.openid.UserClientRoleProtocolMapper("userClientRoleMapper",
            realm_id=realm.id,
            client_id=openid_client.id,
            claim_name="foo")
        ```
        ### Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        user_client_role_mapper = keycloak.openid.UserClientRoleProtocolMapper("userClientRoleMapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            claim_name="foo")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        :param pulumi.Input[str] claim_name: The name of the claim to insert into a token.
        :param pulumi.Input[str] claim_value_type: The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_id_for_role_mappings: The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
        :param pulumi.Input[str] client_role_prefix: A prefix for each Client Role.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[bool] multivalued: Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI.
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
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

            __props__['add_to_access_token'] = add_to_access_token
            __props__['add_to_id_token'] = add_to_id_token
            __props__['add_to_userinfo'] = add_to_userinfo
            if claim_name is None:
                raise TypeError("Missing required property 'claim_name'")
            __props__['claim_name'] = claim_name
            __props__['claim_value_type'] = claim_value_type
            __props__['client_id'] = client_id
            __props__['client_id_for_role_mappings'] = client_id_for_role_mappings
            __props__['client_role_prefix'] = client_role_prefix
            __props__['client_scope_id'] = client_scope_id
            __props__['multivalued'] = multivalued
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(UserClientRoleProtocolMapper, __self__).__init__(
            'keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_to_access_token: Optional[pulumi.Input[bool]] = None,
            add_to_id_token: Optional[pulumi.Input[bool]] = None,
            add_to_userinfo: Optional[pulumi.Input[bool]] = None,
            claim_name: Optional[pulumi.Input[str]] = None,
            claim_value_type: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_id_for_role_mappings: Optional[pulumi.Input[str]] = None,
            client_role_prefix: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            multivalued: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'UserClientRoleProtocolMapper':
        """
        Get an existing UserClientRoleProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        :param pulumi.Input[str] claim_name: The name of the claim to insert into a token.
        :param pulumi.Input[str] claim_value_type: The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_id_for_role_mappings: The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
        :param pulumi.Input[str] client_role_prefix: A prefix for each Client Role.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[bool] multivalued: Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI.
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_to_access_token"] = add_to_access_token
        __props__["add_to_id_token"] = add_to_id_token
        __props__["add_to_userinfo"] = add_to_userinfo
        __props__["claim_name"] = claim_name
        __props__["claim_value_type"] = claim_value_type
        __props__["client_id"] = client_id
        __props__["client_id_for_role_mappings"] = client_id_for_role_mappings
        __props__["client_role_prefix"] = client_role_prefix
        __props__["client_scope_id"] = client_scope_id
        __props__["multivalued"] = multivalued
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return UserClientRoleProtocolMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addToAccessToken")
    def add_to_access_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        """
        return pulumi.get(self, "add_to_access_token")

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        """
        return pulumi.get(self, "add_to_id_token")

    @property
    @pulumi.getter(name="addToUserinfo")
    def add_to_userinfo(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        """
        return pulumi.get(self, "add_to_userinfo")

    @property
    @pulumi.getter(name="claimName")
    def claim_name(self) -> pulumi.Output[str]:
        """
        The name of the claim to insert into a token.
        """
        return pulumi.get(self, "claim_name")

    @property
    @pulumi.getter(name="claimValueType")
    def claim_value_type(self) -> pulumi.Output[Optional[str]]:
        """
        The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
        """
        return pulumi.get(self, "claim_value_type")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientIdForRoleMappings")
    def client_id_for_role_mappings(self) -> pulumi.Output[Optional[str]]:
        """
        The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
        """
        return pulumi.get(self, "client_id_for_role_mappings")

    @property
    @pulumi.getter(name="clientRolePrefix")
    def client_role_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        A prefix for each Client Role.
        """
        return pulumi.get(self, "client_role_prefix")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter
    def multivalued(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
        """
        return pulumi.get(self, "multivalued")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name of this protocol mapper in the GUI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this protocol mapper exists within.
        """
        return pulumi.get(self, "realm_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

