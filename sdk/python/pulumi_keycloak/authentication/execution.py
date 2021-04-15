# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ExecutionArgs', 'Execution']

@pulumi.input_type
class ExecutionArgs:
    def __init__(__self__, *,
                 authenticator: pulumi.Input[str],
                 parent_flow_alias: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 requirement: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Execution resource.
        :param pulumi.Input[str] authenticator: The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        :param pulumi.Input[str] parent_flow_alias: The alias of the flow this execution is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        :param pulumi.Input[str] requirement: The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        pulumi.set(__self__, "authenticator", authenticator)
        pulumi.set(__self__, "parent_flow_alias", parent_flow_alias)
        pulumi.set(__self__, "realm_id", realm_id)
        if requirement is not None:
            pulumi.set(__self__, "requirement", requirement)

    @property
    @pulumi.getter
    def authenticator(self) -> pulumi.Input[str]:
        """
        The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        """
        return pulumi.get(self, "authenticator")

    @authenticator.setter
    def authenticator(self, value: pulumi.Input[str]):
        pulumi.set(self, "authenticator", value)

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> pulumi.Input[str]:
        """
        The alias of the flow this execution is attached to.
        """
        return pulumi.get(self, "parent_flow_alias")

    @parent_flow_alias.setter
    def parent_flow_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_flow_alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm the authentication execution exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def requirement(self) -> Optional[pulumi.Input[str]]:
        """
        The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        return pulumi.get(self, "requirement")

    @requirement.setter
    def requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "requirement", value)


@pulumi.input_type
class _ExecutionState:
    def __init__(__self__, *,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Execution resources.
        :param pulumi.Input[str] authenticator: The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        :param pulumi.Input[str] parent_flow_alias: The alias of the flow this execution is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        :param pulumi.Input[str] requirement: The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        if authenticator is not None:
            pulumi.set(__self__, "authenticator", authenticator)
        if parent_flow_alias is not None:
            pulumi.set(__self__, "parent_flow_alias", parent_flow_alias)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if requirement is not None:
            pulumi.set(__self__, "requirement", requirement)

    @property
    @pulumi.getter
    def authenticator(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        """
        return pulumi.get(self, "authenticator")

    @authenticator.setter
    def authenticator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticator", value)

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of the flow this execution is attached to.
        """
        return pulumi.get(self, "parent_flow_alias")

    @parent_flow_alias.setter
    def parent_flow_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_flow_alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm the authentication execution exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def requirement(self) -> Optional[pulumi.Input[str]]:
        """
        The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        return pulumi.get(self, "requirement")

    @requirement.setter
    def requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "requirement", value)


class Execution(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows for creating and managing an authentication execution within Keycloak.

        An authentication execution is an action that the user or service may or may not take when authenticating through an authentication
        flow.

        > Due to limitations in the Keycloak API, the ordering of authentication executions within a flow must be specified using `depends_on`. Authentication executions that are created first will appear first within the flow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        flow = keycloak.authentication.Flow("flow",
            realm_id=realm.id,
            alias="my-flow-alias")
        # first execution
        execution_one = keycloak.authentication.Execution("executionOne",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="auth-cookie",
            requirement="ALTERNATIVE")
        # second execution
        execution_two = keycloak.authentication.Execution("executionTwo",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector",
            requirement="ALTERNATIVE",
            opts=pulumi.ResourceOptions(depends_on=[execution_one]))
        ```

        ## Import

        Authentication executions can be imported using the formats`{{realmId}}/{{parentFlowAlias}}/{{authenticationExecutionId}}`. Examplebash

        ```sh
         $ pulumi import keycloak:authentication/execution:Execution keycloak_authentication_execution my-realm/my-flow/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticator: The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        :param pulumi.Input[str] parent_flow_alias: The alias of the flow this execution is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        :param pulumi.Input[str] requirement: The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExecutionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing an authentication execution within Keycloak.

        An authentication execution is an action that the user or service may or may not take when authenticating through an authentication
        flow.

        > Due to limitations in the Keycloak API, the ordering of authentication executions within a flow must be specified using `depends_on`. Authentication executions that are created first will appear first within the flow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        flow = keycloak.authentication.Flow("flow",
            realm_id=realm.id,
            alias="my-flow-alias")
        # first execution
        execution_one = keycloak.authentication.Execution("executionOne",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="auth-cookie",
            requirement="ALTERNATIVE")
        # second execution
        execution_two = keycloak.authentication.Execution("executionTwo",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector",
            requirement="ALTERNATIVE",
            opts=pulumi.ResourceOptions(depends_on=[execution_one]))
        ```

        ## Import

        Authentication executions can be imported using the formats`{{realmId}}/{{parentFlowAlias}}/{{authenticationExecutionId}}`. Examplebash

        ```sh
         $ pulumi import keycloak:authentication/execution:Execution keycloak_authentication_execution my-realm/my-flow/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
        ```

        :param str resource_name: The name of the resource.
        :param ExecutionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExecutionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = ExecutionArgs.__new__(ExecutionArgs)

            if authenticator is None and not opts.urn:
                raise TypeError("Missing required property 'authenticator'")
            __props__.__dict__["authenticator"] = authenticator
            if parent_flow_alias is None and not opts.urn:
                raise TypeError("Missing required property 'parent_flow_alias'")
            __props__.__dict__["parent_flow_alias"] = parent_flow_alias
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["requirement"] = requirement
        super(Execution, __self__).__init__(
            'keycloak:authentication/execution:Execution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authenticator: Optional[pulumi.Input[str]] = None,
            parent_flow_alias: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            requirement: Optional[pulumi.Input[str]] = None) -> 'Execution':
        """
        Get an existing Execution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticator: The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        :param pulumi.Input[str] parent_flow_alias: The alias of the flow this execution is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        :param pulumi.Input[str] requirement: The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExecutionState.__new__(_ExecutionState)

        __props__.__dict__["authenticator"] = authenticator
        __props__.__dict__["parent_flow_alias"] = parent_flow_alias
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["requirement"] = requirement
        return Execution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authenticator(self) -> pulumi.Output[str]:
        """
        The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
        """
        return pulumi.get(self, "authenticator")

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> pulumi.Output[str]:
        """
        The alias of the flow this execution is attached to.
        """
        return pulumi.get(self, "parent_flow_alias")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm the authentication execution exists in.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def requirement(self) -> pulumi.Output[Optional[str]]:
        """
        The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
        """
        return pulumi.get(self, "requirement")

