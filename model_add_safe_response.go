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

// AddSafeResponse struct for AddSafeResponse
type AddSafeResponse struct {
	// The name of the Safe to be used when calling Safe APIs.
	SafeUrlId *string `json:"safeUrlId,omitempty"`
	// The name of the Safe.
	SafeName *string `json:"safeName,omitempty"`
	// The unique numerical ID of the Safe.
	SafeNumber *int64 `json:"safeNumber,omitempty"`
	// The description of the Safe.
	Description *string `json:"description,omitempty"`
	// The location of the Safe in the Vault.
	Location *string `json:"location,omitempty"`
	Creator *SafeCreatorResponse `json:"creator,omitempty"`
	// Whether or not to enable Object Level Access Control for the new Safe.
	OlacEnabled *bool `json:"olacEnabled,omitempty"`
	// The name of the CPM user who will manage the new Safe.
	ManagingCPM *string `json:"managingCPM,omitempty"`
	// The number of retained versions of every password that is stored in the Safe.
	NumberOfVersionsRetention *int32 `json:"numberOfVersionsRetention,omitempty"`
	// The number of days that password versions are saved in the Safe.
	NumberOfDaysRetention *int32 `json:"numberOfDaysRetention,omitempty"`
	// Whether or not to automatically purge files after the end of the Object History Retention Period defined in the Safe properties.  Reports Safes and PSM Recording Safes are created automatically with AutoPurgeEnabled set to Yes.These Safes cannot be managed by the CPM.
	AutoPurgeEnabled *bool `json:"autoPurgeEnabled,omitempty"`
	// Unix creation time of the Safe.
	CreationTime *int64 `json:"creationTime,omitempty"`
	// Unix time when the Safe was last updated.
	LastModificationTime *int64 `json:"lastModificationTime,omitempty"`
}

// NewAddSafeResponse instantiates a new AddSafeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSafeResponse() *AddSafeResponse {
	this := AddSafeResponse{}
	return &this
}

// NewAddSafeResponseWithDefaults instantiates a new AddSafeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSafeResponseWithDefaults() *AddSafeResponse {
	this := AddSafeResponse{}
	return &this
}

// GetSafeUrlId returns the SafeUrlId field value if set, zero value otherwise.
func (o *AddSafeResponse) GetSafeUrlId() string {
	if o == nil || o.SafeUrlId == nil {
		var ret string
		return ret
	}
	return *o.SafeUrlId
}

// GetSafeUrlIdOk returns a tuple with the SafeUrlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetSafeUrlIdOk() (*string, bool) {
	if o == nil || o.SafeUrlId == nil {
		return nil, false
	}
	return o.SafeUrlId, true
}

// HasSafeUrlId returns a boolean if a field has been set.
func (o *AddSafeResponse) HasSafeUrlId() bool {
	if o != nil && o.SafeUrlId != nil {
		return true
	}

	return false
}

// SetSafeUrlId gets a reference to the given string and assigns it to the SafeUrlId field.
func (o *AddSafeResponse) SetSafeUrlId(v string) {
	o.SafeUrlId = &v
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *AddSafeResponse) GetSafeName() string {
	if o == nil || o.SafeName == nil {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetSafeNameOk() (*string, bool) {
	if o == nil || o.SafeName == nil {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *AddSafeResponse) HasSafeName() bool {
	if o != nil && o.SafeName != nil {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *AddSafeResponse) SetSafeName(v string) {
	o.SafeName = &v
}

// GetSafeNumber returns the SafeNumber field value if set, zero value otherwise.
func (o *AddSafeResponse) GetSafeNumber() int64 {
	if o == nil || o.SafeNumber == nil {
		var ret int64
		return ret
	}
	return *o.SafeNumber
}

// GetSafeNumberOk returns a tuple with the SafeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetSafeNumberOk() (*int64, bool) {
	if o == nil || o.SafeNumber == nil {
		return nil, false
	}
	return o.SafeNumber, true
}

// HasSafeNumber returns a boolean if a field has been set.
func (o *AddSafeResponse) HasSafeNumber() bool {
	if o != nil && o.SafeNumber != nil {
		return true
	}

	return false
}

// SetSafeNumber gets a reference to the given int64 and assigns it to the SafeNumber field.
func (o *AddSafeResponse) SetSafeNumber(v int64) {
	o.SafeNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSafeResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSafeResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSafeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AddSafeResponse) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AddSafeResponse) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AddSafeResponse) SetLocation(v string) {
	o.Location = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AddSafeResponse) GetCreator() SafeCreatorResponse {
	if o == nil || o.Creator == nil {
		var ret SafeCreatorResponse
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetCreatorOk() (*SafeCreatorResponse, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AddSafeResponse) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given SafeCreatorResponse and assigns it to the Creator field.
func (o *AddSafeResponse) SetCreator(v SafeCreatorResponse) {
	o.Creator = &v
}

// GetOlacEnabled returns the OlacEnabled field value if set, zero value otherwise.
func (o *AddSafeResponse) GetOlacEnabled() bool {
	if o == nil || o.OlacEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OlacEnabled
}

// GetOlacEnabledOk returns a tuple with the OlacEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetOlacEnabledOk() (*bool, bool) {
	if o == nil || o.OlacEnabled == nil {
		return nil, false
	}
	return o.OlacEnabled, true
}

// HasOlacEnabled returns a boolean if a field has been set.
func (o *AddSafeResponse) HasOlacEnabled() bool {
	if o != nil && o.OlacEnabled != nil {
		return true
	}

	return false
}

// SetOlacEnabled gets a reference to the given bool and assigns it to the OlacEnabled field.
func (o *AddSafeResponse) SetOlacEnabled(v bool) {
	o.OlacEnabled = &v
}

// GetManagingCPM returns the ManagingCPM field value if set, zero value otherwise.
func (o *AddSafeResponse) GetManagingCPM() string {
	if o == nil || o.ManagingCPM == nil {
		var ret string
		return ret
	}
	return *o.ManagingCPM
}

// GetManagingCPMOk returns a tuple with the ManagingCPM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetManagingCPMOk() (*string, bool) {
	if o == nil || o.ManagingCPM == nil {
		return nil, false
	}
	return o.ManagingCPM, true
}

// HasManagingCPM returns a boolean if a field has been set.
func (o *AddSafeResponse) HasManagingCPM() bool {
	if o != nil && o.ManagingCPM != nil {
		return true
	}

	return false
}

// SetManagingCPM gets a reference to the given string and assigns it to the ManagingCPM field.
func (o *AddSafeResponse) SetManagingCPM(v string) {
	o.ManagingCPM = &v
}

// GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field value if set, zero value otherwise.
func (o *AddSafeResponse) GetNumberOfVersionsRetention() int32 {
	if o == nil || o.NumberOfVersionsRetention == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfVersionsRetention
}

// GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetNumberOfVersionsRetentionOk() (*int32, bool) {
	if o == nil || o.NumberOfVersionsRetention == nil {
		return nil, false
	}
	return o.NumberOfVersionsRetention, true
}

// HasNumberOfVersionsRetention returns a boolean if a field has been set.
func (o *AddSafeResponse) HasNumberOfVersionsRetention() bool {
	if o != nil && o.NumberOfVersionsRetention != nil {
		return true
	}

	return false
}

// SetNumberOfVersionsRetention gets a reference to the given int32 and assigns it to the NumberOfVersionsRetention field.
func (o *AddSafeResponse) SetNumberOfVersionsRetention(v int32) {
	o.NumberOfVersionsRetention = &v
}

// GetNumberOfDaysRetention returns the NumberOfDaysRetention field value if set, zero value otherwise.
func (o *AddSafeResponse) GetNumberOfDaysRetention() int32 {
	if o == nil || o.NumberOfDaysRetention == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfDaysRetention
}

// GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetNumberOfDaysRetentionOk() (*int32, bool) {
	if o == nil || o.NumberOfDaysRetention == nil {
		return nil, false
	}
	return o.NumberOfDaysRetention, true
}

// HasNumberOfDaysRetention returns a boolean if a field has been set.
func (o *AddSafeResponse) HasNumberOfDaysRetention() bool {
	if o != nil && o.NumberOfDaysRetention != nil {
		return true
	}

	return false
}

// SetNumberOfDaysRetention gets a reference to the given int32 and assigns it to the NumberOfDaysRetention field.
func (o *AddSafeResponse) SetNumberOfDaysRetention(v int32) {
	o.NumberOfDaysRetention = &v
}

// GetAutoPurgeEnabled returns the AutoPurgeEnabled field value if set, zero value otherwise.
func (o *AddSafeResponse) GetAutoPurgeEnabled() bool {
	if o == nil || o.AutoPurgeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoPurgeEnabled
}

// GetAutoPurgeEnabledOk returns a tuple with the AutoPurgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetAutoPurgeEnabledOk() (*bool, bool) {
	if o == nil || o.AutoPurgeEnabled == nil {
		return nil, false
	}
	return o.AutoPurgeEnabled, true
}

// HasAutoPurgeEnabled returns a boolean if a field has been set.
func (o *AddSafeResponse) HasAutoPurgeEnabled() bool {
	if o != nil && o.AutoPurgeEnabled != nil {
		return true
	}

	return false
}

// SetAutoPurgeEnabled gets a reference to the given bool and assigns it to the AutoPurgeEnabled field.
func (o *AddSafeResponse) SetAutoPurgeEnabled(v bool) {
	o.AutoPurgeEnabled = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AddSafeResponse) GetCreationTime() int64 {
	if o == nil || o.CreationTime == nil {
		var ret int64
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetCreationTimeOk() (*int64, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AddSafeResponse) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given int64 and assigns it to the CreationTime field.
func (o *AddSafeResponse) SetCreationTime(v int64) {
	o.CreationTime = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *AddSafeResponse) GetLastModificationTime() int64 {
	if o == nil || o.LastModificationTime == nil {
		var ret int64
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeResponse) GetLastModificationTimeOk() (*int64, bool) {
	if o == nil || o.LastModificationTime == nil {
		return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *AddSafeResponse) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime != nil {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given int64 and assigns it to the LastModificationTime field.
func (o *AddSafeResponse) SetLastModificationTime(v int64) {
	o.LastModificationTime = &v
}

func (o AddSafeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafeUrlId != nil {
		toSerialize["safeUrlId"] = o.SafeUrlId
	}
	if o.SafeName != nil {
		toSerialize["safeName"] = o.SafeName
	}
	if o.SafeNumber != nil {
		toSerialize["safeNumber"] = o.SafeNumber
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.OlacEnabled != nil {
		toSerialize["olacEnabled"] = o.OlacEnabled
	}
	if o.ManagingCPM != nil {
		toSerialize["managingCPM"] = o.ManagingCPM
	}
	if o.NumberOfVersionsRetention != nil {
		toSerialize["numberOfVersionsRetention"] = o.NumberOfVersionsRetention
	}
	if o.NumberOfDaysRetention != nil {
		toSerialize["numberOfDaysRetention"] = o.NumberOfDaysRetention
	}
	if o.AutoPurgeEnabled != nil {
		toSerialize["autoPurgeEnabled"] = o.AutoPurgeEnabled
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.LastModificationTime != nil {
		toSerialize["lastModificationTime"] = o.LastModificationTime
	}
	return json.Marshal(toSerialize)
}

type NullableAddSafeResponse struct {
	value *AddSafeResponse
	isSet bool
}

func (v NullableAddSafeResponse) Get() *AddSafeResponse {
	return v.value
}

func (v *NullableAddSafeResponse) Set(val *AddSafeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSafeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSafeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSafeResponse(val *AddSafeResponse) *NullableAddSafeResponse {
	return &NullableAddSafeResponse{value: val, isSet: true}
}

func (v NullableAddSafeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSafeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


