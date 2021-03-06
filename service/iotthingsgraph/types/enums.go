// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DefinitionLanguage string

// Enum values for DefinitionLanguage
const (
	DefinitionLanguageGraphql DefinitionLanguage = "GRAPHQL"
)

// Values returns all known values for DefinitionLanguage. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DefinitionLanguage) Values() []DefinitionLanguage {
	return []DefinitionLanguage{
		"GRAPHQL",
	}
}

type DeploymentTarget string

// Enum values for DeploymentTarget
const (
	DeploymentTargetGreengrass DeploymentTarget = "GREENGRASS"
	DeploymentTargetCloud      DeploymentTarget = "CLOUD"
)

// Values returns all known values for DeploymentTarget. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentTarget) Values() []DeploymentTarget {
	return []DeploymentTarget{
		"GREENGRASS",
		"CLOUD",
	}
}

type EntityFilterName string

// Enum values for EntityFilterName
const (
	EntityFilterNameName                 EntityFilterName = "NAME"
	EntityFilterNameNamespace            EntityFilterName = "NAMESPACE"
	EntityFilterNameSemantic_type_path   EntityFilterName = "SEMANTIC_TYPE_PATH"
	EntityFilterNameReferenced_entity_id EntityFilterName = "REFERENCED_ENTITY_ID"
)

// Values returns all known values for EntityFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EntityFilterName) Values() []EntityFilterName {
	return []EntityFilterName{
		"NAME",
		"NAMESPACE",
		"SEMANTIC_TYPE_PATH",
		"REFERENCED_ENTITY_ID",
	}
}

type EntityType string

// Enum values for EntityType
const (
	EntityTypeDevice       EntityType = "DEVICE"
	EntityTypeService      EntityType = "SERVICE"
	EntityTypeDevice_model EntityType = "DEVICE_MODEL"
	EntityTypeCapability   EntityType = "CAPABILITY"
	EntityTypeState        EntityType = "STATE"
	EntityTypeAction       EntityType = "ACTION"
	EntityTypeEvent        EntityType = "EVENT"
	EntityTypeProperty     EntityType = "PROPERTY"
	EntityTypeMapping      EntityType = "MAPPING"
	EntityTypeEnum         EntityType = "ENUM"
)

// Values returns all known values for EntityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EntityType) Values() []EntityType {
	return []EntityType{
		"DEVICE",
		"SERVICE",
		"DEVICE_MODEL",
		"CAPABILITY",
		"STATE",
		"ACTION",
		"EVENT",
		"PROPERTY",
		"MAPPING",
		"ENUM",
	}
}

type FlowExecutionEventType string

// Enum values for FlowExecutionEventType
const (
	FlowExecutionEventTypeExecution_started              FlowExecutionEventType = "EXECUTION_STARTED"
	FlowExecutionEventTypeExecution_failed               FlowExecutionEventType = "EXECUTION_FAILED"
	FlowExecutionEventTypeExecution_aborted              FlowExecutionEventType = "EXECUTION_ABORTED"
	FlowExecutionEventTypeExecution_succeeded            FlowExecutionEventType = "EXECUTION_SUCCEEDED"
	FlowExecutionEventTypeStep_started                   FlowExecutionEventType = "STEP_STARTED"
	FlowExecutionEventTypeStep_failed                    FlowExecutionEventType = "STEP_FAILED"
	FlowExecutionEventTypeStep_succeeded                 FlowExecutionEventType = "STEP_SUCCEEDED"
	FlowExecutionEventTypeActivity_scheduled             FlowExecutionEventType = "ACTIVITY_SCHEDULED"
	FlowExecutionEventTypeActivity_started               FlowExecutionEventType = "ACTIVITY_STARTED"
	FlowExecutionEventTypeActivity_failed                FlowExecutionEventType = "ACTIVITY_FAILED"
	FlowExecutionEventTypeActivity_succeeded             FlowExecutionEventType = "ACTIVITY_SUCCEEDED"
	FlowExecutionEventTypeStart_flow_execution_task      FlowExecutionEventType = "START_FLOW_EXECUTION_TASK"
	FlowExecutionEventTypeSchedule_next_ready_steps_task FlowExecutionEventType = "SCHEDULE_NEXT_READY_STEPS_TASK"
	FlowExecutionEventTypeThing_action_task              FlowExecutionEventType = "THING_ACTION_TASK"
	FlowExecutionEventTypeThing_action_task_failed       FlowExecutionEventType = "THING_ACTION_TASK_FAILED"
	FlowExecutionEventTypeThing_action_task_succeeded    FlowExecutionEventType = "THING_ACTION_TASK_SUCCEEDED"
	FlowExecutionEventTypeAcknowledge_task_message       FlowExecutionEventType = "ACKNOWLEDGE_TASK_MESSAGE"
)

// Values returns all known values for FlowExecutionEventType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FlowExecutionEventType) Values() []FlowExecutionEventType {
	return []FlowExecutionEventType{
		"EXECUTION_STARTED",
		"EXECUTION_FAILED",
		"EXECUTION_ABORTED",
		"EXECUTION_SUCCEEDED",
		"STEP_STARTED",
		"STEP_FAILED",
		"STEP_SUCCEEDED",
		"ACTIVITY_SCHEDULED",
		"ACTIVITY_STARTED",
		"ACTIVITY_FAILED",
		"ACTIVITY_SUCCEEDED",
		"START_FLOW_EXECUTION_TASK",
		"SCHEDULE_NEXT_READY_STEPS_TASK",
		"THING_ACTION_TASK",
		"THING_ACTION_TASK_FAILED",
		"THING_ACTION_TASK_SUCCEEDED",
		"ACKNOWLEDGE_TASK_MESSAGE",
	}
}

type FlowExecutionStatus string

// Enum values for FlowExecutionStatus
const (
	FlowExecutionStatusRunning   FlowExecutionStatus = "RUNNING"
	FlowExecutionStatusAborted   FlowExecutionStatus = "ABORTED"
	FlowExecutionStatusSucceeded FlowExecutionStatus = "SUCCEEDED"
	FlowExecutionStatusFailed    FlowExecutionStatus = "FAILED"
)

// Values returns all known values for FlowExecutionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FlowExecutionStatus) Values() []FlowExecutionStatus {
	return []FlowExecutionStatus{
		"RUNNING",
		"ABORTED",
		"SUCCEEDED",
		"FAILED",
	}
}

type FlowTemplateFilterName string

// Enum values for FlowTemplateFilterName
const (
	FlowTemplateFilterNameDevice_model_id FlowTemplateFilterName = "DEVICE_MODEL_ID"
)

// Values returns all known values for FlowTemplateFilterName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FlowTemplateFilterName) Values() []FlowTemplateFilterName {
	return []FlowTemplateFilterName{
		"DEVICE_MODEL_ID",
	}
}

type NamespaceDeletionStatus string

// Enum values for NamespaceDeletionStatus
const (
	NamespaceDeletionStatusIn_progress NamespaceDeletionStatus = "IN_PROGRESS"
	NamespaceDeletionStatusSucceeded   NamespaceDeletionStatus = "SUCCEEDED"
	NamespaceDeletionStatusFailed      NamespaceDeletionStatus = "FAILED"
)

// Values returns all known values for NamespaceDeletionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NamespaceDeletionStatus) Values() []NamespaceDeletionStatus {
	return []NamespaceDeletionStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type NamespaceDeletionStatusErrorCodes string

// Enum values for NamespaceDeletionStatusErrorCodes
const (
	NamespaceDeletionStatusErrorCodesValidation_failed NamespaceDeletionStatusErrorCodes = "VALIDATION_FAILED"
)

// Values returns all known values for NamespaceDeletionStatusErrorCodes. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (NamespaceDeletionStatusErrorCodes) Values() []NamespaceDeletionStatusErrorCodes {
	return []NamespaceDeletionStatusErrorCodes{
		"VALIDATION_FAILED",
	}
}

type SystemInstanceDeploymentStatus string

// Enum values for SystemInstanceDeploymentStatus
const (
	SystemInstanceDeploymentStatusNot_deployed         SystemInstanceDeploymentStatus = "NOT_DEPLOYED"
	SystemInstanceDeploymentStatusBootstrap            SystemInstanceDeploymentStatus = "BOOTSTRAP"
	SystemInstanceDeploymentStatusDeploy_in_progress   SystemInstanceDeploymentStatus = "DEPLOY_IN_PROGRESS"
	SystemInstanceDeploymentStatusDeployed_in_target   SystemInstanceDeploymentStatus = "DEPLOYED_IN_TARGET"
	SystemInstanceDeploymentStatusUndeploy_in_progress SystemInstanceDeploymentStatus = "UNDEPLOY_IN_PROGRESS"
	SystemInstanceDeploymentStatusFailed               SystemInstanceDeploymentStatus = "FAILED"
	SystemInstanceDeploymentStatusPending_delete       SystemInstanceDeploymentStatus = "PENDING_DELETE"
	SystemInstanceDeploymentStatusDeleted_in_target    SystemInstanceDeploymentStatus = "DELETED_IN_TARGET"
)

// Values returns all known values for SystemInstanceDeploymentStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SystemInstanceDeploymentStatus) Values() []SystemInstanceDeploymentStatus {
	return []SystemInstanceDeploymentStatus{
		"NOT_DEPLOYED",
		"BOOTSTRAP",
		"DEPLOY_IN_PROGRESS",
		"DEPLOYED_IN_TARGET",
		"UNDEPLOY_IN_PROGRESS",
		"FAILED",
		"PENDING_DELETE",
		"DELETED_IN_TARGET",
	}
}

type SystemInstanceFilterName string

// Enum values for SystemInstanceFilterName
const (
	SystemInstanceFilterNameSystem_template_id    SystemInstanceFilterName = "SYSTEM_TEMPLATE_ID"
	SystemInstanceFilterNameStatus                SystemInstanceFilterName = "STATUS"
	SystemInstanceFilterNameGreengrass_group_name SystemInstanceFilterName = "GREENGRASS_GROUP_NAME"
)

// Values returns all known values for SystemInstanceFilterName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SystemInstanceFilterName) Values() []SystemInstanceFilterName {
	return []SystemInstanceFilterName{
		"SYSTEM_TEMPLATE_ID",
		"STATUS",
		"GREENGRASS_GROUP_NAME",
	}
}

type SystemTemplateFilterName string

// Enum values for SystemTemplateFilterName
const (
	SystemTemplateFilterNameFlow_template_id SystemTemplateFilterName = "FLOW_TEMPLATE_ID"
)

// Values returns all known values for SystemTemplateFilterName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SystemTemplateFilterName) Values() []SystemTemplateFilterName {
	return []SystemTemplateFilterName{
		"FLOW_TEMPLATE_ID",
	}
}

type UploadStatus string

// Enum values for UploadStatus
const (
	UploadStatusIn_progress UploadStatus = "IN_PROGRESS"
	UploadStatusSucceeded   UploadStatus = "SUCCEEDED"
	UploadStatusFailed      UploadStatus = "FAILED"
)

// Values returns all known values for UploadStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UploadStatus) Values() []UploadStatus {
	return []UploadStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}
