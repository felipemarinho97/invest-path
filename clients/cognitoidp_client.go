package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// ICognitoIDPClient generic client
type ICognitoIDPClient interface {
	// Adds additional user attributes to the user pool schema.
	//
	AddCustomAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.AddCustomAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	// Adds the specified user to the specified group. Calling this action requires
	// developer credentials.
	//
	AdminAddUserToGroup(arg1 context.Context, arg2 *cognitoidentityprovider.AdminAddUserToGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	// Confirms user registration as an admin without using a confirmation code. Works
	// on any user. Calling this action requires developer credentials.
	//
	AdminConfirmSignUp(arg1 context.Context, arg2 *cognitoidentityprovider.AdminConfirmSignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	// Creates a new user in the specified user pool. If MessageAction is not set, the
	// default is to send a welcome message via email or phone (SMS). This message is
	// based on a template that you configured in your call to create or update a user
	// pool. This template includes your custom sign-up instructions and placeholders
	// for user name and temporary password. Alternatively, you can call
	// AdminCreateUser with “SUPPRESS” for the MessageAction parameter, and Amazon
	// Cognito will not send any email. In either case, the user will be in the
	// FORCE_CHANGE_PASSWORD state until they sign in and change their password.
	// AdminCreateUser requires developer credentials.
	//
	AdminCreateUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminCreateUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	// Deletes a user as an administrator. Works on any user. Calling this action
	// requires developer credentials.
	//
	AdminDeleteUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDeleteUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	// Deletes the user attributes in a user pool as an administrator. Works on any
	// user. Calling this action requires developer credentials.
	//
	AdminDeleteUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDeleteUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	// Disables the user from signing in with the specified external (SAML or social)
	// identity provider. If the user to disable is a Cognito User Pools native
	// username + password user, they are not permitted to use their password to
	// sign-in. If the user to disable is a linked external IdP user, any link between
	// that user and an existing user is removed. The next time the external user (no
	// longer attached to the previously linked DestinationUser) signs in, they must
	// create a new user account. See AdminLinkProviderForUser
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminLinkProviderForUser.html).
	// This action is enabled only for admin access and requires developer credentials.
	// The ProviderName must match the value specified when creating an IdP for the
	// pool. To disable a native username + password user, the ProviderName value must
	// be Cognito and the ProviderAttributeName must be Cognito_Subject, with the
	// ProviderAttributeValue being the name that is used in the user pool for the
	// user. The ProviderAttributeName must always be Cognito_Subject for social
	// identity providers. The ProviderAttributeValue must always be the exact subject
	// that was used when the user was originally linked as a source user. For
	// de-linking a SAML identity, there are two scenarios. If the linked identity has
	// not yet been used to sign-in, the ProviderAttributeName and
	// ProviderAttributeValue must be the same values that were used for the SourceUser
	// when the identities were originally linked using  AdminLinkProviderForUser call.
	// (If the linking was done with ProviderAttributeName set to Cognito_Subject, the
	// same applies here). However, if the user has already signed in, the
	// ProviderAttributeName must be Cognito_Subject and ProviderAttributeValue must be
	// the subject of the SAML assertion.
	//
	AdminDisableProviderForUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDisableProviderForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error)
	// Disables the specified user. Calling this action requires developer credentials.
	//
	AdminDisableUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDisableUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	// Enables the specified user as an administrator. Works on any user. Calling this
	// action requires developer credentials.
	//
	AdminEnableUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminEnableUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	// Forgets the device, as an administrator. Calling this action requires developer
	// credentials.
	//
	AdminForgetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.AdminForgetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	// Gets the device, as an administrator. Calling this action requires developer
	// credentials.
	//
	AdminGetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.AdminGetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	// Gets the specified user by user name in a user pool as an administrator. Works
	// on any user. Calling this action requires developer credentials.
	//
	AdminGetUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminGetUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetUserOutput, error)
	// Initiates the authentication flow, as an administrator. Calling this action
	// requires developer credentials.
	//
	AdminInitiateAuth(arg1 context.Context, arg2 *cognitoidentityprovider.AdminInitiateAuthInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	// Links an existing user account in a user pool (DestinationUser) to an identity
	// from an external identity provider (SourceUser) based on a specified attribute
	// name and value from the external identity provider. This allows you to create a
	// link from the existing user account to an external federated user identity that
	// has not yet been used to sign in, so that the federated user identity can be
	// used to sign in as the existing user account. For example, if there is an
	// existing user with a username and password, this API links that user to a
	// federated user identity, so that when the federated user identity is used, the
	// user signs in as the existing user account. The maximum number of federated
	// identities linked to a user is 5. Because this API allows a user with an
	// external federated identity to sign in as an existing user in the user pool, it
	// is critical that it only be used with external identity providers and provider
	// attributes that have been trusted by the application owner. This action is
	// enabled only for admin access and requires developer credentials.
	//
	AdminLinkProviderForUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminLinkProviderForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error)
	// Lists devices, as an administrator. Calling this action requires developer
	// credentials.
	//
	AdminListDevices(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListDevicesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	// Lists the groups that the user belongs to. Calling this action requires
	// developer credentials.
	//
	AdminListGroupsForUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListGroupsForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	// Lists a history of user activity and any risks detected as part of Amazon
	// Cognito advanced security.
	//
	AdminListUserAuthEvents(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListUserAuthEventsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error)
	// Removes the specified user from the specified group. Calling this action
	// requires developer credentials.
	//
	AdminRemoveUserFromGroup(arg1 context.Context, arg2 *cognitoidentityprovider.AdminRemoveUserFromGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	// Resets the specified user's password in a user pool as an administrator. Works
	// on any user. When a developer calls this API, the current password is
	// invalidated, so it must be changed. If a user tries to sign in after the API is
	// called, the app will get a PasswordResetRequiredException exception back and
	// should direct the user down the flow to reset the password, which is the same as
	// the forgot password flow. In addition, if the user pool has phone verification
	// selected and a verified phone number exists for the user, or if email
	// verification is selected and a verified email exists for the user, calling this
	// API will also result in sending a message to the end user with the code to
	// change their password. Calling this action requires developer credentials.
	//
	AdminResetUserPassword(arg1 context.Context, arg2 *cognitoidentityprovider.AdminResetUserPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	// Responds to an authentication challenge, as an administrator. Calling this
	// action requires developer credentials.
	//
	AdminRespondToAuthChallenge(arg1 context.Context, arg2 *cognitoidentityprovider.AdminRespondToAuthChallengeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	// Sets the user's multi-factor authentication (MFA) preference, including which
	// MFA options are enabled and if any are preferred. Only one factor can be set as
	// preferred. The preferred MFA factor will be used to authenticate a user if
	// multiple factors are enabled. If multiple options are enabled and no preference
	// is set, a challenge to choose an MFA option will be returned during sign in.
	//
	AdminSetUserMFAPreference(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserMFAPreferenceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error)
	// Sets the specified user's password in a user pool as an administrator. Works on
	// any user. The password can be temporary or permanent. If it is temporary, the
	// user status will be placed into the FORCE_CHANGE_PASSWORD state. When the user
	// next tries to sign in, the InitiateAuth/AdminInitiateAuth response will contain
	// the NEW_PASSWORD_REQUIRED challenge. If the user does not sign in before it
	// expires, the user will not be able to sign in and their password will need to be
	// reset by an administrator. Once the user has set a new password, or the password
	// is permanent, the user status will be set to Confirmed.
	//
	AdminSetUserPassword(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error)
	// This action is no longer supported. You can use it to configure only SMS MFA.
	// You can't use it to configure TOTP software token MFA. To configure either type
	// of MFA, use AdminSetUserMFAPreference
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminSetUserMFAPreference.html)
	// instead.
	//
	AdminSetUserSettings(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserSettingsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	// Provides feedback for an authentication event as to whether it was from a valid
	// user. This feedback is used for improving the risk evaluation decision for the
	// user pool as part of Amazon Cognito advanced security.
	//
	AdminUpdateAuthEventFeedback(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error)
	// Updates the device status as an administrator. Calling this action requires
	// developer credentials.
	//
	AdminUpdateDeviceStatus(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateDeviceStatusInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	// Updates the specified user's attributes, including developer attributes, as an
	// administrator. Works on any user. For custom attributes, you must prepend the
	// custom: prefix to the attribute name. In addition to updating user attributes,
	// this API can also be used to mark phone and email as verified. Calling this
	// action requires developer credentials.
	//
	AdminUpdateUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	// Signs out users from all devices, as an administrator. It also invalidates all
	// refresh tokens issued to a user. The user's current access and Id tokens remain
	// valid until their expiry. Access and Id tokens expire one hour after they are
	// issued. Calling this action requires developer credentials.
	//
	AdminUserGlobalSignOut(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUserGlobalSignOutInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	// Returns a unique generated shared secret key code for the user account. The
	// request takes an access token or a session string, but not both.
	//
	AssociateSoftwareToken(arg1 context.Context, arg2 *cognitoidentityprovider.AssociateSoftwareTokenInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error)
	// Changes the password for a specified user in a user pool.
	//
	ChangePassword(arg1 context.Context, arg2 *cognitoidentityprovider.ChangePasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ChangePasswordOutput, error)
	// Confirms tracking of the device. This API call is the call that begins device
	// tracking.
	//
	ConfirmDevice(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	// Allows a user to enter a confirmation code to reset a forgotten password.
	//
	ConfirmForgotPassword(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmForgotPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	// Confirms registration of a user and handles the existing alias from a previous
	// user.
	//
	ConfirmSignUp(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmSignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	// Creates a new group in the specified user pool. Calling this action requires
	// developer credentials.
	//
	CreateGroup(arg1 context.Context, arg2 *cognitoidentityprovider.CreateGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error)
	// Creates an identity provider for a user pool.
	//
	CreateIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.CreateIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateIdentityProviderOutput, error)
	// Creates a new OAuth2.0 resource server and defines custom scopes in it.
	//
	CreateResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.CreateResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateResourceServerOutput, error)
	// Creates the user import job.
	//
	CreateUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	// Creates a new Amazon Cognito user pool and sets the password policy for the
	// pool.
	//
	CreateUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	// Creates the user pool client.
	//
	CreateUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	// Creates a new domain for a user pool.
	//
	CreateUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error)
	// Deletes a group. Calling this action requires developer credentials.
	//
	DeleteGroup(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteGroupOutput, error)
	// Deletes an identity provider for a user pool.
	//
	DeleteIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error)
	// Deletes a resource server.
	//
	DeleteResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteResourceServerOutput, error)
	// Allows a user to delete himself or herself.
	//
	DeleteUser(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserOutput, error)
	// Deletes the attributes for a user.
	//
	DeleteUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	// Deletes the specified Amazon Cognito user pool.
	//
	DeleteUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	// Allows the developer to delete the user pool client.
	//
	DeleteUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	// Deletes a domain for a user pool.
	//
	DeleteUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error)
	// Gets information about a specific identity provider.
	//
	DescribeIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	// Describes a resource server.
	//
	DescribeResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	// Describes the risk configuration.
	//
	DescribeRiskConfiguration(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeRiskConfigurationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error)
	// Describes the user import job.
	//
	DescribeUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	// Returns the configuration information and metadata of the specified user pool.
	//
	DescribeUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	// Client method for returning the configuration information and metadata of the
	// specified user pool app client.
	//
	DescribeUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	// Gets information about a domain.
	//
	DescribeUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	// Forgets the specified device.
	//
	ForgetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.ForgetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	// Calling this API causes a message to be sent to the end user with a confirmation
	// code that is required to change the user's password. For the Username parameter,
	// you can use the username or user alias. The method used to send the confirmation
	// code is sent according to the specified AccountRecoverySetting. For more
	// information, see Recovering User Accounts
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-recover-a-user-account.html)
	// in the Amazon Cognito Developer Guide. If neither a verified phone number nor a
	// verified email exists, an InvalidParameterException is thrown. To use the
	// confirmation code for resetting the password, call ConfirmForgotPassword
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_ConfirmForgotPassword.html).
	//
	ForgotPassword(arg1 context.Context, arg2 *cognitoidentityprovider.ForgotPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	// Gets the header information for the .csv file to be used as input for the user
	// import job.
	//
	GetCSVHeader(arg1 context.Context, arg2 *cognitoidentityprovider.GetCSVHeaderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	// Gets the device.
	//
	GetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.GetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error)
	// Gets a group. Calling this action requires developer credentials.
	//
	GetGroup(arg1 context.Context, arg2 *cognitoidentityprovider.GetGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error)
	// Gets the specified identity provider.
	//
	GetIdentityProviderByIdentifier(arg1 context.Context, arg2 *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	// This method takes a user pool ID, and returns the signing certificate.
	//
	GetSigningCertificate(arg1 context.Context, arg2 *cognitoidentityprovider.GetSigningCertificateInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error)
	// Gets the UI Customization information for a particular app client's app UI, if
	// there is something set. If nothing is set for the particular client, but there
	// is an existing pool level customization (app clientId will be ALL), then that is
	// returned. If nothing is present, then an empty shape is returned.
	//
	GetUICustomization(arg1 context.Context, arg2 *cognitoidentityprovider.GetUICustomizationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	// Gets the user attributes and metadata for a user.
	//
	GetUser(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error)
	// Gets the user attribute verification code for the specified attribute name.
	//
	GetUserAttributeVerificationCode(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	// Gets the user pool multi-factor authentication (MFA) configuration.
	//
	GetUserPoolMfaConfig(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserPoolMfaConfigInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error)
	// Signs out users from all devices. It also invalidates all refresh tokens issued
	// to a user. The user's current access and Id tokens remain valid until their
	// expiry. Access and Id tokens expire one hour after they are issued.
	//
	GlobalSignOut(arg1 context.Context, arg2 *cognitoidentityprovider.GlobalSignOutInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	// Initiates the authentication flow.
	//
	InitiateAuth(arg1 context.Context, arg2 *cognitoidentityprovider.InitiateAuthInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.InitiateAuthOutput, error)
	// Lists the devices.
	//
	ListDevices(arg1 context.Context, arg2 *cognitoidentityprovider.ListDevicesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error)
	// Lists the groups associated with a user pool. Calling this action requires
	// developer credentials.
	//
	ListGroups(arg1 context.Context, arg2 *cognitoidentityprovider.ListGroupsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error)
	// Lists information about all identity providers for a user pool.
	//
	ListIdentityProviders(arg1 context.Context, arg2 *cognitoidentityprovider.ListIdentityProvidersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	// Lists the resource servers for a user pool.
	//
	ListResourceServers(arg1 context.Context, arg2 *cognitoidentityprovider.ListResourceServersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error)
	// Lists the tags that are assigned to an Amazon Cognito user pool. A tag is a
	// label that you can apply to user pools to categorize and manage them in
	// different ways, such as by purpose, owner, environment, or other criteria. You
	// can use this action up to 10 times per second, per account.
	//
	ListTagsForResource(arg1 context.Context, arg2 *cognitoidentityprovider.ListTagsForResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error)
	// Lists the user import jobs.
	//
	ListUserImportJobs(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserImportJobsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	// Lists the clients that have been created for the specified user pool.
	//
	ListUserPoolClients(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserPoolClientsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	// Lists the user pools associated with an AWS account.
	//
	ListUserPools(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserPoolsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	// Lists the users in the Amazon Cognito user pool.
	//
	ListUsers(arg1 context.Context, arg2 *cognitoidentityprovider.ListUsersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error)
	// Lists the users in the specified group. Calling this action requires developer
	// credentials.
	//
	ListUsersInGroup(arg1 context.Context, arg2 *cognitoidentityprovider.ListUsersInGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	// Resends the confirmation (for confirmation of registration) to a specific user
	// in the user pool.
	//
	ResendConfirmationCode(arg1 context.Context, arg2 *cognitoidentityprovider.ResendConfirmationCodeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	// Responds to the authentication challenge.
	//
	RespondToAuthChallenge(arg1 context.Context, arg2 *cognitoidentityprovider.RespondToAuthChallengeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	// Configures actions on detected risks. To delete the risk configuration for
	// UserPoolId or ClientId, pass null values for all four configuration types. To
	// enable Amazon Cognito advanced security features, update the user pool to
	// include the UserPoolAddOns keyAdvancedSecurityMode.
	//
	SetRiskConfiguration(arg1 context.Context, arg2 *cognitoidentityprovider.SetRiskConfigurationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetRiskConfigurationOutput, error)
	// Sets the UI customization information for a user pool's built-in app UI. You can
	// specify app UI customization settings for a single client (with a specific
	// clientId) or for all clients (by setting the clientId to ALL). If you specify
	// ALL, the default configuration will be used for every client that has no UI
	// customization set previously. If you specify UI customization settings for a
	// particular client, it will no longer fall back to the ALL configuration. To use
	// this API, your user pool must have a domain associated with it. Otherwise, there
	// is no place to host the app's pages, and the service will throw an error.
	//
	SetUICustomization(arg1 context.Context, arg2 *cognitoidentityprovider.SetUICustomizationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUICustomizationOutput, error)
	// Set the user's multi-factor authentication (MFA) method preference, including
	// which MFA factors are enabled and if any are preferred. Only one factor can be
	// set as preferred. The preferred MFA factor will be used to authenticate a user
	// if multiple factors are enabled. If multiple options are enabled and no
	// preference is set, a challenge to choose an MFA option will be returned during
	// sign in. If an MFA type is enabled for a user, the user will be prompted for MFA
	// during all sign in attempts, unless device tracking is turned on and the device
	// has been trusted. If you would like MFA to be applied selectively based on the
	// assessed risk level of sign in attempts, disable MFA for users and turn on
	// Adaptive Authentication for the user pool.
	//
	SetUserMFAPreference(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserMFAPreferenceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error)
	// Set the user pool multi-factor authentication (MFA) configuration.
	//
	SetUserPoolMfaConfig(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserPoolMfaConfigInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error)
	// This action is no longer supported. You can use it to configure only SMS MFA.
	// You can't use it to configure TOTP software token MFA. To configure either type
	// of MFA, use SetUserMFAPreference
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUserMFAPreference.html)
	// instead.
	//
	SetUserSettings(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserSettingsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	// Registers the user in the specified user pool and creates a user name, password,
	// and user attributes.
	//
	SignUp(arg1 context.Context, arg2 *cognitoidentityprovider.SignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SignUpOutput, error)
	// Starts the user import.
	//
	StartUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.StartUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	// Stops the user import job.
	//
	StopUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.StopUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	// Assigns a set of tags to an Amazon Cognito user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria. Each tag consists of a key and
	// value, both of which you define. A key is a general category for more specific
	// values. For example, if you have two versions of a user pool, one for testing
	// and another for production, you might assign an Environment tag key to both user
	// pools. The value of this key might be Test for one user pool and Production for
	// the other. Tags are useful for cost tracking and access control. You can
	// activate your tags so that they appear on the Billing and Cost Management
	// console, where you can track the costs associated with your user pools. In an
	// IAM policy, you can constrain permissions for user pools based on specific tags
	// or tag values. You can use this action up to 5 times per second, per account. A
	// user pool can have as many as 50 tags.
	//
	TagResource(arg1 context.Context, arg2 *cognitoidentityprovider.TagResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.TagResourceOutput, error)
	// Removes the specified tags from an Amazon Cognito user pool. You can use this
	// action up to 5 times per second, per account
	//
	UntagResource(arg1 context.Context, arg2 *cognitoidentityprovider.UntagResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UntagResourceOutput, error)
	// Provides the feedback for an authentication event whether it was from a valid
	// user or not. This feedback is used for improving the risk evaluation decision
	// for the user pool as part of Amazon Cognito advanced security.
	//
	UpdateAuthEventFeedback(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateAuthEventFeedbackInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error)
	// Updates the device status.
	//
	UpdateDeviceStatus(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateDeviceStatusInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	// Updates the specified group with the specified attributes. Calling this action
	// requires developer credentials. If you don't provide a value for an attribute,
	// it will be set to the default value.
	//
	UpdateGroup(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateGroupOutput, error)
	// Updates identity provider information for a user pool.
	//
	UpdateIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error)
	// Updates the name and scopes of resource server. All other fields are read-only.
	// If you don't provide a value for an attribute, it will be set to the default
	// value.
	//
	UpdateResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateResourceServerOutput, error)
	// Allows a user to update a specific attribute (one at a time).
	//
	UpdateUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	// Updates the specified user pool with the specified attributes. You can get a
	// list of the current user pool settings using DescribeUserPool
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html).
	// If you don't provide a value for an attribute, it will be set to the default
	// value.
	//
	UpdateUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	// Updates the specified user pool app client with the specified attributes. You
	// can get a list of the current user pool app client settings using
	// DescribeUserPoolClient
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPoolClient.html).
	// If you don't provide a value for an attribute, it will be set to the default
	// value.
	//
	UpdateUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	// Updates the Secure Sockets Layer (SSL) certificate for the custom domain for
	// your user pool. You can use this operation to provide the Amazon Resource Name
	// (ARN) of a new certificate to Amazon Cognito. You cannot use it to change the
	// domain for a user pool. A custom domain is used to host the Amazon Cognito
	// hosted UI, which provides sign-up and sign-in pages for your application. When
	// you set up a custom domain, you provide a certificate that you manage with AWS
	// Certificate Manager (ACM). When necessary, you can use this operation to change
	// the certificate that you applied to your custom domain. Usually, this is
	// unnecessary following routine certificate renewal with ACM. When you renew your
	// existing certificate in ACM, the ARN for your certificate remains the same, and
	// your custom domain uses the new certificate automatically. However, if you
	// replace your existing certificate with a new one, ACM gives the new certificate
	// a new ARN. To apply the new certificate to your custom domain, you must provide
	// this ARN to Amazon Cognito. When you add your new certificate in ACM, you must
	// choose US East (N. Virginia) as the AWS Region. After you submit your request,
	// Amazon Cognito requires up to 1 hour to distribute your new certificate to your
	// custom domain. For more information about adding a custom domain to your user
	// pool, see Using Your Own Domain for the Hosted UI
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html).
	//
	UpdateUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error)
	// Use this API to register a user's entered TOTP code and mark the user's software
	// token MFA status as "verified" if successful. The request takes an access token
	// or a session string, but not both.
	//
	VerifySoftwareToken(arg1 context.Context, arg2 *cognitoidentityprovider.VerifySoftwareTokenInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error)
	// Verifies the specified user attributes in the user pool.
	//
	VerifyUserAttribute(arg1 context.Context, arg2 *cognitoidentityprovider.VerifyUserAttributeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
}

// CognitoIDPClientMock generic client mock
type CognitoIDPClientMock struct {
	AddCustomAttributesMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.AddCustomAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	AdminAddUserToGroupMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminAddUserToGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminConfirmSignUpMock               func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminConfirmSignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	AdminCreateUserMock                  func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminCreateUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminDeleteUserMock                  func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDeleteUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminDeleteUserAttributesMock        func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDeleteUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	AdminDisableProviderForUserMock      func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDisableProviderForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error)
	AdminDisableUserMock                 func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDisableUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	AdminEnableUserMock                  func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminEnableUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	AdminForgetDeviceMock                func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminForgetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	AdminGetDeviceMock                   func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminGetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	AdminGetUserMock                     func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminGetUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminInitiateAuthMock                func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminInitiateAuthInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	AdminLinkProviderForUserMock         func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminLinkProviderForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error)
	AdminListDevicesMock                 func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListDevicesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	AdminListGroupsForUserMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListGroupsForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	AdminListUserAuthEventsMock          func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListUserAuthEventsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error)
	AdminRemoveUserFromGroupMock         func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminRemoveUserFromGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminResetUserPasswordMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminResetUserPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	AdminRespondToAuthChallengeMock      func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminRespondToAuthChallengeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	AdminSetUserMFAPreferenceMock        func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserMFAPreferenceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error)
	AdminSetUserPasswordMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error)
	AdminSetUserSettingsMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserSettingsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	AdminUpdateAuthEventFeedbackMock     func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error)
	AdminUpdateDeviceStatusMock          func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateDeviceStatusInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	AdminUpdateUserAttributesMock        func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminUserGlobalSignOutMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUserGlobalSignOutInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	AssociateSoftwareTokenMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.AssociateSoftwareTokenInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error)
	ChangePasswordMock                   func(arg1 context.Context, arg2 *cognitoidentityprovider.ChangePasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ChangePasswordOutput, error)
	ConfirmDeviceMock                    func(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	ConfirmForgotPasswordMock            func(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmForgotPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ConfirmSignUpMock                    func(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmSignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	CreateGroupMock                      func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error)
	CreateIdentityProviderMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateIdentityProviderOutput, error)
	CreateResourceServerMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateResourceServerOutput, error)
	CreateUserImportJobMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	CreateUserPoolMock                   func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	CreateUserPoolClientMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	CreateUserPoolDomainMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error)
	DeleteGroupMock                      func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteGroupOutput, error)
	DeleteIdentityProviderMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error)
	DeleteResourceServerMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteResourceServerOutput, error)
	DeleteUserMock                       func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserOutput, error)
	DeleteUserAttributesMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	DeleteUserPoolMock                   func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	DeleteUserPoolClientMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	DeleteUserPoolDomainMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error)
	DescribeIdentityProviderMock         func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeResourceServerMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	DescribeRiskConfigurationMock        func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeRiskConfigurationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error)
	DescribeUserImportJobMock            func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserPoolMock                 func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolClientMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolDomainMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	ForgetDeviceMock                     func(arg1 context.Context, arg2 *cognitoidentityprovider.ForgetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	ForgotPasswordMock                   func(arg1 context.Context, arg2 *cognitoidentityprovider.ForgotPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	GetCSVHeaderMock                     func(arg1 context.Context, arg2 *cognitoidentityprovider.GetCSVHeaderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetDeviceMock                        func(arg1 context.Context, arg2 *cognitoidentityprovider.GetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetGroupMock                         func(arg1 context.Context, arg2 *cognitoidentityprovider.GetGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error)
	GetIdentityProviderByIdentifierMock  func(arg1 context.Context, arg2 *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	GetSigningCertificateMock            func(arg1 context.Context, arg2 *cognitoidentityprovider.GetSigningCertificateInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error)
	GetUICustomizationMock               func(arg1 context.Context, arg2 *cognitoidentityprovider.GetUICustomizationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	GetUserMock                          func(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserAttributeVerificationCodeMock func(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserPoolMfaConfigMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserPoolMfaConfigInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error)
	GlobalSignOutMock                    func(arg1 context.Context, arg2 *cognitoidentityprovider.GlobalSignOutInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	InitiateAuthMock                     func(arg1 context.Context, arg2 *cognitoidentityprovider.InitiateAuthInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.InitiateAuthOutput, error)
	ListDevicesMock                      func(arg1 context.Context, arg2 *cognitoidentityprovider.ListDevicesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListGroupsMock                       func(arg1 context.Context, arg2 *cognitoidentityprovider.ListGroupsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListIdentityProvidersMock            func(arg1 context.Context, arg2 *cognitoidentityprovider.ListIdentityProvidersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListResourceServersMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.ListResourceServersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error)
	ListTagsForResourceMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.ListTagsForResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error)
	ListUserImportJobsMock               func(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserImportJobsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserPoolClientsMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserPoolClientsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPoolsMock                    func(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserPoolsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUsersMock                        func(arg1 context.Context, arg2 *cognitoidentityprovider.ListUsersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersInGroupMock                 func(arg1 context.Context, arg2 *cognitoidentityprovider.ListUsersInGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	ResendConfirmationCodeMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.ResendConfirmationCodeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	RespondToAuthChallengeMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.RespondToAuthChallengeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	SetRiskConfigurationMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.SetRiskConfigurationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetRiskConfigurationOutput, error)
	SetUICustomizationMock               func(arg1 context.Context, arg2 *cognitoidentityprovider.SetUICustomizationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUICustomizationOutput, error)
	SetUserMFAPreferenceMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserMFAPreferenceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error)
	SetUserPoolMfaConfigMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserPoolMfaConfigInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error)
	SetUserSettingsMock                  func(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserSettingsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	SignUpMock                           func(arg1 context.Context, arg2 *cognitoidentityprovider.SignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SignUpOutput, error)
	StartUserImportJobMock               func(arg1 context.Context, arg2 *cognitoidentityprovider.StartUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	StopUserImportJobMock                func(arg1 context.Context, arg2 *cognitoidentityprovider.StopUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	TagResourceMock                      func(arg1 context.Context, arg2 *cognitoidentityprovider.TagResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.TagResourceOutput, error)
	UntagResourceMock                    func(arg1 context.Context, arg2 *cognitoidentityprovider.UntagResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UntagResourceOutput, error)
	UpdateAuthEventFeedbackMock          func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateAuthEventFeedbackInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error)
	UpdateDeviceStatusMock               func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateDeviceStatusInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	UpdateGroupMock                      func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateGroupOutput, error)
	UpdateIdentityProviderMock           func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error)
	UpdateResourceServerMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateResourceServerOutput, error)
	UpdateUserAttributesMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	UpdateUserPoolMock                   func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	UpdateUserPoolClientMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	UpdateUserPoolDomainMock             func(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error)
	VerifySoftwareTokenMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.VerifySoftwareTokenInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error)
	VerifyUserAttributeMock              func(arg1 context.Context, arg2 *cognitoidentityprovider.VerifyUserAttributeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
}

// AddCustomAttributes mocked funcition
func (m CognitoIDPClientMock) AddCustomAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.AddCustomAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
	return m.AddCustomAttributesMock(arg1, arg2, arg3...)
}

// AdminAddUserToGroup mocked funcition
func (m CognitoIDPClientMock) AdminAddUserToGroup(arg1 context.Context, arg2 *cognitoidentityprovider.AdminAddUserToGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
	return m.AdminAddUserToGroupMock(arg1, arg2, arg3...)
}

// AdminConfirmSignUp mocked funcition
func (m CognitoIDPClientMock) AdminConfirmSignUp(arg1 context.Context, arg2 *cognitoidentityprovider.AdminConfirmSignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error) {
	return m.AdminConfirmSignUpMock(arg1, arg2, arg3...)
}

// AdminCreateUser mocked funcition
func (m CognitoIDPClientMock) AdminCreateUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminCreateUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	return m.AdminCreateUserMock(arg1, arg2, arg3...)
}

// AdminDeleteUser mocked funcition
func (m CognitoIDPClientMock) AdminDeleteUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDeleteUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
	return m.AdminDeleteUserMock(arg1, arg2, arg3...)
}

// AdminDeleteUserAttributes mocked funcition
func (m CognitoIDPClientMock) AdminDeleteUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDeleteUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error) {
	return m.AdminDeleteUserAttributesMock(arg1, arg2, arg3...)
}

// AdminDisableProviderForUser mocked funcition
func (m CognitoIDPClientMock) AdminDisableProviderForUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDisableProviderForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error) {
	return m.AdminDisableProviderForUserMock(arg1, arg2, arg3...)
}

// AdminDisableUser mocked funcition
func (m CognitoIDPClientMock) AdminDisableUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminDisableUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDisableUserOutput, error) {
	return m.AdminDisableUserMock(arg1, arg2, arg3...)
}

// AdminEnableUser mocked funcition
func (m CognitoIDPClientMock) AdminEnableUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminEnableUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminEnableUserOutput, error) {
	return m.AdminEnableUserMock(arg1, arg2, arg3...)
}

// AdminForgetDevice mocked funcition
func (m CognitoIDPClientMock) AdminForgetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.AdminForgetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminForgetDeviceOutput, error) {
	return m.AdminForgetDeviceMock(arg1, arg2, arg3...)
}

// AdminGetDevice mocked funcition
func (m CognitoIDPClientMock) AdminGetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.AdminGetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetDeviceOutput, error) {
	return m.AdminGetDeviceMock(arg1, arg2, arg3...)
}

// AdminGetUser mocked funcition
func (m CognitoIDPClientMock) AdminGetUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminGetUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	return m.AdminGetUserMock(arg1, arg2, arg3...)
}

// AdminInitiateAuth mocked funcition
func (m CognitoIDPClientMock) AdminInitiateAuth(arg1 context.Context, arg2 *cognitoidentityprovider.AdminInitiateAuthInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	return m.AdminInitiateAuthMock(arg1, arg2, arg3...)
}

// AdminLinkProviderForUser mocked funcition
func (m CognitoIDPClientMock) AdminLinkProviderForUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminLinkProviderForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error) {
	return m.AdminLinkProviderForUserMock(arg1, arg2, arg3...)
}

// AdminListDevices mocked funcition
func (m CognitoIDPClientMock) AdminListDevices(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListDevicesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListDevicesOutput, error) {
	return m.AdminListDevicesMock(arg1, arg2, arg3...)
}

// AdminListGroupsForUser mocked funcition
func (m CognitoIDPClientMock) AdminListGroupsForUser(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListGroupsForUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
	return m.AdminListGroupsForUserMock(arg1, arg2, arg3...)
}

// AdminListUserAuthEvents mocked funcition
func (m CognitoIDPClientMock) AdminListUserAuthEvents(arg1 context.Context, arg2 *cognitoidentityprovider.AdminListUserAuthEventsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error) {
	return m.AdminListUserAuthEventsMock(arg1, arg2, arg3...)
}

// AdminRemoveUserFromGroup mocked funcition
func (m CognitoIDPClientMock) AdminRemoveUserFromGroup(arg1 context.Context, arg2 *cognitoidentityprovider.AdminRemoveUserFromGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
	return m.AdminRemoveUserFromGroupMock(arg1, arg2, arg3...)
}

// AdminResetUserPassword mocked funcition
func (m CognitoIDPClientMock) AdminResetUserPassword(arg1 context.Context, arg2 *cognitoidentityprovider.AdminResetUserPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error) {
	return m.AdminResetUserPasswordMock(arg1, arg2, arg3...)
}

// AdminRespondToAuthChallenge mocked funcition
func (m CognitoIDPClientMock) AdminRespondToAuthChallenge(arg1 context.Context, arg2 *cognitoidentityprovider.AdminRespondToAuthChallengeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
	return m.AdminRespondToAuthChallengeMock(arg1, arg2, arg3...)
}

// AdminSetUserMFAPreference mocked funcition
func (m CognitoIDPClientMock) AdminSetUserMFAPreference(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserMFAPreferenceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error) {
	return m.AdminSetUserMFAPreferenceMock(arg1, arg2, arg3...)
}

// AdminSetUserPassword mocked funcition
func (m CognitoIDPClientMock) AdminSetUserPassword(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
	return m.AdminSetUserPasswordMock(arg1, arg2, arg3...)
}

// AdminSetUserSettings mocked funcition
func (m CognitoIDPClientMock) AdminSetUserSettings(arg1 context.Context, arg2 *cognitoidentityprovider.AdminSetUserSettingsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error) {
	return m.AdminSetUserSettingsMock(arg1, arg2, arg3...)
}

// AdminUpdateAuthEventFeedback mocked funcition
func (m CognitoIDPClientMock) AdminUpdateAuthEventFeedback(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error) {
	return m.AdminUpdateAuthEventFeedbackMock(arg1, arg2, arg3...)
}

// AdminUpdateDeviceStatus mocked funcition
func (m CognitoIDPClientMock) AdminUpdateDeviceStatus(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateDeviceStatusInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error) {
	return m.AdminUpdateDeviceStatusMock(arg1, arg2, arg3...)
}

// AdminUpdateUserAttributes mocked funcition
func (m CognitoIDPClientMock) AdminUpdateUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUpdateUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
	return m.AdminUpdateUserAttributesMock(arg1, arg2, arg3...)
}

// AdminUserGlobalSignOut mocked funcition
func (m CognitoIDPClientMock) AdminUserGlobalSignOut(arg1 context.Context, arg2 *cognitoidentityprovider.AdminUserGlobalSignOutInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
	return m.AdminUserGlobalSignOutMock(arg1, arg2, arg3...)
}

// AssociateSoftwareToken mocked funcition
func (m CognitoIDPClientMock) AssociateSoftwareToken(arg1 context.Context, arg2 *cognitoidentityprovider.AssociateSoftwareTokenInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {
	return m.AssociateSoftwareTokenMock(arg1, arg2, arg3...)
}

// ChangePassword mocked funcition
func (m CognitoIDPClientMock) ChangePassword(arg1 context.Context, arg2 *cognitoidentityprovider.ChangePasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ChangePasswordOutput, error) {
	return m.ChangePasswordMock(arg1, arg2, arg3...)
}

// ConfirmDevice mocked funcition
func (m CognitoIDPClientMock) ConfirmDevice(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmDeviceOutput, error) {
	return m.ConfirmDeviceMock(arg1, arg2, arg3...)
}

// ConfirmForgotPassword mocked funcition
func (m CognitoIDPClientMock) ConfirmForgotPassword(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmForgotPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
	return m.ConfirmForgotPasswordMock(arg1, arg2, arg3...)
}

// ConfirmSignUp mocked funcition
func (m CognitoIDPClientMock) ConfirmSignUp(arg1 context.Context, arg2 *cognitoidentityprovider.ConfirmSignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
	return m.ConfirmSignUpMock(arg1, arg2, arg3...)
}

// CreateGroup mocked funcition
func (m CognitoIDPClientMock) CreateGroup(arg1 context.Context, arg2 *cognitoidentityprovider.CreateGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error) {
	return m.CreateGroupMock(arg1, arg2, arg3...)
}

// CreateIdentityProvider mocked funcition
func (m CognitoIDPClientMock) CreateIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.CreateIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateIdentityProviderOutput, error) {
	return m.CreateIdentityProviderMock(arg1, arg2, arg3...)
}

// CreateResourceServer mocked funcition
func (m CognitoIDPClientMock) CreateResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.CreateResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateResourceServerOutput, error) {
	return m.CreateResourceServerMock(arg1, arg2, arg3...)
}

// CreateUserImportJob mocked funcition
func (m CognitoIDPClientMock) CreateUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserImportJobOutput, error) {
	return m.CreateUserImportJobMock(arg1, arg2, arg3...)
}

// CreateUserPool mocked funcition
func (m CognitoIDPClientMock) CreateUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolOutput, error) {
	return m.CreateUserPoolMock(arg1, arg2, arg3...)
}

// CreateUserPoolClient mocked funcition
func (m CognitoIDPClientMock) CreateUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
	return m.CreateUserPoolClientMock(arg1, arg2, arg3...)
}

// CreateUserPoolDomain mocked funcition
func (m CognitoIDPClientMock) CreateUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.CreateUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error) {
	return m.CreateUserPoolDomainMock(arg1, arg2, arg3...)
}

// DeleteGroup mocked funcition
func (m CognitoIDPClientMock) DeleteGroup(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteGroupOutput, error) {
	return m.DeleteGroupMock(arg1, arg2, arg3...)
}

// DeleteIdentityProvider mocked funcition
func (m CognitoIDPClientMock) DeleteIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error) {
	return m.DeleteIdentityProviderMock(arg1, arg2, arg3...)
}

// DeleteResourceServer mocked funcition
func (m CognitoIDPClientMock) DeleteResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteResourceServerOutput, error) {
	return m.DeleteResourceServerMock(arg1, arg2, arg3...)
}

// DeleteUser mocked funcition
func (m CognitoIDPClientMock) DeleteUser(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserOutput, error) {
	return m.DeleteUserMock(arg1, arg2, arg3...)
}

// DeleteUserAttributes mocked funcition
func (m CognitoIDPClientMock) DeleteUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserAttributesOutput, error) {
	return m.DeleteUserAttributesMock(arg1, arg2, arg3...)
}

// DeleteUserPool mocked funcition
func (m CognitoIDPClientMock) DeleteUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolOutput, error) {
	return m.DeleteUserPoolMock(arg1, arg2, arg3...)
}

// DeleteUserPoolClient mocked funcition
func (m CognitoIDPClientMock) DeleteUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error) {
	return m.DeleteUserPoolClientMock(arg1, arg2, arg3...)
}

// DeleteUserPoolDomain mocked funcition
func (m CognitoIDPClientMock) DeleteUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.DeleteUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error) {
	return m.DeleteUserPoolDomainMock(arg1, arg2, arg3...)
}

// DescribeIdentityProvider mocked funcition
func (m CognitoIDPClientMock) DescribeIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
	return m.DescribeIdentityProviderMock(arg1, arg2, arg3...)
}

// DescribeResourceServer mocked funcition
func (m CognitoIDPClientMock) DescribeResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
	return m.DescribeResourceServerMock(arg1, arg2, arg3...)
}

// DescribeRiskConfiguration mocked funcition
func (m CognitoIDPClientMock) DescribeRiskConfiguration(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeRiskConfigurationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
	return m.DescribeRiskConfigurationMock(arg1, arg2, arg3...)
}

// DescribeUserImportJob mocked funcition
func (m CognitoIDPClientMock) DescribeUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
	return m.DescribeUserImportJobMock(arg1, arg2, arg3...)
}

// DescribeUserPool mocked funcition
func (m CognitoIDPClientMock) DescribeUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	return m.DescribeUserPoolMock(arg1, arg2, arg3...)
}

// DescribeUserPoolClient mocked funcition
func (m CognitoIDPClientMock) DescribeUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
	return m.DescribeUserPoolClientMock(arg1, arg2, arg3...)
}

// DescribeUserPoolDomain mocked funcition
func (m CognitoIDPClientMock) DescribeUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.DescribeUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
	return m.DescribeUserPoolDomainMock(arg1, arg2, arg3...)
}

// ForgetDevice mocked funcition
func (m CognitoIDPClientMock) ForgetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.ForgetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ForgetDeviceOutput, error) {
	return m.ForgetDeviceMock(arg1, arg2, arg3...)
}

// ForgotPassword mocked funcition
func (m CognitoIDPClientMock) ForgotPassword(arg1 context.Context, arg2 *cognitoidentityprovider.ForgotPasswordInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
	return m.ForgotPasswordMock(arg1, arg2, arg3...)
}

// GetCSVHeader mocked funcition
func (m CognitoIDPClientMock) GetCSVHeader(arg1 context.Context, arg2 *cognitoidentityprovider.GetCSVHeaderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
	return m.GetCSVHeaderMock(arg1, arg2, arg3...)
}

// GetDevice mocked funcition
func (m CognitoIDPClientMock) GetDevice(arg1 context.Context, arg2 *cognitoidentityprovider.GetDeviceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error) {
	return m.GetDeviceMock(arg1, arg2, arg3...)
}

// GetGroup mocked funcition
func (m CognitoIDPClientMock) GetGroup(arg1 context.Context, arg2 *cognitoidentityprovider.GetGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error) {
	return m.GetGroupMock(arg1, arg2, arg3...)
}

// GetIdentityProviderByIdentifier mocked funcition
func (m CognitoIDPClientMock) GetIdentityProviderByIdentifier(arg1 context.Context, arg2 *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
	return m.GetIdentityProviderByIdentifierMock(arg1, arg2, arg3...)
}

// GetSigningCertificate mocked funcition
func (m CognitoIDPClientMock) GetSigningCertificate(arg1 context.Context, arg2 *cognitoidentityprovider.GetSigningCertificateInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
	return m.GetSigningCertificateMock(arg1, arg2, arg3...)
}

// GetUICustomization mocked funcition
func (m CognitoIDPClientMock) GetUICustomization(arg1 context.Context, arg2 *cognitoidentityprovider.GetUICustomizationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
	return m.GetUICustomizationMock(arg1, arg2, arg3...)
}

// GetUser mocked funcition
func (m CognitoIDPClientMock) GetUser(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error) {
	return m.GetUserMock(arg1, arg2, arg3...)
}

// GetUserAttributeVerificationCode mocked funcition
func (m CognitoIDPClientMock) GetUserAttributeVerificationCode(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
	return m.GetUserAttributeVerificationCodeMock(arg1, arg2, arg3...)
}

// GetUserPoolMfaConfig mocked funcition
func (m CognitoIDPClientMock) GetUserPoolMfaConfig(arg1 context.Context, arg2 *cognitoidentityprovider.GetUserPoolMfaConfigInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
	return m.GetUserPoolMfaConfigMock(arg1, arg2, arg3...)
}

// GlobalSignOut mocked funcition
func (m CognitoIDPClientMock) GlobalSignOut(arg1 context.Context, arg2 *cognitoidentityprovider.GlobalSignOutInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
	return m.GlobalSignOutMock(arg1, arg2, arg3...)
}

// InitiateAuth mocked funcition
func (m CognitoIDPClientMock) InitiateAuth(arg1 context.Context, arg2 *cognitoidentityprovider.InitiateAuthInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	return m.InitiateAuthMock(arg1, arg2, arg3...)
}

// ListDevices mocked funcition
func (m CognitoIDPClientMock) ListDevices(arg1 context.Context, arg2 *cognitoidentityprovider.ListDevicesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error) {
	return m.ListDevicesMock(arg1, arg2, arg3...)
}

// ListGroups mocked funcition
func (m CognitoIDPClientMock) ListGroups(arg1 context.Context, arg2 *cognitoidentityprovider.ListGroupsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error) {
	return m.ListGroupsMock(arg1, arg2, arg3...)
}

// ListIdentityProviders mocked funcition
func (m CognitoIDPClientMock) ListIdentityProviders(arg1 context.Context, arg2 *cognitoidentityprovider.ListIdentityProvidersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
	return m.ListIdentityProvidersMock(arg1, arg2, arg3...)
}

// ListResourceServers mocked funcition
func (m CognitoIDPClientMock) ListResourceServers(arg1 context.Context, arg2 *cognitoidentityprovider.ListResourceServersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error) {
	return m.ListResourceServersMock(arg1, arg2, arg3...)
}

// ListTagsForResource mocked funcition
func (m CognitoIDPClientMock) ListTagsForResource(arg1 context.Context, arg2 *cognitoidentityprovider.ListTagsForResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
	return m.ListTagsForResourceMock(arg1, arg2, arg3...)
}

// ListUserImportJobs mocked funcition
func (m CognitoIDPClientMock) ListUserImportJobs(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserImportJobsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
	return m.ListUserImportJobsMock(arg1, arg2, arg3...)
}

// ListUserPoolClients mocked funcition
func (m CognitoIDPClientMock) ListUserPoolClients(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserPoolClientsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
	return m.ListUserPoolClientsMock(arg1, arg2, arg3...)
}

// ListUserPools mocked funcition
func (m CognitoIDPClientMock) ListUserPools(arg1 context.Context, arg2 *cognitoidentityprovider.ListUserPoolsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
	return m.ListUserPoolsMock(arg1, arg2, arg3...)
}

// ListUsers mocked funcition
func (m CognitoIDPClientMock) ListUsers(arg1 context.Context, arg2 *cognitoidentityprovider.ListUsersInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error) {
	return m.ListUsersMock(arg1, arg2, arg3...)
}

// ListUsersInGroup mocked funcition
func (m CognitoIDPClientMock) ListUsersInGroup(arg1 context.Context, arg2 *cognitoidentityprovider.ListUsersInGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	return m.ListUsersInGroupMock(arg1, arg2, arg3...)
}

// ResendConfirmationCode mocked funcition
func (m CognitoIDPClientMock) ResendConfirmationCode(arg1 context.Context, arg2 *cognitoidentityprovider.ResendConfirmationCodeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
	return m.ResendConfirmationCodeMock(arg1, arg2, arg3...)
}

// RespondToAuthChallenge mocked funcition
func (m CognitoIDPClientMock) RespondToAuthChallenge(arg1 context.Context, arg2 *cognitoidentityprovider.RespondToAuthChallengeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
	return m.RespondToAuthChallengeMock(arg1, arg2, arg3...)
}

// SetRiskConfiguration mocked funcition
func (m CognitoIDPClientMock) SetRiskConfiguration(arg1 context.Context, arg2 *cognitoidentityprovider.SetRiskConfigurationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetRiskConfigurationOutput, error) {
	return m.SetRiskConfigurationMock(arg1, arg2, arg3...)
}

// SetUICustomization mocked funcition
func (m CognitoIDPClientMock) SetUICustomization(arg1 context.Context, arg2 *cognitoidentityprovider.SetUICustomizationInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUICustomizationOutput, error) {
	return m.SetUICustomizationMock(arg1, arg2, arg3...)
}

// SetUserMFAPreference mocked funcition
func (m CognitoIDPClientMock) SetUserMFAPreference(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserMFAPreferenceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {
	return m.SetUserMFAPreferenceMock(arg1, arg2, arg3...)
}

// SetUserPoolMfaConfig mocked funcition
func (m CognitoIDPClientMock) SetUserPoolMfaConfig(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserPoolMfaConfigInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error) {
	return m.SetUserPoolMfaConfigMock(arg1, arg2, arg3...)
}

// SetUserSettings mocked funcition
func (m CognitoIDPClientMock) SetUserSettings(arg1 context.Context, arg2 *cognitoidentityprovider.SetUserSettingsInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SetUserSettingsOutput, error) {
	return m.SetUserSettingsMock(arg1, arg2, arg3...)
}

// SignUp mocked funcition
func (m CognitoIDPClientMock) SignUp(arg1 context.Context, arg2 *cognitoidentityprovider.SignUpInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SignUpOutput, error) {
	return m.SignUpMock(arg1, arg2, arg3...)
}

// StartUserImportJob mocked funcition
func (m CognitoIDPClientMock) StartUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.StartUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.StartUserImportJobOutput, error) {
	return m.StartUserImportJobMock(arg1, arg2, arg3...)
}

// StopUserImportJob mocked funcition
func (m CognitoIDPClientMock) StopUserImportJob(arg1 context.Context, arg2 *cognitoidentityprovider.StopUserImportJobInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.StopUserImportJobOutput, error) {
	return m.StopUserImportJobMock(arg1, arg2, arg3...)
}

// TagResource mocked funcition
func (m CognitoIDPClientMock) TagResource(arg1 context.Context, arg2 *cognitoidentityprovider.TagResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.TagResourceOutput, error) {
	return m.TagResourceMock(arg1, arg2, arg3...)
}

// UntagResource mocked funcition
func (m CognitoIDPClientMock) UntagResource(arg1 context.Context, arg2 *cognitoidentityprovider.UntagResourceInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UntagResourceOutput, error) {
	return m.UntagResourceMock(arg1, arg2, arg3...)
}

// UpdateAuthEventFeedback mocked funcition
func (m CognitoIDPClientMock) UpdateAuthEventFeedback(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateAuthEventFeedbackInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error) {
	return m.UpdateAuthEventFeedbackMock(arg1, arg2, arg3...)
}

// UpdateDeviceStatus mocked funcition
func (m CognitoIDPClientMock) UpdateDeviceStatus(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateDeviceStatusInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error) {
	return m.UpdateDeviceStatusMock(arg1, arg2, arg3...)
}

// UpdateGroup mocked funcition
func (m CognitoIDPClientMock) UpdateGroup(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateGroupInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateGroupOutput, error) {
	return m.UpdateGroupMock(arg1, arg2, arg3...)
}

// UpdateIdentityProvider mocked funcition
func (m CognitoIDPClientMock) UpdateIdentityProvider(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateIdentityProviderInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error) {
	return m.UpdateIdentityProviderMock(arg1, arg2, arg3...)
}

// UpdateResourceServer mocked funcition
func (m CognitoIDPClientMock) UpdateResourceServer(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateResourceServerInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateResourceServerOutput, error) {
	return m.UpdateResourceServerMock(arg1, arg2, arg3...)
}

// UpdateUserAttributes mocked funcition
func (m CognitoIDPClientMock) UpdateUserAttributes(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserAttributesInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
	return m.UpdateUserAttributesMock(arg1, arg2, arg3...)
}

// UpdateUserPool mocked funcition
func (m CognitoIDPClientMock) UpdateUserPool(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolOutput, error) {
	return m.UpdateUserPoolMock(arg1, arg2, arg3...)
}

// UpdateUserPoolClient mocked funcition
func (m CognitoIDPClientMock) UpdateUserPoolClient(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolClientInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error) {
	return m.UpdateUserPoolClientMock(arg1, arg2, arg3...)
}

// UpdateUserPoolDomain mocked funcition
func (m CognitoIDPClientMock) UpdateUserPoolDomain(arg1 context.Context, arg2 *cognitoidentityprovider.UpdateUserPoolDomainInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error) {
	return m.UpdateUserPoolDomainMock(arg1, arg2, arg3...)
}

// VerifySoftwareToken mocked funcition
func (m CognitoIDPClientMock) VerifySoftwareToken(arg1 context.Context, arg2 *cognitoidentityprovider.VerifySoftwareTokenInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {
	return m.VerifySoftwareTokenMock(arg1, arg2, arg3...)
}

// VerifyUserAttribute mocked funcition
func (m CognitoIDPClientMock) VerifyUserAttribute(arg1 context.Context, arg2 *cognitoidentityprovider.VerifyUserAttributeInput, arg3 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.VerifyUserAttributeOutput, error) {
	return m.VerifyUserAttributeMock(arg1, arg2, arg3...)
}
