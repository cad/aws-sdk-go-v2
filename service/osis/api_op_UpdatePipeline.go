// Code generated by smithy-go-codegen DO NOT EDIT.

package osis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/osis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an OpenSearch Ingestion pipeline. For more information, see Updating
// Amazon OpenSearch Ingestion pipelines (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/update-pipeline.html)
// .
func (c *Client) UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) {
	if params == nil {
		params = &UpdatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePipeline", params, optFns, c.addOperationUpdatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePipelineInput struct {

	// The name of the pipeline to update.
	//
	// This member is required.
	PipelineName *string

	// Key-value pairs to configure log publishing.
	LogPublishingOptions *types.LogPublishingOptions

	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs)
	MaxUnits *int32

	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	MinUnits *int32

	// The pipeline configuration in YAML format. The command accepts the pipeline
	// configuration as a string or within a .yaml file. If you provide the
	// configuration as a string, each new line must be escaped with \n .
	PipelineConfigurationBody *string

	noSmithyDocumentSerde
}

type UpdatePipelineOutput struct {

	// Container for information about the updated pipeline.
	Pipeline *types.Pipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePipeline{}, middleware.After)
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
	if err = addOpUpdatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipeline(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "osis",
		OperationName: "UpdatePipeline",
	}
}
