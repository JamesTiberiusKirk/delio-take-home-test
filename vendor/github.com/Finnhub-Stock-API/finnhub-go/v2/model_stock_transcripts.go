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

// StockTranscripts struct for StockTranscripts
type StockTranscripts struct {
	// Transcript's ID used to get the <a href=\"#transcripts\">full transcript</a>.
	Id *string `json:"id,omitempty"`
	// Title.
	Title *string `json:"title,omitempty"`
	// Time of the event.
	Time *string `json:"time,omitempty"`
	// Year of earnings result in the case of earnings call transcript.
	Year *int64 `json:"year,omitempty"`
	// Quarter of earnings result in the case of earnings call transcript.
	Quarter *int64 `json:"quarter,omitempty"`
}

// NewStockTranscripts instantiates a new StockTranscripts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTranscripts() *StockTranscripts {
	this := StockTranscripts{}
	return &this
}

// NewStockTranscriptsWithDefaults instantiates a new StockTranscripts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTranscriptsWithDefaults() *StockTranscripts {
	this := StockTranscripts{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StockTranscripts) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTranscripts) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StockTranscripts) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StockTranscripts) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *StockTranscripts) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTranscripts) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *StockTranscripts) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *StockTranscripts) SetTitle(v string) {
	o.Title = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *StockTranscripts) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTranscripts) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *StockTranscripts) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *StockTranscripts) SetTime(v string) {
	o.Time = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *StockTranscripts) GetYear() int64 {
	if o == nil || o.Year == nil {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTranscripts) GetYearOk() (*int64, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *StockTranscripts) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *StockTranscripts) SetYear(v int64) {
	o.Year = &v
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *StockTranscripts) GetQuarter() int64 {
	if o == nil || o.Quarter == nil {
		var ret int64
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTranscripts) GetQuarterOk() (*int64, bool) {
	if o == nil || o.Quarter == nil {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *StockTranscripts) HasQuarter() bool {
	if o != nil && o.Quarter != nil {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given int64 and assigns it to the Quarter field.
func (o *StockTranscripts) SetQuarter(v int64) {
	o.Quarter = &v
}

func (o StockTranscripts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	if o.Quarter != nil {
		toSerialize["quarter"] = o.Quarter
	}
	return json.Marshal(toSerialize)
}

type NullableStockTranscripts struct {
	value *StockTranscripts
	isSet bool
}

func (v NullableStockTranscripts) Get() *StockTranscripts {
	return v.value
}

func (v *NullableStockTranscripts) Set(val *StockTranscripts) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTranscripts) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTranscripts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTranscripts(val *StockTranscripts) *NullableStockTranscripts {
	return &NullableStockTranscripts{value: val, isSet: true}
}

func (v NullableStockTranscripts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTranscripts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

