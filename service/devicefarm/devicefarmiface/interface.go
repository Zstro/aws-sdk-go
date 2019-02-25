// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package devicefarmiface provides an interface to enable mocking the AWS Device Farm service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package devicefarmiface

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/service/devicefarm"
)

// DeviceFarmAPI provides an interface to enable mocking the
// devicefarm.DeviceFarm service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Device Farm.
//    func myFunc(svc devicefarmiface.DeviceFarmAPI) bool {
//        // Make svc.CreateDevicePool request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := devicefarm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDeviceFarmClient struct {
//        devicefarmiface.DeviceFarmAPI
//    }
//    func (m *mockDeviceFarmClient) CreateDevicePool(input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDeviceFarmClient{}
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
type DeviceFarmAPI interface {
	CreateDevicePool(*devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error)
	CreateDevicePoolWithContext(aws.Context, *devicefarm.CreateDevicePoolInput, ...request.Option) (*devicefarm.CreateDevicePoolOutput, error)
	CreateDevicePoolRequest(*devicefarm.CreateDevicePoolInput) (*request.Request, *devicefarm.CreateDevicePoolOutput)

	CreateInstanceProfile(*devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error)
	CreateInstanceProfileWithContext(aws.Context, *devicefarm.CreateInstanceProfileInput, ...request.Option) (*devicefarm.CreateInstanceProfileOutput, error)
	CreateInstanceProfileRequest(*devicefarm.CreateInstanceProfileInput) (*request.Request, *devicefarm.CreateInstanceProfileOutput)

	CreateNetworkProfile(*devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error)
	CreateNetworkProfileWithContext(aws.Context, *devicefarm.CreateNetworkProfileInput, ...request.Option) (*devicefarm.CreateNetworkProfileOutput, error)
	CreateNetworkProfileRequest(*devicefarm.CreateNetworkProfileInput) (*request.Request, *devicefarm.CreateNetworkProfileOutput)

	CreateProject(*devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *devicefarm.CreateProjectInput, ...request.Option) (*devicefarm.CreateProjectOutput, error)
	CreateProjectRequest(*devicefarm.CreateProjectInput) (*request.Request, *devicefarm.CreateProjectOutput)

	CreateRemoteAccessSession(*devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error)
	CreateRemoteAccessSessionWithContext(aws.Context, *devicefarm.CreateRemoteAccessSessionInput, ...request.Option) (*devicefarm.CreateRemoteAccessSessionOutput, error)
	CreateRemoteAccessSessionRequest(*devicefarm.CreateRemoteAccessSessionInput) (*request.Request, *devicefarm.CreateRemoteAccessSessionOutput)

	CreateUpload(*devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error)
	CreateUploadWithContext(aws.Context, *devicefarm.CreateUploadInput, ...request.Option) (*devicefarm.CreateUploadOutput, error)
	CreateUploadRequest(*devicefarm.CreateUploadInput) (*request.Request, *devicefarm.CreateUploadOutput)

	CreateVPCEConfiguration(*devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error)
	CreateVPCEConfigurationWithContext(aws.Context, *devicefarm.CreateVPCEConfigurationInput, ...request.Option) (*devicefarm.CreateVPCEConfigurationOutput, error)
	CreateVPCEConfigurationRequest(*devicefarm.CreateVPCEConfigurationInput) (*request.Request, *devicefarm.CreateVPCEConfigurationOutput)

	DeleteDevicePool(*devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error)
	DeleteDevicePoolWithContext(aws.Context, *devicefarm.DeleteDevicePoolInput, ...request.Option) (*devicefarm.DeleteDevicePoolOutput, error)
	DeleteDevicePoolRequest(*devicefarm.DeleteDevicePoolInput) (*request.Request, *devicefarm.DeleteDevicePoolOutput)

	DeleteInstanceProfile(*devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileWithContext(aws.Context, *devicefarm.DeleteInstanceProfileInput, ...request.Option) (*devicefarm.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileRequest(*devicefarm.DeleteInstanceProfileInput) (*request.Request, *devicefarm.DeleteInstanceProfileOutput)

	DeleteNetworkProfile(*devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error)
	DeleteNetworkProfileWithContext(aws.Context, *devicefarm.DeleteNetworkProfileInput, ...request.Option) (*devicefarm.DeleteNetworkProfileOutput, error)
	DeleteNetworkProfileRequest(*devicefarm.DeleteNetworkProfileInput) (*request.Request, *devicefarm.DeleteNetworkProfileOutput)

	DeleteProject(*devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *devicefarm.DeleteProjectInput, ...request.Option) (*devicefarm.DeleteProjectOutput, error)
	DeleteProjectRequest(*devicefarm.DeleteProjectInput) (*request.Request, *devicefarm.DeleteProjectOutput)

	DeleteRemoteAccessSession(*devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error)
	DeleteRemoteAccessSessionWithContext(aws.Context, *devicefarm.DeleteRemoteAccessSessionInput, ...request.Option) (*devicefarm.DeleteRemoteAccessSessionOutput, error)
	DeleteRemoteAccessSessionRequest(*devicefarm.DeleteRemoteAccessSessionInput) (*request.Request, *devicefarm.DeleteRemoteAccessSessionOutput)

	DeleteRun(*devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error)
	DeleteRunWithContext(aws.Context, *devicefarm.DeleteRunInput, ...request.Option) (*devicefarm.DeleteRunOutput, error)
	DeleteRunRequest(*devicefarm.DeleteRunInput) (*request.Request, *devicefarm.DeleteRunOutput)

	DeleteUpload(*devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error)
	DeleteUploadWithContext(aws.Context, *devicefarm.DeleteUploadInput, ...request.Option) (*devicefarm.DeleteUploadOutput, error)
	DeleteUploadRequest(*devicefarm.DeleteUploadInput) (*request.Request, *devicefarm.DeleteUploadOutput)

	DeleteVPCEConfiguration(*devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error)
	DeleteVPCEConfigurationWithContext(aws.Context, *devicefarm.DeleteVPCEConfigurationInput, ...request.Option) (*devicefarm.DeleteVPCEConfigurationOutput, error)
	DeleteVPCEConfigurationRequest(*devicefarm.DeleteVPCEConfigurationInput) (*request.Request, *devicefarm.DeleteVPCEConfigurationOutput)

	GetAccountSettings(*devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error)
	GetAccountSettingsWithContext(aws.Context, *devicefarm.GetAccountSettingsInput, ...request.Option) (*devicefarm.GetAccountSettingsOutput, error)
	GetAccountSettingsRequest(*devicefarm.GetAccountSettingsInput) (*request.Request, *devicefarm.GetAccountSettingsOutput)

	GetDevice(*devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error)
	GetDeviceWithContext(aws.Context, *devicefarm.GetDeviceInput, ...request.Option) (*devicefarm.GetDeviceOutput, error)
	GetDeviceRequest(*devicefarm.GetDeviceInput) (*request.Request, *devicefarm.GetDeviceOutput)

	GetDeviceInstance(*devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error)
	GetDeviceInstanceWithContext(aws.Context, *devicefarm.GetDeviceInstanceInput, ...request.Option) (*devicefarm.GetDeviceInstanceOutput, error)
	GetDeviceInstanceRequest(*devicefarm.GetDeviceInstanceInput) (*request.Request, *devicefarm.GetDeviceInstanceOutput)

	GetDevicePool(*devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error)
	GetDevicePoolWithContext(aws.Context, *devicefarm.GetDevicePoolInput, ...request.Option) (*devicefarm.GetDevicePoolOutput, error)
	GetDevicePoolRequest(*devicefarm.GetDevicePoolInput) (*request.Request, *devicefarm.GetDevicePoolOutput)

	GetDevicePoolCompatibility(*devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error)
	GetDevicePoolCompatibilityWithContext(aws.Context, *devicefarm.GetDevicePoolCompatibilityInput, ...request.Option) (*devicefarm.GetDevicePoolCompatibilityOutput, error)
	GetDevicePoolCompatibilityRequest(*devicefarm.GetDevicePoolCompatibilityInput) (*request.Request, *devicefarm.GetDevicePoolCompatibilityOutput)

	GetInstanceProfile(*devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error)
	GetInstanceProfileWithContext(aws.Context, *devicefarm.GetInstanceProfileInput, ...request.Option) (*devicefarm.GetInstanceProfileOutput, error)
	GetInstanceProfileRequest(*devicefarm.GetInstanceProfileInput) (*request.Request, *devicefarm.GetInstanceProfileOutput)

	GetJob(*devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error)
	GetJobWithContext(aws.Context, *devicefarm.GetJobInput, ...request.Option) (*devicefarm.GetJobOutput, error)
	GetJobRequest(*devicefarm.GetJobInput) (*request.Request, *devicefarm.GetJobOutput)

	GetNetworkProfile(*devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error)
	GetNetworkProfileWithContext(aws.Context, *devicefarm.GetNetworkProfileInput, ...request.Option) (*devicefarm.GetNetworkProfileOutput, error)
	GetNetworkProfileRequest(*devicefarm.GetNetworkProfileInput) (*request.Request, *devicefarm.GetNetworkProfileOutput)

	GetOfferingStatus(*devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error)
	GetOfferingStatusWithContext(aws.Context, *devicefarm.GetOfferingStatusInput, ...request.Option) (*devicefarm.GetOfferingStatusOutput, error)
	GetOfferingStatusRequest(*devicefarm.GetOfferingStatusInput) (*request.Request, *devicefarm.GetOfferingStatusOutput)

	GetOfferingStatusPages(*devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool) error
	GetOfferingStatusPagesWithContext(aws.Context, *devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool, ...request.Option) error

	GetProject(*devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error)
	GetProjectWithContext(aws.Context, *devicefarm.GetProjectInput, ...request.Option) (*devicefarm.GetProjectOutput, error)
	GetProjectRequest(*devicefarm.GetProjectInput) (*request.Request, *devicefarm.GetProjectOutput)

	GetRemoteAccessSession(*devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error)
	GetRemoteAccessSessionWithContext(aws.Context, *devicefarm.GetRemoteAccessSessionInput, ...request.Option) (*devicefarm.GetRemoteAccessSessionOutput, error)
	GetRemoteAccessSessionRequest(*devicefarm.GetRemoteAccessSessionInput) (*request.Request, *devicefarm.GetRemoteAccessSessionOutput)

	GetRun(*devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error)
	GetRunWithContext(aws.Context, *devicefarm.GetRunInput, ...request.Option) (*devicefarm.GetRunOutput, error)
	GetRunRequest(*devicefarm.GetRunInput) (*request.Request, *devicefarm.GetRunOutput)

	GetSuite(*devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error)
	GetSuiteWithContext(aws.Context, *devicefarm.GetSuiteInput, ...request.Option) (*devicefarm.GetSuiteOutput, error)
	GetSuiteRequest(*devicefarm.GetSuiteInput) (*request.Request, *devicefarm.GetSuiteOutput)

	GetTest(*devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error)
	GetTestWithContext(aws.Context, *devicefarm.GetTestInput, ...request.Option) (*devicefarm.GetTestOutput, error)
	GetTestRequest(*devicefarm.GetTestInput) (*request.Request, *devicefarm.GetTestOutput)

	GetUpload(*devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error)
	GetUploadWithContext(aws.Context, *devicefarm.GetUploadInput, ...request.Option) (*devicefarm.GetUploadOutput, error)
	GetUploadRequest(*devicefarm.GetUploadInput) (*request.Request, *devicefarm.GetUploadOutput)

	GetVPCEConfiguration(*devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error)
	GetVPCEConfigurationWithContext(aws.Context, *devicefarm.GetVPCEConfigurationInput, ...request.Option) (*devicefarm.GetVPCEConfigurationOutput, error)
	GetVPCEConfigurationRequest(*devicefarm.GetVPCEConfigurationInput) (*request.Request, *devicefarm.GetVPCEConfigurationOutput)

	InstallToRemoteAccessSession(*devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error)
	InstallToRemoteAccessSessionWithContext(aws.Context, *devicefarm.InstallToRemoteAccessSessionInput, ...request.Option) (*devicefarm.InstallToRemoteAccessSessionOutput, error)
	InstallToRemoteAccessSessionRequest(*devicefarm.InstallToRemoteAccessSessionInput) (*request.Request, *devicefarm.InstallToRemoteAccessSessionOutput)

	ListArtifacts(*devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error)
	ListArtifactsWithContext(aws.Context, *devicefarm.ListArtifactsInput, ...request.Option) (*devicefarm.ListArtifactsOutput, error)
	ListArtifactsRequest(*devicefarm.ListArtifactsInput) (*request.Request, *devicefarm.ListArtifactsOutput)

	ListArtifactsPages(*devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool) error
	ListArtifactsPagesWithContext(aws.Context, *devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool, ...request.Option) error

	ListDeviceInstances(*devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error)
	ListDeviceInstancesWithContext(aws.Context, *devicefarm.ListDeviceInstancesInput, ...request.Option) (*devicefarm.ListDeviceInstancesOutput, error)
	ListDeviceInstancesRequest(*devicefarm.ListDeviceInstancesInput) (*request.Request, *devicefarm.ListDeviceInstancesOutput)

	ListDevicePools(*devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error)
	ListDevicePoolsWithContext(aws.Context, *devicefarm.ListDevicePoolsInput, ...request.Option) (*devicefarm.ListDevicePoolsOutput, error)
	ListDevicePoolsRequest(*devicefarm.ListDevicePoolsInput) (*request.Request, *devicefarm.ListDevicePoolsOutput)

	ListDevicePoolsPages(*devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool) error
	ListDevicePoolsPagesWithContext(aws.Context, *devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool, ...request.Option) error

	ListDevices(*devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error)
	ListDevicesWithContext(aws.Context, *devicefarm.ListDevicesInput, ...request.Option) (*devicefarm.ListDevicesOutput, error)
	ListDevicesRequest(*devicefarm.ListDevicesInput) (*request.Request, *devicefarm.ListDevicesOutput)

	ListDevicesPages(*devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool) error
	ListDevicesPagesWithContext(aws.Context, *devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool, ...request.Option) error

	ListInstanceProfiles(*devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error)
	ListInstanceProfilesWithContext(aws.Context, *devicefarm.ListInstanceProfilesInput, ...request.Option) (*devicefarm.ListInstanceProfilesOutput, error)
	ListInstanceProfilesRequest(*devicefarm.ListInstanceProfilesInput) (*request.Request, *devicefarm.ListInstanceProfilesOutput)

	ListJobs(*devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *devicefarm.ListJobsInput, ...request.Option) (*devicefarm.ListJobsOutput, error)
	ListJobsRequest(*devicefarm.ListJobsInput) (*request.Request, *devicefarm.ListJobsOutput)

	ListJobsPages(*devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool, ...request.Option) error

	ListNetworkProfiles(*devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error)
	ListNetworkProfilesWithContext(aws.Context, *devicefarm.ListNetworkProfilesInput, ...request.Option) (*devicefarm.ListNetworkProfilesOutput, error)
	ListNetworkProfilesRequest(*devicefarm.ListNetworkProfilesInput) (*request.Request, *devicefarm.ListNetworkProfilesOutput)

	ListOfferingPromotions(*devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error)
	ListOfferingPromotionsWithContext(aws.Context, *devicefarm.ListOfferingPromotionsInput, ...request.Option) (*devicefarm.ListOfferingPromotionsOutput, error)
	ListOfferingPromotionsRequest(*devicefarm.ListOfferingPromotionsInput) (*request.Request, *devicefarm.ListOfferingPromotionsOutput)

	ListOfferingTransactions(*devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error)
	ListOfferingTransactionsWithContext(aws.Context, *devicefarm.ListOfferingTransactionsInput, ...request.Option) (*devicefarm.ListOfferingTransactionsOutput, error)
	ListOfferingTransactionsRequest(*devicefarm.ListOfferingTransactionsInput) (*request.Request, *devicefarm.ListOfferingTransactionsOutput)

	ListOfferingTransactionsPages(*devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool) error
	ListOfferingTransactionsPagesWithContext(aws.Context, *devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool, ...request.Option) error

	ListOfferings(*devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error)
	ListOfferingsWithContext(aws.Context, *devicefarm.ListOfferingsInput, ...request.Option) (*devicefarm.ListOfferingsOutput, error)
	ListOfferingsRequest(*devicefarm.ListOfferingsInput) (*request.Request, *devicefarm.ListOfferingsOutput)

	ListOfferingsPages(*devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool) error
	ListOfferingsPagesWithContext(aws.Context, *devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool, ...request.Option) error

	ListProjects(*devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *devicefarm.ListProjectsInput, ...request.Option) (*devicefarm.ListProjectsOutput, error)
	ListProjectsRequest(*devicefarm.ListProjectsInput) (*request.Request, *devicefarm.ListProjectsOutput)

	ListProjectsPages(*devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool) error
	ListProjectsPagesWithContext(aws.Context, *devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool, ...request.Option) error

	ListRemoteAccessSessions(*devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error)
	ListRemoteAccessSessionsWithContext(aws.Context, *devicefarm.ListRemoteAccessSessionsInput, ...request.Option) (*devicefarm.ListRemoteAccessSessionsOutput, error)
	ListRemoteAccessSessionsRequest(*devicefarm.ListRemoteAccessSessionsInput) (*request.Request, *devicefarm.ListRemoteAccessSessionsOutput)

	ListRuns(*devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error)
	ListRunsWithContext(aws.Context, *devicefarm.ListRunsInput, ...request.Option) (*devicefarm.ListRunsOutput, error)
	ListRunsRequest(*devicefarm.ListRunsInput) (*request.Request, *devicefarm.ListRunsOutput)

	ListRunsPages(*devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool) error
	ListRunsPagesWithContext(aws.Context, *devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool, ...request.Option) error

	ListSamples(*devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error)
	ListSamplesWithContext(aws.Context, *devicefarm.ListSamplesInput, ...request.Option) (*devicefarm.ListSamplesOutput, error)
	ListSamplesRequest(*devicefarm.ListSamplesInput) (*request.Request, *devicefarm.ListSamplesOutput)

	ListSamplesPages(*devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool) error
	ListSamplesPagesWithContext(aws.Context, *devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool, ...request.Option) error

	ListSuites(*devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error)
	ListSuitesWithContext(aws.Context, *devicefarm.ListSuitesInput, ...request.Option) (*devicefarm.ListSuitesOutput, error)
	ListSuitesRequest(*devicefarm.ListSuitesInput) (*request.Request, *devicefarm.ListSuitesOutput)

	ListSuitesPages(*devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool) error
	ListSuitesPagesWithContext(aws.Context, *devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool, ...request.Option) error

	ListTests(*devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error)
	ListTestsWithContext(aws.Context, *devicefarm.ListTestsInput, ...request.Option) (*devicefarm.ListTestsOutput, error)
	ListTestsRequest(*devicefarm.ListTestsInput) (*request.Request, *devicefarm.ListTestsOutput)

	ListTestsPages(*devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool) error
	ListTestsPagesWithContext(aws.Context, *devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool, ...request.Option) error

	ListUniqueProblems(*devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error)
	ListUniqueProblemsWithContext(aws.Context, *devicefarm.ListUniqueProblemsInput, ...request.Option) (*devicefarm.ListUniqueProblemsOutput, error)
	ListUniqueProblemsRequest(*devicefarm.ListUniqueProblemsInput) (*request.Request, *devicefarm.ListUniqueProblemsOutput)

	ListUniqueProblemsPages(*devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error
	ListUniqueProblemsPagesWithContext(aws.Context, *devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool, ...request.Option) error

	ListUploads(*devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error)
	ListUploadsWithContext(aws.Context, *devicefarm.ListUploadsInput, ...request.Option) (*devicefarm.ListUploadsOutput, error)
	ListUploadsRequest(*devicefarm.ListUploadsInput) (*request.Request, *devicefarm.ListUploadsOutput)

	ListUploadsPages(*devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool) error
	ListUploadsPagesWithContext(aws.Context, *devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool, ...request.Option) error

	ListVPCEConfigurations(*devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error)
	ListVPCEConfigurationsWithContext(aws.Context, *devicefarm.ListVPCEConfigurationsInput, ...request.Option) (*devicefarm.ListVPCEConfigurationsOutput, error)
	ListVPCEConfigurationsRequest(*devicefarm.ListVPCEConfigurationsInput) (*request.Request, *devicefarm.ListVPCEConfigurationsOutput)

	PurchaseOffering(*devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error)
	PurchaseOfferingWithContext(aws.Context, *devicefarm.PurchaseOfferingInput, ...request.Option) (*devicefarm.PurchaseOfferingOutput, error)
	PurchaseOfferingRequest(*devicefarm.PurchaseOfferingInput) (*request.Request, *devicefarm.PurchaseOfferingOutput)

	RenewOffering(*devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error)
	RenewOfferingWithContext(aws.Context, *devicefarm.RenewOfferingInput, ...request.Option) (*devicefarm.RenewOfferingOutput, error)
	RenewOfferingRequest(*devicefarm.RenewOfferingInput) (*request.Request, *devicefarm.RenewOfferingOutput)

	ScheduleRun(*devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error)
	ScheduleRunWithContext(aws.Context, *devicefarm.ScheduleRunInput, ...request.Option) (*devicefarm.ScheduleRunOutput, error)
	ScheduleRunRequest(*devicefarm.ScheduleRunInput) (*request.Request, *devicefarm.ScheduleRunOutput)

	StopJob(*devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error)
	StopJobWithContext(aws.Context, *devicefarm.StopJobInput, ...request.Option) (*devicefarm.StopJobOutput, error)
	StopJobRequest(*devicefarm.StopJobInput) (*request.Request, *devicefarm.StopJobOutput)

	StopRemoteAccessSession(*devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error)
	StopRemoteAccessSessionWithContext(aws.Context, *devicefarm.StopRemoteAccessSessionInput, ...request.Option) (*devicefarm.StopRemoteAccessSessionOutput, error)
	StopRemoteAccessSessionRequest(*devicefarm.StopRemoteAccessSessionInput) (*request.Request, *devicefarm.StopRemoteAccessSessionOutput)

	StopRun(*devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error)
	StopRunWithContext(aws.Context, *devicefarm.StopRunInput, ...request.Option) (*devicefarm.StopRunOutput, error)
	StopRunRequest(*devicefarm.StopRunInput) (*request.Request, *devicefarm.StopRunOutput)

	UpdateDeviceInstance(*devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error)
	UpdateDeviceInstanceWithContext(aws.Context, *devicefarm.UpdateDeviceInstanceInput, ...request.Option) (*devicefarm.UpdateDeviceInstanceOutput, error)
	UpdateDeviceInstanceRequest(*devicefarm.UpdateDeviceInstanceInput) (*request.Request, *devicefarm.UpdateDeviceInstanceOutput)

	UpdateDevicePool(*devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error)
	UpdateDevicePoolWithContext(aws.Context, *devicefarm.UpdateDevicePoolInput, ...request.Option) (*devicefarm.UpdateDevicePoolOutput, error)
	UpdateDevicePoolRequest(*devicefarm.UpdateDevicePoolInput) (*request.Request, *devicefarm.UpdateDevicePoolOutput)

	UpdateInstanceProfile(*devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error)
	UpdateInstanceProfileWithContext(aws.Context, *devicefarm.UpdateInstanceProfileInput, ...request.Option) (*devicefarm.UpdateInstanceProfileOutput, error)
	UpdateInstanceProfileRequest(*devicefarm.UpdateInstanceProfileInput) (*request.Request, *devicefarm.UpdateInstanceProfileOutput)

	UpdateNetworkProfile(*devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error)
	UpdateNetworkProfileWithContext(aws.Context, *devicefarm.UpdateNetworkProfileInput, ...request.Option) (*devicefarm.UpdateNetworkProfileOutput, error)
	UpdateNetworkProfileRequest(*devicefarm.UpdateNetworkProfileInput) (*request.Request, *devicefarm.UpdateNetworkProfileOutput)

	UpdateProject(*devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *devicefarm.UpdateProjectInput, ...request.Option) (*devicefarm.UpdateProjectOutput, error)
	UpdateProjectRequest(*devicefarm.UpdateProjectInput) (*request.Request, *devicefarm.UpdateProjectOutput)

	UpdateUpload(*devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error)
	UpdateUploadWithContext(aws.Context, *devicefarm.UpdateUploadInput, ...request.Option) (*devicefarm.UpdateUploadOutput, error)
	UpdateUploadRequest(*devicefarm.UpdateUploadInput) (*request.Request, *devicefarm.UpdateUploadOutput)

	UpdateVPCEConfiguration(*devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error)
	UpdateVPCEConfigurationWithContext(aws.Context, *devicefarm.UpdateVPCEConfigurationInput, ...request.Option) (*devicefarm.UpdateVPCEConfigurationOutput, error)
	UpdateVPCEConfigurationRequest(*devicefarm.UpdateVPCEConfigurationInput) (*request.Request, *devicefarm.UpdateVPCEConfigurationOutput)
}

var _ DeviceFarmAPI = (*devicefarm.DeviceFarm)(nil)
