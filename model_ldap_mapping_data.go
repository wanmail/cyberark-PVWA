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

// LDAPMappingData struct for LDAPMappingData
type LDAPMappingData struct {
	// The LDAP branch used for the directory queries.
	LDAPBranch string `json:"LDAPBranch"`
	// A list of Vault groups that the user will be added to once mapped.
	VaultGroups []string `json:"VaultGroups,omitempty"`
	// The security attributes and authorizations that are applied when LDAP user accounts in the current mapping are created in the Vault.  Here is the list with the possible authorizations:  AddSafes,   AuditUsers,   AddUpdateUsers,   ResetUsersPasswords,   ActivateUsers,  AddNetworkAreas,  ManageServerFileCategories,  BackupAllSafes,  RestoreAllSafes.
	MappingAuthorizations []string `json:"MappingAuthorizations,omitempty"`
	// The specific Vault location path that users in the current mapping are added under.  This value cannot be updated.
	Location *string `json:"Location,omitempty"`
	// Read-only - The authentication method that the users belong to this map will use to log onto the Vault.
	AuthenticationMethod []string `json:"AuthenticationMethod,omitempty"`
	// Read-only - The interfaces that users in the current mapping can use to access the Vault.  This will only be available to users with Add/Update users permissions.
	UserType *string `json:"UserType,omitempty"`
	// Read-only - Whether users in the current mapping are temporarily inaccessible.
	DisableUser *bool `json:"DisableUser,omitempty"`
	// The number of days that activity records are stored for users in the current mapping before they can be deleted.
	UserActivityLogPeriod *int32 `json:"UserActivityLogPeriod,omitempty"`
	// Read-only - The date after which user accounts in the current mapping are no longer accessible, in Unix time (\"0\" if never).
	UserExpiration *int64 `json:"UserExpiration,omitempty"`
	// Read-only - Whether users in the current mapping can log on to the Vault only starting from specific hours.
	LogonFromHour *int32 `json:"LogonFromHour,omitempty"`
	// Read-only - Whether users in the current mapping can log on to the Vault only up to specific hours.
	LogonToHour *int32 `json:"LogonToHour,omitempty"`
	// Read-only - Unique ID of the directory mapping
	MappingID *int64 `json:"MappingID,omitempty"`
	// The order in which the maps are matched with users and groups   from the External Directory when they are created in the Vault.  This field is read only.
	DirectoryMappingOrder *int32 `json:"DirectoryMappingOrder,omitempty"`
	// The name of the role.  For example: Vault Admins, Safe Managers.
	MappingName string `json:"MappingName"`
	// LDAP users that the filter applies to them are assigned to relevant roles in the system.
	LDAPQuery *string `json:"LDAPQuery,omitempty"`
	// Users who belong to these LDAP groups are assigned to relevant roles in the system.  Note: This field is required if the VaultGroups field was sent.
	DomainGroups []string `json:"DomainGroups,omitempty"`
}

// NewLDAPMappingData instantiates a new LDAPMappingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPMappingData(lDAPBranch string, mappingName string) *LDAPMappingData {
	this := LDAPMappingData{}
	this.LDAPBranch = lDAPBranch
	this.MappingName = mappingName
	return &this
}

// NewLDAPMappingDataWithDefaults instantiates a new LDAPMappingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPMappingDataWithDefaults() *LDAPMappingData {
	this := LDAPMappingData{}
	return &this
}

// GetLDAPBranch returns the LDAPBranch field value
func (o *LDAPMappingData) GetLDAPBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LDAPBranch
}

// GetLDAPBranchOk returns a tuple with the LDAPBranch field value
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetLDAPBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LDAPBranch, true
}

// SetLDAPBranch sets field value
func (o *LDAPMappingData) SetLDAPBranch(v string) {
	o.LDAPBranch = v
}

// GetVaultGroups returns the VaultGroups field value if set, zero value otherwise.
func (o *LDAPMappingData) GetVaultGroups() []string {
	if o == nil || o.VaultGroups == nil {
		var ret []string
		return ret
	}
	return o.VaultGroups
}

// GetVaultGroupsOk returns a tuple with the VaultGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetVaultGroupsOk() ([]string, bool) {
	if o == nil || o.VaultGroups == nil {
		return nil, false
	}
	return o.VaultGroups, true
}

// HasVaultGroups returns a boolean if a field has been set.
func (o *LDAPMappingData) HasVaultGroups() bool {
	if o != nil && o.VaultGroups != nil {
		return true
	}

	return false
}

// SetVaultGroups gets a reference to the given []string and assigns it to the VaultGroups field.
func (o *LDAPMappingData) SetVaultGroups(v []string) {
	o.VaultGroups = v
}

// GetMappingAuthorizations returns the MappingAuthorizations field value if set, zero value otherwise.
func (o *LDAPMappingData) GetMappingAuthorizations() []string {
	if o == nil || o.MappingAuthorizations == nil {
		var ret []string
		return ret
	}
	return o.MappingAuthorizations
}

// GetMappingAuthorizationsOk returns a tuple with the MappingAuthorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetMappingAuthorizationsOk() ([]string, bool) {
	if o == nil || o.MappingAuthorizations == nil {
		return nil, false
	}
	return o.MappingAuthorizations, true
}

// HasMappingAuthorizations returns a boolean if a field has been set.
func (o *LDAPMappingData) HasMappingAuthorizations() bool {
	if o != nil && o.MappingAuthorizations != nil {
		return true
	}

	return false
}

// SetMappingAuthorizations gets a reference to the given []string and assigns it to the MappingAuthorizations field.
func (o *LDAPMappingData) SetMappingAuthorizations(v []string) {
	o.MappingAuthorizations = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LDAPMappingData) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LDAPMappingData) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *LDAPMappingData) SetLocation(v string) {
	o.Location = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *LDAPMappingData) GetAuthenticationMethod() []string {
	if o == nil || o.AuthenticationMethod == nil {
		var ret []string
		return ret
	}
	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetAuthenticationMethodOk() ([]string, bool) {
	if o == nil || o.AuthenticationMethod == nil {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *LDAPMappingData) HasAuthenticationMethod() bool {
	if o != nil && o.AuthenticationMethod != nil {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given []string and assigns it to the AuthenticationMethod field.
func (o *LDAPMappingData) SetAuthenticationMethod(v []string) {
	o.AuthenticationMethod = v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *LDAPMappingData) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *LDAPMappingData) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *LDAPMappingData) SetUserType(v string) {
	o.UserType = &v
}

// GetDisableUser returns the DisableUser field value if set, zero value otherwise.
func (o *LDAPMappingData) GetDisableUser() bool {
	if o == nil || o.DisableUser == nil {
		var ret bool
		return ret
	}
	return *o.DisableUser
}

// GetDisableUserOk returns a tuple with the DisableUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetDisableUserOk() (*bool, bool) {
	if o == nil || o.DisableUser == nil {
		return nil, false
	}
	return o.DisableUser, true
}

// HasDisableUser returns a boolean if a field has been set.
func (o *LDAPMappingData) HasDisableUser() bool {
	if o != nil && o.DisableUser != nil {
		return true
	}

	return false
}

// SetDisableUser gets a reference to the given bool and assigns it to the DisableUser field.
func (o *LDAPMappingData) SetDisableUser(v bool) {
	o.DisableUser = &v
}

// GetUserActivityLogPeriod returns the UserActivityLogPeriod field value if set, zero value otherwise.
func (o *LDAPMappingData) GetUserActivityLogPeriod() int32 {
	if o == nil || o.UserActivityLogPeriod == nil {
		var ret int32
		return ret
	}
	return *o.UserActivityLogPeriod
}

// GetUserActivityLogPeriodOk returns a tuple with the UserActivityLogPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetUserActivityLogPeriodOk() (*int32, bool) {
	if o == nil || o.UserActivityLogPeriod == nil {
		return nil, false
	}
	return o.UserActivityLogPeriod, true
}

// HasUserActivityLogPeriod returns a boolean if a field has been set.
func (o *LDAPMappingData) HasUserActivityLogPeriod() bool {
	if o != nil && o.UserActivityLogPeriod != nil {
		return true
	}

	return false
}

// SetUserActivityLogPeriod gets a reference to the given int32 and assigns it to the UserActivityLogPeriod field.
func (o *LDAPMappingData) SetUserActivityLogPeriod(v int32) {
	o.UserActivityLogPeriod = &v
}

// GetUserExpiration returns the UserExpiration field value if set, zero value otherwise.
func (o *LDAPMappingData) GetUserExpiration() int64 {
	if o == nil || o.UserExpiration == nil {
		var ret int64
		return ret
	}
	return *o.UserExpiration
}

// GetUserExpirationOk returns a tuple with the UserExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetUserExpirationOk() (*int64, bool) {
	if o == nil || o.UserExpiration == nil {
		return nil, false
	}
	return o.UserExpiration, true
}

// HasUserExpiration returns a boolean if a field has been set.
func (o *LDAPMappingData) HasUserExpiration() bool {
	if o != nil && o.UserExpiration != nil {
		return true
	}

	return false
}

// SetUserExpiration gets a reference to the given int64 and assigns it to the UserExpiration field.
func (o *LDAPMappingData) SetUserExpiration(v int64) {
	o.UserExpiration = &v
}

// GetLogonFromHour returns the LogonFromHour field value if set, zero value otherwise.
func (o *LDAPMappingData) GetLogonFromHour() int32 {
	if o == nil || o.LogonFromHour == nil {
		var ret int32
		return ret
	}
	return *o.LogonFromHour
}

// GetLogonFromHourOk returns a tuple with the LogonFromHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetLogonFromHourOk() (*int32, bool) {
	if o == nil || o.LogonFromHour == nil {
		return nil, false
	}
	return o.LogonFromHour, true
}

// HasLogonFromHour returns a boolean if a field has been set.
func (o *LDAPMappingData) HasLogonFromHour() bool {
	if o != nil && o.LogonFromHour != nil {
		return true
	}

	return false
}

// SetLogonFromHour gets a reference to the given int32 and assigns it to the LogonFromHour field.
func (o *LDAPMappingData) SetLogonFromHour(v int32) {
	o.LogonFromHour = &v
}

// GetLogonToHour returns the LogonToHour field value if set, zero value otherwise.
func (o *LDAPMappingData) GetLogonToHour() int32 {
	if o == nil || o.LogonToHour == nil {
		var ret int32
		return ret
	}
	return *o.LogonToHour
}

// GetLogonToHourOk returns a tuple with the LogonToHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetLogonToHourOk() (*int32, bool) {
	if o == nil || o.LogonToHour == nil {
		return nil, false
	}
	return o.LogonToHour, true
}

// HasLogonToHour returns a boolean if a field has been set.
func (o *LDAPMappingData) HasLogonToHour() bool {
	if o != nil && o.LogonToHour != nil {
		return true
	}

	return false
}

// SetLogonToHour gets a reference to the given int32 and assigns it to the LogonToHour field.
func (o *LDAPMappingData) SetLogonToHour(v int32) {
	o.LogonToHour = &v
}

// GetMappingID returns the MappingID field value if set, zero value otherwise.
func (o *LDAPMappingData) GetMappingID() int64 {
	if o == nil || o.MappingID == nil {
		var ret int64
		return ret
	}
	return *o.MappingID
}

// GetMappingIDOk returns a tuple with the MappingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetMappingIDOk() (*int64, bool) {
	if o == nil || o.MappingID == nil {
		return nil, false
	}
	return o.MappingID, true
}

// HasMappingID returns a boolean if a field has been set.
func (o *LDAPMappingData) HasMappingID() bool {
	if o != nil && o.MappingID != nil {
		return true
	}

	return false
}

// SetMappingID gets a reference to the given int64 and assigns it to the MappingID field.
func (o *LDAPMappingData) SetMappingID(v int64) {
	o.MappingID = &v
}

// GetDirectoryMappingOrder returns the DirectoryMappingOrder field value if set, zero value otherwise.
func (o *LDAPMappingData) GetDirectoryMappingOrder() int32 {
	if o == nil || o.DirectoryMappingOrder == nil {
		var ret int32
		return ret
	}
	return *o.DirectoryMappingOrder
}

// GetDirectoryMappingOrderOk returns a tuple with the DirectoryMappingOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetDirectoryMappingOrderOk() (*int32, bool) {
	if o == nil || o.DirectoryMappingOrder == nil {
		return nil, false
	}
	return o.DirectoryMappingOrder, true
}

// HasDirectoryMappingOrder returns a boolean if a field has been set.
func (o *LDAPMappingData) HasDirectoryMappingOrder() bool {
	if o != nil && o.DirectoryMappingOrder != nil {
		return true
	}

	return false
}

// SetDirectoryMappingOrder gets a reference to the given int32 and assigns it to the DirectoryMappingOrder field.
func (o *LDAPMappingData) SetDirectoryMappingOrder(v int32) {
	o.DirectoryMappingOrder = &v
}

// GetMappingName returns the MappingName field value
func (o *LDAPMappingData) GetMappingName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MappingName
}

// GetMappingNameOk returns a tuple with the MappingName field value
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetMappingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MappingName, true
}

// SetMappingName sets field value
func (o *LDAPMappingData) SetMappingName(v string) {
	o.MappingName = v
}

// GetLDAPQuery returns the LDAPQuery field value if set, zero value otherwise.
func (o *LDAPMappingData) GetLDAPQuery() string {
	if o == nil || o.LDAPQuery == nil {
		var ret string
		return ret
	}
	return *o.LDAPQuery
}

// GetLDAPQueryOk returns a tuple with the LDAPQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetLDAPQueryOk() (*string, bool) {
	if o == nil || o.LDAPQuery == nil {
		return nil, false
	}
	return o.LDAPQuery, true
}

// HasLDAPQuery returns a boolean if a field has been set.
func (o *LDAPMappingData) HasLDAPQuery() bool {
	if o != nil && o.LDAPQuery != nil {
		return true
	}

	return false
}

// SetLDAPQuery gets a reference to the given string and assigns it to the LDAPQuery field.
func (o *LDAPMappingData) SetLDAPQuery(v string) {
	o.LDAPQuery = &v
}

// GetDomainGroups returns the DomainGroups field value if set, zero value otherwise.
func (o *LDAPMappingData) GetDomainGroups() []string {
	if o == nil || o.DomainGroups == nil {
		var ret []string
		return ret
	}
	return o.DomainGroups
}

// GetDomainGroupsOk returns a tuple with the DomainGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPMappingData) GetDomainGroupsOk() ([]string, bool) {
	if o == nil || o.DomainGroups == nil {
		return nil, false
	}
	return o.DomainGroups, true
}

// HasDomainGroups returns a boolean if a field has been set.
func (o *LDAPMappingData) HasDomainGroups() bool {
	if o != nil && o.DomainGroups != nil {
		return true
	}

	return false
}

// SetDomainGroups gets a reference to the given []string and assigns it to the DomainGroups field.
func (o *LDAPMappingData) SetDomainGroups(v []string) {
	o.DomainGroups = v
}

func (o LDAPMappingData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LDAPBranch"] = o.LDAPBranch
	}
	if o.VaultGroups != nil {
		toSerialize["VaultGroups"] = o.VaultGroups
	}
	if o.MappingAuthorizations != nil {
		toSerialize["MappingAuthorizations"] = o.MappingAuthorizations
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.AuthenticationMethod != nil {
		toSerialize["AuthenticationMethod"] = o.AuthenticationMethod
	}
	if o.UserType != nil {
		toSerialize["UserType"] = o.UserType
	}
	if o.DisableUser != nil {
		toSerialize["DisableUser"] = o.DisableUser
	}
	if o.UserActivityLogPeriod != nil {
		toSerialize["UserActivityLogPeriod"] = o.UserActivityLogPeriod
	}
	if o.UserExpiration != nil {
		toSerialize["UserExpiration"] = o.UserExpiration
	}
	if o.LogonFromHour != nil {
		toSerialize["LogonFromHour"] = o.LogonFromHour
	}
	if o.LogonToHour != nil {
		toSerialize["LogonToHour"] = o.LogonToHour
	}
	if o.MappingID != nil {
		toSerialize["MappingID"] = o.MappingID
	}
	if o.DirectoryMappingOrder != nil {
		toSerialize["DirectoryMappingOrder"] = o.DirectoryMappingOrder
	}
	if true {
		toSerialize["MappingName"] = o.MappingName
	}
	if o.LDAPQuery != nil {
		toSerialize["LDAPQuery"] = o.LDAPQuery
	}
	if o.DomainGroups != nil {
		toSerialize["DomainGroups"] = o.DomainGroups
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPMappingData struct {
	value *LDAPMappingData
	isSet bool
}

func (v NullableLDAPMappingData) Get() *LDAPMappingData {
	return v.value
}

func (v *NullableLDAPMappingData) Set(val *LDAPMappingData) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPMappingData) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPMappingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPMappingData(val *LDAPMappingData) *NullableLDAPMappingData {
	return &NullableLDAPMappingData{value: val, isSet: true}
}

func (v NullableLDAPMappingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPMappingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


