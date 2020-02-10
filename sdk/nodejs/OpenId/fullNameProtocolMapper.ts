// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## # keycloak.OpenId.FullNameProtocolMapper
 * 
 * Allows for creating and managing full name protocol mappers within
 * Keycloak.
 * 
 * Full name protocol mappers allow you to map a user's first and last name
 * to the OpenID Connect `name` claim in a token. Protocol mappers can be defined
 * for a single client, or they can be defined for a client scope which can
 * be shared between multiple different clients.
 * 
 * ### Example Usage (Client)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 * 
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const openidClient = new keycloak.OpenId.Client("openidClient", {
 *     accessType: "CONFIDENTIAL",
 *     clientId: "test-client",
 *     enabled: true,
 *     realmId: realm.id,
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const fullNameMapper = new keycloak.OpenId.FullNameProtocolMapper("fullNameMapper", {
 *     clientId: openidClient.id,
 *     realmId: realm.id,
 * });
 * ```
 * 
 * ### Example Usage (Client Scope)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 * 
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const clientScope = new keycloak.OpenId.ClientScope("clientScope", {
 *     realmId: realm.id,
 * });
 * const fullNameMapper = new keycloak.OpenId.FullNameProtocolMapper("fullNameMapper", {
 *     clientScopeId: clientScope.id,
 *     realmId: realm.id,
 * });
 * ```
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
 * - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `addToIdToken` - (Optional) Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
 * - `addToAccessToken` - (Optional) Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
 * - `addToUserinfo` - (Optional) Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
 *
 * > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_full_name_protocol_mapper.html.markdown.
 */
export class FullNameProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing FullNameProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FullNameProtocolMapperState, opts?: pulumi.CustomResourceOptions): FullNameProtocolMapper {
        return new FullNameProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:OpenId/fullNameProtocolMapper:FullNameProtocolMapper';

    /**
     * Returns true if the given object is an instance of FullNameProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FullNameProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FullNameProtocolMapper.__pulumiType;
    }

    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    public readonly addToUserinfo!: pulumi.Output<boolean | undefined>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a FullNameProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FullNameProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FullNameProtocolMapperArgs | FullNameProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FullNameProtocolMapperState | undefined;
            inputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            inputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            inputs["addToUserinfo"] = state ? state.addToUserinfo : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as FullNameProtocolMapperArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            inputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            inputs["addToUserinfo"] = args ? args.addToUserinfo : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FullNameProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FullNameProtocolMapper resources.
 */
export interface FullNameProtocolMapperState {
    readonly addToAccessToken?: pulumi.Input<boolean>;
    readonly addToIdToken?: pulumi.Input<boolean>;
    readonly addToUserinfo?: pulumi.Input<boolean>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
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
 * The set of arguments for constructing a FullNameProtocolMapper resource.
 */
export interface FullNameProtocolMapperArgs {
    readonly addToAccessToken?: pulumi.Input<boolean>;
    readonly addToIdToken?: pulumi.Input<boolean>;
    readonly addToUserinfo?: pulumi.Input<boolean>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId: pulumi.Input<string>;
}
