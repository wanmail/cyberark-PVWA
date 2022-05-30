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

// GroupPlatform struct for GroupPlatform
type GroupPlatform struct {
	// Indicates whether a platform is active or inactive. An inactive platform cannot be assigned to accounts but will continue to manage accounts that were already assigned to it
	Active *bool `json:"Active,omitempty"`
	// Unique ID of the platform in the system
	ID *int64 `json:"ID,omitempty"`
	// Unique textual representation of the platform in the system
	PlatformID *string `json:"PlatformID,omitempty"`
	// Display name of the platform
	Name *string `json:"Name,omitempty"`
}

// NewGroupPlatform instantiates a new GroupPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPlatform() *GroupPlatform {
	this := GroupPlatform{}
	return &this
}

// NewGroupPlatformWithDefaults instantiates a new GroupPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPlatformWithDefaults() *GroupPlatform {
	this := GroupPlatform{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GroupPlatform) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPlatform) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GroupPlatform) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GroupPlatform) SetActive(v bool) {
	o.Active = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *GroupPlatform) GetID() int64 {
	if o == nil || o.ID == nil {
		var ret int64
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPlatform) GetIDOk() (*int64, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *GroupPlatform) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given int64 and assigns it to the ID field.
func (o *GroupPlatform) SetID(v int64) {
	o.ID = &v
}

// GetPlatformID returns the PlatformID field value if set, zero value otherwise.
func (o *GroupPlatform) GetPlatformID() string {
	if o == nil || o.PlatformID == nil {
		var ret string
		return ret
	}
	return *o.PlatformID
}

// GetPlatformIDOk returns a tuple with the PlatformID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPlatform) GetPlatformIDOk() (*string, bool) {
	if o == nil || o.PlatformID == nil {
		return nil, false
	}
	return o.PlatformID, true
}

// HasPlatformID returns a boolean if a field has been set.
func (o *GroupPlatform) HasPlatformID() bool {
	if o != nil && o.PlatformID != nil {
		return true
	}

	return false
}

// SetPlatformID gets a reference to the given string and assigns it to the PlatformID field.
func (o *GroupPlatform) SetPlatformID(v string) {
	o.PlatformID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupPlatform) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPlatform) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupPlatform) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupPlatform) SetName(v string) {
	o.Name = &v
}

func (o GroupPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["Active"] = o.Active
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.PlatformID != nil {
		toSerialize["PlatformID"] = o.PlatformID
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGroupPlatform struct {
	value *GroupPlatform
	isSet bool
}

func (v NullableGroupPlatform) Get() *GroupPlatform {
	return v.value
}

func (v *NullableGroupPlatform) Set(val *GroupPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPlatform(val *GroupPlatform) *NullableGroupPlatform {
	return &NullableGroupPlatform{value: val, isSet: true}
}

func (v NullableGroupPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


