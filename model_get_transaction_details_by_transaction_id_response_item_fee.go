/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GetTransactionDetailsByTransactionIDResponseItemFee struct for GetTransactionDetailsByTransactionIDResponseItemFee
type GetTransactionDetailsByTransactionIDResponseItemFee struct {
	// Defines the amount of the transaction fee.
	Amount string `json:"amount"`
	// Defines the unit of the fee amount, e.g. BTC, ETH, XRP.
	Unit string `json:"unit"`
}

// NewGetTransactionDetailsByTransactionIDResponseItemFee instantiates a new GetTransactionDetailsByTransactionIDResponseItemFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDResponseItemFee(amount string, unit string) *GetTransactionDetailsByTransactionIDResponseItemFee {
	this := GetTransactionDetailsByTransactionIDResponseItemFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetTransactionDetailsByTransactionIDResponseItemFeeWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDResponseItemFeeWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemFee {
	this := GetTransactionDetailsByTransactionIDResponseItemFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetTransactionDetailsByTransactionIDResponseItemFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemFee) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetTransactionDetailsByTransactionIDResponseItemFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDResponseItemFee) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetTransactionDetailsByTransactionIDResponseItemFee) SetUnit(v string) {
	o.Unit = v
}

func (o GetTransactionDetailsByTransactionIDResponseItemFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDResponseItemFee struct {
	value *GetTransactionDetailsByTransactionIDResponseItemFee
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDResponseItemFee) Get() *GetTransactionDetailsByTransactionIDResponseItemFee {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseItemFee) Set(val *GetTransactionDetailsByTransactionIDResponseItemFee) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDResponseItemFee) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseItemFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDResponseItemFee(val *GetTransactionDetailsByTransactionIDResponseItemFee) *NullableGetTransactionDetailsByTransactionIDResponseItemFee {
	return &NullableGetTransactionDetailsByTransactionIDResponseItemFee{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDResponseItemFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDResponseItemFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

