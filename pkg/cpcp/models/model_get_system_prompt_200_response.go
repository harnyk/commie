/*
Commie Plugin API

API for communication between Commie and plugins via CPCP (Commie Plugin Connectivity Protocol).

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// checks if the GetSystemPrompt200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSystemPrompt200Response{}

// GetSystemPrompt200Response struct for GetSystemPrompt200Response
type GetSystemPrompt200Response struct {
	// Optional system prompt provided by the plugin
	SystemPrompt NullableString `json:"system_prompt,omitempty"`
}

// NewGetSystemPrompt200Response instantiates a new GetSystemPrompt200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSystemPrompt200Response() *GetSystemPrompt200Response {
	this := GetSystemPrompt200Response{}
	return &this
}

// NewGetSystemPrompt200ResponseWithDefaults instantiates a new GetSystemPrompt200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSystemPrompt200ResponseWithDefaults() *GetSystemPrompt200Response {
	this := GetSystemPrompt200Response{}
	return &this
}

// GetSystemPrompt returns the SystemPrompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetSystemPrompt200Response) GetSystemPrompt() string {
	if o == nil || IsNil(o.SystemPrompt.Get()) {
		var ret string
		return ret
	}
	return *o.SystemPrompt.Get()
}

// GetSystemPromptOk returns a tuple with the SystemPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSystemPrompt200Response) GetSystemPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemPrompt.Get(), o.SystemPrompt.IsSet()
}

// HasSystemPrompt returns a boolean if a field has been set.
func (o *GetSystemPrompt200Response) HasSystemPrompt() bool {
	if o != nil && o.SystemPrompt.IsSet() {
		return true
	}

	return false
}

// SetSystemPrompt gets a reference to the given NullableString and assigns it to the SystemPrompt field.
func (o *GetSystemPrompt200Response) SetSystemPrompt(v string) {
	o.SystemPrompt.Set(&v)
}
// SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil
func (o *GetSystemPrompt200Response) SetSystemPromptNil() {
	o.SystemPrompt.Set(nil)
}

// UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
func (o *GetSystemPrompt200Response) UnsetSystemPrompt() {
	o.SystemPrompt.Unset()
}

func (o GetSystemPrompt200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSystemPrompt200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemPrompt.IsSet() {
		toSerialize["system_prompt"] = o.SystemPrompt.Get()
	}
	return toSerialize, nil
}

type NullableGetSystemPrompt200Response struct {
	value *GetSystemPrompt200Response
	isSet bool
}

func (v NullableGetSystemPrompt200Response) Get() *GetSystemPrompt200Response {
	return v.value
}

func (v *NullableGetSystemPrompt200Response) Set(val *GetSystemPrompt200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSystemPrompt200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSystemPrompt200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSystemPrompt200Response(val *GetSystemPrompt200Response) *NullableGetSystemPrompt200Response {
	return &NullableGetSystemPrompt200Response{value: val, isSet: true}
}

func (v NullableGetSystemPrompt200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSystemPrompt200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


