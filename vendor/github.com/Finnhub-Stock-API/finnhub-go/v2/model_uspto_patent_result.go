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

// UsptoPatentResult struct for UsptoPatentResult
type UsptoPatentResult struct {
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Array of patents.
	Data *[]UsptoPatent `json:"data,omitempty"`
}

// NewUsptoPatentResult instantiates a new UsptoPatentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsptoPatentResult() *UsptoPatentResult {
	this := UsptoPatentResult{}
	return &this
}

// NewUsptoPatentResultWithDefaults instantiates a new UsptoPatentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsptoPatentResultWithDefaults() *UsptoPatentResult {
	this := UsptoPatentResult{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UsptoPatentResult) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatentResult) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UsptoPatentResult) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UsptoPatentResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsptoPatentResult) GetData() []UsptoPatent {
	if o == nil || o.Data == nil {
		var ret []UsptoPatent
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatentResult) GetDataOk() (*[]UsptoPatent, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsptoPatentResult) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []UsptoPatent and assigns it to the Data field.
func (o *UsptoPatentResult) SetData(v []UsptoPatent) {
	o.Data = &v
}

func (o UsptoPatentResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUsptoPatentResult struct {
	value *UsptoPatentResult
	isSet bool
}

func (v NullableUsptoPatentResult) Get() *UsptoPatentResult {
	return v.value
}

func (v *NullableUsptoPatentResult) Set(val *UsptoPatentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUsptoPatentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUsptoPatentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsptoPatentResult(val *UsptoPatentResult) *NullableUsptoPatentResult {
	return &NullableUsptoPatentResult{value: val, isSet: true}
}

func (v NullableUsptoPatentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsptoPatentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

