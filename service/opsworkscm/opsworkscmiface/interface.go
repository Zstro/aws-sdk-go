// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworkscmiface provides an interface to enable mocking the AWS OpsWorks for Chef Automate service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opsworkscmiface

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
	"github.com/journeymidnight/aws-sdk-go/service/opsworkscm"
)

// OpsWorksCMAPI provides an interface to enable mocking the
// opsworkscm.OpsWorksCM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS OpsWorks for Chef Automate.
//    func myFunc(svc opsworkscmiface.OpsWorksCMAPI) bool {
//        // Make svc.AssociateNode request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := opsworkscm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOpsWorksCMClient struct {
//        opsworkscmiface.OpsWorksCMAPI
//    }
//    func (m *mockOpsWorksCMClient) AssociateNode(input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOpsWorksCMClient{}
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
type OpsWorksCMAPI interface {
	AssociateNode(*opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error)
	AssociateNodeWithContext(aws.Context, *opsworkscm.AssociateNodeInput, ...request.Option) (*opsworkscm.AssociateNodeOutput, error)
	AssociateNodeRequest(*opsworkscm.AssociateNodeInput) (*request.Request, *opsworkscm.AssociateNodeOutput)

	CreateBackup(*opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error)
	CreateBackupWithContext(aws.Context, *opsworkscm.CreateBackupInput, ...request.Option) (*opsworkscm.CreateBackupOutput, error)
	CreateBackupRequest(*opsworkscm.CreateBackupInput) (*request.Request, *opsworkscm.CreateBackupOutput)

	CreateServer(*opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error)
	CreateServerWithContext(aws.Context, *opsworkscm.CreateServerInput, ...request.Option) (*opsworkscm.CreateServerOutput, error)
	CreateServerRequest(*opsworkscm.CreateServerInput) (*request.Request, *opsworkscm.CreateServerOutput)

	DeleteBackup(*opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error)
	DeleteBackupWithContext(aws.Context, *opsworkscm.DeleteBackupInput, ...request.Option) (*opsworkscm.DeleteBackupOutput, error)
	DeleteBackupRequest(*opsworkscm.DeleteBackupInput) (*request.Request, *opsworkscm.DeleteBackupOutput)

	DeleteServer(*opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error)
	DeleteServerWithContext(aws.Context, *opsworkscm.DeleteServerInput, ...request.Option) (*opsworkscm.DeleteServerOutput, error)
	DeleteServerRequest(*opsworkscm.DeleteServerInput) (*request.Request, *opsworkscm.DeleteServerOutput)

	DescribeAccountAttributes(*opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesWithContext(aws.Context, *opsworkscm.DescribeAccountAttributesInput, ...request.Option) (*opsworkscm.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesRequest(*opsworkscm.DescribeAccountAttributesInput) (*request.Request, *opsworkscm.DescribeAccountAttributesOutput)

	DescribeBackups(*opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error)
	DescribeBackupsWithContext(aws.Context, *opsworkscm.DescribeBackupsInput, ...request.Option) (*opsworkscm.DescribeBackupsOutput, error)
	DescribeBackupsRequest(*opsworkscm.DescribeBackupsInput) (*request.Request, *opsworkscm.DescribeBackupsOutput)

	DescribeEvents(*opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error)
	DescribeEventsWithContext(aws.Context, *opsworkscm.DescribeEventsInput, ...request.Option) (*opsworkscm.DescribeEventsOutput, error)
	DescribeEventsRequest(*opsworkscm.DescribeEventsInput) (*request.Request, *opsworkscm.DescribeEventsOutput)

	DescribeNodeAssociationStatus(*opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error)
	DescribeNodeAssociationStatusWithContext(aws.Context, *opsworkscm.DescribeNodeAssociationStatusInput, ...request.Option) (*opsworkscm.DescribeNodeAssociationStatusOutput, error)
	DescribeNodeAssociationStatusRequest(*opsworkscm.DescribeNodeAssociationStatusInput) (*request.Request, *opsworkscm.DescribeNodeAssociationStatusOutput)

	DescribeServers(*opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error)
	DescribeServersWithContext(aws.Context, *opsworkscm.DescribeServersInput, ...request.Option) (*opsworkscm.DescribeServersOutput, error)
	DescribeServersRequest(*opsworkscm.DescribeServersInput) (*request.Request, *opsworkscm.DescribeServersOutput)

	DisassociateNode(*opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error)
	DisassociateNodeWithContext(aws.Context, *opsworkscm.DisassociateNodeInput, ...request.Option) (*opsworkscm.DisassociateNodeOutput, error)
	DisassociateNodeRequest(*opsworkscm.DisassociateNodeInput) (*request.Request, *opsworkscm.DisassociateNodeOutput)

	ExportServerEngineAttribute(*opsworkscm.ExportServerEngineAttributeInput) (*opsworkscm.ExportServerEngineAttributeOutput, error)
	ExportServerEngineAttributeWithContext(aws.Context, *opsworkscm.ExportServerEngineAttributeInput, ...request.Option) (*opsworkscm.ExportServerEngineAttributeOutput, error)
	ExportServerEngineAttributeRequest(*opsworkscm.ExportServerEngineAttributeInput) (*request.Request, *opsworkscm.ExportServerEngineAttributeOutput)

	RestoreServer(*opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error)
	RestoreServerWithContext(aws.Context, *opsworkscm.RestoreServerInput, ...request.Option) (*opsworkscm.RestoreServerOutput, error)
	RestoreServerRequest(*opsworkscm.RestoreServerInput) (*request.Request, *opsworkscm.RestoreServerOutput)

	StartMaintenance(*opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error)
	StartMaintenanceWithContext(aws.Context, *opsworkscm.StartMaintenanceInput, ...request.Option) (*opsworkscm.StartMaintenanceOutput, error)
	StartMaintenanceRequest(*opsworkscm.StartMaintenanceInput) (*request.Request, *opsworkscm.StartMaintenanceOutput)

	UpdateServer(*opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error)
	UpdateServerWithContext(aws.Context, *opsworkscm.UpdateServerInput, ...request.Option) (*opsworkscm.UpdateServerOutput, error)
	UpdateServerRequest(*opsworkscm.UpdateServerInput) (*request.Request, *opsworkscm.UpdateServerOutput)

	UpdateServerEngineAttributes(*opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error)
	UpdateServerEngineAttributesWithContext(aws.Context, *opsworkscm.UpdateServerEngineAttributesInput, ...request.Option) (*opsworkscm.UpdateServerEngineAttributesOutput, error)
	UpdateServerEngineAttributesRequest(*opsworkscm.UpdateServerEngineAttributesInput) (*request.Request, *opsworkscm.UpdateServerEngineAttributesOutput)

	WaitUntilNodeAssociated(*opsworkscm.DescribeNodeAssociationStatusInput) error
	WaitUntilNodeAssociatedWithContext(aws.Context, *opsworkscm.DescribeNodeAssociationStatusInput, ...request.WaiterOption) error
}

var _ OpsWorksCMAPI = (*opsworkscm.OpsWorksCM)(nil)
