// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/awserr"
	"github.com/journeymidnight/aws-sdk-go/aws/session"
	"github.com/journeymidnight/aws-sdk-go/service/cloudfront"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

//

func ExampleCloudFront_CreateCloudFrontOriginAccessIdentity_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.CreateCloudFrontOriginAccessIdentityInput{}

	result, err := svc.CreateCloudFrontOriginAccessIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeOriginAccessIdentityAlreadyExists:
				fmt.Println(cloudfront.ErrCodeOriginAccessIdentityAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeTooManyCloudFrontOriginAccessIdentities:
				fmt.Println(cloudfront.ErrCodeTooManyCloudFrontOriginAccessIdentities, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_CreateDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.CreateDistributionInput{}

	result, err := svc.CreateDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeCNAMEAlreadyExists:
				fmt.Println(cloudfront.ErrCodeCNAMEAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeDistributionAlreadyExists:
				fmt.Println(cloudfront.ErrCodeDistributionAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeInvalidOrigin:
				fmt.Println(cloudfront.ErrCodeInvalidOrigin, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeInvalidOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeTooManyTrustedSigners:
				fmt.Println(cloudfront.ErrCodeTooManyTrustedSigners, aerr.Error())
			case cloudfront.ErrCodeTrustedSignerDoesNotExist:
				fmt.Println(cloudfront.ErrCodeTrustedSignerDoesNotExist, aerr.Error())
			case cloudfront.ErrCodeInvalidViewerCertificate:
				fmt.Println(cloudfront.ErrCodeInvalidViewerCertificate, aerr.Error())
			case cloudfront.ErrCodeInvalidMinimumProtocolVersion:
				fmt.Println(cloudfront.ErrCodeInvalidMinimumProtocolVersion, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributionCNAMEs:
				fmt.Println(cloudfront.ErrCodeTooManyDistributionCNAMEs, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributions:
				fmt.Println(cloudfront.ErrCodeTooManyDistributions, aerr.Error())
			case cloudfront.ErrCodeInvalidDefaultRootObject:
				fmt.Println(cloudfront.ErrCodeInvalidDefaultRootObject, aerr.Error())
			case cloudfront.ErrCodeInvalidRelativePath:
				fmt.Println(cloudfront.ErrCodeInvalidRelativePath, aerr.Error())
			case cloudfront.ErrCodeInvalidErrorCode:
				fmt.Println(cloudfront.ErrCodeInvalidErrorCode, aerr.Error())
			case cloudfront.ErrCodeInvalidResponseCode:
				fmt.Println(cloudfront.ErrCodeInvalidResponseCode, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidRequiredProtocol:
				fmt.Println(cloudfront.ErrCodeInvalidRequiredProtocol, aerr.Error())
			case cloudfront.ErrCodeNoSuchOrigin:
				fmt.Println(cloudfront.ErrCodeNoSuchOrigin, aerr.Error())
			case cloudfront.ErrCodeTooManyOrigins:
				fmt.Println(cloudfront.ErrCodeTooManyOrigins, aerr.Error())
			case cloudfront.ErrCodeTooManyCacheBehaviors:
				fmt.Println(cloudfront.ErrCodeTooManyCacheBehaviors, aerr.Error())
			case cloudfront.ErrCodeTooManyCookieNamesInWhiteList:
				fmt.Println(cloudfront.ErrCodeTooManyCookieNamesInWhiteList, aerr.Error())
			case cloudfront.ErrCodeInvalidForwardCookies:
				fmt.Println(cloudfront.ErrCodeInvalidForwardCookies, aerr.Error())
			case cloudfront.ErrCodeTooManyHeadersInForwardedValues:
				fmt.Println(cloudfront.ErrCodeTooManyHeadersInForwardedValues, aerr.Error())
			case cloudfront.ErrCodeInvalidHeadersForS3Origin:
				fmt.Println(cloudfront.ErrCodeInvalidHeadersForS3Origin, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			case cloudfront.ErrCodeTooManyCertificates:
				fmt.Println(cloudfront.ErrCodeTooManyCertificates, aerr.Error())
			case cloudfront.ErrCodeInvalidLocationCode:
				fmt.Println(cloudfront.ErrCodeInvalidLocationCode, aerr.Error())
			case cloudfront.ErrCodeInvalidGeoRestrictionParameter:
				fmt.Println(cloudfront.ErrCodeInvalidGeoRestrictionParameter, aerr.Error())
			case cloudfront.ErrCodeInvalidProtocolSettings:
				fmt.Println(cloudfront.ErrCodeInvalidProtocolSettings, aerr.Error())
			case cloudfront.ErrCodeInvalidTTLOrder:
				fmt.Println(cloudfront.ErrCodeInvalidTTLOrder, aerr.Error())
			case cloudfront.ErrCodeInvalidWebACLId:
				fmt.Println(cloudfront.ErrCodeInvalidWebACLId, aerr.Error())
			case cloudfront.ErrCodeTooManyOriginCustomHeaders:
				fmt.Println(cloudfront.ErrCodeTooManyOriginCustomHeaders, aerr.Error())
			case cloudfront.ErrCodeTooManyQueryStringParameters:
				fmt.Println(cloudfront.ErrCodeTooManyQueryStringParameters, aerr.Error())
			case cloudfront.ErrCodeInvalidQueryStringParameters:
				fmt.Println(cloudfront.ErrCodeInvalidQueryStringParameters, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributionsWithLambdaAssociations:
				fmt.Println(cloudfront.ErrCodeTooManyDistributionsWithLambdaAssociations, aerr.Error())
			case cloudfront.ErrCodeTooManyLambdaFunctionAssociations:
				fmt.Println(cloudfront.ErrCodeTooManyLambdaFunctionAssociations, aerr.Error())
			case cloudfront.ErrCodeInvalidLambdaFunctionAssociation:
				fmt.Println(cloudfront.ErrCodeInvalidLambdaFunctionAssociation, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginReadTimeout:
				fmt.Println(cloudfront.ErrCodeInvalidOriginReadTimeout, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginKeepaliveTimeout:
				fmt.Println(cloudfront.ErrCodeInvalidOriginKeepaliveTimeout, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_CreateDistributionWithTags_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.CreateDistributionWithTagsInput{}

	result, err := svc.CreateDistributionWithTags(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeCNAMEAlreadyExists:
				fmt.Println(cloudfront.ErrCodeCNAMEAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeDistributionAlreadyExists:
				fmt.Println(cloudfront.ErrCodeDistributionAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeInvalidOrigin:
				fmt.Println(cloudfront.ErrCodeInvalidOrigin, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeInvalidOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeTooManyTrustedSigners:
				fmt.Println(cloudfront.ErrCodeTooManyTrustedSigners, aerr.Error())
			case cloudfront.ErrCodeTrustedSignerDoesNotExist:
				fmt.Println(cloudfront.ErrCodeTrustedSignerDoesNotExist, aerr.Error())
			case cloudfront.ErrCodeInvalidViewerCertificate:
				fmt.Println(cloudfront.ErrCodeInvalidViewerCertificate, aerr.Error())
			case cloudfront.ErrCodeInvalidMinimumProtocolVersion:
				fmt.Println(cloudfront.ErrCodeInvalidMinimumProtocolVersion, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributionCNAMEs:
				fmt.Println(cloudfront.ErrCodeTooManyDistributionCNAMEs, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributions:
				fmt.Println(cloudfront.ErrCodeTooManyDistributions, aerr.Error())
			case cloudfront.ErrCodeInvalidDefaultRootObject:
				fmt.Println(cloudfront.ErrCodeInvalidDefaultRootObject, aerr.Error())
			case cloudfront.ErrCodeInvalidRelativePath:
				fmt.Println(cloudfront.ErrCodeInvalidRelativePath, aerr.Error())
			case cloudfront.ErrCodeInvalidErrorCode:
				fmt.Println(cloudfront.ErrCodeInvalidErrorCode, aerr.Error())
			case cloudfront.ErrCodeInvalidResponseCode:
				fmt.Println(cloudfront.ErrCodeInvalidResponseCode, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidRequiredProtocol:
				fmt.Println(cloudfront.ErrCodeInvalidRequiredProtocol, aerr.Error())
			case cloudfront.ErrCodeNoSuchOrigin:
				fmt.Println(cloudfront.ErrCodeNoSuchOrigin, aerr.Error())
			case cloudfront.ErrCodeTooManyOrigins:
				fmt.Println(cloudfront.ErrCodeTooManyOrigins, aerr.Error())
			case cloudfront.ErrCodeTooManyCacheBehaviors:
				fmt.Println(cloudfront.ErrCodeTooManyCacheBehaviors, aerr.Error())
			case cloudfront.ErrCodeTooManyCookieNamesInWhiteList:
				fmt.Println(cloudfront.ErrCodeTooManyCookieNamesInWhiteList, aerr.Error())
			case cloudfront.ErrCodeInvalidForwardCookies:
				fmt.Println(cloudfront.ErrCodeInvalidForwardCookies, aerr.Error())
			case cloudfront.ErrCodeTooManyHeadersInForwardedValues:
				fmt.Println(cloudfront.ErrCodeTooManyHeadersInForwardedValues, aerr.Error())
			case cloudfront.ErrCodeInvalidHeadersForS3Origin:
				fmt.Println(cloudfront.ErrCodeInvalidHeadersForS3Origin, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			case cloudfront.ErrCodeTooManyCertificates:
				fmt.Println(cloudfront.ErrCodeTooManyCertificates, aerr.Error())
			case cloudfront.ErrCodeInvalidLocationCode:
				fmt.Println(cloudfront.ErrCodeInvalidLocationCode, aerr.Error())
			case cloudfront.ErrCodeInvalidGeoRestrictionParameter:
				fmt.Println(cloudfront.ErrCodeInvalidGeoRestrictionParameter, aerr.Error())
			case cloudfront.ErrCodeInvalidProtocolSettings:
				fmt.Println(cloudfront.ErrCodeInvalidProtocolSettings, aerr.Error())
			case cloudfront.ErrCodeInvalidTTLOrder:
				fmt.Println(cloudfront.ErrCodeInvalidTTLOrder, aerr.Error())
			case cloudfront.ErrCodeInvalidWebACLId:
				fmt.Println(cloudfront.ErrCodeInvalidWebACLId, aerr.Error())
			case cloudfront.ErrCodeTooManyOriginCustomHeaders:
				fmt.Println(cloudfront.ErrCodeTooManyOriginCustomHeaders, aerr.Error())
			case cloudfront.ErrCodeInvalidTagging:
				fmt.Println(cloudfront.ErrCodeInvalidTagging, aerr.Error())
			case cloudfront.ErrCodeTooManyQueryStringParameters:
				fmt.Println(cloudfront.ErrCodeTooManyQueryStringParameters, aerr.Error())
			case cloudfront.ErrCodeInvalidQueryStringParameters:
				fmt.Println(cloudfront.ErrCodeInvalidQueryStringParameters, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributionsWithLambdaAssociations:
				fmt.Println(cloudfront.ErrCodeTooManyDistributionsWithLambdaAssociations, aerr.Error())
			case cloudfront.ErrCodeTooManyLambdaFunctionAssociations:
				fmt.Println(cloudfront.ErrCodeTooManyLambdaFunctionAssociations, aerr.Error())
			case cloudfront.ErrCodeInvalidLambdaFunctionAssociation:
				fmt.Println(cloudfront.ErrCodeInvalidLambdaFunctionAssociation, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginReadTimeout:
				fmt.Println(cloudfront.ErrCodeInvalidOriginReadTimeout, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginKeepaliveTimeout:
				fmt.Println(cloudfront.ErrCodeInvalidOriginKeepaliveTimeout, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_CreateInvalidation_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.CreateInvalidationInput{}

	result, err := svc.CreateInvalidation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodeBatchTooLarge:
				fmt.Println(cloudfront.ErrCodeBatchTooLarge, aerr.Error())
			case cloudfront.ErrCodeTooManyInvalidationsInProgress:
				fmt.Println(cloudfront.ErrCodeTooManyInvalidationsInProgress, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_CreateStreamingDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.CreateStreamingDistributionInput{}

	result, err := svc.CreateStreamingDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeCNAMEAlreadyExists:
				fmt.Println(cloudfront.ErrCodeCNAMEAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeStreamingDistributionAlreadyExists:
				fmt.Println(cloudfront.ErrCodeStreamingDistributionAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeInvalidOrigin:
				fmt.Println(cloudfront.ErrCodeInvalidOrigin, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeInvalidOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeTooManyTrustedSigners:
				fmt.Println(cloudfront.ErrCodeTooManyTrustedSigners, aerr.Error())
			case cloudfront.ErrCodeTrustedSignerDoesNotExist:
				fmt.Println(cloudfront.ErrCodeTrustedSignerDoesNotExist, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeTooManyStreamingDistributionCNAMEs:
				fmt.Println(cloudfront.ErrCodeTooManyStreamingDistributionCNAMEs, aerr.Error())
			case cloudfront.ErrCodeTooManyStreamingDistributions:
				fmt.Println(cloudfront.ErrCodeTooManyStreamingDistributions, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_DeleteCloudFrontOriginAccessIdentity_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.DeleteCloudFrontOriginAccessIdentityInput{}

	result, err := svc.DeleteCloudFrontOriginAccessIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeInvalidIfMatchVersion:
				fmt.Println(cloudfront.ErrCodeInvalidIfMatchVersion, aerr.Error())
			case cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodePreconditionFailed:
				fmt.Println(cloudfront.ErrCodePreconditionFailed, aerr.Error())
			case cloudfront.ErrCodeOriginAccessIdentityInUse:
				fmt.Println(cloudfront.ErrCodeOriginAccessIdentityInUse, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_DeleteDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.DeleteDistributionInput{}

	result, err := svc.DeleteDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeDistributionNotDisabled:
				fmt.Println(cloudfront.ErrCodeDistributionNotDisabled, aerr.Error())
			case cloudfront.ErrCodeInvalidIfMatchVersion:
				fmt.Println(cloudfront.ErrCodeInvalidIfMatchVersion, aerr.Error())
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodePreconditionFailed:
				fmt.Println(cloudfront.ErrCodePreconditionFailed, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_DeleteStreamingDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.DeleteStreamingDistributionInput{}

	result, err := svc.DeleteStreamingDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeStreamingDistributionNotDisabled:
				fmt.Println(cloudfront.ErrCodeStreamingDistributionNotDisabled, aerr.Error())
			case cloudfront.ErrCodeInvalidIfMatchVersion:
				fmt.Println(cloudfront.ErrCodeInvalidIfMatchVersion, aerr.Error())
			case cloudfront.ErrCodeNoSuchStreamingDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchStreamingDistribution, aerr.Error())
			case cloudfront.ErrCodePreconditionFailed:
				fmt.Println(cloudfront.ErrCodePreconditionFailed, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetCloudFrontOriginAccessIdentity_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetCloudFrontOriginAccessIdentityInput{}

	result, err := svc.GetCloudFrontOriginAccessIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetCloudFrontOriginAccessIdentityConfig_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetCloudFrontOriginAccessIdentityConfigInput{}

	result, err := svc.GetCloudFrontOriginAccessIdentityConfig(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetDistributionInput{}

	result, err := svc.GetDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetDistributionConfig_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetDistributionConfigInput{}

	result, err := svc.GetDistributionConfig(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetInvalidation_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetInvalidationInput{}

	result, err := svc.GetInvalidation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchInvalidation:
				fmt.Println(cloudfront.ErrCodeNoSuchInvalidation, aerr.Error())
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetStreamingDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetStreamingDistributionInput{}

	result, err := svc.GetStreamingDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchStreamingDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchStreamingDistribution, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_GetStreamingDistributionConfig_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.GetStreamingDistributionConfigInput{}

	result, err := svc.GetStreamingDistributionConfig(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeNoSuchStreamingDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchStreamingDistribution, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_ListCloudFrontOriginAccessIdentities_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}

	result, err := svc.ListCloudFrontOriginAccessIdentities(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_ListDistributions_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.ListDistributionsInput{}

	result, err := svc.ListDistributions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_ListDistributionsByWebACLId_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.ListDistributionsByWebACLIdInput{}

	result, err := svc.ListDistributionsByWebACLId(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidWebACLId:
				fmt.Println(cloudfront.ErrCodeInvalidWebACLId, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_ListInvalidations_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.ListInvalidationsInput{}

	result, err := svc.ListInvalidations(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_ListStreamingDistributions_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.ListStreamingDistributionsInput{}

	result, err := svc.ListStreamingDistributions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_ListTagsForResource_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.ListTagsForResourceInput{}

	result, err := svc.ListTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidTagging:
				fmt.Println(cloudfront.ErrCodeInvalidTagging, aerr.Error())
			case cloudfront.ErrCodeNoSuchResource:
				fmt.Println(cloudfront.ErrCodeNoSuchResource, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_TagResource_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.TagResourceInput{}

	result, err := svc.TagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidTagging:
				fmt.Println(cloudfront.ErrCodeInvalidTagging, aerr.Error())
			case cloudfront.ErrCodeNoSuchResource:
				fmt.Println(cloudfront.ErrCodeNoSuchResource, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_UntagResource_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.UntagResourceInput{}

	result, err := svc.UntagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidTagging:
				fmt.Println(cloudfront.ErrCodeInvalidTagging, aerr.Error())
			case cloudfront.ErrCodeNoSuchResource:
				fmt.Println(cloudfront.ErrCodeNoSuchResource, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_UpdateCloudFrontOriginAccessIdentity_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.UpdateCloudFrontOriginAccessIdentityInput{}

	result, err := svc.UpdateCloudFrontOriginAccessIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeIllegalUpdate:
				fmt.Println(cloudfront.ErrCodeIllegalUpdate, aerr.Error())
			case cloudfront.ErrCodeInvalidIfMatchVersion:
				fmt.Println(cloudfront.ErrCodeInvalidIfMatchVersion, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeNoSuchCloudFrontOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodePreconditionFailed:
				fmt.Println(cloudfront.ErrCodePreconditionFailed, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_UpdateDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.UpdateDistributionInput{}

	result, err := svc.UpdateDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeCNAMEAlreadyExists:
				fmt.Println(cloudfront.ErrCodeCNAMEAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeIllegalUpdate:
				fmt.Println(cloudfront.ErrCodeIllegalUpdate, aerr.Error())
			case cloudfront.ErrCodeInvalidIfMatchVersion:
				fmt.Println(cloudfront.ErrCodeInvalidIfMatchVersion, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeNoSuchDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchDistribution, aerr.Error())
			case cloudfront.ErrCodePreconditionFailed:
				fmt.Println(cloudfront.ErrCodePreconditionFailed, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributionCNAMEs:
				fmt.Println(cloudfront.ErrCodeTooManyDistributionCNAMEs, aerr.Error())
			case cloudfront.ErrCodeInvalidDefaultRootObject:
				fmt.Println(cloudfront.ErrCodeInvalidDefaultRootObject, aerr.Error())
			case cloudfront.ErrCodeInvalidRelativePath:
				fmt.Println(cloudfront.ErrCodeInvalidRelativePath, aerr.Error())
			case cloudfront.ErrCodeInvalidErrorCode:
				fmt.Println(cloudfront.ErrCodeInvalidErrorCode, aerr.Error())
			case cloudfront.ErrCodeInvalidResponseCode:
				fmt.Println(cloudfront.ErrCodeInvalidResponseCode, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeInvalidOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeTooManyTrustedSigners:
				fmt.Println(cloudfront.ErrCodeTooManyTrustedSigners, aerr.Error())
			case cloudfront.ErrCodeTrustedSignerDoesNotExist:
				fmt.Println(cloudfront.ErrCodeTrustedSignerDoesNotExist, aerr.Error())
			case cloudfront.ErrCodeInvalidViewerCertificate:
				fmt.Println(cloudfront.ErrCodeInvalidViewerCertificate, aerr.Error())
			case cloudfront.ErrCodeInvalidMinimumProtocolVersion:
				fmt.Println(cloudfront.ErrCodeInvalidMinimumProtocolVersion, aerr.Error())
			case cloudfront.ErrCodeInvalidRequiredProtocol:
				fmt.Println(cloudfront.ErrCodeInvalidRequiredProtocol, aerr.Error())
			case cloudfront.ErrCodeNoSuchOrigin:
				fmt.Println(cloudfront.ErrCodeNoSuchOrigin, aerr.Error())
			case cloudfront.ErrCodeTooManyOrigins:
				fmt.Println(cloudfront.ErrCodeTooManyOrigins, aerr.Error())
			case cloudfront.ErrCodeTooManyCacheBehaviors:
				fmt.Println(cloudfront.ErrCodeTooManyCacheBehaviors, aerr.Error())
			case cloudfront.ErrCodeTooManyCookieNamesInWhiteList:
				fmt.Println(cloudfront.ErrCodeTooManyCookieNamesInWhiteList, aerr.Error())
			case cloudfront.ErrCodeInvalidForwardCookies:
				fmt.Println(cloudfront.ErrCodeInvalidForwardCookies, aerr.Error())
			case cloudfront.ErrCodeTooManyHeadersInForwardedValues:
				fmt.Println(cloudfront.ErrCodeTooManyHeadersInForwardedValues, aerr.Error())
			case cloudfront.ErrCodeInvalidHeadersForS3Origin:
				fmt.Println(cloudfront.ErrCodeInvalidHeadersForS3Origin, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			case cloudfront.ErrCodeTooManyCertificates:
				fmt.Println(cloudfront.ErrCodeTooManyCertificates, aerr.Error())
			case cloudfront.ErrCodeInvalidLocationCode:
				fmt.Println(cloudfront.ErrCodeInvalidLocationCode, aerr.Error())
			case cloudfront.ErrCodeInvalidGeoRestrictionParameter:
				fmt.Println(cloudfront.ErrCodeInvalidGeoRestrictionParameter, aerr.Error())
			case cloudfront.ErrCodeInvalidTTLOrder:
				fmt.Println(cloudfront.ErrCodeInvalidTTLOrder, aerr.Error())
			case cloudfront.ErrCodeInvalidWebACLId:
				fmt.Println(cloudfront.ErrCodeInvalidWebACLId, aerr.Error())
			case cloudfront.ErrCodeTooManyOriginCustomHeaders:
				fmt.Println(cloudfront.ErrCodeTooManyOriginCustomHeaders, aerr.Error())
			case cloudfront.ErrCodeTooManyQueryStringParameters:
				fmt.Println(cloudfront.ErrCodeTooManyQueryStringParameters, aerr.Error())
			case cloudfront.ErrCodeInvalidQueryStringParameters:
				fmt.Println(cloudfront.ErrCodeInvalidQueryStringParameters, aerr.Error())
			case cloudfront.ErrCodeTooManyDistributionsWithLambdaAssociations:
				fmt.Println(cloudfront.ErrCodeTooManyDistributionsWithLambdaAssociations, aerr.Error())
			case cloudfront.ErrCodeTooManyLambdaFunctionAssociations:
				fmt.Println(cloudfront.ErrCodeTooManyLambdaFunctionAssociations, aerr.Error())
			case cloudfront.ErrCodeInvalidLambdaFunctionAssociation:
				fmt.Println(cloudfront.ErrCodeInvalidLambdaFunctionAssociation, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginReadTimeout:
				fmt.Println(cloudfront.ErrCodeInvalidOriginReadTimeout, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginKeepaliveTimeout:
				fmt.Println(cloudfront.ErrCodeInvalidOriginKeepaliveTimeout, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

//

func ExampleCloudFront_UpdateStreamingDistribution_shared00() {
	svc := cloudfront.New(session.New())
	input := &cloudfront.UpdateStreamingDistributionInput{}

	result, err := svc.UpdateStreamingDistribution(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudfront.ErrCodeAccessDenied:
				fmt.Println(cloudfront.ErrCodeAccessDenied, aerr.Error())
			case cloudfront.ErrCodeCNAMEAlreadyExists:
				fmt.Println(cloudfront.ErrCodeCNAMEAlreadyExists, aerr.Error())
			case cloudfront.ErrCodeIllegalUpdate:
				fmt.Println(cloudfront.ErrCodeIllegalUpdate, aerr.Error())
			case cloudfront.ErrCodeInvalidIfMatchVersion:
				fmt.Println(cloudfront.ErrCodeInvalidIfMatchVersion, aerr.Error())
			case cloudfront.ErrCodeMissingBody:
				fmt.Println(cloudfront.ErrCodeMissingBody, aerr.Error())
			case cloudfront.ErrCodeNoSuchStreamingDistribution:
				fmt.Println(cloudfront.ErrCodeNoSuchStreamingDistribution, aerr.Error())
			case cloudfront.ErrCodePreconditionFailed:
				fmt.Println(cloudfront.ErrCodePreconditionFailed, aerr.Error())
			case cloudfront.ErrCodeTooManyStreamingDistributionCNAMEs:
				fmt.Println(cloudfront.ErrCodeTooManyStreamingDistributionCNAMEs, aerr.Error())
			case cloudfront.ErrCodeInvalidArgument:
				fmt.Println(cloudfront.ErrCodeInvalidArgument, aerr.Error())
			case cloudfront.ErrCodeInvalidOriginAccessIdentity:
				fmt.Println(cloudfront.ErrCodeInvalidOriginAccessIdentity, aerr.Error())
			case cloudfront.ErrCodeTooManyTrustedSigners:
				fmt.Println(cloudfront.ErrCodeTooManyTrustedSigners, aerr.Error())
			case cloudfront.ErrCodeTrustedSignerDoesNotExist:
				fmt.Println(cloudfront.ErrCodeTrustedSignerDoesNotExist, aerr.Error())
			case cloudfront.ErrCodeInconsistentQuantities:
				fmt.Println(cloudfront.ErrCodeInconsistentQuantities, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
