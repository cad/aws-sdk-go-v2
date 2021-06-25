// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns a RepositoryAssociation
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html)
// object that contains information about the requested repository association.
func (c *Client) DescribeRepositoryAssociation(ctx context.Context, params *DescribeRepositoryAssociationInput, optFns ...func(*Options)) (*DescribeRepositoryAssociationOutput, error) {
	if params == nil {
		params = &DescribeRepositoryAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRepositoryAssociation", params, optFns, c.addOperationDescribeRepositoryAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRepositoryAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRepositoryAssociationInput struct {

	// The Amazon Resource Name (ARN) of the RepositoryAssociation
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html)
	// object. You can retrieve this ARN by calling ListRepositoryAssociations
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_ListRepositoryAssociations.html).
	//
	// This member is required.
	AssociationArn *string
}

type DescribeRepositoryAssociationOutput struct {

	// Information about the repository association.
	RepositoryAssociation *types.RepositoryAssociation

	// An array of key-value pairs used to tag an associated repository. A tag is a
	// custom attribute label with two parts:
	//
	// * A tag key (for example, CostCenter,
	// Environment, Project, or Secret). Tag keys are case sensitive.
	//
	// * An optional
	// field known as a tag value (for example, 111122223333, Production, or a team
	// name). Omitting the tag value is the same as using an empty string. Like tag
	// keys, tag values are case sensitive.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeRepositoryAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRepositoryAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRepositoryAssociation{}, middleware.After)
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
	if err = addOpDescribeRepositoryAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRepositoryAssociation(options.Region), middleware.Before); err != nil {
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

// DescribeRepositoryAssociationAPIClient is a client that implements the
// DescribeRepositoryAssociation operation.
type DescribeRepositoryAssociationAPIClient interface {
	DescribeRepositoryAssociation(context.Context, *DescribeRepositoryAssociationInput, ...func(*Options)) (*DescribeRepositoryAssociationOutput, error)
}

var _ DescribeRepositoryAssociationAPIClient = (*Client)(nil)

// RepositoryAssociationSucceededWaiterOptions are waiter options for
// RepositoryAssociationSucceededWaiter
type RepositoryAssociationSucceededWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// RepositoryAssociationSucceededWaiter will use default minimum delay of 10
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, RepositoryAssociationSucceededWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeRepositoryAssociationInput, *DescribeRepositoryAssociationOutput, error) (bool, error)
}

// RepositoryAssociationSucceededWaiter defines the waiters for
// RepositoryAssociationSucceeded
type RepositoryAssociationSucceededWaiter struct {
	client DescribeRepositoryAssociationAPIClient

	options RepositoryAssociationSucceededWaiterOptions
}

// NewRepositoryAssociationSucceededWaiter constructs a
// RepositoryAssociationSucceededWaiter.
func NewRepositoryAssociationSucceededWaiter(client DescribeRepositoryAssociationAPIClient, optFns ...func(*RepositoryAssociationSucceededWaiterOptions)) *RepositoryAssociationSucceededWaiter {
	options := RepositoryAssociationSucceededWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = repositoryAssociationSucceededStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &RepositoryAssociationSucceededWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for RepositoryAssociationSucceeded waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *RepositoryAssociationSucceededWaiter) Wait(ctx context.Context, params *DescribeRepositoryAssociationInput, maxWaitDur time.Duration, optFns ...func(*RepositoryAssociationSucceededWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
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

		out, err := w.client.DescribeRepositoryAssociation(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for RepositoryAssociationSucceeded waiter")
}

func repositoryAssociationSucceededStateRetryable(ctx context.Context, input *DescribeRepositoryAssociationInput, output *DescribeRepositoryAssociationOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("RepositoryAssociation.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Associated"
		value, ok := pathValue.(types.RepositoryAssociationState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.RepositoryAssociationState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("RepositoryAssociation.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Associating"
		value, ok := pathValue.(types.RepositoryAssociationState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.RepositoryAssociationState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeRepositoryAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DescribeRepositoryAssociation",
	}
}
