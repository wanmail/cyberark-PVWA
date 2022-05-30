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

// GetAccountsRequest struct for GetAccountsRequest
type GetAccountsRequest struct {
	// List of keywords separated with space to search in accounts.
	Search *string `json:"search,omitempty"`
	// Property or properties to sort returned accounts, followed by asc (default) or desc to control sort direction. Multiple sorts are comma-separated. Maximum number of properties is 3.
	Sort []string `json:"sort,omitempty"`
	// Offset of the first returned account into the collection of results.
	Offset *int32 `json:"offset,omitempty"`
	// Maximum number of returned accounts. If not specified, the default value is 50. The maximum number that can be specified is 1000.
	Limit *int32 `json:"limit,omitempty"`
	// Search for accounts filtered by specific Safe. Filter=safename eq [safe name]
	Filter *string `json:"filter,omitempty"`
	// Type of search to perform - if keywords are contained in the relevant account properties values or in the start of the properties values (the later enhances performance)
	SearchType *string `json:"searchType,omitempty"`
}

// NewGetAccountsRequest instantiates a new GetAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountsRequest() *GetAccountsRequest {
	this := GetAccountsRequest{}
	return &this
}

// NewGetAccountsRequestWithDefaults instantiates a new GetAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountsRequestWithDefaults() *GetAccountsRequest {
	this := GetAccountsRequest{}
	return &this
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *GetAccountsRequest) GetSearch() string {
	if o == nil || o.Search == nil {
		var ret string
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsRequest) GetSearchOk() (*string, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *GetAccountsRequest) HasSearch() bool {
	if o != nil && o.Search != nil {
		return true
	}

	return false
}

// SetSearch gets a reference to the given string and assigns it to the Search field.
func (o *GetAccountsRequest) SetSearch(v string) {
	o.Search = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *GetAccountsRequest) GetSort() []string {
	if o == nil || o.Sort == nil {
		var ret []string
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsRequest) GetSortOk() ([]string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *GetAccountsRequest) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []string and assigns it to the Sort field.
func (o *GetAccountsRequest) SetSort(v []string) {
	o.Sort = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetAccountsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetAccountsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetAccountsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetAccountsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetAccountsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetAccountsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GetAccountsRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GetAccountsRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GetAccountsRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetSearchType returns the SearchType field value if set, zero value otherwise.
func (o *GetAccountsRequest) GetSearchType() string {
	if o == nil || o.SearchType == nil {
		var ret string
		return ret
	}
	return *o.SearchType
}

// GetSearchTypeOk returns a tuple with the SearchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountsRequest) GetSearchTypeOk() (*string, bool) {
	if o == nil || o.SearchType == nil {
		return nil, false
	}
	return o.SearchType, true
}

// HasSearchType returns a boolean if a field has been set.
func (o *GetAccountsRequest) HasSearchType() bool {
	if o != nil && o.SearchType != nil {
		return true
	}

	return false
}

// SetSearchType gets a reference to the given string and assigns it to the SearchType field.
func (o *GetAccountsRequest) SetSearchType(v string) {
	o.SearchType = &v
}

func (o GetAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.SearchType != nil {
		toSerialize["searchType"] = o.SearchType
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccountsRequest struct {
	value *GetAccountsRequest
	isSet bool
}

func (v NullableGetAccountsRequest) Get() *GetAccountsRequest {
	return v.value
}

func (v *NullableGetAccountsRequest) Set(val *GetAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountsRequest(val *GetAccountsRequest) *NullableGetAccountsRequest {
	return &NullableGetAccountsRequest{value: val, isSet: true}
}

func (v NullableGetAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

