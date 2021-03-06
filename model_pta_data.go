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

// PTAData struct for PTAData
type PTAData struct {
	Incident *Incident `json:"Incident,omitempty"`
	Sessions []Session `json:"Sessions,omitempty"`
}

// NewPTAData instantiates a new PTAData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPTAData() *PTAData {
	this := PTAData{}
	return &this
}

// NewPTADataWithDefaults instantiates a new PTAData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPTADataWithDefaults() *PTAData {
	this := PTAData{}
	return &this
}

// GetIncident returns the Incident field value if set, zero value otherwise.
func (o *PTAData) GetIncident() Incident {
	if o == nil || o.Incident == nil {
		var ret Incident
		return ret
	}
	return *o.Incident
}

// GetIncidentOk returns a tuple with the Incident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTAData) GetIncidentOk() (*Incident, bool) {
	if o == nil || o.Incident == nil {
		return nil, false
	}
	return o.Incident, true
}

// HasIncident returns a boolean if a field has been set.
func (o *PTAData) HasIncident() bool {
	if o != nil && o.Incident != nil {
		return true
	}

	return false
}

// SetIncident gets a reference to the given Incident and assigns it to the Incident field.
func (o *PTAData) SetIncident(v Incident) {
	o.Incident = &v
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *PTAData) GetSessions() []Session {
	if o == nil || o.Sessions == nil {
		var ret []Session
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTAData) GetSessionsOk() ([]Session, bool) {
	if o == nil || o.Sessions == nil {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *PTAData) HasSessions() bool {
	if o != nil && o.Sessions != nil {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []Session and assigns it to the Sessions field.
func (o *PTAData) SetSessions(v []Session) {
	o.Sessions = v
}

func (o PTAData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Incident != nil {
		toSerialize["Incident"] = o.Incident
	}
	if o.Sessions != nil {
		toSerialize["Sessions"] = o.Sessions
	}
	return json.Marshal(toSerialize)
}

type NullablePTAData struct {
	value *PTAData
	isSet bool
}

func (v NullablePTAData) Get() *PTAData {
	return v.value
}

func (v *NullablePTAData) Set(val *PTAData) {
	v.value = val
	v.isSet = true
}

func (v NullablePTAData) IsSet() bool {
	return v.isSet
}

func (v *NullablePTAData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePTAData(val *PTAData) *NullablePTAData {
	return &NullablePTAData{value: val, isSet: true}
}

func (v NullablePTAData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePTAData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


