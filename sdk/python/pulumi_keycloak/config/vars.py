# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'base_path',
    'client_id',
    'client_secret',
    'client_timeout',
    'initial_login',
    'password',
    'realm',
    'root_ca_certificate',
    'tls_insecure_skip_verify',
    'url',
    'username',
]

__config__ = pulumi.Config('keycloak')

base_path = __config__.get('basePath')

client_id = __config__.get('clientId')

client_secret = __config__.get('clientSecret')

client_timeout = __config__.get('clientTimeout') or (_utilities.get_env_int('KEYCLOAK_CLIENT_TIMEOUT') or 5)
"""
Timeout (in seconds) of the Keycloak client
"""

initial_login = __config__.get('initialLogin')
"""
Whether or not to login to Keycloak instance on provider initialization
"""

password = __config__.get('password')

realm = __config__.get('realm')

root_ca_certificate = __config__.get('rootCaCertificate')
"""
Allows x509 calls using an unknown CA certificate (for development purposes)
"""

tls_insecure_skip_verify = __config__.get('tlsInsecureSkipVerify')
"""
Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
should be avoided.
"""

url = __config__.get('url')
"""
The base URL of the Keycloak instance, before `/auth`
"""

username = __config__.get('username')

