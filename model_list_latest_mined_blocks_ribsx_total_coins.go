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

// ListLatestMinedBlocksRIBSXTotalCoins struct for ListLatestMinedBlocksRIBSXTotalCoins
type ListLatestMinedBlocksRIBSXTotalCoins struct {
	// Defines the amount of all coins.
	Amount *string `json:"amount,omitempty"`
	// Defines the unit of the amount of all coins.
	Unit *string `json:"unit,omitempty"`
}

// NewListLatestMinedBlocksRIBSXTotalCoins instantiates a new ListLatestMinedBlocksRIBSXTotalCoins object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLatestMinedBlocksRIBSXTotalCoins() *ListLatestMinedBlocksRIBSXTotalCoins {
	this := ListLatestMinedBlocksRIBSXTotalCoins{}
	return &this
}

// NewListLatestMinedBlocksRIBSXTotalCoinsWithDefaults instantiates a new ListLatestMinedBlocksRIBSXTotalCoins object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLatestMinedBlocksRIBSXTotalCoinsWithDefaults() *ListLatestMinedBlocksRIBSXTotalCoins {
	this := ListLatestMinedBlocksRIBSXTotalCoins{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) SetAmount(v string) {
	o.Amount = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ListLatestMinedBlocksRIBSXTotalCoins) SetUnit(v string) {
	o.Unit = &v
}

func (o ListLatestMinedBlocksRIBSXTotalCoins) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListLatestMinedBlocksRIBSXTotalCoins struct {
	value *ListLatestMinedBlocksRIBSXTotalCoins
	isSet bool
}

func (v NullableListLatestMinedBlocksRIBSXTotalCoins) Get() *ListLatestMinedBlocksRIBSXTotalCoins {
	return v.value
}

func (v *NullableListLatestMinedBlocksRIBSXTotalCoins) Set(val *ListLatestMinedBlocksRIBSXTotalCoins) {
	v.value = val
	v.isSet = true
}

func (v NullableListLatestMinedBlocksRIBSXTotalCoins) IsSet() bool {
	return v.isSet
}

func (v *NullableListLatestMinedBlocksRIBSXTotalCoins) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLatestMinedBlocksRIBSXTotalCoins(val *ListLatestMinedBlocksRIBSXTotalCoins) *NullableListLatestMinedBlocksRIBSXTotalCoins {
	return &NullableListLatestMinedBlocksRIBSXTotalCoins{value: val, isSet: true}
}

func (v NullableListLatestMinedBlocksRIBSXTotalCoins) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLatestMinedBlocksRIBSXTotalCoins) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

