// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing script protocol mappers within Keycloak.
 *
 * Script protocol mappers evaluate a JavaScript function to produce a token claim based on context information.
 *
 * Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
 * multiple different clients.
 *
 * ## Example Usage
 * ### Client)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const openidClient = new keycloak.openid.Client("openidClient", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const scriptMapper = new keycloak.openid.ScriptProtocolMapper("scriptMapper", {
 *     realmId: realm.id,
 *     clientId: openidClient.id,
 *     claimName: "foo",
 *     script: "exports = 'foo';",
 * });
 * ```
 * ### Client Scope)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const clientScope = new keycloak.openid.ClientScope("clientScope", {realmId: realm.id});
 * const scriptMapper = new keycloak.openid.ScriptProtocolMapper("scriptMapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     claimName: "foo",
 *     script: "exports = 'foo';",
 * });
 * ```
 *
 * ## Import
 *
 * Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper script_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 *
 * ```sh
 *  $ pulumi import keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper script_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 */
export class ScriptProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing ScriptProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScriptProtocolMapperState, opts?: pulumi.CustomResourceOptions): ScriptProtocolMapper {
        return new ScriptProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper';

    /**
     * Returns true if the given object is an instance of ScriptProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScriptProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScriptProtocolMapper.__pulumiType;
    }

    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     */
    public readonly addToUserinfo!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the claim to insert into a token.
     */
    public readonly claimName!: pulumi.Output<string>;
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     */
    public readonly claimValueType!: pulumi.Output<string | undefined>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
     */
    public readonly multivalued!: pulumi.Output<boolean | undefined>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * JavaScript code to compute the claim value.
     */
    public readonly script!: pulumi.Output<string>;

    /**
     * Create a ScriptProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScriptProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScriptProtocolMapperArgs | ScriptProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScriptProtocolMapperState | undefined;
            inputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            inputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            inputs["addToUserinfo"] = state ? state.addToUserinfo : undefined;
            inputs["claimName"] = state ? state.claimName : undefined;
            inputs["claimValueType"] = state ? state.claimValueType : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["multivalued"] = state ? state.multivalued : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["script"] = state ? state.script : undefined;
        } else {
            const args = argsOrState as ScriptProtocolMapperArgs | undefined;
            if ((!args || args.claimName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'claimName'");
            }
            if ((!args || args.realmId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.script === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'script'");
            }
            inputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            inputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            inputs["addToUserinfo"] = args ? args.addToUserinfo : undefined;
            inputs["claimName"] = args ? args.claimName : undefined;
            inputs["claimValueType"] = args ? args.claimValueType : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["multivalued"] = args ? args.multivalued : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["script"] = args ? args.script : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ScriptProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScriptProtocolMapper resources.
 */
export interface ScriptProtocolMapperState {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    /**
     * The name of the claim to insert into a token.
     */
    readonly claimName?: pulumi.Input<string>;
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     */
    readonly claimValueType?: pulumi.Input<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
     */
    readonly multivalued?: pulumi.Input<boolean>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * JavaScript code to compute the claim value.
     */
    readonly script?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScriptProtocolMapper resource.
 */
export interface ScriptProtocolMapperArgs {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    /**
     * The name of the claim to insert into a token.
     */
    readonly claimName: pulumi.Input<string>;
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     */
    readonly claimValueType?: pulumi.Input<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
     */
    readonly multivalued?: pulumi.Input<boolean>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * JavaScript code to compute the claim value.
     */
    readonly script: pulumi.Input<string>;
}