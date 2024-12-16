# Private State Tokens Issuer Registration

With the [Private State Tokens](https://developer.chrome.com/en/docs/privacy-sandbox/trust-tokens/) API, a limited amount of information can be conveyed from one browsing context to another (for example, across sites) to help combat fraud, without passive tracking. This readme describes the process of how websites can issue Private State Tokens.
To apply for an issuer (and its key commitments) to be included within Chrome, the issuer website’s operator must open a new [issue](https://github.com/GoogleChrome/private-tokens/issues/new) on [this repository](https://github.com/GoogleChrome/private-tokens) using the “New PST Issuer” template which specifies:

*   Issuer name - Human-readable name representing the issuer website.
*   Product name - Human-readable name describing the specific product or service that will issue Private State Tokens
*   Privacy Policy - Link to the issuer’s Privacy Policy governing the issuance of Private State Tokens.
*   The [origin](https://developer.mozilla.org/en-US/docs/Glossary/Origin) that is used for the issuing service (scheme, host, port). The scheme must be HTTPS.
*   An email or email alias that is monitored by the issuer’s operator for issues regarding the issuer website.
*   A public HTTPS endpoint that responds to key commitment requests as described in the [specification](https://wicg.github.io/trust-token-api/#issuer-public-keys).
*   Purpose - A selection of common purposes for issuing private state tokens. Organizations should select one or more purposes that match their use case and include in their registration. If selecting “other”, organizations should describe their use case in the submission and also consider opening an issue on this repository describing your use case, to enable future improvements to this list. 
*   A description of the issuer, describing the intended purpose of the tokens.
*   Acceptance of the [current Disclosure and Acknowledgement](#disclosure-and-acknowledgement) to be an issuer website.

In order to help sites select the most appropriate purpose as part of issuer registration, the following describes the various purposes in greater detail:
*  Account safety
    *  Account Safety covers a range of account-based protections, including safeguarding users from account takeover attempts as well as protecting sites from account creation abuse.
*  Invalid Traffic Detection in Advertising
    *  Invalid Traffic (IVT) describes any activity that doesn’t come from a real user with genuine interest. This purpose covers IVT detection specifically for advertising use cases.
*  Invalid Traffic Detection for non-advertising use cases
    *  Invalid Traffic (IVT) describes any activity that doesn’t come from a real user with genuine interest. This purpose covers IVT detection for non-advertising use cases.
*   Payment & e-commerce protections
    *  These use cases cover payment and e-commerce protections – safeguarding users and sites against automated or manual techniques to abuse online payment workflows.
*  Bulk content abuse (e.g. Spam) prevention
    * Bulk content and spam prevention helps sites differentiate from organic and non-organic content and content interactions, including preventing spam in user generated content submissions.
*  Other, please specify: {description}
    *  If none of the above selections describe the site’s purpose for issuing Private State Tokens, select this option. Issuers should also describe their use case and consider opening an issue on this repository describing the use case in greater detail, to enable future improvements to this list.


Once an endpoint has been verified (to make sure that the endpoint responds with an [appropriate JSON dictionary](https://wicg.github.io/trust-token-api/#issuer-public-keys)), it will be merged into this repository and Chrome server-side infrastructure will begin fetching those keys roughly at an hourly rate and eventually distributing those keys to Chrome instances. Key commitments are only allowed to change every 60 days, and any rotation faster than that will be ignored.

Issuer entries in this repository will have an expiration date 6 months after the issuer request has been filed, issuers will need to refile before the 6 months elapse to ensure that Chrome continues to receive their key commitments. This expiration is to ensure that issuers are still active and acknowledge the latest specification and PST versions. Expired issuer entries will result in invalid tokens. To avoid interruption, it is recommended that you mark your calendars to refresh your configuration 2 weeks before it expires.

If there’s an emergency key loss/compromise, issuers will need to set a key of “emergency” in their key commitment. This may happen in the event that their company has been hacked and their keys have been compromised. To perform another “emergency” rotation, they will be required to first publicize a notification about the incident by opening a new issue with the appropriate issue template on this repository (this requirement is inspired  by the incident notification requirements in the [CA ecosystem](https://www.mozilla.org/en-US/about/governance/policies/security-group/certs/policy/)), on why they needed to do an emergency key rotation. Issuers who have not yet issued an incident notification and had it acknowledged may continue performing normal frequency key rotations but will be unable to perform additional “emergency” rotations. Key rotations out of the normal cycle can allow malicious issuers to gain additional information about a user based on the timing of their key rotations, so incident notification requirement is intended to prevent issuers from abusing the mechanism and rotating keys too frequently.

### Disclosure and Acknowledgement
By registering as an issuer, you acknowledge the following:

1. I understand the technical restrictions on key rotation frequency of 60 days in the PST API.
2. I understand that my issuer registration will be valid for a period of six months after the key commitment is accepted, and that I will need to re-register in this repository following that six-month period.
3. I understand that in the future renewing my registration for this API may have additional requirements, to reduce the risk of abuse by token issuers.

---
Request Template:

**Summary:** Private State Token Issuer Request - {Issuer Name}

**Description:**

{Issuer Name}

[Origin](https://developer.mozilla.org/en-US/docs/Glossary/Origin): {Origin}

Contact: {Contact Email}

Key Commitment Endpoint URL: {Issuer Key Commitment Endpoint}

Purpose: {Issuer purpose}

Disclosure & Acknowledgement:

1. I understand the technical restrictions on key rotation frequency of 60 days in the PST API.
2. I understand that my issuer registration will be valid for a period of six months after the key commitment is accepted, and that I will need to re-register in this repository following that six-month period.
3. I understand that in the future renewing my registration for this API may have additional requirements, to reduce the risk of abuse by token issuers.
