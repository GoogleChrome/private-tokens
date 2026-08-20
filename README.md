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

* Name `issuerRequestUri`, value string. The value must be a valid URL. The URL must have https
  scheme.
* Name `version`, value must be an integer. Indicates the PVT version. The value must fit
  int32. The browser will ignore if the version is not supported.
* Name `publicKey`, value base64 encoding of the public key. Public key is parsed based on
  the crypto parameters deduced from the value of the `version`.
* Name `publicKeyProof`, value base64 encoding of the public key proof. Public key proof is
  parsed based on the crypto parameters deduced from the value of the `version`.
* Name `expiration`, value *string*. Expiration must fit into int64. Expiration is
  in number of seconds since the unix epoch. The browser will stop using the `public_key` past
  expiration date.
* Name `batchSize`, value integer. Browser will send token requests in batches. Each 
  HTTP request to issuance endpoint will contain `batchSize` many individual
  [TokenRequest](https://github.com/cathieyun/draft-athm/blob/main/draft-yun-privacypass-athm.md#client-to-issuer-request)s.
* Name `redeemerOrigins`, value [array](https://datatracker.ietf.org/doc/html/rfc7159#section-5)
  of strings. Strings must be valid web origins. The scheme must be https.

For an example see the demo key commitment endpoint
[https://privatetokens.dev/.well-known/private-verification-token/key-commitment](https://privatetokens.dev/.well-known/private-verification-token/key-commitment), which returns (on August 20th 2026)

```
{
"issuerRequestUri": "https://privatetokens.dev/.well-known/private-verification-token/issue",
"version": 1,
"publicKey": "AgHPcLgHpe/ASNDfgaOp7gyvULDweWsAw5L1i2wMvi6FAoIgHpg69O7qX35d3rBH+TgWTT/WWnNInST0chCpLqBFA6Ib8wfRFpEUX2VaoJhn+u8n5AEPpLfixlwsl841kK1l",
"publicKeyProof": "trewu/RhAwtPh5gfxo3wC7dbDoWcScC69qrBf28caNrdAL8x5jcYW8ol6A26OiGVsqQB2XoAPeQHtvYxN+kgqw==",
"expiration": "1794854079",
"batchSize": 10,
"redeemerOrigins": ["https://privatetokens.dev"]
}
```
