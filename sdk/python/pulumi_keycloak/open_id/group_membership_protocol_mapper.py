# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GroupMembershipProtocolMapper(pulumi.CustomResource):
    add_to_access_token: pulumi.Output[bool]
    add_to_id_token: pulumi.Output[bool]
    add_to_userinfo: pulumi.Output[bool]
    claim_name: pulumi.Output[str]
    client_id: pulumi.Output[str]
    client_scope_id: pulumi.Output[str]
    full_path: pulumi.Output[bool]
    name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, claim_name=None, client_id=None, client_scope_id=None, full_path=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # OpenId.GroupMembershipProtocolMapper
        
        Allows for creating and managing group membership protocol mappers within
        Keycloak.
        
        Group membership protocol mappers allow you to map a user's group memberships
        to a claim in a token. Protocol mappers can be defined for a single client,
        or they can be defined for a client scope which can be shared between multiple
        different clients.
        
        ### Argument Reference
        
        The following arguments are supported:
        
        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `claim_name` - (Required) The name of the claim to insert into a token.
        - `full_path` - (Optional) Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
        - `add_to_id_token` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        - `add_to_access_token` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        - `add_to_userinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_group_membership_protocol_mapper.html.markdown.
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

            __props__['add_to_access_token'] = add_to_access_token
            __props__['add_to_id_token'] = add_to_id_token
            __props__['add_to_userinfo'] = add_to_userinfo
            if claim_name is None:
                raise TypeError("Missing required property 'claim_name'")
            __props__['claim_name'] = claim_name
            __props__['client_id'] = client_id
            __props__['client_scope_id'] = client_scope_id
            __props__['full_path'] = full_path
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(GroupMembershipProtocolMapper, __self__).__init__(
            'keycloak:OpenId/groupMembershipProtocolMapper:GroupMembershipProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, claim_name=None, client_id=None, client_scope_id=None, full_path=None, name=None, realm_id=None):
        """
        Get an existing GroupMembershipProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_group_membership_protocol_mapper.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["add_to_access_token"] = add_to_access_token
        __props__["add_to_id_token"] = add_to_id_token
        __props__["add_to_userinfo"] = add_to_userinfo
        __props__["claim_name"] = claim_name
        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["full_path"] = full_path
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return GroupMembershipProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

