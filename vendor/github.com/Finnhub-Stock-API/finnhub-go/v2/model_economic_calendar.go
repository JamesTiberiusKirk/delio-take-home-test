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

// EconomicCalendar struct for EconomicCalendar
type EconomicCalendar struct {
	// Array of economic events.
	EconomicCalendar *[]EconomicEvent `json:"economicCalendar,omitempty"`
}

// NewEconomicCalendar instantiates a new EconomicCalendar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEconomicCalendar() *EconomicCalendar {
	this := EconomicCalendar{}
	return &this
}

// NewEconomicCalendarWithDefaults instantiates a new EconomicCalendar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEconomicCalendarWithDefaults() *EconomicCalendar {
	this := EconomicCalendar{}
	return &this
}

// GetEconomicCalendar returns the EconomicCalendar field value if set, zero value otherwise.
func (o *EconomicCalendar) GetEconomicCalendar() []EconomicEvent {
	if o == nil || o.EconomicCalendar == nil {
		var ret []EconomicEvent
		return ret
	}
	return *o.EconomicCalendar
}

// GetEconomicCalendarOk returns a tuple with the EconomicCalendar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicCalendar) GetEconomicCalendarOk() (*[]EconomicEvent, bool) {
	if o == nil || o.EconomicCalendar == nil {
		return nil, false
	}
	return o.EconomicCalendar, true
}

// HasEconomicCalendar returns a boolean if a field has been set.
func (o *EconomicCalendar) HasEconomicCalendar() bool {
	if o != nil && o.EconomicCalendar != nil {
		return true
	}

	return false
}

// SetEconomicCalendar gets a reference to the given []EconomicEvent and assigns it to the EconomicCalendar field.
func (o *EconomicCalendar) SetEconomicCalendar(v []EconomicEvent) {
	o.EconomicCalendar = &v
}

func (o EconomicCalendar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EconomicCalendar != nil {
		toSerialize["economicCalendar"] = o.EconomicCalendar
	}
	return json.Marshal(toSerialize)
}

type NullableEconomicCalendar struct {
	value *EconomicCalendar
	isSet bool
}

func (v NullableEconomicCalendar) Get() *EconomicCalendar {
	return v.value
}

func (v *NullableEconomicCalendar) Set(val *EconomicCalendar) {
	v.value = val
	v.isSet = true
}

func (v NullableEconomicCalendar) IsSet() bool {
	return v.isSet
}

func (v *NullableEconomicCalendar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEconomicCalendar(val *EconomicCalendar) *NullableEconomicCalendar {
	return &NullableEconomicCalendar{value: val, isSet: true}
}

func (v NullableEconomicCalendar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEconomicCalendar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


