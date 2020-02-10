// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.Saml.UserPropertyProtocolMapper
 * 
 * Allows for creating and managing user property protocol mappers for
 * SAML clients within Keycloak.
 * 
 * SAML user property protocol mappers allow you to map properties of the Keycloak
 * user model to an attribute in a SAML assertion. Protocol mappers
 * can be defined for a single client, or they can be defined for a client scope which
 * can be shared between multiple different clients.
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
 * const samlClient = new keycloak.Saml.Client("samlClient", {
 *     clientId: "test-saml-client",
 *     realmId: keycloak_realm_test.id,
 * });
 * const samlUserPropertyMapper = new keycloak.Saml.UserPropertyProtocolMapper("samlUserPropertyMapper", {
 *     clientId: samlClient.id,
 *     realmId: keycloak_realm_test.id,
 *     samlAttributeName: "email",
 *     samlAttributeNameFormat: "Unspecified",
 *     userProperty: "email",
 * });
 * ```
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required if `clientScopeId` is not specified) The SAML client this protocol mapper is attached to.
 * - `clientScopeId` - (Required if `clientId` is not specified) The SAML client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `userProperty` - (Required) The property of the Keycloak user model to map.
 * - `friendlyName` - (Optional) An optional human-friendly name for this attribute.
 * - `samlAttributeName` - (Required) The name of the SAML attribute.
 * - `samlAttributeNameFormat` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
 *
 * > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/saml_user_property_protocol_mapper.html.markdown.
 */
export class UserPropertyProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserPropertyProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPropertyProtocolMapperState, opts?: pulumi.CustomResourceOptions): UserPropertyProtocolMapper {
        return new UserPropertyProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:Saml/userPropertyProtocolMapper:UserPropertyProtocolMapper';

    /**
     * Returns true if the given object is an instance of UserPropertyProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPropertyProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPropertyProtocolMapper.__pulumiType;
    }

    public readonly clientId!: pulumi.Output<string | undefined>;
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly samlAttributeName!: pulumi.Output<string>;
    public readonly samlAttributeNameFormat!: pulumi.Output<string>;
    public readonly userProperty!: pulumi.Output<string>;

    /**
     * Create a UserPropertyProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPropertyProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPropertyProtocolMapperArgs | UserPropertyProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserPropertyProtocolMapperState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["samlAttributeName"] = state ? state.samlAttributeName : undefined;
            inputs["samlAttributeNameFormat"] = state ? state.samlAttributeNameFormat : undefined;
            inputs["userProperty"] = state ? state.userProperty : undefined;
        } else {
            const args = argsOrState as UserPropertyProtocolMapperArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.samlAttributeName === undefined) {
                throw new Error("Missing required property 'samlAttributeName'");
            }
            if (!args || args.samlAttributeNameFormat === undefined) {
                throw new Error("Missing required property 'samlAttributeNameFormat'");
            }
            if (!args || args.userProperty === undefined) {
                throw new Error("Missing required property 'userProperty'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["samlAttributeName"] = args ? args.samlAttributeName : undefined;
            inputs["samlAttributeNameFormat"] = args ? args.samlAttributeNameFormat : undefined;
            inputs["userProperty"] = args ? args.userProperty : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserPropertyProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
 */
export interface UserPropertyProtocolMapperState {
    readonly clientId?: pulumi.Input<string>;
    readonly clientScopeId?: pulumi.Input<string>;
    readonly friendlyName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly samlAttributeName?: pulumi.Input<string>;
    readonly samlAttributeNameFormat?: pulumi.Input<string>;
    readonly userProperty?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPropertyProtocolMapper resource.
 */
export interface UserPropertyProtocolMapperArgs {
    readonly clientId?: pulumi.Input<string>;
    readonly clientScopeId?: pulumi.Input<string>;
    readonly friendlyName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly samlAttributeName: pulumi.Input<string>;
    readonly samlAttributeNameFormat: pulumi.Input<string>;
    readonly userProperty: pulumi.Input<string>;
}
