# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetClientAuthorizationPolicyResult',
    'AwaitableGetClientAuthorizationPolicyResult',
    'get_client_authorization_policy',
]

@pulumi.output_type
class GetClientAuthorizationPolicyResult:
    """
    A collection of values returned by getClientAuthorizationPolicy.
    """
    def __init__(__self__, decision_strategy=None, id=None, logic=None, name=None, owner=None, policies=None, realm_id=None, resource_server_id=None, resources=None, scopes=None, type=None):
        if decision_strategy and not isinstance(decision_strategy, str):
            raise TypeError("Expected argument 'decision_strategy' to be a str")
        pulumi.set(__self__, "decision_strategy", decision_strategy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logic and not isinstance(logic, str):
            raise TypeError("Expected argument 'logic' to be a str")
        pulumi.set(__self__, "logic", logic)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if resource_server_id and not isinstance(resource_server_id, str):
            raise TypeError("Expected argument 'resource_server_id' to be a str")
        pulumi.set(__self__, "resource_server_id", resource_server_id)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> str:
        """
        (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        """
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def logic(self) -> str:
        """
        (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
        """
        return pulumi.get(self, "logic")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        (Computed) The ID of the owning resource. Applies to resources.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def policies(self) -> Sequence[str]:
        """
        (Computed) The IDs of the policies that must be applied to scopes/resources for this policy/permission. Applies to policies and permissions.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> str:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter
    def resources(self) -> Sequence[str]:
        """
        (Computed) The IDs of the resources that this permission applies to. Applies to resource-based permissions.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        (Computed) The IDs of the scopes that this permission applies to. Applies to scope-based permissions.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        (Computed) The type of this policy / permission. For permissions, this could be `resource` or `scope`. For policies, this could be any type of authorization policy, such as `js`.
        """
        return pulumi.get(self, "type")


class AwaitableGetClientAuthorizationPolicyResult(GetClientAuthorizationPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientAuthorizationPolicyResult(
            decision_strategy=self.decision_strategy,
            id=self.id,
            logic=self.logic,
            name=self.name,
            owner=self.owner,
            policies=self.policies,
            realm_id=self.realm_id,
            resource_server_id=self.resource_server_id,
            resources=self.resources,
            scopes=self.scopes,
            type=self.type)


def get_client_authorization_policy(name: Optional[str] = None,
                                    realm_id: Optional[str] = None,
                                    resource_server_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientAuthorizationPolicyResult:
    """
    This data source can be used to fetch policy and permission information for an OpenID client that has authorization enabled.

    ## Example Usage

    In this example, we'll create a new OpenID client with authorization enabled. This will cause Keycloak to create a default
    permission for this client called "Default Permission". We'll use the `openid.getClientAuthorizationPolicy` data
    source to fetch information about this permission, so we can use it to create a new resource-based authorization permission.

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    client_with_authz = keycloak.openid.Client("clientWithAuthz",
        client_id="client-with-authz",
        realm_id=realm.id,
        access_type="CONFIDENTIAL",
        service_accounts_enabled=True,
        authorization=keycloak.openid.ClientAuthorizationArgs(
            policy_enforcement_mode="ENFORCING",
        ))
    default_permission = client_with_authz.resource_server_id.apply(lambda resource_server_id: keycloak.openid.get_client_authorization_policy(realm_id=keycloak_realm["test"]["id"],
        resource_server_id=resource_server_id,
        name="Default Permission"))
    resource = keycloak.openid.ClientAuthorizationResource("resource",
        resource_server_id=client_with_authz.resource_server_id,
        realm_id=keycloak_realm["test"]["id"],
        uris=["/endpoint/*"],
        attributes={
            "foo": "bar",
        })
    permission = keycloak.openid.ClientAuthorizationPermission("permission",
        resource_server_id=client_with_authz.resource_server_id,
        realm_id=keycloak_realm["test"]["id"],
        policies=[default_permission.id],
        resources=[resource.id])
    ```


    :param str name: The name of the authorization policy.
    :param str realm_id: The realm this authorization policy exists within.
    :param str resource_server_id: The ID of the resource server this authorization policy is attached to.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['realmId'] = realm_id
    __args__['resourceServerId'] = resource_server_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:openid/getClientAuthorizationPolicy:getClientAuthorizationPolicy', __args__, opts=opts, typ=GetClientAuthorizationPolicyResult).value

    return AwaitableGetClientAuthorizationPolicyResult(
        decision_strategy=__ret__.decision_strategy,
        id=__ret__.id,
        logic=__ret__.logic,
        name=__ret__.name,
        owner=__ret__.owner,
        policies=__ret__.policies,
        realm_id=__ret__.realm_id,
        resource_server_id=__ret__.resource_server_id,
        resources=__ret__.resources,
        scopes=__ret__.scopes,
        type=__ret__.type)
