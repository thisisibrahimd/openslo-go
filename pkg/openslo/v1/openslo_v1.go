// An Alert Condition allows you to define under which conditions an alert for an SLO needs
// to be triggered.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type AlertCondition struct {
	// The version of specification format for this particular entity that this is written                   
	// against.                                                                                              
	APIVersion                                                                            APIVersion         `json:"apiVersion"`
	Kind                                                                                  AlertConditionKind `json:"kind"`
	Metadata                                                                              Metadata           `json:"metadata"`
	Spec                                                                                  AlertConditionSpec `json:"spec"`
}

type Metadata struct {
	// key <> value pairs which should be used to define implementation / system specific                             
	// metadata about the SLO. For example, it can be metadata about a dashboard url, or how to                       
	// name a metric created by the SLI, etc.                                                                         
	Annotations                                                                                map[string]interface{} `json:"annotations,omitempty"`
	DisplayName                                                                                *string                `json:"displayName,omitempty"`
	// key <> value pairs of labels that can be used as metadata relevant to users                                    
	Labels                                                                                     map[string]*Label      `json:"labels,omitempty"`
	Name                                                                                       string                 `json:"name"`
}

type AlertConditionSpec struct {
	Condition   PurpleCondition `json:"condition"`
	Description *string         `json:"description,omitempty"`
	Severity    string          `json:"severity"`
}

type PurpleCondition struct {
	AlertAfter     *string     `json:"alertAfter,omitempty"`
	Kind           *PurpleKind `json:"kind,omitempty"`
	LookbackWindow *string     `json:"lookbackWindow,omitempty"`
	Op             *Op         `json:"op,omitempty"`
	Threshold      *float64    `json:"threshold,omitempty"`
}

// An Alert Notification Target defines the possible targets where alert notifications
// should be delivered to. For example, this can be a web-hook, Slack or any other custom
// target.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type AlertNotificationTarget struct {
	// The version of specification format for this particular entity that this is written                            
	// against.                                                                                                       
	APIVersion                                                                            APIVersion                  `json:"apiVersion"`
	Kind                                                                                  AlertNotificationTargetKind `json:"kind"`
	Metadata                                                                              Metadata                    `json:"metadata"`
	Spec                                                                                  AlertNotificationTargetSpec `json:"spec"`
}

type AlertNotificationTargetSpec struct {
	Description *string `json:"description,omitempty"`
	Target      string  `json:"target"`
}

// An Alert Policy allows you to define the alert conditions for an SLO.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type AlertPolicy struct {
	// The version of specification format for this particular entity that this is written                
	// against.                                                                                           
	APIVersion                                                                            APIVersion      `json:"apiVersion"`
	Kind                                                                                  AlertPolicyKind `json:"kind"`
	Metadata                                                                              Metadata        `json:"metadata"`
	Spec                                                                                  AlertPolicySpec `json:"spec"`
}

type AlertPolicySpec struct {
	AlertWhenBreaching  bool                 `json:"alertWhenBreaching"`
	AlertWhenNoData     bool                 `json:"alertWhenNoData"`
	AlertWhenResolved   bool                 `json:"alertWhenResolved"`
	Conditions          []ConditionElement   `json:"conditions"`
	Description         *string              `json:"description,omitempty"`
	NotificationTargets []NotificationTarget `json:"notificationTargets"`
}

type ConditionElement struct {
	ConditionRef *string             `json:"conditionRef,omitempty"`
	Kind         *AlertConditionKind `json:"kind,omitempty"`
	Metadata     *ConditionMetadata  `json:"metadata,omitempty"`
	Spec         *SpecClass          `json:"spec,omitempty"`
}

type ConditionMetadata struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}

// An Alert Condition allows you to define under which conditions an alert for an SLO needs
// to be triggered.
type SpecClass struct {
	Spec *PurpleSpec `json:"spec,omitempty"`
}

type PurpleSpec struct {
	Condition   PurpleCondition `json:"condition"`
	Description *string         `json:"description,omitempty"`
	Severity    string          `json:"severity"`
}

type NotificationTarget struct {
	Kind      *AlertConditionKind     `json:"kind,omitempty"`
	Metadata  *ConditionMetadata      `json:"metadata,omitempty"`
	Spec      *NotificationTargetSpec `json:"spec,omitempty"`
	TargetRef *string                 `json:"targetRef,omitempty"`
}

// An Alert Notification Target defines the possible targets where alert notifications
// should be delivered to. For example, this can be a web-hook, Slack or any other custom
// target.
type NotificationTargetSpec struct {
	Spec *FluffySpec `json:"spec,omitempty"`
}

type FluffySpec struct {
	Description *string `json:"description,omitempty"`
	Target      string  `json:"target"`
}

// A budget adjustment is a date-time range from when events should be ignored in error
// budget calculations
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type BudgetAdjustment struct {
	// The version of specification format for this particular entity that this is written                     
	// against.                                                                                                
	APIVersion                                                                            APIVersion           `json:"apiVersion"`
	Kind                                                                                  BudgetAdjustmentKind `json:"kind"`
	Metadata                                                                              Metadata             `json:"metadata"`
	Spec                                                                                  BudgetAdjustmentSpec `json:"spec"`
}

import "time"

type BudgetAdjustmentSpec struct {
	Description  *string            `json:"description,omitempty"`
	EndTime      time.Time          `json:"endTime"`
	IndicatorRef *string            `json:"indicatorRef,omitempty"`
	Service      *string            `json:"service,omitempty"`
	StartTime    time.Time          `json:"startTime"`
	TimeWindow   []PurpleTimeWindow `json:"timeWindow"`
	Indicator    interface{}        `json:"indicator"`
}

type PurpleTimeWindow struct {
	Duration string `json:"duration"`
}

// A DataSource represents connection details with a particular metric source.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type Datasource struct {
	// The version of specification format for this particular entity that this is written               
	// against.                                                                                          
	APIVersion                                                                            APIVersion     `json:"apiVersion"`
	Kind                                                                                  DatasourceKind `json:"kind"`
	Metadata                                                                              Metadata       `json:"metadata"`
	Spec                                                                                  DatasourceSpec `json:"spec"`
}

type DatasourceSpec struct {
	// Fields used for creating a connection with particular datasource e.g. AccessKeys,                       
	// SecretKeys, etc. Everything that is valid YAML can be put here.                                         
	ConnectionDetails                                                                   map[string]interface{} `json:"connectionDetails"`
	Type                                                                                string                 `json:"type"`
}

// A Service is a high-level grouping of SLO.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type Service struct {
	// The version of specification format for this particular entity that this is written            
	// against.                                                                                       
	APIVersion                                                                            APIVersion  `json:"apiVersion"`
	Kind                                                                                  ServiceKind `json:"kind"`
	Metadata                                                                              Metadata    `json:"metadata"`
	Spec                                                                                  ServiceSpec `json:"spec"`
}

type ServiceSpec struct {
	Description *string `json:"description,omitempty"`
}

// A service level indicator (SLI) represents how to read metrics from data sources.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type SLI struct {
	// The version of specification format for this particular entity that this is written           
	// against.                                                                                      
	APIVersion                                                                            APIVersion `json:"apiVersion"`
	Kind                                                                                  SLIKind    `json:"kind"`
	Metadata                                                                              Metadata   `json:"metadata"`
	Spec                                                                                  SLISpec    `json:"spec"`
}

type SLISpec struct {
	RatioMetric     *RatioMetric     `json:"ratioMetric,omitempty"`
	ThresholdMetric *ThresholdMetric `json:"thresholdMetric,omitempty"`
}

type RatioMetric struct {
	Bad     *Bad  `json:"bad,omitempty"`
	Counter bool  `json:"counter"`
	Good    *Good `json:"good,omitempty"`
	Total   Total `json:"total"`
}

type Bad struct {
	MetricSource MetricSource `json:"metricSource"`
}

type MetricSource struct {
	MetricSourceRef *string                `json:"metricSourceRef,omitempty"`
	Spec            map[string]interface{} `json:"spec"`
	Type            *string                `json:"type,omitempty"`
}

type Good struct {
	MetricSource MetricSource `json:"metricSource"`
}

type Total struct {
	MetricSource MetricSource `json:"metricSource"`
}

type ThresholdMetric struct {
	MetricSource MetricSource `json:"metricSource"`
}

// A service level objective (SLO) is a target value or a range of values for a service
// level that is described by a service level indicator (SLI).
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type Slo struct {
	// The version of specification format for this particular entity that this is written           
	// against.                                                                                      
	APIVersion                                                                            APIVersion `json:"apiVersion"`
	Kind                                                                                  SLOKind    `json:"kind"`
	Metadata                                                                              Metadata   `json:"metadata"`
	Spec                                                                                  SLOSpec    `json:"spec"`
}

type SLOSpec struct {
	AlertPolicies   []string           `json:"alertPolicies,omitempty"`
	BudgetingMethod BudgetingMethod    `json:"budgetingMethod"`
	Description     *string            `json:"description,omitempty"`
	Indicator       *Indicator         `json:"indicator,omitempty"`
	IndicatorRef    *string            `json:"indicatorRef,omitempty"`
	Objectives      []Objective        `json:"objectives"`
	Service         string             `json:"service"`
	TimeWindow      []FluffyTimeWindow `json:"timeWindow,omitempty"`
}

// A service level indicator (SLI) represents how to read metrics from data sources.
type Indicator struct {
	Metadata *Metadata      `json:"metadata,omitempty"`
	Spec     *IndicatorSpec `json:"spec,omitempty"`
}

type IndicatorSpec struct {
	RatioMetric     *RatioMetric     `json:"ratioMetric,omitempty"`
	ThresholdMetric *ThresholdMetric `json:"thresholdMetric,omitempty"`
}

type Objective struct {
	DisplayName     *string          `json:"displayName,omitempty"`
	Op              *Op              `json:"op,omitempty"`
	Target          float64          `json:"target"`
	TimeSliceTarget *float64         `json:"timeSliceTarget,omitempty"`
	TimeSliceWindow *TimeSliceWindow `json:"timeSliceWindow"`
	Value           *float64         `json:"value,omitempty"`
}

type FluffyTimeWindow struct {
	Duration string `json:"duration"`
}

type APIVersion string

const (
	OpensloV1 APIVersion = "openslo/v1"
)

type AlertConditionKind string

const (
	KindAlertCondition AlertConditionKind = "AlertCondition"
)

type PurpleKind string

const (
	Burnrate PurpleKind = "burnrate"
)

type Op string

const (
	Gt  Op = "gt"
	Gte Op = "gte"
	LTE Op = "lte"
	Lt  Op = "lt"
)

type AlertNotificationTargetKind string

const (
	KindAlertNotificationTarget AlertNotificationTargetKind = "AlertNotificationTarget"
)

type AlertPolicyKind string

const (
	KindAlertPolicy AlertPolicyKind = "AlertPolicy"
)

type BudgetAdjustmentKind string

const (
	KindBudgetAdjustment BudgetAdjustmentKind = "BudgetAdjustment"
)

type DatasourceKind string

const (
	DataSource DatasourceKind = "DataSource"
)

type ServiceKind string

const (
	KindService ServiceKind = "Service"
)

type SLIKind string

const (
	KindSLI SLIKind = "SLI"
)

type SLOKind string

const (
	KindSLO SLOKind = "SLO"
)

type BudgetingMethod string

const (
	Occurrences BudgetingMethod = "Occurrences"
	Timeslices  BudgetingMethod = "Timeslices"
)

type Label struct {
	AnythingArray []interface{}
	String        *string
}

type TimeSliceWindow struct {
	Double *float64
	String *string
}
