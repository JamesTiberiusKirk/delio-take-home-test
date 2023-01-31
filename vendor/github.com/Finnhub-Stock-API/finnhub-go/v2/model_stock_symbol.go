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

// StockSymbol struct for StockSymbol
type StockSymbol struct {
	// Symbol description
	Description *string `json:"description,omitempty"`
	// Display symbol name.
	DisplaySymbol *string `json:"displaySymbol,omitempty"`
	// Unique symbol used to identify this symbol used in <code>/stock/candle</code> endpoint.
	Symbol *string `json:"symbol,omitempty"`
	// Security type.
	Type *string `json:"type,omitempty"`
	// Primary exchange's MIC.
	Mic *string `json:"mic,omitempty"`
	// FIGI identifier.
	Figi *string `json:"figi,omitempty"`
	// Global Share Class FIGI.
	ShareClassFIGI *string `json:"shareClassFIGI,omitempty"`
	// Price's currency. This might be different from the reporting currency of fundamental data.
	Currency *string `json:"currency,omitempty"`
	// Alternative ticker for exchanges with multiple tickers for 1 stock such as BSE.
	Symbol2 *string `json:"symbol2,omitempty"`
	// ISIN. This field is only available for EU stocks and selected Asian markets. Entitlement from Finnhub is required to access this field.
	Isin *string `json:"isin,omitempty"`
}

// NewStockSymbol instantiates a new StockSymbol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockSymbol() *StockSymbol {
	this := StockSymbol{}
	return &this
}

// NewStockSymbolWithDefaults instantiates a new StockSymbol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockSymbolWithDefaults() *StockSymbol {
	this := StockSymbol{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StockSymbol) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StockSymbol) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StockSymbol) SetDescription(v string) {
	o.Description = &v
}

// GetDisplaySymbol returns the DisplaySymbol field value if set, zero value otherwise.
func (o *StockSymbol) GetDisplaySymbol() string {
	if o == nil || o.DisplaySymbol == nil {
		var ret string
		return ret
	}
	return *o.DisplaySymbol
}

// GetDisplaySymbolOk returns a tuple with the DisplaySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetDisplaySymbolOk() (*string, bool) {
	if o == nil || o.DisplaySymbol == nil {
		return nil, false
	}
	return o.DisplaySymbol, true
}

// HasDisplaySymbol returns a boolean if a field has been set.
func (o *StockSymbol) HasDisplaySymbol() bool {
	if o != nil && o.DisplaySymbol != nil {
		return true
	}

	return false
}

// SetDisplaySymbol gets a reference to the given string and assigns it to the DisplaySymbol field.
func (o *StockSymbol) SetDisplaySymbol(v string) {
	o.DisplaySymbol = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *StockSymbol) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *StockSymbol) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *StockSymbol) SetSymbol(v string) {
	o.Symbol = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StockSymbol) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StockSymbol) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StockSymbol) SetType(v string) {
	o.Type = &v
}

// GetMic returns the Mic field value if set, zero value otherwise.
func (o *StockSymbol) GetMic() string {
	if o == nil || o.Mic == nil {
		var ret string
		return ret
	}
	return *o.Mic
}

// GetMicOk returns a tuple with the Mic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetMicOk() (*string, bool) {
	if o == nil || o.Mic == nil {
		return nil, false
	}
	return o.Mic, true
}

// HasMic returns a boolean if a field has been set.
func (o *StockSymbol) HasMic() bool {
	if o != nil && o.Mic != nil {
		return true
	}

	return false
}

// SetMic gets a reference to the given string and assigns it to the Mic field.
func (o *StockSymbol) SetMic(v string) {
	o.Mic = &v
}

// GetFigi returns the Figi field value if set, zero value otherwise.
func (o *StockSymbol) GetFigi() string {
	if o == nil || o.Figi == nil {
		var ret string
		return ret
	}
	return *o.Figi
}

// GetFigiOk returns a tuple with the Figi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetFigiOk() (*string, bool) {
	if o == nil || o.Figi == nil {
		return nil, false
	}
	return o.Figi, true
}

// HasFigi returns a boolean if a field has been set.
func (o *StockSymbol) HasFigi() bool {
	if o != nil && o.Figi != nil {
		return true
	}

	return false
}

// SetFigi gets a reference to the given string and assigns it to the Figi field.
func (o *StockSymbol) SetFigi(v string) {
	o.Figi = &v
}

// GetShareClassFIGI returns the ShareClassFIGI field value if set, zero value otherwise.
func (o *StockSymbol) GetShareClassFIGI() string {
	if o == nil || o.ShareClassFIGI == nil {
		var ret string
		return ret
	}
	return *o.ShareClassFIGI
}

// GetShareClassFIGIOk returns a tuple with the ShareClassFIGI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetShareClassFIGIOk() (*string, bool) {
	if o == nil || o.ShareClassFIGI == nil {
		return nil, false
	}
	return o.ShareClassFIGI, true
}

// HasShareClassFIGI returns a boolean if a field has been set.
func (o *StockSymbol) HasShareClassFIGI() bool {
	if o != nil && o.ShareClassFIGI != nil {
		return true
	}

	return false
}

// SetShareClassFIGI gets a reference to the given string and assigns it to the ShareClassFIGI field.
func (o *StockSymbol) SetShareClassFIGI(v string) {
	o.ShareClassFIGI = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *StockSymbol) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *StockSymbol) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *StockSymbol) SetCurrency(v string) {
	o.Currency = &v
}

// GetSymbol2 returns the Symbol2 field value if set, zero value otherwise.
func (o *StockSymbol) GetSymbol2() string {
	if o == nil || o.Symbol2 == nil {
		var ret string
		return ret
	}
	return *o.Symbol2
}

// GetSymbol2Ok returns a tuple with the Symbol2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetSymbol2Ok() (*string, bool) {
	if o == nil || o.Symbol2 == nil {
		return nil, false
	}
	return o.Symbol2, true
}

// HasSymbol2 returns a boolean if a field has been set.
func (o *StockSymbol) HasSymbol2() bool {
	if o != nil && o.Symbol2 != nil {
		return true
	}

	return false
}

// SetSymbol2 gets a reference to the given string and assigns it to the Symbol2 field.
func (o *StockSymbol) SetSymbol2(v string) {
	o.Symbol2 = &v
}

// GetIsin returns the Isin field value if set, zero value otherwise.
func (o *StockSymbol) GetIsin() string {
	if o == nil || o.Isin == nil {
		var ret string
		return ret
	}
	return *o.Isin
}

// GetIsinOk returns a tuple with the Isin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockSymbol) GetIsinOk() (*string, bool) {
	if o == nil || o.Isin == nil {
		return nil, false
	}
	return o.Isin, true
}

// HasIsin returns a boolean if a field has been set.
func (o *StockSymbol) HasIsin() bool {
	if o != nil && o.Isin != nil {
		return true
	}

	return false
}

// SetIsin gets a reference to the given string and assigns it to the Isin field.
func (o *StockSymbol) SetIsin(v string) {
	o.Isin = &v
}

func (o StockSymbol) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Mic != nil {
		toSerialize["mic"] = o.Mic
	}
	if o.Figi != nil {
		toSerialize["figi"] = o.Figi
	}
	if o.ShareClassFIGI != nil {
		toSerialize["shareClassFIGI"] = o.ShareClassFIGI
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Symbol2 != nil {
		toSerialize["symbol2"] = o.Symbol2
	}
	if o.Isin != nil {
		toSerialize["isin"] = o.Isin
	}
	return json.Marshal(toSerialize)
}

type NullableStockSymbol struct {
	value *StockSymbol
	isSet bool
}

func (v NullableStockSymbol) Get() *StockSymbol {
	return v.value
}

func (v *NullableStockSymbol) Set(val *StockSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableStockSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableStockSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockSymbol(val *StockSymbol) *NullableStockSymbol {
	return &NullableStockSymbol{value: val, isSet: true}
}

func (v NullableStockSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

