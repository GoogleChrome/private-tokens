# Private Verification Tokens Issuer Registration

This readme describes the process of how origins (websites) can register as
issuers of Private Verification Tokens (PVTs). Registration is
required for any page that wishes to use PVTs. To apply to become an issuer the
issuer website’s operator must open a new issue on this repository using the
“New PVT Issuer” template which specifies:

* Organization/Company Name: The entity operating the origin.
* Contact Email: An email address for the PVT team to communicate regarding the registration.
* Key Commitment URL: The full URL for the key commitment endpoint.
* Batch Size: Chrome will create token requests in batch using the batch size
  specified.
* Expiry: Timestamp when the registration expires. In seconds since the unix
  epoch.
* Intended Use Case: A clear description of how PVTs will be used on the
  registered origin (e.g., "To reduce captchas and friction for users navigating
  to the page by verifying humanness").
* Acknowledgement: Confirmation that the origin understands and will adhere to PVT policies,
  including not using PVTs for user tracking  or cross-site correlation.

Key commitment endpoint requirements must align with and extend the Privacy Pass Issuer Configuration described in [RFC 9578 Section 4](https://www.rfc-editor.org/rfc/rfc9578.html#name-configuration).

The key commitment endpoint must return an `application/private-token-issuer-directory` HTTP 200 response containing an Issuer Directory JSON object:

* `issuer-request-uri` (string): The token request URL for generating access tokens (as an absolute URL or relative to the directory object).
* `token-keys` (array of objects): A list of Issuer Public Keys for the issuance protocol. Each key object contains:
  * `token-type` (integer): The token type identifier (as defined in [RFC 9578 Section 8.2](https://www.rfc-editor.org/rfc/rfc9578.html#section-8.2)).
  * `token-key` (string): Base64url-encoded public key (per [RFC 4648](https://datatracker.ietf.org/doc/html/rfc4648), including padding) for use with the issuance protocol.
  * `not-before` (integer, optional): UNIX timestamp in seconds since January 1, 1970 UTC indicating when the key can be used.
* `version` (integer): Indicates the PVT version (must fit in int32).
* `batch-size` (integer): The batch size for token requests. The browser will create token requests in batches using this size.
* `redeemer-origins` (array of strings): A list of valid web origins (HTTPS scheme required) permitted to redeem tokens issued by this origin.

Example response:

```json
{
  "issuer-request-uri": "https://issuer.example.net/request",
  "token-keys": [
    {
      "token-type": 2,
      "token-key": "AyuAAk7oGNcJGWeAqEr_4IeJ9XFSn8zBrM4H7qLfL8ZfA19qbrhL6pwTYRFUar2GQ8R8O0PlPp56h5a6G5JNCU4Dt_Ft8K2Cy9i9agTtQnEHrdWj1LqEDps0Gju6wdm3_hk=",
      "not-before": 1770000000
    }
  ],
  "version": 1,
  "batch-size": 10,
  "redeemer-origins": [
    "https://example.com"
  ]
}
```

Key expiration and key rotation should be accomplished through a mixture of HTTP `Cache-Control` headers and `not-before` timestamps as described in [RFC 9578 Section 4](https://www.rfc-editor.org/rfc/rfc9578.html#name-configuration):
* **HTTP `Cache-Control` Headers**: Issuers use HTTP cache directives (e.g., `Cache-Control: max-age=86400`) to set the caching lifetime of the directory resource, governing how frequently clients revalidate and fetch updated configurations.
* **`not-before` Timestamps**: Issuers can stage future keys in `token-keys` using `not-before` UNIX timestamps. Clients will refrain from using a key until its `not-before` timestamp has passed, enabling smooth scheduled key rotations.

Once the response from key commitment endpoint has been verified (to make sure
that the endpoint responds with an appropriate JSON directory object), it will be
merged into this repository and Chrome server-side infrastructure will begin
fetching those keys roughly at an hourly rate and eventually distributing those
keys to Chrome instances. Key commitments are only allowed to change every 60
days, and any rotation faster than that will be ignored.

Request Template:

Organization/Company Name: {CompanyName}
Contact Email: {Email}
Key Commitment URL: {KeyCommitmentURL}
Batch Size: {BatchSize}
Expiry: {ExpiryInSecondsSinceUnixEpoch}
Intended Use Case: {UseCaseDescription}
Acknowledgement:

By registering as an issuer, you acknowledge the following:
* I understand the technical restrictions on key rotation frequency of 60 days.
* I understand that my issuer registration will be valid for a period of six months after the key commitment is accepted, and that I will need to re-register in this repository following that six-month period.
* I understand that in the future renewing my registration may have additional requirements, to reduce the risk of abuse.
