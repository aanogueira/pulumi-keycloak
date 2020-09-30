// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Outputs
{

    [OutputType]
    public sealed class RealmSecurityDefensesHeaders
    {
        /// <summary>
        /// Sets the Content Security Policy, which can be used for prevent pages from being included by non-origin iframes. More information can be found in the [W3C-CSP](https://www.w3.org/TR/CSP/) Abstract.
        /// </summary>
        public readonly string? ContentSecurityPolicy;
        /// <summary>
        /// Used for testing Content Security Policies.
        /// </summary>
        public readonly string? ContentSecurityPolicyReportOnly;
        /// <summary>
        /// The Script-Transport-Security HTTP header tells browsers to always use HTTPS.
        /// </summary>
        public readonly string? StrictTransportSecurity;
        /// <summary>
        /// Sets the X-Content-Type-Options, which can be used for prevent MIME-sniffing a response away from the declared content-type
        /// </summary>
        public readonly string? XContentTypeOptions;
        /// <summary>
        /// Sets the x-frame-option, which can be used to prevent pages from being included by non-origin iframes. More information can be found in the [RFC7034](https://tools.ietf.org/html/rfc7034)
        /// </summary>
        public readonly string? XFrameOptions;
        /// <summary>
        /// Prevent pages from appearing in search engines.
        /// </summary>
        public readonly string? XRobotsTag;
        /// <summary>
        /// This header configures the Cross-site scripting (XSS) filter in your browser.
        /// </summary>
        public readonly string? XXssProtection;

        [OutputConstructor]
        private RealmSecurityDefensesHeaders(
            string? contentSecurityPolicy,

            string? contentSecurityPolicyReportOnly,

            string? strictTransportSecurity,

            string? xContentTypeOptions,

            string? xFrameOptions,

            string? xRobotsTag,

            string? xXssProtection)
        {
            ContentSecurityPolicy = contentSecurityPolicy;
            ContentSecurityPolicyReportOnly = contentSecurityPolicyReportOnly;
            StrictTransportSecurity = strictTransportSecurity;
            XContentTypeOptions = xContentTypeOptions;
            XFrameOptions = xFrameOptions;
            XRobotsTag = xRobotsTag;
            XXssProtection = xXssProtection;
        }
    }
}
