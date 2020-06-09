// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class UserClientRoleProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserClientRoleProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserClientRoleProtocolMapperState, opts?: pulumi.CustomResourceOptions): UserClientRoleProtocolMapper {
        return new UserClientRoleProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper';

    /**
     * Returns true if the given object is an instance of UserClientRoleProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserClientRoleProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserClientRoleProtocolMapper.__pulumiType;
    }

    /**
     * Indicates if the attribute should be a claim in the access token.
     */
    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the attribute should be a claim in the id token.
     */
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     */
    public readonly addToUserinfo!: pulumi.Output<boolean | undefined>;
    public readonly claimName!: pulumi.Output<string>;
    /**
     * Claim type used when serializing tokens.
     */
    public readonly claimValueType!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * Client ID for role mappings.
     */
    public readonly clientIdForRoleMappings!: pulumi.Output<string | undefined>;
    /**
     * Prefix that will be added to each client role.
     */
    public readonly clientRolePrefix!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether this attribute is a single value or an array of values.
     */
    public readonly multivalued!: pulumi.Output<boolean | undefined>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a UserClientRoleProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserClientRoleProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserClientRoleProtocolMapperArgs | UserClientRoleProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserClientRoleProtocolMapperState | undefined;
            inputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            inputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            inputs["addToUserinfo"] = state ? state.addToUserinfo : undefined;
            inputs["claimName"] = state ? state.claimName : undefined;
            inputs["claimValueType"] = state ? state.claimValueType : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientIdForRoleMappings"] = state ? state.clientIdForRoleMappings : undefined;
            inputs["clientRolePrefix"] = state ? state.clientRolePrefix : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["multivalued"] = state ? state.multivalued : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as UserClientRoleProtocolMapperArgs | undefined;
            if (!args || args.claimName === undefined) {
                throw new Error("Missing required property 'claimName'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            inputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            inputs["addToUserinfo"] = args ? args.addToUserinfo : undefined;
            inputs["claimName"] = args ? args.claimName : undefined;
            inputs["claimValueType"] = args ? args.claimValueType : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientIdForRoleMappings"] = args ? args.clientIdForRoleMappings : undefined;
            inputs["clientRolePrefix"] = args ? args.clientRolePrefix : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["multivalued"] = args ? args.multivalued : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserClientRoleProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserClientRoleProtocolMapper resources.
 */
export interface UserClientRoleProtocolMapperState {
    /**
     * Indicates if the attribute should be a claim in the access token.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should be a claim in the id token.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    readonly claimName?: pulumi.Input<string>;
    /**
     * Claim type used when serializing tokens.
     */
    readonly claimValueType?: pulumi.Input<string>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client ID for role mappings.
     */
    readonly clientIdForRoleMappings?: pulumi.Input<string>;
    /**
     * Prefix that will be added to each client role.
     */
    readonly clientRolePrefix?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates whether this attribute is a single value or an array of values.
     */
    readonly multivalued?: pulumi.Input<boolean>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserClientRoleProtocolMapper resource.
 */
export interface UserClientRoleProtocolMapperArgs {
    /**
     * Indicates if the attribute should be a claim in the access token.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should be a claim in the id token.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    readonly claimName: pulumi.Input<string>;
    /**
     * Claim type used when serializing tokens.
     */
    readonly claimValueType?: pulumi.Input<string>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client ID for role mappings.
     */
    readonly clientIdForRoleMappings?: pulumi.Input<string>;
    /**
     * Prefix that will be added to each client role.
     */
    readonly clientRolePrefix?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates whether this attribute is a single value or an array of values.
     */
    readonly multivalued?: pulumi.Input<boolean>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId: pulumi.Input<string>;
}
