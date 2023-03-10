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

// IsinChangeInfo struct for IsinChangeInfo
type IsinChangeInfo struct {
	// Event's date.
	AtDate *string `json:"atDate,omitempty"`
	// Old ISIN.
	OldIsin *string `json:"oldIsin,omitempty"`
	// New ISIN.
	NewIsin *string `json:"newIsin,omitempty"`
}

// NewIsinChangeInfo instantiates a new IsinChangeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsinChangeInfo() *IsinChangeInfo {
	this := IsinChangeInfo{}
	return &this
}

// NewIsinChangeInfoWithDefaults instantiates a new IsinChangeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsinChangeInfoWithDefaults() *IsinChangeInfo {
	this := IsinChangeInfo{}
	return &this
}

// GetAtDate returns the AtDate field value if set, zero value otherwise.
func (o *IsinChangeInfo) GetAtDate() string {
	if o == nil || o.AtDate == nil {
		var ret string
		return ret
	}
	return *o.AtDate
}

// GetAtDateOk returns a tuple with the AtDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsinChangeInfo) GetAtDateOk() (*string, bool) {
	if o == nil || o.AtDate == nil {
		return nil, false
	}
	return o.AtDate, true
}

// HasAtDate returns a boolean if a field has been set.
func (o *IsinChangeInfo) HasAtDate() bool {
	if o != nil && o.AtDate != nil {
		return true
	}

	return false
}

// SetAtDate gets a reference to the given string and assigns it to the AtDate field.
func (o *IsinChangeInfo) SetAtDate(v string) {
	o.AtDate = &v
}

// GetOldIsin returns the OldIsin field value if set, zero value otherwise.
func (o *IsinChangeInfo) GetOldIsin() string {
	if o == nil || o.OldIsin == nil {
		var ret string
		return ret
	}
	return *o.OldIsin
}

// GetOldIsinOk returns a tuple with the OldIsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsinChangeInfo) GetOldIsinOk() (*string, bool) {
	if o == nil || o.OldIsin == nil {
		return nil, false
	}
	return o.OldIsin, true
}

// HasOldIsin returns a boolean if a field has been set.
func (o *IsinChangeInfo) HasOldIsin() bool {
	if o != nil && o.OldIsin != nil {
		return true
	}

	return false
}

// SetOldIsin gets a reference to the given string and assigns it to the OldIsin field.
func (o *IsinChangeInfo) SetOldIsin(v string) {
	o.OldIsin = &v
}

// GetNewIsin returns the NewIsin field value if set, zero value otherwise.
func (o *IsinChangeInfo) GetNewIsin() string {
	if o == nil || o.NewIsin == nil {
		var ret string
		return ret
	}
	return *o.NewIsin
}

// GetNewIsinOk returns a tuple with the NewIsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsinChangeInfo) GetNewIsinOk() (*string, bool) {
	if o == nil || o.NewIsin == nil {
		return nil, false
	}
	return o.NewIsin, true
}

// HasNewIsin returns a boolean if a field has been set.
func (o *IsinChangeInfo) HasNewIsin() bool {
	if o != nil && o.NewIsin != nil {
		return true
	}

	return false
}

// SetNewIsin gets a reference to the given string and assigns it to the NewIsin field.
func (o *IsinChangeInfo) SetNewIsin(v string) {
	o.NewIsin = &v
}

func (o IsinChangeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AtDate != nil {
		toSerialize["atDate"] = o.AtDate
	}
	if o.OldIsin != nil {
		toSerialize["oldIsin"] = o.OldIsin
	}
	if o.NewIsin != nil {
		toSerialize["newIsin"] = o.NewIsin
	}
	return json.Marshal(toSerialize)
}

type NullableIsinChangeInfo struct {
	value *IsinChangeInfo
	isSet bool
}

func (v NullableIsinChangeInfo) Get() *IsinChangeInfo {
	return v.value
}

func (v *NullableIsinChangeInfo) Set(val *IsinChangeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIsinChangeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIsinChangeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsinChangeInfo(val *IsinChangeInfo) *NullableIsinChangeInfo {
	return &NullableIsinChangeInfo{value: val, isSet: true}
}

func (v NullableIsinChangeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsinChangeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


