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

// GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee struct for GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee
type GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee struct {
	// Defines the amount of the fee.
	Amount string `json:"amount"`
	// Defines the unit of the fee.
	Unit string `json:"unit"`
}

// NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee instantiates a new GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee(amount string, unit string) *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee {
	this := GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemFeeWithDefaults instantiates a new GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemFeeWithDefaults() *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee {
	this := GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) SetUnit(v string) {
	o.Unit = v
}

func (o GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee struct {
	value *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee
	isSet bool
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) Get() *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee {
	return v.value
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) Set(val *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee(val *GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) *NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee {
	return &NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee{value: val, isSet: true}
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDResponseItemFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


