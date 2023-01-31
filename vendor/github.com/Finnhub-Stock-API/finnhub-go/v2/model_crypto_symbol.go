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

// CryptoSymbol struct for CryptoSymbol
type CryptoSymbol struct {
	// Symbol description
	Description *string `json:"description,omitempty"`
	// Display symbol name.
	DisplaySymbol *string `json:"displaySymbol,omitempty"`
	// Unique symbol used to identify this symbol used in <code>/crypto/candle</code> endpoint.
	Symbol *string `json:"symbol,omitempty"`
}

// NewCryptoSymbol instantiates a new CryptoSymbol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoSymbol() *CryptoSymbol {
	this := CryptoSymbol{}
	return &this
}

// NewCryptoSymbolWithDefaults instantiates a new CryptoSymbol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoSymbolWithDefaults() *CryptoSymbol {
	this := CryptoSymbol{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CryptoSymbol) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoSymbol) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CryptoSymbol) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CryptoSymbol) SetDescription(v string) {
	o.Description = &v
}

// GetDisplaySymbol returns the DisplaySymbol field value if set, zero value otherwise.
func (o *CryptoSymbol) GetDisplaySymbol() string {
	if o == nil || o.DisplaySymbol == nil {
		var ret string
		return ret
	}
	return *o.DisplaySymbol
}

// GetDisplaySymbolOk returns a tuple with the DisplaySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoSymbol) GetDisplaySymbolOk() (*string, bool) {
	if o == nil || o.DisplaySymbol == nil {
		return nil, false
	}
	return o.DisplaySymbol, true
}

// HasDisplaySymbol returns a boolean if a field has been set.
func (o *CryptoSymbol) HasDisplaySymbol() bool {
	if o != nil && o.DisplaySymbol != nil {
		return true
	}

	return false
}

// SetDisplaySymbol gets a reference to the given string and assigns it to the DisplaySymbol field.
func (o *CryptoSymbol) SetDisplaySymbol(v string) {
	o.DisplaySymbol = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CryptoSymbol) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoSymbol) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CryptoSymbol) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CryptoSymbol) SetSymbol(v string) {
	o.Symbol = &v
}

func (o CryptoSymbol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplaySymbol != nil {
		toSerialize["displaySymbol"] = o.DisplaySymbol
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableCryptoSymbol struct {
	value *CryptoSymbol
	isSet bool
}

func (v NullableCryptoSymbol) Get() *CryptoSymbol {
	return v.value
}

func (v *NullableCryptoSymbol) Set(val *CryptoSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoSymbol(val *CryptoSymbol) *NullableCryptoSymbol {
	return &NullableCryptoSymbol{value: val, isSet: true}
}

func (v NullableCryptoSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


