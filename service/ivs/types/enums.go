// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChannelLatencyMode string

// Enum values for ChannelLatencyMode
const (
	ChannelLatencyModeNormalLatency ChannelLatencyMode = "NORMAL"
	ChannelLatencyModeLowLatency    ChannelLatencyMode = "LOW"
)

// Values returns all known values for ChannelLatencyMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelLatencyMode) Values() []ChannelLatencyMode {
	return []ChannelLatencyMode{
		"NORMAL",
		"LOW",
	}
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypeBasicChannelType      ChannelType = "BASIC"
	ChannelTypeStandardChannelType   ChannelType = "STANDARD"
	ChannelTypeAdvancedSDChannelType ChannelType = "ADVANCED_SD"
	ChannelTypeAdvancedHDChannelType ChannelType = "ADVANCED_HD"
)

// Values returns all known values for ChannelType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChannelType) Values() []ChannelType {
	return []ChannelType{
		"BASIC",
		"STANDARD",
		"ADVANCED_SD",
		"ADVANCED_HD",
	}
}

type RecordingConfigurationState string

// Enum values for RecordingConfigurationState
const (
	RecordingConfigurationStateCreating     RecordingConfigurationState = "CREATING"
	RecordingConfigurationStateCreateFailed RecordingConfigurationState = "CREATE_FAILED"
	RecordingConfigurationStateActive       RecordingConfigurationState = "ACTIVE"
)

// Values returns all known values for RecordingConfigurationState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecordingConfigurationState) Values() []RecordingConfigurationState {
	return []RecordingConfigurationState{
		"CREATING",
		"CREATE_FAILED",
		"ACTIVE",
	}
}

type RecordingMode string

// Enum values for RecordingMode
const (
	RecordingModeDisabled RecordingMode = "DISABLED"
	RecordingModeInterval RecordingMode = "INTERVAL"
)

// Values returns all known values for RecordingMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecordingMode) Values() []RecordingMode {
	return []RecordingMode{
		"DISABLED",
		"INTERVAL",
	}
}

type StreamHealth string

// Enum values for StreamHealth
const (
	StreamHealthStreamHealthy StreamHealth = "HEALTHY"
	StreamHealthStarving      StreamHealth = "STARVING"
	StreamHealthUnknown       StreamHealth = "UNKNOWN"
)

// Values returns all known values for StreamHealth. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StreamHealth) Values() []StreamHealth {
	return []StreamHealth{
		"HEALTHY",
		"STARVING",
		"UNKNOWN",
	}
}

type StreamState string

// Enum values for StreamState
const (
	StreamStateStreamLive    StreamState = "LIVE"
	StreamStateStreamOffline StreamState = "OFFLINE"
)

// Values returns all known values for StreamState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StreamState) Values() []StreamState {
	return []StreamState{
		"LIVE",
		"OFFLINE",
	}
}

type TranscodePreset string

// Enum values for TranscodePreset
const (
	TranscodePresetHigherBandwidthTranscodePreset      TranscodePreset = "HIGHER_BANDWIDTH_DELIVERY"
	TranscodePresetConstrainedBandwidthTranscodePreset TranscodePreset = "CONSTRAINED_BANDWIDTH_DELIVERY"
)

// Values returns all known values for TranscodePreset. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscodePreset) Values() []TranscodePreset {
	return []TranscodePreset{
		"HIGHER_BANDWIDTH_DELIVERY",
		"CONSTRAINED_BANDWIDTH_DELIVERY",
	}
}
