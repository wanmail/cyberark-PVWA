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

// ConfirmerStatus struct for ConfirmerStatus
type ConfirmerStatus struct {
	Type *int32 `json:"Type,omitempty"`
	ID *int64 `json:"ID,omitempty"`
	Name *string `json:"Name,omitempty"`
	Action *int32 `json:"Action,omitempty"`
	Reason *string `json:"Reason,omitempty"`
	ActionDate *int64 `json:"ActionDate,omitempty"`
	AdditionalDetails *map[string]string `json:"AdditionalDetails,omitempty"`
	Members []ConfirmerMember `json:"Members,omitempty"`
}

// NewConfirmerStatus instantiates a new ConfirmerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmerStatus() *ConfirmerStatus {
	this := ConfirmerStatus{}
	return &this
}

// NewConfirmerStatusWithDefaults instantiates a new ConfirmerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmerStatusWithDefaults() *ConfirmerStatus {
	this := ConfirmerStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *ConfirmerStatus) SetType(v int32) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetID() int64 {
	if o == nil || o.ID == nil {
		var ret int64
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetIDOk() (*int64, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given int64 and assigns it to the ID field.
func (o *ConfirmerStatus) SetID(v int64) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfirmerStatus) SetName(v string) {
	o.Name = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetAction() int32 {
	if o == nil || o.Action == nil {
		var ret int32
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetActionOk() (*int32, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given int32 and assigns it to the Action field.
func (o *ConfirmerStatus) SetAction(v int32) {
	o.Action = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ConfirmerStatus) SetReason(v string) {
	o.Reason = &v
}

// GetActionDate returns the ActionDate field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetActionDate() int64 {
	if o == nil || o.ActionDate == nil {
		var ret int64
		return ret
	}
	return *o.ActionDate
}

// GetActionDateOk returns a tuple with the ActionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetActionDateOk() (*int64, bool) {
	if o == nil || o.ActionDate == nil {
		return nil, false
	}
	return o.ActionDate, true
}

// HasActionDate returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasActionDate() bool {
	if o != nil && o.ActionDate != nil {
		return true
	}

	return false
}

// SetActionDate gets a reference to the given int64 and assigns it to the ActionDate field.
func (o *ConfirmerStatus) SetActionDate(v int64) {
	o.ActionDate = &v
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetAdditionalDetails() map[string]string {
	if o == nil || o.AdditionalDetails == nil {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetAdditionalDetailsOk() (*map[string]string, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given map[string]string and assigns it to the AdditionalDetails field.
func (o *ConfirmerStatus) SetAdditionalDetails(v map[string]string) {
	o.AdditionalDetails = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ConfirmerStatus) GetMembers() []ConfirmerMember {
	if o == nil || o.Members == nil {
		var ret []ConfirmerMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmerStatus) GetMembersOk() ([]ConfirmerMember, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ConfirmerStatus) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ConfirmerMember and assigns it to the Members field.
func (o *ConfirmerStatus) SetMembers(v []ConfirmerMember) {
	o.Members = v
}

func (o ConfirmerStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.ActionDate != nil {
		toSerialize["ActionDate"] = o.ActionDate
	}
	if o.AdditionalDetails != nil {
		toSerialize["AdditionalDetails"] = o.AdditionalDetails
	}
	if o.Members != nil {
		toSerialize["Members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableConfirmerStatus struct {
	value *ConfirmerStatus
	isSet bool
}

func (v NullableConfirmerStatus) Get() *ConfirmerStatus {
	return v.value
}

func (v *NullableConfirmerStatus) Set(val *ConfirmerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmerStatus(val *ConfirmerStatus) *NullableConfirmerStatus {
	return &NullableConfirmerStatus{value: val, isSet: true}
}

func (v NullableConfirmerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


