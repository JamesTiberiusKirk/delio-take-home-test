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

// ForexCandles struct for ForexCandles
type ForexCandles struct {
	// List of open prices for returned candles.
	O *[]float32 `json:"o,omitempty"`
	// List of high prices for returned candles.
	H *[]float32 `json:"h,omitempty"`
	// List of low prices for returned candles.
	L *[]float32 `json:"l,omitempty"`
	// List of close prices for returned candles.
	C *[]float32 `json:"c,omitempty"`
	// List of volume data for returned candles.
	V *[]float32 `json:"v,omitempty"`
	// List of timestamp for returned candles.
	T *[]float32 `json:"t,omitempty"`
	// Status of the response. This field can either be ok or no_data.
	S *string `json:"s,omitempty"`
}

// NewForexCandles instantiates a new ForexCandles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForexCandles() *ForexCandles {
	this := ForexCandles{}
	return &this
}

// NewForexCandlesWithDefaults instantiates a new ForexCandles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForexCandlesWithDefaults() *ForexCandles {
	this := ForexCandles{}
	return &this
}

// GetO returns the O field value if set, zero value otherwise.
func (o *ForexCandles) GetO() []float32 {
	if o == nil || o.O == nil {
		var ret []float32
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetOOk() (*[]float32, bool) {
	if o == nil || o.O == nil {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *ForexCandles) HasO() bool {
	if o != nil && o.O != nil {
		return true
	}

	return false
}

// SetO gets a reference to the given []float32 and assigns it to the O field.
func (o *ForexCandles) SetO(v []float32) {
	o.O = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *ForexCandles) GetH() []float32 {
	if o == nil || o.H == nil {
		var ret []float32
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetHOk() (*[]float32, bool) {
	if o == nil || o.H == nil {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *ForexCandles) HasH() bool {
	if o != nil && o.H != nil {
		return true
	}

	return false
}

// SetH gets a reference to the given []float32 and assigns it to the H field.
func (o *ForexCandles) SetH(v []float32) {
	o.H = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *ForexCandles) GetL() []float32 {
	if o == nil || o.L == nil {
		var ret []float32
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetLOk() (*[]float32, bool) {
	if o == nil || o.L == nil {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *ForexCandles) HasL() bool {
	if o != nil && o.L != nil {
		return true
	}

	return false
}

// SetL gets a reference to the given []float32 and assigns it to the L field.
func (o *ForexCandles) SetL(v []float32) {
	o.L = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ForexCandles) GetC() []float32 {
	if o == nil || o.C == nil {
		var ret []float32
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetCOk() (*[]float32, bool) {
	if o == nil || o.C == nil {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *ForexCandles) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given []float32 and assigns it to the C field.
func (o *ForexCandles) SetC(v []float32) {
	o.C = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *ForexCandles) GetV() []float32 {
	if o == nil || o.V == nil {
		var ret []float32
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetVOk() (*[]float32, bool) {
	if o == nil || o.V == nil {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *ForexCandles) HasV() bool {
	if o != nil && o.V != nil {
		return true
	}

	return false
}

// SetV gets a reference to the given []float32 and assigns it to the V field.
func (o *ForexCandles) SetV(v []float32) {
	o.V = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ForexCandles) GetT() []float32 {
	if o == nil || o.T == nil {
		var ret []float32
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetTOk() (*[]float32, bool) {
	if o == nil || o.T == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ForexCandles) HasT() bool {
	if o != nil && o.T != nil {
		return true
	}

	return false
}

// SetT gets a reference to the given []float32 and assigns it to the T field.
func (o *ForexCandles) SetT(v []float32) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ForexCandles) GetS() string {
	if o == nil || o.S == nil {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForexCandles) GetSOk() (*string, bool) {
	if o == nil || o.S == nil {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *ForexCandles) HasS() bool {
	if o != nil && o.S != nil {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ForexCandles) SetS(v string) {
	o.S = &v
}

func (o ForexCandles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.O != nil {
		toSerialize["o"] = o.O
	}
	if o.H != nil {
		toSerialize["h"] = o.H
	}
	if o.L != nil {
		toSerialize["l"] = o.L
	}
	if o.C != nil {
		toSerialize["c"] = o.C
	}
	if o.V != nil {
		toSerialize["v"] = o.V
	}
	if o.T != nil {
		toSerialize["t"] = o.T
	}
	if o.S != nil {
		toSerialize["s"] = o.S
	}
	return json.Marshal(toSerialize)
}

type NullableForexCandles struct {
	value *ForexCandles
	isSet bool
}

func (v NullableForexCandles) Get() *ForexCandles {
	return v.value
}

func (v *NullableForexCandles) Set(val *ForexCandles) {
	v.value = val
	v.isSet = true
}

func (v NullableForexCandles) IsSet() bool {
	return v.isSet
}

func (v *NullableForexCandles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForexCandles(val *ForexCandles) *NullableForexCandles {
	return &NullableForexCandles{value: val, isSet: true}
}

func (v NullableForexCandles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForexCandles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

