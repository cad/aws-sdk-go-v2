// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Get detail data for a service instance. A service instance is an instantiation
// of service template, which is running in a specific environment.
func (c *Client) GetServiceInstance(ctx context.Context, params *GetServiceInstanceInput, optFns ...func(*Options)) (*GetServiceInstanceOutput, error) {
	if params == nil {
		params = &GetServiceInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceInstance", params, optFns, c.addOperationGetServiceInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceInstanceInput struct {

	// The name of a service instance that you want to get the detail data for.
	//
	// This member is required.
	Name *string

	// The name of the service that the service instance belongs to.
	//
	// This member is required.
	ServiceName *string
}

type GetServiceInstanceOutput struct {

	// The service instance detail data that's returned by AWS Proton.
	//
	// This member is required.
	ServiceInstance *types.ServiceInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetServiceInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetServiceInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetServiceInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetServiceInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetServiceInstanceAPIClient is a client that implements the GetServiceInstance
// operation.
type GetServiceInstanceAPIClient interface {
	GetServiceInstance(context.Context, *GetServiceInstanceInput, ...func(*Options)) (*GetServiceInstanceOutput, error)
}

var _ GetServiceInstanceAPIClient = (*Client)(nil)

// ServiceInstanceDeployedWaiterOptions are waiter options for
// ServiceInstanceDeployedWaiter
type ServiceInstanceDeployedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ServiceInstanceDeployedWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ServiceInstanceDeployedWaiter will use default max delay of 4999
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetServiceInstanceInput, *GetServiceInstanceOutput, error) (bool, error)
}

// ServiceInstanceDeployedWaiter defines the waiters for ServiceInstanceDeployed
type ServiceInstanceDeployedWaiter struct {
	client GetServiceInstanceAPIClient

	options ServiceInstanceDeployedWaiterOptions
}

// NewServiceInstanceDeployedWaiter constructs a ServiceInstanceDeployedWaiter.
func NewServiceInstanceDeployedWaiter(client GetServiceInstanceAPIClient, optFns ...func(*ServiceInstanceDeployedWaiterOptions)) *ServiceInstanceDeployedWaiter {
	options := ServiceInstanceDeployedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 4999 * time.Second
	options.Retryable = serviceInstanceDeployedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ServiceInstanceDeployedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ServiceInstanceDeployed waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ServiceInstanceDeployedWaiter) Wait(ctx context.Context, params *GetServiceInstanceInput, maxWaitDur time.Duration, optFns ...func(*ServiceInstanceDeployedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 4999 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetServiceInstance(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ServiceInstanceDeployed waiter")
}

func serviceInstanceDeployedStateRetryable(ctx context.Context, input *GetServiceInstanceInput, output *GetServiceInstanceOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("serviceInstance.deploymentStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUCCEEDED"
		value, ok := pathValue.(types.DeploymentStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.DeploymentStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("serviceInstance.deploymentStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.DeploymentStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.DeploymentStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}
