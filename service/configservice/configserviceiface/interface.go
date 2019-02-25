// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package configserviceiface provides an interface to enable mocking the AWS Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package configserviceiface

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/service/configservice"
)

// ConfigServiceAPI provides an interface to enable mocking the
// configservice.ConfigService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Config.
//    func myFunc(svc configserviceiface.ConfigServiceAPI) bool {
//        // Make svc.BatchGetAggregateResourceConfig request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := configservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConfigServiceClient struct {
//        configserviceiface.ConfigServiceAPI
//    }
//    func (m *mockConfigServiceClient) BatchGetAggregateResourceConfig(input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConfigServiceClient{}
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
type ConfigServiceAPI interface {
	BatchGetAggregateResourceConfig(*configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error)
	BatchGetAggregateResourceConfigWithContext(aws.Context, *configservice.BatchGetAggregateResourceConfigInput, ...request.Option) (*configservice.BatchGetAggregateResourceConfigOutput, error)
	BatchGetAggregateResourceConfigRequest(*configservice.BatchGetAggregateResourceConfigInput) (*request.Request, *configservice.BatchGetAggregateResourceConfigOutput)

	BatchGetResourceConfig(*configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error)
	BatchGetResourceConfigWithContext(aws.Context, *configservice.BatchGetResourceConfigInput, ...request.Option) (*configservice.BatchGetResourceConfigOutput, error)
	BatchGetResourceConfigRequest(*configservice.BatchGetResourceConfigInput) (*request.Request, *configservice.BatchGetResourceConfigOutput)

	DeleteAggregationAuthorization(*configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error)
	DeleteAggregationAuthorizationWithContext(aws.Context, *configservice.DeleteAggregationAuthorizationInput, ...request.Option) (*configservice.DeleteAggregationAuthorizationOutput, error)
	DeleteAggregationAuthorizationRequest(*configservice.DeleteAggregationAuthorizationInput) (*request.Request, *configservice.DeleteAggregationAuthorizationOutput)

	DeleteConfigRule(*configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error)
	DeleteConfigRuleWithContext(aws.Context, *configservice.DeleteConfigRuleInput, ...request.Option) (*configservice.DeleteConfigRuleOutput, error)
	DeleteConfigRuleRequest(*configservice.DeleteConfigRuleInput) (*request.Request, *configservice.DeleteConfigRuleOutput)

	DeleteConfigurationAggregator(*configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error)
	DeleteConfigurationAggregatorWithContext(aws.Context, *configservice.DeleteConfigurationAggregatorInput, ...request.Option) (*configservice.DeleteConfigurationAggregatorOutput, error)
	DeleteConfigurationAggregatorRequest(*configservice.DeleteConfigurationAggregatorInput) (*request.Request, *configservice.DeleteConfigurationAggregatorOutput)

	DeleteConfigurationRecorder(*configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error)
	DeleteConfigurationRecorderWithContext(aws.Context, *configservice.DeleteConfigurationRecorderInput, ...request.Option) (*configservice.DeleteConfigurationRecorderOutput, error)
	DeleteConfigurationRecorderRequest(*configservice.DeleteConfigurationRecorderInput) (*request.Request, *configservice.DeleteConfigurationRecorderOutput)

	DeleteDeliveryChannel(*configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)
	DeleteDeliveryChannelWithContext(aws.Context, *configservice.DeleteDeliveryChannelInput, ...request.Option) (*configservice.DeleteDeliveryChannelOutput, error)
	DeleteDeliveryChannelRequest(*configservice.DeleteDeliveryChannelInput) (*request.Request, *configservice.DeleteDeliveryChannelOutput)

	DeleteEvaluationResults(*configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error)
	DeleteEvaluationResultsWithContext(aws.Context, *configservice.DeleteEvaluationResultsInput, ...request.Option) (*configservice.DeleteEvaluationResultsOutput, error)
	DeleteEvaluationResultsRequest(*configservice.DeleteEvaluationResultsInput) (*request.Request, *configservice.DeleteEvaluationResultsOutput)

	DeletePendingAggregationRequest(*configservice.DeletePendingAggregationRequestInput) (*configservice.DeletePendingAggregationRequestOutput, error)
	DeletePendingAggregationRequestWithContext(aws.Context, *configservice.DeletePendingAggregationRequestInput, ...request.Option) (*configservice.DeletePendingAggregationRequestOutput, error)
	DeletePendingAggregationRequestRequest(*configservice.DeletePendingAggregationRequestInput) (*request.Request, *configservice.DeletePendingAggregationRequestOutput)

	DeleteRetentionConfiguration(*configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error)
	DeleteRetentionConfigurationWithContext(aws.Context, *configservice.DeleteRetentionConfigurationInput, ...request.Option) (*configservice.DeleteRetentionConfigurationOutput, error)
	DeleteRetentionConfigurationRequest(*configservice.DeleteRetentionConfigurationInput) (*request.Request, *configservice.DeleteRetentionConfigurationOutput)

	DeliverConfigSnapshot(*configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)
	DeliverConfigSnapshotWithContext(aws.Context, *configservice.DeliverConfigSnapshotInput, ...request.Option) (*configservice.DeliverConfigSnapshotOutput, error)
	DeliverConfigSnapshotRequest(*configservice.DeliverConfigSnapshotInput) (*request.Request, *configservice.DeliverConfigSnapshotOutput)

	DescribeAggregateComplianceByConfigRules(*configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
	DescribeAggregateComplianceByConfigRulesWithContext(aws.Context, *configservice.DescribeAggregateComplianceByConfigRulesInput, ...request.Option) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
	DescribeAggregateComplianceByConfigRulesRequest(*configservice.DescribeAggregateComplianceByConfigRulesInput) (*request.Request, *configservice.DescribeAggregateComplianceByConfigRulesOutput)

	DescribeAggregationAuthorizations(*configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error)
	DescribeAggregationAuthorizationsWithContext(aws.Context, *configservice.DescribeAggregationAuthorizationsInput, ...request.Option) (*configservice.DescribeAggregationAuthorizationsOutput, error)
	DescribeAggregationAuthorizationsRequest(*configservice.DescribeAggregationAuthorizationsInput) (*request.Request, *configservice.DescribeAggregationAuthorizationsOutput)

	DescribeComplianceByConfigRule(*configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByConfigRuleWithContext(aws.Context, *configservice.DescribeComplianceByConfigRuleInput, ...request.Option) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByConfigRuleRequest(*configservice.DescribeComplianceByConfigRuleInput) (*request.Request, *configservice.DescribeComplianceByConfigRuleOutput)

	DescribeComplianceByResource(*configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeComplianceByResourceWithContext(aws.Context, *configservice.DescribeComplianceByResourceInput, ...request.Option) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeComplianceByResourceRequest(*configservice.DescribeComplianceByResourceInput) (*request.Request, *configservice.DescribeComplianceByResourceOutput)

	DescribeConfigRuleEvaluationStatus(*configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRuleEvaluationStatusWithContext(aws.Context, *configservice.DescribeConfigRuleEvaluationStatusInput, ...request.Option) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRuleEvaluationStatusRequest(*configservice.DescribeConfigRuleEvaluationStatusInput) (*request.Request, *configservice.DescribeConfigRuleEvaluationStatusOutput)

	DescribeConfigRules(*configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigRulesWithContext(aws.Context, *configservice.DescribeConfigRulesInput, ...request.Option) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigRulesRequest(*configservice.DescribeConfigRulesInput) (*request.Request, *configservice.DescribeConfigRulesOutput)

	DescribeConfigurationAggregatorSourcesStatus(*configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
	DescribeConfigurationAggregatorSourcesStatusWithContext(aws.Context, *configservice.DescribeConfigurationAggregatorSourcesStatusInput, ...request.Option) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
	DescribeConfigurationAggregatorSourcesStatusRequest(*configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*request.Request, *configservice.DescribeConfigurationAggregatorSourcesStatusOutput)

	DescribeConfigurationAggregators(*configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error)
	DescribeConfigurationAggregatorsWithContext(aws.Context, *configservice.DescribeConfigurationAggregatorsInput, ...request.Option) (*configservice.DescribeConfigurationAggregatorsOutput, error)
	DescribeConfigurationAggregatorsRequest(*configservice.DescribeConfigurationAggregatorsInput) (*request.Request, *configservice.DescribeConfigurationAggregatorsOutput)

	DescribeConfigurationRecorderStatus(*configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorderStatusWithContext(aws.Context, *configservice.DescribeConfigurationRecorderStatusInput, ...request.Option) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorderStatusRequest(*configservice.DescribeConfigurationRecorderStatusInput) (*request.Request, *configservice.DescribeConfigurationRecorderStatusOutput)

	DescribeConfigurationRecorders(*configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecordersWithContext(aws.Context, *configservice.DescribeConfigurationRecordersInput, ...request.Option) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecordersRequest(*configservice.DescribeConfigurationRecordersInput) (*request.Request, *configservice.DescribeConfigurationRecordersOutput)

	DescribeDeliveryChannelStatus(*configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannelStatusWithContext(aws.Context, *configservice.DescribeDeliveryChannelStatusInput, ...request.Option) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannelStatusRequest(*configservice.DescribeDeliveryChannelStatusInput) (*request.Request, *configservice.DescribeDeliveryChannelStatusOutput)

	DescribeDeliveryChannels(*configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeDeliveryChannelsWithContext(aws.Context, *configservice.DescribeDeliveryChannelsInput, ...request.Option) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeDeliveryChannelsRequest(*configservice.DescribeDeliveryChannelsInput) (*request.Request, *configservice.DescribeDeliveryChannelsOutput)

	DescribePendingAggregationRequests(*configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error)
	DescribePendingAggregationRequestsWithContext(aws.Context, *configservice.DescribePendingAggregationRequestsInput, ...request.Option) (*configservice.DescribePendingAggregationRequestsOutput, error)
	DescribePendingAggregationRequestsRequest(*configservice.DescribePendingAggregationRequestsInput) (*request.Request, *configservice.DescribePendingAggregationRequestsOutput)

	DescribeRetentionConfigurations(*configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error)
	DescribeRetentionConfigurationsWithContext(aws.Context, *configservice.DescribeRetentionConfigurationsInput, ...request.Option) (*configservice.DescribeRetentionConfigurationsOutput, error)
	DescribeRetentionConfigurationsRequest(*configservice.DescribeRetentionConfigurationsInput) (*request.Request, *configservice.DescribeRetentionConfigurationsOutput)

	GetAggregateComplianceDetailsByConfigRule(*configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
	GetAggregateComplianceDetailsByConfigRuleWithContext(aws.Context, *configservice.GetAggregateComplianceDetailsByConfigRuleInput, ...request.Option) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
	GetAggregateComplianceDetailsByConfigRuleRequest(*configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*request.Request, *configservice.GetAggregateComplianceDetailsByConfigRuleOutput)

	GetAggregateConfigRuleComplianceSummary(*configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
	GetAggregateConfigRuleComplianceSummaryWithContext(aws.Context, *configservice.GetAggregateConfigRuleComplianceSummaryInput, ...request.Option) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
	GetAggregateConfigRuleComplianceSummaryRequest(*configservice.GetAggregateConfigRuleComplianceSummaryInput) (*request.Request, *configservice.GetAggregateConfigRuleComplianceSummaryOutput)

	GetAggregateDiscoveredResourceCounts(*configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
	GetAggregateDiscoveredResourceCountsWithContext(aws.Context, *configservice.GetAggregateDiscoveredResourceCountsInput, ...request.Option) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
	GetAggregateDiscoveredResourceCountsRequest(*configservice.GetAggregateDiscoveredResourceCountsInput) (*request.Request, *configservice.GetAggregateDiscoveredResourceCountsOutput)

	GetAggregateResourceConfig(*configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error)
	GetAggregateResourceConfigWithContext(aws.Context, *configservice.GetAggregateResourceConfigInput, ...request.Option) (*configservice.GetAggregateResourceConfigOutput, error)
	GetAggregateResourceConfigRequest(*configservice.GetAggregateResourceConfigInput) (*request.Request, *configservice.GetAggregateResourceConfigOutput)

	GetComplianceDetailsByConfigRule(*configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByConfigRuleWithContext(aws.Context, *configservice.GetComplianceDetailsByConfigRuleInput, ...request.Option) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByConfigRuleRequest(*configservice.GetComplianceDetailsByConfigRuleInput) (*request.Request, *configservice.GetComplianceDetailsByConfigRuleOutput)

	GetComplianceDetailsByResource(*configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceDetailsByResourceWithContext(aws.Context, *configservice.GetComplianceDetailsByResourceInput, ...request.Option) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceDetailsByResourceRequest(*configservice.GetComplianceDetailsByResourceInput) (*request.Request, *configservice.GetComplianceDetailsByResourceOutput)

	GetComplianceSummaryByConfigRule(*configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByConfigRuleWithContext(aws.Context, *configservice.GetComplianceSummaryByConfigRuleInput, ...request.Option) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByConfigRuleRequest(*configservice.GetComplianceSummaryByConfigRuleInput) (*request.Request, *configservice.GetComplianceSummaryByConfigRuleOutput)

	GetComplianceSummaryByResourceType(*configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetComplianceSummaryByResourceTypeWithContext(aws.Context, *configservice.GetComplianceSummaryByResourceTypeInput, ...request.Option) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetComplianceSummaryByResourceTypeRequest(*configservice.GetComplianceSummaryByResourceTypeInput) (*request.Request, *configservice.GetComplianceSummaryByResourceTypeOutput)

	GetDiscoveredResourceCounts(*configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetDiscoveredResourceCountsWithContext(aws.Context, *configservice.GetDiscoveredResourceCountsInput, ...request.Option) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetDiscoveredResourceCountsRequest(*configservice.GetDiscoveredResourceCountsInput) (*request.Request, *configservice.GetDiscoveredResourceCountsOutput)

	GetResourceConfigHistory(*configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)
	GetResourceConfigHistoryWithContext(aws.Context, *configservice.GetResourceConfigHistoryInput, ...request.Option) (*configservice.GetResourceConfigHistoryOutput, error)
	GetResourceConfigHistoryRequest(*configservice.GetResourceConfigHistoryInput) (*request.Request, *configservice.GetResourceConfigHistoryOutput)

	GetResourceConfigHistoryPages(*configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool) error
	GetResourceConfigHistoryPagesWithContext(aws.Context, *configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool, ...request.Option) error

	ListAggregateDiscoveredResources(*configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
	ListAggregateDiscoveredResourcesWithContext(aws.Context, *configservice.ListAggregateDiscoveredResourcesInput, ...request.Option) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
	ListAggregateDiscoveredResourcesRequest(*configservice.ListAggregateDiscoveredResourcesInput) (*request.Request, *configservice.ListAggregateDiscoveredResourcesOutput)

	ListDiscoveredResources(*configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesWithContext(aws.Context, *configservice.ListDiscoveredResourcesInput, ...request.Option) (*configservice.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesRequest(*configservice.ListDiscoveredResourcesInput) (*request.Request, *configservice.ListDiscoveredResourcesOutput)

	PutAggregationAuthorization(*configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error)
	PutAggregationAuthorizationWithContext(aws.Context, *configservice.PutAggregationAuthorizationInput, ...request.Option) (*configservice.PutAggregationAuthorizationOutput, error)
	PutAggregationAuthorizationRequest(*configservice.PutAggregationAuthorizationInput) (*request.Request, *configservice.PutAggregationAuthorizationOutput)

	PutConfigRule(*configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error)
	PutConfigRuleWithContext(aws.Context, *configservice.PutConfigRuleInput, ...request.Option) (*configservice.PutConfigRuleOutput, error)
	PutConfigRuleRequest(*configservice.PutConfigRuleInput) (*request.Request, *configservice.PutConfigRuleOutput)

	PutConfigurationAggregator(*configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error)
	PutConfigurationAggregatorWithContext(aws.Context, *configservice.PutConfigurationAggregatorInput, ...request.Option) (*configservice.PutConfigurationAggregatorOutput, error)
	PutConfigurationAggregatorRequest(*configservice.PutConfigurationAggregatorInput) (*request.Request, *configservice.PutConfigurationAggregatorOutput)

	PutConfigurationRecorder(*configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)
	PutConfigurationRecorderWithContext(aws.Context, *configservice.PutConfigurationRecorderInput, ...request.Option) (*configservice.PutConfigurationRecorderOutput, error)
	PutConfigurationRecorderRequest(*configservice.PutConfigurationRecorderInput) (*request.Request, *configservice.PutConfigurationRecorderOutput)

	PutDeliveryChannel(*configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)
	PutDeliveryChannelWithContext(aws.Context, *configservice.PutDeliveryChannelInput, ...request.Option) (*configservice.PutDeliveryChannelOutput, error)
	PutDeliveryChannelRequest(*configservice.PutDeliveryChannelInput) (*request.Request, *configservice.PutDeliveryChannelOutput)

	PutEvaluations(*configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
	PutEvaluationsWithContext(aws.Context, *configservice.PutEvaluationsInput, ...request.Option) (*configservice.PutEvaluationsOutput, error)
	PutEvaluationsRequest(*configservice.PutEvaluationsInput) (*request.Request, *configservice.PutEvaluationsOutput)

	PutRetentionConfiguration(*configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error)
	PutRetentionConfigurationWithContext(aws.Context, *configservice.PutRetentionConfigurationInput, ...request.Option) (*configservice.PutRetentionConfigurationOutput, error)
	PutRetentionConfigurationRequest(*configservice.PutRetentionConfigurationInput) (*request.Request, *configservice.PutRetentionConfigurationOutput)

	StartConfigRulesEvaluation(*configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error)
	StartConfigRulesEvaluationWithContext(aws.Context, *configservice.StartConfigRulesEvaluationInput, ...request.Option) (*configservice.StartConfigRulesEvaluationOutput, error)
	StartConfigRulesEvaluationRequest(*configservice.StartConfigRulesEvaluationInput) (*request.Request, *configservice.StartConfigRulesEvaluationOutput)

	StartConfigurationRecorder(*configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)
	StartConfigurationRecorderWithContext(aws.Context, *configservice.StartConfigurationRecorderInput, ...request.Option) (*configservice.StartConfigurationRecorderOutput, error)
	StartConfigurationRecorderRequest(*configservice.StartConfigurationRecorderInput) (*request.Request, *configservice.StartConfigurationRecorderOutput)

	StopConfigurationRecorder(*configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
	StopConfigurationRecorderWithContext(aws.Context, *configservice.StopConfigurationRecorderInput, ...request.Option) (*configservice.StopConfigurationRecorderOutput, error)
	StopConfigurationRecorderRequest(*configservice.StopConfigurationRecorderInput) (*request.Request, *configservice.StopConfigurationRecorderOutput)
}

var _ ConfigServiceAPI = (*configservice.ConfigService)(nil)
