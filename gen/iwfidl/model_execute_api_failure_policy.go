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

// ExecuteApiFailurePolicy the model 'ExecuteApiFailurePolicy'
type ExecuteApiFailurePolicy string

// List of ExecuteApiFailurePolicy
const (
	FAIL_WORKFLOW_ON_EXECUTE_API_FAILURE ExecuteApiFailurePolicy = "FAIL_WORKFLOW_ON_EXECUTE_API_FAILURE"
	PROCEED_TO_CONFIGURED_STATE          ExecuteApiFailurePolicy = "PROCEED_TO_CONFIGURED_STATE"
)

// All allowed values of ExecuteApiFailurePolicy enum
var AllowedExecuteApiFailurePolicyEnumValues = []ExecuteApiFailurePolicy{
	"FAIL_WORKFLOW_ON_EXECUTE_API_FAILURE",
	"PROCEED_TO_CONFIGURED_STATE",
}

func (v *ExecuteApiFailurePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExecuteApiFailurePolicy(value)
	for _, existing := range AllowedExecuteApiFailurePolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExecuteApiFailurePolicy", value)
}

// NewExecuteApiFailurePolicyFromValue returns a pointer to a valid ExecuteApiFailurePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExecuteApiFailurePolicyFromValue(v string) (*ExecuteApiFailurePolicy, error) {
	ev := ExecuteApiFailurePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExecuteApiFailurePolicy: valid values are %v", v, AllowedExecuteApiFailurePolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExecuteApiFailurePolicy) IsValid() bool {
	for _, existing := range AllowedExecuteApiFailurePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecuteApiFailurePolicy value
func (v ExecuteApiFailurePolicy) Ptr() *ExecuteApiFailurePolicy {
	return &v
}

type NullableExecuteApiFailurePolicy struct {
	value *ExecuteApiFailurePolicy
	isSet bool
}

func (v NullableExecuteApiFailurePolicy) Get() *ExecuteApiFailurePolicy {
	return v.value
}

func (v *NullableExecuteApiFailurePolicy) Set(val *ExecuteApiFailurePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteApiFailurePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteApiFailurePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteApiFailurePolicy(val *ExecuteApiFailurePolicy) *NullableExecuteApiFailurePolicy {
	return &NullableExecuteApiFailurePolicy{value: val, isSet: true}
}

func (v NullableExecuteApiFailurePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteApiFailurePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
