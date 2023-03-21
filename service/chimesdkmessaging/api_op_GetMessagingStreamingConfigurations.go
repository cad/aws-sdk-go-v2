// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the data streaming configuration for an AppInstance. For more
// information, see Streaming messaging data
// (https://docs.aws.amazon.com/chime-sdk/latest/dg/streaming-export.html) in the
// Amazon Chime SDK Developer Guide.
func (c *Client) GetMessagingStreamingConfigurations(ctx context.Context, params *GetMessagingStreamingConfigurationsInput, optFns ...func(*Options)) (*GetMessagingStreamingConfigurationsOutput, error) {
	if params == nil {
		params = &GetMessagingStreamingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMessagingStreamingConfigurations", params, optFns, c.addOperationGetMessagingStreamingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMessagingStreamingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMessagingStreamingConfigurationsInput struct {

	// The ARN of the streaming configurations.
	//
	// This member is required.
	AppInstanceArn *string

	noSmithyDocumentSerde
}

type GetMessagingStreamingConfigurationsOutput struct {

	// The streaming settings.
	StreamingConfigurations []types.StreamingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMessagingStreamingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMessagingStreamingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMessagingStreamingConfigurations{}, middleware.After)
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
	if err = addOpGetMessagingStreamingConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMessagingStreamingConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMessagingStreamingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetMessagingStreamingConfigurations",
	}
}
