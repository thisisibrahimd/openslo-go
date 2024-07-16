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

// checks if the AlertPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertPolicy{}

// AlertPolicy struct for AlertPolicy
type AlertPolicy struct {
	ApiVersion *OpensloApiVersion `json:"apiVersion,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Spec *AlertPolicySpec `json:"spec,omitempty"`
}

// NewAlertPolicy instantiates a new AlertPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertPolicy() *AlertPolicy {
	this := AlertPolicy{}
	return &this
}

// NewAlertPolicyWithDefaults instantiates a new AlertPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertPolicyWithDefaults() *AlertPolicy {
	this := AlertPolicy{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AlertPolicy) GetApiVersion() OpensloApiVersion {
	if o == nil || IsNil(o.ApiVersion) {
		var ret OpensloApiVersion
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertPolicy) GetApiVersionOk() (*OpensloApiVersion, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AlertPolicy) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given OpensloApiVersion and assigns it to the ApiVersion field.
func (o *AlertPolicy) SetApiVersion(v OpensloApiVersion) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AlertPolicy) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertPolicy) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AlertPolicy) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AlertPolicy) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AlertPolicy) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertPolicy) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AlertPolicy) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *AlertPolicy) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *AlertPolicy) GetSpec() AlertPolicySpec {
	if o == nil || IsNil(o.Spec) {
		var ret AlertPolicySpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertPolicy) GetSpecOk() (*AlertPolicySpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *AlertPolicy) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given AlertPolicySpec and assigns it to the Spec field.
func (o *AlertPolicy) SetSpec(v AlertPolicySpec) {
	o.Spec = &v
}

func (o AlertPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

type NullableAlertPolicy struct {
	value *AlertPolicy
	isSet bool
}

func (v NullableAlertPolicy) Get() *AlertPolicy {
	return v.value
}

func (v *NullableAlertPolicy) Set(val *AlertPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertPolicy(val *AlertPolicy) *NullableAlertPolicy {
	return &NullableAlertPolicy{value: val, isSet: true}
}

func (v NullableAlertPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


