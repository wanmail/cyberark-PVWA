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

// AddAccountGroupData struct for AddAccountGroupData
type AddAccountGroupData struct {
	// The name of the newly created group.
	GroupName string `json:"GroupName"`
	// The name of the platform of the group.  The associated platform must be set to PolicyType = Group.
	GroupPlatformId string `json:"GroupPlatformId"`
	// The name of the Safe where the group will be created.
	Safe string `json:"Safe"`
}

// NewAddAccountGroupData instantiates a new AddAccountGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAccountGroupData(groupName string, groupPlatformId string, safe string) *AddAccountGroupData {
	this := AddAccountGroupData{}
	this.GroupName = groupName
	this.GroupPlatformId = groupPlatformId
	this.Safe = safe
	return &this
}

// NewAddAccountGroupDataWithDefaults instantiates a new AddAccountGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAccountGroupDataWithDefaults() *AddAccountGroupData {
	this := AddAccountGroupData{}
	return &this
}

// GetGroupName returns the GroupName field value
func (o *AddAccountGroupData) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *AddAccountGroupData) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *AddAccountGroupData) SetGroupName(v string) {
	o.GroupName = v
}

// GetGroupPlatformId returns the GroupPlatformId field value
func (o *AddAccountGroupData) GetGroupPlatformId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupPlatformId
}

// GetGroupPlatformIdOk returns a tuple with the GroupPlatformId field value
// and a boolean to check if the value has been set.
func (o *AddAccountGroupData) GetGroupPlatformIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupPlatformId, true
}

// SetGroupPlatformId sets field value
func (o *AddAccountGroupData) SetGroupPlatformId(v string) {
	o.GroupPlatformId = v
}

// GetSafe returns the Safe field value
func (o *AddAccountGroupData) GetSafe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Safe
}

// GetSafeOk returns a tuple with the Safe field value
// and a boolean to check if the value has been set.
func (o *AddAccountGroupData) GetSafeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Safe, true
}

// SetSafe sets field value
func (o *AddAccountGroupData) SetSafe(v string) {
	o.Safe = v
}

func (o AddAccountGroupData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["GroupName"] = o.GroupName
	}
	if true {
		toSerialize["GroupPlatformId"] = o.GroupPlatformId
	}
	if true {
		toSerialize["Safe"] = o.Safe
	}
	return json.Marshal(toSerialize)
}

type NullableAddAccountGroupData struct {
	value *AddAccountGroupData
	isSet bool
}

func (v NullableAddAccountGroupData) Get() *AddAccountGroupData {
	return v.value
}

func (v *NullableAddAccountGroupData) Set(val *AddAccountGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAccountGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAccountGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAccountGroupData(val *AddAccountGroupData) *NullableAddAccountGroupData {
	return &NullableAddAccountGroupData{value: val, isSet: true}
}

func (v NullableAddAccountGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAccountGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


