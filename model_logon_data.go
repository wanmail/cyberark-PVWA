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

// LogonData struct for LogonData
type LogonData struct {
	// The name of the user who will logon to the Vault
	UserName *string `json:"UserName,omitempty"`
	// The password of the user
	Password *string `json:"Password,omitempty"`
	// The new password of the user. This parameter is optional, and enables you to change a password  This parameter can be used with the following authentication methods:  CyberArk,  LDAP
	NewPassword *string `json:"NewPassword,omitempty"`
	Type *string `json:"Type,omitempty"`
	AdditionalInfo *string `json:"AdditionalInfo,omitempty"`
	SecureMode *bool `json:"SecureMode,omitempty"`
	// Set this parameter to true In order to allow more than one connection for the same user simultaneously.    Up to 300 concurrent sessions.
	ConcurrentSession *bool `json:"concurrentSession,omitempty"`
}

// NewLogonData instantiates a new LogonData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogonData() *LogonData {
	this := LogonData{}
	return &this
}

// NewLogonDataWithDefaults instantiates a new LogonData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogonDataWithDefaults() *LogonData {
	this := LogonData{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *LogonData) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *LogonData) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *LogonData) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LogonData) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LogonData) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LogonData) SetPassword(v string) {
	o.Password = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *LogonData) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *LogonData) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *LogonData) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogonData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogonData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogonData) SetType(v string) {
	o.Type = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *LogonData) GetAdditionalInfo() string {
	if o == nil || o.AdditionalInfo == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || o.AdditionalInfo == nil {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *LogonData) HasAdditionalInfo() bool {
	if o != nil && o.AdditionalInfo != nil {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *LogonData) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetSecureMode returns the SecureMode field value if set, zero value otherwise.
func (o *LogonData) GetSecureMode() bool {
	if o == nil || o.SecureMode == nil {
		var ret bool
		return ret
	}
	return *o.SecureMode
}

// GetSecureModeOk returns a tuple with the SecureMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetSecureModeOk() (*bool, bool) {
	if o == nil || o.SecureMode == nil {
		return nil, false
	}
	return o.SecureMode, true
}

// HasSecureMode returns a boolean if a field has been set.
func (o *LogonData) HasSecureMode() bool {
	if o != nil && o.SecureMode != nil {
		return true
	}

	return false
}

// SetSecureMode gets a reference to the given bool and assigns it to the SecureMode field.
func (o *LogonData) SetSecureMode(v bool) {
	o.SecureMode = &v
}

// GetConcurrentSession returns the ConcurrentSession field value if set, zero value otherwise.
func (o *LogonData) GetConcurrentSession() bool {
	if o == nil || o.ConcurrentSession == nil {
		var ret bool
		return ret
	}
	return *o.ConcurrentSession
}

// GetConcurrentSessionOk returns a tuple with the ConcurrentSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogonData) GetConcurrentSessionOk() (*bool, bool) {
	if o == nil || o.ConcurrentSession == nil {
		return nil, false
	}
	return o.ConcurrentSession, true
}

// HasConcurrentSession returns a boolean if a field has been set.
func (o *LogonData) HasConcurrentSession() bool {
	if o != nil && o.ConcurrentSession != nil {
		return true
	}

	return false
}

// SetConcurrentSession gets a reference to the given bool and assigns it to the ConcurrentSession field.
func (o *LogonData) SetConcurrentSession(v bool) {
	o.ConcurrentSession = &v
}

func (o LogonData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.NewPassword != nil {
		toSerialize["NewPassword"] = o.NewPassword
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.AdditionalInfo != nil {
		toSerialize["AdditionalInfo"] = o.AdditionalInfo
	}
	if o.SecureMode != nil {
		toSerialize["SecureMode"] = o.SecureMode
	}
	if o.ConcurrentSession != nil {
		toSerialize["concurrentSession"] = o.ConcurrentSession
	}
	return json.Marshal(toSerialize)
}

type NullableLogonData struct {
	value *LogonData
	isSet bool
}

func (v NullableLogonData) Get() *LogonData {
	return v.value
}

func (v *NullableLogonData) Set(val *LogonData) {
	v.value = val
	v.isSet = true
}

func (v NullableLogonData) IsSet() bool {
	return v.isSet
}

func (v *NullableLogonData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogonData(val *LogonData) *NullableLogonData {
	return &NullableLogonData{value: val, isSet: true}
}

func (v NullableLogonData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogonData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


