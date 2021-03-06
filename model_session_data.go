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

// SessionData struct for SessionData
type SessionData struct {
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

// NewSessionData instantiates a new SessionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionData() *SessionData {
	this := SessionData{}
	return &this
}

// NewSessionDataWithDefaults instantiates a new SessionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionDataWithDefaults() *SessionData {
	this := SessionData{}
	return &this
}

// GetSessionID returns the SessionID field value if set, zero value otherwise.
func (o *SessionData) GetSessionID() string {
	if o == nil || o.SessionID == nil {
		var ret string
		return ret
	}
	return *o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetSessionIDOk() (*string, bool) {
	if o == nil || o.SessionID == nil {
		return nil, false
	}
	return o.SessionID, true
}

// HasSessionID returns a boolean if a field has been set.
func (o *SessionData) HasSessionID() bool {
	if o != nil && o.SessionID != nil {
		return true
	}

	return false
}

// SetSessionID gets a reference to the given string and assigns it to the SessionID field.
func (o *SessionData) SetSessionID(v string) {
	o.SessionID = &v
}

// GetSessionGuid returns the SessionGuid field value if set, zero value otherwise.
func (o *SessionData) GetSessionGuid() string {
	if o == nil || o.SessionGuid == nil {
		var ret string
		return ret
	}
	return *o.SessionGuid
}

// GetSessionGuidOk returns a tuple with the SessionGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetSessionGuidOk() (*string, bool) {
	if o == nil || o.SessionGuid == nil {
		return nil, false
	}
	return o.SessionGuid, true
}

// HasSessionGuid returns a boolean if a field has been set.
func (o *SessionData) HasSessionGuid() bool {
	if o != nil && o.SessionGuid != nil {
		return true
	}

	return false
}

// SetSessionGuid gets a reference to the given string and assigns it to the SessionGuid field.
func (o *SessionData) SetSessionGuid(v string) {
	o.SessionGuid = &v
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *SessionData) GetSafeName() string {
	if o == nil || o.SafeName == nil {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetSafeNameOk() (*string, bool) {
	if o == nil || o.SafeName == nil {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *SessionData) HasSafeName() bool {
	if o != nil && o.SafeName != nil {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *SessionData) SetSafeName(v string) {
	o.SafeName = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *SessionData) GetFolderName() string {
	if o == nil || o.FolderName == nil {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetFolderNameOk() (*string, bool) {
	if o == nil || o.FolderName == nil {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *SessionData) HasFolderName() bool {
	if o != nil && o.FolderName != nil {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *SessionData) SetFolderName(v string) {
	o.FolderName = &v
}

// GetIsLive returns the IsLive field value if set, zero value otherwise.
func (o *SessionData) GetIsLive() bool {
	if o == nil || o.IsLive == nil {
		var ret bool
		return ret
	}
	return *o.IsLive
}

// GetIsLiveOk returns a tuple with the IsLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetIsLiveOk() (*bool, bool) {
	if o == nil || o.IsLive == nil {
		return nil, false
	}
	return o.IsLive, true
}

// HasIsLive returns a boolean if a field has been set.
func (o *SessionData) HasIsLive() bool {
	if o != nil && o.IsLive != nil {
		return true
	}

	return false
}

// SetIsLive gets a reference to the given bool and assigns it to the IsLive field.
func (o *SessionData) SetIsLive(v bool) {
	o.IsLive = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *SessionData) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *SessionData) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *SessionData) SetFileName(v string) {
	o.FileName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SessionData) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SessionData) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *SessionData) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SessionData) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SessionData) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *SessionData) SetEnd(v int64) {
	o.End = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SessionData) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SessionData) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *SessionData) SetDuration(v int64) {
	o.Duration = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SessionData) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SessionData) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *SessionData) SetUser(v string) {
	o.User = &v
}

// GetRemoteMachine returns the RemoteMachine field value if set, zero value otherwise.
func (o *SessionData) GetRemoteMachine() string {
	if o == nil || o.RemoteMachine == nil {
		var ret string
		return ret
	}
	return *o.RemoteMachine
}

// GetRemoteMachineOk returns a tuple with the RemoteMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetRemoteMachineOk() (*string, bool) {
	if o == nil || o.RemoteMachine == nil {
		return nil, false
	}
	return o.RemoteMachine, true
}

// HasRemoteMachine returns a boolean if a field has been set.
func (o *SessionData) HasRemoteMachine() bool {
	if o != nil && o.RemoteMachine != nil {
		return true
	}

	return false
}

// SetRemoteMachine gets a reference to the given string and assigns it to the RemoteMachine field.
func (o *SessionData) SetRemoteMachine(v string) {
	o.RemoteMachine = &v
}

// GetProtectionDate returns the ProtectionDate field value if set, zero value otherwise.
func (o *SessionData) GetProtectionDate() int64 {
	if o == nil || o.ProtectionDate == nil {
		var ret int64
		return ret
	}
	return *o.ProtectionDate
}

// GetProtectionDateOk returns a tuple with the ProtectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetProtectionDateOk() (*int64, bool) {
	if o == nil || o.ProtectionDate == nil {
		return nil, false
	}
	return o.ProtectionDate, true
}

// HasProtectionDate returns a boolean if a field has been set.
func (o *SessionData) HasProtectionDate() bool {
	if o != nil && o.ProtectionDate != nil {
		return true
	}

	return false
}

// SetProtectionDate gets a reference to the given int64 and assigns it to the ProtectionDate field.
func (o *SessionData) SetProtectionDate(v int64) {
	o.ProtectionDate = &v
}

// GetProtectedBy returns the ProtectedBy field value if set, zero value otherwise.
func (o *SessionData) GetProtectedBy() string {
	if o == nil || o.ProtectedBy == nil {
		var ret string
		return ret
	}
	return *o.ProtectedBy
}

// GetProtectedByOk returns a tuple with the ProtectedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetProtectedByOk() (*string, bool) {
	if o == nil || o.ProtectedBy == nil {
		return nil, false
	}
	return o.ProtectedBy, true
}

// HasProtectedBy returns a boolean if a field has been set.
func (o *SessionData) HasProtectedBy() bool {
	if o != nil && o.ProtectedBy != nil {
		return true
	}

	return false
}

// SetProtectedBy gets a reference to the given string and assigns it to the ProtectedBy field.
func (o *SessionData) SetProtectedBy(v string) {
	o.ProtectedBy = &v
}

// GetProtectionEnabled returns the ProtectionEnabled field value if set, zero value otherwise.
func (o *SessionData) GetProtectionEnabled() bool {
	if o == nil || o.ProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ProtectionEnabled
}

// GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetProtectionEnabledOk() (*bool, bool) {
	if o == nil || o.ProtectionEnabled == nil {
		return nil, false
	}
	return o.ProtectionEnabled, true
}

// HasProtectionEnabled returns a boolean if a field has been set.
func (o *SessionData) HasProtectionEnabled() bool {
	if o != nil && o.ProtectionEnabled != nil {
		return true
	}

	return false
}

// SetProtectionEnabled gets a reference to the given bool and assigns it to the ProtectionEnabled field.
func (o *SessionData) SetProtectionEnabled(v bool) {
	o.ProtectionEnabled = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *SessionData) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *SessionData) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *SessionData) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetAccountPlatformID returns the AccountPlatformID field value if set, zero value otherwise.
func (o *SessionData) GetAccountPlatformID() string {
	if o == nil || o.AccountPlatformID == nil {
		var ret string
		return ret
	}
	return *o.AccountPlatformID
}

// GetAccountPlatformIDOk returns a tuple with the AccountPlatformID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetAccountPlatformIDOk() (*string, bool) {
	if o == nil || o.AccountPlatformID == nil {
		return nil, false
	}
	return o.AccountPlatformID, true
}

// HasAccountPlatformID returns a boolean if a field has been set.
func (o *SessionData) HasAccountPlatformID() bool {
	if o != nil && o.AccountPlatformID != nil {
		return true
	}

	return false
}

// SetAccountPlatformID gets a reference to the given string and assigns it to the AccountPlatformID field.
func (o *SessionData) SetAccountPlatformID(v string) {
	o.AccountPlatformID = &v
}

// GetAccountAddress returns the AccountAddress field value if set, zero value otherwise.
func (o *SessionData) GetAccountAddress() string {
	if o == nil || o.AccountAddress == nil {
		var ret string
		return ret
	}
	return *o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetAccountAddressOk() (*string, bool) {
	if o == nil || o.AccountAddress == nil {
		return nil, false
	}
	return o.AccountAddress, true
}

// HasAccountAddress returns a boolean if a field has been set.
func (o *SessionData) HasAccountAddress() bool {
	if o != nil && o.AccountAddress != nil {
		return true
	}

	return false
}

// SetAccountAddress gets a reference to the given string and assigns it to the AccountAddress field.
func (o *SessionData) SetAccountAddress(v string) {
	o.AccountAddress = &v
}

// GetPIMSuCommand returns the PIMSuCommand field value if set, zero value otherwise.
func (o *SessionData) GetPIMSuCommand() string {
	if o == nil || o.PIMSuCommand == nil {
		var ret string
		return ret
	}
	return *o.PIMSuCommand
}

// GetPIMSuCommandOk returns a tuple with the PIMSuCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetPIMSuCommandOk() (*string, bool) {
	if o == nil || o.PIMSuCommand == nil {
		return nil, false
	}
	return o.PIMSuCommand, true
}

// HasPIMSuCommand returns a boolean if a field has been set.
func (o *SessionData) HasPIMSuCommand() bool {
	if o != nil && o.PIMSuCommand != nil {
		return true
	}

	return false
}

// SetPIMSuCommand gets a reference to the given string and assigns it to the PIMSuCommand field.
func (o *SessionData) SetPIMSuCommand(v string) {
	o.PIMSuCommand = &v
}

// GetPIMSuCWD returns the PIMSuCWD field value if set, zero value otherwise.
func (o *SessionData) GetPIMSuCWD() string {
	if o == nil || o.PIMSuCWD == nil {
		var ret string
		return ret
	}
	return *o.PIMSuCWD
}

// GetPIMSuCWDOk returns a tuple with the PIMSuCWD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetPIMSuCWDOk() (*string, bool) {
	if o == nil || o.PIMSuCWD == nil {
		return nil, false
	}
	return o.PIMSuCWD, true
}

// HasPIMSuCWD returns a boolean if a field has been set.
func (o *SessionData) HasPIMSuCWD() bool {
	if o != nil && o.PIMSuCWD != nil {
		return true
	}

	return false
}

// SetPIMSuCWD gets a reference to the given string and assigns it to the PIMSuCWD field.
func (o *SessionData) SetPIMSuCWD(v string) {
	o.PIMSuCWD = &v
}

// GetConnectionComponentID returns the ConnectionComponentID field value if set, zero value otherwise.
func (o *SessionData) GetConnectionComponentID() string {
	if o == nil || o.ConnectionComponentID == nil {
		var ret string
		return ret
	}
	return *o.ConnectionComponentID
}

// GetConnectionComponentIDOk returns a tuple with the ConnectionComponentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetConnectionComponentIDOk() (*string, bool) {
	if o == nil || o.ConnectionComponentID == nil {
		return nil, false
	}
	return o.ConnectionComponentID, true
}

// HasConnectionComponentID returns a boolean if a field has been set.
func (o *SessionData) HasConnectionComponentID() bool {
	if o != nil && o.ConnectionComponentID != nil {
		return true
	}

	return false
}

// SetConnectionComponentID gets a reference to the given string and assigns it to the ConnectionComponentID field.
func (o *SessionData) SetConnectionComponentID(v string) {
	o.ConnectionComponentID = &v
}

// GetPSMRecordingEntity returns the PSMRecordingEntity field value if set, zero value otherwise.
func (o *SessionData) GetPSMRecordingEntity() string {
	if o == nil || o.PSMRecordingEntity == nil {
		var ret string
		return ret
	}
	return *o.PSMRecordingEntity
}

// GetPSMRecordingEntityOk returns a tuple with the PSMRecordingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetPSMRecordingEntityOk() (*string, bool) {
	if o == nil || o.PSMRecordingEntity == nil {
		return nil, false
	}
	return o.PSMRecordingEntity, true
}

// HasPSMRecordingEntity returns a boolean if a field has been set.
func (o *SessionData) HasPSMRecordingEntity() bool {
	if o != nil && o.PSMRecordingEntity != nil {
		return true
	}

	return false
}

// SetPSMRecordingEntity gets a reference to the given string and assigns it to the PSMRecordingEntity field.
func (o *SessionData) SetPSMRecordingEntity(v string) {
	o.PSMRecordingEntity = &v
}

// GetTicketID returns the TicketID field value if set, zero value otherwise.
func (o *SessionData) GetTicketID() string {
	if o == nil || o.TicketID == nil {
		var ret string
		return ret
	}
	return *o.TicketID
}

// GetTicketIDOk returns a tuple with the TicketID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetTicketIDOk() (*string, bool) {
	if o == nil || o.TicketID == nil {
		return nil, false
	}
	return o.TicketID, true
}

// HasTicketID returns a boolean if a field has been set.
func (o *SessionData) HasTicketID() bool {
	if o != nil && o.TicketID != nil {
		return true
	}

	return false
}

// SetTicketID gets a reference to the given string and assigns it to the TicketID field.
func (o *SessionData) SetTicketID(v string) {
	o.TicketID = &v
}

// GetFromIP returns the FromIP field value if set, zero value otherwise.
func (o *SessionData) GetFromIP() string {
	if o == nil || o.FromIP == nil {
		var ret string
		return ret
	}
	return *o.FromIP
}

// GetFromIPOk returns a tuple with the FromIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetFromIPOk() (*string, bool) {
	if o == nil || o.FromIP == nil {
		return nil, false
	}
	return o.FromIP, true
}

// HasFromIP returns a boolean if a field has been set.
func (o *SessionData) HasFromIP() bool {
	if o != nil && o.FromIP != nil {
		return true
	}

	return false
}

// SetFromIP gets a reference to the given string and assigns it to the FromIP field.
func (o *SessionData) SetFromIP(v string) {
	o.FromIP = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SessionData) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SessionData) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SessionData) SetProtocol(v string) {
	o.Protocol = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *SessionData) GetClient() string {
	if o == nil || o.Client == nil {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetClientOk() (*string, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *SessionData) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *SessionData) SetClient(v string) {
	o.Client = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *SessionData) GetRiskScore() int32 {
	if o == nil || o.RiskScore == nil {
		var ret int32
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetRiskScoreOk() (*int32, bool) {
	if o == nil || o.RiskScore == nil {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *SessionData) HasRiskScore() bool {
	if o != nil && o.RiskScore != nil {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given int32 and assigns it to the RiskScore field.
func (o *SessionData) SetRiskScore(v int32) {
	o.RiskScore = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SessionData) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SessionData) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *SessionData) SetSeverity(v string) {
	o.Severity = &v
}

// GetIncidentDetails returns the IncidentDetails field value if set, zero value otherwise.
func (o *SessionData) GetIncidentDetails() PTAData {
	if o == nil || o.IncidentDetails == nil {
		var ret PTAData
		return ret
	}
	return *o.IncidentDetails
}

// GetIncidentDetailsOk returns a tuple with the IncidentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetIncidentDetailsOk() (*PTAData, bool) {
	if o == nil || o.IncidentDetails == nil {
		return nil, false
	}
	return o.IncidentDetails, true
}

// HasIncidentDetails returns a boolean if a field has been set.
func (o *SessionData) HasIncidentDetails() bool {
	if o != nil && o.IncidentDetails != nil {
		return true
	}

	return false
}

// SetIncidentDetails gets a reference to the given PTAData and assigns it to the IncidentDetails field.
func (o *SessionData) SetIncidentDetails(v PTAData) {
	o.IncidentDetails = &v
}

// GetRawProperties returns the RawProperties field value if set, zero value otherwise.
func (o *SessionData) GetRawProperties() map[string]string {
	if o == nil || o.RawProperties == nil {
		var ret map[string]string
		return ret
	}
	return *o.RawProperties
}

// GetRawPropertiesOk returns a tuple with the RawProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetRawPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.RawProperties == nil {
		return nil, false
	}
	return o.RawProperties, true
}

// HasRawProperties returns a boolean if a field has been set.
func (o *SessionData) HasRawProperties() bool {
	if o != nil && o.RawProperties != nil {
		return true
	}

	return false
}

// SetRawProperties gets a reference to the given map[string]string and assigns it to the RawProperties field.
func (o *SessionData) SetRawProperties(v map[string]string) {
	o.RawProperties = &v
}

// GetRecordingFiles returns the RecordingFiles field value if set, zero value otherwise.
func (o *SessionData) GetRecordingFiles() []RecordingFile {
	if o == nil || o.RecordingFiles == nil {
		var ret []RecordingFile
		return ret
	}
	return o.RecordingFiles
}

// GetRecordingFilesOk returns a tuple with the RecordingFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetRecordingFilesOk() ([]RecordingFile, bool) {
	if o == nil || o.RecordingFiles == nil {
		return nil, false
	}
	return o.RecordingFiles, true
}

// HasRecordingFiles returns a boolean if a field has been set.
func (o *SessionData) HasRecordingFiles() bool {
	if o != nil && o.RecordingFiles != nil {
		return true
	}

	return false
}

// SetRecordingFiles gets a reference to the given []RecordingFile and assigns it to the RecordingFiles field.
func (o *SessionData) SetRecordingFiles(v []RecordingFile) {
	o.RecordingFiles = v
}

// GetRecordedActivities returns the RecordedActivities field value if set, zero value otherwise.
func (o *SessionData) GetRecordedActivities() []RecordedCommand {
	if o == nil || o.RecordedActivities == nil {
		var ret []RecordedCommand
		return ret
	}
	return o.RecordedActivities
}

// GetRecordedActivitiesOk returns a tuple with the RecordedActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetRecordedActivitiesOk() ([]RecordedCommand, bool) {
	if o == nil || o.RecordedActivities == nil {
		return nil, false
	}
	return o.RecordedActivities, true
}

// HasRecordedActivities returns a boolean if a field has been set.
func (o *SessionData) HasRecordedActivities() bool {
	if o != nil && o.RecordedActivities != nil {
		return true
	}

	return false
}

// SetRecordedActivities gets a reference to the given []RecordedCommand and assigns it to the RecordedActivities field.
func (o *SessionData) SetRecordedActivities(v []RecordedCommand) {
	o.RecordedActivities = v
}

// GetVideoSize returns the VideoSize field value if set, zero value otherwise.
func (o *SessionData) GetVideoSize() int64 {
	if o == nil || o.VideoSize == nil {
		var ret int64
		return ret
	}
	return *o.VideoSize
}

// GetVideoSizeOk returns a tuple with the VideoSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetVideoSizeOk() (*int64, bool) {
	if o == nil || o.VideoSize == nil {
		return nil, false
	}
	return o.VideoSize, true
}

// HasVideoSize returns a boolean if a field has been set.
func (o *SessionData) HasVideoSize() bool {
	if o != nil && o.VideoSize != nil {
		return true
	}

	return false
}

// SetVideoSize gets a reference to the given int64 and assigns it to the VideoSize field.
func (o *SessionData) SetVideoSize(v int64) {
	o.VideoSize = &v
}

// GetTextSize returns the TextSize field value if set, zero value otherwise.
func (o *SessionData) GetTextSize() int64 {
	if o == nil || o.TextSize == nil {
		var ret int64
		return ret
	}
	return *o.TextSize
}

// GetTextSizeOk returns a tuple with the TextSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetTextSizeOk() (*int64, bool) {
	if o == nil || o.TextSize == nil {
		return nil, false
	}
	return o.TextSize, true
}

// HasTextSize returns a boolean if a field has been set.
func (o *SessionData) HasTextSize() bool {
	if o != nil && o.TextSize != nil {
		return true
	}

	return false
}

// SetTextSize gets a reference to the given int64 and assigns it to the TextSize field.
func (o *SessionData) SetTextSize(v int64) {
	o.TextSize = &v
}

// GetDetailsUrl returns the DetailsUrl field value if set, zero value otherwise.
func (o *SessionData) GetDetailsUrl() string {
	if o == nil || o.DetailsUrl == nil {
		var ret string
		return ret
	}
	return *o.DetailsUrl
}

// GetDetailsUrlOk returns a tuple with the DetailsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionData) GetDetailsUrlOk() (*string, bool) {
	if o == nil || o.DetailsUrl == nil {
		return nil, false
	}
	return o.DetailsUrl, true
}

// HasDetailsUrl returns a boolean if a field has been set.
func (o *SessionData) HasDetailsUrl() bool {
	if o != nil && o.DetailsUrl != nil {
		return true
	}

	return false
}

// SetDetailsUrl gets a reference to the given string and assigns it to the DetailsUrl field.
func (o *SessionData) SetDetailsUrl(v string) {
	o.DetailsUrl = &v
}

func (o SessionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableSessionData struct {
	value *SessionData
	isSet bool
}

func (v NullableSessionData) Get() *SessionData {
	return v.value
}

func (v *NullableSessionData) Set(val *SessionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionData(val *SessionData) *NullableSessionData {
	return &NullableSessionData{value: val, isSet: true}
}

func (v NullableSessionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


