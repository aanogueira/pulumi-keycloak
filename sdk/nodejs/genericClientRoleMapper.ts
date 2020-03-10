// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class GenericClientRoleMapper extends pulumi.CustomResource {
    /**
     * Get an existing GenericClientRoleMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GenericClientRoleMapperState, opts?: pulumi.CustomResourceOptions): GenericClientRoleMapper {
        return new GenericClientRoleMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/genericClientRoleMapper:GenericClientRoleMapper';

    /**
     * Returns true if the given object is an instance of GenericClientRoleMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GenericClientRoleMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GenericClientRoleMapper.__pulumiType;
    }

    public readonly clientId!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly roleId!: pulumi.Output<string>;

    /**
     * Create a GenericClientRoleMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GenericClientRoleMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GenericClientRoleMapperArgs | GenericClientRoleMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GenericClientRoleMapperState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as GenericClientRoleMapperArgs | undefined;
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.roleId === undefined) {
                throw new Error("Missing required property 'roleId'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["roleId"] = args ? args.roleId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GenericClientRoleMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GenericClientRoleMapper resources.
 */
export interface GenericClientRoleMapperState {
    readonly clientId?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GenericClientRoleMapper resource.
 */
export interface GenericClientRoleMapperArgs {
    readonly clientId: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly roleId: pulumi.Input<string>;
}
