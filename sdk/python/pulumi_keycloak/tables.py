# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

_SNAKE_TO_CAMEL_CASE_TABLE = {
    "access_code_lifespan": "accessCodeLifespan",
    "access_code_lifespan_login": "accessCodeLifespanLogin",
    "access_code_lifespan_user_action": "accessCodeLifespanUserAction",
    "access_token_lifespan": "accessTokenLifespan",
    "access_token_lifespan_for_implicit_flow": "accessTokenLifespanForImplicitFlow",
    "access_type": "accessType",
    "account_theme": "accountTheme",
    "action_token_generated_by_admin_lifespan": "actionTokenGeneratedByAdminLifespan",
    "action_token_generated_by_user_lifespan": "actionTokenGeneratedByUserLifespan",
    "add_read_token_role_on_create": "addReadTokenRoleOnCreate",
    "add_to_access_token": "addToAccessToken",
    "add_to_id_token": "addToIdToken",
    "add_to_userinfo": "addToUserinfo",
    "admin_theme": "adminTheme",
    "always_read_value_from_ldap": "alwaysReadValueFromLdap",
    "assertion_consumer_post_url": "assertionConsumerPostUrl",
    "assertion_consumer_redirect_url": "assertionConsumerRedirectUrl",
    "attribute_friendly_name": "attributeFriendlyName",
    "attribute_name": "attributeName",
    "attribute_value": "attributeValue",
    "authenticate_by_default": "authenticateByDefault",
    "authorization_url": "authorizationUrl",
    "backchannel_supported": "backchannelSupported",
    "base_url": "baseUrl",
    "batch_size_for_sync": "batchSizeForSync",
    "bind_credential": "bindCredential",
    "bind_dn": "bindDn",
    "browser_flow": "browserFlow",
    "cache_policy": "cachePolicy",
    "changed_sync_period": "changedSyncPeriod",
    "claim_name": "claimName",
    "claim_value": "claimValue",
    "claim_value_type": "claimValueType",
    "client_authentication_flow": "clientAuthenticationFlow",
    "client_id": "clientId",
    "client_scope_id": "clientScopeId",
    "client_secret": "clientSecret",
    "client_signature_required": "clientSignatureRequired",
    "client_timeout": "clientTimeout",
    "composite_roles": "compositeRoles",
    "connection_timeout": "connectionTimeout",
    "connection_url": "connectionUrl",
    "consent_screen_text": "consentScreenText",
    "custom_user_search_filter": "customUserSearchFilter",
    "decision_strategy": "decisionStrategy",
    "default_action": "defaultAction",
    "default_scopes": "defaultScopes",
    "direct_access_grants_enabled": "directAccessGrantsEnabled",
    "direct_grant_flow": "directGrantFlow",
    "display_name": "displayName",
    "docker_authentication_flow": "dockerAuthenticationFlow",
    "drop_non_existing_groups_during_sync": "dropNonExistingGroupsDuringSync",
    "duplicate_emails_allowed": "duplicateEmailsAllowed",
    "edit_mode": "editMode",
    "edit_username_allowed": "editUsernameAllowed",
    "email_theme": "emailTheme",
    "exclude_session_state_from_auth_response": "excludeSessionStateFromAuthResponse",
    "extra_config": "extraConfig",
    "federated_identities": "federatedIdentities",
    "first_broker_login_flow_alias": "firstBrokerLoginFlowAlias",
    "first_name": "firstName",
    "force_authn": "forceAuthn",
    "force_post_binding": "forcePostBinding",
    "friendly_name": "friendlyName",
    "front_channel_logout": "frontChannelLogout",
    "full_path": "fullPath",
    "full_scope_allowed": "fullScopeAllowed",
    "full_sync_period": "fullSyncPeriod",
    "group_id": "groupId",
    "group_ids": "groupIds",
    "group_name_ldap_attribute": "groupNameLdapAttribute",
    "group_object_classes": "groupObjectClasses",
    "groups_ldap_filter": "groupsLdapFilter",
    "hide_on_login_page": "hideOnLoginPage",
    "icon_uri": "iconUri",
    "identity_provider_alias": "identityProviderAlias",
    "idp_initiated_sso_relay_state": "idpInitiatedSsoRelayState",
    "idp_initiated_sso_url_name": "idpInitiatedSsoUrlName",
    "ignore_missing_groups": "ignoreMissingGroups",
    "implicit_flow_enabled": "implicitFlowEnabled",
    "import_enabled": "importEnabled",
    "include_authn_statement": "includeAuthnStatement",
    "included_client_audience": "includedClientAudience",
    "included_custom_audience": "includedCustomAudience",
    "initial_login": "initialLogin",
    "initial_password": "initialPassword",
    "internal_id": "internalId",
    "is_mandatory_in_ldap": "isMandatoryInLdap",
    "jwks_url": "jwksUrl",
    "last_name": "lastName",
    "ldap_attribute": "ldapAttribute",
    "ldap_full_name_attribute": "ldapFullNameAttribute",
    "ldap_groups_dn": "ldapGroupsDn",
    "ldap_password_policy_hints_enabled": "ldapPasswordPolicyHintsEnabled",
    "ldap_user_federation_id": "ldapUserFederationId",
    "link_only": "linkOnly",
    "login_hint": "loginHint",
    "login_theme": "loginTheme",
    "login_with_email_allowed": "loginWithEmailAllowed",
    "logout_service_post_binding_url": "logoutServicePostBindingUrl",
    "logout_service_redirect_binding_url": "logoutServiceRedirectBindingUrl",
    "logout_url": "logoutUrl",
    "mapped_group_attributes": "mappedGroupAttributes",
    "master_saml_processing_url": "masterSamlProcessingUrl",
    "memberof_ldap_attribute": "memberofLdapAttribute",
    "membership_attribute_type": "membershipAttributeType",
    "membership_ldap_attribute": "membershipLdapAttribute",
    "membership_user_ldap_attribute": "membershipUserLdapAttribute",
    "name_id_format": "nameIdFormat",
    "name_id_policy_format": "nameIdPolicyFormat",
    "offline_session_idle_timeout": "offlineSessionIdleTimeout",
    "offline_session_max_lifespan": "offlineSessionMaxLifespan",
    "optional_scopes": "optionalScopes",
    "owner_managed_access": "ownerManagedAccess",
    "parent_id": "parentId",
    "password_policy": "passwordPolicy",
    "pkce_code_challenge_method": "pkceCodeChallengeMethod",
    "post_binding_authn_request": "postBindingAuthnRequest",
    "post_binding_logout": "postBindingLogout",
    "post_binding_response": "postBindingResponse",
    "post_broker_login_flow_alias": "postBrokerLoginFlowAlias",
    "preserve_group_inheritance": "preserveGroupInheritance",
    "protocol_mapper": "protocolMapper",
    "provider_id": "providerId",
    "rdn_ldap_attribute": "rdnLdapAttribute",
    "read_only": "readOnly",
    "read_timeout": "readTimeout",
    "realm_id": "realmId",
    "realm_role_prefix": "realmRolePrefix",
    "refresh_token_max_reuse": "refreshTokenMaxReuse",
    "registration_allowed": "registrationAllowed",
    "registration_email_as_username": "registrationEmailAsUsername",
    "registration_flow": "registrationFlow",
    "remember_me": "rememberMe",
    "reset_credentials_flow": "resetCredentialsFlow",
    "reset_password_allowed": "resetPasswordAllowed",
    "resource_server_id": "resourceServerId",
    "revoke_refresh_token": "revokeRefreshToken",
    "role_id": "roleId",
    "role_ids": "roleIds",
    "root_url": "rootUrl",
    "saml_attribute_name": "samlAttributeName",
    "saml_attribute_name_format": "samlAttributeNameFormat",
    "search_scope": "searchScope",
    "security_defenses": "securityDefenses",
    "service_account_user_id": "serviceAccountUserId",
    "service_accounts_enabled": "serviceAccountsEnabled",
    "sign_assertions": "signAssertions",
    "sign_documents": "signDocuments",
    "signature_algorithm": "signatureAlgorithm",
    "signing_certificate": "signingCertificate",
    "signing_private_key": "signingPrivateKey",
    "single_logout_service_url": "singleLogoutServiceUrl",
    "single_sign_on_service_url": "singleSignOnServiceUrl",
    "smtp_server": "smtpServer",
    "ssl_required": "sslRequired",
    "sso_session_idle_timeout": "ssoSessionIdleTimeout",
    "sso_session_max_lifespan": "ssoSessionMaxLifespan",
    "standard_flow_enabled": "standardFlowEnabled",
    "store_token": "storeToken",
    "sync_registrations": "syncRegistrations",
    "token_url": "tokenUrl",
    "trust_email": "trustEmail",
    "ui_locales": "uiLocales",
    "use_truststore_spi": "useTruststoreSpi",
    "user_attribute": "userAttribute",
    "user_info_url": "userInfoUrl",
    "user_model_attribute": "userModelAttribute",
    "user_object_classes": "userObjectClasses",
    "user_property": "userProperty",
    "user_roles_retrieve_strategy": "userRolesRetrieveStrategy",
    "user_session": "userSession",
    "username_ldap_attribute": "usernameLdapAttribute",
    "users_dn": "usersDn",
    "uuid_ldap_attribute": "uuidLdapAttribute",
    "valid_redirect_uris": "validRedirectUris",
    "validate_password_policy": "validatePasswordPolicy",
    "validate_signature": "validateSignature",
    "verify_email": "verifyEmail",
    "want_assertions_encrypted": "wantAssertionsEncrypted",
    "want_assertions_signed": "wantAssertionsSigned",
    "web_origins": "webOrigins",
    "write_only": "writeOnly",
    "xml_sign_key_info_key_name_transformer": "xmlSignKeyInfoKeyNameTransformer",
}

_CAMEL_TO_SNAKE_CASE_TABLE = {
    "accessCodeLifespan": "access_code_lifespan",
    "accessCodeLifespanLogin": "access_code_lifespan_login",
    "accessCodeLifespanUserAction": "access_code_lifespan_user_action",
    "accessTokenLifespan": "access_token_lifespan",
    "accessTokenLifespanForImplicitFlow": "access_token_lifespan_for_implicit_flow",
    "accessType": "access_type",
    "accountTheme": "account_theme",
    "actionTokenGeneratedByAdminLifespan": "action_token_generated_by_admin_lifespan",
    "actionTokenGeneratedByUserLifespan": "action_token_generated_by_user_lifespan",
    "addReadTokenRoleOnCreate": "add_read_token_role_on_create",
    "addToAccessToken": "add_to_access_token",
    "addToIdToken": "add_to_id_token",
    "addToUserinfo": "add_to_userinfo",
    "adminTheme": "admin_theme",
    "alwaysReadValueFromLdap": "always_read_value_from_ldap",
    "assertionConsumerPostUrl": "assertion_consumer_post_url",
    "assertionConsumerRedirectUrl": "assertion_consumer_redirect_url",
    "attributeFriendlyName": "attribute_friendly_name",
    "attributeName": "attribute_name",
    "attributeValue": "attribute_value",
    "authenticateByDefault": "authenticate_by_default",
    "authorizationUrl": "authorization_url",
    "backchannelSupported": "backchannel_supported",
    "baseUrl": "base_url",
    "batchSizeForSync": "batch_size_for_sync",
    "bindCredential": "bind_credential",
    "bindDn": "bind_dn",
    "browserFlow": "browser_flow",
    "cachePolicy": "cache_policy",
    "changedSyncPeriod": "changed_sync_period",
    "claimName": "claim_name",
    "claimValue": "claim_value",
    "claimValueType": "claim_value_type",
    "clientAuthenticationFlow": "client_authentication_flow",
    "clientId": "client_id",
    "clientScopeId": "client_scope_id",
    "clientSecret": "client_secret",
    "clientSignatureRequired": "client_signature_required",
    "clientTimeout": "client_timeout",
    "compositeRoles": "composite_roles",
    "connectionTimeout": "connection_timeout",
    "connectionUrl": "connection_url",
    "consentScreenText": "consent_screen_text",
    "customUserSearchFilter": "custom_user_search_filter",
    "decisionStrategy": "decision_strategy",
    "defaultAction": "default_action",
    "defaultScopes": "default_scopes",
    "directAccessGrantsEnabled": "direct_access_grants_enabled",
    "directGrantFlow": "direct_grant_flow",
    "displayName": "display_name",
    "dockerAuthenticationFlow": "docker_authentication_flow",
    "dropNonExistingGroupsDuringSync": "drop_non_existing_groups_during_sync",
    "duplicateEmailsAllowed": "duplicate_emails_allowed",
    "editMode": "edit_mode",
    "editUsernameAllowed": "edit_username_allowed",
    "emailTheme": "email_theme",
    "excludeSessionStateFromAuthResponse": "exclude_session_state_from_auth_response",
    "extraConfig": "extra_config",
    "federatedIdentities": "federated_identities",
    "firstBrokerLoginFlowAlias": "first_broker_login_flow_alias",
    "firstName": "first_name",
    "forceAuthn": "force_authn",
    "forcePostBinding": "force_post_binding",
    "friendlyName": "friendly_name",
    "frontChannelLogout": "front_channel_logout",
    "fullPath": "full_path",
    "fullScopeAllowed": "full_scope_allowed",
    "fullSyncPeriod": "full_sync_period",
    "groupId": "group_id",
    "groupIds": "group_ids",
    "groupNameLdapAttribute": "group_name_ldap_attribute",
    "groupObjectClasses": "group_object_classes",
    "groupsLdapFilter": "groups_ldap_filter",
    "hideOnLoginPage": "hide_on_login_page",
    "iconUri": "icon_uri",
    "identityProviderAlias": "identity_provider_alias",
    "idpInitiatedSsoRelayState": "idp_initiated_sso_relay_state",
    "idpInitiatedSsoUrlName": "idp_initiated_sso_url_name",
    "ignoreMissingGroups": "ignore_missing_groups",
    "implicitFlowEnabled": "implicit_flow_enabled",
    "importEnabled": "import_enabled",
    "includeAuthnStatement": "include_authn_statement",
    "includedClientAudience": "included_client_audience",
    "includedCustomAudience": "included_custom_audience",
    "initialLogin": "initial_login",
    "initialPassword": "initial_password",
    "internalId": "internal_id",
    "isMandatoryInLdap": "is_mandatory_in_ldap",
    "jwksUrl": "jwks_url",
    "lastName": "last_name",
    "ldapAttribute": "ldap_attribute",
    "ldapFullNameAttribute": "ldap_full_name_attribute",
    "ldapGroupsDn": "ldap_groups_dn",
    "ldapPasswordPolicyHintsEnabled": "ldap_password_policy_hints_enabled",
    "ldapUserFederationId": "ldap_user_federation_id",
    "linkOnly": "link_only",
    "loginHint": "login_hint",
    "loginTheme": "login_theme",
    "loginWithEmailAllowed": "login_with_email_allowed",
    "logoutServicePostBindingUrl": "logout_service_post_binding_url",
    "logoutServiceRedirectBindingUrl": "logout_service_redirect_binding_url",
    "logoutUrl": "logout_url",
    "mappedGroupAttributes": "mapped_group_attributes",
    "masterSamlProcessingUrl": "master_saml_processing_url",
    "memberofLdapAttribute": "memberof_ldap_attribute",
    "membershipAttributeType": "membership_attribute_type",
    "membershipLdapAttribute": "membership_ldap_attribute",
    "membershipUserLdapAttribute": "membership_user_ldap_attribute",
    "nameIdFormat": "name_id_format",
    "nameIdPolicyFormat": "name_id_policy_format",
    "offlineSessionIdleTimeout": "offline_session_idle_timeout",
    "offlineSessionMaxLifespan": "offline_session_max_lifespan",
    "optionalScopes": "optional_scopes",
    "ownerManagedAccess": "owner_managed_access",
    "parentId": "parent_id",
    "passwordPolicy": "password_policy",
    "pkceCodeChallengeMethod": "pkce_code_challenge_method",
    "postBindingAuthnRequest": "post_binding_authn_request",
    "postBindingLogout": "post_binding_logout",
    "postBindingResponse": "post_binding_response",
    "postBrokerLoginFlowAlias": "post_broker_login_flow_alias",
    "preserveGroupInheritance": "preserve_group_inheritance",
    "protocolMapper": "protocol_mapper",
    "providerId": "provider_id",
    "rdnLdapAttribute": "rdn_ldap_attribute",
    "readOnly": "read_only",
    "readTimeout": "read_timeout",
    "realmId": "realm_id",
    "realmRolePrefix": "realm_role_prefix",
    "refreshTokenMaxReuse": "refresh_token_max_reuse",
    "registrationAllowed": "registration_allowed",
    "registrationEmailAsUsername": "registration_email_as_username",
    "registrationFlow": "registration_flow",
    "rememberMe": "remember_me",
    "resetCredentialsFlow": "reset_credentials_flow",
    "resetPasswordAllowed": "reset_password_allowed",
    "resourceServerId": "resource_server_id",
    "revokeRefreshToken": "revoke_refresh_token",
    "roleId": "role_id",
    "roleIds": "role_ids",
    "rootUrl": "root_url",
    "samlAttributeName": "saml_attribute_name",
    "samlAttributeNameFormat": "saml_attribute_name_format",
    "searchScope": "search_scope",
    "securityDefenses": "security_defenses",
    "serviceAccountUserId": "service_account_user_id",
    "serviceAccountsEnabled": "service_accounts_enabled",
    "signAssertions": "sign_assertions",
    "signDocuments": "sign_documents",
    "signatureAlgorithm": "signature_algorithm",
    "signingCertificate": "signing_certificate",
    "signingPrivateKey": "signing_private_key",
    "singleLogoutServiceUrl": "single_logout_service_url",
    "singleSignOnServiceUrl": "single_sign_on_service_url",
    "smtpServer": "smtp_server",
    "sslRequired": "ssl_required",
    "ssoSessionIdleTimeout": "sso_session_idle_timeout",
    "ssoSessionMaxLifespan": "sso_session_max_lifespan",
    "standardFlowEnabled": "standard_flow_enabled",
    "storeToken": "store_token",
    "syncRegistrations": "sync_registrations",
    "tokenUrl": "token_url",
    "trustEmail": "trust_email",
    "uiLocales": "ui_locales",
    "useTruststoreSpi": "use_truststore_spi",
    "userAttribute": "user_attribute",
    "userInfoUrl": "user_info_url",
    "userModelAttribute": "user_model_attribute",
    "userObjectClasses": "user_object_classes",
    "userProperty": "user_property",
    "userRolesRetrieveStrategy": "user_roles_retrieve_strategy",
    "userSession": "user_session",
    "usernameLdapAttribute": "username_ldap_attribute",
    "usersDn": "users_dn",
    "uuidLdapAttribute": "uuid_ldap_attribute",
    "validRedirectUris": "valid_redirect_uris",
    "validatePasswordPolicy": "validate_password_policy",
    "validateSignature": "validate_signature",
    "verifyEmail": "verify_email",
    "wantAssertionsEncrypted": "want_assertions_encrypted",
    "wantAssertionsSigned": "want_assertions_signed",
    "webOrigins": "web_origins",
    "writeOnly": "write_only",
    "xmlSignKeyInfoKeyNameTransformer": "xml_sign_key_info_key_name_transformer",
}