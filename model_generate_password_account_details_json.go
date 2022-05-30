/*
Privileged Access Security REST API

Display the PVWA REST APIs below for a description of how to use them and try them out. Access information about additional REST APIs through the online documentation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GeneratePasswordAccountDetailsJson struct for GeneratePasswordAccountDetailsJson
type GeneratePasswordAccountDetailsJson struct {
	AccountId string `json:"AccountId"`
}

// NewGeneratePasswordAccountDetailsJson instantiates a new GeneratePasswordAccountDetailsJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePasswordAccountDetailsJson(accountId string) *GeneratePasswordAccountDetailsJson {
	this := GeneratePasswordAccountDetailsJson{}
	this.AccountId = accountId
	return &this
}

// NewGeneratePasswordAccountDetailsJsonWithDefaults instantiates a new GeneratePasswordAccountDetailsJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePasswordAccountDetailsJsonWithDefaults() *GeneratePasswordAccountDetailsJson {
	this := GeneratePasswordAccountDetailsJson{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *GeneratePasswordAccountDetailsJson) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *GeneratePasswordAccountDetailsJson) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *GeneratePasswordAccountDetailsJson) SetAccountId(v string) {
	o.AccountId = v
}

func (o GeneratePasswordAccountDetailsJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AccountId"] = o.AccountId
	}
	return json.Marshal(toSerialize)
}

type NullableGeneratePasswordAccountDetailsJson struct {
	value *GeneratePasswordAccountDetailsJson
	isSet bool
}

func (v NullableGeneratePasswordAccountDetailsJson) Get() *GeneratePasswordAccountDetailsJson {
	return v.value
}

func (v *NullableGeneratePasswordAccountDetailsJson) Set(val *GeneratePasswordAccountDetailsJson) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePasswordAccountDetailsJson) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePasswordAccountDetailsJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePasswordAccountDetailsJson(val *GeneratePasswordAccountDetailsJson) *NullableGeneratePasswordAccountDetailsJson {
	return &NullableGeneratePasswordAccountDetailsJson{value: val, isSet: true}
}

func (v NullableGeneratePasswordAccountDetailsJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratePasswordAccountDetailsJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


