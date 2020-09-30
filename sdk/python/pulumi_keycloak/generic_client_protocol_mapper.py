# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['GenericClientProtocolMapper']


class GenericClientProtocolMapper(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_mapper: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows for creating and managing protocol mappers for both types of clients (openid-connect and saml) within Keycloak.

        There are two uses cases for using this resource:
        * If you implemented a custom protocol mapper, this resource can be used to configure it
        * If the provider doesn't support a particular protocol mapper, this resource can be used instead.

        Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
        Therefore, if possible, a specific mapper should be used.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        saml_client = keycloak.saml.Client("samlClient",
            realm_id=realm.id,
            client_id="test-client")
        saml_hardcode_attribute_mapper = keycloak.GenericClientProtocolMapper("samlHardcodeAttributeMapper",
            realm_id=realm.id,
            client_id=saml_client.id,
            protocol="saml",
            protocol_mapper="saml-hardcode-attribute-mapper",
            config={
                "attribute.name": "name",
                "attribute.nameformat": "Basic",
                "attribute.value": "value",
                "friendly.name": "display name",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The client this protocol mapper is attached to.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[Mapping[str, Any]] config: A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI.
        :param pulumi.Input[str] protocol: The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        :param pulumi.Input[str] protocol_mapper: The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
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

            __props__['client_id'] = client_id
            __props__['client_scope_id'] = client_scope_id
            if config is None:
                raise TypeError("Missing required property 'config'")
            __props__['config'] = config
            __props__['name'] = name
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if protocol_mapper is None:
                raise TypeError("Missing required property 'protocol_mapper'")
            __props__['protocol_mapper'] = protocol_mapper
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(GenericClientProtocolMapper, __self__).__init__(
            'keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            protocol_mapper: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'GenericClientProtocolMapper':
        """
        Get an existing GenericClientProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The client this protocol mapper is attached to.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[Mapping[str, Any]] config: A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI.
        :param pulumi.Input[str] protocol: The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        :param pulumi.Input[str] protocol_mapper: The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["config"] = config
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["protocol_mapper"] = protocol_mapper
        __props__["realm_id"] = realm_id
        return GenericClientProtocolMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client this protocol mapper is attached to.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The mapper's associated client scope. Cannot be used at the same time as client_id.
        """
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name of this protocol mapper in the GUI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolMapper")
    def protocol_mapper(self) -> pulumi.Output[str]:
        """
        The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
        """
        return pulumi.get(self, "protocol_mapper")

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

