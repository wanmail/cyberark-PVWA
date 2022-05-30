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

// UserPhone struct for UserPhone
type UserPhone struct {
	HomeNumber *string `json:"homeNumber,omitempty"`
	BusinessNumber *string `json:"businessNumber,omitempty"`
	CellularNumber *string `json:"cellularNumber,omitempty"`
	FaxNumber *string `json:"faxNumber,omitempty"`
	PagerNumber *string `json:"pagerNumber,omitempty"`
}

// NewUserPhone instantiates a new UserPhone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPhone() *UserPhone {
	this := UserPhone{}
	return &this
}

// NewUserPhoneWithDefaults instantiates a new UserPhone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPhoneWithDefaults() *UserPhone {
	this := UserPhone{}
	return &this
}

// GetHomeNumber returns the HomeNumber field value if set, zero value otherwise.
func (o *UserPhone) GetHomeNumber() string {
	if o == nil || o.HomeNumber == nil {
		var ret string
		return ret
	}
	return *o.HomeNumber
}

// GetHomeNumberOk returns a tuple with the HomeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPhone) GetHomeNumberOk() (*string, bool) {
	if o == nil || o.HomeNumber == nil {
		return nil, false
	}
	return o.HomeNumber, true
}

// HasHomeNumber returns a boolean if a field has been set.
func (o *UserPhone) HasHomeNumber() bool {
	if o != nil && o.HomeNumber != nil {
		return true
	}

	return false
}

// SetHomeNumber gets a reference to the given string and assigns it to the HomeNumber field.
func (o *UserPhone) SetHomeNumber(v string) {
	o.HomeNumber = &v
}

// GetBusinessNumber returns the BusinessNumber field value if set, zero value otherwise.
func (o *UserPhone) GetBusinessNumber() string {
	if o == nil || o.BusinessNumber == nil {
		var ret string
		return ret
	}
	return *o.BusinessNumber
}

// GetBusinessNumberOk returns a tuple with the BusinessNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPhone) GetBusinessNumberOk() (*string, bool) {
	if o == nil || o.BusinessNumber == nil {
		return nil, false
	}
	return o.BusinessNumber, true
}

// HasBusinessNumber returns a boolean if a field has been set.
func (o *UserPhone) HasBusinessNumber() bool {
	if o != nil && o.BusinessNumber != nil {
		return true
	}

	return false
}

// SetBusinessNumber gets a reference to the given string and assigns it to the BusinessNumber field.
func (o *UserPhone) SetBusinessNumber(v string) {
	o.BusinessNumber = &v
}

// GetCellularNumber returns the CellularNumber field value if set, zero value otherwise.
func (o *UserPhone) GetCellularNumber() string {
	if o == nil || o.CellularNumber == nil {
		var ret string
		return ret
	}
	return *o.CellularNumber
}

// GetCellularNumberOk returns a tuple with the CellularNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPhone) GetCellularNumberOk() (*string, bool) {
	if o == nil || o.CellularNumber == nil {
		return nil, false
	}
	return o.CellularNumber, true
}

// HasCellularNumber returns a boolean if a field has been set.
func (o *UserPhone) HasCellularNumber() bool {
	if o != nil && o.CellularNumber != nil {
		return true
	}

	return false
}

// SetCellularNumber gets a reference to the given string and assigns it to the CellularNumber field.
func (o *UserPhone) SetCellularNumber(v string) {
	o.CellularNumber = &v
}

// GetFaxNumber returns the FaxNumber field value if set, zero value otherwise.
func (o *UserPhone) GetFaxNumber() string {
	if o == nil || o.FaxNumber == nil {
		var ret string
		return ret
	}
	return *o.FaxNumber
}

// GetFaxNumberOk returns a tuple with the FaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPhone) GetFaxNumberOk() (*string, bool) {
	if o == nil || o.FaxNumber == nil {
		return nil, false
	}
	return o.FaxNumber, true
}

// HasFaxNumber returns a boolean if a field has been set.
func (o *UserPhone) HasFaxNumber() bool {
	if o != nil && o.FaxNumber != nil {
		return true
	}

	return false
}

// SetFaxNumber gets a reference to the given string and assigns it to the FaxNumber field.
func (o *UserPhone) SetFaxNumber(v string) {
	o.FaxNumber = &v
}

// GetPagerNumber returns the PagerNumber field value if set, zero value otherwise.
func (o *UserPhone) GetPagerNumber() string {
	if o == nil || o.PagerNumber == nil {
		var ret string
		return ret
	}
	return *o.PagerNumber
}

// GetPagerNumberOk returns a tuple with the PagerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPhone) GetPagerNumberOk() (*string, bool) {
	if o == nil || o.PagerNumber == nil {
		return nil, false
	}
	return o.PagerNumber, true
}

// HasPagerNumber returns a boolean if a field has been set.
func (o *UserPhone) HasPagerNumber() bool {
	if o != nil && o.PagerNumber != nil {
		return true
	}

	return false
}

// SetPagerNumber gets a reference to the given string and assigns it to the PagerNumber field.
func (o *UserPhone) SetPagerNumber(v string) {
	o.PagerNumber = &v
}

func (o UserPhone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HomeNumber != nil {
		toSerialize["homeNumber"] = o.HomeNumber
	}
	if o.BusinessNumber != nil {
		toSerialize["businessNumber"] = o.BusinessNumber
	}
	if o.CellularNumber != nil {
		toSerialize["cellularNumber"] = o.CellularNumber
	}
	if o.FaxNumber != nil {
		toSerialize["faxNumber"] = o.FaxNumber
	}
	if o.PagerNumber != nil {
		toSerialize["pagerNumber"] = o.PagerNumber
	}
	return json.Marshal(toSerialize)
}

type NullableUserPhone struct {
	value *UserPhone
	isSet bool
}

func (v NullableUserPhone) Get() *UserPhone {
	return v.value
}

func (v *NullableUserPhone) Set(val *UserPhone) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPhone) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPhone(val *UserPhone) *NullableUserPhone {
	return &NullableUserPhone{value: val, isSet: true}
}

func (v NullableUserPhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


