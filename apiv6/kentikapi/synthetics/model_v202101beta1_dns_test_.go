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
)

// V202101beta1DnsTest struct for V202101beta1DnsTest
type V202101beta1DnsTest struct {
	Target *string `json:"target,omitempty"`
}

// NewV202101beta1DnsTest instantiates a new V202101beta1DnsTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1DnsTest() *V202101beta1DnsTest {
	this := V202101beta1DnsTest{}
	return &this
}

// NewV202101beta1DnsTestWithDefaults instantiates a new V202101beta1DnsTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1DnsTestWithDefaults() *V202101beta1DnsTest {
	this := V202101beta1DnsTest{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *V202101beta1DnsTest) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1DnsTest) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *V202101beta1DnsTest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *V202101beta1DnsTest) SetTarget(v string) {
	o.Target = &v
}

func (o V202101beta1DnsTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1DnsTest struct {
	value *V202101beta1DnsTest
	isSet bool
}

func (v NullableV202101beta1DnsTest) Get() *V202101beta1DnsTest {
	return v.value
}

func (v *NullableV202101beta1DnsTest) Set(val *V202101beta1DnsTest) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1DnsTest) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1DnsTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1DnsTest(val *V202101beta1DnsTest) *NullableV202101beta1DnsTest {
	return &NullableV202101beta1DnsTest{value: val, isSet: true}
}

func (v NullableV202101beta1DnsTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1DnsTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
