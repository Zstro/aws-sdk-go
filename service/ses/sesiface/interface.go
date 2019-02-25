// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sesiface provides an interface to enable mocking the Amazon Simple Email Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sesiface

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/service/ses"
)

// SESAPI provides an interface to enable mocking the
// ses.SES service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Email Service.
//    func myFunc(svc sesiface.SESAPI) bool {
//        // Make svc.CloneReceiptRuleSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ses.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSESClient struct {
//        sesiface.SESAPI
//    }
//    func (m *mockSESClient) CloneReceiptRuleSet(input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSESClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SESAPI interface {
	CloneReceiptRuleSet(*ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error)
	CloneReceiptRuleSetWithContext(aws.Context, *ses.CloneReceiptRuleSetInput, ...request.Option) (*ses.CloneReceiptRuleSetOutput, error)
	CloneReceiptRuleSetRequest(*ses.CloneReceiptRuleSetInput) (*request.Request, *ses.CloneReceiptRuleSetOutput)

	CreateConfigurationSet(*ses.CreateConfigurationSetInput) (*ses.CreateConfigurationSetOutput, error)
	CreateConfigurationSetWithContext(aws.Context, *ses.CreateConfigurationSetInput, ...request.Option) (*ses.CreateConfigurationSetOutput, error)
	CreateConfigurationSetRequest(*ses.CreateConfigurationSetInput) (*request.Request, *ses.CreateConfigurationSetOutput)

	CreateConfigurationSetEventDestination(*ses.CreateConfigurationSetEventDestinationInput) (*ses.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationWithContext(aws.Context, *ses.CreateConfigurationSetEventDestinationInput, ...request.Option) (*ses.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationRequest(*ses.CreateConfigurationSetEventDestinationInput) (*request.Request, *ses.CreateConfigurationSetEventDestinationOutput)

	CreateConfigurationSetTrackingOptions(*ses.CreateConfigurationSetTrackingOptionsInput) (*ses.CreateConfigurationSetTrackingOptionsOutput, error)
	CreateConfigurationSetTrackingOptionsWithContext(aws.Context, *ses.CreateConfigurationSetTrackingOptionsInput, ...request.Option) (*ses.CreateConfigurationSetTrackingOptionsOutput, error)
	CreateConfigurationSetTrackingOptionsRequest(*ses.CreateConfigurationSetTrackingOptionsInput) (*request.Request, *ses.CreateConfigurationSetTrackingOptionsOutput)

	CreateCustomVerificationEmailTemplate(*ses.CreateCustomVerificationEmailTemplateInput) (*ses.CreateCustomVerificationEmailTemplateOutput, error)
	CreateCustomVerificationEmailTemplateWithContext(aws.Context, *ses.CreateCustomVerificationEmailTemplateInput, ...request.Option) (*ses.CreateCustomVerificationEmailTemplateOutput, error)
	CreateCustomVerificationEmailTemplateRequest(*ses.CreateCustomVerificationEmailTemplateInput) (*request.Request, *ses.CreateCustomVerificationEmailTemplateOutput)

	CreateReceiptFilter(*ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error)
	CreateReceiptFilterWithContext(aws.Context, *ses.CreateReceiptFilterInput, ...request.Option) (*ses.CreateReceiptFilterOutput, error)
	CreateReceiptFilterRequest(*ses.CreateReceiptFilterInput) (*request.Request, *ses.CreateReceiptFilterOutput)

	CreateReceiptRule(*ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error)
	CreateReceiptRuleWithContext(aws.Context, *ses.CreateReceiptRuleInput, ...request.Option) (*ses.CreateReceiptRuleOutput, error)
	CreateReceiptRuleRequest(*ses.CreateReceiptRuleInput) (*request.Request, *ses.CreateReceiptRuleOutput)

	CreateReceiptRuleSet(*ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error)
	CreateReceiptRuleSetWithContext(aws.Context, *ses.CreateReceiptRuleSetInput, ...request.Option) (*ses.CreateReceiptRuleSetOutput, error)
	CreateReceiptRuleSetRequest(*ses.CreateReceiptRuleSetInput) (*request.Request, *ses.CreateReceiptRuleSetOutput)

	CreateTemplate(*ses.CreateTemplateInput) (*ses.CreateTemplateOutput, error)
	CreateTemplateWithContext(aws.Context, *ses.CreateTemplateInput, ...request.Option) (*ses.CreateTemplateOutput, error)
	CreateTemplateRequest(*ses.CreateTemplateInput) (*request.Request, *ses.CreateTemplateOutput)

	DeleteConfigurationSet(*ses.DeleteConfigurationSetInput) (*ses.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetWithContext(aws.Context, *ses.DeleteConfigurationSetInput, ...request.Option) (*ses.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetRequest(*ses.DeleteConfigurationSetInput) (*request.Request, *ses.DeleteConfigurationSetOutput)

	DeleteConfigurationSetEventDestination(*ses.DeleteConfigurationSetEventDestinationInput) (*ses.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationWithContext(aws.Context, *ses.DeleteConfigurationSetEventDestinationInput, ...request.Option) (*ses.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationRequest(*ses.DeleteConfigurationSetEventDestinationInput) (*request.Request, *ses.DeleteConfigurationSetEventDestinationOutput)

	DeleteConfigurationSetTrackingOptions(*ses.DeleteConfigurationSetTrackingOptionsInput) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error)
	DeleteConfigurationSetTrackingOptionsWithContext(aws.Context, *ses.DeleteConfigurationSetTrackingOptionsInput, ...request.Option) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error)
	DeleteConfigurationSetTrackingOptionsRequest(*ses.DeleteConfigurationSetTrackingOptionsInput) (*request.Request, *ses.DeleteConfigurationSetTrackingOptionsOutput)

	DeleteCustomVerificationEmailTemplate(*ses.DeleteCustomVerificationEmailTemplateInput) (*ses.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteCustomVerificationEmailTemplateWithContext(aws.Context, *ses.DeleteCustomVerificationEmailTemplateInput, ...request.Option) (*ses.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteCustomVerificationEmailTemplateRequest(*ses.DeleteCustomVerificationEmailTemplateInput) (*request.Request, *ses.DeleteCustomVerificationEmailTemplateOutput)

	DeleteIdentity(*ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)
	DeleteIdentityWithContext(aws.Context, *ses.DeleteIdentityInput, ...request.Option) (*ses.DeleteIdentityOutput, error)
	DeleteIdentityRequest(*ses.DeleteIdentityInput) (*request.Request, *ses.DeleteIdentityOutput)

	DeleteIdentityPolicy(*ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error)
	DeleteIdentityPolicyWithContext(aws.Context, *ses.DeleteIdentityPolicyInput, ...request.Option) (*ses.DeleteIdentityPolicyOutput, error)
	DeleteIdentityPolicyRequest(*ses.DeleteIdentityPolicyInput) (*request.Request, *ses.DeleteIdentityPolicyOutput)

	DeleteReceiptFilter(*ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error)
	DeleteReceiptFilterWithContext(aws.Context, *ses.DeleteReceiptFilterInput, ...request.Option) (*ses.DeleteReceiptFilterOutput, error)
	DeleteReceiptFilterRequest(*ses.DeleteReceiptFilterInput) (*request.Request, *ses.DeleteReceiptFilterOutput)

	DeleteReceiptRule(*ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error)
	DeleteReceiptRuleWithContext(aws.Context, *ses.DeleteReceiptRuleInput, ...request.Option) (*ses.DeleteReceiptRuleOutput, error)
	DeleteReceiptRuleRequest(*ses.DeleteReceiptRuleInput) (*request.Request, *ses.DeleteReceiptRuleOutput)

	DeleteReceiptRuleSet(*ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error)
	DeleteReceiptRuleSetWithContext(aws.Context, *ses.DeleteReceiptRuleSetInput, ...request.Option) (*ses.DeleteReceiptRuleSetOutput, error)
	DeleteReceiptRuleSetRequest(*ses.DeleteReceiptRuleSetInput) (*request.Request, *ses.DeleteReceiptRuleSetOutput)

	DeleteTemplate(*ses.DeleteTemplateInput) (*ses.DeleteTemplateOutput, error)
	DeleteTemplateWithContext(aws.Context, *ses.DeleteTemplateInput, ...request.Option) (*ses.DeleteTemplateOutput, error)
	DeleteTemplateRequest(*ses.DeleteTemplateInput) (*request.Request, *ses.DeleteTemplateOutput)

	DeleteVerifiedEmailAddress(*ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)
	DeleteVerifiedEmailAddressWithContext(aws.Context, *ses.DeleteVerifiedEmailAddressInput, ...request.Option) (*ses.DeleteVerifiedEmailAddressOutput, error)
	DeleteVerifiedEmailAddressRequest(*ses.DeleteVerifiedEmailAddressInput) (*request.Request, *ses.DeleteVerifiedEmailAddressOutput)

	DescribeActiveReceiptRuleSet(*ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error)
	DescribeActiveReceiptRuleSetWithContext(aws.Context, *ses.DescribeActiveReceiptRuleSetInput, ...request.Option) (*ses.DescribeActiveReceiptRuleSetOutput, error)
	DescribeActiveReceiptRuleSetRequest(*ses.DescribeActiveReceiptRuleSetInput) (*request.Request, *ses.DescribeActiveReceiptRuleSetOutput)

	DescribeConfigurationSet(*ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error)
	DescribeConfigurationSetWithContext(aws.Context, *ses.DescribeConfigurationSetInput, ...request.Option) (*ses.DescribeConfigurationSetOutput, error)
	DescribeConfigurationSetRequest(*ses.DescribeConfigurationSetInput) (*request.Request, *ses.DescribeConfigurationSetOutput)

	DescribeReceiptRule(*ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error)
	DescribeReceiptRuleWithContext(aws.Context, *ses.DescribeReceiptRuleInput, ...request.Option) (*ses.DescribeReceiptRuleOutput, error)
	DescribeReceiptRuleRequest(*ses.DescribeReceiptRuleInput) (*request.Request, *ses.DescribeReceiptRuleOutput)

	DescribeReceiptRuleSet(*ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error)
	DescribeReceiptRuleSetWithContext(aws.Context, *ses.DescribeReceiptRuleSetInput, ...request.Option) (*ses.DescribeReceiptRuleSetOutput, error)
	DescribeReceiptRuleSetRequest(*ses.DescribeReceiptRuleSetInput) (*request.Request, *ses.DescribeReceiptRuleSetOutput)

	GetAccountSendingEnabled(*ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error)
	GetAccountSendingEnabledWithContext(aws.Context, *ses.GetAccountSendingEnabledInput, ...request.Option) (*ses.GetAccountSendingEnabledOutput, error)
	GetAccountSendingEnabledRequest(*ses.GetAccountSendingEnabledInput) (*request.Request, *ses.GetAccountSendingEnabledOutput)

	GetCustomVerificationEmailTemplate(*ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error)
	GetCustomVerificationEmailTemplateWithContext(aws.Context, *ses.GetCustomVerificationEmailTemplateInput, ...request.Option) (*ses.GetCustomVerificationEmailTemplateOutput, error)
	GetCustomVerificationEmailTemplateRequest(*ses.GetCustomVerificationEmailTemplateInput) (*request.Request, *ses.GetCustomVerificationEmailTemplateOutput)

	GetIdentityDkimAttributes(*ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error)
	GetIdentityDkimAttributesWithContext(aws.Context, *ses.GetIdentityDkimAttributesInput, ...request.Option) (*ses.GetIdentityDkimAttributesOutput, error)
	GetIdentityDkimAttributesRequest(*ses.GetIdentityDkimAttributesInput) (*request.Request, *ses.GetIdentityDkimAttributesOutput)

	GetIdentityMailFromDomainAttributes(*ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error)
	GetIdentityMailFromDomainAttributesWithContext(aws.Context, *ses.GetIdentityMailFromDomainAttributesInput, ...request.Option) (*ses.GetIdentityMailFromDomainAttributesOutput, error)
	GetIdentityMailFromDomainAttributesRequest(*ses.GetIdentityMailFromDomainAttributesInput) (*request.Request, *ses.GetIdentityMailFromDomainAttributesOutput)

	GetIdentityNotificationAttributes(*ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)
	GetIdentityNotificationAttributesWithContext(aws.Context, *ses.GetIdentityNotificationAttributesInput, ...request.Option) (*ses.GetIdentityNotificationAttributesOutput, error)
	GetIdentityNotificationAttributesRequest(*ses.GetIdentityNotificationAttributesInput) (*request.Request, *ses.GetIdentityNotificationAttributesOutput)

	GetIdentityPolicies(*ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error)
	GetIdentityPoliciesWithContext(aws.Context, *ses.GetIdentityPoliciesInput, ...request.Option) (*ses.GetIdentityPoliciesOutput, error)
	GetIdentityPoliciesRequest(*ses.GetIdentityPoliciesInput) (*request.Request, *ses.GetIdentityPoliciesOutput)

	GetIdentityVerificationAttributes(*ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)
	GetIdentityVerificationAttributesWithContext(aws.Context, *ses.GetIdentityVerificationAttributesInput, ...request.Option) (*ses.GetIdentityVerificationAttributesOutput, error)
	GetIdentityVerificationAttributesRequest(*ses.GetIdentityVerificationAttributesInput) (*request.Request, *ses.GetIdentityVerificationAttributesOutput)

	GetSendQuota(*ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)
	GetSendQuotaWithContext(aws.Context, *ses.GetSendQuotaInput, ...request.Option) (*ses.GetSendQuotaOutput, error)
	GetSendQuotaRequest(*ses.GetSendQuotaInput) (*request.Request, *ses.GetSendQuotaOutput)

	GetSendStatistics(*ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)
	GetSendStatisticsWithContext(aws.Context, *ses.GetSendStatisticsInput, ...request.Option) (*ses.GetSendStatisticsOutput, error)
	GetSendStatisticsRequest(*ses.GetSendStatisticsInput) (*request.Request, *ses.GetSendStatisticsOutput)

	GetTemplate(*ses.GetTemplateInput) (*ses.GetTemplateOutput, error)
	GetTemplateWithContext(aws.Context, *ses.GetTemplateInput, ...request.Option) (*ses.GetTemplateOutput, error)
	GetTemplateRequest(*ses.GetTemplateInput) (*request.Request, *ses.GetTemplateOutput)

	ListConfigurationSets(*ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error)
	ListConfigurationSetsWithContext(aws.Context, *ses.ListConfigurationSetsInput, ...request.Option) (*ses.ListConfigurationSetsOutput, error)
	ListConfigurationSetsRequest(*ses.ListConfigurationSetsInput) (*request.Request, *ses.ListConfigurationSetsOutput)

	ListCustomVerificationEmailTemplates(*ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error)
	ListCustomVerificationEmailTemplatesWithContext(aws.Context, *ses.ListCustomVerificationEmailTemplatesInput, ...request.Option) (*ses.ListCustomVerificationEmailTemplatesOutput, error)
	ListCustomVerificationEmailTemplatesRequest(*ses.ListCustomVerificationEmailTemplatesInput) (*request.Request, *ses.ListCustomVerificationEmailTemplatesOutput)

	ListCustomVerificationEmailTemplatesPages(*ses.ListCustomVerificationEmailTemplatesInput, func(*ses.ListCustomVerificationEmailTemplatesOutput, bool) bool) error
	ListCustomVerificationEmailTemplatesPagesWithContext(aws.Context, *ses.ListCustomVerificationEmailTemplatesInput, func(*ses.ListCustomVerificationEmailTemplatesOutput, bool) bool, ...request.Option) error

	ListIdentities(*ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)
	ListIdentitiesWithContext(aws.Context, *ses.ListIdentitiesInput, ...request.Option) (*ses.ListIdentitiesOutput, error)
	ListIdentitiesRequest(*ses.ListIdentitiesInput) (*request.Request, *ses.ListIdentitiesOutput)

	ListIdentitiesPages(*ses.ListIdentitiesInput, func(*ses.ListIdentitiesOutput, bool) bool) error
	ListIdentitiesPagesWithContext(aws.Context, *ses.ListIdentitiesInput, func(*ses.ListIdentitiesOutput, bool) bool, ...request.Option) error

	ListIdentityPolicies(*ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error)
	ListIdentityPoliciesWithContext(aws.Context, *ses.ListIdentityPoliciesInput, ...request.Option) (*ses.ListIdentityPoliciesOutput, error)
	ListIdentityPoliciesRequest(*ses.ListIdentityPoliciesInput) (*request.Request, *ses.ListIdentityPoliciesOutput)

	ListReceiptFilters(*ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error)
	ListReceiptFiltersWithContext(aws.Context, *ses.ListReceiptFiltersInput, ...request.Option) (*ses.ListReceiptFiltersOutput, error)
	ListReceiptFiltersRequest(*ses.ListReceiptFiltersInput) (*request.Request, *ses.ListReceiptFiltersOutput)

	ListReceiptRuleSets(*ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error)
	ListReceiptRuleSetsWithContext(aws.Context, *ses.ListReceiptRuleSetsInput, ...request.Option) (*ses.ListReceiptRuleSetsOutput, error)
	ListReceiptRuleSetsRequest(*ses.ListReceiptRuleSetsInput) (*request.Request, *ses.ListReceiptRuleSetsOutput)

	ListTemplates(*ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error)
	ListTemplatesWithContext(aws.Context, *ses.ListTemplatesInput, ...request.Option) (*ses.ListTemplatesOutput, error)
	ListTemplatesRequest(*ses.ListTemplatesInput) (*request.Request, *ses.ListTemplatesOutput)

	ListVerifiedEmailAddresses(*ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)
	ListVerifiedEmailAddressesWithContext(aws.Context, *ses.ListVerifiedEmailAddressesInput, ...request.Option) (*ses.ListVerifiedEmailAddressesOutput, error)
	ListVerifiedEmailAddressesRequest(*ses.ListVerifiedEmailAddressesInput) (*request.Request, *ses.ListVerifiedEmailAddressesOutput)

	PutIdentityPolicy(*ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error)
	PutIdentityPolicyWithContext(aws.Context, *ses.PutIdentityPolicyInput, ...request.Option) (*ses.PutIdentityPolicyOutput, error)
	PutIdentityPolicyRequest(*ses.PutIdentityPolicyInput) (*request.Request, *ses.PutIdentityPolicyOutput)

	ReorderReceiptRuleSet(*ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error)
	ReorderReceiptRuleSetWithContext(aws.Context, *ses.ReorderReceiptRuleSetInput, ...request.Option) (*ses.ReorderReceiptRuleSetOutput, error)
	ReorderReceiptRuleSetRequest(*ses.ReorderReceiptRuleSetInput) (*request.Request, *ses.ReorderReceiptRuleSetOutput)

	SendBounce(*ses.SendBounceInput) (*ses.SendBounceOutput, error)
	SendBounceWithContext(aws.Context, *ses.SendBounceInput, ...request.Option) (*ses.SendBounceOutput, error)
	SendBounceRequest(*ses.SendBounceInput) (*request.Request, *ses.SendBounceOutput)

	SendBulkTemplatedEmail(*ses.SendBulkTemplatedEmailInput) (*ses.SendBulkTemplatedEmailOutput, error)
	SendBulkTemplatedEmailWithContext(aws.Context, *ses.SendBulkTemplatedEmailInput, ...request.Option) (*ses.SendBulkTemplatedEmailOutput, error)
	SendBulkTemplatedEmailRequest(*ses.SendBulkTemplatedEmailInput) (*request.Request, *ses.SendBulkTemplatedEmailOutput)

	SendCustomVerificationEmail(*ses.SendCustomVerificationEmailInput) (*ses.SendCustomVerificationEmailOutput, error)
	SendCustomVerificationEmailWithContext(aws.Context, *ses.SendCustomVerificationEmailInput, ...request.Option) (*ses.SendCustomVerificationEmailOutput, error)
	SendCustomVerificationEmailRequest(*ses.SendCustomVerificationEmailInput) (*request.Request, *ses.SendCustomVerificationEmailOutput)

	SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error)
	SendEmailWithContext(aws.Context, *ses.SendEmailInput, ...request.Option) (*ses.SendEmailOutput, error)
	SendEmailRequest(*ses.SendEmailInput) (*request.Request, *ses.SendEmailOutput)

	SendRawEmail(*ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)
	SendRawEmailWithContext(aws.Context, *ses.SendRawEmailInput, ...request.Option) (*ses.SendRawEmailOutput, error)
	SendRawEmailRequest(*ses.SendRawEmailInput) (*request.Request, *ses.SendRawEmailOutput)

	SendTemplatedEmail(*ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error)
	SendTemplatedEmailWithContext(aws.Context, *ses.SendTemplatedEmailInput, ...request.Option) (*ses.SendTemplatedEmailOutput, error)
	SendTemplatedEmailRequest(*ses.SendTemplatedEmailInput) (*request.Request, *ses.SendTemplatedEmailOutput)

	SetActiveReceiptRuleSet(*ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error)
	SetActiveReceiptRuleSetWithContext(aws.Context, *ses.SetActiveReceiptRuleSetInput, ...request.Option) (*ses.SetActiveReceiptRuleSetOutput, error)
	SetActiveReceiptRuleSetRequest(*ses.SetActiveReceiptRuleSetInput) (*request.Request, *ses.SetActiveReceiptRuleSetOutput)

	SetIdentityDkimEnabled(*ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error)
	SetIdentityDkimEnabledWithContext(aws.Context, *ses.SetIdentityDkimEnabledInput, ...request.Option) (*ses.SetIdentityDkimEnabledOutput, error)
	SetIdentityDkimEnabledRequest(*ses.SetIdentityDkimEnabledInput) (*request.Request, *ses.SetIdentityDkimEnabledOutput)

	SetIdentityFeedbackForwardingEnabled(*ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)
	SetIdentityFeedbackForwardingEnabledWithContext(aws.Context, *ses.SetIdentityFeedbackForwardingEnabledInput, ...request.Option) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)
	SetIdentityFeedbackForwardingEnabledRequest(*ses.SetIdentityFeedbackForwardingEnabledInput) (*request.Request, *ses.SetIdentityFeedbackForwardingEnabledOutput)

	SetIdentityHeadersInNotificationsEnabled(*ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error)
	SetIdentityHeadersInNotificationsEnabledWithContext(aws.Context, *ses.SetIdentityHeadersInNotificationsEnabledInput, ...request.Option) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error)
	SetIdentityHeadersInNotificationsEnabledRequest(*ses.SetIdentityHeadersInNotificationsEnabledInput) (*request.Request, *ses.SetIdentityHeadersInNotificationsEnabledOutput)

	SetIdentityMailFromDomain(*ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error)
	SetIdentityMailFromDomainWithContext(aws.Context, *ses.SetIdentityMailFromDomainInput, ...request.Option) (*ses.SetIdentityMailFromDomainOutput, error)
	SetIdentityMailFromDomainRequest(*ses.SetIdentityMailFromDomainInput) (*request.Request, *ses.SetIdentityMailFromDomainOutput)

	SetIdentityNotificationTopic(*ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)
	SetIdentityNotificationTopicWithContext(aws.Context, *ses.SetIdentityNotificationTopicInput, ...request.Option) (*ses.SetIdentityNotificationTopicOutput, error)
	SetIdentityNotificationTopicRequest(*ses.SetIdentityNotificationTopicInput) (*request.Request, *ses.SetIdentityNotificationTopicOutput)

	SetReceiptRulePosition(*ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error)
	SetReceiptRulePositionWithContext(aws.Context, *ses.SetReceiptRulePositionInput, ...request.Option) (*ses.SetReceiptRulePositionOutput, error)
	SetReceiptRulePositionRequest(*ses.SetReceiptRulePositionInput) (*request.Request, *ses.SetReceiptRulePositionOutput)

	TestRenderTemplate(*ses.TestRenderTemplateInput) (*ses.TestRenderTemplateOutput, error)
	TestRenderTemplateWithContext(aws.Context, *ses.TestRenderTemplateInput, ...request.Option) (*ses.TestRenderTemplateOutput, error)
	TestRenderTemplateRequest(*ses.TestRenderTemplateInput) (*request.Request, *ses.TestRenderTemplateOutput)

	UpdateAccountSendingEnabled(*ses.UpdateAccountSendingEnabledInput) (*ses.UpdateAccountSendingEnabledOutput, error)
	UpdateAccountSendingEnabledWithContext(aws.Context, *ses.UpdateAccountSendingEnabledInput, ...request.Option) (*ses.UpdateAccountSendingEnabledOutput, error)
	UpdateAccountSendingEnabledRequest(*ses.UpdateAccountSendingEnabledInput) (*request.Request, *ses.UpdateAccountSendingEnabledOutput)

	UpdateConfigurationSetEventDestination(*ses.UpdateConfigurationSetEventDestinationInput) (*ses.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationWithContext(aws.Context, *ses.UpdateConfigurationSetEventDestinationInput, ...request.Option) (*ses.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationRequest(*ses.UpdateConfigurationSetEventDestinationInput) (*request.Request, *ses.UpdateConfigurationSetEventDestinationOutput)

	UpdateConfigurationSetReputationMetricsEnabled(*ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error)
	UpdateConfigurationSetReputationMetricsEnabledWithContext(aws.Context, *ses.UpdateConfigurationSetReputationMetricsEnabledInput, ...request.Option) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error)
	UpdateConfigurationSetReputationMetricsEnabledRequest(*ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*request.Request, *ses.UpdateConfigurationSetReputationMetricsEnabledOutput)

	UpdateConfigurationSetSendingEnabled(*ses.UpdateConfigurationSetSendingEnabledInput) (*ses.UpdateConfigurationSetSendingEnabledOutput, error)
	UpdateConfigurationSetSendingEnabledWithContext(aws.Context, *ses.UpdateConfigurationSetSendingEnabledInput, ...request.Option) (*ses.UpdateConfigurationSetSendingEnabledOutput, error)
	UpdateConfigurationSetSendingEnabledRequest(*ses.UpdateConfigurationSetSendingEnabledInput) (*request.Request, *ses.UpdateConfigurationSetSendingEnabledOutput)

	UpdateConfigurationSetTrackingOptions(*ses.UpdateConfigurationSetTrackingOptionsInput) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error)
	UpdateConfigurationSetTrackingOptionsWithContext(aws.Context, *ses.UpdateConfigurationSetTrackingOptionsInput, ...request.Option) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error)
	UpdateConfigurationSetTrackingOptionsRequest(*ses.UpdateConfigurationSetTrackingOptionsInput) (*request.Request, *ses.UpdateConfigurationSetTrackingOptionsOutput)

	UpdateCustomVerificationEmailTemplate(*ses.UpdateCustomVerificationEmailTemplateInput) (*ses.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateCustomVerificationEmailTemplateWithContext(aws.Context, *ses.UpdateCustomVerificationEmailTemplateInput, ...request.Option) (*ses.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateCustomVerificationEmailTemplateRequest(*ses.UpdateCustomVerificationEmailTemplateInput) (*request.Request, *ses.UpdateCustomVerificationEmailTemplateOutput)

	UpdateReceiptRule(*ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error)
	UpdateReceiptRuleWithContext(aws.Context, *ses.UpdateReceiptRuleInput, ...request.Option) (*ses.UpdateReceiptRuleOutput, error)
	UpdateReceiptRuleRequest(*ses.UpdateReceiptRuleInput) (*request.Request, *ses.UpdateReceiptRuleOutput)

	UpdateTemplate(*ses.UpdateTemplateInput) (*ses.UpdateTemplateOutput, error)
	UpdateTemplateWithContext(aws.Context, *ses.UpdateTemplateInput, ...request.Option) (*ses.UpdateTemplateOutput, error)
	UpdateTemplateRequest(*ses.UpdateTemplateInput) (*request.Request, *ses.UpdateTemplateOutput)

	VerifyDomainDkim(*ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error)
	VerifyDomainDkimWithContext(aws.Context, *ses.VerifyDomainDkimInput, ...request.Option) (*ses.VerifyDomainDkimOutput, error)
	VerifyDomainDkimRequest(*ses.VerifyDomainDkimInput) (*request.Request, *ses.VerifyDomainDkimOutput)

	VerifyDomainIdentity(*ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)
	VerifyDomainIdentityWithContext(aws.Context, *ses.VerifyDomainIdentityInput, ...request.Option) (*ses.VerifyDomainIdentityOutput, error)
	VerifyDomainIdentityRequest(*ses.VerifyDomainIdentityInput) (*request.Request, *ses.VerifyDomainIdentityOutput)

	VerifyEmailAddress(*ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)
	VerifyEmailAddressWithContext(aws.Context, *ses.VerifyEmailAddressInput, ...request.Option) (*ses.VerifyEmailAddressOutput, error)
	VerifyEmailAddressRequest(*ses.VerifyEmailAddressInput) (*request.Request, *ses.VerifyEmailAddressOutput)

	VerifyEmailIdentity(*ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)
	VerifyEmailIdentityWithContext(aws.Context, *ses.VerifyEmailIdentityInput, ...request.Option) (*ses.VerifyEmailIdentityOutput, error)
	VerifyEmailIdentityRequest(*ses.VerifyEmailIdentityInput) (*request.Request, *ses.VerifyEmailIdentityOutput)

	WaitUntilIdentityExists(*ses.GetIdentityVerificationAttributesInput) error
	WaitUntilIdentityExistsWithContext(aws.Context, *ses.GetIdentityVerificationAttributesInput, ...request.WaiterOption) error
}

var _ SESAPI = (*ses.SES)(nil)
