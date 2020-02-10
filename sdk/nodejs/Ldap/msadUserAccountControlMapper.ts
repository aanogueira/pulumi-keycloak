// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.Ldap.MsadUserAccountControlMapper
 * 
 * Allows for creating and managing MSAD user account control mappers for Keycloak
 * users federated via LDAP.
 * 
 * The MSAD (Microsoft Active Directory) user account control mapper is specific
 * to LDAP user federation providers that are pulling from AD, and it can propagate
 * AD user state to Keycloak in order to enforce settings like expired passwords
 * or disabled accounts.
 * 
 * ### Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 * 
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "test",
 * });
 * const ldapUserFederation = new keycloak.Ldap.UserFederation("ldapUserFederation", {
 *     bindCredential: "admin",
 *     bindDn: "cn=admin,dc=example,dc=org",
 *     connectionUrl: "ldap://my-ad-server",
 *     rdnLdapAttribute: "cn",
 *     realmId: realm.id,
 *     userObjectClasses: [
 *         "person",
 *         "organizationalPerson",
 *         "user",
 *     ],
 *     usernameLdapAttribute: "cn",
 *     usersDn: "dc=example,dc=org",
 *     uuidLdapAttribute: "objectGUID",
 * });
 * const msadUserAccountControlMapper = new keycloak.Ldap.MsadUserAccountControlMapper("msadUserAccountControlMapper", {
 *     ldapUserFederationId: ldapUserFederation.id,
 *     realmId: realm.id,
 * });
 * ```
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realmId` - (Required) The realm that this LDAP mapper will exist in.
 * - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
 * - `name` - (Required) Display name of this mapper when displayed in the console.
 * - `ldapPasswordPolicyHintsEnabled` - (Optional) When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
 *
 * > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/ldap_msad_user_account_control_mapper.html.markdown.
 */
export class MsadUserAccountControlMapper extends pulumi.CustomResource {
    /**
     * Get an existing MsadUserAccountControlMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MsadUserAccountControlMapperState, opts?: pulumi.CustomResourceOptions): MsadUserAccountControlMapper {
        return new MsadUserAccountControlMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:Ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper';

    /**
     * Returns true if the given object is an instance of MsadUserAccountControlMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MsadUserAccountControlMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MsadUserAccountControlMapper.__pulumiType;
    }

    public readonly ldapPasswordPolicyHintsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a MsadUserAccountControlMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MsadUserAccountControlMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MsadUserAccountControlMapperArgs | MsadUserAccountControlMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MsadUserAccountControlMapperState | undefined;
            inputs["ldapPasswordPolicyHintsEnabled"] = state ? state.ldapPasswordPolicyHintsEnabled : undefined;
            inputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as MsadUserAccountControlMapperArgs | undefined;
            if (!args || args.ldapUserFederationId === undefined) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["ldapPasswordPolicyHintsEnabled"] = args ? args.ldapPasswordPolicyHintsEnabled : undefined;
            inputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MsadUserAccountControlMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MsadUserAccountControlMapper resources.
 */
export interface MsadUserAccountControlMapperState {
    readonly ldapPasswordPolicyHintsEnabled?: pulumi.Input<boolean>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MsadUserAccountControlMapper resource.
 */
export interface MsadUserAccountControlMapperArgs {
    readonly ldapPasswordPolicyHintsEnabled?: pulumi.Input<boolean>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId: pulumi.Input<string>;
}
