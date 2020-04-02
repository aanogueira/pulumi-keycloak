# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClientResult:
    """
    A collection of values returned by getClient.
    """
    def __init__(__self__, access_type=None, authorization=None, client_id=None, client_secret=None, description=None, direct_access_grants_enabled=None, enabled=None, full_scope_allowed=None, id=None, implicit_flow_enabled=None, name=None, realm_id=None, resource_server_id=None, service_account_user_id=None, service_accounts_enabled=None, standard_flow_enabled=None, valid_redirect_uris=None, web_origins=None):
        if access_type and not isinstance(access_type, str):
            raise TypeError("Expected argument 'access_type' to be a str")
        __self__.access_type = access_type
        if authorization and not isinstance(authorization, dict):
            raise TypeError("Expected argument 'authorization' to be a dict")
        __self__.authorization = authorization
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        __self__.client_id = client_id
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        __self__.client_secret = client_secret
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if direct_access_grants_enabled and not isinstance(direct_access_grants_enabled, bool):
            raise TypeError("Expected argument 'direct_access_grants_enabled' to be a bool")
        __self__.direct_access_grants_enabled = direct_access_grants_enabled
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        if full_scope_allowed and not isinstance(full_scope_allowed, bool):
            raise TypeError("Expected argument 'full_scope_allowed' to be a bool")
        __self__.full_scope_allowed = full_scope_allowed
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if implicit_flow_enabled and not isinstance(implicit_flow_enabled, bool):
            raise TypeError("Expected argument 'implicit_flow_enabled' to be a bool")
        __self__.implicit_flow_enabled = implicit_flow_enabled
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        __self__.realm_id = realm_id
        if resource_server_id and not isinstance(resource_server_id, str):
            raise TypeError("Expected argument 'resource_server_id' to be a str")
        __self__.resource_server_id = resource_server_id
        if service_account_user_id and not isinstance(service_account_user_id, str):
            raise TypeError("Expected argument 'service_account_user_id' to be a str")
        __self__.service_account_user_id = service_account_user_id
        if service_accounts_enabled and not isinstance(service_accounts_enabled, bool):
            raise TypeError("Expected argument 'service_accounts_enabled' to be a bool")
        __self__.service_accounts_enabled = service_accounts_enabled
        if standard_flow_enabled and not isinstance(standard_flow_enabled, bool):
            raise TypeError("Expected argument 'standard_flow_enabled' to be a bool")
        __self__.standard_flow_enabled = standard_flow_enabled
        if valid_redirect_uris and not isinstance(valid_redirect_uris, list):
            raise TypeError("Expected argument 'valid_redirect_uris' to be a list")
        __self__.valid_redirect_uris = valid_redirect_uris
        if web_origins and not isinstance(web_origins, list):
            raise TypeError("Expected argument 'web_origins' to be a list")
        __self__.web_origins = web_origins
class AwaitableGetClientResult(GetClientResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientResult(
            access_type=self.access_type,
            authorization=self.authorization,
            client_id=self.client_id,
            client_secret=self.client_secret,
            description=self.description,
            direct_access_grants_enabled=self.direct_access_grants_enabled,
            enabled=self.enabled,
            full_scope_allowed=self.full_scope_allowed,
            id=self.id,
            implicit_flow_enabled=self.implicit_flow_enabled,
            name=self.name,
            realm_id=self.realm_id,
            resource_server_id=self.resource_server_id,
            service_account_user_id=self.service_account_user_id,
            service_accounts_enabled=self.service_accounts_enabled,
            standard_flow_enabled=self.standard_flow_enabled,
            valid_redirect_uris=self.valid_redirect_uris,
            web_origins=self.web_origins)

def get_client(client_id=None,realm_id=None,opts=None):
    """
    ## # openid.Client data source

    This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.

    ### Argument Reference

    The following arguments are supported:

    - `realm_id` - (Required) The realm id.
    - `client_id` - (Required) The client id.

    ### Attributes Reference

    See the docs for the `openid.Client` resource for details on the exported attributes.

    > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/d/keycloak_openid_client.html.markdown.
    """
    __args__ = dict()


    __args__['clientId'] = client_id
    __args__['realmId'] = realm_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:openid/getClient:getClient', __args__, opts=opts).value

    return AwaitableGetClientResult(
        access_type=__ret__.get('accessType'),
        authorization=__ret__.get('authorization'),
        client_id=__ret__.get('clientId'),
        client_secret=__ret__.get('clientSecret'),
        description=__ret__.get('description'),
        direct_access_grants_enabled=__ret__.get('directAccessGrantsEnabled'),
        enabled=__ret__.get('enabled'),
        full_scope_allowed=__ret__.get('fullScopeAllowed'),
        id=__ret__.get('id'),
        implicit_flow_enabled=__ret__.get('implicitFlowEnabled'),
        name=__ret__.get('name'),
        realm_id=__ret__.get('realmId'),
        resource_server_id=__ret__.get('resourceServerId'),
        service_account_user_id=__ret__.get('serviceAccountUserId'),
        service_accounts_enabled=__ret__.get('serviceAccountsEnabled'),
        standard_flow_enabled=__ret__.get('standardFlowEnabled'),
        valid_redirect_uris=__ret__.get('validRedirectUris'),
        web_origins=__ret__.get('webOrigins'))