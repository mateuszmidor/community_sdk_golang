/*
 * Cloud Export Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudexport

import (
	"encoding/json"
)

// V202101beta1AwsProperties struct for V202101beta1AwsProperties
type V202101beta1AwsProperties struct {
	Bucket          *string `json:"bucket,omitempty"`
	IamRoleArn      *string `json:"iamRoleArn,omitempty"`
	Region          *string `json:"region,omitempty"`
	DeleteAfterRead *bool   `json:"deleteAfterRead,omitempty"`
	MultipleBuckets *bool   `json:"multipleBuckets,omitempty"`
}

// NewV202101beta1AwsProperties instantiates a new V202101beta1AwsProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1AwsProperties() *V202101beta1AwsProperties {
	this := V202101beta1AwsProperties{}
	return &this
}

// NewV202101beta1AwsPropertiesWithDefaults instantiates a new V202101beta1AwsProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1AwsPropertiesWithDefaults() *V202101beta1AwsProperties {
	this := V202101beta1AwsProperties{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *V202101beta1AwsProperties) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1AwsProperties) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *V202101beta1AwsProperties) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *V202101beta1AwsProperties) SetBucket(v string) {
	o.Bucket = &v
}

// GetIamRoleArn returns the IamRoleArn field value if set, zero value otherwise.
func (o *V202101beta1AwsProperties) GetIamRoleArn() string {
	if o == nil || o.IamRoleArn == nil {
		var ret string
		return ret
	}
	return *o.IamRoleArn
}

// GetIamRoleArnOk returns a tuple with the IamRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1AwsProperties) GetIamRoleArnOk() (*string, bool) {
	if o == nil || o.IamRoleArn == nil {
		return nil, false
	}
	return o.IamRoleArn, true
}

// HasIamRoleArn returns a boolean if a field has been set.
func (o *V202101beta1AwsProperties) HasIamRoleArn() bool {
	if o != nil && o.IamRoleArn != nil {
		return true
	}

	return false
}

// SetIamRoleArn gets a reference to the given string and assigns it to the IamRoleArn field.
func (o *V202101beta1AwsProperties) SetIamRoleArn(v string) {
	o.IamRoleArn = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *V202101beta1AwsProperties) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1AwsProperties) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *V202101beta1AwsProperties) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *V202101beta1AwsProperties) SetRegion(v string) {
	o.Region = &v
}

// GetDeleteAfterRead returns the DeleteAfterRead field value if set, zero value otherwise.
func (o *V202101beta1AwsProperties) GetDeleteAfterRead() bool {
	if o == nil || o.DeleteAfterRead == nil {
		var ret bool
		return ret
	}
	return *o.DeleteAfterRead
}

// GetDeleteAfterReadOk returns a tuple with the DeleteAfterRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1AwsProperties) GetDeleteAfterReadOk() (*bool, bool) {
	if o == nil || o.DeleteAfterRead == nil {
		return nil, false
	}
	return o.DeleteAfterRead, true
}

// HasDeleteAfterRead returns a boolean if a field has been set.
func (o *V202101beta1AwsProperties) HasDeleteAfterRead() bool {
	if o != nil && o.DeleteAfterRead != nil {
		return true
	}

	return false
}

// SetDeleteAfterRead gets a reference to the given bool and assigns it to the DeleteAfterRead field.
func (o *V202101beta1AwsProperties) SetDeleteAfterRead(v bool) {
	o.DeleteAfterRead = &v
}

// GetMultipleBuckets returns the MultipleBuckets field value if set, zero value otherwise.
func (o *V202101beta1AwsProperties) GetMultipleBuckets() bool {
	if o == nil || o.MultipleBuckets == nil {
		var ret bool
		return ret
	}
	return *o.MultipleBuckets
}

// GetMultipleBucketsOk returns a tuple with the MultipleBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1AwsProperties) GetMultipleBucketsOk() (*bool, bool) {
	if o == nil || o.MultipleBuckets == nil {
		return nil, false
	}
	return o.MultipleBuckets, true
}

// HasMultipleBuckets returns a boolean if a field has been set.
func (o *V202101beta1AwsProperties) HasMultipleBuckets() bool {
	if o != nil && o.MultipleBuckets != nil {
		return true
	}

	return false
}

// SetMultipleBuckets gets a reference to the given bool and assigns it to the MultipleBuckets field.
func (o *V202101beta1AwsProperties) SetMultipleBuckets(v bool) {
	o.MultipleBuckets = &v
}

func (o V202101beta1AwsProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.IamRoleArn != nil {
		toSerialize["iamRoleArn"] = o.IamRoleArn
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.DeleteAfterRead != nil {
		toSerialize["deleteAfterRead"] = o.DeleteAfterRead
	}
	if o.MultipleBuckets != nil {
		toSerialize["multipleBuckets"] = o.MultipleBuckets
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1AwsProperties struct {
	value *V202101beta1AwsProperties
	isSet bool
}

func (v NullableV202101beta1AwsProperties) Get() *V202101beta1AwsProperties {
	return v.value
}

func (v *NullableV202101beta1AwsProperties) Set(val *V202101beta1AwsProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1AwsProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1AwsProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1AwsProperties(val *V202101beta1AwsProperties) *NullableV202101beta1AwsProperties {
	return &NullableV202101beta1AwsProperties{value: val, isSet: true}
}

func (v NullableV202101beta1AwsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1AwsProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
