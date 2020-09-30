// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing an authentication execution within Keycloak.
 *
 * An authentication execution is an action that the user or service may or may not take when authenticating through an authentication
 * flow.
 *
 * > Due to limitations in the Keycloak API, the ordering of authentication executions within a flow must be specified using `dependsOn`. Authentication executions that are created first will appear first within the flow.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const flow = new keycloak.authentication.Flow("flow", {
 *     realmId: realm.id,
 *     alias: "my-flow-alias",
 * });
 * // first execution
 * const executionOne = new keycloak.authentication.Execution("executionOne", {
 *     realmId: realm.id,
 *     parentFlowAlias: flow.alias,
 *     authenticator: "auth-cookie",
 *     requirement: "ALTERNATIVE",
 * });
 * // second execution
 * const executionTwo = new keycloak.authentication.Execution("executionTwo", {
 *     realmId: realm.id,
 *     parentFlowAlias: flow.alias,
 *     authenticator: "identity-provider-redirector",
 *     requirement: "ALTERNATIVE",
 * }, {
 *     dependsOn: [executionOne],
 * });
 * ```
 */
export class Execution extends pulumi.CustomResource {
    /**
     * Get an existing Execution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExecutionState, opts?: pulumi.CustomResourceOptions): Execution {
        return new Execution(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:authentication/execution:Execution';

    /**
     * Returns true if the given object is an instance of Execution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Execution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Execution.__pulumiType;
    }

    /**
     * The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
     */
    public readonly authenticator!: pulumi.Output<string>;
    /**
     * The alias of the flow this execution is attached to.
     */
    public readonly parentFlowAlias!: pulumi.Output<string>;
    /**
     * The realm the authentication execution exists in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
     */
    public readonly requirement!: pulumi.Output<string | undefined>;

    /**
     * Create a Execution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExecutionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExecutionArgs | ExecutionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ExecutionState | undefined;
            inputs["authenticator"] = state ? state.authenticator : undefined;
            inputs["parentFlowAlias"] = state ? state.parentFlowAlias : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["requirement"] = state ? state.requirement : undefined;
        } else {
            const args = argsOrState as ExecutionArgs | undefined;
            if (!args || args.authenticator === undefined) {
                throw new Error("Missing required property 'authenticator'");
            }
            if (!args || args.parentFlowAlias === undefined) {
                throw new Error("Missing required property 'parentFlowAlias'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["authenticator"] = args ? args.authenticator : undefined;
            inputs["parentFlowAlias"] = args ? args.parentFlowAlias : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["requirement"] = args ? args.requirement : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Execution.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Execution resources.
 */
export interface ExecutionState {
    /**
     * The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
     */
    readonly authenticator?: pulumi.Input<string>;
    /**
     * The alias of the flow this execution is attached to.
     */
    readonly parentFlowAlias?: pulumi.Input<string>;
    /**
     * The realm the authentication execution exists in.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
     */
    readonly requirement?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Execution resource.
 */
export interface ExecutionArgs {
    /**
     * The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
     */
    readonly authenticator: pulumi.Input<string>;
    /**
     * The alias of the flow this execution is attached to.
     */
    readonly parentFlowAlias: pulumi.Input<string>;
    /**
     * The realm the authentication execution exists in.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
     */
    readonly requirement?: pulumi.Input<string>;
}
