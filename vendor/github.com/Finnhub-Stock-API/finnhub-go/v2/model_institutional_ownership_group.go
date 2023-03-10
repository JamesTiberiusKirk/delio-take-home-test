/*
Finnhub API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// InstitutionalOwnershipGroup struct for InstitutionalOwnershipGroup
type InstitutionalOwnershipGroup struct {
	// Report date.
	ReportDate *string `json:"reportDate,omitempty"`
	// Array of institutional investors.
	Ownership *[]InstitutionalOwnershipInfo `json:"ownership,omitempty"`
}

// NewInstitutionalOwnershipGroup instantiates a new InstitutionalOwnershipGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionalOwnershipGroup() *InstitutionalOwnershipGroup {
	this := InstitutionalOwnershipGroup{}
	return &this
}

// NewInstitutionalOwnershipGroupWithDefaults instantiates a new InstitutionalOwnershipGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionalOwnershipGroupWithDefaults() *InstitutionalOwnershipGroup {
	this := InstitutionalOwnershipGroup{}
	return &this
}

// GetReportDate returns the ReportDate field value if set, zero value otherwise.
func (o *InstitutionalOwnershipGroup) GetReportDate() string {
	if o == nil || o.ReportDate == nil {
		var ret string
		return ret
	}
	return *o.ReportDate
}

// GetReportDateOk returns a tuple with the ReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalOwnershipGroup) GetReportDateOk() (*string, bool) {
	if o == nil || o.ReportDate == nil {
		return nil, false
	}
	return o.ReportDate, true
}

// HasReportDate returns a boolean if a field has been set.
func (o *InstitutionalOwnershipGroup) HasReportDate() bool {
	if o != nil && o.ReportDate != nil {
		return true
	}

	return false
}

// SetReportDate gets a reference to the given string and assigns it to the ReportDate field.
func (o *InstitutionalOwnershipGroup) SetReportDate(v string) {
	o.ReportDate = &v
}

// GetOwnership returns the Ownership field value if set, zero value otherwise.
func (o *InstitutionalOwnershipGroup) GetOwnership() []InstitutionalOwnershipInfo {
	if o == nil || o.Ownership == nil {
		var ret []InstitutionalOwnershipInfo
		return ret
	}
	return *o.Ownership
}

// GetOwnershipOk returns a tuple with the Ownership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalOwnershipGroup) GetOwnershipOk() (*[]InstitutionalOwnershipInfo, bool) {
	if o == nil || o.Ownership == nil {
		return nil, false
	}
	return o.Ownership, true
}

// HasOwnership returns a boolean if a field has been set.
func (o *InstitutionalOwnershipGroup) HasOwnership() bool {
	if o != nil && o.Ownership != nil {
		return true
	}

	return false
}

// SetOwnership gets a reference to the given []InstitutionalOwnershipInfo and assigns it to the Ownership field.
func (o *InstitutionalOwnershipGroup) SetOwnership(v []InstitutionalOwnershipInfo) {
	o.Ownership = &v
}

func (o InstitutionalOwnershipGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportDate != nil {
		toSerialize["reportDate"] = o.ReportDate
	}
	if o.Ownership != nil {
		toSerialize["ownership"] = o.Ownership
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionalOwnershipGroup struct {
	value *InstitutionalOwnershipGroup
	isSet bool
}

func (v NullableInstitutionalOwnershipGroup) Get() *InstitutionalOwnershipGroup {
	return v.value
}

func (v *NullableInstitutionalOwnershipGroup) Set(val *InstitutionalOwnershipGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionalOwnershipGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionalOwnershipGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionalOwnershipGroup(val *InstitutionalOwnershipGroup) *NullableInstitutionalOwnershipGroup {
	return &NullableInstitutionalOwnershipGroup{value: val, isSet: true}
}

func (v NullableInstitutionalOwnershipGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionalOwnershipGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


