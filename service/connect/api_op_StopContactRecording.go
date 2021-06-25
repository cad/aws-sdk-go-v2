// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops recording a call when a contact is being recorded. StopContactRecording is
// a one-time action. If you use StopContactRecording to stop recording an ongoing
// call, you can't use StartContactRecording to restart it. For scenarios where the
// recording has started and you want to suspend it for sensitive information (for
// example, to collect a credit card number), and then restart it, use
// SuspendContactRecording and ResumeContactRecording. Only voice recordings are
// supported at this time.
func (c *Client) StopContactRecording(ctx context.Context, params *StopContactRecordingInput, optFns ...func(*Options)) (*StopContactRecordingOutput, error) {
	if params == nil {
		params = &StopContactRecordingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopContactRecording", params, optFns, c.addOperationStopContactRecordingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopContactRecordingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopContactRecordingInput struct {

	// The identifier of the contact.
	//
	// This member is required.
	ContactId *string

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// This member is required.
	InitialContactId *string

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string
}

type StopContactRecordingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationStopContactRecordingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopContactRecording{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopContactRecording{}, middleware.After)
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
	if err = addOpStopContactRecordingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopContactRecording(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopContactRecording(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StopContactRecording",
	}
}
