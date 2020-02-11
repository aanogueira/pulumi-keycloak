// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AttributeToRoleIdentityMapper extends pulumi.CustomResource {
    /**
     * Get an existing AttributeToRoleIdentityMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttributeToRoleIdentityMapperState, opts?: pulumi.CustomResourceOptions): AttributeToRoleIdentityMapper {
        return new AttributeToRoleIdentityMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper';

    /**
     * Returns true if the given object is an instance of AttributeToRoleIdentityMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttributeToRoleIdentityMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttributeToRoleIdentityMapper.__pulumiType;
    }

    /**
     * Attribute Friendly Name
     */
    public readonly attributeFriendlyName!: pulumi.Output<string | undefined>;
    /**
     * Attribute Name
     */
    public readonly attributeName!: pulumi.Output<string | undefined>;
    /**
     * Attribute Value
     */
    public readonly attributeValue!: pulumi.Output<string | undefined>;
    /**
     * OIDC Claim Name
     */
    public readonly claimName!: pulumi.Output<string | undefined>;
    /**
     * OIDC Claim Value
     */
    public readonly claimValue!: pulumi.Output<string | undefined>;
    /**
     * IDP Alias
     */
    public readonly identityProviderAlias!: pulumi.Output<string>;
    /**
     * IDP Mapper Name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Realm Name
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Role Name
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a AttributeToRoleIdentityMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttributeToRoleIdentityMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttributeToRoleIdentityMapperArgs | AttributeToRoleIdentityMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AttributeToRoleIdentityMapperState | undefined;
            inputs["attributeFriendlyName"] = state ? state.attributeFriendlyName : undefined;
            inputs["attributeName"] = state ? state.attributeName : undefined;
            inputs["attributeValue"] = state ? state.attributeValue : undefined;
            inputs["claimName"] = state ? state.claimName : undefined;
            inputs["claimValue"] = state ? state.claimValue : undefined;
            inputs["identityProviderAlias"] = state ? state.identityProviderAlias : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realm"] = state ? state.realm : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as AttributeToRoleIdentityMapperArgs | undefined;
            if (!args || args.identityProviderAlias === undefined) {
                throw new Error("Missing required property 'identityProviderAlias'");
            }
            if (!args || args.realm === undefined) {
                throw new Error("Missing required property 'realm'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["attributeFriendlyName"] = args ? args.attributeFriendlyName : undefined;
            inputs["attributeName"] = args ? args.attributeName : undefined;
            inputs["attributeValue"] = args ? args.attributeValue : undefined;
            inputs["claimName"] = args ? args.claimName : undefined;
            inputs["claimValue"] = args ? args.claimValue : undefined;
            inputs["identityProviderAlias"] = args ? args.identityProviderAlias : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realm"] = args ? args.realm : undefined;
            inputs["role"] = args ? args.role : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AttributeToRoleIdentityMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttributeToRoleIdentityMapper resources.
 */
export interface AttributeToRoleIdentityMapperState {
    /**
     * Attribute Friendly Name
     */
    readonly attributeFriendlyName?: pulumi.Input<string>;
    /**
     * Attribute Name
     */
    readonly attributeName?: pulumi.Input<string>;
    /**
     * Attribute Value
     */
    readonly attributeValue?: pulumi.Input<string>;
    /**
     * OIDC Claim Name
     */
    readonly claimName?: pulumi.Input<string>;
    /**
     * OIDC Claim Value
     */
    readonly claimValue?: pulumi.Input<string>;
    /**
     * IDP Alias
     */
    readonly identityProviderAlias?: pulumi.Input<string>;
    /**
     * IDP Mapper Name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Realm Name
     */
    readonly realm?: pulumi.Input<string>;
    /**
     * Role Name
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AttributeToRoleIdentityMapper resource.
 */
export interface AttributeToRoleIdentityMapperArgs {
    /**
     * Attribute Friendly Name
     */
    readonly attributeFriendlyName?: pulumi.Input<string>;
    /**
     * Attribute Name
     */
    readonly attributeName?: pulumi.Input<string>;
    /**
     * Attribute Value
     */
    readonly attributeValue?: pulumi.Input<string>;
    /**
     * OIDC Claim Name
     */
    readonly claimName?: pulumi.Input<string>;
    /**
     * OIDC Claim Value
     */
    readonly claimValue?: pulumi.Input<string>;
    /**
     * IDP Alias
     */
    readonly identityProviderAlias: pulumi.Input<string>;
    /**
     * IDP Mapper Name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Realm Name
     */
    readonly realm: pulumi.Input<string>;
    /**
     * Role Name
     */
    readonly role: pulumi.Input<string>;
}