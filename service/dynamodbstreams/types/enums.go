// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type KeyType string

// Enum values for KeyType
const (
	KeyTypeHash  KeyType = "HASH"
	KeyTypeRange KeyType = "RANGE"
)

// Values returns all known values for KeyType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (KeyType) Values() []KeyType {
	return []KeyType{
		"HASH",
		"RANGE",
	}
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeInsert OperationType = "INSERT"
	OperationTypeModify OperationType = "MODIFY"
	OperationTypeRemove OperationType = "REMOVE"
)

// Values returns all known values for OperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationType) Values() []OperationType {
	return []OperationType{
		"INSERT",
		"MODIFY",
		"REMOVE",
	}
}

type ShardIteratorType string

// Enum values for ShardIteratorType
const (
	ShardIteratorTypeTrim_horizon          ShardIteratorType = "TRIM_HORIZON"
	ShardIteratorTypeLatest                ShardIteratorType = "LATEST"
	ShardIteratorTypeAt_sequence_number    ShardIteratorType = "AT_SEQUENCE_NUMBER"
	ShardIteratorTypeAfter_sequence_number ShardIteratorType = "AFTER_SEQUENCE_NUMBER"
)

// Values returns all known values for ShardIteratorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShardIteratorType) Values() []ShardIteratorType {
	return []ShardIteratorType{
		"TRIM_HORIZON",
		"LATEST",
		"AT_SEQUENCE_NUMBER",
		"AFTER_SEQUENCE_NUMBER",
	}
}

type StreamStatus string

// Enum values for StreamStatus
const (
	StreamStatusEnabling  StreamStatus = "ENABLING"
	StreamStatusEnabled   StreamStatus = "ENABLED"
	StreamStatusDisabling StreamStatus = "DISABLING"
	StreamStatusDisabled  StreamStatus = "DISABLED"
)

// Values returns all known values for StreamStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StreamStatus) Values() []StreamStatus {
	return []StreamStatus{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
	}
}

type StreamViewType string

// Enum values for StreamViewType
const (
	StreamViewTypeNew_image          StreamViewType = "NEW_IMAGE"
	StreamViewTypeOld_image          StreamViewType = "OLD_IMAGE"
	StreamViewTypeNew_and_old_images StreamViewType = "NEW_AND_OLD_IMAGES"
	StreamViewTypeKeys_only          StreamViewType = "KEYS_ONLY"
)

// Values returns all known values for StreamViewType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StreamViewType) Values() []StreamViewType {
	return []StreamViewType{
		"NEW_IMAGE",
		"OLD_IMAGE",
		"NEW_AND_OLD_IMAGES",
		"KEYS_ONLY",
	}
}
