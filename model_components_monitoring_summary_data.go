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

// ComponentsMonitoringSummaryData struct for ComponentsMonitoringSummaryData
type ComponentsMonitoringSummaryData struct {
	Components []ComponentMonitoringSummaryData `json:"Components,omitempty"`
	Vaults []ComponentMonitoringSummaryVaultData `json:"Vaults,omitempty"`
}

// NewComponentsMonitoringSummaryData instantiates a new ComponentsMonitoringSummaryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentsMonitoringSummaryData() *ComponentsMonitoringSummaryData {
	this := ComponentsMonitoringSummaryData{}
	return &this
}

// NewComponentsMonitoringSummaryDataWithDefaults instantiates a new ComponentsMonitoringSummaryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentsMonitoringSummaryDataWithDefaults() *ComponentsMonitoringSummaryData {
	this := ComponentsMonitoringSummaryData{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ComponentsMonitoringSummaryData) GetComponents() []ComponentMonitoringSummaryData {
	if o == nil || o.Components == nil {
		var ret []ComponentMonitoringSummaryData
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsMonitoringSummaryData) GetComponentsOk() ([]ComponentMonitoringSummaryData, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ComponentsMonitoringSummaryData) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ComponentMonitoringSummaryData and assigns it to the Components field.
func (o *ComponentsMonitoringSummaryData) SetComponents(v []ComponentMonitoringSummaryData) {
	o.Components = v
}

// GetVaults returns the Vaults field value if set, zero value otherwise.
func (o *ComponentsMonitoringSummaryData) GetVaults() []ComponentMonitoringSummaryVaultData {
	if o == nil || o.Vaults == nil {
		var ret []ComponentMonitoringSummaryVaultData
		return ret
	}
	return o.Vaults
}

// GetVaultsOk returns a tuple with the Vaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsMonitoringSummaryData) GetVaultsOk() ([]ComponentMonitoringSummaryVaultData, bool) {
	if o == nil || o.Vaults == nil {
		return nil, false
	}
	return o.Vaults, true
}

// HasVaults returns a boolean if a field has been set.
func (o *ComponentsMonitoringSummaryData) HasVaults() bool {
	if o != nil && o.Vaults != nil {
		return true
	}

	return false
}

// SetVaults gets a reference to the given []ComponentMonitoringSummaryVaultData and assigns it to the Vaults field.
func (o *ComponentsMonitoringSummaryData) SetVaults(v []ComponentMonitoringSummaryVaultData) {
	o.Vaults = v
}

func (o ComponentsMonitoringSummaryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["Components"] = o.Components
	}
	if o.Vaults != nil {
		toSerialize["Vaults"] = o.Vaults
	}
	return json.Marshal(toSerialize)
}

type NullableComponentsMonitoringSummaryData struct {
	value *ComponentsMonitoringSummaryData
	isSet bool
}

func (v NullableComponentsMonitoringSummaryData) Get() *ComponentsMonitoringSummaryData {
	return v.value
}

func (v *NullableComponentsMonitoringSummaryData) Set(val *ComponentsMonitoringSummaryData) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentsMonitoringSummaryData) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentsMonitoringSummaryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentsMonitoringSummaryData(val *ComponentsMonitoringSummaryData) *NullableComponentsMonitoringSummaryData {
	return &NullableComponentsMonitoringSummaryData{value: val, isSet: true}
}

func (v NullableComponentsMonitoringSummaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentsMonitoringSummaryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


