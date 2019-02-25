// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package resourcegroupstaggingapiiface provides an interface to enable mocking the AWS Resource Groups Tagging API service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package resourcegroupstaggingapiiface

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/service/resourcegroupstaggingapi"
)

// ResourceGroupsTaggingAPIAPI provides an interface to enable mocking the
// resourcegroupstaggingapi.ResourceGroupsTaggingAPI service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Resource Groups Tagging API.
//    func myFunc(svc resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI) bool {
//        // Make svc.GetResources request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := resourcegroupstaggingapi.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockResourceGroupsTaggingAPIClient struct {
//        resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
//    }
//    func (m *mockResourceGroupsTaggingAPIClient) GetResources(input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockResourceGroupsTaggingAPIClient{}
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
type ResourceGroupsTaggingAPIAPI interface {
	GetResources(*resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error)
	GetResourcesWithContext(aws.Context, *resourcegroupstaggingapi.GetResourcesInput, ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error)
	GetResourcesRequest(*resourcegroupstaggingapi.GetResourcesInput) (*request.Request, *resourcegroupstaggingapi.GetResourcesOutput)

	GetResourcesPages(*resourcegroupstaggingapi.GetResourcesInput, func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool) error
	GetResourcesPagesWithContext(aws.Context, *resourcegroupstaggingapi.GetResourcesInput, func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool, ...request.Option) error

	GetTagKeys(*resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
	GetTagKeysWithContext(aws.Context, *resourcegroupstaggingapi.GetTagKeysInput, ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
	GetTagKeysRequest(*resourcegroupstaggingapi.GetTagKeysInput) (*request.Request, *resourcegroupstaggingapi.GetTagKeysOutput)

	GetTagKeysPages(*resourcegroupstaggingapi.GetTagKeysInput, func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool) error
	GetTagKeysPagesWithContext(aws.Context, *resourcegroupstaggingapi.GetTagKeysInput, func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool, ...request.Option) error

	GetTagValues(*resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
	GetTagValuesWithContext(aws.Context, *resourcegroupstaggingapi.GetTagValuesInput, ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
	GetTagValuesRequest(*resourcegroupstaggingapi.GetTagValuesInput) (*request.Request, *resourcegroupstaggingapi.GetTagValuesOutput)

	GetTagValuesPages(*resourcegroupstaggingapi.GetTagValuesInput, func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool) error
	GetTagValuesPagesWithContext(aws.Context, *resourcegroupstaggingapi.GetTagValuesInput, func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool, ...request.Option) error

	TagResources(*resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error)
	TagResourcesWithContext(aws.Context, *resourcegroupstaggingapi.TagResourcesInput, ...request.Option) (*resourcegroupstaggingapi.TagResourcesOutput, error)
	TagResourcesRequest(*resourcegroupstaggingapi.TagResourcesInput) (*request.Request, *resourcegroupstaggingapi.TagResourcesOutput)

	UntagResources(*resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error)
	UntagResourcesWithContext(aws.Context, *resourcegroupstaggingapi.UntagResourcesInput, ...request.Option) (*resourcegroupstaggingapi.UntagResourcesOutput, error)
	UntagResourcesRequest(*resourcegroupstaggingapi.UntagResourcesInput) (*request.Request, *resourcegroupstaggingapi.UntagResourcesOutput)
}

var _ ResourceGroupsTaggingAPIAPI = (*resourcegroupstaggingapi.ResourceGroupsTaggingAPI)(nil)
