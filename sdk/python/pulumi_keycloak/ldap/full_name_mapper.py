# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FullNameMapper(pulumi.CustomResource):
    ldap_full_name_attribute: pulumi.Output[str]
    ldap_user_federation_id: pulumi.Output[str]
    name: pulumi.Output[str]
    read_only: pulumi.Output[bool]
    realm_id: pulumi.Output[str]
    write_only: pulumi.Output[bool]
    def __init__(__self__, resource_name, opts=None, ldap_full_name_attribute=None, ldap_user_federation_id=None, name=None, read_only=None, realm_id=None, write_only=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # Ldap.FullNameMapper
        
        Allows for creating and managing full name mappers for Keycloak users federated
        via LDAP.
        
        The LDAP full name mapper can map a user's full name from an LDAP attribute
        to the first and last name attributes of a Keycloak user.
        
        ### Argument Reference
        
        The following arguments are supported:
        
        - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
        - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
        - `name` - (Required) Display name of this mapper when displayed in the console.
        - `ldap_full_name_attribute` - (Required) The name of the LDAP attribute containing the user's full name.
        - `read_only` - (Optional) When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
        - `write_only` - (Optional) When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/ldap_full_name_mapper.html.markdown.
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

            if ldap_full_name_attribute is None:
                raise TypeError("Missing required property 'ldap_full_name_attribute'")
            __props__['ldap_full_name_attribute'] = ldap_full_name_attribute
            if ldap_user_federation_id is None:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__['ldap_user_federation_id'] = ldap_user_federation_id
            __props__['name'] = name
            __props__['read_only'] = read_only
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            __props__['write_only'] = write_only
        super(FullNameMapper, __self__).__init__(
            'keycloak:Ldap/fullNameMapper:FullNameMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ldap_full_name_attribute=None, ldap_user_federation_id=None, name=None, read_only=None, realm_id=None, write_only=None):
        """
        Get an existing FullNameMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/ldap_full_name_mapper.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["ldap_full_name_attribute"] = ldap_full_name_attribute
        __props__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__["name"] = name
        __props__["read_only"] = read_only
        __props__["realm_id"] = realm_id
        __props__["write_only"] = write_only
        return FullNameMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

