# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ScriptProtocolMapper']


class ScriptProtocolMapper(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name_format: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 single_value_attribute: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows for creating and managing script protocol mappers for SAML clients within Keycloak.

        Script protocol mappers evaluate a JavaScript function to produce an attribute value based on context information.

        Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
        multiple different clients.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        saml_client = keycloak.saml.Client("samlClient",
            realm_id=keycloak_realm["test"]["id"],
            client_id="saml-client")
        saml_script_mapper = keycloak.saml.ScriptProtocolMapper("samlScriptMapper",
            realm_id=keycloak_realm["test"]["id"],
            client_id=saml_client.id,
            script="exports = 'foo';",
            saml_attribute_name="displayName",
            saml_attribute_name_format="Unspecified")
        ```

        ## Import

        Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash

        ```sh
         $ pulumi import keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper saml_script_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
        ```

        ```sh
         $ pulumi import keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper saml_script_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] friendly_name: An optional human-friendly name for this attribute.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI.
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        :param pulumi.Input[str] saml_attribute_name: The name of the SAML attribute.
        :param pulumi.Input[str] saml_attribute_name_format: The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
        :param pulumi.Input[str] script: JavaScript code to compute the attribute value.
        :param pulumi.Input[bool] single_value_attribute: When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
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
            __props__['friendly_name'] = friendly_name
            __props__['name'] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if saml_attribute_name is None and not opts.urn:
                raise TypeError("Missing required property 'saml_attribute_name'")
            __props__['saml_attribute_name'] = saml_attribute_name
            if saml_attribute_name_format is None and not opts.urn:
                raise TypeError("Missing required property 'saml_attribute_name_format'")
            __props__['saml_attribute_name_format'] = saml_attribute_name_format
            if script is None and not opts.urn:
                raise TypeError("Missing required property 'script'")
            __props__['script'] = script
            __props__['single_value_attribute'] = single_value_attribute
        super(ScriptProtocolMapper, __self__).__init__(
            'keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            saml_attribute_name: Optional[pulumi.Input[str]] = None,
            saml_attribute_name_format: Optional[pulumi.Input[str]] = None,
            script: Optional[pulumi.Input[str]] = None,
            single_value_attribute: Optional[pulumi.Input[bool]] = None) -> 'ScriptProtocolMapper':
        """
        Get an existing ScriptProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] friendly_name: An optional human-friendly name for this attribute.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI.
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        :param pulumi.Input[str] saml_attribute_name: The name of the SAML attribute.
        :param pulumi.Input[str] saml_attribute_name_format: The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
        :param pulumi.Input[str] script: JavaScript code to compute the attribute value.
        :param pulumi.Input[bool] single_value_attribute: When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["friendly_name"] = friendly_name
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        __props__["saml_attribute_name"] = saml_attribute_name
        __props__["saml_attribute_name_format"] = saml_attribute_name_format
        __props__["script"] = script
        __props__["single_value_attribute"] = single_value_attribute
        return ScriptProtocolMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        """
        An optional human-friendly name for this attribute.
        """
        return pulumi.get(self, "friendly_name")

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

    @property
    @pulumi.getter(name="samlAttributeName")
    def saml_attribute_name(self) -> pulumi.Output[str]:
        """
        The name of the SAML attribute.
        """
        return pulumi.get(self, "saml_attribute_name")

    @property
    @pulumi.getter(name="samlAttributeNameFormat")
    def saml_attribute_name_format(self) -> pulumi.Output[str]:
        """
        The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
        """
        return pulumi.get(self, "saml_attribute_name_format")

    @property
    @pulumi.getter
    def script(self) -> pulumi.Output[str]:
        """
        JavaScript code to compute the attribute value.
        """
        return pulumi.get(self, "script")

    @property
    @pulumi.getter(name="singleValueAttribute")
    def single_value_attribute(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
        """
        return pulumi.get(self, "single_value_attribute")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
