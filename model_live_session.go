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

// LiveSession struct for LiveSession
type LiveSession struct {
	CanTerminate *bool `json:"CanTerminate,omitempty"`
	CanMonitor *bool `json:"CanMonitor,omitempty"`
	CanSuspend *bool `json:"CanSuspend,omitempty"`
	SessionID *string `json:"SessionID,omitempty"`
	SessionGuid *string `json:"SessionGuid,omitempty"`
	SafeName *string `json:"SafeName,omitempty"`
	FolderName *string `json:"FolderName,omitempty"`
	IsLive *bool `json:"IsLive,omitempty"`
	FileName *string `json:"FileName,omitempty"`
	Start *int64 `json:"Start,omitempty"`
	End *int64 `json:"End,omitempty"`
	Duration *int64 `json:"Duration,omitempty"`
	User *string `json:"User,omitempty"`
	RemoteMachine *string `json:"RemoteMachine,omitempty"`
	ProtectionDate *int64 `json:"ProtectionDate,omitempty"`
	ProtectedBy *string `json:"ProtectedBy,omitempty"`
	ProtectionEnabled *bool `json:"ProtectionEnabled,omitempty"`
	AccountUsername *string `json:"AccountUsername,omitempty"`
	AccountPlatformID *string `json:"AccountPlatformID,omitempty"`
	AccountAddress *string `json:"AccountAddress,omitempty"`
	PIMSuCommand *string `json:"PIMSuCommand,omitempty"`
	PIMSuCWD *string `json:"PIMSuCWD,omitempty"`
	ConnectionComponentID *string `json:"ConnectionComponentID,omitempty"`
	PSMRecordingEntity *string `json:"PSMRecordingEntity,omitempty"`
	TicketID *string `json:"TicketID,omitempty"`
	FromIP *string `json:"FromIP,omitempty"`
	Protocol *string `json:"Protocol,omitempty"`
	Client *string `json:"Client,omitempty"`
	RiskScore *int32 `json:"RiskScore,omitempty"`
	Severity *string `json:"Severity,omitempty"`
	IncidentDetails *PTAData `json:"IncidentDetails,omitempty"`
	RawProperties *map[string]string `json:"RawProperties,omitempty"`
	RecordingFiles []RecordingFile `json:"RecordingFiles,omitempty"`
	RecordedActivities []RecordedCommand `json:"RecordedActivities,omitempty"`
	VideoSize *int64 `json:"VideoSize,omitempty"`
	TextSize *int64 `json:"TextSize,omitempty"`
	DetailsUrl *string `json:"DetailsUrl,omitempty"`
}

// NewLiveSession instantiates a new LiveSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveSession() *LiveSession {
	this := LiveSession{}
	return &this
}

// NewLiveSessionWithDefaults instantiates a new LiveSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveSessionWithDefaults() *LiveSession {
	this := LiveSession{}
	return &this
}

// GetCanTerminate returns the CanTerminate field value if set, zero value otherwise.
func (o *LiveSession) GetCanTerminate() bool {
	if o == nil || o.CanTerminate == nil {
		var ret bool
		return ret
	}
	return *o.CanTerminate
}

// GetCanTerminateOk returns a tuple with the CanTerminate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetCanTerminateOk() (*bool, bool) {
	if o == nil || o.CanTerminate == nil {
		return nil, false
	}
	return o.CanTerminate, true
}

// HasCanTerminate returns a boolean if a field has been set.
func (o *LiveSession) HasCanTerminate() bool {
	if o != nil && o.CanTerminate != nil {
		return true
	}

	return false
}

// SetCanTerminate gets a reference to the given bool and assigns it to the CanTerminate field.
func (o *LiveSession) SetCanTerminate(v bool) {
	o.CanTerminate = &v
}

// GetCanMonitor returns the CanMonitor field value if set, zero value otherwise.
func (o *LiveSession) GetCanMonitor() bool {
	if o == nil || o.CanMonitor == nil {
		var ret bool
		return ret
	}
	return *o.CanMonitor
}

// GetCanMonitorOk returns a tuple with the CanMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetCanMonitorOk() (*bool, bool) {
	if o == nil || o.CanMonitor == nil {
		return nil, false
	}
	return o.CanMonitor, true
}

// HasCanMonitor returns a boolean if a field has been set.
func (o *LiveSession) HasCanMonitor() bool {
	if o != nil && o.CanMonitor != nil {
		return true
	}

	return false
}

// SetCanMonitor gets a reference to the given bool and assigns it to the CanMonitor field.
func (o *LiveSession) SetCanMonitor(v bool) {
	o.CanMonitor = &v
}

// GetCanSuspend returns the CanSuspend field value if set, zero value otherwise.
func (o *LiveSession) GetCanSuspend() bool {
	if o == nil || o.CanSuspend == nil {
		var ret bool
		return ret
	}
	return *o.CanSuspend
}

// GetCanSuspendOk returns a tuple with the CanSuspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetCanSuspendOk() (*bool, bool) {
	if o == nil || o.CanSuspend == nil {
		return nil, false
	}
	return o.CanSuspend, true
}

// HasCanSuspend returns a boolean if a field has been set.
func (o *LiveSession) HasCanSuspend() bool {
	if o != nil && o.CanSuspend != nil {
		return true
	}

	return false
}

// SetCanSuspend gets a reference to the given bool and assigns it to the CanSuspend field.
func (o *LiveSession) SetCanSuspend(v bool) {
	o.CanSuspend = &v
}

// GetSessionID returns the SessionID field value if set, zero value otherwise.
func (o *LiveSession) GetSessionID() string {
	if o == nil || o.SessionID == nil {
		var ret string
		return ret
	}
	return *o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetSessionIDOk() (*string, bool) {
	if o == nil || o.SessionID == nil {
		return nil, false
	}
	return o.SessionID, true
}

// HasSessionID returns a boolean if a field has been set.
func (o *LiveSession) HasSessionID() bool {
	if o != nil && o.SessionID != nil {
		return true
	}

	return false
}

// SetSessionID gets a reference to the given string and assigns it to the SessionID field.
func (o *LiveSession) SetSessionID(v string) {
	o.SessionID = &v
}

// GetSessionGuid returns the SessionGuid field value if set, zero value otherwise.
func (o *LiveSession) GetSessionGuid() string {
	if o == nil || o.SessionGuid == nil {
		var ret string
		return ret
	}
	return *o.SessionGuid
}

// GetSessionGuidOk returns a tuple with the SessionGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetSessionGuidOk() (*string, bool) {
	if o == nil || o.SessionGuid == nil {
		return nil, false
	}
	return o.SessionGuid, true
}

// HasSessionGuid returns a boolean if a field has been set.
func (o *LiveSession) HasSessionGuid() bool {
	if o != nil && o.SessionGuid != nil {
		return true
	}

	return false
}

// SetSessionGuid gets a reference to the given string and assigns it to the SessionGuid field.
func (o *LiveSession) SetSessionGuid(v string) {
	o.SessionGuid = &v
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *LiveSession) GetSafeName() string {
	if o == nil || o.SafeName == nil {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetSafeNameOk() (*string, bool) {
	if o == nil || o.SafeName == nil {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *LiveSession) HasSafeName() bool {
	if o != nil && o.SafeName != nil {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *LiveSession) SetSafeName(v string) {
	o.SafeName = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *LiveSession) GetFolderName() string {
	if o == nil || o.FolderName == nil {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetFolderNameOk() (*string, bool) {
	if o == nil || o.FolderName == nil {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *LiveSession) HasFolderName() bool {
	if o != nil && o.FolderName != nil {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *LiveSession) SetFolderName(v string) {
	o.FolderName = &v
}

// GetIsLive returns the IsLive field value if set, zero value otherwise.
func (o *LiveSession) GetIsLive() bool {
	if o == nil || o.IsLive == nil {
		var ret bool
		return ret
	}
	return *o.IsLive
}

// GetIsLiveOk returns a tuple with the IsLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetIsLiveOk() (*bool, bool) {
	if o == nil || o.IsLive == nil {
		return nil, false
	}
	return o.IsLive, true
}

// HasIsLive returns a boolean if a field has been set.
func (o *LiveSession) HasIsLive() bool {
	if o != nil && o.IsLive != nil {
		return true
	}

	return false
}

// SetIsLive gets a reference to the given bool and assigns it to the IsLive field.
func (o *LiveSession) SetIsLive(v bool) {
	o.IsLive = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *LiveSession) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *LiveSession) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *LiveSession) SetFileName(v string) {
	o.FileName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *LiveSession) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *LiveSession) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *LiveSession) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *LiveSession) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *LiveSession) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *LiveSession) SetEnd(v int64) {
	o.End = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LiveSession) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LiveSession) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *LiveSession) SetDuration(v int64) {
	o.Duration = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LiveSession) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LiveSession) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *LiveSession) SetUser(v string) {
	o.User = &v
}

// GetRemoteMachine returns the RemoteMachine field value if set, zero value otherwise.
func (o *LiveSession) GetRemoteMachine() string {
	if o == nil || o.RemoteMachine == nil {
		var ret string
		return ret
	}
	return *o.RemoteMachine
}

// GetRemoteMachineOk returns a tuple with the RemoteMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetRemoteMachineOk() (*string, bool) {
	if o == nil || o.RemoteMachine == nil {
		return nil, false
	}
	return o.RemoteMachine, true
}

// HasRemoteMachine returns a boolean if a field has been set.
func (o *LiveSession) HasRemoteMachine() bool {
	if o != nil && o.RemoteMachine != nil {
		return true
	}

	return false
}

// SetRemoteMachine gets a reference to the given string and assigns it to the RemoteMachine field.
func (o *LiveSession) SetRemoteMachine(v string) {
	o.RemoteMachine = &v
}

// GetProtectionDate returns the ProtectionDate field value if set, zero value otherwise.
func (o *LiveSession) GetProtectionDate() int64 {
	if o == nil || o.ProtectionDate == nil {
		var ret int64
		return ret
	}
	return *o.ProtectionDate
}

// GetProtectionDateOk returns a tuple with the ProtectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetProtectionDateOk() (*int64, bool) {
	if o == nil || o.ProtectionDate == nil {
		return nil, false
	}
	return o.ProtectionDate, true
}

// HasProtectionDate returns a boolean if a field has been set.
func (o *LiveSession) HasProtectionDate() bool {
	if o != nil && o.ProtectionDate != nil {
		return true
	}

	return false
}

// SetProtectionDate gets a reference to the given int64 and assigns it to the ProtectionDate field.
func (o *LiveSession) SetProtectionDate(v int64) {
	o.ProtectionDate = &v
}

// GetProtectedBy returns the ProtectedBy field value if set, zero value otherwise.
func (o *LiveSession) GetProtectedBy() string {
	if o == nil || o.ProtectedBy == nil {
		var ret string
		return ret
	}
	return *o.ProtectedBy
}

// GetProtectedByOk returns a tuple with the ProtectedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetProtectedByOk() (*string, bool) {
	if o == nil || o.ProtectedBy == nil {
		return nil, false
	}
	return o.ProtectedBy, true
}

// HasProtectedBy returns a boolean if a field has been set.
func (o *LiveSession) HasProtectedBy() bool {
	if o != nil && o.ProtectedBy != nil {
		return true
	}

	return false
}

// SetProtectedBy gets a reference to the given string and assigns it to the ProtectedBy field.
func (o *LiveSession) SetProtectedBy(v string) {
	o.ProtectedBy = &v
}

// GetProtectionEnabled returns the ProtectionEnabled field value if set, zero value otherwise.
func (o *LiveSession) GetProtectionEnabled() bool {
	if o == nil || o.ProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ProtectionEnabled
}

// GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetProtectionEnabledOk() (*bool, bool) {
	if o == nil || o.ProtectionEnabled == nil {
		return nil, false
	}
	return o.ProtectionEnabled, true
}

// HasProtectionEnabled returns a boolean if a field has been set.
func (o *LiveSession) HasProtectionEnabled() bool {
	if o != nil && o.ProtectionEnabled != nil {
		return true
	}

	return false
}

// SetProtectionEnabled gets a reference to the given bool and assigns it to the ProtectionEnabled field.
func (o *LiveSession) SetProtectionEnabled(v bool) {
	o.ProtectionEnabled = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *LiveSession) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *LiveSession) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *LiveSession) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetAccountPlatformID returns the AccountPlatformID field value if set, zero value otherwise.
func (o *LiveSession) GetAccountPlatformID() string {
	if o == nil || o.AccountPlatformID == nil {
		var ret string
		return ret
	}
	return *o.AccountPlatformID
}

// GetAccountPlatformIDOk returns a tuple with the AccountPlatformID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetAccountPlatformIDOk() (*string, bool) {
	if o == nil || o.AccountPlatformID == nil {
		return nil, false
	}
	return o.AccountPlatformID, true
}

// HasAccountPlatformID returns a boolean if a field has been set.
func (o *LiveSession) HasAccountPlatformID() bool {
	if o != nil && o.AccountPlatformID != nil {
		return true
	}

	return false
}

// SetAccountPlatformID gets a reference to the given string and assigns it to the AccountPlatformID field.
func (o *LiveSession) SetAccountPlatformID(v string) {
	o.AccountPlatformID = &v
}

// GetAccountAddress returns the AccountAddress field value if set, zero value otherwise.
func (o *LiveSession) GetAccountAddress() string {
	if o == nil || o.AccountAddress == nil {
		var ret string
		return ret
	}
	return *o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetAccountAddressOk() (*string, bool) {
	if o == nil || o.AccountAddress == nil {
		return nil, false
	}
	return o.AccountAddress, true
}

// HasAccountAddress returns a boolean if a field has been set.
func (o *LiveSession) HasAccountAddress() bool {
	if o != nil && o.AccountAddress != nil {
		return true
	}

	return false
}

// SetAccountAddress gets a reference to the given string and assigns it to the AccountAddress field.
func (o *LiveSession) SetAccountAddress(v string) {
	o.AccountAddress = &v
}

// GetPIMSuCommand returns the PIMSuCommand field value if set, zero value otherwise.
func (o *LiveSession) GetPIMSuCommand() string {
	if o == nil || o.PIMSuCommand == nil {
		var ret string
		return ret
	}
	return *o.PIMSuCommand
}

// GetPIMSuCommandOk returns a tuple with the PIMSuCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetPIMSuCommandOk() (*string, bool) {
	if o == nil || o.PIMSuCommand == nil {
		return nil, false
	}
	return o.PIMSuCommand, true
}

// HasPIMSuCommand returns a boolean if a field has been set.
func (o *LiveSession) HasPIMSuCommand() bool {
	if o != nil && o.PIMSuCommand != nil {
		return true
	}

	return false
}

// SetPIMSuCommand gets a reference to the given string and assigns it to the PIMSuCommand field.
func (o *LiveSession) SetPIMSuCommand(v string) {
	o.PIMSuCommand = &v
}

// GetPIMSuCWD returns the PIMSuCWD field value if set, zero value otherwise.
func (o *LiveSession) GetPIMSuCWD() string {
	if o == nil || o.PIMSuCWD == nil {
		var ret string
		return ret
	}
	return *o.PIMSuCWD
}

// GetPIMSuCWDOk returns a tuple with the PIMSuCWD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetPIMSuCWDOk() (*string, bool) {
	if o == nil || o.PIMSuCWD == nil {
		return nil, false
	}
	return o.PIMSuCWD, true
}

// HasPIMSuCWD returns a boolean if a field has been set.
func (o *LiveSession) HasPIMSuCWD() bool {
	if o != nil && o.PIMSuCWD != nil {
		return true
	}

	return false
}

// SetPIMSuCWD gets a reference to the given string and assigns it to the PIMSuCWD field.
func (o *LiveSession) SetPIMSuCWD(v string) {
	o.PIMSuCWD = &v
}

// GetConnectionComponentID returns the ConnectionComponentID field value if set, zero value otherwise.
func (o *LiveSession) GetConnectionComponentID() string {
	if o == nil || o.ConnectionComponentID == nil {
		var ret string
		return ret
	}
	return *o.ConnectionComponentID
}

// GetConnectionComponentIDOk returns a tuple with the ConnectionComponentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetConnectionComponentIDOk() (*string, bool) {
	if o == nil || o.ConnectionComponentID == nil {
		return nil, false
	}
	return o.ConnectionComponentID, true
}

// HasConnectionComponentID returns a boolean if a field has been set.
func (o *LiveSession) HasConnectionComponentID() bool {
	if o != nil && o.ConnectionComponentID != nil {
		return true
	}

	return false
}

// SetConnectionComponentID gets a reference to the given string and assigns it to the ConnectionComponentID field.
func (o *LiveSession) SetConnectionComponentID(v string) {
	o.ConnectionComponentID = &v
}

// GetPSMRecordingEntity returns the PSMRecordingEntity field value if set, zero value otherwise.
func (o *LiveSession) GetPSMRecordingEntity() string {
	if o == nil || o.PSMRecordingEntity == nil {
		var ret string
		return ret
	}
	return *o.PSMRecordingEntity
}

// GetPSMRecordingEntityOk returns a tuple with the PSMRecordingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetPSMRecordingEntityOk() (*string, bool) {
	if o == nil || o.PSMRecordingEntity == nil {
		return nil, false
	}
	return o.PSMRecordingEntity, true
}

// HasPSMRecordingEntity returns a boolean if a field has been set.
func (o *LiveSession) HasPSMRecordingEntity() bool {
	if o != nil && o.PSMRecordingEntity != nil {
		return true
	}

	return false
}

// SetPSMRecordingEntity gets a reference to the given string and assigns it to the PSMRecordingEntity field.
func (o *LiveSession) SetPSMRecordingEntity(v string) {
	o.PSMRecordingEntity = &v
}

// GetTicketID returns the TicketID field value if set, zero value otherwise.
func (o *LiveSession) GetTicketID() string {
	if o == nil || o.TicketID == nil {
		var ret string
		return ret
	}
	return *o.TicketID
}

// GetTicketIDOk returns a tuple with the TicketID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetTicketIDOk() (*string, bool) {
	if o == nil || o.TicketID == nil {
		return nil, false
	}
	return o.TicketID, true
}

// HasTicketID returns a boolean if a field has been set.
func (o *LiveSession) HasTicketID() bool {
	if o != nil && o.TicketID != nil {
		return true
	}

	return false
}

// SetTicketID gets a reference to the given string and assigns it to the TicketID field.
func (o *LiveSession) SetTicketID(v string) {
	o.TicketID = &v
}

// GetFromIP returns the FromIP field value if set, zero value otherwise.
func (o *LiveSession) GetFromIP() string {
	if o == nil || o.FromIP == nil {
		var ret string
		return ret
	}
	return *o.FromIP
}

// GetFromIPOk returns a tuple with the FromIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetFromIPOk() (*string, bool) {
	if o == nil || o.FromIP == nil {
		return nil, false
	}
	return o.FromIP, true
}

// HasFromIP returns a boolean if a field has been set.
func (o *LiveSession) HasFromIP() bool {
	if o != nil && o.FromIP != nil {
		return true
	}

	return false
}

// SetFromIP gets a reference to the given string and assigns it to the FromIP field.
func (o *LiveSession) SetFromIP(v string) {
	o.FromIP = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *LiveSession) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *LiveSession) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *LiveSession) SetProtocol(v string) {
	o.Protocol = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *LiveSession) GetClient() string {
	if o == nil || o.Client == nil {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetClientOk() (*string, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *LiveSession) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *LiveSession) SetClient(v string) {
	o.Client = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *LiveSession) GetRiskScore() int32 {
	if o == nil || o.RiskScore == nil {
		var ret int32
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetRiskScoreOk() (*int32, bool) {
	if o == nil || o.RiskScore == nil {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *LiveSession) HasRiskScore() bool {
	if o != nil && o.RiskScore != nil {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given int32 and assigns it to the RiskScore field.
func (o *LiveSession) SetRiskScore(v int32) {
	o.RiskScore = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *LiveSession) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *LiveSession) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *LiveSession) SetSeverity(v string) {
	o.Severity = &v
}

// GetIncidentDetails returns the IncidentDetails field value if set, zero value otherwise.
func (o *LiveSession) GetIncidentDetails() PTAData {
	if o == nil || o.IncidentDetails == nil {
		var ret PTAData
		return ret
	}
	return *o.IncidentDetails
}

// GetIncidentDetailsOk returns a tuple with the IncidentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetIncidentDetailsOk() (*PTAData, bool) {
	if o == nil || o.IncidentDetails == nil {
		return nil, false
	}
	return o.IncidentDetails, true
}

// HasIncidentDetails returns a boolean if a field has been set.
func (o *LiveSession) HasIncidentDetails() bool {
	if o != nil && o.IncidentDetails != nil {
		return true
	}

	return false
}

// SetIncidentDetails gets a reference to the given PTAData and assigns it to the IncidentDetails field.
func (o *LiveSession) SetIncidentDetails(v PTAData) {
	o.IncidentDetails = &v
}

// GetRawProperties returns the RawProperties field value if set, zero value otherwise.
func (o *LiveSession) GetRawProperties() map[string]string {
	if o == nil || o.RawProperties == nil {
		var ret map[string]string
		return ret
	}
	return *o.RawProperties
}

// GetRawPropertiesOk returns a tuple with the RawProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetRawPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.RawProperties == nil {
		return nil, false
	}
	return o.RawProperties, true
}

// HasRawProperties returns a boolean if a field has been set.
func (o *LiveSession) HasRawProperties() bool {
	if o != nil && o.RawProperties != nil {
		return true
	}

	return false
}

// SetRawProperties gets a reference to the given map[string]string and assigns it to the RawProperties field.
func (o *LiveSession) SetRawProperties(v map[string]string) {
	o.RawProperties = &v
}

// GetRecordingFiles returns the RecordingFiles field value if set, zero value otherwise.
func (o *LiveSession) GetRecordingFiles() []RecordingFile {
	if o == nil || o.RecordingFiles == nil {
		var ret []RecordingFile
		return ret
	}
	return o.RecordingFiles
}

// GetRecordingFilesOk returns a tuple with the RecordingFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetRecordingFilesOk() ([]RecordingFile, bool) {
	if o == nil || o.RecordingFiles == nil {
		return nil, false
	}
	return o.RecordingFiles, true
}

// HasRecordingFiles returns a boolean if a field has been set.
func (o *LiveSession) HasRecordingFiles() bool {
	if o != nil && o.RecordingFiles != nil {
		return true
	}

	return false
}

// SetRecordingFiles gets a reference to the given []RecordingFile and assigns it to the RecordingFiles field.
func (o *LiveSession) SetRecordingFiles(v []RecordingFile) {
	o.RecordingFiles = v
}

// GetRecordedActivities returns the RecordedActivities field value if set, zero value otherwise.
func (o *LiveSession) GetRecordedActivities() []RecordedCommand {
	if o == nil || o.RecordedActivities == nil {
		var ret []RecordedCommand
		return ret
	}
	return o.RecordedActivities
}

// GetRecordedActivitiesOk returns a tuple with the RecordedActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetRecordedActivitiesOk() ([]RecordedCommand, bool) {
	if o == nil || o.RecordedActivities == nil {
		return nil, false
	}
	return o.RecordedActivities, true
}

// HasRecordedActivities returns a boolean if a field has been set.
func (o *LiveSession) HasRecordedActivities() bool {
	if o != nil && o.RecordedActivities != nil {
		return true
	}

	return false
}

// SetRecordedActivities gets a reference to the given []RecordedCommand and assigns it to the RecordedActivities field.
func (o *LiveSession) SetRecordedActivities(v []RecordedCommand) {
	o.RecordedActivities = v
}

// GetVideoSize returns the VideoSize field value if set, zero value otherwise.
func (o *LiveSession) GetVideoSize() int64 {
	if o == nil || o.VideoSize == nil {
		var ret int64
		return ret
	}
	return *o.VideoSize
}

// GetVideoSizeOk returns a tuple with the VideoSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetVideoSizeOk() (*int64, bool) {
	if o == nil || o.VideoSize == nil {
		return nil, false
	}
	return o.VideoSize, true
}

// HasVideoSize returns a boolean if a field has been set.
func (o *LiveSession) HasVideoSize() bool {
	if o != nil && o.VideoSize != nil {
		return true
	}

	return false
}

// SetVideoSize gets a reference to the given int64 and assigns it to the VideoSize field.
func (o *LiveSession) SetVideoSize(v int64) {
	o.VideoSize = &v
}

// GetTextSize returns the TextSize field value if set, zero value otherwise.
func (o *LiveSession) GetTextSize() int64 {
	if o == nil || o.TextSize == nil {
		var ret int64
		return ret
	}
	return *o.TextSize
}

// GetTextSizeOk returns a tuple with the TextSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetTextSizeOk() (*int64, bool) {
	if o == nil || o.TextSize == nil {
		return nil, false
	}
	return o.TextSize, true
}

// HasTextSize returns a boolean if a field has been set.
func (o *LiveSession) HasTextSize() bool {
	if o != nil && o.TextSize != nil {
		return true
	}

	return false
}

// SetTextSize gets a reference to the given int64 and assigns it to the TextSize field.
func (o *LiveSession) SetTextSize(v int64) {
	o.TextSize = &v
}

// GetDetailsUrl returns the DetailsUrl field value if set, zero value otherwise.
func (o *LiveSession) GetDetailsUrl() string {
	if o == nil || o.DetailsUrl == nil {
		var ret string
		return ret
	}
	return *o.DetailsUrl
}

// GetDetailsUrlOk returns a tuple with the DetailsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveSession) GetDetailsUrlOk() (*string, bool) {
	if o == nil || o.DetailsUrl == nil {
		return nil, false
	}
	return o.DetailsUrl, true
}

// HasDetailsUrl returns a boolean if a field has been set.
func (o *LiveSession) HasDetailsUrl() bool {
	if o != nil && o.DetailsUrl != nil {
		return true
	}

	return false
}

// SetDetailsUrl gets a reference to the given string and assigns it to the DetailsUrl field.
func (o *LiveSession) SetDetailsUrl(v string) {
	o.DetailsUrl = &v
}

func (o LiveSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanTerminate != nil {
		toSerialize["CanTerminate"] = o.CanTerminate
	}
	if o.CanMonitor != nil {
		toSerialize["CanMonitor"] = o.CanMonitor
	}
	if o.CanSuspend != nil {
		toSerialize["CanSuspend"] = o.CanSuspend
	}
	if o.SessionID != nil {
		toSerialize["SessionID"] = o.SessionID
	}
	if o.SessionGuid != nil {
		toSerialize["SessionGuid"] = o.SessionGuid
	}
	if o.SafeName != nil {
		toSerialize["SafeName"] = o.SafeName
	}
	if o.FolderName != nil {
		toSerialize["FolderName"] = o.FolderName
	}
	if o.IsLive != nil {
		toSerialize["IsLive"] = o.IsLive
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.Start != nil {
		toSerialize["Start"] = o.Start
	}
	if o.End != nil {
		toSerialize["End"] = o.End
	}
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	if o.RemoteMachine != nil {
		toSerialize["RemoteMachine"] = o.RemoteMachine
	}
	if o.ProtectionDate != nil {
		toSerialize["ProtectionDate"] = o.ProtectionDate
	}
	if o.ProtectedBy != nil {
		toSerialize["ProtectedBy"] = o.ProtectedBy
	}
	if o.ProtectionEnabled != nil {
		toSerialize["ProtectionEnabled"] = o.ProtectionEnabled
	}
	if o.AccountUsername != nil {
		toSerialize["AccountUsername"] = o.AccountUsername
	}
	if o.AccountPlatformID != nil {
		toSerialize["AccountPlatformID"] = o.AccountPlatformID
	}
	if o.AccountAddress != nil {
		toSerialize["AccountAddress"] = o.AccountAddress
	}
	if o.PIMSuCommand != nil {
		toSerialize["PIMSuCommand"] = o.PIMSuCommand
	}
	if o.PIMSuCWD != nil {
		toSerialize["PIMSuCWD"] = o.PIMSuCWD
	}
	if o.ConnectionComponentID != nil {
		toSerialize["ConnectionComponentID"] = o.ConnectionComponentID
	}
	if o.PSMRecordingEntity != nil {
		toSerialize["PSMRecordingEntity"] = o.PSMRecordingEntity
	}
	if o.TicketID != nil {
		toSerialize["TicketID"] = o.TicketID
	}
	if o.FromIP != nil {
		toSerialize["FromIP"] = o.FromIP
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Client != nil {
		toSerialize["Client"] = o.Client
	}
	if o.RiskScore != nil {
		toSerialize["RiskScore"] = o.RiskScore
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.IncidentDetails != nil {
		toSerialize["IncidentDetails"] = o.IncidentDetails
	}
	if o.RawProperties != nil {
		toSerialize["RawProperties"] = o.RawProperties
	}
	if o.RecordingFiles != nil {
		toSerialize["RecordingFiles"] = o.RecordingFiles
	}
	if o.RecordedActivities != nil {
		toSerialize["RecordedActivities"] = o.RecordedActivities
	}
	if o.VideoSize != nil {
		toSerialize["VideoSize"] = o.VideoSize
	}
	if o.TextSize != nil {
		toSerialize["TextSize"] = o.TextSize
	}
	if o.DetailsUrl != nil {
		toSerialize["DetailsUrl"] = o.DetailsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableLiveSession struct {
	value *LiveSession
	isSet bool
}

func (v NullableLiveSession) Get() *LiveSession {
	return v.value
}

func (v *NullableLiveSession) Set(val *LiveSession) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveSession) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveSession(val *LiveSession) *NullableLiveSession {
	return &NullableLiveSession{value: val, isSet: true}
}

func (v NullableLiveSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

