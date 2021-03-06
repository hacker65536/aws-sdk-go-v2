// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ArtifactCategory string

// Enum values for ArtifactCategory
const (
	ArtifactCategoryScreenshot ArtifactCategory = "SCREENSHOT"
	ArtifactCategoryFile       ArtifactCategory = "FILE"
	ArtifactCategoryLog        ArtifactCategory = "LOG"
)

// Values returns all known values for ArtifactCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactCategory) Values() []ArtifactCategory {
	return []ArtifactCategory{
		"SCREENSHOT",
		"FILE",
		"LOG",
	}
}

type ArtifactType string

// Enum values for ArtifactType
const (
	ArtifactTypeUnknown                  ArtifactType = "UNKNOWN"
	ArtifactTypeScreenshot               ArtifactType = "SCREENSHOT"
	ArtifactTypeDevice_log               ArtifactType = "DEVICE_LOG"
	ArtifactTypeMessage_log              ArtifactType = "MESSAGE_LOG"
	ArtifactTypeVideo_log                ArtifactType = "VIDEO_LOG"
	ArtifactTypeResult_log               ArtifactType = "RESULT_LOG"
	ArtifactTypeService_log              ArtifactType = "SERVICE_LOG"
	ArtifactTypeWebkit_log               ArtifactType = "WEBKIT_LOG"
	ArtifactTypeInstrumentation_output   ArtifactType = "INSTRUMENTATION_OUTPUT"
	ArtifactTypeExerciser_monkey_output  ArtifactType = "EXERCISER_MONKEY_OUTPUT"
	ArtifactTypeCalabash_json_output     ArtifactType = "CALABASH_JSON_OUTPUT"
	ArtifactTypeCalabash_pretty_output   ArtifactType = "CALABASH_PRETTY_OUTPUT"
	ArtifactTypeCalabash_standard_output ArtifactType = "CALABASH_STANDARD_OUTPUT"
	ArtifactTypeCalabash_java_xml_output ArtifactType = "CALABASH_JAVA_XML_OUTPUT"
	ArtifactTypeAutomation_output        ArtifactType = "AUTOMATION_OUTPUT"
	ArtifactTypeAppium_server_output     ArtifactType = "APPIUM_SERVER_OUTPUT"
	ArtifactTypeAppium_java_output       ArtifactType = "APPIUM_JAVA_OUTPUT"
	ArtifactTypeAppium_java_xml_output   ArtifactType = "APPIUM_JAVA_XML_OUTPUT"
	ArtifactTypeAppium_python_output     ArtifactType = "APPIUM_PYTHON_OUTPUT"
	ArtifactTypeAppium_python_xml_output ArtifactType = "APPIUM_PYTHON_XML_OUTPUT"
	ArtifactTypeExplorer_event_log       ArtifactType = "EXPLORER_EVENT_LOG"
	ArtifactTypeExplorer_summary_log     ArtifactType = "EXPLORER_SUMMARY_LOG"
	ArtifactTypeApplication_crash_report ArtifactType = "APPLICATION_CRASH_REPORT"
	ArtifactTypeXctest_log               ArtifactType = "XCTEST_LOG"
	ArtifactTypeVideo                    ArtifactType = "VIDEO"
	ArtifactTypeCustomer_artifact        ArtifactType = "CUSTOMER_ARTIFACT"
	ArtifactTypeCustomer_artifact_log    ArtifactType = "CUSTOMER_ARTIFACT_LOG"
	ArtifactTypeTestspec_output          ArtifactType = "TESTSPEC_OUTPUT"
)

// Values returns all known values for ArtifactType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ArtifactType) Values() []ArtifactType {
	return []ArtifactType{
		"UNKNOWN",
		"SCREENSHOT",
		"DEVICE_LOG",
		"MESSAGE_LOG",
		"VIDEO_LOG",
		"RESULT_LOG",
		"SERVICE_LOG",
		"WEBKIT_LOG",
		"INSTRUMENTATION_OUTPUT",
		"EXERCISER_MONKEY_OUTPUT",
		"CALABASH_JSON_OUTPUT",
		"CALABASH_PRETTY_OUTPUT",
		"CALABASH_STANDARD_OUTPUT",
		"CALABASH_JAVA_XML_OUTPUT",
		"AUTOMATION_OUTPUT",
		"APPIUM_SERVER_OUTPUT",
		"APPIUM_JAVA_OUTPUT",
		"APPIUM_JAVA_XML_OUTPUT",
		"APPIUM_PYTHON_OUTPUT",
		"APPIUM_PYTHON_XML_OUTPUT",
		"EXPLORER_EVENT_LOG",
		"EXPLORER_SUMMARY_LOG",
		"APPLICATION_CRASH_REPORT",
		"XCTEST_LOG",
		"VIDEO",
		"CUSTOMER_ARTIFACT",
		"CUSTOMER_ARTIFACT_LOG",
		"TESTSPEC_OUTPUT",
	}
}

type BillingMethod string

// Enum values for BillingMethod
const (
	BillingMethodMetered   BillingMethod = "METERED"
	BillingMethodUnmetered BillingMethod = "UNMETERED"
)

// Values returns all known values for BillingMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BillingMethod) Values() []BillingMethod {
	return []BillingMethod{
		"METERED",
		"UNMETERED",
	}
}

type CurrencyCode string

// Enum values for CurrencyCode
const (
	CurrencyCodeUsd CurrencyCode = "USD"
)

// Values returns all known values for CurrencyCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CurrencyCode) Values() []CurrencyCode {
	return []CurrencyCode{
		"USD",
	}
}

type DeviceAttribute string

// Enum values for DeviceAttribute
const (
	DeviceAttributeArn                   DeviceAttribute = "ARN"
	DeviceAttributePlatform              DeviceAttribute = "PLATFORM"
	DeviceAttributeForm_factor           DeviceAttribute = "FORM_FACTOR"
	DeviceAttributeManufacturer          DeviceAttribute = "MANUFACTURER"
	DeviceAttributeRemote_access_enabled DeviceAttribute = "REMOTE_ACCESS_ENABLED"
	DeviceAttributeRemote_debug_enabled  DeviceAttribute = "REMOTE_DEBUG_ENABLED"
	DeviceAttributeAppium_version        DeviceAttribute = "APPIUM_VERSION"
	DeviceAttributeInstance_arn          DeviceAttribute = "INSTANCE_ARN"
	DeviceAttributeInstance_labels       DeviceAttribute = "INSTANCE_LABELS"
	DeviceAttributeFleet_type            DeviceAttribute = "FLEET_TYPE"
	DeviceAttributeOs_version            DeviceAttribute = "OS_VERSION"
	DeviceAttributeModel                 DeviceAttribute = "MODEL"
	DeviceAttributeAvailability          DeviceAttribute = "AVAILABILITY"
)

// Values returns all known values for DeviceAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceAttribute) Values() []DeviceAttribute {
	return []DeviceAttribute{
		"ARN",
		"PLATFORM",
		"FORM_FACTOR",
		"MANUFACTURER",
		"REMOTE_ACCESS_ENABLED",
		"REMOTE_DEBUG_ENABLED",
		"APPIUM_VERSION",
		"INSTANCE_ARN",
		"INSTANCE_LABELS",
		"FLEET_TYPE",
		"OS_VERSION",
		"MODEL",
		"AVAILABILITY",
	}
}

type DeviceAvailability string

// Enum values for DeviceAvailability
const (
	DeviceAvailabilityTemporary_not_available DeviceAvailability = "TEMPORARY_NOT_AVAILABLE"
	DeviceAvailabilityBusy                    DeviceAvailability = "BUSY"
	DeviceAvailabilityAvailable               DeviceAvailability = "AVAILABLE"
	DeviceAvailabilityHighly_available        DeviceAvailability = "HIGHLY_AVAILABLE"
)

// Values returns all known values for DeviceAvailability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceAvailability) Values() []DeviceAvailability {
	return []DeviceAvailability{
		"TEMPORARY_NOT_AVAILABLE",
		"BUSY",
		"AVAILABLE",
		"HIGHLY_AVAILABLE",
	}
}

type DeviceFilterAttribute string

// Enum values for DeviceFilterAttribute
const (
	DeviceFilterAttributeArn                   DeviceFilterAttribute = "ARN"
	DeviceFilterAttributePlatform              DeviceFilterAttribute = "PLATFORM"
	DeviceFilterAttributeOs_version            DeviceFilterAttribute = "OS_VERSION"
	DeviceFilterAttributeModel                 DeviceFilterAttribute = "MODEL"
	DeviceFilterAttributeAvailability          DeviceFilterAttribute = "AVAILABILITY"
	DeviceFilterAttributeForm_factor           DeviceFilterAttribute = "FORM_FACTOR"
	DeviceFilterAttributeManufacturer          DeviceFilterAttribute = "MANUFACTURER"
	DeviceFilterAttributeRemote_access_enabled DeviceFilterAttribute = "REMOTE_ACCESS_ENABLED"
	DeviceFilterAttributeRemote_debug_enabled  DeviceFilterAttribute = "REMOTE_DEBUG_ENABLED"
	DeviceFilterAttributeInstance_arn          DeviceFilterAttribute = "INSTANCE_ARN"
	DeviceFilterAttributeInstance_labels       DeviceFilterAttribute = "INSTANCE_LABELS"
	DeviceFilterAttributeFleet_type            DeviceFilterAttribute = "FLEET_TYPE"
)

// Values returns all known values for DeviceFilterAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceFilterAttribute) Values() []DeviceFilterAttribute {
	return []DeviceFilterAttribute{
		"ARN",
		"PLATFORM",
		"OS_VERSION",
		"MODEL",
		"AVAILABILITY",
		"FORM_FACTOR",
		"MANUFACTURER",
		"REMOTE_ACCESS_ENABLED",
		"REMOTE_DEBUG_ENABLED",
		"INSTANCE_ARN",
		"INSTANCE_LABELS",
		"FLEET_TYPE",
	}
}

type DeviceFormFactor string

// Enum values for DeviceFormFactor
const (
	DeviceFormFactorPhone  DeviceFormFactor = "PHONE"
	DeviceFormFactorTablet DeviceFormFactor = "TABLET"
)

// Values returns all known values for DeviceFormFactor. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceFormFactor) Values() []DeviceFormFactor {
	return []DeviceFormFactor{
		"PHONE",
		"TABLET",
	}
}

type DevicePlatform string

// Enum values for DevicePlatform
const (
	DevicePlatformAndroid DevicePlatform = "ANDROID"
	DevicePlatformIos     DevicePlatform = "IOS"
)

// Values returns all known values for DevicePlatform. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DevicePlatform) Values() []DevicePlatform {
	return []DevicePlatform{
		"ANDROID",
		"IOS",
	}
}

type DevicePoolType string

// Enum values for DevicePoolType
const (
	DevicePoolTypeCurated DevicePoolType = "CURATED"
	DevicePoolTypePrivate DevicePoolType = "PRIVATE"
)

// Values returns all known values for DevicePoolType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DevicePoolType) Values() []DevicePoolType {
	return []DevicePoolType{
		"CURATED",
		"PRIVATE",
	}
}

type ExecutionResult string

// Enum values for ExecutionResult
const (
	ExecutionResultPending ExecutionResult = "PENDING"
	ExecutionResultPassed  ExecutionResult = "PASSED"
	ExecutionResultWarned  ExecutionResult = "WARNED"
	ExecutionResultFailed  ExecutionResult = "FAILED"
	ExecutionResultSkipped ExecutionResult = "SKIPPED"
	ExecutionResultErrored ExecutionResult = "ERRORED"
	ExecutionResultStopped ExecutionResult = "STOPPED"
)

// Values returns all known values for ExecutionResult. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutionResult) Values() []ExecutionResult {
	return []ExecutionResult{
		"PENDING",
		"PASSED",
		"WARNED",
		"FAILED",
		"SKIPPED",
		"ERRORED",
		"STOPPED",
	}
}

type ExecutionResultCode string

// Enum values for ExecutionResultCode
const (
	ExecutionResultCodeParsing_failed            ExecutionResultCode = "PARSING_FAILED"
	ExecutionResultCodeVpc_endpoint_setup_failed ExecutionResultCode = "VPC_ENDPOINT_SETUP_FAILED"
)

// Values returns all known values for ExecutionResultCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutionResultCode) Values() []ExecutionResultCode {
	return []ExecutionResultCode{
		"PARSING_FAILED",
		"VPC_ENDPOINT_SETUP_FAILED",
	}
}

type ExecutionStatus string

// Enum values for ExecutionStatus
const (
	ExecutionStatusPending             ExecutionStatus = "PENDING"
	ExecutionStatusPending_concurrnecy ExecutionStatus = "PENDING_CONCURRENCY"
	ExecutionStatusPending_device      ExecutionStatus = "PENDING_DEVICE"
	ExecutionStatusProcessing          ExecutionStatus = "PROCESSING"
	ExecutionStatusScheduling          ExecutionStatus = "SCHEDULING"
	ExecutionStatusPreparing           ExecutionStatus = "PREPARING"
	ExecutionStatusRunning             ExecutionStatus = "RUNNING"
	ExecutionStatusCompleted           ExecutionStatus = "COMPLETED"
	ExecutionStatusStopping            ExecutionStatus = "STOPPING"
)

// Values returns all known values for ExecutionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutionStatus) Values() []ExecutionStatus {
	return []ExecutionStatus{
		"PENDING",
		"PENDING_CONCURRENCY",
		"PENDING_DEVICE",
		"PROCESSING",
		"SCHEDULING",
		"PREPARING",
		"RUNNING",
		"COMPLETED",
		"STOPPING",
	}
}

type InstanceStatus string

// Enum values for InstanceStatus
const (
	InstanceStatusIn_use        InstanceStatus = "IN_USE"
	InstanceStatusPreparing     InstanceStatus = "PREPARING"
	InstanceStatusAvailable     InstanceStatus = "AVAILABLE"
	InstanceStatusNot_available InstanceStatus = "NOT_AVAILABLE"
)

// Values returns all known values for InstanceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceStatus) Values() []InstanceStatus {
	return []InstanceStatus{
		"IN_USE",
		"PREPARING",
		"AVAILABLE",
		"NOT_AVAILABLE",
	}
}

type InteractionMode string

// Enum values for InteractionMode
const (
	InteractionModeInteractive InteractionMode = "INTERACTIVE"
	InteractionModeNo_video    InteractionMode = "NO_VIDEO"
	InteractionModeVideo_only  InteractionMode = "VIDEO_ONLY"
)

// Values returns all known values for InteractionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InteractionMode) Values() []InteractionMode {
	return []InteractionMode{
		"INTERACTIVE",
		"NO_VIDEO",
		"VIDEO_ONLY",
	}
}

type NetworkProfileType string

// Enum values for NetworkProfileType
const (
	NetworkProfileTypeCurated NetworkProfileType = "CURATED"
	NetworkProfileTypePrivate NetworkProfileType = "PRIVATE"
)

// Values returns all known values for NetworkProfileType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NetworkProfileType) Values() []NetworkProfileType {
	return []NetworkProfileType{
		"CURATED",
		"PRIVATE",
	}
}

type OfferingTransactionType string

// Enum values for OfferingTransactionType
const (
	OfferingTransactionTypePurchase OfferingTransactionType = "PURCHASE"
	OfferingTransactionTypeRenew    OfferingTransactionType = "RENEW"
	OfferingTransactionTypeSystem   OfferingTransactionType = "SYSTEM"
)

// Values returns all known values for OfferingTransactionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OfferingTransactionType) Values() []OfferingTransactionType {
	return []OfferingTransactionType{
		"PURCHASE",
		"RENEW",
		"SYSTEM",
	}
}

type OfferingType string

// Enum values for OfferingType
const (
	OfferingTypeRecurring OfferingType = "RECURRING"
)

// Values returns all known values for OfferingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OfferingType) Values() []OfferingType {
	return []OfferingType{
		"RECURRING",
	}
}

type RecurringChargeFrequency string

// Enum values for RecurringChargeFrequency
const (
	RecurringChargeFrequencyMonthly RecurringChargeFrequency = "MONTHLY"
)

// Values returns all known values for RecurringChargeFrequency. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecurringChargeFrequency) Values() []RecurringChargeFrequency {
	return []RecurringChargeFrequency{
		"MONTHLY",
	}
}

type RuleOperator string

// Enum values for RuleOperator
const (
	RuleOperatorEquals                 RuleOperator = "EQUALS"
	RuleOperatorLess_than              RuleOperator = "LESS_THAN"
	RuleOperatorLess_than_or_equals    RuleOperator = "LESS_THAN_OR_EQUALS"
	RuleOperatorGreater_than           RuleOperator = "GREATER_THAN"
	RuleOperatorGreater_than_or_equals RuleOperator = "GREATER_THAN_OR_EQUALS"
	RuleOperatorIn                     RuleOperator = "IN"
	RuleOperatorNot_in                 RuleOperator = "NOT_IN"
	RuleOperatorContains               RuleOperator = "CONTAINS"
)

// Values returns all known values for RuleOperator. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RuleOperator) Values() []RuleOperator {
	return []RuleOperator{
		"EQUALS",
		"LESS_THAN",
		"LESS_THAN_OR_EQUALS",
		"GREATER_THAN",
		"GREATER_THAN_OR_EQUALS",
		"IN",
		"NOT_IN",
		"CONTAINS",
	}
}

type SampleType string

// Enum values for SampleType
const (
	SampleTypeCpu                 SampleType = "CPU"
	SampleTypeMemory              SampleType = "MEMORY"
	SampleTypeThreads             SampleType = "THREADS"
	SampleTypeRx_rate             SampleType = "RX_RATE"
	SampleTypeTx_rate             SampleType = "TX_RATE"
	SampleTypeRx                  SampleType = "RX"
	SampleTypeTx                  SampleType = "TX"
	SampleTypeNative_frames       SampleType = "NATIVE_FRAMES"
	SampleTypeNative_fps          SampleType = "NATIVE_FPS"
	SampleTypeNative_min_drawtime SampleType = "NATIVE_MIN_DRAWTIME"
	SampleTypeNative_avg_drawtime SampleType = "NATIVE_AVG_DRAWTIME"
	SampleTypeNative_max_drawtime SampleType = "NATIVE_MAX_DRAWTIME"
	SampleTypeOpengl_frames       SampleType = "OPENGL_FRAMES"
	SampleTypeOpengl_fps          SampleType = "OPENGL_FPS"
	SampleTypeOpengl_min_drawtime SampleType = "OPENGL_MIN_DRAWTIME"
	SampleTypeOpengl_avg_drawtime SampleType = "OPENGL_AVG_DRAWTIME"
	SampleTypeOpengl_max_drawtime SampleType = "OPENGL_MAX_DRAWTIME"
)

// Values returns all known values for SampleType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SampleType) Values() []SampleType {
	return []SampleType{
		"CPU",
		"MEMORY",
		"THREADS",
		"RX_RATE",
		"TX_RATE",
		"RX",
		"TX",
		"NATIVE_FRAMES",
		"NATIVE_FPS",
		"NATIVE_MIN_DRAWTIME",
		"NATIVE_AVG_DRAWTIME",
		"NATIVE_MAX_DRAWTIME",
		"OPENGL_FRAMES",
		"OPENGL_FPS",
		"OPENGL_MIN_DRAWTIME",
		"OPENGL_AVG_DRAWTIME",
		"OPENGL_MAX_DRAWTIME",
	}
}

type TestGridSessionArtifactCategory string

// Enum values for TestGridSessionArtifactCategory
const (
	TestGridSessionArtifactCategoryVideo TestGridSessionArtifactCategory = "VIDEO"
	TestGridSessionArtifactCategoryLog   TestGridSessionArtifactCategory = "LOG"
)

// Values returns all known values for TestGridSessionArtifactCategory. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TestGridSessionArtifactCategory) Values() []TestGridSessionArtifactCategory {
	return []TestGridSessionArtifactCategory{
		"VIDEO",
		"LOG",
	}
}

type TestGridSessionArtifactType string

// Enum values for TestGridSessionArtifactType
const (
	TestGridSessionArtifactTypeUnknown      TestGridSessionArtifactType = "UNKNOWN"
	TestGridSessionArtifactTypeVideo        TestGridSessionArtifactType = "VIDEO"
	TestGridSessionArtifactTypeSelenium_log TestGridSessionArtifactType = "SELENIUM_LOG"
)

// Values returns all known values for TestGridSessionArtifactType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TestGridSessionArtifactType) Values() []TestGridSessionArtifactType {
	return []TestGridSessionArtifactType{
		"UNKNOWN",
		"VIDEO",
		"SELENIUM_LOG",
	}
}

type TestGridSessionStatus string

// Enum values for TestGridSessionStatus
const (
	TestGridSessionStatusActive  TestGridSessionStatus = "ACTIVE"
	TestGridSessionStatusClosed  TestGridSessionStatus = "CLOSED"
	TestGridSessionStatusErrored TestGridSessionStatus = "ERRORED"
)

// Values returns all known values for TestGridSessionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TestGridSessionStatus) Values() []TestGridSessionStatus {
	return []TestGridSessionStatus{
		"ACTIVE",
		"CLOSED",
		"ERRORED",
	}
}

type TestType string

// Enum values for TestType
const (
	TestTypeBuiltin_fuzz            TestType = "BUILTIN_FUZZ"
	TestTypeBuiltin_explorer        TestType = "BUILTIN_EXPLORER"
	TestTypeWeb_performance_profile TestType = "WEB_PERFORMANCE_PROFILE"
	TestTypeAppium_java_junit       TestType = "APPIUM_JAVA_JUNIT"
	TestTypeAppium_java_testng      TestType = "APPIUM_JAVA_TESTNG"
	TestTypeAppium_python           TestType = "APPIUM_PYTHON"
	TestTypeAppium_node             TestType = "APPIUM_NODE"
	TestTypeAppium_ruby             TestType = "APPIUM_RUBY"
	TestTypeAppium_web_java_junit   TestType = "APPIUM_WEB_JAVA_JUNIT"
	TestTypeAppium_web_java_testng  TestType = "APPIUM_WEB_JAVA_TESTNG"
	TestTypeAppium_web_python       TestType = "APPIUM_WEB_PYTHON"
	TestTypeAppium_web_node         TestType = "APPIUM_WEB_NODE"
	TestTypeAppium_web_ruby         TestType = "APPIUM_WEB_RUBY"
	TestTypeCalabash                TestType = "CALABASH"
	TestTypeInstrumentation         TestType = "INSTRUMENTATION"
	TestTypeUiautomation            TestType = "UIAUTOMATION"
	TestTypeUiautomator             TestType = "UIAUTOMATOR"
	TestTypeXctest                  TestType = "XCTEST"
	TestTypeXctest_ui               TestType = "XCTEST_UI"
	TestTypeRemote_access_record    TestType = "REMOTE_ACCESS_RECORD"
	TestTypeRemote_access_replay    TestType = "REMOTE_ACCESS_REPLAY"
)

// Values returns all known values for TestType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TestType) Values() []TestType {
	return []TestType{
		"BUILTIN_FUZZ",
		"BUILTIN_EXPLORER",
		"WEB_PERFORMANCE_PROFILE",
		"APPIUM_JAVA_JUNIT",
		"APPIUM_JAVA_TESTNG",
		"APPIUM_PYTHON",
		"APPIUM_NODE",
		"APPIUM_RUBY",
		"APPIUM_WEB_JAVA_JUNIT",
		"APPIUM_WEB_JAVA_TESTNG",
		"APPIUM_WEB_PYTHON",
		"APPIUM_WEB_NODE",
		"APPIUM_WEB_RUBY",
		"CALABASH",
		"INSTRUMENTATION",
		"UIAUTOMATION",
		"UIAUTOMATOR",
		"XCTEST",
		"XCTEST_UI",
		"REMOTE_ACCESS_RECORD",
		"REMOTE_ACCESS_REPLAY",
	}
}

type UploadCategory string

// Enum values for UploadCategory
const (
	UploadCategoryCurated UploadCategory = "CURATED"
	UploadCategoryPrivate UploadCategory = "PRIVATE"
)

// Values returns all known values for UploadCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UploadCategory) Values() []UploadCategory {
	return []UploadCategory{
		"CURATED",
		"PRIVATE",
	}
}

type UploadStatus string

// Enum values for UploadStatus
const (
	UploadStatusInitialized UploadStatus = "INITIALIZED"
	UploadStatusProcessing  UploadStatus = "PROCESSING"
	UploadStatusSucceeded   UploadStatus = "SUCCEEDED"
	UploadStatusFailed      UploadStatus = "FAILED"
)

// Values returns all known values for UploadStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UploadStatus) Values() []UploadStatus {
	return []UploadStatus{
		"INITIALIZED",
		"PROCESSING",
		"SUCCEEDED",
		"FAILED",
	}
}

type UploadType string

// Enum values for UploadType
const (
	UploadTypeAndroid_app                         UploadType = "ANDROID_APP"
	UploadTypeIos_app                             UploadType = "IOS_APP"
	UploadTypeWeb_app                             UploadType = "WEB_APP"
	UploadTypeExternal_data                       UploadType = "EXTERNAL_DATA"
	UploadTypeAppium_java_junit_test_package      UploadType = "APPIUM_JAVA_JUNIT_TEST_PACKAGE"
	UploadTypeAppium_java_testng_test_package     UploadType = "APPIUM_JAVA_TESTNG_TEST_PACKAGE"
	UploadTypeAppium_python_test_package          UploadType = "APPIUM_PYTHON_TEST_PACKAGE"
	UploadTypeAppium_node_test_package            UploadType = "APPIUM_NODE_TEST_PACKAGE"
	UploadTypeAppium_ruby_test_package            UploadType = "APPIUM_RUBY_TEST_PACKAGE"
	UploadTypeAppium_web_java_junit_test_package  UploadType = "APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE"
	UploadTypeAppium_web_java_testng_test_package UploadType = "APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE"
	UploadTypeAppium_web_python_test_package      UploadType = "APPIUM_WEB_PYTHON_TEST_PACKAGE"
	UploadTypeAppium_web_node_test_package        UploadType = "APPIUM_WEB_NODE_TEST_PACKAGE"
	UploadTypeAppium_web_ruby_test_package        UploadType = "APPIUM_WEB_RUBY_TEST_PACKAGE"
	UploadTypeCalabash_test_package               UploadType = "CALABASH_TEST_PACKAGE"
	UploadTypeInstrumentation_test_package        UploadType = "INSTRUMENTATION_TEST_PACKAGE"
	UploadTypeUiautomation_test_package           UploadType = "UIAUTOMATION_TEST_PACKAGE"
	UploadTypeUiautomator_test_package            UploadType = "UIAUTOMATOR_TEST_PACKAGE"
	UploadTypeXctest_test_package                 UploadType = "XCTEST_TEST_PACKAGE"
	UploadTypeXctest_ui_test_package              UploadType = "XCTEST_UI_TEST_PACKAGE"
	UploadTypeAppium_java_junit_test_spec         UploadType = "APPIUM_JAVA_JUNIT_TEST_SPEC"
	UploadTypeAppium_java_testng_test_spec        UploadType = "APPIUM_JAVA_TESTNG_TEST_SPEC"
	UploadTypeAppium_python_test_spec             UploadType = "APPIUM_PYTHON_TEST_SPEC"
	UploadTypeAppium_node_test_spec               UploadType = "APPIUM_NODE_TEST_SPEC"
	UploadTypeAppium_ruby_test_spec               UploadType = "APPIUM_RUBY_TEST_SPEC"
	UploadTypeAppium_web_java_junit_test_spec     UploadType = "APPIUM_WEB_JAVA_JUNIT_TEST_SPEC"
	UploadTypeAppium_web_java_testng_test_spec    UploadType = "APPIUM_WEB_JAVA_TESTNG_TEST_SPEC"
	UploadTypeAppium_web_python_test_spec         UploadType = "APPIUM_WEB_PYTHON_TEST_SPEC"
	UploadTypeAppium_web_node_test_spec           UploadType = "APPIUM_WEB_NODE_TEST_SPEC"
	UploadTypeAppium_web_ruby_test_spec           UploadType = "APPIUM_WEB_RUBY_TEST_SPEC"
	UploadTypeInstrumentation_test_spec           UploadType = "INSTRUMENTATION_TEST_SPEC"
	UploadTypeXctest_ui_test_spec                 UploadType = "XCTEST_UI_TEST_SPEC"
)

// Values returns all known values for UploadType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UploadType) Values() []UploadType {
	return []UploadType{
		"ANDROID_APP",
		"IOS_APP",
		"WEB_APP",
		"EXTERNAL_DATA",
		"APPIUM_JAVA_JUNIT_TEST_PACKAGE",
		"APPIUM_JAVA_TESTNG_TEST_PACKAGE",
		"APPIUM_PYTHON_TEST_PACKAGE",
		"APPIUM_NODE_TEST_PACKAGE",
		"APPIUM_RUBY_TEST_PACKAGE",
		"APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE",
		"APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE",
		"APPIUM_WEB_PYTHON_TEST_PACKAGE",
		"APPIUM_WEB_NODE_TEST_PACKAGE",
		"APPIUM_WEB_RUBY_TEST_PACKAGE",
		"CALABASH_TEST_PACKAGE",
		"INSTRUMENTATION_TEST_PACKAGE",
		"UIAUTOMATION_TEST_PACKAGE",
		"UIAUTOMATOR_TEST_PACKAGE",
		"XCTEST_TEST_PACKAGE",
		"XCTEST_UI_TEST_PACKAGE",
		"APPIUM_JAVA_JUNIT_TEST_SPEC",
		"APPIUM_JAVA_TESTNG_TEST_SPEC",
		"APPIUM_PYTHON_TEST_SPEC",
		"APPIUM_NODE_TEST_SPEC",
		"APPIUM_RUBY_TEST_SPEC",
		"APPIUM_WEB_JAVA_JUNIT_TEST_SPEC",
		"APPIUM_WEB_JAVA_TESTNG_TEST_SPEC",
		"APPIUM_WEB_PYTHON_TEST_SPEC",
		"APPIUM_WEB_NODE_TEST_SPEC",
		"APPIUM_WEB_RUBY_TEST_SPEC",
		"INSTRUMENTATION_TEST_SPEC",
		"XCTEST_UI_TEST_SPEC",
	}
}
