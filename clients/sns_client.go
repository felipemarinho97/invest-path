package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// ISNSClient generic client
type ISNSClient interface {
	// Adds a statement to a topic's access control policy, granting access for the
	// specified AWS accounts to the specified actions.
	//
	AddPermission(arg1 context.Context, arg2 *sns.AddPermissionInput, arg3 ...func(*sns.Options)) (*sns.AddPermissionOutput, error)
	// Accepts a phone number and indicates whether the phone holder has opted out of
	// receiving SMS messages from your account. You cannot send SMS messages to a
	// number that is opted out. To resume sending messages, you can opt in the number
	// by using the OptInPhoneNumber action.
	//
	CheckIfPhoneNumberIsOptedOut(arg1 context.Context, arg2 *sns.CheckIfPhoneNumberIsOptedOutInput, arg3 ...func(*sns.Options)) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	// Verifies an endpoint owner's intent to receive messages by validating the token
	// sent to the endpoint by an earlier Subscribe action. If the token is valid, the
	// action creates a new subscription and returns its Amazon Resource Name (ARN).
	// This call requires an AWS signature only when the AuthenticateOnUnsubscribe flag
	// is set to "true".
	//
	ConfirmSubscription(arg1 context.Context, arg2 *sns.ConfirmSubscriptionInput, arg3 ...func(*sns.Options)) (*sns.ConfirmSubscriptionOutput, error)
	// Creates a platform application object for one of the supported push notification
	// services, such as APNS and GCM (Firebase Cloud Messaging), to which devices and
	// mobile apps may register. You must specify PlatformPrincipal and
	// PlatformCredential attributes when using the CreatePlatformApplication action.
	// PlatformPrincipal and PlatformCredential are received from the notification
	// service.
	//
	// * For ADM, PlatformPrincipal is client id and PlatformCredential is
	// client secret.
	//
	// * For Baidu, PlatformPrincipal is API key and PlatformCredential
	// is secret key.
	//
	// * For APNS and APNS_SANDBOX, PlatformPrincipal is SSL
	// certificate and PlatformCredential is private key.
	//
	// * For GCM (Firebase Cloud
	// Messaging), there is no PlatformPrincipal and the PlatformCredential is API
	// key.
	//
	// * For MPNS, PlatformPrincipal is TLS certificate and PlatformCredential is
	// private key.
	//
	// * For WNS, PlatformPrincipal is Package Security Identifier and
	// PlatformCredential is secret key.
	//
	// You can use the returned
	// PlatformApplicationArn as an attribute for the CreatePlatformEndpoint action.
	//
	CreatePlatformApplication(arg1 context.Context, arg2 *sns.CreatePlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.CreatePlatformApplicationOutput, error)
	// Creates an endpoint for a device and mobile app on one of the supported push
	// notification services, such as GCM (Firebase Cloud Messaging) and APNS.
	// CreatePlatformEndpoint requires the PlatformApplicationArn that is returned from
	// CreatePlatformApplication. You can use the returned EndpointArn to send a
	// message to a mobile app or by the Subscribe action for subscription to a topic.
	// The CreatePlatformEndpoint action is idempotent, so if the requester already
	// owns an endpoint with the same device token and attributes, that endpoint's ARN
	// is returned without creating a new endpoint. For more information, see Using
	// Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). When using
	// CreatePlatformEndpoint with Baidu, two attributes must be provided: ChannelId
	// and UserId. The token field must also contain the ChannelId. For more
	// information, see Creating an Amazon SNS Endpoint for Baidu
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html).
	//
	CreatePlatformEndpoint(arg1 context.Context, arg2 *sns.CreatePlatformEndpointInput, arg3 ...func(*sns.Options)) (*sns.CreatePlatformEndpointOutput, error)
	// Creates a topic to which notifications can be published. Users can create at
	// most 100,000 standard topics (at most 1,000 FIFO topics). For more information,
	// see https://aws.amazon.com/sns (http://aws.amazon.com/sns/). This action is
	// idempotent, so if the requester already owns a topic with the specified name,
	// that topic's ARN is returned without creating a new topic.
	//
	CreateTopic(arg1 context.Context, arg2 *sns.CreateTopicInput, arg3 ...func(*sns.Options)) (*sns.CreateTopicOutput, error)
	// Deletes the endpoint for a device and mobile app from Amazon SNS. This action is
	// idempotent. For more information, see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). When you delete
	// an endpoint that is also subscribed to a topic, then you must also unsubscribe
	// the endpoint from the topic.
	//
	DeleteEndpoint(arg1 context.Context, arg2 *sns.DeleteEndpointInput, arg3 ...func(*sns.Options)) (*sns.DeleteEndpointOutput, error)
	// Deletes a platform application object for one of the supported push notification
	// services, such as APNS and GCM (Firebase Cloud Messaging). For more information,
	// see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
	//
	DeletePlatformApplication(arg1 context.Context, arg2 *sns.DeletePlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.DeletePlatformApplicationOutput, error)
	// Deletes a topic and all its subscriptions. Deleting a topic might prevent some
	// messages previously sent to the topic from being delivered to subscribers. This
	// action is idempotent, so deleting a topic that does not exist does not result in
	// an error.
	//
	DeleteTopic(arg1 context.Context, arg2 *sns.DeleteTopicInput, arg3 ...func(*sns.Options)) (*sns.DeleteTopicOutput, error)
	// Retrieves the endpoint attributes for a device on one of the supported push
	// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
	// information, see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
	//
	GetEndpointAttributes(arg1 context.Context, arg2 *sns.GetEndpointAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error)
	// Retrieves the attributes of the platform application object for the supported
	// push notification services, such as APNS and GCM (Firebase Cloud Messaging). For
	// more information, see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
	//
	GetPlatformApplicationAttributes(arg1 context.Context, arg2 *sns.GetPlatformApplicationAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error)
	// Returns the settings for sending SMS messages from your account. These settings
	// are set with the SetSMSAttributes action.
	//
	GetSMSAttributes(arg1 context.Context, arg2 *sns.GetSMSAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error)
	// Returns all of the properties of a subscription.
	//
	GetSubscriptionAttributes(arg1 context.Context, arg2 *sns.GetSubscriptionAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error)
	// Returns all of the properties of a topic. Topic properties returned might differ
	// based on the authorization of the user.
	//
	GetTopicAttributes(arg1 context.Context, arg2 *sns.GetTopicAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	// Lists the endpoints and endpoint attributes for devices in a supported push
	// notification service, such as GCM (Firebase Cloud Messaging) and APNS. The
	// results for ListEndpointsByPlatformApplication are paginated and return a
	// limited list of endpoints, up to 100. If additional records are available after
	// the first page results, then a NextToken string will be returned. To receive the
	// next page, you call ListEndpointsByPlatformApplication again using the NextToken
	// string received from the previous call. When there are no more records to
	// return, NextToken will be null. For more information, see Using Amazon SNS
	// Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). This action is
	// throttled at 30 transactions per second (TPS).
	//
	ListEndpointsByPlatformApplication(arg1 context.Context, arg2 *sns.ListEndpointsByPlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	// Returns a list of phone numbers that are opted out, meaning you cannot send SMS
	// messages to them. The results for ListPhoneNumbersOptedOut are paginated, and
	// each page returns up to 100 phone numbers. If additional phone numbers are
	// available after the first page of results, then a NextToken string will be
	// returned. To receive the next page, you call ListPhoneNumbersOptedOut again
	// using the NextToken string received from the previous call. When there are no
	// more records to return, NextToken will be null.
	//
	ListPhoneNumbersOptedOut(arg1 context.Context, arg2 *sns.ListPhoneNumbersOptedOutInput, arg3 ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error)
	// Lists the platform application objects for the supported push notification
	// services, such as APNS and GCM (Firebase Cloud Messaging). The results for
	// ListPlatformApplications are paginated and return a limited list of
	// applications, up to 100. If additional records are available after the first
	// page results, then a NextToken string will be returned. To receive the next
	// page, you call ListPlatformApplications using the NextToken string received from
	// the previous call. When there are no more records to return, NextToken will be
	// null. For more information, see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). This action is
	// throttled at 15 transactions per second (TPS).
	//
	ListPlatformApplications(arg1 context.Context, arg2 *sns.ListPlatformApplicationsInput, arg3 ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error)
	// Returns a list of the requester's subscriptions. Each call returns a limited
	// list of subscriptions, up to 100. If there are more subscriptions, a NextToken
	// is also returned. Use the NextToken parameter in a new ListSubscriptions call to
	// get further results. This action is throttled at 30 transactions per second
	// (TPS).
	//
	ListSubscriptions(arg1 context.Context, arg2 *sns.ListSubscriptionsInput, arg3 ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	// Returns a list of the subscriptions to a specific topic. Each call returns a
	// limited list of subscriptions, up to 100. If there are more subscriptions, a
	// NextToken is also returned. Use the NextToken parameter in a new
	// ListSubscriptionsByTopic call to get further results. This action is throttled
	// at 30 transactions per second (TPS).
	//
	ListSubscriptionsByTopic(arg1 context.Context, arg2 *sns.ListSubscriptionsByTopicInput, arg3 ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error)
	// List all tags added to the specified Amazon SNS topic. For an overview, see
	// Amazon SNS Tags (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html) in the
	// Amazon Simple Notification Service Developer Guide.
	//
	ListTagsForResource(arg1 context.Context, arg2 *sns.ListTagsForResourceInput, arg3 ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
	// Returns a list of the requester's topics. Each call returns a limited list of
	// topics, up to 100. If there are more topics, a NextToken is also returned. Use
	// the NextToken parameter in a new ListTopics call to get further results. This
	// action is throttled at 30 transactions per second (TPS).
	//
	ListTopics(arg1 context.Context, arg2 *sns.ListTopicsInput, arg3 ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
	// Use this request to opt in a phone number that is opted out, which enables you
	// to resume sending SMS messages to the number. You can opt in a phone number only
	// once every 30 days.
	//
	OptInPhoneNumber(arg1 context.Context, arg2 *sns.OptInPhoneNumberInput, arg3 ...func(*sns.Options)) (*sns.OptInPhoneNumberOutput, error)
	// Sends a message to an Amazon SNS topic, a text message (SMS message) directly to
	// a phone number, or a message to a mobile platform endpoint (when you specify the
	// TargetArn). If you send a message to a topic, Amazon SNS delivers the message to
	// each endpoint that is subscribed to the topic. The format of the message depends
	// on the notification protocol for each subscribed endpoint. When a messageId is
	// returned, the message has been saved and Amazon SNS will attempt to deliver it
	// shortly. To use the Publish action for sending a message to a mobile endpoint,
	// such as an app on a Kindle device or mobile phone, you must specify the
	// EndpointArn for the TargetArn parameter. The EndpointArn is returned when making
	// a call with the CreatePlatformEndpoint action. For more information about
	// formatting messages, see Send Custom Platform-Specific Payloads in Messages to
	// Mobile Devices
	// (https://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-custommessage.html).
	// You can publish messages only to topics and endpoints in the same AWS Region.
	//
	Publish(arg1 context.Context, arg2 *sns.PublishInput, arg3 ...func(*sns.Options)) (*sns.PublishOutput, error)
	// Removes a statement from a topic's access control policy.
	//
	RemovePermission(arg1 context.Context, arg2 *sns.RemovePermissionInput, arg3 ...func(*sns.Options)) (*sns.RemovePermissionOutput, error)
	// Sets the attributes for an endpoint for a device on one of the supported push
	// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
	// information, see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
	//
	SetEndpointAttributes(arg1 context.Context, arg2 *sns.SetEndpointAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetEndpointAttributesOutput, error)
	// Sets the attributes of the platform application object for the supported push
	// notification services, such as APNS and GCM (Firebase Cloud Messaging). For more
	// information, see Using Amazon SNS Mobile Push Notifications
	// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). For information
	// on configuring attributes for message delivery status, see Using Amazon SNS
	// Application Attributes for Message Delivery Status
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html).
	//
	SetPlatformApplicationAttributes(arg1 context.Context, arg2 *sns.SetPlatformApplicationAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetPlatformApplicationAttributesOutput, error)
	// Use this request to set the default settings for sending SMS messages and
	// receiving daily SMS usage reports. You can override some of these settings for a
	// single message when you use the Publish action with the
	// MessageAttributes.entry.N parameter. For more information, see Publishing to a
	// mobile phone
	// (https://docs.aws.amazon.com/sns/latest/dg/sms_publish-to-phone.html) in the
	// Amazon SNS Developer Guide. To use this operation, you must grant the Amazon SNS
	// service principal (sns.amazonaws.com) permission to perform the s3:ListBucket
	// action.
	//
	SetSMSAttributes(arg1 context.Context, arg2 *sns.SetSMSAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetSMSAttributesOutput, error)
	// Allows a subscription owner to set an attribute of the subscription to a new
	// value.
	//
	SetSubscriptionAttributes(arg1 context.Context, arg2 *sns.SetSubscriptionAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetSubscriptionAttributesOutput, error)
	// Allows a topic owner to set an attribute of the topic to a new value.
	//
	SetTopicAttributes(arg1 context.Context, arg2 *sns.SetTopicAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetTopicAttributesOutput, error)
	// Subscribes an endpoint to an Amazon SNS topic. If the endpoint type is HTTP/S or
	// email, or if the endpoint and the topic are not in the same AWS account, the
	// endpoint owner must run the ConfirmSubscription action to confirm the
	// subscription. You call the ConfirmSubscription action with the token from the
	// subscription response. Confirmation tokens are valid for three days. This action
	// is throttled at 100 transactions per second (TPS).
	//
	Subscribe(arg1 context.Context, arg2 *sns.SubscribeInput, arg3 ...func(*sns.Options)) (*sns.SubscribeOutput, error)
	// Add tags to the specified Amazon SNS topic. For an overview, see Amazon SNS Tags
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html) in the Amazon SNS
	// Developer Guide. When you use topic tags, keep the following guidelines in
	// mind:
	//
	// * Adding more than 50 tags to a topic isn't recommended.
	//
	// * Tags don't
	// have any semantic meaning. Amazon SNS interprets tags as character strings.
	//
	// *
	// Tags are case-sensitive.
	//
	// * A new tag with a key identical to that of an
	// existing tag overwrites the existing tag.
	//
	// * Tagging actions are limited to 10
	// TPS per AWS account, per AWS region. If your application requires a higher
	// throughput, file a technical support request
	// (https://console.aws.amazon.com/support/home#/case/create?issueType=technical).
	//
	TagResource(arg1 context.Context, arg2 *sns.TagResourceInput, arg3 ...func(*sns.Options)) (*sns.TagResourceOutput, error)
	// Deletes a subscription. If the subscription requires authentication for
	// deletion, only the owner of the subscription or the topic's owner can
	// unsubscribe, and an AWS signature is required. If the Unsubscribe call does not
	// require authentication and the requester is not the subscription owner, a final
	// cancellation message is delivered to the endpoint, so that the endpoint owner
	// can easily resubscribe to the topic if the Unsubscribe request was unintended.
	// This action is throttled at 100 transactions per second (TPS).
	//
	Unsubscribe(arg1 context.Context, arg2 *sns.UnsubscribeInput, arg3 ...func(*sns.Options)) (*sns.UnsubscribeOutput, error)
	// Remove tags from the specified Amazon SNS topic. For an overview, see Amazon SNS
	// Tags (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html) in the Amazon SNS
	// Developer Guide.
	//
	UntagResource(arg1 context.Context, arg2 *sns.UntagResourceInput, arg3 ...func(*sns.Options)) (*sns.UntagResourceOutput, error)
}

// SNSClientMock generic client mock
type SNSClientMock struct {
	AddPermissionMock                      func(arg1 context.Context, arg2 *sns.AddPermissionInput, arg3 ...func(*sns.Options)) (*sns.AddPermissionOutput, error)
	CheckIfPhoneNumberIsOptedOutMock       func(arg1 context.Context, arg2 *sns.CheckIfPhoneNumberIsOptedOutInput, arg3 ...func(*sns.Options)) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	ConfirmSubscriptionMock                func(arg1 context.Context, arg2 *sns.ConfirmSubscriptionInput, arg3 ...func(*sns.Options)) (*sns.ConfirmSubscriptionOutput, error)
	CreatePlatformApplicationMock          func(arg1 context.Context, arg2 *sns.CreatePlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.CreatePlatformApplicationOutput, error)
	CreatePlatformEndpointMock             func(arg1 context.Context, arg2 *sns.CreatePlatformEndpointInput, arg3 ...func(*sns.Options)) (*sns.CreatePlatformEndpointOutput, error)
	CreateTopicMock                        func(arg1 context.Context, arg2 *sns.CreateTopicInput, arg3 ...func(*sns.Options)) (*sns.CreateTopicOutput, error)
	DeleteEndpointMock                     func(arg1 context.Context, arg2 *sns.DeleteEndpointInput, arg3 ...func(*sns.Options)) (*sns.DeleteEndpointOutput, error)
	DeletePlatformApplicationMock          func(arg1 context.Context, arg2 *sns.DeletePlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.DeletePlatformApplicationOutput, error)
	DeleteTopicMock                        func(arg1 context.Context, arg2 *sns.DeleteTopicInput, arg3 ...func(*sns.Options)) (*sns.DeleteTopicOutput, error)
	GetEndpointAttributesMock              func(arg1 context.Context, arg2 *sns.GetEndpointAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error)
	GetPlatformApplicationAttributesMock   func(arg1 context.Context, arg2 *sns.GetPlatformApplicationAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetSMSAttributesMock                   func(arg1 context.Context, arg2 *sns.GetSMSAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error)
	GetSubscriptionAttributesMock          func(arg1 context.Context, arg2 *sns.GetSubscriptionAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error)
	GetTopicAttributesMock                 func(arg1 context.Context, arg2 *sns.GetTopicAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	ListEndpointsByPlatformApplicationMock func(arg1 context.Context, arg2 *sns.ListEndpointsByPlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListPhoneNumbersOptedOutMock           func(arg1 context.Context, arg2 *sns.ListPhoneNumbersOptedOutInput, arg3 ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPlatformApplicationsMock           func(arg1 context.Context, arg2 *sns.ListPlatformApplicationsInput, arg3 ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error)
	ListSubscriptionsMock                  func(arg1 context.Context, arg2 *sns.ListSubscriptionsInput, arg3 ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsByTopicMock           func(arg1 context.Context, arg2 *sns.ListSubscriptionsByTopicInput, arg3 ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error)
	ListTagsForResourceMock                func(arg1 context.Context, arg2 *sns.ListTagsForResourceInput, arg3 ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
	ListTopicsMock                         func(arg1 context.Context, arg2 *sns.ListTopicsInput, arg3 ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
	OptInPhoneNumberMock                   func(arg1 context.Context, arg2 *sns.OptInPhoneNumberInput, arg3 ...func(*sns.Options)) (*sns.OptInPhoneNumberOutput, error)
	PublishMock                            func(arg1 context.Context, arg2 *sns.PublishInput, arg3 ...func(*sns.Options)) (*sns.PublishOutput, error)
	RemovePermissionMock                   func(arg1 context.Context, arg2 *sns.RemovePermissionInput, arg3 ...func(*sns.Options)) (*sns.RemovePermissionOutput, error)
	SetEndpointAttributesMock              func(arg1 context.Context, arg2 *sns.SetEndpointAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetEndpointAttributesOutput, error)
	SetPlatformApplicationAttributesMock   func(arg1 context.Context, arg2 *sns.SetPlatformApplicationAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetPlatformApplicationAttributesOutput, error)
	SetSMSAttributesMock                   func(arg1 context.Context, arg2 *sns.SetSMSAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetSMSAttributesOutput, error)
	SetSubscriptionAttributesMock          func(arg1 context.Context, arg2 *sns.SetSubscriptionAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetSubscriptionAttributesOutput, error)
	SetTopicAttributesMock                 func(arg1 context.Context, arg2 *sns.SetTopicAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetTopicAttributesOutput, error)
	SubscribeMock                          func(arg1 context.Context, arg2 *sns.SubscribeInput, arg3 ...func(*sns.Options)) (*sns.SubscribeOutput, error)
	TagResourceMock                        func(arg1 context.Context, arg2 *sns.TagResourceInput, arg3 ...func(*sns.Options)) (*sns.TagResourceOutput, error)
	UnsubscribeMock                        func(arg1 context.Context, arg2 *sns.UnsubscribeInput, arg3 ...func(*sns.Options)) (*sns.UnsubscribeOutput, error)
	UntagResourceMock                      func(arg1 context.Context, arg2 *sns.UntagResourceInput, arg3 ...func(*sns.Options)) (*sns.UntagResourceOutput, error)
}

// AddPermission mocked funcition
func (m SNSClientMock) AddPermission(arg1 context.Context, arg2 *sns.AddPermissionInput, arg3 ...func(*sns.Options)) (*sns.AddPermissionOutput, error) {
	return m.AddPermissionMock(arg1, arg2, arg3...)
}

// CheckIfPhoneNumberIsOptedOut mocked funcition
func (m SNSClientMock) CheckIfPhoneNumberIsOptedOut(arg1 context.Context, arg2 *sns.CheckIfPhoneNumberIsOptedOutInput, arg3 ...func(*sns.Options)) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
	return m.CheckIfPhoneNumberIsOptedOutMock(arg1, arg2, arg3...)
}

// ConfirmSubscription mocked funcition
func (m SNSClientMock) ConfirmSubscription(arg1 context.Context, arg2 *sns.ConfirmSubscriptionInput, arg3 ...func(*sns.Options)) (*sns.ConfirmSubscriptionOutput, error) {
	return m.ConfirmSubscriptionMock(arg1, arg2, arg3...)
}

// CreatePlatformApplication mocked funcition
func (m SNSClientMock) CreatePlatformApplication(arg1 context.Context, arg2 *sns.CreatePlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.CreatePlatformApplicationOutput, error) {
	return m.CreatePlatformApplicationMock(arg1, arg2, arg3...)
}

// CreatePlatformEndpoint mocked funcition
func (m SNSClientMock) CreatePlatformEndpoint(arg1 context.Context, arg2 *sns.CreatePlatformEndpointInput, arg3 ...func(*sns.Options)) (*sns.CreatePlatformEndpointOutput, error) {
	return m.CreatePlatformEndpointMock(arg1, arg2, arg3...)
}

// CreateTopic mocked funcition
func (m SNSClientMock) CreateTopic(arg1 context.Context, arg2 *sns.CreateTopicInput, arg3 ...func(*sns.Options)) (*sns.CreateTopicOutput, error) {
	return m.CreateTopicMock(arg1, arg2, arg3...)
}

// DeleteEndpoint mocked funcition
func (m SNSClientMock) DeleteEndpoint(arg1 context.Context, arg2 *sns.DeleteEndpointInput, arg3 ...func(*sns.Options)) (*sns.DeleteEndpointOutput, error) {
	return m.DeleteEndpointMock(arg1, arg2, arg3...)
}

// DeletePlatformApplication mocked funcition
func (m SNSClientMock) DeletePlatformApplication(arg1 context.Context, arg2 *sns.DeletePlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.DeletePlatformApplicationOutput, error) {
	return m.DeletePlatformApplicationMock(arg1, arg2, arg3...)
}

// DeleteTopic mocked funcition
func (m SNSClientMock) DeleteTopic(arg1 context.Context, arg2 *sns.DeleteTopicInput, arg3 ...func(*sns.Options)) (*sns.DeleteTopicOutput, error) {
	return m.DeleteTopicMock(arg1, arg2, arg3...)
}

// GetEndpointAttributes mocked funcition
func (m SNSClientMock) GetEndpointAttributes(arg1 context.Context, arg2 *sns.GetEndpointAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error) {
	return m.GetEndpointAttributesMock(arg1, arg2, arg3...)
}

// GetPlatformApplicationAttributes mocked funcition
func (m SNSClientMock) GetPlatformApplicationAttributes(arg1 context.Context, arg2 *sns.GetPlatformApplicationAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error) {
	return m.GetPlatformApplicationAttributesMock(arg1, arg2, arg3...)
}

// GetSMSAttributes mocked funcition
func (m SNSClientMock) GetSMSAttributes(arg1 context.Context, arg2 *sns.GetSMSAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error) {
	return m.GetSMSAttributesMock(arg1, arg2, arg3...)
}

// GetSubscriptionAttributes mocked funcition
func (m SNSClientMock) GetSubscriptionAttributes(arg1 context.Context, arg2 *sns.GetSubscriptionAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error) {
	return m.GetSubscriptionAttributesMock(arg1, arg2, arg3...)
}

// GetTopicAttributes mocked funcition
func (m SNSClientMock) GetTopicAttributes(arg1 context.Context, arg2 *sns.GetTopicAttributesInput, arg3 ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error) {
	return m.GetTopicAttributesMock(arg1, arg2, arg3...)
}

// ListEndpointsByPlatformApplication mocked funcition
func (m SNSClientMock) ListEndpointsByPlatformApplication(arg1 context.Context, arg2 *sns.ListEndpointsByPlatformApplicationInput, arg3 ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	return m.ListEndpointsByPlatformApplicationMock(arg1, arg2, arg3...)
}

// ListPhoneNumbersOptedOut mocked funcition
func (m SNSClientMock) ListPhoneNumbersOptedOut(arg1 context.Context, arg2 *sns.ListPhoneNumbersOptedOutInput, arg3 ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	return m.ListPhoneNumbersOptedOutMock(arg1, arg2, arg3...)
}

// ListPlatformApplications mocked funcition
func (m SNSClientMock) ListPlatformApplications(arg1 context.Context, arg2 *sns.ListPlatformApplicationsInput, arg3 ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error) {
	return m.ListPlatformApplicationsMock(arg1, arg2, arg3...)
}

// ListSubscriptions mocked funcition
func (m SNSClientMock) ListSubscriptions(arg1 context.Context, arg2 *sns.ListSubscriptionsInput, arg3 ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error) {
	return m.ListSubscriptionsMock(arg1, arg2, arg3...)
}

// ListSubscriptionsByTopic mocked funcition
func (m SNSClientMock) ListSubscriptionsByTopic(arg1 context.Context, arg2 *sns.ListSubscriptionsByTopicInput, arg3 ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error) {
	return m.ListSubscriptionsByTopicMock(arg1, arg2, arg3...)
}

// ListTagsForResource mocked funcition
func (m SNSClientMock) ListTagsForResource(arg1 context.Context, arg2 *sns.ListTagsForResourceInput, arg3 ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error) {
	return m.ListTagsForResourceMock(arg1, arg2, arg3...)
}

// ListTopics mocked funcition
func (m SNSClientMock) ListTopics(arg1 context.Context, arg2 *sns.ListTopicsInput, arg3 ...func(*sns.Options)) (*sns.ListTopicsOutput, error) {
	return m.ListTopicsMock(arg1, arg2, arg3...)
}

// OptInPhoneNumber mocked funcition
func (m SNSClientMock) OptInPhoneNumber(arg1 context.Context, arg2 *sns.OptInPhoneNumberInput, arg3 ...func(*sns.Options)) (*sns.OptInPhoneNumberOutput, error) {
	return m.OptInPhoneNumberMock(arg1, arg2, arg3...)
}

// Publish mocked funcition
func (m SNSClientMock) Publish(arg1 context.Context, arg2 *sns.PublishInput, arg3 ...func(*sns.Options)) (*sns.PublishOutput, error) {
	return m.PublishMock(arg1, arg2, arg3...)
}

// RemovePermission mocked funcition
func (m SNSClientMock) RemovePermission(arg1 context.Context, arg2 *sns.RemovePermissionInput, arg3 ...func(*sns.Options)) (*sns.RemovePermissionOutput, error) {
	return m.RemovePermissionMock(arg1, arg2, arg3...)
}

// SetEndpointAttributes mocked funcition
func (m SNSClientMock) SetEndpointAttributes(arg1 context.Context, arg2 *sns.SetEndpointAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetEndpointAttributesOutput, error) {
	return m.SetEndpointAttributesMock(arg1, arg2, arg3...)
}

// SetPlatformApplicationAttributes mocked funcition
func (m SNSClientMock) SetPlatformApplicationAttributes(arg1 context.Context, arg2 *sns.SetPlatformApplicationAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetPlatformApplicationAttributesOutput, error) {
	return m.SetPlatformApplicationAttributesMock(arg1, arg2, arg3...)
}

// SetSMSAttributes mocked funcition
func (m SNSClientMock) SetSMSAttributes(arg1 context.Context, arg2 *sns.SetSMSAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetSMSAttributesOutput, error) {
	return m.SetSMSAttributesMock(arg1, arg2, arg3...)
}

// SetSubscriptionAttributes mocked funcition
func (m SNSClientMock) SetSubscriptionAttributes(arg1 context.Context, arg2 *sns.SetSubscriptionAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetSubscriptionAttributesOutput, error) {
	return m.SetSubscriptionAttributesMock(arg1, arg2, arg3...)
}

// SetTopicAttributes mocked funcition
func (m SNSClientMock) SetTopicAttributes(arg1 context.Context, arg2 *sns.SetTopicAttributesInput, arg3 ...func(*sns.Options)) (*sns.SetTopicAttributesOutput, error) {
	return m.SetTopicAttributesMock(arg1, arg2, arg3...)
}

// Subscribe mocked funcition
func (m SNSClientMock) Subscribe(arg1 context.Context, arg2 *sns.SubscribeInput, arg3 ...func(*sns.Options)) (*sns.SubscribeOutput, error) {
	return m.SubscribeMock(arg1, arg2, arg3...)
}

// TagResource mocked funcition
func (m SNSClientMock) TagResource(arg1 context.Context, arg2 *sns.TagResourceInput, arg3 ...func(*sns.Options)) (*sns.TagResourceOutput, error) {
	return m.TagResourceMock(arg1, arg2, arg3...)
}

// Unsubscribe mocked funcition
func (m SNSClientMock) Unsubscribe(arg1 context.Context, arg2 *sns.UnsubscribeInput, arg3 ...func(*sns.Options)) (*sns.UnsubscribeOutput, error) {
	return m.UnsubscribeMock(arg1, arg2, arg3...)
}

// UntagResource mocked funcition
func (m SNSClientMock) UntagResource(arg1 context.Context, arg2 *sns.UntagResourceInput, arg3 ...func(*sns.Options)) (*sns.UntagResourceOutput, error) {
	return m.UntagResourceMock(arg1, arg2, arg3...)
}
