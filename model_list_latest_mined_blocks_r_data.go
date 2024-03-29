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

// ListLatestMinedBlocksRData struct for ListLatestMinedBlocksRData
type ListLatestMinedBlocksRData struct {
	// Defines how many items should be returned in the response per page basis.
	Limit int32 `json:"limit"`
	// The starting index of the response items, i.e. where the response should start listing the returned items.
	Offset int32 `json:"offset"`
	// Defines the total number of items returned in the response.
	Total int32 `json:"total"`
	Items []ListLatestMinedBlocksRI `json:"items"`
}

// NewListLatestMinedBlocksRData instantiates a new ListLatestMinedBlocksRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLatestMinedBlocksRData(limit int32, offset int32, total int32, items []ListLatestMinedBlocksRI) *ListLatestMinedBlocksRData {
	this := ListLatestMinedBlocksRData{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Items = items
	return &this
}

// NewListLatestMinedBlocksRDataWithDefaults instantiates a new ListLatestMinedBlocksRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLatestMinedBlocksRDataWithDefaults() *ListLatestMinedBlocksRData {
	this := ListLatestMinedBlocksRData{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListLatestMinedBlocksRData) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRData) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListLatestMinedBlocksRData) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ListLatestMinedBlocksRData) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRData) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListLatestMinedBlocksRData) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *ListLatestMinedBlocksRData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListLatestMinedBlocksRData) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ListLatestMinedBlocksRData) GetItems() []ListLatestMinedBlocksRI {
	if o == nil {
		var ret []ListLatestMinedBlocksRI
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRData) GetItemsOk() ([]ListLatestMinedBlocksRI, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListLatestMinedBlocksRData) SetItems(v []ListLatestMinedBlocksRI) {
	o.Items = v
}

func (o ListLatestMinedBlocksRData) MarshalJSON() ([]byte, error) {
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

type NullableListLatestMinedBlocksRData struct {
	value *ListLatestMinedBlocksRData
	isSet bool
}

func (v NullableListLatestMinedBlocksRData) Get() *ListLatestMinedBlocksRData {
	return v.value
}

func (v *NullableListLatestMinedBlocksRData) Set(val *ListLatestMinedBlocksRData) {
	v.value = val
	v.isSet = true
}

func (v NullableListLatestMinedBlocksRData) IsSet() bool {
	return v.isSet
}

func (v *NullableListLatestMinedBlocksRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLatestMinedBlocksRData(val *ListLatestMinedBlocksRData) *NullableListLatestMinedBlocksRData {
	return &NullableListLatestMinedBlocksRData{value: val, isSet: true}
}

func (v NullableListLatestMinedBlocksRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLatestMinedBlocksRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


