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

// ChangeCredentialsNowProperties struct for ChangeCredentialsNowProperties
type ChangeCredentialsNowProperties struct {
	// Whether or not the CPM will change the credentials in all the accounts that belong to the same group.   This parameter is only relevant for accounts that belong to an account group, and if this parameter does not belong to a group,   it will be ignored.If this account is part of an account group and this value is not specified, the default value will be applied.
	ChangeEntireGroup *bool `json:"ChangeEntireGroup,omitempty"`
}

// NewChangeCredentialsNowProperties instantiates a new ChangeCredentialsNowProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeCredentialsNowProperties() *ChangeCredentialsNowProperties {
	this := ChangeCredentialsNowProperties{}
	return &this
}

// NewChangeCredentialsNowPropertiesWithDefaults instantiates a new ChangeCredentialsNowProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeCredentialsNowPropertiesWithDefaults() *ChangeCredentialsNowProperties {
	this := ChangeCredentialsNowProperties{}
	return &this
}

// GetChangeEntireGroup returns the ChangeEntireGroup field value if set, zero value otherwise.
func (o *ChangeCredentialsNowProperties) GetChangeEntireGroup() bool {
	if o == nil || o.ChangeEntireGroup == nil {
		var ret bool
		return ret
	}
	return *o.ChangeEntireGroup
}

// GetChangeEntireGroupOk returns a tuple with the ChangeEntireGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeCredentialsNowProperties) GetChangeEntireGroupOk() (*bool, bool) {
	if o == nil || o.ChangeEntireGroup == nil {
		return nil, false
	}
	return o.ChangeEntireGroup, true
}

// HasChangeEntireGroup returns a boolean if a field has been set.
func (o *ChangeCredentialsNowProperties) HasChangeEntireGroup() bool {
	if o != nil && o.ChangeEntireGroup != nil {
		return true
	}

	return false
}

// SetChangeEntireGroup gets a reference to the given bool and assigns it to the ChangeEntireGroup field.
func (o *ChangeCredentialsNowProperties) SetChangeEntireGroup(v bool) {
	o.ChangeEntireGroup = &v
}

func (o ChangeCredentialsNowProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeEntireGroup != nil {
		toSerialize["ChangeEntireGroup"] = o.ChangeEntireGroup
	}
	return json.Marshal(toSerialize)
}

type NullableChangeCredentialsNowProperties struct {
	value *ChangeCredentialsNowProperties
	isSet bool
}

func (v NullableChangeCredentialsNowProperties) Get() *ChangeCredentialsNowProperties {
	return v.value
}

func (v *NullableChangeCredentialsNowProperties) Set(val *ChangeCredentialsNowProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeCredentialsNowProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeCredentialsNowProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeCredentialsNowProperties(val *ChangeCredentialsNowProperties) *NullableChangeCredentialsNowProperties {
	return &NullableChangeCredentialsNowProperties{value: val, isSet: true}
}

func (v NullableChangeCredentialsNowProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeCredentialsNowProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


