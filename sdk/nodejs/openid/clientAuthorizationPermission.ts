// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ClientAuthorizationPermission extends pulumi.CustomResource {
    /**
     * Get an existing ClientAuthorizationPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientAuthorizationPermissionState, opts?: pulumi.CustomResourceOptions): ClientAuthorizationPermission {
        return new ClientAuthorizationPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission';

    /**
     * Returns true if the given object is an instance of ClientAuthorizationPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientAuthorizationPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientAuthorizationPermission.__pulumiType;
    }

    public readonly decisionStrategy!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly policies!: pulumi.Output<string[] | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;
    public readonly resources!: pulumi.Output<string[] | undefined>;
    public readonly scopes!: pulumi.Output<string[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ClientAuthorizationPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientAuthorizationPermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientAuthorizationPermissionArgs | ClientAuthorizationPermissionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClientAuthorizationPermissionState | undefined;
            inputs["decisionStrategy"] = state ? state.decisionStrategy : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["resourceServerId"] = state ? state.resourceServerId : undefined;
            inputs["resources"] = state ? state.resources : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ClientAuthorizationPermissionArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.resourceServerId === undefined) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            inputs["decisionStrategy"] = args ? args.decisionStrategy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["resourceServerId"] = args ? args.resourceServerId : undefined;
            inputs["resources"] = args ? args.resources : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClientAuthorizationPermission.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientAuthorizationPermission resources.
 */
export interface ClientAuthorizationPermissionState {
    readonly decisionStrategy?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    readonly realmId?: pulumi.Input<string>;
    readonly resourceServerId?: pulumi.Input<string>;
    readonly resources?: pulumi.Input<pulumi.Input<string>[]>;
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientAuthorizationPermission resource.
 */
export interface ClientAuthorizationPermissionArgs {
    readonly decisionStrategy?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    readonly realmId: pulumi.Input<string>;
    readonly resourceServerId: pulumi.Input<string>;
    readonly resources?: pulumi.Input<pulumi.Input<string>[]>;
    readonly scopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly type?: pulumi.Input<string>;
}
