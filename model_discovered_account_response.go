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

// DiscoveredAccountResponse struct for DiscoveredAccountResponse
type DiscoveredAccountResponse struct {
	// The ID of the on-boarded or pending account.
	Id *string `json:"id,omitempty"`
	// The status of the account. The possible values are:  addedAccount - the account was on-boarded to the Vault.  addedAsPending - the account was added to the Pending Accounts list.  updatedAccount - the account is already exists and new dependencies were added to it.  updatedPending - the pending account already exists and it was updated if needed.  alreadyExists - the account is already exists and no new dependencies were added.
	Status *string `json:"status,omitempty"`
	// The name of the account user.
	UserName *string `json:"userName,omitempty"`
	// The name or address of the machine where the account is located.
	Address *string `json:"address,omitempty"`
	// The name of the safe of the account.
	Safe *string `json:"safe,omitempty"`
	// The dependencies of the account.
	Dependencies []DiscoveredDependencyResponse `json:"dependencies,omitempty"`
}

// NewDiscoveredAccountResponse instantiates a new DiscoveredAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredAccountResponse() *DiscoveredAccountResponse {
	this := DiscoveredAccountResponse{}
	return &this
}

// NewDiscoveredAccountResponseWithDefaults instantiates a new DiscoveredAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredAccountResponseWithDefaults() *DiscoveredAccountResponse {
	this := DiscoveredAccountResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscoveredAccountResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAccountResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscoveredAccountResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiscoveredAccountResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DiscoveredAccountResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAccountResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DiscoveredAccountResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DiscoveredAccountResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DiscoveredAccountResponse) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAccountResponse) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DiscoveredAccountResponse) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DiscoveredAccountResponse) SetUserName(v string) {
	o.UserName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DiscoveredAccountResponse) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAccountResponse) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DiscoveredAccountResponse) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DiscoveredAccountResponse) SetAddress(v string) {
	o.Address = &v
}

// GetSafe returns the Safe field value if set, zero value otherwise.
func (o *DiscoveredAccountResponse) GetSafe() string {
	if o == nil || o.Safe == nil {
		var ret string
		return ret
	}
	return *o.Safe
}

// GetSafeOk returns a tuple with the Safe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAccountResponse) GetSafeOk() (*string, bool) {
	if o == nil || o.Safe == nil {
		return nil, false
	}
	return o.Safe, true
}

// HasSafe returns a boolean if a field has been set.
func (o *DiscoveredAccountResponse) HasSafe() bool {
	if o != nil && o.Safe != nil {
		return true
	}

	return false
}

// SetSafe gets a reference to the given string and assigns it to the Safe field.
func (o *DiscoveredAccountResponse) SetSafe(v string) {
	o.Safe = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *DiscoveredAccountResponse) GetDependencies() []DiscoveredDependencyResponse {
	if o == nil || o.Dependencies == nil {
		var ret []DiscoveredDependencyResponse
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAccountResponse) GetDependenciesOk() ([]DiscoveredDependencyResponse, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *DiscoveredAccountResponse) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []DiscoveredDependencyResponse and assigns it to the Dependencies field.
func (o *DiscoveredAccountResponse) SetDependencies(v []DiscoveredDependencyResponse) {
	o.Dependencies = v
}

func (o DiscoveredAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Safe != nil {
		toSerialize["safe"] = o.Safe
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	return json.Marshal(toSerialize)
}

type NullableDiscoveredAccountResponse struct {
	value *DiscoveredAccountResponse
	isSet bool
}

func (v NullableDiscoveredAccountResponse) Get() *DiscoveredAccountResponse {
	return v.value
}

func (v *NullableDiscoveredAccountResponse) Set(val *DiscoveredAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredAccountResponse(val *DiscoveredAccountResponse) *NullableDiscoveredAccountResponse {
	return &NullableDiscoveredAccountResponse{value: val, isSet: true}
}

func (v NullableDiscoveredAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


