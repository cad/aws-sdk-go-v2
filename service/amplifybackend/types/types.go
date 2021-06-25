// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The authentication settings for accessing provisioned data models in your
// Amplify project.
type BackendAPIAppSyncAuthSettings struct {

	// The Amazon Cognito user pool ID, if Amazon Cognito was used as an authentication
	// setting to access your data models.
	CognitoUserPoolId *string

	// The API key description for API_KEY, if it was used as an authentication
	// mechanism to access your data models.
	Description *string

	// The API key expiration time for API_KEY, if it was used as an authentication
	// mechanism to access your data models.
	ExpirationTime float64

	// The expiry time for the OpenID authentication mechanism.
	OpenIDAuthTTL *string

	// The clientID for openID, if openID was used as an authentication setting to
	// access your data models.
	OpenIDClientId *string

	// The expiry time for the OpenID authentication mechanism.
	OpenIDIatTTL *string

	// The openID issuer URL, if openID was used as an authentication setting to access
	// your data models.
	OpenIDIssueURL *string

	// The OpenID provider name, if OpenID was used as an authentication mechanism to
	// access your data models.
	OpenIDProviderName *string
}

// Describes the auth types for your configured data models.
type BackendAPIAuthType struct {

	// Describes the authentication mode.
	Mode Mode

	// Describes settings for the authentication mode.
	Settings *BackendAPIAppSyncAuthSettings
}

// Describes the conflict resolution configuration for your data model configured
// in your Amplify project.
type BackendAPIConflictResolution struct {

	// The strategy for conflict resolution.
	ResolutionStrategy ResolutionStrategy
}

// The resource config for the data model, configured as a part of the Amplify
// project.
type BackendAPIResourceConfig struct {

	// Additional authentication methods used to interact with your data models.
	AdditionalAuthTypes []BackendAPIAuthType

	// The API name used to interact with the data model, configured as a part of the
	// amplify project.
	ApiName *string

	// The conflict resolution strategy for your data stored in the data models.
	ConflictResolution *BackendAPIConflictResolution

	// The default authentication type for interacting with the configured data models
	// in your Amplify project.
	DefaultAuthType *BackendAPIAuthType

	// The service used to provision and interact with the data model.
	Service *string

	// The definition of the data model in the annotated transform of the GraphQL
	// schema.
	TransformSchema *string
}

// Describes third-party social federation configurations for allowing your app
// users to sign in using OAuth.
type BackendAuthSocialProviderConfig struct {

	// Describes the client_id which can be obtained from the third-party social
	// federation provider.
	ClientId *string

	// Describes the client_secret which can be obtained from third-party social
	// federation providers.
	ClientSecret *string
}

// The response object for this operation.
type BackendJobRespObj struct {

	// The app ID.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment.
	//
	// This member is required.
	BackendEnvironmentName *string

	// The time when the job was created.
	CreateTime *string

	// If the request fails, this error is returned.
	Error *string

	// The ID for the job.
	JobId *string

	// The name of the operation.
	Operation *string

	// The current status of the request.
	Status *string

	// The time when the job was last updated.
	UpdateTime *string
}

// Describes the forgot password policy for authenticating into the Amplify app.
type CreateBackendAuthForgotPasswordConfig struct {

	// Describes which mode to use (either SMS or email) to deliver messages to app
	// users that want to recover their password.
	//
	// This member is required.
	DeliveryMethod DeliveryMethod

	// The configuration for the email sent when an app user forgets their password.
	EmailSettings *EmailSettings

	// The configuration for the SMS message sent when an app user forgets their
	// password.
	SmsSettings *SmsSettings
}

// Describes authorization configurations for the auth resources, configured as a
// part of your Amplify project.
type CreateBackendAuthIdentityPoolConfig struct {

	// Name of the Amazon Cognito identity pool used for authorization.
	//
	// This member is required.
	IdentityPoolName *string

	// Set to true or false based on whether you want to enable guest authorization to
	// your Amplify app.
	//
	// This member is required.
	UnauthenticatedLogin bool
}

// Describes whether multi-factor authentication policies should be applied for
// your Amazon Cognito user pool configured as a part of your Amplify project.
type CreateBackendAuthMFAConfig struct {

	// Describes whether MFA should be [ON, OFF, OPTIONAL] for authentication in your
	// Amplify project.
	//
	// This member is required.
	MFAMode MFAMode

	// Describes the configuration settings and methods for your Amplify app users to
	// use MFA.
	Settings *Settings
}

// Creates the OAuth configuration for your Amplify project.
type CreateBackendAuthOAuthConfig struct {

	// The OAuth grant type that you use to allow app users to authenticate from your
	// Amplify app.
	//
	// This member is required.
	OAuthGrantType OAuthGrantType

	// List of OAuth-related flows used to allow your app users to authenticate from
	// your Amplify app.
	//
	// This member is required.
	OAuthScopes []OAuthScopesElement

	// The redirected URI for signing in to your Amplify app.
	//
	// This member is required.
	RedirectSignInURIs []string

	// Redirect URLs used by OAuth when a user signs out of an Amplify app.
	//
	// This member is required.
	RedirectSignOutURIs []string

	// The domain prefix for your Amplify app.
	DomainPrefix *string

	// The settings for using social providers to access your Amplify app.
	SocialProviderSettings *SocialProviderSettings
}

// The password policy configuration for the backend to your Amplify project.
type CreateBackendAuthPasswordPolicyConfig struct {

	// The minimum length of the password used to access the backend of your Amplify
	// project.
	//
	// This member is required.
	MinimumLength float64

	// Additional constraints for the password used to access the backend of your
	// Amplify project.
	AdditionalConstraints []AdditionalConstraintsElement
}

// Defines the resource configuration when creating an auth resource in your
// Amplify project.
type CreateBackendAuthResourceConfig struct {

	// Defines whether you want to configure only authentication or both authentication
	// and authorization settings.
	//
	// This member is required.
	AuthResources AuthResources

	// Defines the service name to use when configuring an authentication resource in
	// your Amplify project.
	//
	// This member is required.
	Service Service

	// Describes authentication configuration for the Amazon Cognito user pool,
	// provisioned as a part of your auth resource in the Amplify project.
	//
	// This member is required.
	UserPoolConfigs *CreateBackendAuthUserPoolConfig

	// Describes the authorization configuration for the Amazon Cognito identity pool,
	// provisioned as a part of your auth resource in the Amplify project.
	IdentityPoolConfigs *CreateBackendAuthIdentityPoolConfig
}

// Describes the Amazon Cognito user pool configuration for the auth resource to be
// configured for your Amplify project.
type CreateBackendAuthUserPoolConfig struct {

	// The required attributes to sign up new users in the user pool.
	//
	// This member is required.
	RequiredSignUpAttributes []RequiredSignUpAttributesElement

	// Describes the sign-in methods that your Amplify app users use to log in using
	// the Amazon Cognito user pool, configured as a part of your Amplify project.
	//
	// This member is required.
	SignInMethod SignInMethod

	// The Amazon Cognito user pool name.
	//
	// This member is required.
	UserPoolName *string

	// Describes the forgotten password policy for your Amazon Cognito user pool,
	// configured as a part of your Amplify project.
	ForgotPassword *CreateBackendAuthForgotPasswordConfig

	// Describes whether multi-factor authentication policies should be applied for
	// your Amazon Cognito user pool configured as a part of your Amplify project.
	Mfa *CreateBackendAuthMFAConfig

	// Describes the OAuth policy and rules for your Amazon Cognito user pool,
	// configured as a part of your Amplify project.
	OAuth *CreateBackendAuthOAuthConfig

	// Describes the password policy for your Amazon Cognito user pool, configured as a
	// part of your Amplify project.
	PasswordPolicy *CreateBackendAuthPasswordPolicyConfig
}

// The configuration for the email sent when an app user forgets their password.
type EmailSettings struct {

	// The body of the email.
	EmailMessage *string

	// The subject of the email.
	EmailSubject *string
}

// The request object for this operation.
type LoginAuthConfigReqObj struct {

	// The Amazon Cognito identity pool ID used for the Amplify Admin UI login
	// authorization.
	AwsCognitoIdentityPoolId *string

	// The AWS Region for the Amplify Admin UI login.
	AwsCognitoRegion *string

	// The Amazon Cognito user pool ID used for Amplify Admin UI login authentication.
	AwsUserPoolsId *string

	// The web client ID for the Amazon Cognito user pools.
	AwsUserPoolsWebClientId *string
}

// Defines the resource configuration for the data model in your Amplify project.
type ResourceConfig struct {
}

// The settings of your MFA configuration for the backend of your Amplify project.
type Settings struct {

	// The supported MFA types.
	MfaTypes []MfaTypesElement

	// The body of the SMS message.
	SmsMessage *string
}

// SMS settings for authentication.
type SmsSettings struct {

	// The body of the SMS message.
	SmsMessage *string
}

// The settings for using the social identity providers for access to your Amplify
// app.
type SocialProviderSettings struct {

	// Describes third-party social federation configurations for allowing your app
	// users to sign in using OAuth.
	Facebook *BackendAuthSocialProviderConfig

	// Describes third-party social federation configurations for allowing your app
	// users to sign in using OAuth.
	Google *BackendAuthSocialProviderConfig

	// Describes third-party social federation configurations for allowing your app
	// users to sign in using OAuth.
	LoginWithAmazon *BackendAuthSocialProviderConfig
}

// Describes the forgot password policy for authenticating into the Amplify app.
type UpdateBackendAuthForgotPasswordConfig struct {

	// Describes which mode to use (either SMS or email) to deliver messages to app
	// users that want to recover their password.
	DeliveryMethod DeliveryMethod

	// The configuration for the email sent when an app user forgets their password.
	EmailSettings *EmailSettings

	// The configuration for the SMS message sent when an Amplify app user forgets
	// their password.
	SmsSettings *SmsSettings
}

// Describes the authorization configuration for the Amazon Cognito identity pool,
// provisioned as a part of your auth resource in the Amplify project.
type UpdateBackendAuthIdentityPoolConfig struct {

	// A boolean value which can be set to allow or disallow guest-level authorization
	// into your Amplify app.
	UnauthenticatedLogin bool
}

// Updates the multi-factor authentication (MFA) configuration for the backend of
// your Amplify project.
type UpdateBackendAuthMFAConfig struct {

	// The MFA mode for the backend of your Amplify project.
	MFAMode MFAMode

	// The settings of your MFA configuration for the backend of your Amplify project.
	Settings *Settings
}

// The OAuth configurations for authenticating users into your Amplify app.
type UpdateBackendAuthOAuthConfig struct {

	// The Amazon Cognito domain prefix used to create a hosted UI for authentication.
	DomainPrefix *string

	// The OAuth grant type to allow app users to authenticate from your Amplify app.
	OAuthGrantType OAuthGrantType

	// The list of OAuth-related flows that can allow users to authenticate from your
	// Amplify app.
	OAuthScopes []OAuthScopesElement

	// Redirect URLs used by OAuth when a user signs in to an Amplify app.
	RedirectSignInURIs []string

	// Redirect URLs used by OAuth when a user signs out of an Amplify app.
	RedirectSignOutURIs []string

	// Describes third-party social federation configurations for allowing your users
	// to sign in with OAuth.
	SocialProviderSettings *SocialProviderSettings
}

// Describes the password policy for your Amazon Cognito user pool configured as a
// part of your Amplify project.
type UpdateBackendAuthPasswordPolicyConfig struct {

	// Describes additional constraints on password requirements to sign in to the auth
	// resource, configured as a part of your Amplify project.
	AdditionalConstraints []AdditionalConstraintsElement

	// Describes the minimum length of the password required to sign in to the auth
	// resource, configured as a part of your Amplify project.
	MinimumLength float64
}

// Defines the resource configuration when updating an authentication resource in
// your Amplify project.
type UpdateBackendAuthResourceConfig struct {

	// Defines the service name to use when configuring an authentication resource in
	// your Amplify project.
	//
	// This member is required.
	AuthResources AuthResources

	// Defines the service name to use when configuring an authentication resource in
	// your Amplify project.
	//
	// This member is required.
	Service Service

	// Describes the authentication configuration for the Amazon Cognito user pool,
	// provisioned as a part of your auth resource in the Amplify project.
	//
	// This member is required.
	UserPoolConfigs *UpdateBackendAuthUserPoolConfig

	// Describes the authorization configuration for the Amazon Cognito identity pool,
	// provisioned as a part of your auth resource in the Amplify project.
	IdentityPoolConfigs *UpdateBackendAuthIdentityPoolConfig
}

// Describes the Amazon Cognito user pool configuration for the authorization
// resource to be configured for your Amplify project on an update.
type UpdateBackendAuthUserPoolConfig struct {

	// Describes the forgot password policy for your Amazon Cognito user pool,
	// configured as a part of your Amplify project.
	ForgotPassword *UpdateBackendAuthForgotPasswordConfig

	// Describes whether multi-factor authentication policies should be applied for
	// your Amazon Cognito user pool configured as a part of your Amplify project.
	Mfa *UpdateBackendAuthMFAConfig

	// Describes the OAuth policy and rules for your Amazon Cognito user pool,
	// configured as a part of your Amplify project.
	OAuth *UpdateBackendAuthOAuthConfig

	// Describes the password policy for your Amazon Cognito user pool, configured as a
	// part of your Amplify project.
	PasswordPolicy *UpdateBackendAuthPasswordPolicyConfig
}
