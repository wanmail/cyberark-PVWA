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

// UpdateSafeRequestBody struct for UpdateSafeRequestBody
type UpdateSafeRequestBody struct {
	// The name of the Safe.
	SafeName *string `json:"SafeName,omitempty"`
	// The description of the Safe.
	Description *string `json:"Description,omitempty"`
	// The location of the Safe in the Vault.
	Location *string `json:"Location,omitempty"`
	// The number of days that password versions are saved in the Safe.
	NumberOfDaysRetention *int32 `json:"NumberOfDaysRetention,omitempty"`
	// The number of retained versions of every password that is stored in the Safe.
	NumberOfVersionsRetention *int32 `json:"NumberOfVersionsRetention,omitempty"`
	// Whether or not to enable Object Level Access Control for the Safe.
	OLACEnabled *bool `json:"OLACEnabled,omitempty"`
	// The name of the CPM user who will manage the Safe.
	ManagingCPM *string `json:"ManagingCPM,omitempty"`
}

// NewUpdateSafeRequestBody instantiates a new UpdateSafeRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSafeRequestBody() *UpdateSafeRequestBody {
	this := UpdateSafeRequestBody{}
	return &this
}

// NewUpdateSafeRequestBodyWithDefaults instantiates a new UpdateSafeRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSafeRequestBodyWithDefaults() *UpdateSafeRequestBody {
	this := UpdateSafeRequestBody{}
	return &this
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetSafeName() string {
	if o == nil || o.SafeName == nil {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetSafeNameOk() (*string, bool) {
	if o == nil || o.SafeName == nil {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasSafeName() bool {
	if o != nil && o.SafeName != nil {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *UpdateSafeRequestBody) SetSafeName(v string) {
	o.SafeName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateSafeRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UpdateSafeRequestBody) SetLocation(v string) {
	o.Location = &v
}

// GetNumberOfDaysRetention returns the NumberOfDaysRetention field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetNumberOfDaysRetention() int32 {
	if o == nil || o.NumberOfDaysRetention == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfDaysRetention
}

// GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetNumberOfDaysRetentionOk() (*int32, bool) {
	if o == nil || o.NumberOfDaysRetention == nil {
		return nil, false
	}
	return o.NumberOfDaysRetention, true
}

// HasNumberOfDaysRetention returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasNumberOfDaysRetention() bool {
	if o != nil && o.NumberOfDaysRetention != nil {
		return true
	}

	return false
}

// SetNumberOfDaysRetention gets a reference to the given int32 and assigns it to the NumberOfDaysRetention field.
func (o *UpdateSafeRequestBody) SetNumberOfDaysRetention(v int32) {
	o.NumberOfDaysRetention = &v
}

// GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetNumberOfVersionsRetention() int32 {
	if o == nil || o.NumberOfVersionsRetention == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfVersionsRetention
}

// GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetNumberOfVersionsRetentionOk() (*int32, bool) {
	if o == nil || o.NumberOfVersionsRetention == nil {
		return nil, false
	}
	return o.NumberOfVersionsRetention, true
}

// HasNumberOfVersionsRetention returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasNumberOfVersionsRetention() bool {
	if o != nil && o.NumberOfVersionsRetention != nil {
		return true
	}

	return false
}

// SetNumberOfVersionsRetention gets a reference to the given int32 and assigns it to the NumberOfVersionsRetention field.
func (o *UpdateSafeRequestBody) SetNumberOfVersionsRetention(v int32) {
	o.NumberOfVersionsRetention = &v
}

// GetOLACEnabled returns the OLACEnabled field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetOLACEnabled() bool {
	if o == nil || o.OLACEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OLACEnabled
}

// GetOLACEnabledOk returns a tuple with the OLACEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetOLACEnabledOk() (*bool, bool) {
	if o == nil || o.OLACEnabled == nil {
		return nil, false
	}
	return o.OLACEnabled, true
}

// HasOLACEnabled returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasOLACEnabled() bool {
	if o != nil && o.OLACEnabled != nil {
		return true
	}

	return false
}

// SetOLACEnabled gets a reference to the given bool and assigns it to the OLACEnabled field.
func (o *UpdateSafeRequestBody) SetOLACEnabled(v bool) {
	o.OLACEnabled = &v
}

// GetManagingCPM returns the ManagingCPM field value if set, zero value otherwise.
func (o *UpdateSafeRequestBody) GetManagingCPM() string {
	if o == nil || o.ManagingCPM == nil {
		var ret string
		return ret
	}
	return *o.ManagingCPM
}

// GetManagingCPMOk returns a tuple with the ManagingCPM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSafeRequestBody) GetManagingCPMOk() (*string, bool) {
	if o == nil || o.ManagingCPM == nil {
		return nil, false
	}
	return o.ManagingCPM, true
}

// HasManagingCPM returns a boolean if a field has been set.
func (o *UpdateSafeRequestBody) HasManagingCPM() bool {
	if o != nil && o.ManagingCPM != nil {
		return true
	}

	return false
}

// SetManagingCPM gets a reference to the given string and assigns it to the ManagingCPM field.
func (o *UpdateSafeRequestBody) SetManagingCPM(v string) {
	o.ManagingCPM = &v
}

func (o UpdateSafeRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafeName != nil {
		toSerialize["SafeName"] = o.SafeName
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.NumberOfDaysRetention != nil {
		toSerialize["NumberOfDaysRetention"] = o.NumberOfDaysRetention
	}
	if o.NumberOfVersionsRetention != nil {
		toSerialize["NumberOfVersionsRetention"] = o.NumberOfVersionsRetention
	}
	if o.OLACEnabled != nil {
		toSerialize["OLACEnabled"] = o.OLACEnabled
	}
	if o.ManagingCPM != nil {
		toSerialize["ManagingCPM"] = o.ManagingCPM
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSafeRequestBody struct {
	value *UpdateSafeRequestBody
	isSet bool
}

func (v NullableUpdateSafeRequestBody) Get() *UpdateSafeRequestBody {
	return v.value
}

func (v *NullableUpdateSafeRequestBody) Set(val *UpdateSafeRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSafeRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSafeRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSafeRequestBody(val *UpdateSafeRequestBody) *NullableUpdateSafeRequestBody {
	return &NullableUpdateSafeRequestBody{value: val, isSet: true}
}

func (v NullableUpdateSafeRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSafeRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


