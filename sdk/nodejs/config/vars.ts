// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("keycloak");

export let clientId: string | undefined = __config.get("clientId") || utilities.getEnv("KEYCLOAK_CLIENT_ID");
export let clientSecret: string | undefined = __config.get("clientSecret") || utilities.getEnv("KEYCLOAK_CLIENT_SECRET");
/**
 * Timeout (in seconds) of the Keycloak client
 */
export let clientTimeout: number | undefined = __config.getObject<number>("clientTimeout") || (utilities.getEnvNumber("KEYCLOAK_CLIENT_TIMEOUT") || 5);
/**
 * Whether or not to login to Keycloak instance on provider initialization
 */
export let initialLogin: boolean | undefined = __config.getObject<boolean>("initialLogin");
export let password: string | undefined = __config.get("password") || utilities.getEnv("KEYCLOAK_PASSWORD");
export let realm: string | undefined = __config.get("realm") || (utilities.getEnv("KEYCLOAK_REALM") || "master");
/**
 * The base URL of the Keycloak instance, before `/auth`
 */
export let url: string | undefined = __config.get("url") || utilities.getEnv("KEYCLOAK_URL");
export let username: string | undefined = __config.get("username") || utilities.getEnv("KEYCLOAK_USER");