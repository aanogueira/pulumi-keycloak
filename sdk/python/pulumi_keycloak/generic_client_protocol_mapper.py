# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GenericClientProtocolMapper(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    config: pulumi.Output[dict]
    name: pulumi.Output[str]
    protocol: pulumi.Output[str]
    protocol_mapper: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, client_id=None, config=None, name=None, protocol=None, protocol_mapper=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # .GenericClientProtocolMapper
        
        Allows for creating and managing protocol mapper for both types of clients (openid-connect and saml) within Keycloak.
        
        There are two uses cases for using this resource:
        * If you implemented a custom protocol mapper, this resource can be used to configure it
        * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
        
        Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors. 
        Therefore, if possible, a specific mapper should be used.
        
        ### Argument Reference
        
        The following arguments are supported:
        
        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required) The client this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `protocol` - (Required) The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        - `protocol_mapper` - (Required) The name of the protocol mapper. The protocol mapper must be
           compatible with the specified client.
        - `config` - (Required) A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/generic_client_protocol_mapper.html.markdown.
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

            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
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
    def get(resource_name, id, opts=None, client_id=None, config=None, name=None, protocol=None, protocol_mapper=None, realm_id=None):
        """
        Get an existing GenericClientProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/generic_client_protocol_mapper.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["client_id"] = client_id
        __props__["config"] = config
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["protocol_mapper"] = protocol_mapper
        __props__["realm_id"] = realm_id
        return GenericClientProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
