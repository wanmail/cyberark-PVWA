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

// GetDiscoveredAccountsRequest struct for GetDiscoveredAccountsRequest
type GetDiscoveredAccountsRequest struct {
	// Searches according to REST standard (search=\"search word\").    Search is supported for userName and address.
	Search *string `json:"search,omitempty"`
	// The type of search to perform.     The keyword can either be contained within the account property values, or at the beginning of the value  specified in the search parameter.    When using a keyword at the beginning of a value, performance is enhanced.    Valid values: contains (default) or startswith
	SearchType *string `json:"searchType,omitempty"`
	// The offset of the first returned accounts into the list of results.
	Offset *int32 `json:"offset,omitempty"`
	// The maximum number of returned accounts.     If not specified, the server limits the results to 100.    The maximum number that can be specified is 1000.
	Limit *int32 `json:"limit,omitempty"`
	// Search for accounts filtered by:    • platformType - Get discovered accounts of a specific platform, using the platform type name.        Usage: filter=platformType eq Windows Server Local            Type: string            Valid values:                o Windows Server Local                    o Windows Desktop Local                    o Windows Domain                    o Unix                    o Unix SSH Key                    o AWS                    o AWS Access Keys                    o Azure Password Management        • privileged - Get only privileged or non-privileged discovered accounts.        Usage: filter=privileged eq true            Type: boolean            Valid values: true/false        • accountEnabled - Get only enabled or disabled discovered accounts.           Usage: filter=accountEnabled eq true            Type: boolean            Valid values: true/false                  Note: To use more than one filter, you can use the AND operator.         filter=platformType eq Windows Server Local AND privileged eq true
	Filter *string `json:"filter,omitempty"`
}

// NewGetDiscoveredAccountsRequest instantiates a new GetDiscoveredAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDiscoveredAccountsRequest() *GetDiscoveredAccountsRequest {
	this := GetDiscoveredAccountsRequest{}
	return &this
}

// NewGetDiscoveredAccountsRequestWithDefaults instantiates a new GetDiscoveredAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDiscoveredAccountsRequestWithDefaults() *GetDiscoveredAccountsRequest {
	this := GetDiscoveredAccountsRequest{}
	return &this
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsRequest) GetSearch() string {
	if o == nil || o.Search == nil {
		var ret string
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsRequest) GetSearchOk() (*string, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsRequest) HasSearch() bool {
	if o != nil && o.Search != nil {
		return true
	}

	return false
}

// SetSearch gets a reference to the given string and assigns it to the Search field.
func (o *GetDiscoveredAccountsRequest) SetSearch(v string) {
	o.Search = &v
}

// GetSearchType returns the SearchType field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsRequest) GetSearchType() string {
	if o == nil || o.SearchType == nil {
		var ret string
		return ret
	}
	return *o.SearchType
}

// GetSearchTypeOk returns a tuple with the SearchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsRequest) GetSearchTypeOk() (*string, bool) {
	if o == nil || o.SearchType == nil {
		return nil, false
	}
	return o.SearchType, true
}

// HasSearchType returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsRequest) HasSearchType() bool {
	if o != nil && o.SearchType != nil {
		return true
	}

	return false
}

// SetSearchType gets a reference to the given string and assigns it to the SearchType field.
func (o *GetDiscoveredAccountsRequest) SetSearchType(v string) {
	o.SearchType = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetDiscoveredAccountsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetDiscoveredAccountsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GetDiscoveredAccountsRequest) SetFilter(v string) {
	o.Filter = &v
}

func (o GetDiscoveredAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.SearchType != nil {
		toSerialize["searchType"] = o.SearchType
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
	return json.Marshal(toSerialize)
}

type NullableGetDiscoveredAccountsRequest struct {
	value *GetDiscoveredAccountsRequest
	isSet bool
}

func (v NullableGetDiscoveredAccountsRequest) Get() *GetDiscoveredAccountsRequest {
	return v.value
}

func (v *NullableGetDiscoveredAccountsRequest) Set(val *GetDiscoveredAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoveredAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoveredAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoveredAccountsRequest(val *GetDiscoveredAccountsRequest) *NullableGetDiscoveredAccountsRequest {
	return &NullableGetDiscoveredAccountsRequest{value: val, isSet: true}
}

func (v NullableGetDiscoveredAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoveredAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


