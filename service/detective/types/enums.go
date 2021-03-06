// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type MemberDisabledReason string

// Enum values for MemberDisabledReason
const (
	MemberDisabledReasonVolume_too_high MemberDisabledReason = "VOLUME_TOO_HIGH"
	MemberDisabledReasonVolume_unknown  MemberDisabledReason = "VOLUME_UNKNOWN"
)

// Values returns all known values for MemberDisabledReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MemberDisabledReason) Values() []MemberDisabledReason {
	return []MemberDisabledReason{
		"VOLUME_TOO_HIGH",
		"VOLUME_UNKNOWN",
	}
}

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusInvited                  MemberStatus = "INVITED"
	MemberStatusVerification_in_progress MemberStatus = "VERIFICATION_IN_PROGRESS"
	MemberStatusVerification_failed      MemberStatus = "VERIFICATION_FAILED"
	MemberStatusEnabled                  MemberStatus = "ENABLED"
	MemberStatusAccepted_but_disabled    MemberStatus = "ACCEPTED_BUT_DISABLED"
)

// Values returns all known values for MemberStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MemberStatus) Values() []MemberStatus {
	return []MemberStatus{
		"INVITED",
		"VERIFICATION_IN_PROGRESS",
		"VERIFICATION_FAILED",
		"ENABLED",
		"ACCEPTED_BUT_DISABLED",
	}
}
