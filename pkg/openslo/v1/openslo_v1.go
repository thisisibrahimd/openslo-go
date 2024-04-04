// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    alertCondition, err := UnmarshalAlertCondition(bytes)
//    bytes, err = alertCondition.Marshal()
//
//    alertNotificationTarget, err := UnmarshalAlertNotificationTarget(bytes)
//    bytes, err = alertNotificationTarget.Marshal()
//
//    alertPolicy, err := UnmarshalAlertPolicy(bytes)
//    bytes, err = alertPolicy.Marshal()
//
//    budgetAdjustment, err := UnmarshalBudgetAdjustment(bytes)
//    bytes, err = budgetAdjustment.Marshal()
//
//    datasource, err := UnmarshalDatasource(bytes)
//    bytes, err = datasource.Marshal()
//
//    service, err := UnmarshalService(bytes)
//    bytes, err = service.Marshal()
//
//    sLI, err := UnmarshalSLI(bytes)
//    bytes, err = sLI.Marshal()
//
//    slo, err := UnmarshalSlo(bytes)
//    bytes, err = slo.Marshal()

package openslo_v1

import "bytes"
import "errors"
import "time"

import "encoding/json"

func UnmarshalAlertCondition(data []byte) (AlertCondition, error) {
	var r AlertCondition
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlertCondition) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAlertNotificationTarget(data []byte) (AlertNotificationTarget, error) {
	var r AlertNotificationTarget
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlertNotificationTarget) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAlertPolicy(data []byte) (AlertPolicy, error) {
	var r AlertPolicy
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlertPolicy) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBudgetAdjustment(data []byte) (BudgetAdjustment, error) {
	var r BudgetAdjustment
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BudgetAdjustment) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasource(data []byte) (Datasource, error) {
	var r Datasource
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Datasource) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalService(data []byte) (Service, error) {
	var r Service
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Service) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSLI(data []byte) (SLI, error) {
	var r SLI
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SLI) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlo(data []byte) (Slo, error) {
	var r Slo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Slo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// An Alert Condition allows you to define under which conditions an alert for an SLO needs
// to be triggered.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type AlertCondition struct {
	// The version of specification format for this particular entity that this is written                   
	// against.                                                                                              
	APIVersion                                                                            APIVersion         `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  AlertConditionKind `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata           `yaml:"metadata" json:"metadata"`
	Spec                                                                                  AlertConditionSpec `yaml:"spec" json:"spec"`
}

type Metadata struct {
	// key <> value pairs which should be used to define implementation / system specific                             
	// metadata about the SLO. For example, it can be metadata about a dashboard url, or how to                       
	// name a metric created by the SLI, etc.                                                                         
	Annotations                                                                                map[string]interface{} `yaml:"annotations,omitempty" json:"annotations,omitempty"`
	DisplayName                                                                                *string                `yaml:"displayName,omitempty" json:"displayName,omitempty"`
	// key <> value pairs of labels that can be used as metadata relevant to users                                    
	Labels                                                                                     map[string]*Label      `yaml:"labels,omitempty" json:"labels,omitempty"`
	Name                                                                                       string                 `yaml:"name" json:"name"`
}

type AlertConditionSpec struct {
	Condition   PurpleCondition `yaml:"condition" json:"condition"`
	Description *string         `yaml:"description,omitempty" json:"description,omitempty"`
	Severity    string          `yaml:"severity" json:"severity"`
}

type PurpleCondition struct {
	AlertAfter     *string     `yaml:"alertAfter,omitempty" json:"alertAfter,omitempty"`
	Kind           *PurpleKind `yaml:"kind,omitempty" json:"kind,omitempty"`
	LookbackWindow *string     `yaml:"lookbackWindow,omitempty" json:"lookbackWindow,omitempty"`
	Op             *Op         `yaml:"op,omitempty" json:"op,omitempty"`
	Threshold      *float64    `yaml:"threshold,omitempty" json:"threshold,omitempty"`
}

// An Alert Notification Target defines the possible targets where alert notifications
// should be delivered to. For example, this can be a web-hook, Slack or any other custom
// target.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type AlertNotificationTarget struct {
	// The version of specification format for this particular entity that this is written                            
	// against.                                                                                                       
	APIVersion                                                                            APIVersion                  `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  AlertNotificationTargetKind `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata                    `yaml:"metadata" json:"metadata"`
	Spec                                                                                  AlertNotificationTargetSpec `yaml:"spec" json:"spec"`
}

type AlertNotificationTargetSpec struct {
	Description *string `yaml:"description,omitempty" json:"description,omitempty"`
	Target      string  `yaml:"target" json:"target"`
}

// An Alert Policy allows you to define the alert conditions for an SLO.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type AlertPolicy struct {
	// The version of specification format for this particular entity that this is written                
	// against.                                                                                           
	APIVersion                                                                            APIVersion      `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  AlertPolicyKind `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata        `yaml:"metadata" json:"metadata"`
	Spec                                                                                  AlertPolicySpec `yaml:"spec" json:"spec"`
}

type AlertPolicySpec struct {
	AlertWhenBreaching  bool                 `yaml:"alertWhenBreaching" json:"alertWhenBreaching"`
	AlertWhenNoData     bool                 `yaml:"alertWhenNoData" json:"alertWhenNoData"`
	AlertWhenResolved   bool                 `yaml:"alertWhenResolved" json:"alertWhenResolved"`
	Conditions          []ConditionElement   `yaml:"conditions" json:"conditions"`
	Description         *string              `yaml:"description,omitempty" json:"description,omitempty"`
	NotificationTargets []NotificationTarget `yaml:"notificationTargets" json:"notificationTargets"`
}

type ConditionElement struct {
	ConditionRef *string             `yaml:"conditionRef,omitempty" json:"conditionRef,omitempty"`
	Kind         *AlertConditionKind `yaml:"kind,omitempty" json:"kind,omitempty"`
	Metadata     *ConditionMetadata  `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec         *SpecClass          `yaml:"spec,omitempty" json:"spec,omitempty"`
}

type ConditionMetadata struct {
	Metadata *Metadata `yaml:"metadata,omitempty" json:"metadata,omitempty"`
}

// An Alert Condition allows you to define under which conditions an alert for an SLO needs
// to be triggered.
type SpecClass struct {
	Spec *PurpleSpec `yaml:"spec,omitempty" json:"spec,omitempty"`
}

type PurpleSpec struct {
	Condition   PurpleCondition `yaml:"condition" json:"condition"`
	Description *string         `yaml:"description,omitempty" json:"description,omitempty"`
	Severity    string          `yaml:"severity" json:"severity"`
}

type NotificationTarget struct {
	Kind      *AlertConditionKind     `yaml:"kind,omitempty" json:"kind,omitempty"`
	Metadata  *ConditionMetadata      `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec      *NotificationTargetSpec `yaml:"spec,omitempty" json:"spec,omitempty"`
	TargetRef *string                 `yaml:"targetRef,omitempty" json:"targetRef,omitempty"`
}

// An Alert Notification Target defines the possible targets where alert notifications
// should be delivered to. For example, this can be a web-hook, Slack or any other custom
// target.
type NotificationTargetSpec struct {
	Spec *FluffySpec `yaml:"spec,omitempty" json:"spec,omitempty"`
}

type FluffySpec struct {
	Description *string `yaml:"description,omitempty" json:"description,omitempty"`
	Target      string  `yaml:"target" json:"target"`
}

// A budget adjustment is a date-time range from when events should be ignored in error
// budget calculations
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type BudgetAdjustment struct {
	// The version of specification format for this particular entity that this is written                     
	// against.                                                                                                
	APIVersion                                                                            APIVersion           `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  BudgetAdjustmentKind `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata             `yaml:"metadata" json:"metadata"`
	Spec                                                                                  BudgetAdjustmentSpec `yaml:"spec" json:"spec"`
}

type BudgetAdjustmentSpec struct {
	Description  *string            `yaml:"description,omitempty" json:"description,omitempty"`
	EndTime      time.Time          `yaml:"endTime" json:"endTime"`
	IndicatorRef *string            `yaml:"indicatorRef,omitempty" json:"indicatorRef,omitempty"`
	Service      *string            `yaml:"service,omitempty" json:"service,omitempty"`
	StartTime    time.Time          `yaml:"startTime" json:"startTime"`
	TimeWindow   []PurpleTimeWindow `yaml:"timeWindow" json:"timeWindow"`
	Indicator    interface{}        `yaml:"indicator,omitempty" json:"indicator,omitempty"`
}

type PurpleTimeWindow struct {
	Duration string `yaml:"duration" json:"duration"`
}

// A DataSource represents connection details with a particular metric source.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type Datasource struct {
	// The version of specification format for this particular entity that this is written               
	// against.                                                                                          
	APIVersion                                                                            APIVersion     `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  DatasourceKind `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata       `yaml:"metadata" json:"metadata"`
	Spec                                                                                  DatasourceSpec `yaml:"spec" json:"spec"`
}

type DatasourceSpec struct {
	// Fields used for creating a connection with particular datasource e.g. AccessKeys,                       
	// SecretKeys, etc. Everything that is valid YAML can be put here.                                         
	ConnectionDetails                                                                   map[string]interface{} `yaml:"connectionDetails" json:"connectionDetails"`
	Type                                                                                string                 `yaml:"type" json:"type"`
}

// A Service is a high-level grouping of SLO.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type Service struct {
	// The version of specification format for this particular entity that this is written            
	// against.                                                                                       
	APIVersion                                                                            APIVersion  `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  ServiceKind `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata    `yaml:"metadata" json:"metadata"`
	Spec                                                                                  ServiceSpec `yaml:"spec" json:"spec"`
}

type ServiceSpec struct {
	Description *string `yaml:"description,omitempty" json:"description,omitempty"`
}

// A service level indicator (SLI) represents how to read metrics from data sources.
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type SLI struct {
	// The version of specification format for this particular entity that this is written           
	// against.                                                                                      
	APIVersion                                                                            APIVersion `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  SLIKind    `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata   `yaml:"metadata" json:"metadata"`
	Spec                                                                                  SLISpec    `yaml:"spec" json:"spec"`
}

type SLISpec struct {
	RatioMetric     *RatioMetric     `yaml:"ratioMetric,omitempty" json:"ratioMetric,omitempty"`
	ThresholdMetric *ThresholdMetric `yaml:"thresholdMetric,omitempty" json:"thresholdMetric,omitempty"`
}

type RatioMetric struct {
	Bad     *Bad  `yaml:"bad,omitempty" json:"bad,omitempty"`
	Counter bool  `yaml:"counter" json:"counter"`
	Good    *Good `yaml:"good,omitempty" json:"good,omitempty"`
	Total   Total `yaml:"total" json:"total"`
}

type Bad struct {
	MetricSource MetricSource `yaml:"metricSource" json:"metricSource"`
}

type MetricSource struct {
	MetricSourceRef *string                `yaml:"metricSourceRef,omitempty" json:"metricSourceRef,omitempty"`
	Spec            map[string]interface{} `yaml:"spec" json:"spec"`
	Type            *string                `yaml:"type,omitempty" json:"type,omitempty"`
}

type Good struct {
	MetricSource MetricSource `yaml:"metricSource" json:"metricSource"`
}

type Total struct {
	MetricSource MetricSource `yaml:"metricSource" json:"metricSource"`
}

type ThresholdMetric struct {
	MetricSource MetricSource `yaml:"metricSource" json:"metricSource"`
}

// A service level objective (SLO) is a target value or a range of values for a service
// level that is described by a service level indicator (SLI).
//
// The OpenSLO General Schema lays out the basic structure of an OpenSLO document.
type Slo struct {
	// The version of specification format for this particular entity that this is written           
	// against.                                                                                      
	APIVersion                                                                            APIVersion `yaml:"apiVersion" json:"apiVersion"`
	Kind                                                                                  SLOKind    `yaml:"kind" json:"kind"`
	Metadata                                                                              Metadata   `yaml:"metadata" json:"metadata"`
	Spec                                                                                  SLOSpec    `yaml:"spec" json:"spec"`
}

type SLOSpec struct {
	AlertPolicies   []string           `yaml:"alertPolicies,omitempty" json:"alertPolicies,omitempty"`
	BudgetingMethod BudgetingMethod    `yaml:"budgetingMethod" json:"budgetingMethod"`
	Description     *string            `yaml:"description,omitempty" json:"description,omitempty"`
	Indicator       *Indicator         `yaml:"indicator,omitempty" json:"indicator,omitempty"`
	IndicatorRef    *string            `yaml:"indicatorRef,omitempty" json:"indicatorRef,omitempty"`
	Objectives      []Objective        `yaml:"objectives" json:"objectives"`
	Service         string             `yaml:"service" json:"service"`
	TimeWindow      []FluffyTimeWindow `yaml:"timeWindow,omitempty" json:"timeWindow,omitempty"`
}

// A service level indicator (SLI) represents how to read metrics from data sources.
type Indicator struct {
	Metadata *Metadata      `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec     *IndicatorSpec `yaml:"spec,omitempty" json:"spec,omitempty"`
}

type IndicatorSpec struct {
	RatioMetric     *RatioMetric     `yaml:"ratioMetric,omitempty" json:"ratioMetric,omitempty"`
	ThresholdMetric *ThresholdMetric `yaml:"thresholdMetric,omitempty" json:"thresholdMetric,omitempty"`
}

type Objective struct {
	DisplayName     *string          `yaml:"displayName,omitempty" json:"displayName,omitempty"`
	Op              *Op              `yaml:"op,omitempty" json:"op,omitempty"`
	Target          float64          `yaml:"target" json:"target"`
	TimeSliceTarget *float64         `yaml:"timeSliceTarget,omitempty" json:"timeSliceTarget,omitempty"`
	TimeSliceWindow *TimeSliceWindow `yaml:"timeSliceWindow,omitempty" json:"timeSliceWindow,omitempty"`
	Value           *float64         `yaml:"value,omitempty" json:"value,omitempty"`
}

type FluffyTimeWindow struct {
	Duration string `yaml:"duration" json:"duration"`
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

func (x *Label) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.AnythingArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Label) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, false, nil, false)
}

type TimeSliceWindow struct {
	Double *float64
	String *string
}

func (x *TimeSliceWindow) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *TimeSliceWindow) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
