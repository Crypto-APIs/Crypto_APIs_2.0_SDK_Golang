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

// GetXRPRippleBlockDetailsByBlockHeightRITotalFees Defines the total fees included in the specific block.
type GetXRPRippleBlockDetailsByBlockHeightRITotalFees struct {
	// Defines the amount of all fees included in the specific block.
	Amount string `json:"amount"`
	// Defines the unit of all fees included in the specific block.
	Unit string `json:"unit"`
}

// NewGetXRPRippleBlockDetailsByBlockHeightRITotalFees instantiates a new GetXRPRippleBlockDetailsByBlockHeightRITotalFees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleBlockDetailsByBlockHeightRITotalFees(amount string, unit string) *GetXRPRippleBlockDetailsByBlockHeightRITotalFees {
	this := GetXRPRippleBlockDetailsByBlockHeightRITotalFees{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetXRPRippleBlockDetailsByBlockHeightRITotalFeesWithDefaults instantiates a new GetXRPRippleBlockDetailsByBlockHeightRITotalFees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleBlockDetailsByBlockHeightRITotalFeesWithDefaults() *GetXRPRippleBlockDetailsByBlockHeightRITotalFees {
	this := GetXRPRippleBlockDetailsByBlockHeightRITotalFees{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) SetUnit(v string) {
	o.Unit = v
}

func (o GetXRPRippleBlockDetailsByBlockHeightRITotalFees) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees struct {
	value *GetXRPRippleBlockDetailsByBlockHeightRITotalFees
	isSet bool
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees) Get() *GetXRPRippleBlockDetailsByBlockHeightRITotalFees {
	return v.value
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees) Set(val *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees(val *GetXRPRippleBlockDetailsByBlockHeightRITotalFees) *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees {
	return &NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees{value: val, isSet: true}
}

func (v NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHeightRITotalFees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


