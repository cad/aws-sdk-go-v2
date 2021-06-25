// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimev2

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteSession struct {
}

func (*validateOpDeleteSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSession struct {
}

func (*validateOpGetSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutSession struct {
}

func (*validateOpPutSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRecognizeText struct {
}

func (*validateOpRecognizeText) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRecognizeText) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RecognizeTextInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRecognizeTextInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRecognizeUtterance struct {
}

func (*validateOpRecognizeUtterance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRecognizeUtterance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RecognizeUtteranceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRecognizeUtteranceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSession{}, middleware.After)
}

func addOpGetSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSession{}, middleware.After)
}

func addOpPutSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutSession{}, middleware.After)
}

func addOpRecognizeTextValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRecognizeText{}, middleware.After)
}

func addOpRecognizeUtteranceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRecognizeUtterance{}, middleware.After)
}

func validateActiveContext(v *types.ActiveContext) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActiveContext"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.TimeToLive == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimeToLive"))
	} else if v.TimeToLive != nil {
		if err := validateActiveContextTimeToLive(v.TimeToLive); err != nil {
			invalidParams.AddNested("TimeToLive", err.(smithy.InvalidParamsError))
		}
	}
	if v.ContextAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContextAttributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateActiveContextsList(v []types.ActiveContext) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActiveContextsList"}
	for i := range v {
		if err := validateActiveContext(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateActiveContextTimeToLive(v *types.ActiveContextTimeToLive) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActiveContextTimeToLive"}
	if v.TimeToLiveInSeconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimeToLiveInSeconds"))
	}
	if v.TurnsToLive == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TurnsToLive"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateButton(v *types.Button) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Button"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateButtonsList(v []types.Button) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ButtonsList"}
	for i := range v {
		if err := validateButton(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDialogAction(v *types.DialogAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DialogAction"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImageResponseCard(v *types.ImageResponseCard) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImageResponseCard"}
	if v.Title == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Title"))
	}
	if v.Buttons != nil {
		if err := validateButtonsList(v.Buttons); err != nil {
			invalidParams.AddNested("Buttons", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIntent(v *types.Intent) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Intent"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Slots != nil {
		if err := validateSlots(v.Slots); err != nil {
			invalidParams.AddNested("Slots", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessage(v *types.Message) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Message"}
	if len(v.ContentType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.ImageResponseCard != nil {
		if err := validateImageResponseCard(v.ImageResponseCard); err != nil {
			invalidParams.AddNested("ImageResponseCard", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessages(v []types.Message) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Messages"}
	for i := range v {
		if err := validateMessage(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSessionState(v *types.SessionState) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SessionState"}
	if v.DialogAction != nil {
		if err := validateDialogAction(v.DialogAction); err != nil {
			invalidParams.AddNested("DialogAction", err.(smithy.InvalidParamsError))
		}
	}
	if v.Intent != nil {
		if err := validateIntent(v.Intent); err != nil {
			invalidParams.AddNested("Intent", err.(smithy.InvalidParamsError))
		}
	}
	if v.ActiveContexts != nil {
		if err := validateActiveContextsList(v.ActiveContexts); err != nil {
			invalidParams.AddNested("ActiveContexts", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSlot(v *types.Slot) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Slot"}
	if v.Value != nil {
		if err := validateValue(v.Value); err != nil {
			invalidParams.AddNested("Value", err.(smithy.InvalidParamsError))
		}
	}
	if v.Values != nil {
		if err := validateValues(v.Values); err != nil {
			invalidParams.AddNested("Values", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSlots(v map[string]types.Slot) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Slots"}
	for key := range v {
		value := v[key]
		if err := validateSlot(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateValue(v *types.Value) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Value"}
	if v.InterpretedValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InterpretedValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateValues(v []types.Slot) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Values"}
	for i := range v {
		if err := validateSlot(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSessionInput(v *DeleteSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSessionInput"}
	if v.BotId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotId"))
	}
	if v.BotAliasId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAliasId"))
	}
	if v.LocaleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LocaleId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSessionInput(v *GetSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSessionInput"}
	if v.BotId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotId"))
	}
	if v.BotAliasId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAliasId"))
	}
	if v.LocaleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LocaleId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutSessionInput(v *PutSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutSessionInput"}
	if v.BotId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotId"))
	}
	if v.BotAliasId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAliasId"))
	}
	if v.LocaleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LocaleId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.Messages != nil {
		if err := validateMessages(v.Messages); err != nil {
			invalidParams.AddNested("Messages", err.(smithy.InvalidParamsError))
		}
	}
	if v.SessionState == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionState"))
	} else if v.SessionState != nil {
		if err := validateSessionState(v.SessionState); err != nil {
			invalidParams.AddNested("SessionState", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRecognizeTextInput(v *RecognizeTextInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RecognizeTextInput"}
	if v.BotId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotId"))
	}
	if v.BotAliasId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAliasId"))
	}
	if v.LocaleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LocaleId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if v.SessionState != nil {
		if err := validateSessionState(v.SessionState); err != nil {
			invalidParams.AddNested("SessionState", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRecognizeUtteranceInput(v *RecognizeUtteranceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RecognizeUtteranceInput"}
	if v.BotId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotId"))
	}
	if v.BotAliasId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAliasId"))
	}
	if v.LocaleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LocaleId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.RequestContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RequestContentType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
