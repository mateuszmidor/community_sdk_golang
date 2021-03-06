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
	"fmt"
)

// V202101beta1TaskState the model 'V202101beta1TaskState'
type V202101beta1TaskState string

// List of v202101beta1TaskState
const (
	V202101BETA1TASKSTATE_UNSPECIFIED V202101beta1TaskState = "TASK_STATE_UNSPECIFIED"
	V202101BETA1TASKSTATE_CREATED     V202101beta1TaskState = "TASK_STATE_CREATED"
	V202101BETA1TASKSTATE_UPDATED     V202101beta1TaskState = "TASK_STATE_UPDATED"
	V202101BETA1TASKSTATE_DELETED     V202101beta1TaskState = "TASK_STATE_DELETED"
)

func (v *V202101beta1TaskState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V202101beta1TaskState(value)
	for _, existing := range []V202101beta1TaskState{"TASK_STATE_UNSPECIFIED", "TASK_STATE_CREATED", "TASK_STATE_UPDATED", "TASK_STATE_DELETED"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V202101beta1TaskState", value)
}

// Ptr returns reference to v202101beta1TaskState value
func (v V202101beta1TaskState) Ptr() *V202101beta1TaskState {
	return &v
}

type NullableV202101beta1TaskState struct {
	value *V202101beta1TaskState
	isSet bool
}

func (v NullableV202101beta1TaskState) Get() *V202101beta1TaskState {
	return v.value
}

func (v *NullableV202101beta1TaskState) Set(val *V202101beta1TaskState) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1TaskState) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1TaskState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1TaskState(val *V202101beta1TaskState) *NullableV202101beta1TaskState {
	return &NullableV202101beta1TaskState{value: val, isSet: true}
}

func (v NullableV202101beta1TaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1TaskState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
