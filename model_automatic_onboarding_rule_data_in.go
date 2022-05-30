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

// AutomaticOnboardingRuleDataIn struct for AutomaticOnboardingRuleDataIn
type AutomaticOnboardingRuleDataIn struct {
	// The ID of the platform that will be associated to the onboarded account.
	TargetPlatformId string `json:"TargetPlatformId"`
	// The name of the Safe where the onboarded account will be stored.
	TargetSafeName string `json:"TargetSafeName"`
	IsAdminIDFilter *bool `json:"IsAdminIDFilter,omitempty"`
	MachineTypeFilter *int32 `json:"MachineTypeFilter,omitempty"`
	// The System Type by which to filter.
	SystemTypeFilter int32 `json:"SystemTypeFilter"`
	// The name of the user by which to filter.
	UserNameFilter *string `json:"UserNameFilter,omitempty"`
	RuleName *string `json:"RuleName,omitempty"`
	RuleDescription *string `json:"RuleDescription,omitempty"`
	UserNameMethod *int32 `json:"UserNameMethod,omitempty"`
	// The Machine Type by which to filter.
	AddressFilter *string `json:"AddressFilter,omitempty"`
	AddressMethod *int32 `json:"AddressMethod,omitempty"`
	AccountCategoryFilter *int32 `json:"AccountCategoryFilter,omitempty"`
}

// NewAutomaticOnboardingRuleDataIn instantiates a new AutomaticOnboardingRuleDataIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticOnboardingRuleDataIn(targetPlatformId string, targetSafeName string, systemTypeFilter int32) *AutomaticOnboardingRuleDataIn {
	this := AutomaticOnboardingRuleDataIn{}
	this.TargetPlatformId = targetPlatformId
	this.TargetSafeName = targetSafeName
	this.SystemTypeFilter = systemTypeFilter
	return &this
}

// NewAutomaticOnboardingRuleDataInWithDefaults instantiates a new AutomaticOnboardingRuleDataIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticOnboardingRuleDataInWithDefaults() *AutomaticOnboardingRuleDataIn {
	this := AutomaticOnboardingRuleDataIn{}
	return &this
}

// GetTargetPlatformId returns the TargetPlatformId field value
func (o *AutomaticOnboardingRuleDataIn) GetTargetPlatformId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPlatformId
}

// GetTargetPlatformIdOk returns a tuple with the TargetPlatformId field value
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetTargetPlatformIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPlatformId, true
}

// SetTargetPlatformId sets field value
func (o *AutomaticOnboardingRuleDataIn) SetTargetPlatformId(v string) {
	o.TargetPlatformId = v
}

// GetTargetSafeName returns the TargetSafeName field value
func (o *AutomaticOnboardingRuleDataIn) GetTargetSafeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetSafeName
}

// GetTargetSafeNameOk returns a tuple with the TargetSafeName field value
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetTargetSafeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetSafeName, true
}

// SetTargetSafeName sets field value
func (o *AutomaticOnboardingRuleDataIn) SetTargetSafeName(v string) {
	o.TargetSafeName = v
}

// GetIsAdminIDFilter returns the IsAdminIDFilter field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetIsAdminIDFilter() bool {
	if o == nil || o.IsAdminIDFilter == nil {
		var ret bool
		return ret
	}
	return *o.IsAdminIDFilter
}

// GetIsAdminIDFilterOk returns a tuple with the IsAdminIDFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetIsAdminIDFilterOk() (*bool, bool) {
	if o == nil || o.IsAdminIDFilter == nil {
		return nil, false
	}
	return o.IsAdminIDFilter, true
}

// HasIsAdminIDFilter returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasIsAdminIDFilter() bool {
	if o != nil && o.IsAdminIDFilter != nil {
		return true
	}

	return false
}

// SetIsAdminIDFilter gets a reference to the given bool and assigns it to the IsAdminIDFilter field.
func (o *AutomaticOnboardingRuleDataIn) SetIsAdminIDFilter(v bool) {
	o.IsAdminIDFilter = &v
}

// GetMachineTypeFilter returns the MachineTypeFilter field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetMachineTypeFilter() int32 {
	if o == nil || o.MachineTypeFilter == nil {
		var ret int32
		return ret
	}
	return *o.MachineTypeFilter
}

// GetMachineTypeFilterOk returns a tuple with the MachineTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetMachineTypeFilterOk() (*int32, bool) {
	if o == nil || o.MachineTypeFilter == nil {
		return nil, false
	}
	return o.MachineTypeFilter, true
}

// HasMachineTypeFilter returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasMachineTypeFilter() bool {
	if o != nil && o.MachineTypeFilter != nil {
		return true
	}

	return false
}

// SetMachineTypeFilter gets a reference to the given int32 and assigns it to the MachineTypeFilter field.
func (o *AutomaticOnboardingRuleDataIn) SetMachineTypeFilter(v int32) {
	o.MachineTypeFilter = &v
}

// GetSystemTypeFilter returns the SystemTypeFilter field value
func (o *AutomaticOnboardingRuleDataIn) GetSystemTypeFilter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SystemTypeFilter
}

// GetSystemTypeFilterOk returns a tuple with the SystemTypeFilter field value
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetSystemTypeFilterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemTypeFilter, true
}

// SetSystemTypeFilter sets field value
func (o *AutomaticOnboardingRuleDataIn) SetSystemTypeFilter(v int32) {
	o.SystemTypeFilter = v
}

// GetUserNameFilter returns the UserNameFilter field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetUserNameFilter() string {
	if o == nil || o.UserNameFilter == nil {
		var ret string
		return ret
	}
	return *o.UserNameFilter
}

// GetUserNameFilterOk returns a tuple with the UserNameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetUserNameFilterOk() (*string, bool) {
	if o == nil || o.UserNameFilter == nil {
		return nil, false
	}
	return o.UserNameFilter, true
}

// HasUserNameFilter returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasUserNameFilter() bool {
	if o != nil && o.UserNameFilter != nil {
		return true
	}

	return false
}

// SetUserNameFilter gets a reference to the given string and assigns it to the UserNameFilter field.
func (o *AutomaticOnboardingRuleDataIn) SetUserNameFilter(v string) {
	o.UserNameFilter = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetRuleName() string {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetRuleNameOk() (*string, bool) {
	if o == nil || o.RuleName == nil {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasRuleName() bool {
	if o != nil && o.RuleName != nil {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *AutomaticOnboardingRuleDataIn) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleDescription returns the RuleDescription field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetRuleDescription() string {
	if o == nil || o.RuleDescription == nil {
		var ret string
		return ret
	}
	return *o.RuleDescription
}

// GetRuleDescriptionOk returns a tuple with the RuleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetRuleDescriptionOk() (*string, bool) {
	if o == nil || o.RuleDescription == nil {
		return nil, false
	}
	return o.RuleDescription, true
}

// HasRuleDescription returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasRuleDescription() bool {
	if o != nil && o.RuleDescription != nil {
		return true
	}

	return false
}

// SetRuleDescription gets a reference to the given string and assigns it to the RuleDescription field.
func (o *AutomaticOnboardingRuleDataIn) SetRuleDescription(v string) {
	o.RuleDescription = &v
}

// GetUserNameMethod returns the UserNameMethod field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetUserNameMethod() int32 {
	if o == nil || o.UserNameMethod == nil {
		var ret int32
		return ret
	}
	return *o.UserNameMethod
}

// GetUserNameMethodOk returns a tuple with the UserNameMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetUserNameMethodOk() (*int32, bool) {
	if o == nil || o.UserNameMethod == nil {
		return nil, false
	}
	return o.UserNameMethod, true
}

// HasUserNameMethod returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasUserNameMethod() bool {
	if o != nil && o.UserNameMethod != nil {
		return true
	}

	return false
}

// SetUserNameMethod gets a reference to the given int32 and assigns it to the UserNameMethod field.
func (o *AutomaticOnboardingRuleDataIn) SetUserNameMethod(v int32) {
	o.UserNameMethod = &v
}

// GetAddressFilter returns the AddressFilter field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetAddressFilter() string {
	if o == nil || o.AddressFilter == nil {
		var ret string
		return ret
	}
	return *o.AddressFilter
}

// GetAddressFilterOk returns a tuple with the AddressFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetAddressFilterOk() (*string, bool) {
	if o == nil || o.AddressFilter == nil {
		return nil, false
	}
	return o.AddressFilter, true
}

// HasAddressFilter returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasAddressFilter() bool {
	if o != nil && o.AddressFilter != nil {
		return true
	}

	return false
}

// SetAddressFilter gets a reference to the given string and assigns it to the AddressFilter field.
func (o *AutomaticOnboardingRuleDataIn) SetAddressFilter(v string) {
	o.AddressFilter = &v
}

// GetAddressMethod returns the AddressMethod field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetAddressMethod() int32 {
	if o == nil || o.AddressMethod == nil {
		var ret int32
		return ret
	}
	return *o.AddressMethod
}

// GetAddressMethodOk returns a tuple with the AddressMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetAddressMethodOk() (*int32, bool) {
	if o == nil || o.AddressMethod == nil {
		return nil, false
	}
	return o.AddressMethod, true
}

// HasAddressMethod returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasAddressMethod() bool {
	if o != nil && o.AddressMethod != nil {
		return true
	}

	return false
}

// SetAddressMethod gets a reference to the given int32 and assigns it to the AddressMethod field.
func (o *AutomaticOnboardingRuleDataIn) SetAddressMethod(v int32) {
	o.AddressMethod = &v
}

// GetAccountCategoryFilter returns the AccountCategoryFilter field value if set, zero value otherwise.
func (o *AutomaticOnboardingRuleDataIn) GetAccountCategoryFilter() int32 {
	if o == nil || o.AccountCategoryFilter == nil {
		var ret int32
		return ret
	}
	return *o.AccountCategoryFilter
}

// GetAccountCategoryFilterOk returns a tuple with the AccountCategoryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticOnboardingRuleDataIn) GetAccountCategoryFilterOk() (*int32, bool) {
	if o == nil || o.AccountCategoryFilter == nil {
		return nil, false
	}
	return o.AccountCategoryFilter, true
}

// HasAccountCategoryFilter returns a boolean if a field has been set.
func (o *AutomaticOnboardingRuleDataIn) HasAccountCategoryFilter() bool {
	if o != nil && o.AccountCategoryFilter != nil {
		return true
	}

	return false
}

// SetAccountCategoryFilter gets a reference to the given int32 and assigns it to the AccountCategoryFilter field.
func (o *AutomaticOnboardingRuleDataIn) SetAccountCategoryFilter(v int32) {
	o.AccountCategoryFilter = &v
}

func (o AutomaticOnboardingRuleDataIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["TargetPlatformId"] = o.TargetPlatformId
	}
	if true {
		toSerialize["TargetSafeName"] = o.TargetSafeName
	}
	if o.IsAdminIDFilter != nil {
		toSerialize["IsAdminIDFilter"] = o.IsAdminIDFilter
	}
	if o.MachineTypeFilter != nil {
		toSerialize["MachineTypeFilter"] = o.MachineTypeFilter
	}
	if true {
		toSerialize["SystemTypeFilter"] = o.SystemTypeFilter
	}
	if o.UserNameFilter != nil {
		toSerialize["UserNameFilter"] = o.UserNameFilter
	}
	if o.RuleName != nil {
		toSerialize["RuleName"] = o.RuleName
	}
	if o.RuleDescription != nil {
		toSerialize["RuleDescription"] = o.RuleDescription
	}
	if o.UserNameMethod != nil {
		toSerialize["UserNameMethod"] = o.UserNameMethod
	}
	if o.AddressFilter != nil {
		toSerialize["AddressFilter"] = o.AddressFilter
	}
	if o.AddressMethod != nil {
		toSerialize["AddressMethod"] = o.AddressMethod
	}
	if o.AccountCategoryFilter != nil {
		toSerialize["AccountCategoryFilter"] = o.AccountCategoryFilter
	}
	return json.Marshal(toSerialize)
}

type NullableAutomaticOnboardingRuleDataIn struct {
	value *AutomaticOnboardingRuleDataIn
	isSet bool
}

func (v NullableAutomaticOnboardingRuleDataIn) Get() *AutomaticOnboardingRuleDataIn {
	return v.value
}

func (v *NullableAutomaticOnboardingRuleDataIn) Set(val *AutomaticOnboardingRuleDataIn) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticOnboardingRuleDataIn) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticOnboardingRuleDataIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticOnboardingRuleDataIn(val *AutomaticOnboardingRuleDataIn) *NullableAutomaticOnboardingRuleDataIn {
	return &NullableAutomaticOnboardingRuleDataIn{value: val, isSet: true}
}

func (v NullableAutomaticOnboardingRuleDataIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticOnboardingRuleDataIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


