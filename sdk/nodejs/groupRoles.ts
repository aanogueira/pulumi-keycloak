// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/group_roles.html.markdown.
 */
export class GroupRoles extends pulumi.CustomResource {
    /**
     * Get an existing GroupRoles resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupRolesState, opts?: pulumi.CustomResourceOptions): GroupRoles {
        return new GroupRoles(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/groupRoles:GroupRoles';

    /**
     * Returns true if the given object is an instance of GroupRoles.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupRoles {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupRoles.__pulumiType;
    }

    public readonly groupId!: pulumi.Output<string | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly roleIds!: pulumi.Output<string[]>;

    /**
     * Create a GroupRoles resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupRolesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupRolesArgs | GroupRolesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupRolesState | undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["roleIds"] = state ? state.roleIds : undefined;
        } else {
            const args = argsOrState as GroupRolesArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.roleIds === undefined) {
                throw new Error("Missing required property 'roleIds'");
            }
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["roleIds"] = args ? args.roleIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupRoles.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupRoles resources.
 */
export interface GroupRolesState {
    readonly groupId?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly roleIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a GroupRoles resource.
 */
export interface GroupRolesArgs {
    readonly groupId?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly roleIds: pulumi.Input<pulumi.Input<string>[]>;
}
