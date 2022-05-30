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

// GetUserSSHKeysModel struct for GetUserSSHKeysModel
type GetUserSSHKeysModel struct {
	// The key type filter
	KeyStoreTypeFilter *int32 `json:"keyStoreTypeFilter,omitempty"`
}

// NewGetUserSSHKeysModel instantiates a new GetUserSSHKeysModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserSSHKeysModel() *GetUserSSHKeysModel {
	this := GetUserSSHKeysModel{}
	return &this
}

// NewGetUserSSHKeysModelWithDefaults instantiates a new GetUserSSHKeysModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserSSHKeysModelWithDefaults() *GetUserSSHKeysModel {
	this := GetUserSSHKeysModel{}
	return &this
}

// GetKeyStoreTypeFilter returns the KeyStoreTypeFilter field value if set, zero value otherwise.
func (o *GetUserSSHKeysModel) GetKeyStoreTypeFilter() int32 {
	if o == nil || o.KeyStoreTypeFilter == nil {
		var ret int32
		return ret
	}
	return *o.KeyStoreTypeFilter
}

// GetKeyStoreTypeFilterOk returns a tuple with the KeyStoreTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserSSHKeysModel) GetKeyStoreTypeFilterOk() (*int32, bool) {
	if o == nil || o.KeyStoreTypeFilter == nil {
		return nil, false
	}
	return o.KeyStoreTypeFilter, true
}

// HasKeyStoreTypeFilter returns a boolean if a field has been set.
func (o *GetUserSSHKeysModel) HasKeyStoreTypeFilter() bool {
	if o != nil && o.KeyStoreTypeFilter != nil {
		return true
	}

	return false
}

// SetKeyStoreTypeFilter gets a reference to the given int32 and assigns it to the KeyStoreTypeFilter field.
func (o *GetUserSSHKeysModel) SetKeyStoreTypeFilter(v int32) {
	o.KeyStoreTypeFilter = &v
}

func (o GetUserSSHKeysModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyStoreTypeFilter != nil {
		toSerialize["keyStoreTypeFilter"] = o.KeyStoreTypeFilter
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserSSHKeysModel struct {
	value *GetUserSSHKeysModel
	isSet bool
}

func (v NullableGetUserSSHKeysModel) Get() *GetUserSSHKeysModel {
	return v.value
}

func (v *NullableGetUserSSHKeysModel) Set(val *GetUserSSHKeysModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserSSHKeysModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserSSHKeysModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserSSHKeysModel(val *GetUserSSHKeysModel) *NullableGetUserSSHKeysModel {
	return &NullableGetUserSSHKeysModel{value: val, isSet: true}
}

func (v NullableGetUserSSHKeysModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserSSHKeysModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


