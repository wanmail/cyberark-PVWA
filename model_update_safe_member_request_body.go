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

// UpdateSafeMemberRequestBody struct for UpdateSafeMemberRequestBody
type UpdateSafeMemberRequestBody struct {
	MembershipExpirationDate *int64 `json:"MembershipExpirationDate,omitempty"`
	Permissions map[string]bool `json:"Permissions"`
}

// NewUpdateSafeMemberRequestBody instantiates a new UpdateSafeMemberRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSafeMemberRequestBody(permissions map[string]bool) *UpdateSafeMemberRequestBody {
	this := UpdateSafeMemberRequestBody{}
	this.Permissions = permissions
	return &this
}

// NewUpdateSafeMemberRequestBodyWithDefaults instantiates a new UpdateSafeMemberRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSafeMemberRequestBodyWithDefaults() *UpdateSafeMemberRequestBody {
	this := UpdateSafeMemberRequestBody{}
	return &this
}

// GetMembershipExpirationDate returns the MembershipExpirationDate field value if set, zero value otherwise.
func (o *UpdateSafeMemberRequestBody) GetMembershipExpirationDate() int64 {
	if o == nil || o.MembershipExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.MembershipExpirationDate
}

// GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeMemberRequestBody) GetMembershipExpirationDateOk() (*int64, bool) {
	if o == nil || o.MembershipExpirationDate == nil {
		return nil, false
	}
	return o.MembershipExpirationDate, true
}

// HasMembershipExpirationDate returns a boolean if a field has been set.
func (o *UpdateSafeMemberRequestBody) HasMembershipExpirationDate() bool {
	if o != nil && o.MembershipExpirationDate != nil {
		return true
	}

	return false
}

// SetMembershipExpirationDate gets a reference to the given int64 and assigns it to the MembershipExpirationDate field.
func (o *UpdateSafeMemberRequestBody) SetMembershipExpirationDate(v int64) {
	o.MembershipExpirationDate = &v
}

// GetPermissions returns the Permissions field value
func (o *UpdateSafeMemberRequestBody) GetPermissions() map[string]bool {
	if o == nil {
		var ret map[string]bool
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UpdateSafeMemberRequestBody) GetPermissionsOk() (*map[string]bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *UpdateSafeMemberRequestBody) SetPermissions(v map[string]bool) {
	o.Permissions = v
}

func (o UpdateSafeMemberRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MembershipExpirationDate != nil {
		toSerialize["MembershipExpirationDate"] = o.MembershipExpirationDate
	}
	if true {
		toSerialize["Permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSafeMemberRequestBody struct {
	value *UpdateSafeMemberRequestBody
	isSet bool
}

func (v NullableUpdateSafeMemberRequestBody) Get() *UpdateSafeMemberRequestBody {
	return v.value
}

func (v *NullableUpdateSafeMemberRequestBody) Set(val *UpdateSafeMemberRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSafeMemberRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSafeMemberRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSafeMemberRequestBody(val *UpdateSafeMemberRequestBody) *NullableUpdateSafeMemberRequestBody {
	return &NullableUpdateSafeMemberRequestBody{value: val, isSet: true}
}

func (v NullableUpdateSafeMemberRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSafeMemberRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


