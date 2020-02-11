// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # keycloak..User
 * 
 * Allows for creating and managing Users within Keycloak.
 * 
 * This resource was created primarily to enable the acceptance tests for the `keycloak..Group` resource.
 * Creating users within Keycloak is not recommended. Instead, users should be federated from external sources
 * by configuring user federation providers or identity providers.
 * 
 * ### Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 * 
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const user = new keycloak.User("user", {
 *     email: "bob@domain.com",
 *     enabled: true,
 *     firstName: "Bob",
 *     lastName: "Bobson",
 *     realmId: realm.id,
 *     username: "bob",
 * });
 * const userWithInitialPassword = new keycloak.User("userWithInitialPassword", {
 *     email: "alice@domain.com",
 *     enabled: true,
 *     firstName: "Alice",
 *     initialPassword: {
 *         temporary: true,
 *         value: "some password",
 *     },
 *     lastName: "Aliceberg",
 *     realmId: realm.id,
 *     username: "alice",
 * });
 * ```
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realmId` - (Required) The realm this user belongs to.
 * - `username` - (Required) The unique username of this user.
 * - `initialPassword` (Optional) When given, the user's initial password will be set.
 *    This attribute is only respected during initial user creation.
 *     - `value` (Required) The initial password.
 *     - `temporary` (Optional) If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
 * - `enabled` - (Optional) When false, this user cannot log in. Defaults to `true`.
 * - `email` - (Optional) The user's email.
 * - `firstName` - (Optional) The user's first name.
 * - `lastName` - (Optional) The user's last name.
 *
 * > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/user.html.markdown.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    public readonly attributes!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly email!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly federatedIdentities!: pulumi.Output<outputs.UserFederatedIdentity[] | undefined>;
    public readonly firstName!: pulumi.Output<string | undefined>;
    public readonly initialPassword!: pulumi.Output<outputs.UserInitialPassword | undefined>;
    public readonly lastName!: pulumi.Output<string | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["federatedIdentities"] = state ? state.federatedIdentities : undefined;
            inputs["firstName"] = state ? state.firstName : undefined;
            inputs["initialPassword"] = state ? state.initialPassword : undefined;
            inputs["lastName"] = state ? state.lastName : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.username === undefined) {
                throw new Error("Missing required property 'username'");
            }
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["federatedIdentities"] = args ? args.federatedIdentities : undefined;
            inputs["firstName"] = args ? args.firstName : undefined;
            inputs["initialPassword"] = args ? args.initialPassword : undefined;
            inputs["lastName"] = args ? args.lastName : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    readonly email?: pulumi.Input<string>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly federatedIdentities?: pulumi.Input<pulumi.Input<inputs.UserFederatedIdentity>[]>;
    readonly firstName?: pulumi.Input<string>;
    readonly initialPassword?: pulumi.Input<inputs.UserInitialPassword>;
    readonly lastName?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    readonly email?: pulumi.Input<string>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly federatedIdentities?: pulumi.Input<pulumi.Input<inputs.UserFederatedIdentity>[]>;
    readonly firstName?: pulumi.Input<string>;
    readonly initialPassword?: pulumi.Input<inputs.UserInitialPassword>;
    readonly lastName?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly username: pulumi.Input<string>;
}