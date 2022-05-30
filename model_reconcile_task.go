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

// ReconcileTask struct for ReconcileTask
type ReconcileTask struct {
	// Whether or not passwords will be reconciled automatically after the CPM detects a password on remote machine that is not synchronized with its corresponding in the server
	AutomaticReconcileWhenUnsynced *bool `json:"AutomaticReconcileWhenUnsynced,omitempty"`
	// Indicates whether the action can be initiated manually.
	AllowManual *bool `json:"AllowManual,omitempty"`
}

// NewReconcileTask instantiates a new ReconcileTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReconcileTask() *ReconcileTask {
	this := ReconcileTask{}
	return &this
}

// NewReconcileTaskWithDefaults instantiates a new ReconcileTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReconcileTaskWithDefaults() *ReconcileTask {
	this := ReconcileTask{}
	return &this
}

// GetAutomaticReconcileWhenUnsynced returns the AutomaticReconcileWhenUnsynced field value if set, zero value otherwise.
func (o *ReconcileTask) GetAutomaticReconcileWhenUnsynced() bool {
	if o == nil || o.AutomaticReconcileWhenUnsynced == nil {
		var ret bool
		return ret
	}
	return *o.AutomaticReconcileWhenUnsynced
}

// GetAutomaticReconcileWhenUnsyncedOk returns a tuple with the AutomaticReconcileWhenUnsynced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReconcileTask) GetAutomaticReconcileWhenUnsyncedOk() (*bool, bool) {
	if o == nil || o.AutomaticReconcileWhenUnsynced == nil {
		return nil, false
	}
	return o.AutomaticReconcileWhenUnsynced, true
}

// HasAutomaticReconcileWhenUnsynced returns a boolean if a field has been set.
func (o *ReconcileTask) HasAutomaticReconcileWhenUnsynced() bool {
	if o != nil && o.AutomaticReconcileWhenUnsynced != nil {
		return true
	}

	return false
}

// SetAutomaticReconcileWhenUnsynced gets a reference to the given bool and assigns it to the AutomaticReconcileWhenUnsynced field.
func (o *ReconcileTask) SetAutomaticReconcileWhenUnsynced(v bool) {
	o.AutomaticReconcileWhenUnsynced = &v
}

// GetAllowManual returns the AllowManual field value if set, zero value otherwise.
func (o *ReconcileTask) GetAllowManual() bool {
	if o == nil || o.AllowManual == nil {
		var ret bool
		return ret
	}
	return *o.AllowManual
}

// GetAllowManualOk returns a tuple with the AllowManual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReconcileTask) GetAllowManualOk() (*bool, bool) {
	if o == nil || o.AllowManual == nil {
		return nil, false
	}
	return o.AllowManual, true
}

// HasAllowManual returns a boolean if a field has been set.
func (o *ReconcileTask) HasAllowManual() bool {
	if o != nil && o.AllowManual != nil {
		return true
	}

	return false
}

// SetAllowManual gets a reference to the given bool and assigns it to the AllowManual field.
func (o *ReconcileTask) SetAllowManual(v bool) {
	o.AllowManual = &v
}

func (o ReconcileTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutomaticReconcileWhenUnsynced != nil {
		toSerialize["AutomaticReconcileWhenUnsynced"] = o.AutomaticReconcileWhenUnsynced
	}
	if o.AllowManual != nil {
		toSerialize["AllowManual"] = o.AllowManual
	}
	return json.Marshal(toSerialize)
}

type NullableReconcileTask struct {
	value *ReconcileTask
	isSet bool
}

func (v NullableReconcileTask) Get() *ReconcileTask {
	return v.value
}

func (v *NullableReconcileTask) Set(val *ReconcileTask) {
	v.value = val
	v.isSet = true
}

func (v NullableReconcileTask) IsSet() bool {
	return v.isSet
}

func (v *NullableReconcileTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReconcileTask(val *ReconcileTask) *NullableReconcileTask {
	return &NullableReconcileTask{value: val, isSet: true}
}

func (v NullableReconcileTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReconcileTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

