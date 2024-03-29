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

// GetTransactionDetailsByTransactionIDFromCallbackRIFee struct for GetTransactionDetailsByTransactionIDFromCallbackRIFee
type GetTransactionDetailsByTransactionIDFromCallbackRIFee struct {
	// When isConfirmed is True - Defines the amount of the transaction fee  When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value.
	Amount string `json:"amount"`
	// Defines the unit of the fee.
	Unit string `json:"unit"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIFee instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIFee(amount string, unit string) *GetTransactionDetailsByTransactionIDFromCallbackRIFee {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIFeeWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIFeeWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIFee {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIFee) SetUnit(v string) {
	o.Unit = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIFee
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIFee {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIFee) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIFee(val *GetTransactionDetailsByTransactionIDFromCallbackRIFee) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


