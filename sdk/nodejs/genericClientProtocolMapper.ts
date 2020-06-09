// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # keycloak..GenericClientProtocolMapper
 *
 * Allows for creating and managing protocol mapper for both types of clients (openid-connect and saml) within Keycloak.
 *
 * There are two uses cases for using this resource:
 * * If you implemented a custom protocol mapper, this resource can be used to configure it
 * * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
 *
 * Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors. 
 * Therefore, if possible, a specific mapper should be used.
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
 * const samlClient = new keycloak.saml.Client("samlClient", {
 *     clientId: "test-client",
 *     realmId: realm.id,
 * });
 * const samlHardcodeAttributeMapper = new keycloak.GenericClientProtocolMapper("samlHardcodeAttributeMapper", {
 *     clientId: samlClient.id,
 *     config: {
 *         "attribute.name": "name",
 *         "attribute.nameformat": "Basic",
 *         "attribute.value": "value",
 *         "friendly.name": "display name",
 *     },
 *     protocol: "saml",
 *     protocolMapper: "saml-hardcode-attribute-mapper",
 *     realmId: realm.id,
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required) The client this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `protocol` - (Required) The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
 * - `protocolMapper` - (Required) The name of the protocol mapper. The protocol mapper must be
 *    compatible with the specified client.
 * - `config` - (Required) A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
 */
export class GenericClientProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing GenericClientProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GenericClientProtocolMapperState, opts?: pulumi.CustomResourceOptions): GenericClientProtocolMapper {
        return new GenericClientProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper';

    /**
     * Returns true if the given object is an instance of GenericClientProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GenericClientProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GenericClientProtocolMapper.__pulumiType;
    }

    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    public readonly config!: pulumi.Output<{[key: string]: any}>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protocol of the client (openid-connect / saml).
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The type of the protocol mapper.
     */
    public readonly protocolMapper!: pulumi.Output<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a GenericClientProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GenericClientProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GenericClientProtocolMapperArgs | GenericClientProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GenericClientProtocolMapperState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["protocolMapper"] = state ? state.protocolMapper : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as GenericClientProtocolMapperArgs | undefined;
            if (!args || args.config === undefined) {
                throw new Error("Missing required property 'config'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.protocolMapper === undefined) {
                throw new Error("Missing required property 'protocolMapper'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["protocolMapper"] = args ? args.protocolMapper : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GenericClientProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GenericClientProtocolMapper resources.
 */
export interface GenericClientProtocolMapperState {
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    readonly config?: pulumi.Input<{[key: string]: any}>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol of the client (openid-connect / saml).
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The type of the protocol mapper.
     */
    readonly protocolMapper?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GenericClientProtocolMapper resource.
 */
export interface GenericClientProtocolMapperArgs {
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    readonly config: pulumi.Input<{[key: string]: any}>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol of the client (openid-connect / saml).
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The type of the protocol mapper.
     */
    readonly protocolMapper: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId: pulumi.Input<string>;
}
