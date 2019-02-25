// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticacheiface provides an interface to enable mocking the Amazon ElastiCache service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticacheiface

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/service/elasticache"
)

// ElastiCacheAPI provides an interface to enable mocking the
// elasticache.ElastiCache service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon ElastiCache.
//    func myFunc(svc elasticacheiface.ElastiCacheAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := elasticache.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockElastiCacheClient struct {
//        elasticacheiface.ElastiCacheAPI
//    }
//    func (m *mockElastiCacheClient) AddTagsToResource(input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockElastiCacheClient{}
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
type ElastiCacheAPI interface {
	AddTagsToResource(*elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error)
	AddTagsToResourceWithContext(aws.Context, *elasticache.AddTagsToResourceInput, ...request.Option) (*elasticache.TagListMessage, error)
	AddTagsToResourceRequest(*elasticache.AddTagsToResourceInput) (*request.Request, *elasticache.TagListMessage)

	AuthorizeCacheSecurityGroupIngress(*elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error)
	AuthorizeCacheSecurityGroupIngressWithContext(aws.Context, *elasticache.AuthorizeCacheSecurityGroupIngressInput, ...request.Option) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error)
	AuthorizeCacheSecurityGroupIngressRequest(*elasticache.AuthorizeCacheSecurityGroupIngressInput) (*request.Request, *elasticache.AuthorizeCacheSecurityGroupIngressOutput)

	CopySnapshot(*elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error)
	CopySnapshotWithContext(aws.Context, *elasticache.CopySnapshotInput, ...request.Option) (*elasticache.CopySnapshotOutput, error)
	CopySnapshotRequest(*elasticache.CopySnapshotInput) (*request.Request, *elasticache.CopySnapshotOutput)

	CreateCacheCluster(*elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error)
	CreateCacheClusterWithContext(aws.Context, *elasticache.CreateCacheClusterInput, ...request.Option) (*elasticache.CreateCacheClusterOutput, error)
	CreateCacheClusterRequest(*elasticache.CreateCacheClusterInput) (*request.Request, *elasticache.CreateCacheClusterOutput)

	CreateCacheParameterGroup(*elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error)
	CreateCacheParameterGroupWithContext(aws.Context, *elasticache.CreateCacheParameterGroupInput, ...request.Option) (*elasticache.CreateCacheParameterGroupOutput, error)
	CreateCacheParameterGroupRequest(*elasticache.CreateCacheParameterGroupInput) (*request.Request, *elasticache.CreateCacheParameterGroupOutput)

	CreateCacheSecurityGroup(*elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error)
	CreateCacheSecurityGroupWithContext(aws.Context, *elasticache.CreateCacheSecurityGroupInput, ...request.Option) (*elasticache.CreateCacheSecurityGroupOutput, error)
	CreateCacheSecurityGroupRequest(*elasticache.CreateCacheSecurityGroupInput) (*request.Request, *elasticache.CreateCacheSecurityGroupOutput)

	CreateCacheSubnetGroup(*elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error)
	CreateCacheSubnetGroupWithContext(aws.Context, *elasticache.CreateCacheSubnetGroupInput, ...request.Option) (*elasticache.CreateCacheSubnetGroupOutput, error)
	CreateCacheSubnetGroupRequest(*elasticache.CreateCacheSubnetGroupInput) (*request.Request, *elasticache.CreateCacheSubnetGroupOutput)

	CreateReplicationGroup(*elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error)
	CreateReplicationGroupWithContext(aws.Context, *elasticache.CreateReplicationGroupInput, ...request.Option) (*elasticache.CreateReplicationGroupOutput, error)
	CreateReplicationGroupRequest(*elasticache.CreateReplicationGroupInput) (*request.Request, *elasticache.CreateReplicationGroupOutput)

	CreateSnapshot(*elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error)
	CreateSnapshotWithContext(aws.Context, *elasticache.CreateSnapshotInput, ...request.Option) (*elasticache.CreateSnapshotOutput, error)
	CreateSnapshotRequest(*elasticache.CreateSnapshotInput) (*request.Request, *elasticache.CreateSnapshotOutput)

	DecreaseReplicaCount(*elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error)
	DecreaseReplicaCountWithContext(aws.Context, *elasticache.DecreaseReplicaCountInput, ...request.Option) (*elasticache.DecreaseReplicaCountOutput, error)
	DecreaseReplicaCountRequest(*elasticache.DecreaseReplicaCountInput) (*request.Request, *elasticache.DecreaseReplicaCountOutput)

	DeleteCacheCluster(*elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error)
	DeleteCacheClusterWithContext(aws.Context, *elasticache.DeleteCacheClusterInput, ...request.Option) (*elasticache.DeleteCacheClusterOutput, error)
	DeleteCacheClusterRequest(*elasticache.DeleteCacheClusterInput) (*request.Request, *elasticache.DeleteCacheClusterOutput)

	DeleteCacheParameterGroup(*elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error)
	DeleteCacheParameterGroupWithContext(aws.Context, *elasticache.DeleteCacheParameterGroupInput, ...request.Option) (*elasticache.DeleteCacheParameterGroupOutput, error)
	DeleteCacheParameterGroupRequest(*elasticache.DeleteCacheParameterGroupInput) (*request.Request, *elasticache.DeleteCacheParameterGroupOutput)

	DeleteCacheSecurityGroup(*elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error)
	DeleteCacheSecurityGroupWithContext(aws.Context, *elasticache.DeleteCacheSecurityGroupInput, ...request.Option) (*elasticache.DeleteCacheSecurityGroupOutput, error)
	DeleteCacheSecurityGroupRequest(*elasticache.DeleteCacheSecurityGroupInput) (*request.Request, *elasticache.DeleteCacheSecurityGroupOutput)

	DeleteCacheSubnetGroup(*elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error)
	DeleteCacheSubnetGroupWithContext(aws.Context, *elasticache.DeleteCacheSubnetGroupInput, ...request.Option) (*elasticache.DeleteCacheSubnetGroupOutput, error)
	DeleteCacheSubnetGroupRequest(*elasticache.DeleteCacheSubnetGroupInput) (*request.Request, *elasticache.DeleteCacheSubnetGroupOutput)

	DeleteReplicationGroup(*elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error)
	DeleteReplicationGroupWithContext(aws.Context, *elasticache.DeleteReplicationGroupInput, ...request.Option) (*elasticache.DeleteReplicationGroupOutput, error)
	DeleteReplicationGroupRequest(*elasticache.DeleteReplicationGroupInput) (*request.Request, *elasticache.DeleteReplicationGroupOutput)

	DeleteSnapshot(*elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error)
	DeleteSnapshotWithContext(aws.Context, *elasticache.DeleteSnapshotInput, ...request.Option) (*elasticache.DeleteSnapshotOutput, error)
	DeleteSnapshotRequest(*elasticache.DeleteSnapshotInput) (*request.Request, *elasticache.DeleteSnapshotOutput)

	DescribeCacheClusters(*elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error)
	DescribeCacheClustersWithContext(aws.Context, *elasticache.DescribeCacheClustersInput, ...request.Option) (*elasticache.DescribeCacheClustersOutput, error)
	DescribeCacheClustersRequest(*elasticache.DescribeCacheClustersInput) (*request.Request, *elasticache.DescribeCacheClustersOutput)

	DescribeCacheClustersPages(*elasticache.DescribeCacheClustersInput, func(*elasticache.DescribeCacheClustersOutput, bool) bool) error
	DescribeCacheClustersPagesWithContext(aws.Context, *elasticache.DescribeCacheClustersInput, func(*elasticache.DescribeCacheClustersOutput, bool) bool, ...request.Option) error

	DescribeCacheEngineVersions(*elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error)
	DescribeCacheEngineVersionsWithContext(aws.Context, *elasticache.DescribeCacheEngineVersionsInput, ...request.Option) (*elasticache.DescribeCacheEngineVersionsOutput, error)
	DescribeCacheEngineVersionsRequest(*elasticache.DescribeCacheEngineVersionsInput) (*request.Request, *elasticache.DescribeCacheEngineVersionsOutput)

	DescribeCacheEngineVersionsPages(*elasticache.DescribeCacheEngineVersionsInput, func(*elasticache.DescribeCacheEngineVersionsOutput, bool) bool) error
	DescribeCacheEngineVersionsPagesWithContext(aws.Context, *elasticache.DescribeCacheEngineVersionsInput, func(*elasticache.DescribeCacheEngineVersionsOutput, bool) bool, ...request.Option) error

	DescribeCacheParameterGroups(*elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error)
	DescribeCacheParameterGroupsWithContext(aws.Context, *elasticache.DescribeCacheParameterGroupsInput, ...request.Option) (*elasticache.DescribeCacheParameterGroupsOutput, error)
	DescribeCacheParameterGroupsRequest(*elasticache.DescribeCacheParameterGroupsInput) (*request.Request, *elasticache.DescribeCacheParameterGroupsOutput)

	DescribeCacheParameterGroupsPages(*elasticache.DescribeCacheParameterGroupsInput, func(*elasticache.DescribeCacheParameterGroupsOutput, bool) bool) error
	DescribeCacheParameterGroupsPagesWithContext(aws.Context, *elasticache.DescribeCacheParameterGroupsInput, func(*elasticache.DescribeCacheParameterGroupsOutput, bool) bool, ...request.Option) error

	DescribeCacheParameters(*elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error)
	DescribeCacheParametersWithContext(aws.Context, *elasticache.DescribeCacheParametersInput, ...request.Option) (*elasticache.DescribeCacheParametersOutput, error)
	DescribeCacheParametersRequest(*elasticache.DescribeCacheParametersInput) (*request.Request, *elasticache.DescribeCacheParametersOutput)

	DescribeCacheParametersPages(*elasticache.DescribeCacheParametersInput, func(*elasticache.DescribeCacheParametersOutput, bool) bool) error
	DescribeCacheParametersPagesWithContext(aws.Context, *elasticache.DescribeCacheParametersInput, func(*elasticache.DescribeCacheParametersOutput, bool) bool, ...request.Option) error

	DescribeCacheSecurityGroups(*elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error)
	DescribeCacheSecurityGroupsWithContext(aws.Context, *elasticache.DescribeCacheSecurityGroupsInput, ...request.Option) (*elasticache.DescribeCacheSecurityGroupsOutput, error)
	DescribeCacheSecurityGroupsRequest(*elasticache.DescribeCacheSecurityGroupsInput) (*request.Request, *elasticache.DescribeCacheSecurityGroupsOutput)

	DescribeCacheSecurityGroupsPages(*elasticache.DescribeCacheSecurityGroupsInput, func(*elasticache.DescribeCacheSecurityGroupsOutput, bool) bool) error
	DescribeCacheSecurityGroupsPagesWithContext(aws.Context, *elasticache.DescribeCacheSecurityGroupsInput, func(*elasticache.DescribeCacheSecurityGroupsOutput, bool) bool, ...request.Option) error

	DescribeCacheSubnetGroups(*elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
	DescribeCacheSubnetGroupsWithContext(aws.Context, *elasticache.DescribeCacheSubnetGroupsInput, ...request.Option) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
	DescribeCacheSubnetGroupsRequest(*elasticache.DescribeCacheSubnetGroupsInput) (*request.Request, *elasticache.DescribeCacheSubnetGroupsOutput)

	DescribeCacheSubnetGroupsPages(*elasticache.DescribeCacheSubnetGroupsInput, func(*elasticache.DescribeCacheSubnetGroupsOutput, bool) bool) error
	DescribeCacheSubnetGroupsPagesWithContext(aws.Context, *elasticache.DescribeCacheSubnetGroupsInput, func(*elasticache.DescribeCacheSubnetGroupsOutput, bool) bool, ...request.Option) error

	DescribeEngineDefaultParameters(*elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error)
	DescribeEngineDefaultParametersWithContext(aws.Context, *elasticache.DescribeEngineDefaultParametersInput, ...request.Option) (*elasticache.DescribeEngineDefaultParametersOutput, error)
	DescribeEngineDefaultParametersRequest(*elasticache.DescribeEngineDefaultParametersInput) (*request.Request, *elasticache.DescribeEngineDefaultParametersOutput)

	DescribeEngineDefaultParametersPages(*elasticache.DescribeEngineDefaultParametersInput, func(*elasticache.DescribeEngineDefaultParametersOutput, bool) bool) error
	DescribeEngineDefaultParametersPagesWithContext(aws.Context, *elasticache.DescribeEngineDefaultParametersInput, func(*elasticache.DescribeEngineDefaultParametersOutput, bool) bool, ...request.Option) error

	DescribeEvents(*elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error)
	DescribeEventsWithContext(aws.Context, *elasticache.DescribeEventsInput, ...request.Option) (*elasticache.DescribeEventsOutput, error)
	DescribeEventsRequest(*elasticache.DescribeEventsInput) (*request.Request, *elasticache.DescribeEventsOutput)

	DescribeEventsPages(*elasticache.DescribeEventsInput, func(*elasticache.DescribeEventsOutput, bool) bool) error
	DescribeEventsPagesWithContext(aws.Context, *elasticache.DescribeEventsInput, func(*elasticache.DescribeEventsOutput, bool) bool, ...request.Option) error

	DescribeReplicationGroups(*elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error)
	DescribeReplicationGroupsWithContext(aws.Context, *elasticache.DescribeReplicationGroupsInput, ...request.Option) (*elasticache.DescribeReplicationGroupsOutput, error)
	DescribeReplicationGroupsRequest(*elasticache.DescribeReplicationGroupsInput) (*request.Request, *elasticache.DescribeReplicationGroupsOutput)

	DescribeReplicationGroupsPages(*elasticache.DescribeReplicationGroupsInput, func(*elasticache.DescribeReplicationGroupsOutput, bool) bool) error
	DescribeReplicationGroupsPagesWithContext(aws.Context, *elasticache.DescribeReplicationGroupsInput, func(*elasticache.DescribeReplicationGroupsOutput, bool) bool, ...request.Option) error

	DescribeReservedCacheNodes(*elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error)
	DescribeReservedCacheNodesWithContext(aws.Context, *elasticache.DescribeReservedCacheNodesInput, ...request.Option) (*elasticache.DescribeReservedCacheNodesOutput, error)
	DescribeReservedCacheNodesRequest(*elasticache.DescribeReservedCacheNodesInput) (*request.Request, *elasticache.DescribeReservedCacheNodesOutput)

	DescribeReservedCacheNodesPages(*elasticache.DescribeReservedCacheNodesInput, func(*elasticache.DescribeReservedCacheNodesOutput, bool) bool) error
	DescribeReservedCacheNodesPagesWithContext(aws.Context, *elasticache.DescribeReservedCacheNodesInput, func(*elasticache.DescribeReservedCacheNodesOutput, bool) bool, ...request.Option) error

	DescribeReservedCacheNodesOfferings(*elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
	DescribeReservedCacheNodesOfferingsWithContext(aws.Context, *elasticache.DescribeReservedCacheNodesOfferingsInput, ...request.Option) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
	DescribeReservedCacheNodesOfferingsRequest(*elasticache.DescribeReservedCacheNodesOfferingsInput) (*request.Request, *elasticache.DescribeReservedCacheNodesOfferingsOutput)

	DescribeReservedCacheNodesOfferingsPages(*elasticache.DescribeReservedCacheNodesOfferingsInput, func(*elasticache.DescribeReservedCacheNodesOfferingsOutput, bool) bool) error
	DescribeReservedCacheNodesOfferingsPagesWithContext(aws.Context, *elasticache.DescribeReservedCacheNodesOfferingsInput, func(*elasticache.DescribeReservedCacheNodesOfferingsOutput, bool) bool, ...request.Option) error

	DescribeSnapshots(*elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error)
	DescribeSnapshotsWithContext(aws.Context, *elasticache.DescribeSnapshotsInput, ...request.Option) (*elasticache.DescribeSnapshotsOutput, error)
	DescribeSnapshotsRequest(*elasticache.DescribeSnapshotsInput) (*request.Request, *elasticache.DescribeSnapshotsOutput)

	DescribeSnapshotsPages(*elasticache.DescribeSnapshotsInput, func(*elasticache.DescribeSnapshotsOutput, bool) bool) error
	DescribeSnapshotsPagesWithContext(aws.Context, *elasticache.DescribeSnapshotsInput, func(*elasticache.DescribeSnapshotsOutput, bool) bool, ...request.Option) error

	IncreaseReplicaCount(*elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error)
	IncreaseReplicaCountWithContext(aws.Context, *elasticache.IncreaseReplicaCountInput, ...request.Option) (*elasticache.IncreaseReplicaCountOutput, error)
	IncreaseReplicaCountRequest(*elasticache.IncreaseReplicaCountInput) (*request.Request, *elasticache.IncreaseReplicaCountOutput)

	ListAllowedNodeTypeModifications(*elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error)
	ListAllowedNodeTypeModificationsWithContext(aws.Context, *elasticache.ListAllowedNodeTypeModificationsInput, ...request.Option) (*elasticache.ListAllowedNodeTypeModificationsOutput, error)
	ListAllowedNodeTypeModificationsRequest(*elasticache.ListAllowedNodeTypeModificationsInput) (*request.Request, *elasticache.ListAllowedNodeTypeModificationsOutput)

	ListTagsForResource(*elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error)
	ListTagsForResourceWithContext(aws.Context, *elasticache.ListTagsForResourceInput, ...request.Option) (*elasticache.TagListMessage, error)
	ListTagsForResourceRequest(*elasticache.ListTagsForResourceInput) (*request.Request, *elasticache.TagListMessage)

	ModifyCacheCluster(*elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error)
	ModifyCacheClusterWithContext(aws.Context, *elasticache.ModifyCacheClusterInput, ...request.Option) (*elasticache.ModifyCacheClusterOutput, error)
	ModifyCacheClusterRequest(*elasticache.ModifyCacheClusterInput) (*request.Request, *elasticache.ModifyCacheClusterOutput)

	ModifyCacheParameterGroup(*elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)
	ModifyCacheParameterGroupWithContext(aws.Context, *elasticache.ModifyCacheParameterGroupInput, ...request.Option) (*elasticache.CacheParameterGroupNameMessage, error)
	ModifyCacheParameterGroupRequest(*elasticache.ModifyCacheParameterGroupInput) (*request.Request, *elasticache.CacheParameterGroupNameMessage)

	ModifyCacheSubnetGroup(*elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error)
	ModifyCacheSubnetGroupWithContext(aws.Context, *elasticache.ModifyCacheSubnetGroupInput, ...request.Option) (*elasticache.ModifyCacheSubnetGroupOutput, error)
	ModifyCacheSubnetGroupRequest(*elasticache.ModifyCacheSubnetGroupInput) (*request.Request, *elasticache.ModifyCacheSubnetGroupOutput)

	ModifyReplicationGroup(*elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error)
	ModifyReplicationGroupWithContext(aws.Context, *elasticache.ModifyReplicationGroupInput, ...request.Option) (*elasticache.ModifyReplicationGroupOutput, error)
	ModifyReplicationGroupRequest(*elasticache.ModifyReplicationGroupInput) (*request.Request, *elasticache.ModifyReplicationGroupOutput)

	ModifyReplicationGroupShardConfiguration(*elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error)
	ModifyReplicationGroupShardConfigurationWithContext(aws.Context, *elasticache.ModifyReplicationGroupShardConfigurationInput, ...request.Option) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error)
	ModifyReplicationGroupShardConfigurationRequest(*elasticache.ModifyReplicationGroupShardConfigurationInput) (*request.Request, *elasticache.ModifyReplicationGroupShardConfigurationOutput)

	PurchaseReservedCacheNodesOffering(*elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error)
	PurchaseReservedCacheNodesOfferingWithContext(aws.Context, *elasticache.PurchaseReservedCacheNodesOfferingInput, ...request.Option) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error)
	PurchaseReservedCacheNodesOfferingRequest(*elasticache.PurchaseReservedCacheNodesOfferingInput) (*request.Request, *elasticache.PurchaseReservedCacheNodesOfferingOutput)

	RebootCacheCluster(*elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error)
	RebootCacheClusterWithContext(aws.Context, *elasticache.RebootCacheClusterInput, ...request.Option) (*elasticache.RebootCacheClusterOutput, error)
	RebootCacheClusterRequest(*elasticache.RebootCacheClusterInput) (*request.Request, *elasticache.RebootCacheClusterOutput)

	RemoveTagsFromResource(*elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error)
	RemoveTagsFromResourceWithContext(aws.Context, *elasticache.RemoveTagsFromResourceInput, ...request.Option) (*elasticache.TagListMessage, error)
	RemoveTagsFromResourceRequest(*elasticache.RemoveTagsFromResourceInput) (*request.Request, *elasticache.TagListMessage)

	ResetCacheParameterGroup(*elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)
	ResetCacheParameterGroupWithContext(aws.Context, *elasticache.ResetCacheParameterGroupInput, ...request.Option) (*elasticache.CacheParameterGroupNameMessage, error)
	ResetCacheParameterGroupRequest(*elasticache.ResetCacheParameterGroupInput) (*request.Request, *elasticache.CacheParameterGroupNameMessage)

	RevokeCacheSecurityGroupIngress(*elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error)
	RevokeCacheSecurityGroupIngressWithContext(aws.Context, *elasticache.RevokeCacheSecurityGroupIngressInput, ...request.Option) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error)
	RevokeCacheSecurityGroupIngressRequest(*elasticache.RevokeCacheSecurityGroupIngressInput) (*request.Request, *elasticache.RevokeCacheSecurityGroupIngressOutput)

	TestFailover(*elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error)
	TestFailoverWithContext(aws.Context, *elasticache.TestFailoverInput, ...request.Option) (*elasticache.TestFailoverOutput, error)
	TestFailoverRequest(*elasticache.TestFailoverInput) (*request.Request, *elasticache.TestFailoverOutput)

	WaitUntilCacheClusterAvailable(*elasticache.DescribeCacheClustersInput) error
	WaitUntilCacheClusterAvailableWithContext(aws.Context, *elasticache.DescribeCacheClustersInput, ...request.WaiterOption) error

	WaitUntilCacheClusterDeleted(*elasticache.DescribeCacheClustersInput) error
	WaitUntilCacheClusterDeletedWithContext(aws.Context, *elasticache.DescribeCacheClustersInput, ...request.WaiterOption) error

	WaitUntilReplicationGroupAvailable(*elasticache.DescribeReplicationGroupsInput) error
	WaitUntilReplicationGroupAvailableWithContext(aws.Context, *elasticache.DescribeReplicationGroupsInput, ...request.WaiterOption) error

	WaitUntilReplicationGroupDeleted(*elasticache.DescribeReplicationGroupsInput) error
	WaitUntilReplicationGroupDeletedWithContext(aws.Context, *elasticache.DescribeReplicationGroupsInput, ...request.WaiterOption) error
}

var _ ElastiCacheAPI = (*elasticache.ElastiCache)(nil)
