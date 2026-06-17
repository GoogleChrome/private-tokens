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

Key commintment endpoint requirements are documented at
[Key Commtiment Endpoint Requirements](https://github.com/GoogleChrome/private-tokens/blob/mainbug/README.md#key-commitment-endpoint-requirements) section.

Once the response from key commitment endpoint has been verified (to make sure
that the endpoint responds with an appropriate JSON dictionary), it will be
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
