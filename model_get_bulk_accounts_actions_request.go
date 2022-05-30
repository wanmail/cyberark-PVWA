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

// GetBulkAccountsActionsRequest struct for GetBulkAccountsActionsRequest
type GetBulkAccountsActionsRequest struct {
	// Maximum number of returned bulk account actions. If not specified, the default value is 10. The maximum number that can be specified is 1000.
	Limit *int32 `json:"Limit,omitempty"`
	// Search for bulk account actions filtered by spesific status. e.g. Filter=status eq InProgress
	Filter *string `json:"Filter,omitempty"`
}

// NewGetBulkAccountsActionsRequest instantiates a new GetBulkAccountsActionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBulkAccountsActionsRequest() *GetBulkAccountsActionsRequest {
	this := GetBulkAccountsActionsRequest{}
	return &this
}

// NewGetBulkAccountsActionsRequestWithDefaults instantiates a new GetBulkAccountsActionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBulkAccountsActionsRequestWithDefaults() *GetBulkAccountsActionsRequest {
	this := GetBulkAccountsActionsRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetBulkAccountsActionsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBulkAccountsActionsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetBulkAccountsActionsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetBulkAccountsActionsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GetBulkAccountsActionsRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBulkAccountsActionsRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GetBulkAccountsActionsRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GetBulkAccountsActionsRequest) SetFilter(v string) {
	o.Filter = &v
}

func (o GetBulkAccountsActionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	if o.Filter != nil {
		toSerialize["Filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableGetBulkAccountsActionsRequest struct {
	value *GetBulkAccountsActionsRequest
	isSet bool
}

func (v NullableGetBulkAccountsActionsRequest) Get() *GetBulkAccountsActionsRequest {
	return v.value
}

func (v *NullableGetBulkAccountsActionsRequest) Set(val *GetBulkAccountsActionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBulkAccountsActionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBulkAccountsActionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBulkAccountsActionsRequest(val *GetBulkAccountsActionsRequest) *NullableGetBulkAccountsActionsRequest {
	return &NullableGetBulkAccountsActionsRequest{value: val, isSet: true}
}

func (v NullableGetBulkAccountsActionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBulkAccountsActionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

