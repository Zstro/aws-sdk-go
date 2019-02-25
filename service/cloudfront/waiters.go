// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"time"

	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/request"
)

// WaitUntilDistributionDeployed uses the CloudFront API operation
// GetDistribution to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudFront) WaitUntilDistributionDeployed(input *GetDistributionInput) error {
	return c.WaitUntilDistributionDeployedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDistributionDeployedWithContext is an extended version of WaitUntilDistributionDeployed.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFront) WaitUntilDistributionDeployedWithContext(ctx aws.Context, input *GetDistributionInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDistributionDeployed",
		MaxAttempts: 25,
		Delay:       request.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Distribution.Status",
				Expected: "Deployed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetDistributionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetDistributionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInvalidationCompleted uses the CloudFront API operation
// GetInvalidation to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudFront) WaitUntilInvalidationCompleted(input *GetInvalidationInput) error {
	return c.WaitUntilInvalidationCompletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInvalidationCompletedWithContext is an extended version of WaitUntilInvalidationCompleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFront) WaitUntilInvalidationCompletedWithContext(ctx aws.Context, input *GetInvalidationInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInvalidationCompleted",
		MaxAttempts: 30,
		Delay:       request.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Invalidation.Status",
				Expected: "Completed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetInvalidationInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetInvalidationRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingDistributionDeployed uses the CloudFront API operation
// GetStreamingDistribution to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudFront) WaitUntilStreamingDistributionDeployed(input *GetStreamingDistributionInput) error {
	return c.WaitUntilStreamingDistributionDeployedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingDistributionDeployedWithContext is an extended version of WaitUntilStreamingDistributionDeployed.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFront) WaitUntilStreamingDistributionDeployedWithContext(ctx aws.Context, input *GetStreamingDistributionInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStreamingDistributionDeployed",
		MaxAttempts: 25,
		Delay:       request.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "StreamingDistribution.Status",
				Expected: "Deployed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetStreamingDistributionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetStreamingDistributionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
