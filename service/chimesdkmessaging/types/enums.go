// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AllowNotifications string

// Enum values for AllowNotifications
const (
	AllowNotificationsAll      AllowNotifications = "ALL"
	AllowNotificationsNone     AllowNotifications = "NONE"
	AllowNotificationsFiltered AllowNotifications = "FILTERED"
)

// Values returns all known values for AllowNotifications. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AllowNotifications) Values() []AllowNotifications {
	return []AllowNotifications{
		"ALL",
		"NONE",
		"FILTERED",
	}
}

type ChannelMembershipType string

// Enum values for ChannelMembershipType
const (
	ChannelMembershipTypeDefault ChannelMembershipType = "DEFAULT"
	ChannelMembershipTypeHidden  ChannelMembershipType = "HIDDEN"
)

// Values returns all known values for ChannelMembershipType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelMembershipType) Values() []ChannelMembershipType {
	return []ChannelMembershipType{
		"DEFAULT",
		"HIDDEN",
	}
}

type ChannelMessagePersistenceType string

// Enum values for ChannelMessagePersistenceType
const (
	ChannelMessagePersistenceTypePersistent    ChannelMessagePersistenceType = "PERSISTENT"
	ChannelMessagePersistenceTypeNonPersistent ChannelMessagePersistenceType = "NON_PERSISTENT"
)

// Values returns all known values for ChannelMessagePersistenceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ChannelMessagePersistenceType) Values() []ChannelMessagePersistenceType {
	return []ChannelMessagePersistenceType{
		"PERSISTENT",
		"NON_PERSISTENT",
	}
}

type ChannelMessageStatus string

// Enum values for ChannelMessageStatus
const (
	ChannelMessageStatusSent    ChannelMessageStatus = "SENT"
	ChannelMessageStatusPending ChannelMessageStatus = "PENDING"
	ChannelMessageStatusFailed  ChannelMessageStatus = "FAILED"
	ChannelMessageStatusDenied  ChannelMessageStatus = "DENIED"
)

// Values returns all known values for ChannelMessageStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelMessageStatus) Values() []ChannelMessageStatus {
	return []ChannelMessageStatus{
		"SENT",
		"PENDING",
		"FAILED",
		"DENIED",
	}
}

type ChannelMessageType string

// Enum values for ChannelMessageType
const (
	ChannelMessageTypeStandard ChannelMessageType = "STANDARD"
	ChannelMessageTypeControl  ChannelMessageType = "CONTROL"
)

// Values returns all known values for ChannelMessageType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelMessageType) Values() []ChannelMessageType {
	return []ChannelMessageType{
		"STANDARD",
		"CONTROL",
	}
}

type ChannelMode string

// Enum values for ChannelMode
const (
	ChannelModeUnrestricted ChannelMode = "UNRESTRICTED"
	ChannelModeRestricted   ChannelMode = "RESTRICTED"
)

// Values returns all known values for ChannelMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChannelMode) Values() []ChannelMode {
	return []ChannelMode{
		"UNRESTRICTED",
		"RESTRICTED",
	}
}

type ChannelPrivacy string

// Enum values for ChannelPrivacy
const (
	ChannelPrivacyPublic  ChannelPrivacy = "PUBLIC"
	ChannelPrivacyPrivate ChannelPrivacy = "PRIVATE"
)

// Values returns all known values for ChannelPrivacy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelPrivacy) Values() []ChannelPrivacy {
	return []ChannelPrivacy{
		"PUBLIC",
		"PRIVATE",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadRequest                           ErrorCode = "BadRequest"
	ErrorCodeConflict                             ErrorCode = "Conflict"
	ErrorCodeForbidden                            ErrorCode = "Forbidden"
	ErrorCodeNotFound                             ErrorCode = "NotFound"
	ErrorCodePreconditionFailed                   ErrorCode = "PreconditionFailed"
	ErrorCodeResourceLimitExceeded                ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServiceFailure                       ErrorCode = "ServiceFailure"
	ErrorCodeAccessDenied                         ErrorCode = "AccessDenied"
	ErrorCodeServiceUnavailable                   ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled                            ErrorCode = "Throttled"
	ErrorCodeThrottling                           ErrorCode = "Throttling"
	ErrorCodeUnauthorized                         ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable                        ErrorCode = "Unprocessable"
	ErrorCodeVoiceConnectorGroupAssociationsExist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhoneNumberAssociationsExist         ErrorCode = "PhoneNumberAssociationsExist"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
	}
}

type FallbackAction string

// Enum values for FallbackAction
const (
	FallbackActionContinue FallbackAction = "CONTINUE"
	FallbackActionAbort    FallbackAction = "ABORT"
)

// Values returns all known values for FallbackAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FallbackAction) Values() []FallbackAction {
	return []FallbackAction{
		"CONTINUE",
		"ABORT",
	}
}

type InvocationType string

// Enum values for InvocationType
const (
	InvocationTypeAsync InvocationType = "ASYNC"
)

// Values returns all known values for InvocationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InvocationType) Values() []InvocationType {
	return []InvocationType{
		"ASYNC",
	}
}

type MessagingDataType string

// Enum values for MessagingDataType
const (
	MessagingDataTypeChannel        MessagingDataType = "Channel"
	MessagingDataTypeChannelMessage MessagingDataType = "ChannelMessage"
)

// Values returns all known values for MessagingDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MessagingDataType) Values() []MessagingDataType {
	return []MessagingDataType{
		"Channel",
		"ChannelMessage",
	}
}

type PushNotificationType string

// Enum values for PushNotificationType
const (
	PushNotificationTypeDefault PushNotificationType = "DEFAULT"
	PushNotificationTypeVoip    PushNotificationType = "VOIP"
)

// Values returns all known values for PushNotificationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PushNotificationType) Values() []PushNotificationType {
	return []PushNotificationType{
		"DEFAULT",
		"VOIP",
	}
}

type SearchFieldKey string

// Enum values for SearchFieldKey
const (
	SearchFieldKeyMembers SearchFieldKey = "MEMBERS"
)

// Values returns all known values for SearchFieldKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SearchFieldKey) Values() []SearchFieldKey {
	return []SearchFieldKey{
		"MEMBERS",
	}
}

type SearchFieldOperator string

// Enum values for SearchFieldOperator
const (
	SearchFieldOperatorEquals   SearchFieldOperator = "EQUALS"
	SearchFieldOperatorIncludes SearchFieldOperator = "INCLUDES"
)

// Values returns all known values for SearchFieldOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SearchFieldOperator) Values() []SearchFieldOperator {
	return []SearchFieldOperator{
		"EQUALS",
		"INCLUDES",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}
