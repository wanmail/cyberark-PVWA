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

// WorkflowPolicy struct for WorkflowPolicy
type WorkflowPolicy struct {
	// Indicates whether the policy is active.
	IsActive *bool `json:"IsActive,omitempty"`
	// Indicates whether the policy is an exception to the master policy.
	IsAnException *bool `json:"IsAnException,omitempty"`
}

// NewWorkflowPolicy instantiates a new WorkflowPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPolicy() *WorkflowPolicy {
	this := WorkflowPolicy{}
	return &this
}

// NewWorkflowPolicyWithDefaults instantiates a new WorkflowPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPolicyWithDefaults() *WorkflowPolicy {
	this := WorkflowPolicy{}
	return &this
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *WorkflowPolicy) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPolicy) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *WorkflowPolicy) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *WorkflowPolicy) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsAnException returns the IsAnException field value if set, zero value otherwise.
func (o *WorkflowPolicy) GetIsAnException() bool {
	if o == nil || o.IsAnException == nil {
		var ret bool
		return ret
	}
	return *o.IsAnException
}

// GetIsAnExceptionOk returns a tuple with the IsAnException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPolicy) GetIsAnExceptionOk() (*bool, bool) {
	if o == nil || o.IsAnException == nil {
		return nil, false
	}
	return o.IsAnException, true
}

// HasIsAnException returns a boolean if a field has been set.
func (o *WorkflowPolicy) HasIsAnException() bool {
	if o != nil && o.IsAnException != nil {
		return true
	}

	return false
}

// SetIsAnException gets a reference to the given bool and assigns it to the IsAnException field.
func (o *WorkflowPolicy) SetIsAnException(v bool) {
	o.IsAnException = &v
}

func (o WorkflowPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}
	if o.IsAnException != nil {
		toSerialize["IsAnException"] = o.IsAnException
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowPolicy struct {
	value *WorkflowPolicy
	isSet bool
}

func (v NullableWorkflowPolicy) Get() *WorkflowPolicy {
	return v.value
}

func (v *NullableWorkflowPolicy) Set(val *WorkflowPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPolicy(val *WorkflowPolicy) *NullableWorkflowPolicy {
	return &NullableWorkflowPolicy{value: val, isSet: true}
}

func (v NullableWorkflowPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


