# Private Tokens

This repository is used to keep record of and register issuers and their configurations for Private Token APIs for use in Chrome.

The supported APIs and the process for registering are listed below:

## Private State Tokens

For instructions on how to register a Private State Token issuer, please read [Private State Tokens Issuer Registration](https://github.com/GoogleChrome/private-tokens/blob/main/PST-Registration.md).

For clarity on the Private State Token API, please read the [Private State Tokens explainer](https://github.com/WICG/trust-token-api/).

## Private Verification Tokens

PVTs work for the registered origins. Registration process is explained in
[PVT-Registration.md](https://github.com/GoogleChrome/private-tokens/blob/main/PVT-Registration.md).

### Key Commtiment Endpoint Requirements

Issuers specify a key commitment endpoint during registration. The key commitment
endpoint must return a response in the following form.

```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: <Length of response>

<JSON response>
```

where response is a [JSON object](https://datatracker.ietf.org/doc/html/rfc7159#section-4)
with the following name/value pairs.

* Name `domain`, value [string](https://datatracker.ietf.org/doc/html/rfc7159#section-7).
  The string must be a registrable domain of the `issuer_origin`.
* Name `issuer_origin`, value string. The value must be a valid URL. The URL must have https
  scheme. The URL must have same registrable domain (eTLD+1) as the one specified in `domain`.
* Name `version`, value must be an integer. Indicates the PVT version. The value must fit
  int32. The browser will ignore if the version is not supported.
* Name `public_key`, value base64 encoding of the public key. Public key is parsed based on
  the crypto parameters deduced from the value of the `version`.
* Name `key_id`, value integer. Key id must fit into uint8. Key id will be used in token
  requests as specified in the privacy pass specs.
* Name `expiration`, value *string*. Expiration must fit into int64. Expiration is
  in number of seconds since the unix epoch. The browser will stop using the `public_key` past
  expiration date.
* Name `batch_size`, value integer. Browser will send token requests in batches. Each 
  HTTP request to issuance endpoint will contain `batch_size` many individual
  [TokenRequest](https://github.com/cathieyun/draft-athm/blob/main/draft-yun-privacypass-athm.md#client-to-issuer-request).
* Name `redeemer_origins`, value [array](https://datatracker.ietf.org/doc/html/rfc7159#section-5)
  of strings. Strings must be valid web origins. The scheme must be https.

For an example see the demo key commitment endpoint
[https://privatetokens.dev/.well-known/private-verification-token/key-commitment](https://privatetokens.dev/.well-known/private-verification-token/key-commitment), which returns (on June 8th 2026)

```
{
"domain": "privatetokens.dev",
"issuer_origin": "https://pvtissuer.privatetokens.dev",
"version": 1,
"public_key": "AyuAAk7oGNcJGWeAqEr/4IeJ9XFSn8zBrM4H7qLfL8ZfA19qbrhL6pwTYRFUar2GQ8R8O0PlPp56h5a6G5JNCU4Dt/Ft8K2Cy9i9agTtQnEHrdWj1LqEDps0Gju6wdm3/hk=",
"key_id": 3,
"expiration": "184368811",
"batch_size": 10,
"redeemer_origins": ["https://privatetokens.dev"]
}
```
