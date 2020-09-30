// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing LDAP user federation providers within Keycloak.
//
// Keycloak can use an LDAP user federation provider to federate users to Keycloak
// from a directory system such as LDAP or Active Directory. Federated users
// will exist within the realm and will be able to log in to clients. Federated
// users can have their attributes defined using mappers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/ldap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ldap.NewUserFederation(ctx, "ldapUserFederation", &ldap.UserFederationArgs{
// 			RealmId:               realm.ID(),
// 			Enabled:               pulumi.Bool(true),
// 			UsernameLdapAttribute: pulumi.String("cn"),
// 			RdnLdapAttribute:      pulumi.String("cn"),
// 			UuidLdapAttribute:     pulumi.String("entryDN"),
// 			UserObjectClasses: pulumi.StringArray{
// 				pulumi.String("simpleSecurityObject"),
// 				pulumi.String("organizationalRole"),
// 			},
// 			ConnectionUrl:     pulumi.String("ldap://openldap"),
// 			UsersDn:           pulumi.String("dc=example,dc=org"),
// 			BindDn:            pulumi.String("cn=admin,dc=example,dc=org"),
// 			BindCredential:    pulumi.String("admin"),
// 			ConnectionTimeout: pulumi.String("5s"),
// 			ReadTimeout:       pulumi.String("10s"),
// 			Kerberos: &ldap.UserFederationKerberosArgs{
// 				KerberosRealm:   pulumi.String("FOO.LOCAL"),
// 				ServerPrincipal: pulumi.String("HTTP/host.foo.com@FOO.LOCAL"),
// 				Keytab:          pulumi.String("/etc/host.keytab"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type UserFederation struct {
	pulumi.CustomResourceState

	// The number of users to sync within a single transaction. Defaults to `1000`.
	BatchSizeForSync pulumi.IntPtrOutput `pulumi:"batchSizeForSync"`
	// Password of LDAP admin. This attribute must be set if `bindDn` is set.
	BindCredential pulumi.StringPtrOutput `pulumi:"bindCredential"`
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if `bindCredential` is set.
	BindDn pulumi.StringPtrOutput `pulumi:"bindDn"`
	// A block containing the cache settings.
	Cache UserFederationCachePtrOutput `pulumi:"cache"`
	// **Deprecated** Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	//
	// Deprecated: use cache.policy instead
	CachePolicy pulumi.StringPtrOutput `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod pulumi.IntPtrOutput `pulumi:"changedSyncPeriod"`
	// LDAP connection timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ConnectionTimeout pulumi.StringPtrOutput `pulumi:"connectionTimeout"`
	// Connection URL to the LDAP server.
	ConnectionUrl pulumi.StringOutput `pulumi:"connectionUrl"`
	// Additional LDAP filter for filtering searched users. Must begin with `(` and end with `)`.
	CustomUserSearchFilter pulumi.StringPtrOutput `pulumi:"customUserSearchFilter"`
	// Can be one of `READ_ONLY`, `WRITABLE`, or `UNSYNCED`. `UNSYNCED` allows user data to be imported but not synced back to LDAP. Defaults to `READ_ONLY`.
	EditMode pulumi.StringPtrOutput `pulumi:"editMode"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrOutput `pulumi:"fullSyncPeriod"`
	// When `true`, LDAP users will be imported into the Keycloak database. Defaults to `true`.
	ImportEnabled pulumi.BoolPtrOutput `pulumi:"importEnabled"`
	// A block containing the kerberos settings.
	Kerberos UserFederationKerberosPtrOutput `pulumi:"kerberos"`
	// Display name of the provider when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to `true`.
	Pagination pulumi.BoolPtrOutput `pulumi:"pagination"`
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute pulumi.StringOutput `pulumi:"rdnLdapAttribute"`
	// LDAP read timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ReadTimeout pulumi.StringPtrOutput `pulumi:"readTimeout"`
	// The realm that this provider will provide user federation for.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Can be one of `ONE_LEVEL` or `SUBTREE`:
	// - `ONE_LEVEL`: Only search for users in the DN specified by `userDn`.
	// - `SUBTREE`: Search entire LDAP subtree.
	SearchScope pulumi.StringPtrOutput `pulumi:"searchScope"`
	// When `true`, newly created users will be synced back to LDAP. Defaults to `false`.
	SyncRegistrations pulumi.BoolPtrOutput `pulumi:"syncRegistrations"`
	// Can be one of `ALWAYS`, `ONLY_FOR_LDAPS`, or `NEVER`:
	UseTruststoreSpi pulumi.StringPtrOutput `pulumi:"useTruststoreSpi"`
	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	UserObjectClasses pulumi.StringArrayOutput `pulumi:"userObjectClasses"`
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute pulumi.StringOutput `pulumi:"usernameLdapAttribute"`
	// Full DN of LDAP tree where your users are.
	UsersDn pulumi.StringOutput `pulumi:"usersDn"`
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute pulumi.StringOutput `pulumi:"uuidLdapAttribute"`
	// When `true`, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy pulumi.BoolPtrOutput `pulumi:"validatePasswordPolicy"`
	// Can be one of `OTHER`, `EDIRECTORY`, `AD`, `RHDS`, or `TIVOLI`. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to `OTHER`.
	Vendor pulumi.StringPtrOutput `pulumi:"vendor"`
}

// NewUserFederation registers a new resource with the given unique name, arguments, and options.
func NewUserFederation(ctx *pulumi.Context,
	name string, args *UserFederationArgs, opts ...pulumi.ResourceOption) (*UserFederation, error) {
	if args == nil || args.ConnectionUrl == nil {
		return nil, errors.New("missing required argument 'ConnectionUrl'")
	}
	if args == nil || args.RdnLdapAttribute == nil {
		return nil, errors.New("missing required argument 'RdnLdapAttribute'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.UserObjectClasses == nil {
		return nil, errors.New("missing required argument 'UserObjectClasses'")
	}
	if args == nil || args.UsernameLdapAttribute == nil {
		return nil, errors.New("missing required argument 'UsernameLdapAttribute'")
	}
	if args == nil || args.UsersDn == nil {
		return nil, errors.New("missing required argument 'UsersDn'")
	}
	if args == nil || args.UuidLdapAttribute == nil {
		return nil, errors.New("missing required argument 'UuidLdapAttribute'")
	}
	if args == nil {
		args = &UserFederationArgs{}
	}
	var resource UserFederation
	err := ctx.RegisterResource("keycloak:ldap/userFederation:UserFederation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFederation gets an existing UserFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFederation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFederationState, opts ...pulumi.ResourceOption) (*UserFederation, error) {
	var resource UserFederation
	err := ctx.ReadResource("keycloak:ldap/userFederation:UserFederation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFederation resources.
type userFederationState struct {
	// The number of users to sync within a single transaction. Defaults to `1000`.
	BatchSizeForSync *int `pulumi:"batchSizeForSync"`
	// Password of LDAP admin. This attribute must be set if `bindDn` is set.
	BindCredential *string `pulumi:"bindCredential"`
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if `bindCredential` is set.
	BindDn *string `pulumi:"bindDn"`
	// A block containing the cache settings.
	Cache *UserFederationCache `pulumi:"cache"`
	// **Deprecated** Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	//
	// Deprecated: use cache.policy instead
	CachePolicy *string `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod *int `pulumi:"changedSyncPeriod"`
	// LDAP connection timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ConnectionTimeout *string `pulumi:"connectionTimeout"`
	// Connection URL to the LDAP server.
	ConnectionUrl *string `pulumi:"connectionUrl"`
	// Additional LDAP filter for filtering searched users. Must begin with `(` and end with `)`.
	CustomUserSearchFilter *string `pulumi:"customUserSearchFilter"`
	// Can be one of `READ_ONLY`, `WRITABLE`, or `UNSYNCED`. `UNSYNCED` allows user data to be imported but not synced back to LDAP. Defaults to `READ_ONLY`.
	EditMode *string `pulumi:"editMode"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *int `pulumi:"fullSyncPeriod"`
	// When `true`, LDAP users will be imported into the Keycloak database. Defaults to `true`.
	ImportEnabled *bool `pulumi:"importEnabled"`
	// A block containing the kerberos settings.
	Kerberos *UserFederationKerberos `pulumi:"kerberos"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to `true`.
	Pagination *bool `pulumi:"pagination"`
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority *int `pulumi:"priority"`
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute *string `pulumi:"rdnLdapAttribute"`
	// LDAP read timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ReadTimeout *string `pulumi:"readTimeout"`
	// The realm that this provider will provide user federation for.
	RealmId *string `pulumi:"realmId"`
	// Can be one of `ONE_LEVEL` or `SUBTREE`:
	// - `ONE_LEVEL`: Only search for users in the DN specified by `userDn`.
	// - `SUBTREE`: Search entire LDAP subtree.
	SearchScope *string `pulumi:"searchScope"`
	// When `true`, newly created users will be synced back to LDAP. Defaults to `false`.
	SyncRegistrations *bool `pulumi:"syncRegistrations"`
	// Can be one of `ALWAYS`, `ONLY_FOR_LDAPS`, or `NEVER`:
	UseTruststoreSpi *string `pulumi:"useTruststoreSpi"`
	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	UserObjectClasses []string `pulumi:"userObjectClasses"`
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute *string `pulumi:"usernameLdapAttribute"`
	// Full DN of LDAP tree where your users are.
	UsersDn *string `pulumi:"usersDn"`
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute *string `pulumi:"uuidLdapAttribute"`
	// When `true`, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy *bool `pulumi:"validatePasswordPolicy"`
	// Can be one of `OTHER`, `EDIRECTORY`, `AD`, `RHDS`, or `TIVOLI`. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to `OTHER`.
	Vendor *string `pulumi:"vendor"`
}

type UserFederationState struct {
	// The number of users to sync within a single transaction. Defaults to `1000`.
	BatchSizeForSync pulumi.IntPtrInput
	// Password of LDAP admin. This attribute must be set if `bindDn` is set.
	BindCredential pulumi.StringPtrInput
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if `bindCredential` is set.
	BindDn pulumi.StringPtrInput
	// A block containing the cache settings.
	Cache UserFederationCachePtrInput
	// **Deprecated** Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	//
	// Deprecated: use cache.policy instead
	CachePolicy pulumi.StringPtrInput
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod pulumi.IntPtrInput
	// LDAP connection timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ConnectionTimeout pulumi.StringPtrInput
	// Connection URL to the LDAP server.
	ConnectionUrl pulumi.StringPtrInput
	// Additional LDAP filter for filtering searched users. Must begin with `(` and end with `)`.
	CustomUserSearchFilter pulumi.StringPtrInput
	// Can be one of `READ_ONLY`, `WRITABLE`, or `UNSYNCED`. `UNSYNCED` allows user data to be imported but not synced back to LDAP. Defaults to `READ_ONLY`.
	EditMode pulumi.StringPtrInput
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrInput
	// When `true`, LDAP users will be imported into the Keycloak database. Defaults to `true`.
	ImportEnabled pulumi.BoolPtrInput
	// A block containing the kerberos settings.
	Kerberos UserFederationKerberosPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to `true`.
	Pagination pulumi.BoolPtrInput
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority pulumi.IntPtrInput
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute pulumi.StringPtrInput
	// LDAP read timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ReadTimeout pulumi.StringPtrInput
	// The realm that this provider will provide user federation for.
	RealmId pulumi.StringPtrInput
	// Can be one of `ONE_LEVEL` or `SUBTREE`:
	// - `ONE_LEVEL`: Only search for users in the DN specified by `userDn`.
	// - `SUBTREE`: Search entire LDAP subtree.
	SearchScope pulumi.StringPtrInput
	// When `true`, newly created users will be synced back to LDAP. Defaults to `false`.
	SyncRegistrations pulumi.BoolPtrInput
	// Can be one of `ALWAYS`, `ONLY_FOR_LDAPS`, or `NEVER`:
	UseTruststoreSpi pulumi.StringPtrInput
	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	UserObjectClasses pulumi.StringArrayInput
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute pulumi.StringPtrInput
	// Full DN of LDAP tree where your users are.
	UsersDn pulumi.StringPtrInput
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute pulumi.StringPtrInput
	// When `true`, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy pulumi.BoolPtrInput
	// Can be one of `OTHER`, `EDIRECTORY`, `AD`, `RHDS`, or `TIVOLI`. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to `OTHER`.
	Vendor pulumi.StringPtrInput
}

func (UserFederationState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFederationState)(nil)).Elem()
}

type userFederationArgs struct {
	// The number of users to sync within a single transaction. Defaults to `1000`.
	BatchSizeForSync *int `pulumi:"batchSizeForSync"`
	// Password of LDAP admin. This attribute must be set if `bindDn` is set.
	BindCredential *string `pulumi:"bindCredential"`
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if `bindCredential` is set.
	BindDn *string `pulumi:"bindDn"`
	// A block containing the cache settings.
	Cache *UserFederationCache `pulumi:"cache"`
	// **Deprecated** Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	//
	// Deprecated: use cache.policy instead
	CachePolicy *string `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod *int `pulumi:"changedSyncPeriod"`
	// LDAP connection timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ConnectionTimeout *string `pulumi:"connectionTimeout"`
	// Connection URL to the LDAP server.
	ConnectionUrl string `pulumi:"connectionUrl"`
	// Additional LDAP filter for filtering searched users. Must begin with `(` and end with `)`.
	CustomUserSearchFilter *string `pulumi:"customUserSearchFilter"`
	// Can be one of `READ_ONLY`, `WRITABLE`, or `UNSYNCED`. `UNSYNCED` allows user data to be imported but not synced back to LDAP. Defaults to `READ_ONLY`.
	EditMode *string `pulumi:"editMode"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *int `pulumi:"fullSyncPeriod"`
	// When `true`, LDAP users will be imported into the Keycloak database. Defaults to `true`.
	ImportEnabled *bool `pulumi:"importEnabled"`
	// A block containing the kerberos settings.
	Kerberos *UserFederationKerberos `pulumi:"kerberos"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to `true`.
	Pagination *bool `pulumi:"pagination"`
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority *int `pulumi:"priority"`
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute string `pulumi:"rdnLdapAttribute"`
	// LDAP read timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ReadTimeout *string `pulumi:"readTimeout"`
	// The realm that this provider will provide user federation for.
	RealmId string `pulumi:"realmId"`
	// Can be one of `ONE_LEVEL` or `SUBTREE`:
	// - `ONE_LEVEL`: Only search for users in the DN specified by `userDn`.
	// - `SUBTREE`: Search entire LDAP subtree.
	SearchScope *string `pulumi:"searchScope"`
	// When `true`, newly created users will be synced back to LDAP. Defaults to `false`.
	SyncRegistrations *bool `pulumi:"syncRegistrations"`
	// Can be one of `ALWAYS`, `ONLY_FOR_LDAPS`, or `NEVER`:
	UseTruststoreSpi *string `pulumi:"useTruststoreSpi"`
	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	UserObjectClasses []string `pulumi:"userObjectClasses"`
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute string `pulumi:"usernameLdapAttribute"`
	// Full DN of LDAP tree where your users are.
	UsersDn string `pulumi:"usersDn"`
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute string `pulumi:"uuidLdapAttribute"`
	// When `true`, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy *bool `pulumi:"validatePasswordPolicy"`
	// Can be one of `OTHER`, `EDIRECTORY`, `AD`, `RHDS`, or `TIVOLI`. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to `OTHER`.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a UserFederation resource.
type UserFederationArgs struct {
	// The number of users to sync within a single transaction. Defaults to `1000`.
	BatchSizeForSync pulumi.IntPtrInput
	// Password of LDAP admin. This attribute must be set if `bindDn` is set.
	BindCredential pulumi.StringPtrInput
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if `bindCredential` is set.
	BindDn pulumi.StringPtrInput
	// A block containing the cache settings.
	Cache UserFederationCachePtrInput
	// **Deprecated** Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	//
	// Deprecated: use cache.policy instead
	CachePolicy pulumi.StringPtrInput
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod pulumi.IntPtrInput
	// LDAP connection timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ConnectionTimeout pulumi.StringPtrInput
	// Connection URL to the LDAP server.
	ConnectionUrl pulumi.StringInput
	// Additional LDAP filter for filtering searched users. Must begin with `(` and end with `)`.
	CustomUserSearchFilter pulumi.StringPtrInput
	// Can be one of `READ_ONLY`, `WRITABLE`, or `UNSYNCED`. `UNSYNCED` allows user data to be imported but not synced back to LDAP. Defaults to `READ_ONLY`.
	EditMode pulumi.StringPtrInput
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrInput
	// When `true`, LDAP users will be imported into the Keycloak database. Defaults to `true`.
	ImportEnabled pulumi.BoolPtrInput
	// A block containing the kerberos settings.
	Kerberos UserFederationKerberosPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to `true`.
	Pagination pulumi.BoolPtrInput
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority pulumi.IntPtrInput
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute pulumi.StringInput
	// LDAP read timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
	ReadTimeout pulumi.StringPtrInput
	// The realm that this provider will provide user federation for.
	RealmId pulumi.StringInput
	// Can be one of `ONE_LEVEL` or `SUBTREE`:
	// - `ONE_LEVEL`: Only search for users in the DN specified by `userDn`.
	// - `SUBTREE`: Search entire LDAP subtree.
	SearchScope pulumi.StringPtrInput
	// When `true`, newly created users will be synced back to LDAP. Defaults to `false`.
	SyncRegistrations pulumi.BoolPtrInput
	// Can be one of `ALWAYS`, `ONLY_FOR_LDAPS`, or `NEVER`:
	UseTruststoreSpi pulumi.StringPtrInput
	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	UserObjectClasses pulumi.StringArrayInput
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute pulumi.StringInput
	// Full DN of LDAP tree where your users are.
	UsersDn pulumi.StringInput
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute pulumi.StringInput
	// When `true`, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy pulumi.BoolPtrInput
	// Can be one of `OTHER`, `EDIRECTORY`, `AD`, `RHDS`, or `TIVOLI`. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to `OTHER`.
	Vendor pulumi.StringPtrInput
}

func (UserFederationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFederationArgs)(nil)).Elem()
}
