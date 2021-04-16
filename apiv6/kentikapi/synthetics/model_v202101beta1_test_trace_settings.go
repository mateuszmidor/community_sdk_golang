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

// V202101beta1TestTraceSettings struct for V202101beta1TestTraceSettings
type V202101beta1TestTraceSettings struct {
	Period   *float32 `json:"period,omitempty"`
	Count    *float32 `json:"count,omitempty"`
	Protocol *string  `json:"protocol,omitempty"`
	Port     *float32 `json:"port,omitempty"`
	Expiry   *float32 `json:"expiry,omitempty"`
	Limit    *float32 `json:"limit,omitempty"`
}

// NewV202101beta1TestTraceSettings instantiates a new V202101beta1TestTraceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1TestTraceSettings() *V202101beta1TestTraceSettings {
	this := V202101beta1TestTraceSettings{}
	return &this
}

// NewV202101beta1TestTraceSettingsWithDefaults instantiates a new V202101beta1TestTraceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1TestTraceSettingsWithDefaults() *V202101beta1TestTraceSettings {
	this := V202101beta1TestTraceSettings{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *V202101beta1TestTraceSettings) GetPeriod() float32 {
	if o == nil || o.Period == nil {
		var ret float32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestTraceSettings) GetPeriodOk() (*float32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *V202101beta1TestTraceSettings) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given float32 and assigns it to the Period field.
func (o *V202101beta1TestTraceSettings) SetPeriod(v float32) {
	o.Period = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V202101beta1TestTraceSettings) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestTraceSettings) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V202101beta1TestTraceSettings) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *V202101beta1TestTraceSettings) SetCount(v float32) {
	o.Count = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V202101beta1TestTraceSettings) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestTraceSettings) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V202101beta1TestTraceSettings) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *V202101beta1TestTraceSettings) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V202101beta1TestTraceSettings) GetPort() float32 {
	if o == nil || o.Port == nil {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestTraceSettings) GetPortOk() (*float32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V202101beta1TestTraceSettings) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *V202101beta1TestTraceSettings) SetPort(v float32) {
	o.Port = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *V202101beta1TestTraceSettings) GetExpiry() float32 {
	if o == nil || o.Expiry == nil {
		var ret float32
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestTraceSettings) GetExpiryOk() (*float32, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *V202101beta1TestTraceSettings) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given float32 and assigns it to the Expiry field.
func (o *V202101beta1TestTraceSettings) SetExpiry(v float32) {
	o.Expiry = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *V202101beta1TestTraceSettings) GetLimit() float32 {
	if o == nil || o.Limit == nil {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TestTraceSettings) GetLimitOk() (*float32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *V202101beta1TestTraceSettings) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *V202101beta1TestTraceSettings) SetLimit(v float32) {
	o.Limit = &v
}

func (o V202101beta1TestTraceSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1TestTraceSettings struct {
	value *V202101beta1TestTraceSettings
	isSet bool
}

func (v NullableV202101beta1TestTraceSettings) Get() *V202101beta1TestTraceSettings {
	return v.value
}

func (v *NullableV202101beta1TestTraceSettings) Set(val *V202101beta1TestTraceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1TestTraceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1TestTraceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1TestTraceSettings(val *V202101beta1TestTraceSettings) *NullableV202101beta1TestTraceSettings {
	return &NullableV202101beta1TestTraceSettings{value: val, isSet: true}
}

func (v NullableV202101beta1TestTraceSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1TestTraceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}