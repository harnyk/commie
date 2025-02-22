/*
Commie Plugin API

API for communication between Commie and plugins via CPCP (Commie Plugin Connectivity Protocol).

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RunTools200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTools200Response{}

// RunTools200Response struct for RunTools200Response
type RunTools200Response struct {
	ToolOutputs []ToolOutput `json:"tool_outputs"`
}

type _RunTools200Response RunTools200Response

// NewRunTools200Response instantiates a new RunTools200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTools200Response(toolOutputs []ToolOutput) *RunTools200Response {
	this := RunTools200Response{}
	this.ToolOutputs = toolOutputs
	return &this
}

// NewRunTools200ResponseWithDefaults instantiates a new RunTools200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTools200ResponseWithDefaults() *RunTools200Response {
	this := RunTools200Response{}
	return &this
}

// GetToolOutputs returns the ToolOutputs field value
func (o *RunTools200Response) GetToolOutputs() []ToolOutput {
	if o == nil {
		var ret []ToolOutput
		return ret
	}

	return o.ToolOutputs
}

// GetToolOutputsOk returns a tuple with the ToolOutputs field value
// and a boolean to check if the value has been set.
func (o *RunTools200Response) GetToolOutputsOk() ([]ToolOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolOutputs, true
}

// SetToolOutputs sets field value
func (o *RunTools200Response) SetToolOutputs(v []ToolOutput) {
	o.ToolOutputs = v
}

func (o RunTools200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTools200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tool_outputs"] = o.ToolOutputs
	return toSerialize, nil
}

func (o *RunTools200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tool_outputs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRunTools200Response := _RunTools200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunTools200Response)

	if err != nil {
		return err
	}

	*o = RunTools200Response(varRunTools200Response)

	return err
}

type NullableRunTools200Response struct {
	value *RunTools200Response
	isSet bool
}

func (v NullableRunTools200Response) Get() *RunTools200Response {
	return v.value
}

func (v *NullableRunTools200Response) Set(val *RunTools200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTools200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTools200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTools200Response(val *RunTools200Response) *NullableRunTools200Response {
	return &NullableRunTools200Response{value: val, isSet: true}
}

func (v NullableRunTools200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTools200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


