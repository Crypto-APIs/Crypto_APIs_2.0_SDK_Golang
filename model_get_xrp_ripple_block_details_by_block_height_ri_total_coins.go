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

// GetXRPRippleBlockDetailsByBlockHeightRITotalCoins Represents the total Coins of the block as a string.
type GetXRPRippleBlockDetailsByBlockHeightRITotalCoins struct {
	// Represents the total amount of all Coins of the block as a string.
	Amount string `json:"amount"`
	// Represents the unit of total Coins of the block as a string.
	Unit string `json:"unit"`
}

// NewGetXRPRippleBlockDetailsByBlockHeightRITotalCoins instantiates a new GetXRPRippleBlockDetailsByBlockHeightRITotalCoins object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleBlockDetailsByBlockHeightRITotalCoins(amount string, unit string) *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins {
	this := GetXRPRippleBlockDetailsByBlockHeightRITotalCoins{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetXRPRippleBlockDetailsByBlockHeightRITotalCoinsWithDefaults instantiates a new GetXRPRippleBlockDetailsByBlockHeightRITotalCoins object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleBlockDetailsByBlockHeightRITotalCoinsWithDefaults() *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins {
	this := GetXRPRippleBlockDetailsByBlockHeightRITotalCoins{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) SetUnit(v string) {
	o.Unit = v
}

func (o GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins struct {
	value *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins
	isSet bool
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins) Get() *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins {
	return v.value
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins) Set(val *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins(val *GetXRPRippleBlockDetailsByBlockHeightRITotalCoins) *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins {
	return &NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins{value: val, isSet: true}
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalCoins) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


