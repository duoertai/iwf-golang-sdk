/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// checks if the StateDecision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateDecision{}

// StateDecision struct for StateDecision
type StateDecision struct {
	NextStates []StateMovement `json:"nextStates,omitempty"`
}

// NewStateDecision instantiates a new StateDecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateDecision() *StateDecision {
	this := StateDecision{}
	return &this
}

// NewStateDecisionWithDefaults instantiates a new StateDecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateDecisionWithDefaults() *StateDecision {
	this := StateDecision{}
	return &this
}

// GetNextStates returns the NextStates field value if set, zero value otherwise.
func (o *StateDecision) GetNextStates() []StateMovement {
	if o == nil || IsNil(o.NextStates) {
		var ret []StateMovement
		return ret
	}
	return o.NextStates
}

// GetNextStatesOk returns a tuple with the NextStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateDecision) GetNextStatesOk() ([]StateMovement, bool) {
	if o == nil || IsNil(o.NextStates) {
		return nil, false
	}
	return o.NextStates, true
}

// HasNextStates returns a boolean if a field has been set.
func (o *StateDecision) HasNextStates() bool {
	if o != nil && !IsNil(o.NextStates) {
		return true
	}

	return false
}

// SetNextStates gets a reference to the given []StateMovement and assigns it to the NextStates field.
func (o *StateDecision) SetNextStates(v []StateMovement) {
	o.NextStates = v
}

func (o StateDecision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateDecision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextStates) {
		toSerialize["nextStates"] = o.NextStates
	}
	return toSerialize, nil
}

type NullableStateDecision struct {
	value *StateDecision
	isSet bool
}

func (v NullableStateDecision) Get() *StateDecision {
	return v.value
}

func (v *NullableStateDecision) Set(val *StateDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableStateDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableStateDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateDecision(val *StateDecision) *NullableStateDecision {
	return &NullableStateDecision{value: val, isSet: true}
}

func (v NullableStateDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


