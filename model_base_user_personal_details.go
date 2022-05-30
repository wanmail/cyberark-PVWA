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

// BaseUserPersonalDetails struct for BaseUserPersonalDetails
type BaseUserPersonalDetails struct {
	FirstName *string `json:"firstName,omitempty"`
	MiddleName *string `json:"middleName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Department *string `json:"department,omitempty"`
}

// NewBaseUserPersonalDetails instantiates a new BaseUserPersonalDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseUserPersonalDetails() *BaseUserPersonalDetails {
	this := BaseUserPersonalDetails{}
	return &this
}

// NewBaseUserPersonalDetailsWithDefaults instantiates a new BaseUserPersonalDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseUserPersonalDetailsWithDefaults() *BaseUserPersonalDetails {
	this := BaseUserPersonalDetails{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BaseUserPersonalDetails) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUserPersonalDetails) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BaseUserPersonalDetails) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BaseUserPersonalDetails) SetFirstName(v string) {
	o.FirstName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *BaseUserPersonalDetails) GetMiddleName() string {
	if o == nil || o.MiddleName == nil {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUserPersonalDetails) GetMiddleNameOk() (*string, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *BaseUserPersonalDetails) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *BaseUserPersonalDetails) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BaseUserPersonalDetails) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUserPersonalDetails) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BaseUserPersonalDetails) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BaseUserPersonalDetails) SetLastName(v string) {
	o.LastName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BaseUserPersonalDetails) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUserPersonalDetails) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BaseUserPersonalDetails) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *BaseUserPersonalDetails) SetOrganization(v string) {
	o.Organization = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *BaseUserPersonalDetails) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseUserPersonalDetails) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *BaseUserPersonalDetails) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *BaseUserPersonalDetails) SetDepartment(v string) {
	o.Department = &v
}

func (o BaseUserPersonalDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.MiddleName != nil {
		toSerialize["middleName"] = o.MiddleName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	return json.Marshal(toSerialize)
}

type NullableBaseUserPersonalDetails struct {
	value *BaseUserPersonalDetails
	isSet bool
}

func (v NullableBaseUserPersonalDetails) Get() *BaseUserPersonalDetails {
	return v.value
}

func (v *NullableBaseUserPersonalDetails) Set(val *BaseUserPersonalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseUserPersonalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseUserPersonalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseUserPersonalDetails(val *BaseUserPersonalDetails) *NullableBaseUserPersonalDetails {
	return &NullableBaseUserPersonalDetails{value: val, isSet: true}
}

func (v NullableBaseUserPersonalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseUserPersonalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


