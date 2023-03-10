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

// SymbolChange struct for SymbolChange
type SymbolChange struct {
	// From date.
	FromDate *string `json:"fromDate,omitempty"`
	// To date.
	ToDate *string `json:"toDate,omitempty"`
	// Array of symbol change events.
	Data *[]SymbolChangeInfo `json:"data,omitempty"`
}

// NewSymbolChange instantiates a new SymbolChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolChange() *SymbolChange {
	this := SymbolChange{}
	return &this
}

// NewSymbolChangeWithDefaults instantiates a new SymbolChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolChangeWithDefaults() *SymbolChange {
	this := SymbolChange{}
	return &this
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *SymbolChange) GetFromDate() string {
	if o == nil || o.FromDate == nil {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolChange) GetFromDateOk() (*string, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *SymbolChange) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *SymbolChange) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *SymbolChange) GetToDate() string {
	if o == nil || o.ToDate == nil {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolChange) GetToDateOk() (*string, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *SymbolChange) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *SymbolChange) SetToDate(v string) {
	o.ToDate = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SymbolChange) GetData() []SymbolChangeInfo {
	if o == nil || o.Data == nil {
		var ret []SymbolChangeInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolChange) GetDataOk() (*[]SymbolChangeInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SymbolChange) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SymbolChangeInfo and assigns it to the Data field.
func (o *SymbolChange) SetData(v []SymbolChangeInfo) {
	o.Data = &v
}

func (o SymbolChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromDate != nil {
		toSerialize["fromDate"] = o.FromDate
	}
	if o.ToDate != nil {
		toSerialize["toDate"] = o.ToDate
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSymbolChange struct {
	value *SymbolChange
	isSet bool
}

func (v NullableSymbolChange) Get() *SymbolChange {
	return v.value
}

func (v *NullableSymbolChange) Set(val *SymbolChange) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolChange) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolChange(val *SymbolChange) *NullableSymbolChange {
	return &NullableSymbolChange{value: val, isSet: true}
}

func (v NullableSymbolChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


