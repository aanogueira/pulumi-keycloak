# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class User(pulumi.CustomResource):
    attributes: pulumi.Output[dict]
    email: pulumi.Output[str]
    enabled: pulumi.Output[bool]
    federated_identities: pulumi.Output[list]
    first_name: pulumi.Output[str]
    initial_password: pulumi.Output[dict]
    last_name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    username: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, attributes=None, email=None, enabled=None, federated_identities=None, first_name=None, initial_password=None, last_name=None, realm_id=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # .User

        Allows for creating and managing Users within Keycloak.

        This resource was created primarily to enable the acceptance tests for the `.Group` resource.
        Creating users within Keycloak is not recommended. Instead, users should be federated from external sources
        by configuring user federation providers or identity providers.

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this user belongs to.
        - `username` - (Required) The unique username of this user.
        - `initial_password` (Optional) When given, the user's initial password will be set.
           This attribute is only respected during initial user creation.
            - `value` (Required) The initial password.
            - `temporary` (Optional) If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
        - `enabled` - (Optional) When false, this user cannot log in. Defaults to `true`.
        - `email` - (Optional) The user's email.
        - `first_name` - (Optional) The user's first name.
        - `last_name` - (Optional) The user's last name.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **federated_identities** object supports the following:

          * `identityProvider` (`pulumi.Input[str]`)
          * `userId` (`pulumi.Input[str]`)
          * `userName` (`pulumi.Input[str]`)

        The **initial_password** object supports the following:

          * `temporary` (`pulumi.Input[bool]`)
          * `value` (`pulumi.Input[str]`)
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

            __props__['attributes'] = attributes
            __props__['email'] = email
            __props__['enabled'] = enabled
            __props__['federated_identities'] = federated_identities
            __props__['first_name'] = first_name
            __props__['initial_password'] = initial_password
            __props__['last_name'] = last_name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(User, __self__).__init__(
            'keycloak:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, attributes=None, email=None, enabled=None, federated_identities=None, first_name=None, initial_password=None, last_name=None, realm_id=None, username=None):
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **federated_identities** object supports the following:

          * `identityProvider` (`pulumi.Input[str]`)
          * `userId` (`pulumi.Input[str]`)
          * `userName` (`pulumi.Input[str]`)

        The **initial_password** object supports the following:

          * `temporary` (`pulumi.Input[bool]`)
          * `value` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["email"] = email
        __props__["enabled"] = enabled
        __props__["federated_identities"] = federated_identities
        __props__["first_name"] = first_name
        __props__["initial_password"] = initial_password
        __props__["last_name"] = last_name
        __props__["realm_id"] = realm_id
        __props__["username"] = username
        return User(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

