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

// TargetPlatform struct for TargetPlatform
type TargetPlatform struct {
	// Indicates whether a platform is active or inactive. An inactive platform cannot be assigned to accounts but will continue to manage accounts that were already assigned to it
	Active *bool `json:"Active,omitempty"`
	// The type of system to target belong to
	SystemType *string `json:"SystemType,omitempty"`
	// Regular expression of safes in which accounts from this platform can be managed
	AllowedSafes *string `json:"AllowedSafes,omitempty"`
	PrivilegedAccessWorkflows *PrivilegedAccessWorkflowsPolicy `json:"PrivilegedAccessWorkflows,omitempty"`
	CredentialsManagementPolicy *CredentialsManagementPolicy `json:"CredentialsManagementPolicy,omitempty"`
	PrivilegedSessionManagement *PrivilegedSessionManagementBase `json:"PrivilegedSessionManagement,omitempty"`
	// Unique ID of the platform in the system
	ID *int64 `json:"ID,omitempty"`
	// Unique textual representation of the platform in the system
	PlatformID *string `json:"PlatformID,omitempty"`
	// Display name of the platform
	Name *string `json:"Name,omitempty"`
}

// NewTargetPlatform instantiates a new TargetPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetPlatform() *TargetPlatform {
	this := TargetPlatform{}
	return &this
}

// NewTargetPlatformWithDefaults instantiates a new TargetPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetPlatformWithDefaults() *TargetPlatform {
	this := TargetPlatform{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TargetPlatform) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TargetPlatform) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TargetPlatform) SetActive(v bool) {
	o.Active = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *TargetPlatform) GetSystemType() string {
	if o == nil || o.SystemType == nil {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetSystemTypeOk() (*string, bool) {
	if o == nil || o.SystemType == nil {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *TargetPlatform) HasSystemType() bool {
	if o != nil && o.SystemType != nil {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *TargetPlatform) SetSystemType(v string) {
	o.SystemType = &v
}

// GetAllowedSafes returns the AllowedSafes field value if set, zero value otherwise.
func (o *TargetPlatform) GetAllowedSafes() string {
	if o == nil || o.AllowedSafes == nil {
		var ret string
		return ret
	}
	return *o.AllowedSafes
}

// GetAllowedSafesOk returns a tuple with the AllowedSafes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetAllowedSafesOk() (*string, bool) {
	if o == nil || o.AllowedSafes == nil {
		return nil, false
	}
	return o.AllowedSafes, true
}

// HasAllowedSafes returns a boolean if a field has been set.
func (o *TargetPlatform) HasAllowedSafes() bool {
	if o != nil && o.AllowedSafes != nil {
		return true
	}

	return false
}

// SetAllowedSafes gets a reference to the given string and assigns it to the AllowedSafes field.
func (o *TargetPlatform) SetAllowedSafes(v string) {
	o.AllowedSafes = &v
}

// GetPrivilegedAccessWorkflows returns the PrivilegedAccessWorkflows field value if set, zero value otherwise.
func (o *TargetPlatform) GetPrivilegedAccessWorkflows() PrivilegedAccessWorkflowsPolicy {
	if o == nil || o.PrivilegedAccessWorkflows == nil {
		var ret PrivilegedAccessWorkflowsPolicy
		return ret
	}
	return *o.PrivilegedAccessWorkflows
}

// GetPrivilegedAccessWorkflowsOk returns a tuple with the PrivilegedAccessWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetPrivilegedAccessWorkflowsOk() (*PrivilegedAccessWorkflowsPolicy, bool) {
	if o == nil || o.PrivilegedAccessWorkflows == nil {
		return nil, false
	}
	return o.PrivilegedAccessWorkflows, true
}

// HasPrivilegedAccessWorkflows returns a boolean if a field has been set.
func (o *TargetPlatform) HasPrivilegedAccessWorkflows() bool {
	if o != nil && o.PrivilegedAccessWorkflows != nil {
		return true
	}

	return false
}

// SetPrivilegedAccessWorkflows gets a reference to the given PrivilegedAccessWorkflowsPolicy and assigns it to the PrivilegedAccessWorkflows field.
func (o *TargetPlatform) SetPrivilegedAccessWorkflows(v PrivilegedAccessWorkflowsPolicy) {
	o.PrivilegedAccessWorkflows = &v
}

// GetCredentialsManagementPolicy returns the CredentialsManagementPolicy field value if set, zero value otherwise.
func (o *TargetPlatform) GetCredentialsManagementPolicy() CredentialsManagementPolicy {
	if o == nil || o.CredentialsManagementPolicy == nil {
		var ret CredentialsManagementPolicy
		return ret
	}
	return *o.CredentialsManagementPolicy
}

// GetCredentialsManagementPolicyOk returns a tuple with the CredentialsManagementPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetCredentialsManagementPolicyOk() (*CredentialsManagementPolicy, bool) {
	if o == nil || o.CredentialsManagementPolicy == nil {
		return nil, false
	}
	return o.CredentialsManagementPolicy, true
}

// HasCredentialsManagementPolicy returns a boolean if a field has been set.
func (o *TargetPlatform) HasCredentialsManagementPolicy() bool {
	if o != nil && o.CredentialsManagementPolicy != nil {
		return true
	}

	return false
}

// SetCredentialsManagementPolicy gets a reference to the given CredentialsManagementPolicy and assigns it to the CredentialsManagementPolicy field.
func (o *TargetPlatform) SetCredentialsManagementPolicy(v CredentialsManagementPolicy) {
	o.CredentialsManagementPolicy = &v
}

// GetPrivilegedSessionManagement returns the PrivilegedSessionManagement field value if set, zero value otherwise.
func (o *TargetPlatform) GetPrivilegedSessionManagement() PrivilegedSessionManagementBase {
	if o == nil || o.PrivilegedSessionManagement == nil {
		var ret PrivilegedSessionManagementBase
		return ret
	}
	return *o.PrivilegedSessionManagement
}

// GetPrivilegedSessionManagementOk returns a tuple with the PrivilegedSessionManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetPrivilegedSessionManagementOk() (*PrivilegedSessionManagementBase, bool) {
	if o == nil || o.PrivilegedSessionManagement == nil {
		return nil, false
	}
	return o.PrivilegedSessionManagement, true
}

// HasPrivilegedSessionManagement returns a boolean if a field has been set.
func (o *TargetPlatform) HasPrivilegedSessionManagement() bool {
	if o != nil && o.PrivilegedSessionManagement != nil {
		return true
	}

	return false
}

// SetPrivilegedSessionManagement gets a reference to the given PrivilegedSessionManagementBase and assigns it to the PrivilegedSessionManagement field.
func (o *TargetPlatform) SetPrivilegedSessionManagement(v PrivilegedSessionManagementBase) {
	o.PrivilegedSessionManagement = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TargetPlatform) GetID() int64 {
	if o == nil || o.ID == nil {
		var ret int64
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetIDOk() (*int64, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TargetPlatform) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given int64 and assigns it to the ID field.
func (o *TargetPlatform) SetID(v int64) {
	o.ID = &v
}

// GetPlatformID returns the PlatformID field value if set, zero value otherwise.
func (o *TargetPlatform) GetPlatformID() string {
	if o == nil || o.PlatformID == nil {
		var ret string
		return ret
	}
	return *o.PlatformID
}

// GetPlatformIDOk returns a tuple with the PlatformID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetPlatformIDOk() (*string, bool) {
	if o == nil || o.PlatformID == nil {
		return nil, false
	}
	return o.PlatformID, true
}

// HasPlatformID returns a boolean if a field has been set.
func (o *TargetPlatform) HasPlatformID() bool {
	if o != nil && o.PlatformID != nil {
		return true
	}

	return false
}

// SetPlatformID gets a reference to the given string and assigns it to the PlatformID field.
func (o *TargetPlatform) SetPlatformID(v string) {
	o.PlatformID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TargetPlatform) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPlatform) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TargetPlatform) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TargetPlatform) SetName(v string) {
	o.Name = &v
}

func (o TargetPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["Active"] = o.Active
	}
	if o.SystemType != nil {
		toSerialize["SystemType"] = o.SystemType
	}
	if o.AllowedSafes != nil {
		toSerialize["AllowedSafes"] = o.AllowedSafes
	}
	if o.PrivilegedAccessWorkflows != nil {
		toSerialize["PrivilegedAccessWorkflows"] = o.PrivilegedAccessWorkflows
	}
	if o.CredentialsManagementPolicy != nil {
		toSerialize["CredentialsManagementPolicy"] = o.CredentialsManagementPolicy
	}
	if o.PrivilegedSessionManagement != nil {
		toSerialize["PrivilegedSessionManagement"] = o.PrivilegedSessionManagement
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

type NullableTargetPlatform struct {
	value *TargetPlatform
	isSet bool
}

func (v NullableTargetPlatform) Get() *TargetPlatform {
	return v.value
}

func (v *NullableTargetPlatform) Set(val *TargetPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetPlatform(val *TargetPlatform) *NullableTargetPlatform {
	return &NullableTargetPlatform{value: val, isSet: true}
}

func (v NullableTargetPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

