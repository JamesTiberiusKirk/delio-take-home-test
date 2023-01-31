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

// Dividends struct for Dividends
type Dividends struct {
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Ex-Dividend date.
	Date *string `json:"date,omitempty"`
	// Amount in local currency.
	Amount *float32 `json:"amount,omitempty"`
	// Adjusted dividend.
	AdjustedAmount *float32 `json:"adjustedAmount,omitempty"`
	// Pay date.
	PayDate *string `json:"payDate,omitempty"`
	// Record date.
	RecordDate *string `json:"recordDate,omitempty"`
	// Declaration date.
	DeclarationDate *string `json:"declarationDate,omitempty"`
	// Currency.
	Currency *string `json:"currency,omitempty"`
}

// NewDividends instantiates a new Dividends object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDividends() *Dividends {
	this := Dividends{}
	return &this
}

// NewDividendsWithDefaults instantiates a new Dividends object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDividendsWithDefaults() *Dividends {
	this := Dividends{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Dividends) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Dividends) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Dividends) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Dividends) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Dividends) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Dividends) SetDate(v string) {
	o.Date = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Dividends) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Dividends) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Dividends) SetAmount(v float32) {
	o.Amount = &v
}

// GetAdjustedAmount returns the AdjustedAmount field value if set, zero value otherwise.
func (o *Dividends) GetAdjustedAmount() float32 {
	if o == nil || o.AdjustedAmount == nil {
		var ret float32
		return ret
	}
	return *o.AdjustedAmount
}

// GetAdjustedAmountOk returns a tuple with the AdjustedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetAdjustedAmountOk() (*float32, bool) {
	if o == nil || o.AdjustedAmount == nil {
		return nil, false
	}
	return o.AdjustedAmount, true
}

// HasAdjustedAmount returns a boolean if a field has been set.
func (o *Dividends) HasAdjustedAmount() bool {
	if o != nil && o.AdjustedAmount != nil {
		return true
	}

	return false
}

// SetAdjustedAmount gets a reference to the given float32 and assigns it to the AdjustedAmount field.
func (o *Dividends) SetAdjustedAmount(v float32) {
	o.AdjustedAmount = &v
}

// GetPayDate returns the PayDate field value if set, zero value otherwise.
func (o *Dividends) GetPayDate() string {
	if o == nil || o.PayDate == nil {
		var ret string
		return ret
	}
	return *o.PayDate
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetPayDateOk() (*string, bool) {
	if o == nil || o.PayDate == nil {
		return nil, false
	}
	return o.PayDate, true
}

// HasPayDate returns a boolean if a field has been set.
func (o *Dividends) HasPayDate() bool {
	if o != nil && o.PayDate != nil {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given string and assigns it to the PayDate field.
func (o *Dividends) SetPayDate(v string) {
	o.PayDate = &v
}

// GetRecordDate returns the RecordDate field value if set, zero value otherwise.
func (o *Dividends) GetRecordDate() string {
	if o == nil || o.RecordDate == nil {
		var ret string
		return ret
	}
	return *o.RecordDate
}

// GetRecordDateOk returns a tuple with the RecordDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetRecordDateOk() (*string, bool) {
	if o == nil || o.RecordDate == nil {
		return nil, false
	}
	return o.RecordDate, true
}

// HasRecordDate returns a boolean if a field has been set.
func (o *Dividends) HasRecordDate() bool {
	if o != nil && o.RecordDate != nil {
		return true
	}

	return false
}

// SetRecordDate gets a reference to the given string and assigns it to the RecordDate field.
func (o *Dividends) SetRecordDate(v string) {
	o.RecordDate = &v
}

// GetDeclarationDate returns the DeclarationDate field value if set, zero value otherwise.
func (o *Dividends) GetDeclarationDate() string {
	if o == nil || o.DeclarationDate == nil {
		var ret string
		return ret
	}
	return *o.DeclarationDate
}

// GetDeclarationDateOk returns a tuple with the DeclarationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetDeclarationDateOk() (*string, bool) {
	if o == nil || o.DeclarationDate == nil {
		return nil, false
	}
	return o.DeclarationDate, true
}

// HasDeclarationDate returns a boolean if a field has been set.
func (o *Dividends) HasDeclarationDate() bool {
	if o != nil && o.DeclarationDate != nil {
		return true
	}

	return false
}

// SetDeclarationDate gets a reference to the given string and assigns it to the DeclarationDate field.
func (o *Dividends) SetDeclarationDate(v string) {
	o.DeclarationDate = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Dividends) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dividends) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Dividends) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Dividends) SetCurrency(v string) {
	o.Currency = &v
}

func (o Dividends) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.AdjustedAmount != nil {
		toSerialize["adjustedAmount"] = o.AdjustedAmount
	}
	if o.PayDate != nil {
		toSerialize["payDate"] = o.PayDate
	}
	if o.RecordDate != nil {
		toSerialize["recordDate"] = o.RecordDate
	}
	if o.DeclarationDate != nil {
		toSerialize["declarationDate"] = o.DeclarationDate
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableDividends struct {
	value *Dividends
	isSet bool
}

func (v NullableDividends) Get() *Dividends {
	return v.value
}

func (v *NullableDividends) Set(val *Dividends) {
	v.value = val
	v.isSet = true
}

func (v NullableDividends) IsSet() bool {
	return v.isSet
}

func (v *NullableDividends) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDividends(val *Dividends) *NullableDividends {
	return &NullableDividends{value: val, isSet: true}
}

func (v NullableDividends) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDividends) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


