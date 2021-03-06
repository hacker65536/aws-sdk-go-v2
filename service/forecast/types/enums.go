// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttributeType string

// Enum values for AttributeType
const (
	AttributeTypeString    AttributeType = "string"
	AttributeTypeInteger   AttributeType = "integer"
	AttributeTypeFloat     AttributeType = "float"
	AttributeTypeTimestamp AttributeType = "timestamp"
)

// Values returns all known values for AttributeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttributeType) Values() []AttributeType {
	return []AttributeType{
		"string",
		"integer",
		"float",
		"timestamp",
	}
}

type DatasetType string

// Enum values for DatasetType
const (
	DatasetTypeTarget_time_series  DatasetType = "TARGET_TIME_SERIES"
	DatasetTypeRelated_time_series DatasetType = "RELATED_TIME_SERIES"
	DatasetTypeItem_metadata       DatasetType = "ITEM_METADATA"
)

// Values returns all known values for DatasetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DatasetType) Values() []DatasetType {
	return []DatasetType{
		"TARGET_TIME_SERIES",
		"RELATED_TIME_SERIES",
		"ITEM_METADATA",
	}
}

type Domain string

// Enum values for Domain
const (
	DomainRetail             Domain = "RETAIL"
	DomainCustom             Domain = "CUSTOM"
	DomainInventory_planning Domain = "INVENTORY_PLANNING"
	DomainEc2_capacity       Domain = "EC2_CAPACITY"
	DomainWork_force         Domain = "WORK_FORCE"
	DomainWeb_traffic        Domain = "WEB_TRAFFIC"
	DomainMetrics            Domain = "METRICS"
)

// Values returns all known values for Domain. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Domain) Values() []Domain {
	return []Domain{
		"RETAIL",
		"CUSTOM",
		"INVENTORY_PLANNING",
		"EC2_CAPACITY",
		"WORK_FORCE",
		"WEB_TRAFFIC",
		"METRICS",
	}
}

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeSummary  EvaluationType = "SUMMARY"
	EvaluationTypeComputed EvaluationType = "COMPUTED"
)

// Values returns all known values for EvaluationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EvaluationType) Values() []EvaluationType {
	return []EvaluationType{
		"SUMMARY",
		"COMPUTED",
	}
}

type FeaturizationMethodName string

// Enum values for FeaturizationMethodName
const (
	FeaturizationMethodNameFilling FeaturizationMethodName = "filling"
)

// Values returns all known values for FeaturizationMethodName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FeaturizationMethodName) Values() []FeaturizationMethodName {
	return []FeaturizationMethodName{
		"filling",
	}
}

type FilterConditionString string

// Enum values for FilterConditionString
const (
	FilterConditionStringIs     FilterConditionString = "IS"
	FilterConditionStringIs_not FilterConditionString = "IS_NOT"
)

// Values returns all known values for FilterConditionString. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterConditionString) Values() []FilterConditionString {
	return []FilterConditionString{
		"IS",
		"IS_NOT",
	}
}

type ScalingType string

// Enum values for ScalingType
const (
	ScalingTypeAuto               ScalingType = "Auto"
	ScalingTypeLinear             ScalingType = "Linear"
	ScalingTypeLogarithmic        ScalingType = "Logarithmic"
	ScalingTypeReverselogarithmic ScalingType = "ReverseLogarithmic"
)

// Values returns all known values for ScalingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScalingType) Values() []ScalingType {
	return []ScalingType{
		"Auto",
		"Linear",
		"Logarithmic",
		"ReverseLogarithmic",
	}
}
