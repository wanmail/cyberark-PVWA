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

// AutomaticSecretManagement struct for AutomaticSecretManagement
type AutomaticSecretManagement struct {
	// Whether the account secret is automatically managed by the CPM. Default: true.
	AutomaticManagementEnabled *bool `json:"automaticManagementEnabled,omitempty"`
	// Reason for disabling automatic secret management
	ManualManagementReason *string `json:"manualManagementReason,omitempty"`
	// Account management status.
	Status *string `json:"status,omitempty"`
	// Last modified date of the account.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`
	// Last reconciled date of the account.
	LastReconciledTime *int64 `json:"lastReconciledTime,omitempty"`
	// Last verified date of the account.
	LastVerifiedTime *int64 `json:"lastVerifiedTime,omitempty"`
}

// NewAutomaticSecretManagement instantiates a new AutomaticSecretManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticSecretManagement() *AutomaticSecretManagement {
	this := AutomaticSecretManagement{}
	return &this
}

// NewAutomaticSecretManagementWithDefaults instantiates a new AutomaticSecretManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticSecretManagementWithDefaults() *AutomaticSecretManagement {
	this := AutomaticSecretManagement{}
	return &this
}

// GetAutomaticManagementEnabled returns the AutomaticManagementEnabled field value if set, zero value otherwise.
func (o *AutomaticSecretManagement) GetAutomaticManagementEnabled() bool {
	if o == nil || o.AutomaticManagementEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutomaticManagementEnabled
}

// GetAutomaticManagementEnabledOk returns a tuple with the AutomaticManagementEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticSecretManagement) GetAutomaticManagementEnabledOk() (*bool, bool) {
	if o == nil || o.AutomaticManagementEnabled == nil {
		return nil, false
	}
	return o.AutomaticManagementEnabled, true
}

// HasAutomaticManagementEnabled returns a boolean if a field has been set.
func (o *AutomaticSecretManagement) HasAutomaticManagementEnabled() bool {
	if o != nil && o.AutomaticManagementEnabled != nil {
		return true
	}

	return false
}

// SetAutomaticManagementEnabled gets a reference to the given bool and assigns it to the AutomaticManagementEnabled field.
func (o *AutomaticSecretManagement) SetAutomaticManagementEnabled(v bool) {
	o.AutomaticManagementEnabled = &v
}

// GetManualManagementReason returns the ManualManagementReason field value if set, zero value otherwise.
func (o *AutomaticSecretManagement) GetManualManagementReason() string {
	if o == nil || o.ManualManagementReason == nil {
		var ret string
		return ret
	}
	return *o.ManualManagementReason
}

// GetManualManagementReasonOk returns a tuple with the ManualManagementReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticSecretManagement) GetManualManagementReasonOk() (*string, bool) {
	if o == nil || o.ManualManagementReason == nil {
		return nil, false
	}
	return o.ManualManagementReason, true
}

// HasManualManagementReason returns a boolean if a field has been set.
func (o *AutomaticSecretManagement) HasManualManagementReason() bool {
	if o != nil && o.ManualManagementReason != nil {
		return true
	}

	return false
}

// SetManualManagementReason gets a reference to the given string and assigns it to the ManualManagementReason field.
func (o *AutomaticSecretManagement) SetManualManagementReason(v string) {
	o.ManualManagementReason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AutomaticSecretManagement) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticSecretManagement) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AutomaticSecretManagement) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AutomaticSecretManagement) SetStatus(v string) {
	o.Status = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *AutomaticSecretManagement) GetLastModifiedTime() int64 {
	if o == nil || o.LastModifiedTime == nil {
		var ret int64
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticSecretManagement) GetLastModifiedTimeOk() (*int64, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *AutomaticSecretManagement) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given int64 and assigns it to the LastModifiedTime field.
func (o *AutomaticSecretManagement) SetLastModifiedTime(v int64) {
	o.LastModifiedTime = &v
}

// GetLastReconciledTime returns the LastReconciledTime field value if set, zero value otherwise.
func (o *AutomaticSecretManagement) GetLastReconciledTime() int64 {
	if o == nil || o.LastReconciledTime == nil {
		var ret int64
		return ret
	}
	return *o.LastReconciledTime
}

// GetLastReconciledTimeOk returns a tuple with the LastReconciledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticSecretManagement) GetLastReconciledTimeOk() (*int64, bool) {
	if o == nil || o.LastReconciledTime == nil {
		return nil, false
	}
	return o.LastReconciledTime, true
}

// HasLastReconciledTime returns a boolean if a field has been set.
func (o *AutomaticSecretManagement) HasLastReconciledTime() bool {
	if o != nil && o.LastReconciledTime != nil {
		return true
	}

	return false
}

// SetLastReconciledTime gets a reference to the given int64 and assigns it to the LastReconciledTime field.
func (o *AutomaticSecretManagement) SetLastReconciledTime(v int64) {
	o.LastReconciledTime = &v
}

// GetLastVerifiedTime returns the LastVerifiedTime field value if set, zero value otherwise.
func (o *AutomaticSecretManagement) GetLastVerifiedTime() int64 {
	if o == nil || o.LastVerifiedTime == nil {
		var ret int64
		return ret
	}
	return *o.LastVerifiedTime
}

// GetLastVerifiedTimeOk returns a tuple with the LastVerifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticSecretManagement) GetLastVerifiedTimeOk() (*int64, bool) {
	if o == nil || o.LastVerifiedTime == nil {
		return nil, false
	}
	return o.LastVerifiedTime, true
}

// HasLastVerifiedTime returns a boolean if a field has been set.
func (o *AutomaticSecretManagement) HasLastVerifiedTime() bool {
	if o != nil && o.LastVerifiedTime != nil {
		return true
	}

	return false
}

// SetLastVerifiedTime gets a reference to the given int64 and assigns it to the LastVerifiedTime field.
func (o *AutomaticSecretManagement) SetLastVerifiedTime(v int64) {
	o.LastVerifiedTime = &v
}

func (o AutomaticSecretManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutomaticManagementEnabled != nil {
		toSerialize["automaticManagementEnabled"] = o.AutomaticManagementEnabled
	}
	if o.ManualManagementReason != nil {
		toSerialize["manualManagementReason"] = o.ManualManagementReason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.LastReconciledTime != nil {
		toSerialize["lastReconciledTime"] = o.LastReconciledTime
	}
	if o.LastVerifiedTime != nil {
		toSerialize["lastVerifiedTime"] = o.LastVerifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableAutomaticSecretManagement struct {
	value *AutomaticSecretManagement
	isSet bool
}

func (v NullableAutomaticSecretManagement) Get() *AutomaticSecretManagement {
	return v.value
}

func (v *NullableAutomaticSecretManagement) Set(val *AutomaticSecretManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticSecretManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticSecretManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticSecretManagement(val *AutomaticSecretManagement) *NullableAutomaticSecretManagement {
	return &NullableAutomaticSecretManagement{value: val, isSet: true}
}

func (v NullableAutomaticSecretManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticSecretManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


