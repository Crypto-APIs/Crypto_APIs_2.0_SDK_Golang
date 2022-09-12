/*
CryptoAPIs

Crypto APIs is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2021-03-20
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListXRPRippleTransactionsByAddressAndTimeRangeRData struct for ListXRPRippleTransactionsByAddressAndTimeRangeRData
type ListXRPRippleTransactionsByAddressAndTimeRangeRData struct {
	// Defines how many items should be returned in the response per page basis.
	Limit int32 `json:"limit"`
	// The starting index of the response items, i.e. where the response should start listing the returned items.
	Offset int32 `json:"offset"`
	// Defines the total number of items returned in the response.
	Total int32 `json:"total"`
	Items []ListXRPRippleTransactionsByAddressAndTimeRangeRI `json:"items"`
}

// NewListXRPRippleTransactionsByAddressAndTimeRangeRData instantiates a new ListXRPRippleTransactionsByAddressAndTimeRangeRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByAddressAndTimeRangeRData(limit int32, offset int32, total int32, items []ListXRPRippleTransactionsByAddressAndTimeRangeRI) *ListXRPRippleTransactionsByAddressAndTimeRangeRData {
	this := ListXRPRippleTransactionsByAddressAndTimeRangeRData{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Items = items
	return &this
}

// NewListXRPRippleTransactionsByAddressAndTimeRangeRDataWithDefaults instantiates a new ListXRPRippleTransactionsByAddressAndTimeRangeRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByAddressAndTimeRangeRDataWithDefaults() *ListXRPRippleTransactionsByAddressAndTimeRangeRData {
	this := ListXRPRippleTransactionsByAddressAndTimeRangeRData{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetItems() []ListXRPRippleTransactionsByAddressAndTimeRangeRI {
	if o == nil {
		var ret []ListXRPRippleTransactionsByAddressAndTimeRangeRI
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) GetItemsOk() ([]ListXRPRippleTransactionsByAddressAndTimeRangeRI, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRData) SetItems(v []ListXRPRippleTransactionsByAddressAndTimeRangeRI) {
	o.Items = v
}

func (o ListXRPRippleTransactionsByAddressAndTimeRangeRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByAddressAndTimeRangeRData struct {
	value *ListXRPRippleTransactionsByAddressAndTimeRangeRData
	isSet bool
}

func (v NullableListXRPRippleTransactionsByAddressAndTimeRangeRData) Get() *ListXRPRippleTransactionsByAddressAndTimeRangeRData {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByAddressAndTimeRangeRData) Set(val *ListXRPRippleTransactionsByAddressAndTimeRangeRData) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByAddressAndTimeRangeRData) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByAddressAndTimeRangeRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByAddressAndTimeRangeRData(val *ListXRPRippleTransactionsByAddressAndTimeRangeRData) *NullableListXRPRippleTransactionsByAddressAndTimeRangeRData {
	return &NullableListXRPRippleTransactionsByAddressAndTimeRangeRData{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByAddressAndTimeRangeRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByAddressAndTimeRangeRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


