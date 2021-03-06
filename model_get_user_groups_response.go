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

// GetUserGroupsResponse struct for GetUserGroupsResponse
type GetUserGroupsResponse struct {
	Value []UserGroup `json:"value,omitempty"`
	Count *int32 `json:"count,omitempty"`
	NextLink *string `json:"nextLink,omitempty"`
}

// NewGetUserGroupsResponse instantiates a new GetUserGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserGroupsResponse() *GetUserGroupsResponse {
	this := GetUserGroupsResponse{}
	return &this
}

// NewGetUserGroupsResponseWithDefaults instantiates a new GetUserGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserGroupsResponseWithDefaults() *GetUserGroupsResponse {
	this := GetUserGroupsResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetUserGroupsResponse) GetValue() []UserGroup {
	if o == nil || o.Value == nil {
		var ret []UserGroup
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGroupsResponse) GetValueOk() ([]UserGroup, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetUserGroupsResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []UserGroup and assigns it to the Value field.
func (o *GetUserGroupsResponse) SetValue(v []UserGroup) {
	o.Value = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetUserGroupsResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGroupsResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetUserGroupsResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetUserGroupsResponse) SetCount(v int32) {
	o.Count = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *GetUserGroupsResponse) GetNextLink() string {
	if o == nil || o.NextLink == nil {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGroupsResponse) GetNextLinkOk() (*string, bool) {
	if o == nil || o.NextLink == nil {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *GetUserGroupsResponse) HasNextLink() bool {
	if o != nil && o.NextLink != nil {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *GetUserGroupsResponse) SetNextLink(v string) {
	o.NextLink = &v
}

func (o GetUserGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.NextLink != nil {
		toSerialize["nextLink"] = o.NextLink
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserGroupsResponse struct {
	value *GetUserGroupsResponse
	isSet bool
}

func (v NullableGetUserGroupsResponse) Get() *GetUserGroupsResponse {
	return v.value
}

func (v *NullableGetUserGroupsResponse) Set(val *GetUserGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserGroupsResponse(val *GetUserGroupsResponse) *NullableGetUserGroupsResponse {
	return &NullableGetUserGroupsResponse{value: val, isSet: true}
}

func (v NullableGetUserGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


