// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing

import (
	"fmt"

	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/awsutil"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
)

const opDescribeServices = "DescribeServices"

// DescribeServicesRequest generates a "aws/request.Request" representing the
// client's request for the DescribeServices operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeServices for more information on using the DescribeServices
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeServicesRequest method.
//    req, resp := client.DescribeServicesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/DescribeServices
func (c *Pricing) DescribeServicesRequest(input *DescribeServicesInput) (req *request.Request, output *DescribeServicesOutput) {
	op := &request.Operation{
		Name:       opDescribeServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeServicesInput{}
	}

	output = &DescribeServicesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeServices API operation for AWS Price List Service.
//
// Returns the metadata for one service or a list of the metadata for all services.
// Use this without a service code to get the service codes for all services.
// Use it with a service code, such as AmazonEC2, to get information specific
// to that service, such as the attribute names available for that service.
// For example, some of the attribute names available for EC2 are volumeType,
// maxIopsVolume, operation, locationType, and instanceCapacity10xlarge.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Price List Service's
// API operation DescribeServices for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalErrorException "InternalErrorException"
//   An error on the server occurred during the processing of your request. Try
//   again later.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   One or more parameters had an invalid value.
//
//   * ErrCodeNotFoundException "NotFoundException"
//   The requested resource can't be found.
//
//   * ErrCodeInvalidNextTokenException "InvalidNextTokenException"
//   The pagination token is invalid. Try again without a pagination token.
//
//   * ErrCodeExpiredNextTokenException "ExpiredNextTokenException"
//   The pagination token expired. Try again without a pagination token.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/DescribeServices
func (c *Pricing) DescribeServices(input *DescribeServicesInput) (*DescribeServicesOutput, error) {
	req, out := c.DescribeServicesRequest(input)
	return out, req.Send()
}

// DescribeServicesWithContext is the same as DescribeServices with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeServices for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) DescribeServicesWithContext(ctx aws.Context, input *DescribeServicesInput, opts ...request.Option) (*DescribeServicesOutput, error) {
	req, out := c.DescribeServicesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// DescribeServicesPages iterates over the pages of a DescribeServices operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See DescribeServices method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a DescribeServices operation.
//    pageNum := 0
//    err := client.DescribeServicesPages(params,
//        func(page *DescribeServicesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *Pricing) DescribeServicesPages(input *DescribeServicesInput, fn func(*DescribeServicesOutput, bool) bool) error {
	return c.DescribeServicesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// DescribeServicesPagesWithContext same as DescribeServicesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) DescribeServicesPagesWithContext(ctx aws.Context, input *DescribeServicesInput, fn func(*DescribeServicesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *DescribeServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*DescribeServicesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opGetAttributeValues = "GetAttributeValues"

// GetAttributeValuesRequest generates a "aws/request.Request" representing the
// client's request for the GetAttributeValues operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetAttributeValues for more information on using the GetAttributeValues
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetAttributeValuesRequest method.
//    req, resp := client.GetAttributeValuesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetAttributeValues
func (c *Pricing) GetAttributeValuesRequest(input *GetAttributeValuesInput) (req *request.Request, output *GetAttributeValuesOutput) {
	op := &request.Operation{
		Name:       opGetAttributeValues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetAttributeValuesInput{}
	}

	output = &GetAttributeValuesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetAttributeValues API operation for AWS Price List Service.
//
// Returns a list of attribute values. Attibutes are similar to the details
// in a Price List API offer file. For a list of available attributes, see Offer
// File Definitions (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html#pps-defs)
// in the AWS Billing and Cost Management User Guide (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Price List Service's
// API operation GetAttributeValues for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalErrorException "InternalErrorException"
//   An error on the server occurred during the processing of your request. Try
//   again later.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   One or more parameters had an invalid value.
//
//   * ErrCodeNotFoundException "NotFoundException"
//   The requested resource can't be found.
//
//   * ErrCodeInvalidNextTokenException "InvalidNextTokenException"
//   The pagination token is invalid. Try again without a pagination token.
//
//   * ErrCodeExpiredNextTokenException "ExpiredNextTokenException"
//   The pagination token expired. Try again without a pagination token.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetAttributeValues
func (c *Pricing) GetAttributeValues(input *GetAttributeValuesInput) (*GetAttributeValuesOutput, error) {
	req, out := c.GetAttributeValuesRequest(input)
	return out, req.Send()
}

// GetAttributeValuesWithContext is the same as GetAttributeValues with the addition of
// the ability to pass a context and additional request options.
//
// See GetAttributeValues for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) GetAttributeValuesWithContext(ctx aws.Context, input *GetAttributeValuesInput, opts ...request.Option) (*GetAttributeValuesOutput, error) {
	req, out := c.GetAttributeValuesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// GetAttributeValuesPages iterates over the pages of a GetAttributeValues operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See GetAttributeValues method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a GetAttributeValues operation.
//    pageNum := 0
//    err := client.GetAttributeValuesPages(params,
//        func(page *GetAttributeValuesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *Pricing) GetAttributeValuesPages(input *GetAttributeValuesInput, fn func(*GetAttributeValuesOutput, bool) bool) error {
	return c.GetAttributeValuesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// GetAttributeValuesPagesWithContext same as GetAttributeValuesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) GetAttributeValuesPagesWithContext(ctx aws.Context, input *GetAttributeValuesInput, fn func(*GetAttributeValuesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *GetAttributeValuesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetAttributeValuesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*GetAttributeValuesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opGetProducts = "GetProducts"

// GetProductsRequest generates a "aws/request.Request" representing the
// client's request for the GetProducts operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetProducts for more information on using the GetProducts
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetProductsRequest method.
//    req, resp := client.GetProductsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetProducts
func (c *Pricing) GetProductsRequest(input *GetProductsInput) (req *request.Request, output *GetProductsOutput) {
	op := &request.Operation{
		Name:       opGetProducts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetProductsInput{}
	}

	output = &GetProductsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetProducts API operation for AWS Price List Service.
//
// Returns a list of all products that match the filter criteria.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Price List Service's
// API operation GetProducts for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalErrorException "InternalErrorException"
//   An error on the server occurred during the processing of your request. Try
//   again later.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   One or more parameters had an invalid value.
//
//   * ErrCodeNotFoundException "NotFoundException"
//   The requested resource can't be found.
//
//   * ErrCodeInvalidNextTokenException "InvalidNextTokenException"
//   The pagination token is invalid. Try again without a pagination token.
//
//   * ErrCodeExpiredNextTokenException "ExpiredNextTokenException"
//   The pagination token expired. Try again without a pagination token.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetProducts
func (c *Pricing) GetProducts(input *GetProductsInput) (*GetProductsOutput, error) {
	req, out := c.GetProductsRequest(input)
	return out, req.Send()
}

// GetProductsWithContext is the same as GetProducts with the addition of
// the ability to pass a context and additional request options.
//
// See GetProducts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) GetProductsWithContext(ctx aws.Context, input *GetProductsInput, opts ...request.Option) (*GetProductsOutput, error) {
	req, out := c.GetProductsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// GetProductsPages iterates over the pages of a GetProducts operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See GetProducts method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a GetProducts operation.
//    pageNum := 0
//    err := client.GetProductsPages(params,
//        func(page *GetProductsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *Pricing) GetProductsPages(input *GetProductsInput, fn func(*GetProductsOutput, bool) bool) error {
	return c.GetProductsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// GetProductsPagesWithContext same as GetProductsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) GetProductsPagesWithContext(ctx aws.Context, input *GetProductsInput, fn func(*GetProductsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *GetProductsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetProductsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*GetProductsOutput), !p.HasNextPage())
	}
	return p.Err()
}

// The values of a given attribute, such as Throughput Optimized HDD or Provisioned
// IOPS for the Amazon EC2volumeType attribute.
type AttributeValue struct {
	_ struct{} `type:"structure"`

	// The specific value of an attributeName.
	Value *string `type:"string"`
}

// String returns the string representation
func (s AttributeValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AttributeValue) GoString() string {
	return s.String()
}

// SetValue sets the Value field's value.
func (s *AttributeValue) SetValue(v string) *AttributeValue {
	s.Value = &v
	return s
}

type DescribeServicesInput struct {
	_ struct{} `type:"structure"`

	// The format version that you want the response to be in.
	//
	// Valid values are: aws_v1
	FormatVersion *string `type:"string"`

	// The maximum number of results that you want returned in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The code for the service whose information you want to retrieve, such as
	// AmazonEC2. You can use the ServiceCode to filter the results in a GetProducts
	// call. To retrieve a list of all services, leave this blank.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServicesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeServicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *DescribeServicesInput) SetFormatVersion(v string) *DescribeServicesInput {
	s.FormatVersion = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeServicesInput) SetMaxResults(v int64) *DescribeServicesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeServicesInput) SetNextToken(v string) *DescribeServicesInput {
	s.NextToken = &v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *DescribeServicesInput) SetServiceCode(v string) *DescribeServicesInput {
	s.ServiceCode = &v
	return s
}

type DescribeServicesOutput struct {
	_ struct{} `type:"structure"`

	// The format version of the response. For example, aws_v1.
	FormatVersion *string `type:"string"`

	// The pagination token for the next set of retreivable results.
	NextToken *string `type:"string"`

	// The service metadata for the service or services in the response.
	Services []*Service `type:"list"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesOutput) GoString() string {
	return s.String()
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *DescribeServicesOutput) SetFormatVersion(v string) *DescribeServicesOutput {
	s.FormatVersion = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeServicesOutput) SetNextToken(v string) *DescribeServicesOutput {
	s.NextToken = &v
	return s
}

// SetServices sets the Services field's value.
func (s *DescribeServicesOutput) SetServices(v []*Service) *DescribeServicesOutput {
	s.Services = v
	return s
}

// The constraints that you want all returned products to match.
type Filter struct {
	_ struct{} `type:"structure"`

	// The product metadata field that you want to filter on. You can filter by
	// just the service code to see all products for a specific service, filter
	// by just the attribute name to see a specific attribute for multiple services,
	// or use both a service code and an attribute name to retrieve only products
	// that match both fields.
	//
	// Valid values include: ServiceCode, and all attribute names
	//
	// For example, you can filter by the AmazonEC2 service code and the volumeType
	// attribute name to get the prices for only Amazon EC2 volumes.
	//
	// Field is a required field
	Field *string `type:"string" required:"true"`

	// The type of filter that you want to use.
	//
	// Valid values are: TERM_MATCH. TERM_MATCH returns only products that match
	// both the given filter field and the given value.
	//
	// Type is a required field
	Type *string `type:"string" required:"true" enum:"FilterType"`

	// The service code or attribute value that you want to filter by. If you are
	// filtering by service code this is the actual service code, such as AmazonEC2.
	// If you are filtering by attribute name, this is the attribute value that
	// you want the returned products to match, such as a Provisioned IOPS volume.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Filter) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Filter"}
	if s.Field == nil {
		invalidParams.Add(request.NewErrParamRequired("Field"))
	}
	if s.Type == nil {
		invalidParams.Add(request.NewErrParamRequired("Type"))
	}
	if s.Value == nil {
		invalidParams.Add(request.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetField sets the Field field's value.
func (s *Filter) SetField(v string) *Filter {
	s.Field = &v
	return s
}

// SetType sets the Type field's value.
func (s *Filter) SetType(v string) *Filter {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *Filter) SetValue(v string) *Filter {
	s.Value = &v
	return s
}

type GetAttributeValuesInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute that you want to retrieve the values for, such
	// as volumeType.
	//
	// AttributeName is a required field
	AttributeName *string `type:"string" required:"true"`

	// The maximum number of results to return in response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The service code for the service whose attributes you want to retrieve. For
	// example, if you want the retrieve an EC2 attribute, use AmazonEC2.
	//
	// ServiceCode is a required field
	ServiceCode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAttributeValuesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAttributeValuesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAttributeValuesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetAttributeValuesInput"}
	if s.AttributeName == nil {
		invalidParams.Add(request.NewErrParamRequired("AttributeName"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ServiceCode == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributeName sets the AttributeName field's value.
func (s *GetAttributeValuesInput) SetAttributeName(v string) *GetAttributeValuesInput {
	s.AttributeName = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetAttributeValuesInput) SetMaxResults(v int64) *GetAttributeValuesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetAttributeValuesInput) SetNextToken(v string) *GetAttributeValuesInput {
	s.NextToken = &v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *GetAttributeValuesInput) SetServiceCode(v string) *GetAttributeValuesInput {
	s.ServiceCode = &v
	return s
}

type GetAttributeValuesOutput struct {
	_ struct{} `type:"structure"`

	// The list of values for an attribute. For example, Throughput Optimized HDD
	// and Provisioned IOPS are two available values for the AmazonEC2volumeType.
	AttributeValues []*AttributeValue `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetAttributeValuesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAttributeValuesOutput) GoString() string {
	return s.String()
}

// SetAttributeValues sets the AttributeValues field's value.
func (s *GetAttributeValuesOutput) SetAttributeValues(v []*AttributeValue) *GetAttributeValuesOutput {
	s.AttributeValues = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetAttributeValuesOutput) SetNextToken(v string) *GetAttributeValuesOutput {
	s.NextToken = &v
	return s
}

type GetProductsInput struct {
	_ struct{} `type:"structure"`

	// The list of filters that limit the returned products. only products that
	// match all filters are returned.
	Filters []*Filter `type:"list"`

	// The format version that you want the response to be in.
	//
	// Valid values are: aws_v1
	FormatVersion *string `type:"string"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The code for the service whose products you want to retrieve.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s GetProductsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProductsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProductsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetProductsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilters sets the Filters field's value.
func (s *GetProductsInput) SetFilters(v []*Filter) *GetProductsInput {
	s.Filters = v
	return s
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *GetProductsInput) SetFormatVersion(v string) *GetProductsInput {
	s.FormatVersion = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetProductsInput) SetMaxResults(v int64) *GetProductsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetProductsInput) SetNextToken(v string) *GetProductsInput {
	s.NextToken = &v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *GetProductsInput) SetServiceCode(v string) *GetProductsInput {
	s.ServiceCode = &v
	return s
}

type GetProductsOutput struct {
	_ struct{} `type:"structure"`

	// The format version of the response. For example, aws_v1.
	FormatVersion *string `type:"string"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`

	// The list of products that match your filters. The list contains both the
	// product metadata and the price information.
	PriceList []aws.JSONValue `type:"list"`
}

// String returns the string representation
func (s GetProductsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProductsOutput) GoString() string {
	return s.String()
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *GetProductsOutput) SetFormatVersion(v string) *GetProductsOutput {
	s.FormatVersion = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetProductsOutput) SetNextToken(v string) *GetProductsOutput {
	s.NextToken = &v
	return s
}

// SetPriceList sets the PriceList field's value.
func (s *GetProductsOutput) SetPriceList(v []aws.JSONValue) *GetProductsOutput {
	s.PriceList = v
	return s
}

// The metadata for a service, such as the service code and available attribute
// names.
type Service struct {
	_ struct{} `type:"structure"`

	// The attributes that are available for this service.
	AttributeNames []*string `type:"list"`

	// The code for the AWS service.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s Service) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Service) GoString() string {
	return s.String()
}

// SetAttributeNames sets the AttributeNames field's value.
func (s *Service) SetAttributeNames(v []*string) *Service {
	s.AttributeNames = v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *Service) SetServiceCode(v string) *Service {
	s.ServiceCode = &v
	return s
}

const (
	// FilterTypeTermMatch is a FilterType enum value
	FilterTypeTermMatch = "TERM_MATCH"
)
