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

// AddSafeMemberResponse struct for AddSafeMemberResponse
type AddSafeMemberResponse struct {
	// The name of the Safe to be used when calling Safe APIs.
	SafeUrlId *string `json:"safeUrlId,omitempty"`
	// The unique name of the Safe.
	SafeName *string `json:"safeName,omitempty"`
	// The unique numerical ID of the Safe.
	SafeNumber *int64 `json:"safeNumber,omitempty"`
	// The Vault user ID, Domain user ID, or group ID of the Safe member.
	MemberId *int64 `json:"memberId,omitempty"`
	// The Vault user name, Domain user name or group name of the Safe member.
	MemberName *string `json:"memberName,omitempty"`
	// The member type.  Valid values: user\\group
	MemberType *string `json:"memberType,omitempty"`
	// The member's expiration date for this Safe.  For members that do not have an expiration date, this value will be null.
	MembershipExpirationDate *int64 `json:"membershipExpirationDate,omitempty"`
	// Whether or not the membership for the Safe is expired.For expired members, the value will be True.
	IsExpiredMembershipEnable *bool `json:"isExpiredMembershipEnable,omitempty"`
	// Whether the member is a predefined user or group of the Vault.
	IsPredefinedUser *bool `json:"isPredefinedUser,omitempty"`
	// Whether or not the current user can update the permissions of the member
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
	// <p>The permissions that the user or group has in this Safe.</p>  <p>    <B>useAccounts</B> - Use accounts but not view passwords.</p>  <p>    <B>retrieveAccounts</B> - Retrieve and view accounts in the Safe.</p>  <p>    <B>listAccounts</B> - View Accounts list.</p>  <p>    <B>addAccounts</B> - Add accounts in the Safe. Users who have this permission automatically have UpdateAccountProperties as well.</p>  <p>    <B>updateAccountContent</B> - Update existing account content.</p>  <p>    <B>updateAccountProperties</B> - Update existing account properties.</p>  <p>    <B>initiateCPMAccountManagementOperations</B> - Initiate password management operations through CPM, such as changing, verifying, and reconciling passwords. When this parameter is set to False, the SpecifyNextAccountContent parameter is also automatically set to False.</p>  <p>    <B>specifyNextAccountContent</B> - Specify the password that is used when the CPM changes the password value. This parameter can only be specified when the InitiateCPMAccountManagementOperations parameter is set to True. When InitiateCPMAccountManagementOperations is set to False this parameter is automatically set to False.</p>  <p>    <B>renameAccounts</B> - Rename existing accounts in the Safe.</p>  <p>    <B>deleteAccounts</B> - Delete existing passwords in the Safe.</p>  <p>    <B>unlockAccounts</B> - Unlock accounts that are locked by other users.</p>  <p>    <B>manageSafe</B> - Perform administrative tasksin the Safe, including: Update Safe properties, Recover the Safe, Delete the Safe.</p>  <p>    <B>manageSafeMembers</B> - Add and remove Safe members, and update their authorizations in the Safe.</p>  <p>    <B>backupSafe</B> - Create a backup of a Safe and its contents, and store in another location.</p>  <p>    <B>viewAuditLog</B> - View account and user activity in the Safe.</p>  <p>    <B>viewSafeMembers</B> - View Safe member's permissions.</p>  <p>    <B>requestsAuthorizationLevel1</B> - Request Authorization Level 1.</p>  <p>    <B>requestsAuthorizationLevel2</B> - Request Authorization Level 2.</p>  <p>    <B>accessWithoutConfirmation</B> - Access the Safe without confirmation from authorized users. This overrides the Safe properties that specify that Safe members require confirmation to access the Safe.</p>  <p>    <B>createFolders</B> - Create folders in the Safe.</p>  <p>    <B>deleteFolders</B> - Delete folders from the Safe.</p>  <p>    <B>moveAccountsAndFolders</B> - Move accounts and folders in the Safe to different folders and subfolders.</p>
	Permissions *map[string]bool `json:"permissions,omitempty"`
}

// NewAddSafeMemberResponse instantiates a new AddSafeMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSafeMemberResponse() *AddSafeMemberResponse {
	this := AddSafeMemberResponse{}
	return &this
}

// NewAddSafeMemberResponseWithDefaults instantiates a new AddSafeMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSafeMemberResponseWithDefaults() *AddSafeMemberResponse {
	this := AddSafeMemberResponse{}
	return &this
}

// GetSafeUrlId returns the SafeUrlId field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetSafeUrlId() string {
	if o == nil || o.SafeUrlId == nil {
		var ret string
		return ret
	}
	return *o.SafeUrlId
}

// GetSafeUrlIdOk returns a tuple with the SafeUrlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetSafeUrlIdOk() (*string, bool) {
	if o == nil || o.SafeUrlId == nil {
		return nil, false
	}
	return o.SafeUrlId, true
}

// HasSafeUrlId returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasSafeUrlId() bool {
	if o != nil && o.SafeUrlId != nil {
		return true
	}

	return false
}

// SetSafeUrlId gets a reference to the given string and assigns it to the SafeUrlId field.
func (o *AddSafeMemberResponse) SetSafeUrlId(v string) {
	o.SafeUrlId = &v
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetSafeName() string {
	if o == nil || o.SafeName == nil {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetSafeNameOk() (*string, bool) {
	if o == nil || o.SafeName == nil {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasSafeName() bool {
	if o != nil && o.SafeName != nil {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *AddSafeMemberResponse) SetSafeName(v string) {
	o.SafeName = &v
}

// GetSafeNumber returns the SafeNumber field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetSafeNumber() int64 {
	if o == nil || o.SafeNumber == nil {
		var ret int64
		return ret
	}
	return *o.SafeNumber
}

// GetSafeNumberOk returns a tuple with the SafeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetSafeNumberOk() (*int64, bool) {
	if o == nil || o.SafeNumber == nil {
		return nil, false
	}
	return o.SafeNumber, true
}

// HasSafeNumber returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasSafeNumber() bool {
	if o != nil && o.SafeNumber != nil {
		return true
	}

	return false
}

// SetSafeNumber gets a reference to the given int64 and assigns it to the SafeNumber field.
func (o *AddSafeMemberResponse) SetSafeNumber(v int64) {
	o.SafeNumber = &v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetMemberId() int64 {
	if o == nil || o.MemberId == nil {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetMemberIdOk() (*int64, bool) {
	if o == nil || o.MemberId == nil {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasMemberId() bool {
	if o != nil && o.MemberId != nil {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *AddSafeMemberResponse) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetMemberName returns the MemberName field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetMemberName() string {
	if o == nil || o.MemberName == nil {
		var ret string
		return ret
	}
	return *o.MemberName
}

// GetMemberNameOk returns a tuple with the MemberName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetMemberNameOk() (*string, bool) {
	if o == nil || o.MemberName == nil {
		return nil, false
	}
	return o.MemberName, true
}

// HasMemberName returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasMemberName() bool {
	if o != nil && o.MemberName != nil {
		return true
	}

	return false
}

// SetMemberName gets a reference to the given string and assigns it to the MemberName field.
func (o *AddSafeMemberResponse) SetMemberName(v string) {
	o.MemberName = &v
}

// GetMemberType returns the MemberType field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetMemberType() string {
	if o == nil || o.MemberType == nil {
		var ret string
		return ret
	}
	return *o.MemberType
}

// GetMemberTypeOk returns a tuple with the MemberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetMemberTypeOk() (*string, bool) {
	if o == nil || o.MemberType == nil {
		return nil, false
	}
	return o.MemberType, true
}

// HasMemberType returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasMemberType() bool {
	if o != nil && o.MemberType != nil {
		return true
	}

	return false
}

// SetMemberType gets a reference to the given string and assigns it to the MemberType field.
func (o *AddSafeMemberResponse) SetMemberType(v string) {
	o.MemberType = &v
}

// GetMembershipExpirationDate returns the MembershipExpirationDate field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetMembershipExpirationDate() int64 {
	if o == nil || o.MembershipExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.MembershipExpirationDate
}

// GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetMembershipExpirationDateOk() (*int64, bool) {
	if o == nil || o.MembershipExpirationDate == nil {
		return nil, false
	}
	return o.MembershipExpirationDate, true
}

// HasMembershipExpirationDate returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasMembershipExpirationDate() bool {
	if o != nil && o.MembershipExpirationDate != nil {
		return true
	}

	return false
}

// SetMembershipExpirationDate gets a reference to the given int64 and assigns it to the MembershipExpirationDate field.
func (o *AddSafeMemberResponse) SetMembershipExpirationDate(v int64) {
	o.MembershipExpirationDate = &v
}

// GetIsExpiredMembershipEnable returns the IsExpiredMembershipEnable field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetIsExpiredMembershipEnable() bool {
	if o == nil || o.IsExpiredMembershipEnable == nil {
		var ret bool
		return ret
	}
	return *o.IsExpiredMembershipEnable
}

// GetIsExpiredMembershipEnableOk returns a tuple with the IsExpiredMembershipEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetIsExpiredMembershipEnableOk() (*bool, bool) {
	if o == nil || o.IsExpiredMembershipEnable == nil {
		return nil, false
	}
	return o.IsExpiredMembershipEnable, true
}

// HasIsExpiredMembershipEnable returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasIsExpiredMembershipEnable() bool {
	if o != nil && o.IsExpiredMembershipEnable != nil {
		return true
	}

	return false
}

// SetIsExpiredMembershipEnable gets a reference to the given bool and assigns it to the IsExpiredMembershipEnable field.
func (o *AddSafeMemberResponse) SetIsExpiredMembershipEnable(v bool) {
	o.IsExpiredMembershipEnable = &v
}

// GetIsPredefinedUser returns the IsPredefinedUser field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetIsPredefinedUser() bool {
	if o == nil || o.IsPredefinedUser == nil {
		var ret bool
		return ret
	}
	return *o.IsPredefinedUser
}

// GetIsPredefinedUserOk returns a tuple with the IsPredefinedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetIsPredefinedUserOk() (*bool, bool) {
	if o == nil || o.IsPredefinedUser == nil {
		return nil, false
	}
	return o.IsPredefinedUser, true
}

// HasIsPredefinedUser returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasIsPredefinedUser() bool {
	if o != nil && o.IsPredefinedUser != nil {
		return true
	}

	return false
}

// SetIsPredefinedUser gets a reference to the given bool and assigns it to the IsPredefinedUser field.
func (o *AddSafeMemberResponse) SetIsPredefinedUser(v bool) {
	o.IsPredefinedUser = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *AddSafeMemberResponse) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AddSafeMemberResponse) GetPermissions() map[string]bool {
	if o == nil || o.Permissions == nil {
		var ret map[string]bool
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSafeMemberResponse) GetPermissionsOk() (*map[string]bool, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AddSafeMemberResponse) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]bool and assigns it to the Permissions field.
func (o *AddSafeMemberResponse) SetPermissions(v map[string]bool) {
	o.Permissions = &v
}

func (o AddSafeMemberResponse) MarshalJSON() ([]byte, error) {
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
	if o.MemberId != nil {
		toSerialize["memberId"] = o.MemberId
	}
	if o.MemberName != nil {
		toSerialize["memberName"] = o.MemberName
	}
	if o.MemberType != nil {
		toSerialize["memberType"] = o.MemberType
	}
	if o.MembershipExpirationDate != nil {
		toSerialize["membershipExpirationDate"] = o.MembershipExpirationDate
	}
	if o.IsExpiredMembershipEnable != nil {
		toSerialize["isExpiredMembershipEnable"] = o.IsExpiredMembershipEnable
	}
	if o.IsPredefinedUser != nil {
		toSerialize["isPredefinedUser"] = o.IsPredefinedUser
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAddSafeMemberResponse struct {
	value *AddSafeMemberResponse
	isSet bool
}

func (v NullableAddSafeMemberResponse) Get() *AddSafeMemberResponse {
	return v.value
}

func (v *NullableAddSafeMemberResponse) Set(val *AddSafeMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSafeMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSafeMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSafeMemberResponse(val *AddSafeMemberResponse) *NullableAddSafeMemberResponse {
	return &NullableAddSafeMemberResponse{value: val, isSet: true}
}

func (v NullableAddSafeMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSafeMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


