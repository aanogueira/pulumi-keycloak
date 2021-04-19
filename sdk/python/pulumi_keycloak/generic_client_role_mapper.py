# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GenericClientRoleMapperArgs', 'GenericClientRoleMapper']

@pulumi.input_type
class GenericClientRoleMapperArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 role_id: pulumi.Input[str],
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GenericClientRoleMapper resource.
        :param pulumi.Input[str] realm_id: The realm this role mapper exists within.
        :param pulumi.Input[str] role_id: The ID of the role to be added to this role mapper.
        :param pulumi.Input[str] client_id: The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        :param pulumi.Input[str] client_scope_id: The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "role_id", role_id)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm this role mapper exists within.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        The ID of the role to be added to this role mapper.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        """
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)


@pulumi.input_type
class _GenericClientRoleMapperState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GenericClientRoleMapper resources.
        :param pulumi.Input[str] client_id: The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        :param pulumi.Input[str] client_scope_id: The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        :param pulumi.Input[str] realm_id: The realm this role mapper exists within.
        :param pulumi.Input[str] role_id: The ID of the role to be added to this role mapper.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        """
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm this role mapper exists within.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the role to be added to this role mapper.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)


class GenericClientRoleMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allow for creating and managing a client's scope mappings within Keycloak.

        By default, all the user role mappings of the user are added as claims within the token (OIDC) or assertion (SAML). When
        `full_scope_allowed` is set to `false` for a client, role scope mapping allows you to limit the roles that get declared
        inside an access token for a client.

        ## Example Usage
        ### Realm Role To Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="BEARER-ONLY")
        realm_role = keycloak.Role("realmRole",
            realm_id=realm.id,
            description="My Realm Role")
        client_role_mapper = keycloak.GenericClientRoleMapper("clientRoleMapper",
            realm_id=realm.id,
            client_id=client.id,
            role_id=realm_role.id)
        ```
        ### Client Role To Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_a = keycloak.openid.Client("clientA",
            realm_id=realm.id,
            client_id="client-a",
            enabled=True,
            access_type="BEARER-ONLY",
            full_scope_allowed=False)
        client_role_a = keycloak.Role("clientRoleA",
            realm_id=realm.id,
            client_id=client_a.id,
            description="My Client Role")
        client_b = keycloak.openid.Client("clientB",
            realm_id=realm.id,
            client_id="client-b",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role_b = keycloak.Role("clientRoleB",
            realm_id=realm.id,
            client_id=client_b.id,
            description="My Client Role")
        client_b_role_mapper = keycloak.GenericClientRoleMapper("clientBRoleMapper",
            realm_id=realm.id,
            client_id=keycloak_client["client_b"]["id"],
            role_id=client_role_a.id)
        ```
        ### Realm Role To Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        realm_role = keycloak.Role("realmRole",
            realm_id=realm.id,
            description="My Realm Role")
        client_role_mapper = keycloak.GenericClientRoleMapper("clientRoleMapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            role_id=realm_role.id)
        ```
        ### Client Role To Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role = keycloak.Role("clientRole",
            realm_id=realm.id,
            client_id=client.id,
            description="My Client Role")
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        client_b_role_mapper = keycloak.GenericClientRoleMapper("clientBRoleMapper",
            realm_id=realm.id,
            client_scope_id=keycloak_client_scope["client_scope"]["id"],
            role_id=client_role.id)
        ```

        ## Import

        Generic client role mappers can be imported using one of the following two formats- When mapping a role to a client, use the format `{{realmId}}/client/{{clientId}}/scope-mappings/{{roleClientId}}/{{roleId}}` - When mapping a role to a client scope, use the format `{{realmId}}/client-scope/{{clientScopeId}}/scope-mappings/{{roleClientId}}/{{roleId}}` Examplebash

        ```sh
         $ pulumi import keycloak:index/genericClientRoleMapper:GenericClientRoleMapper client_role_mapper my-realm/client/23888550-5dcd-41f6-85ba-554233021e9c/scope-mappings/ce51f004-bdfb-4dd5-a963-c4487d2dec5b/ff3aa49f-bc07-4030-8783-41918c3614a3
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        :param pulumi.Input[str] client_scope_id: The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        :param pulumi.Input[str] realm_id: The realm this role mapper exists within.
        :param pulumi.Input[str] role_id: The ID of the role to be added to this role mapper.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GenericClientRoleMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allow for creating and managing a client's scope mappings within Keycloak.

        By default, all the user role mappings of the user are added as claims within the token (OIDC) or assertion (SAML). When
        `full_scope_allowed` is set to `false` for a client, role scope mapping allows you to limit the roles that get declared
        inside an access token for a client.

        ## Example Usage
        ### Realm Role To Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="BEARER-ONLY")
        realm_role = keycloak.Role("realmRole",
            realm_id=realm.id,
            description="My Realm Role")
        client_role_mapper = keycloak.GenericClientRoleMapper("clientRoleMapper",
            realm_id=realm.id,
            client_id=client.id,
            role_id=realm_role.id)
        ```
        ### Client Role To Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_a = keycloak.openid.Client("clientA",
            realm_id=realm.id,
            client_id="client-a",
            enabled=True,
            access_type="BEARER-ONLY",
            full_scope_allowed=False)
        client_role_a = keycloak.Role("clientRoleA",
            realm_id=realm.id,
            client_id=client_a.id,
            description="My Client Role")
        client_b = keycloak.openid.Client("clientB",
            realm_id=realm.id,
            client_id="client-b",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role_b = keycloak.Role("clientRoleB",
            realm_id=realm.id,
            client_id=client_b.id,
            description="My Client Role")
        client_b_role_mapper = keycloak.GenericClientRoleMapper("clientBRoleMapper",
            realm_id=realm.id,
            client_id=keycloak_client["client_b"]["id"],
            role_id=client_role_a.id)
        ```
        ### Realm Role To Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        realm_role = keycloak.Role("realmRole",
            realm_id=realm.id,
            description="My Realm Role")
        client_role_mapper = keycloak.GenericClientRoleMapper("clientRoleMapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            role_id=realm_role.id)
        ```
        ### Client Role To Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role = keycloak.Role("clientRole",
            realm_id=realm.id,
            client_id=client.id,
            description="My Client Role")
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        client_b_role_mapper = keycloak.GenericClientRoleMapper("clientBRoleMapper",
            realm_id=realm.id,
            client_scope_id=keycloak_client_scope["client_scope"]["id"],
            role_id=client_role.id)
        ```

        ## Import

        Generic client role mappers can be imported using one of the following two formats- When mapping a role to a client, use the format `{{realmId}}/client/{{clientId}}/scope-mappings/{{roleClientId}}/{{roleId}}` - When mapping a role to a client scope, use the format `{{realmId}}/client-scope/{{clientScopeId}}/scope-mappings/{{roleClientId}}/{{roleId}}` Examplebash

        ```sh
         $ pulumi import keycloak:index/genericClientRoleMapper:GenericClientRoleMapper client_role_mapper my-realm/client/23888550-5dcd-41f6-85ba-554233021e9c/scope-mappings/ce51f004-bdfb-4dd5-a963-c4487d2dec5b/ff3aa49f-bc07-4030-8783-41918c3614a3
        ```

        :param str resource_name: The name of the resource.
        :param GenericClientRoleMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GenericClientRoleMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GenericClientRoleMapperArgs.__new__(GenericClientRoleMapperArgs)

            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_scope_id"] = client_scope_id
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
        super(GenericClientRoleMapper, __self__).__init__(
            'keycloak:index/genericClientRoleMapper:GenericClientRoleMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None) -> 'GenericClientRoleMapper':
        """
        Get an existing GenericClientRoleMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        :param pulumi.Input[str] client_scope_id: The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        :param pulumi.Input[str] realm_id: The realm this role mapper exists within.
        :param pulumi.Input[str] role_id: The ID of the role to be added to this role mapper.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GenericClientRoleMapperState.__new__(_GenericClientRoleMapperState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_scope_id"] = client_scope_id
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["role_id"] = role_id
        return GenericClientRoleMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the client this role mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the client scope this role mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        """
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this role mapper exists within.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The ID of the role to be added to this role mapper.
        """
        return pulumi.get(self, "role_id")

