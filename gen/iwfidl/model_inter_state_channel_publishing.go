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

// InterStateChannelPublishing struct for InterStateChannelPublishing
type InterStateChannelPublishing struct {
	ChannelName string `json:"channelName"`
	Value *EncodedObject `json:"value,omitempty"`
}

// NewInterStateChannelPublishing instantiates a new InterStateChannelPublishing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterStateChannelPublishing(channelName string) *InterStateChannelPublishing {
	this := InterStateChannelPublishing{}
	this.ChannelName = channelName
	return &this
}

// NewInterStateChannelPublishingWithDefaults instantiates a new InterStateChannelPublishing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterStateChannelPublishingWithDefaults() *InterStateChannelPublishing {
	this := InterStateChannelPublishing{}
	return &this
}

// GetChannelName returns the ChannelName field value
func (o *InterStateChannelPublishing) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelPublishing) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *InterStateChannelPublishing) SetChannelName(v string) {
	o.ChannelName = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterStateChannelPublishing) GetValue() EncodedObject {
	if o == nil || o.Value == nil {
		var ret EncodedObject
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterStateChannelPublishing) GetValueOk() (*EncodedObject, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterStateChannelPublishing) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given EncodedObject and assigns it to the Value field.
func (o *InterStateChannelPublishing) SetValue(v EncodedObject) {
	o.Value = &v
}

func (o InterStateChannelPublishing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channelName"] = o.ChannelName
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInterStateChannelPublishing struct {
	value *InterStateChannelPublishing
	isSet bool
}

func (v NullableInterStateChannelPublishing) Get() *InterStateChannelPublishing {
	return v.value
}

func (v *NullableInterStateChannelPublishing) Set(val *InterStateChannelPublishing) {
	v.value = val
	v.isSet = true
}

func (v NullableInterStateChannelPublishing) IsSet() bool {
	return v.isSet
}

func (v *NullableInterStateChannelPublishing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterStateChannelPublishing(val *InterStateChannelPublishing) *NullableInterStateChannelPublishing {
	return &NullableInterStateChannelPublishing{value: val, isSet: true}
}

func (v NullableInterStateChannelPublishing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterStateChannelPublishing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

