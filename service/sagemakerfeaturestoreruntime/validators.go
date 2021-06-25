// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerfeaturestoreruntime

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchGetRecord struct {
}

func (*validateOpBatchGetRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRecord struct {
}

func (*validateOpDeleteRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRecord struct {
}

func (*validateOpGetRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRecord struct {
}

func (*validateOpPutRecord) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRecordInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetRecord{}, middleware.After)
}

func addOpDeleteRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRecord{}, middleware.After)
}

func addOpGetRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRecord{}, middleware.After)
}

func addOpPutRecordValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRecord{}, middleware.After)
}

func validateBatchGetRecordIdentifier(v *types.BatchGetRecordIdentifier) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetRecordIdentifier"}
	if v.FeatureGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureGroupName"))
	}
	if v.RecordIdentifiersValueAsString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RecordIdentifiersValueAsString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBatchGetRecordIdentifiers(v []types.BatchGetRecordIdentifier) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetRecordIdentifiers"}
	for i := range v {
		if err := validateBatchGetRecordIdentifier(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFeatureValue(v *types.FeatureValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FeatureValue"}
	if v.FeatureName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureName"))
	}
	if v.ValueAsString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ValueAsString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRecord(v []types.FeatureValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Record"}
	for i := range v {
		if err := validateFeatureValue(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetRecordInput(v *BatchGetRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetRecordInput"}
	if v.Identifiers == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifiers"))
	} else if v.Identifiers != nil {
		if err := validateBatchGetRecordIdentifiers(v.Identifiers); err != nil {
			invalidParams.AddNested("Identifiers", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRecordInput(v *DeleteRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRecordInput"}
	if v.FeatureGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureGroupName"))
	}
	if v.RecordIdentifierValueAsString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RecordIdentifierValueAsString"))
	}
	if v.EventTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRecordInput(v *GetRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRecordInput"}
	if v.FeatureGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureGroupName"))
	}
	if v.RecordIdentifierValueAsString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RecordIdentifierValueAsString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRecordInput(v *PutRecordInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRecordInput"}
	if v.FeatureGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FeatureGroupName"))
	}
	if v.Record == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Record"))
	} else if v.Record != nil {
		if err := validateRecord(v.Record); err != nil {
			invalidParams.AddNested("Record", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
