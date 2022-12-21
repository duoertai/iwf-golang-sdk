/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
	"fmt"
)

// TimerStatus the model 'TimerStatus'
type TimerStatus string

// List of TimerStatus
const (
	SCHEDULED TimerStatus = "SCHEDULED"
	FIRED TimerStatus = "FIRED"
)

// All allowed values of TimerStatus enum
var AllowedTimerStatusEnumValues = []TimerStatus{
	"SCHEDULED",
	"FIRED",
}

func (v *TimerStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimerStatus(value)
	for _, existing := range AllowedTimerStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimerStatus", value)
}

// NewTimerStatusFromValue returns a pointer to a valid TimerStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimerStatusFromValue(v string) (*TimerStatus, error) {
	ev := TimerStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimerStatus: valid values are %v", v, AllowedTimerStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimerStatus) IsValid() bool {
	for _, existing := range AllowedTimerStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimerStatus value
func (v TimerStatus) Ptr() *TimerStatus {
	return &v
}

type NullableTimerStatus struct {
	value *TimerStatus
	isSet bool
}

func (v NullableTimerStatus) Get() *TimerStatus {
	return v.value
}

func (v *NullableTimerStatus) Set(val *TimerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerStatus(val *TimerStatus) *NullableTimerStatus {
	return &NullableTimerStatus{value: val, isSet: true}
}

func (v NullableTimerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

