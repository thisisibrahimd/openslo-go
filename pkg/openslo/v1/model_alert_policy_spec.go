/*
OpenSLO V1 Schema

OpenSLO V1 Schema

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openslo_v1

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AlertPolicySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertPolicySpec{}

// AlertPolicySpec struct for AlertPolicySpec
type AlertPolicySpec struct {
	Description *string `json:"description,omitempty"`
	AlertWhenNoData bool `json:"alertWhenNoData"`
	AlertWhenResolved bool `json:"alertWhenResolved"`
	AlertWhenBreaching bool `json:"alertWhenBreaching"`
	Conditions []AlertPolicyCondition `json:"conditions"`
	NotificationTargets []AlertPolicyNotificationTarget `json:"notificationTargets"`
}

type _AlertPolicySpec AlertPolicySpec

// NewAlertPolicySpec instantiates a new AlertPolicySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertPolicySpec(alertWhenNoData bool, alertWhenResolved bool, alertWhenBreaching bool, conditions []AlertPolicyCondition, notificationTargets []AlertPolicyNotificationTarget) *AlertPolicySpec {
	this := AlertPolicySpec{}
	this.AlertWhenNoData = alertWhenNoData
	this.AlertWhenResolved = alertWhenResolved
	this.AlertWhenBreaching = alertWhenBreaching
	this.Conditions = conditions
	this.NotificationTargets = notificationTargets
	return &this
}

// NewAlertPolicySpecWithDefaults instantiates a new AlertPolicySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertPolicySpecWithDefaults() *AlertPolicySpec {
	this := AlertPolicySpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertPolicySpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertPolicySpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertPolicySpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertPolicySpec) SetDescription(v string) {
	o.Description = &v
}

// GetAlertWhenNoData returns the AlertWhenNoData field value
func (o *AlertPolicySpec) GetAlertWhenNoData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlertWhenNoData
}

// GetAlertWhenNoDataOk returns a tuple with the AlertWhenNoData field value
// and a boolean to check if the value has been set.
func (o *AlertPolicySpec) GetAlertWhenNoDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertWhenNoData, true
}

// SetAlertWhenNoData sets field value
func (o *AlertPolicySpec) SetAlertWhenNoData(v bool) {
	o.AlertWhenNoData = v
}

// GetAlertWhenResolved returns the AlertWhenResolved field value
func (o *AlertPolicySpec) GetAlertWhenResolved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlertWhenResolved
}

// GetAlertWhenResolvedOk returns a tuple with the AlertWhenResolved field value
// and a boolean to check if the value has been set.
func (o *AlertPolicySpec) GetAlertWhenResolvedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertWhenResolved, true
}

// SetAlertWhenResolved sets field value
func (o *AlertPolicySpec) SetAlertWhenResolved(v bool) {
	o.AlertWhenResolved = v
}

// GetAlertWhenBreaching returns the AlertWhenBreaching field value
func (o *AlertPolicySpec) GetAlertWhenBreaching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlertWhenBreaching
}

// GetAlertWhenBreachingOk returns a tuple with the AlertWhenBreaching field value
// and a boolean to check if the value has been set.
func (o *AlertPolicySpec) GetAlertWhenBreachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertWhenBreaching, true
}

// SetAlertWhenBreaching sets field value
func (o *AlertPolicySpec) SetAlertWhenBreaching(v bool) {
	o.AlertWhenBreaching = v
}

// GetConditions returns the Conditions field value
func (o *AlertPolicySpec) GetConditions() []AlertPolicyCondition {
	if o == nil {
		var ret []AlertPolicyCondition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *AlertPolicySpec) GetConditionsOk() ([]AlertPolicyCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *AlertPolicySpec) SetConditions(v []AlertPolicyCondition) {
	o.Conditions = v
}

// GetNotificationTargets returns the NotificationTargets field value
func (o *AlertPolicySpec) GetNotificationTargets() []AlertPolicyNotificationTarget {
	if o == nil {
		var ret []AlertPolicyNotificationTarget
		return ret
	}

	return o.NotificationTargets
}

// GetNotificationTargetsOk returns a tuple with the NotificationTargets field value
// and a boolean to check if the value has been set.
func (o *AlertPolicySpec) GetNotificationTargetsOk() ([]AlertPolicyNotificationTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationTargets, true
}

// SetNotificationTargets sets field value
func (o *AlertPolicySpec) SetNotificationTargets(v []AlertPolicyNotificationTarget) {
	o.NotificationTargets = v
}

func (o AlertPolicySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertPolicySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["alertWhenNoData"] = o.AlertWhenNoData
	toSerialize["alertWhenResolved"] = o.AlertWhenResolved
	toSerialize["alertWhenBreaching"] = o.AlertWhenBreaching
	toSerialize["conditions"] = o.Conditions
	toSerialize["notificationTargets"] = o.NotificationTargets
	return toSerialize, nil
}

func (o *AlertPolicySpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alertWhenNoData",
		"alertWhenResolved",
		"alertWhenBreaching",
		"conditions",
		"notificationTargets",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlertPolicySpec := _AlertPolicySpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertPolicySpec)

	if err != nil {
		return err
	}

	*o = AlertPolicySpec(varAlertPolicySpec)

	return err
}

type NullableAlertPolicySpec struct {
	value *AlertPolicySpec
	isSet bool
}

func (v NullableAlertPolicySpec) Get() *AlertPolicySpec {
	return v.value
}

func (v *NullableAlertPolicySpec) Set(val *AlertPolicySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertPolicySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertPolicySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertPolicySpec(val *AlertPolicySpec) *NullableAlertPolicySpec {
	return &NullableAlertPolicySpec{value: val, isSet: true}
}

func (v NullableAlertPolicySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertPolicySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


