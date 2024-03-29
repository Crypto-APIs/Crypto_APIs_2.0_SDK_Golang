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

// GetTransactionDetailsByTransactionIDRIBSECGasPrice struct for GetTransactionDetailsByTransactionIDRIBSECGasPrice
type GetTransactionDetailsByTransactionIDRIBSECGasPrice struct {
	// Represents the price offered to the miner to purchase this amount of gas.
	Amount string `json:"amount"`
	// Defines the unit of the gas price amount, e.g. BTC, ETH, XRP.
	Unit string `json:"unit"`
}

// NewGetTransactionDetailsByTransactionIDRIBSECGasPrice instantiates a new GetTransactionDetailsByTransactionIDRIBSECGasPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSECGasPrice(amount string, unit string) *GetTransactionDetailsByTransactionIDRIBSECGasPrice {
	this := GetTransactionDetailsByTransactionIDRIBSECGasPrice{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSECGasPriceWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSECGasPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSECGasPriceWithDefaults() *GetTransactionDetailsByTransactionIDRIBSECGasPrice {
	this := GetTransactionDetailsByTransactionIDRIBSECGasPrice{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetTransactionDetailsByTransactionIDRIBSECGasPrice) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSECGasPrice) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSECGasPrice) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetTransactionDetailsByTransactionIDRIBSECGasPrice) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSECGasPrice) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSECGasPrice) SetUnit(v string) {
	o.Unit = v
}

func (o GetTransactionDetailsByTransactionIDRIBSECGasPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice struct {
	value *GetTransactionDetailsByTransactionIDRIBSECGasPrice
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice) Get() *GetTransactionDetailsByTransactionIDRIBSECGasPrice {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice) Set(val *GetTransactionDetailsByTransactionIDRIBSECGasPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSECGasPrice(val *GetTransactionDetailsByTransactionIDRIBSECGasPrice) *NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice {
	return &NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSECGasPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


