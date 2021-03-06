/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synthetics

import (
	"encoding/json"
	"time"
)

// V202101beta1Health struct for V202101beta1Health
type V202101beta1Health struct {
	Health *string    `json:"health,omitempty"`
	Time   *time.Time `json:"time,omitempty"`
}

// NewV202101beta1Health instantiates a new V202101beta1Health object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1Health() *V202101beta1Health {
	this := V202101beta1Health{}
	return &this
}

// NewV202101beta1HealthWithDefaults instantiates a new V202101beta1Health object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1HealthWithDefaults() *V202101beta1Health {
	this := V202101beta1Health{}
	return &this
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *V202101beta1Health) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Health) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *V202101beta1Health) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *V202101beta1Health) SetHealth(v string) {
	o.Health = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V202101beta1Health) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Health) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V202101beta1Health) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *V202101beta1Health) SetTime(v time.Time) {
	o.Time = &v
}

func (o V202101beta1Health) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1Health struct {
	value *V202101beta1Health
	isSet bool
}

func (v NullableV202101beta1Health) Get() *V202101beta1Health {
	return v.value
}

func (v *NullableV202101beta1Health) Set(val *V202101beta1Health) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1Health) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1Health) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1Health(val *V202101beta1Health) *NullableV202101beta1Health {
	return &NullableV202101beta1Health{value: val, isSet: true}
}

func (v NullableV202101beta1Health) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1Health) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
