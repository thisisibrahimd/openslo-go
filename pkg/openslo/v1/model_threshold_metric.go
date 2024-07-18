/*
OpenSLO V1 Schema

OpenSLO V1 Schema

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openslo_v1

import (
	"encoding/json"
)

// checks if the ThresholdMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdMetric{}

// ThresholdMetric struct for ThresholdMetric
type ThresholdMetric struct {
	MetricSource *MetricSource `json:"metricSource,omitempty"`
}

// NewThresholdMetric instantiates a new ThresholdMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdMetric() *ThresholdMetric {
	this := ThresholdMetric{}
	return &this
}

// NewThresholdMetricWithDefaults instantiates a new ThresholdMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdMetricWithDefaults() *ThresholdMetric {
	this := ThresholdMetric{}
	return &this
}

// GetMetricSource returns the MetricSource field value if set, zero value otherwise.
func (o *ThresholdMetric) GetMetricSource() MetricSource {
	if o == nil || IsNil(o.MetricSource) {
		var ret MetricSource
		return ret
	}
	return *o.MetricSource
}

// GetMetricSourceOk returns a tuple with the MetricSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMetric) GetMetricSourceOk() (*MetricSource, bool) {
	if o == nil || IsNil(o.MetricSource) {
		return nil, false
	}
	return o.MetricSource, true
}

// HasMetricSource returns a boolean if a field has been set.
func (o *ThresholdMetric) HasMetricSource() bool {
	if o != nil && !IsNil(o.MetricSource) {
		return true
	}

	return false
}

// SetMetricSource gets a reference to the given MetricSource and assigns it to the MetricSource field.
func (o *ThresholdMetric) SetMetricSource(v MetricSource) {
	o.MetricSource = &v
}

func (o ThresholdMetric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricSource) {
		toSerialize["metricSource"] = o.MetricSource
	}
	return toSerialize, nil
}

type NullableThresholdMetric struct {
	value *ThresholdMetric
	isSet bool
}

func (v NullableThresholdMetric) Get() *ThresholdMetric {
	return v.value
}

func (v *NullableThresholdMetric) Set(val *ThresholdMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdMetric(val *ThresholdMetric) *NullableThresholdMetric {
	return &NullableThresholdMetric{value: val, isSet: true}
}

func (v NullableThresholdMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
