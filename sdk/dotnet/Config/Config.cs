// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Keycloak
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("keycloak");

        public static string? ClientId { get; set; } = __config.Get("clientId") ?? Utilities.GetEnv("KEYCLOAK_CLIENT_ID");

        public static string? ClientSecret { get; set; } = __config.Get("clientSecret") ?? Utilities.GetEnv("KEYCLOAK_CLIENT_SECRET");

        /// <summary>
        /// Timeout (in seconds) of the Keycloak client
        /// </summary>
        public static int? ClientTimeout { get; set; } = __config.GetInt32("clientTimeout") ?? Utilities.GetEnvInt32("KEYCLOAK_CLIENT_TIMEOUT") ?? 5;

        /// <summary>
        /// Whether or not to login to Keycloak instance on provider initialization
        /// </summary>
        public static bool? InitialLogin { get; set; } = __config.GetBoolean("initialLogin");

        public static string? Password { get; set; } = __config.Get("password") ?? Utilities.GetEnv("KEYCLOAK_PASSWORD");

        public static string? Realm { get; set; } = __config.Get("realm") ?? Utilities.GetEnv("KEYCLOAK_REALM") ?? "master";

        /// <summary>
        /// The base URL of the Keycloak instance, before `/auth`
        /// </summary>
        public static string? Url { get; set; } = __config.Get("url") ?? Utilities.GetEnv("KEYCLOAK_URL");

        public static string? Username { get; set; } = __config.Get("username") ?? Utilities.GetEnv("KEYCLOAK_USER");

    }
    namespace ConfigTypes
    {
    }
}