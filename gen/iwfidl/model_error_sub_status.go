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

// ErrorSubStatus the model 'ErrorSubStatus'
type ErrorSubStatus string

// List of ErrorSubStatus
const (
	UNCATEGORIZED_SUB_STATUS            ErrorSubStatus = "UNCATEGORIZED_SUB_STATUS"
	WORKFLOW_ALREADY_STARTED_SUB_STATUS ErrorSubStatus = "WORKFLOW_ALREADY_STARTED_SUB_STATUS"
	WORKFLOW_NOT_EXISTS_SUB_STATUS      ErrorSubStatus = "WORKFLOW_NOT_EXISTS_SUB_STATUS"
	WORKER_API_ERROR                    ErrorSubStatus = "WORKER_API_ERROR"
	LONG_POLL_TIME_OUT_SUB_STATUS       ErrorSubStatus = "LONG_POLL_TIME_OUT_SUB_STATUS"
)

// All allowed values of ErrorSubStatus enum
var AllowedErrorSubStatusEnumValues = []ErrorSubStatus{
	"UNCATEGORIZED_SUB_STATUS",
	"WORKFLOW_ALREADY_STARTED_SUB_STATUS",
	"WORKFLOW_NOT_EXISTS_SUB_STATUS",
	"WORKER_API_ERROR",
	"LONG_POLL_TIME_OUT_SUB_STATUS",
}

func (v *ErrorSubStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorSubStatus(value)
	for _, existing := range AllowedErrorSubStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorSubStatus", value)
}

// NewErrorSubStatusFromValue returns a pointer to a valid ErrorSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorSubStatusFromValue(v string) (*ErrorSubStatus, error) {
	ev := ErrorSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorSubStatus: valid values are %v", v, AllowedErrorSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorSubStatus) IsValid() bool {
	for _, existing := range AllowedErrorSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorSubStatus value
func (v ErrorSubStatus) Ptr() *ErrorSubStatus {
	return &v
}

type NullableErrorSubStatus struct {
	value *ErrorSubStatus
	isSet bool
}

func (v NullableErrorSubStatus) Get() *ErrorSubStatus {
	return v.value
}

func (v *NullableErrorSubStatus) Set(val *ErrorSubStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSubStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSubStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSubStatus(val *ErrorSubStatus) *NullableErrorSubStatus {
	return &NullableErrorSubStatus{value: val, isSet: true}
}

func (v NullableErrorSubStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSubStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
