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

// V202101beta1IPInfo struct for V202101beta1IPInfo
type V202101beta1IPInfo struct {
	Ip  *string          `json:"ip,omitempty"`
	Asn *V202101beta1ASN `json:"asn,omitempty"`
	Geo *V202101beta1Geo `json:"geo,omitempty"`
	Dns *V202101beta1DNS `json:"dns,omitempty"`
}

// NewV202101beta1IPInfo instantiates a new V202101beta1IPInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1IPInfo() *V202101beta1IPInfo {
	this := V202101beta1IPInfo{}
	return &this
}

// NewV202101beta1IPInfoWithDefaults instantiates a new V202101beta1IPInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1IPInfoWithDefaults() *V202101beta1IPInfo {
	this := V202101beta1IPInfo{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *V202101beta1IPInfo) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1IPInfo) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *V202101beta1IPInfo) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *V202101beta1IPInfo) SetIp(v string) {
	o.Ip = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *V202101beta1IPInfo) GetAsn() V202101beta1ASN {
	if o == nil || o.Asn == nil {
		var ret V202101beta1ASN
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1IPInfo) GetAsnOk() (*V202101beta1ASN, bool) {
	if o == nil || o.Asn == nil {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *V202101beta1IPInfo) HasAsn() bool {
	if o != nil && o.Asn != nil {
		return true
	}

	return false
}

// SetAsn gets a reference to the given V202101beta1ASN and assigns it to the Asn field.
func (o *V202101beta1IPInfo) SetAsn(v V202101beta1ASN) {
	o.Asn = &v
}

// GetGeo returns the Geo field value if set, zero value otherwise.
func (o *V202101beta1IPInfo) GetGeo() V202101beta1Geo {
	if o == nil || o.Geo == nil {
		var ret V202101beta1Geo
		return ret
	}
	return *o.Geo
}

// GetGeoOk returns a tuple with the Geo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1IPInfo) GetGeoOk() (*V202101beta1Geo, bool) {
	if o == nil || o.Geo == nil {
		return nil, false
	}
	return o.Geo, true
}

// HasGeo returns a boolean if a field has been set.
func (o *V202101beta1IPInfo) HasGeo() bool {
	if o != nil && o.Geo != nil {
		return true
	}

	return false
}

// SetGeo gets a reference to the given V202101beta1Geo and assigns it to the Geo field.
func (o *V202101beta1IPInfo) SetGeo(v V202101beta1Geo) {
	o.Geo = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *V202101beta1IPInfo) GetDns() V202101beta1DNS {
	if o == nil || o.Dns == nil {
		var ret V202101beta1DNS
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1IPInfo) GetDnsOk() (*V202101beta1DNS, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *V202101beta1IPInfo) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given V202101beta1DNS and assigns it to the Dns field.
func (o *V202101beta1IPInfo) SetDns(v V202101beta1DNS) {
	o.Dns = &v
}

func (o V202101beta1IPInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Asn != nil {
		toSerialize["asn"] = o.Asn
	}
	if o.Geo != nil {
		toSerialize["geo"] = o.Geo
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1IPInfo struct {
	value *V202101beta1IPInfo
	isSet bool
}

func (v NullableV202101beta1IPInfo) Get() *V202101beta1IPInfo {
	return v.value
}

func (v *NullableV202101beta1IPInfo) Set(val *V202101beta1IPInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1IPInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1IPInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1IPInfo(val *V202101beta1IPInfo) *NullableV202101beta1IPInfo {
	return &NullableV202101beta1IPInfo{value: val, isSet: true}
}

func (v NullableV202101beta1IPInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1IPInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
