# Casdoor
This section describes how to configure Casdoor as an authentication provider for Nginx UI, contributed by @Jraaay.

Casdoor is a powerful and comprehensive identity authentication solution that supports OAuth 2.0, SAML 2.0, LDAP, AD, and multiple social login methods.
By integrating Casdoor, Nginx UI can leverage these features to improve security and user experience.

## Endpoint
- Type: `string`

This is the Endpoint of the Casdoor server. You need to make sure that Nginx UI can access this URL.

## ClientId
- Type: `string`

This is the Client ID generated by Casdoor for your application.
It is used to identify your application during the authentication process.

## ClientSecret
- Type: `string`

This is the Client Secret generated by Casdoor for your application.
It is necessary to keep your application secure.

## Certificate
- Type: `string`

This is the certificate used during the authentication process.
Make sure it is valid and trusted.

## Organization
- Type: `string`

This is the organization name you set in Casdoor.
It will use this information to process authentication requests.

## Application
- Type: `string`

This is the application name you created in Casdoor.

## RedirectUri
- Type: `string`

This is the URI that users will be redirected to after successful login or authorization.
It should be consistent with the Redirect URI in the Casdoor application configuration.