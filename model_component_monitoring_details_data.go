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

// ComponentMonitoringDetailsData struct for ComponentMonitoringDetailsData
type ComponentMonitoringDetailsData struct {
	ComponentIP *string `json:"ComponentIP,omitempty"`
	ComponentUserName *string `json:"ComponentUserName,omitempty"`
	ComponentVersion *string `json:"ComponentVersion,omitempty"`
	ComponentSpecificStat *int64 `json:"ComponentSpecificStat,omitempty"`
	IsLoggedOn *bool `json:"IsLoggedOn,omitempty"`
	LastLogonDate *int64 `json:"LastLogonDate,omitempty"`
}

// NewComponentMonitoringDetailsData instantiates a new ComponentMonitoringDetailsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentMonitoringDetailsData() *ComponentMonitoringDetailsData {
	this := ComponentMonitoringDetailsData{}
	return &this
}

// NewComponentMonitoringDetailsDataWithDefaults instantiates a new ComponentMonitoringDetailsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentMonitoringDetailsDataWithDefaults() *ComponentMonitoringDetailsData {
	this := ComponentMonitoringDetailsData{}
	return &this
}

// GetComponentIP returns the ComponentIP field value if set, zero value otherwise.
func (o *ComponentMonitoringDetailsData) GetComponentIP() string {
	if o == nil || o.ComponentIP == nil {
		var ret string
		return ret
	}
	return *o.ComponentIP
}

// GetComponentIPOk returns a tuple with the ComponentIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringDetailsData) GetComponentIPOk() (*string, bool) {
	if o == nil || o.ComponentIP == nil {
		return nil, false
	}
	return o.ComponentIP, true
}

// HasComponentIP returns a boolean if a field has been set.
func (o *ComponentMonitoringDetailsData) HasComponentIP() bool {
	if o != nil && o.ComponentIP != nil {
		return true
	}

	return false
}

// SetComponentIP gets a reference to the given string and assigns it to the ComponentIP field.
func (o *ComponentMonitoringDetailsData) SetComponentIP(v string) {
	o.ComponentIP = &v
}

// GetComponentUserName returns the ComponentUserName field value if set, zero value otherwise.
func (o *ComponentMonitoringDetailsData) GetComponentUserName() string {
	if o == nil || o.ComponentUserName == nil {
		var ret string
		return ret
	}
	return *o.ComponentUserName
}

// GetComponentUserNameOk returns a tuple with the ComponentUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringDetailsData) GetComponentUserNameOk() (*string, bool) {
	if o == nil || o.ComponentUserName == nil {
		return nil, false
	}
	return o.ComponentUserName, true
}

// HasComponentUserName returns a boolean if a field has been set.
func (o *ComponentMonitoringDetailsData) HasComponentUserName() bool {
	if o != nil && o.ComponentUserName != nil {
		return true
	}

	return false
}

// SetComponentUserName gets a reference to the given string and assigns it to the ComponentUserName field.
func (o *ComponentMonitoringDetailsData) SetComponentUserName(v string) {
	o.ComponentUserName = &v
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise.
func (o *ComponentMonitoringDetailsData) GetComponentVersion() string {
	if o == nil || o.ComponentVersion == nil {
		var ret string
		return ret
	}
	return *o.ComponentVersion
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringDetailsData) GetComponentVersionOk() (*string, bool) {
	if o == nil || o.ComponentVersion == nil {
		return nil, false
	}
	return o.ComponentVersion, true
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *ComponentMonitoringDetailsData) HasComponentVersion() bool {
	if o != nil && o.ComponentVersion != nil {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given string and assigns it to the ComponentVersion field.
func (o *ComponentMonitoringDetailsData) SetComponentVersion(v string) {
	o.ComponentVersion = &v
}

// GetComponentSpecificStat returns the ComponentSpecificStat field value if set, zero value otherwise.
func (o *ComponentMonitoringDetailsData) GetComponentSpecificStat() int64 {
	if o == nil || o.ComponentSpecificStat == nil {
		var ret int64
		return ret
	}
	return *o.ComponentSpecificStat
}

// GetComponentSpecificStatOk returns a tuple with the ComponentSpecificStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringDetailsData) GetComponentSpecificStatOk() (*int64, bool) {
	if o == nil || o.ComponentSpecificStat == nil {
		return nil, false
	}
	return o.ComponentSpecificStat, true
}

// HasComponentSpecificStat returns a boolean if a field has been set.
func (o *ComponentMonitoringDetailsData) HasComponentSpecificStat() bool {
	if o != nil && o.ComponentSpecificStat != nil {
		return true
	}

	return false
}

// SetComponentSpecificStat gets a reference to the given int64 and assigns it to the ComponentSpecificStat field.
func (o *ComponentMonitoringDetailsData) SetComponentSpecificStat(v int64) {
	o.ComponentSpecificStat = &v
}

// GetIsLoggedOn returns the IsLoggedOn field value if set, zero value otherwise.
func (o *ComponentMonitoringDetailsData) GetIsLoggedOn() bool {
	if o == nil || o.IsLoggedOn == nil {
		var ret bool
		return ret
	}
	return *o.IsLoggedOn
}

// GetIsLoggedOnOk returns a tuple with the IsLoggedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringDetailsData) GetIsLoggedOnOk() (*bool, bool) {
	if o == nil || o.IsLoggedOn == nil {
		return nil, false
	}
	return o.IsLoggedOn, true
}

// HasIsLoggedOn returns a boolean if a field has been set.
func (o *ComponentMonitoringDetailsData) HasIsLoggedOn() bool {
	if o != nil && o.IsLoggedOn != nil {
		return true
	}

	return false
}

// SetIsLoggedOn gets a reference to the given bool and assigns it to the IsLoggedOn field.
func (o *ComponentMonitoringDetailsData) SetIsLoggedOn(v bool) {
	o.IsLoggedOn = &v
}

// GetLastLogonDate returns the LastLogonDate field value if set, zero value otherwise.
func (o *ComponentMonitoringDetailsData) GetLastLogonDate() int64 {
	if o == nil || o.LastLogonDate == nil {
		var ret int64
		return ret
	}
	return *o.LastLogonDate
}

// GetLastLogonDateOk returns a tuple with the LastLogonDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringDetailsData) GetLastLogonDateOk() (*int64, bool) {
	if o == nil || o.LastLogonDate == nil {
		return nil, false
	}
	return o.LastLogonDate, true
}

// HasLastLogonDate returns a boolean if a field has been set.
func (o *ComponentMonitoringDetailsData) HasLastLogonDate() bool {
	if o != nil && o.LastLogonDate != nil {
		return true
	}

	return false
}

// SetLastLogonDate gets a reference to the given int64 and assigns it to the LastLogonDate field.
func (o *ComponentMonitoringDetailsData) SetLastLogonDate(v int64) {
	o.LastLogonDate = &v
}

func (o ComponentMonitoringDetailsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComponentIP != nil {
		toSerialize["ComponentIP"] = o.ComponentIP
	}
	if o.ComponentUserName != nil {
		toSerialize["ComponentUserName"] = o.ComponentUserName
	}
	if o.ComponentVersion != nil {
		toSerialize["ComponentVersion"] = o.ComponentVersion
	}
	if o.ComponentSpecificStat != nil {
		toSerialize["ComponentSpecificStat"] = o.ComponentSpecificStat
	}
	if o.IsLoggedOn != nil {
		toSerialize["IsLoggedOn"] = o.IsLoggedOn
	}
	if o.LastLogonDate != nil {
		toSerialize["LastLogonDate"] = o.LastLogonDate
	}
	return json.Marshal(toSerialize)
}

type NullableComponentMonitoringDetailsData struct {
	value *ComponentMonitoringDetailsData
	isSet bool
}

func (v NullableComponentMonitoringDetailsData) Get() *ComponentMonitoringDetailsData {
	return v.value
}

func (v *NullableComponentMonitoringDetailsData) Set(val *ComponentMonitoringDetailsData) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentMonitoringDetailsData) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentMonitoringDetailsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentMonitoringDetailsData(val *ComponentMonitoringDetailsData) *NullableComponentMonitoringDetailsData {
	return &NullableComponentMonitoringDetailsData{value: val, isSet: true}
}

func (v NullableComponentMonitoringDetailsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentMonitoringDetailsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


