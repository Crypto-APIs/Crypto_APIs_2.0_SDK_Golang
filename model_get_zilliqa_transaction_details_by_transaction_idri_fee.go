/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GetZilliqaTransactionDetailsByTransactionIDRIFee Represents the transaction fee.
type GetZilliqaTransactionDetailsByTransactionIDRIFee struct {
	// Represents the amount of the transaction fee.
	Amount string `json:"amount"`
	// Represents the unit of the transaction fee.
	Unit string `json:"unit"`
}

// NewGetZilliqaTransactionDetailsByTransactionIDRIFee instantiates a new GetZilliqaTransactionDetailsByTransactionIDRIFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZilliqaTransactionDetailsByTransactionIDRIFee(amount string, unit string) *GetZilliqaTransactionDetailsByTransactionIDRIFee {
	this := GetZilliqaTransactionDetailsByTransactionIDRIFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewGetZilliqaTransactionDetailsByTransactionIDRIFeeWithDefaults instantiates a new GetZilliqaTransactionDetailsByTransactionIDRIFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZilliqaTransactionDetailsByTransactionIDRIFeeWithDefaults() *GetZilliqaTransactionDetailsByTransactionIDRIFee {
	this := GetZilliqaTransactionDetailsByTransactionIDRIFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRIFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaTransactionDetailsByTransactionIDRIFee) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRIFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRIFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaTransactionDetailsByTransactionIDRIFee) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRIFee) SetUnit(v string) {
	o.Unit = v
}

func (o GetZilliqaTransactionDetailsByTransactionIDRIFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableGetZilliqaTransactionDetailsByTransactionIDRIFee struct {
	value *GetZilliqaTransactionDetailsByTransactionIDRIFee
	isSet bool
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRIFee) Get() *GetZilliqaTransactionDetailsByTransactionIDRIFee {
	return v.value
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRIFee) Set(val *GetZilliqaTransactionDetailsByTransactionIDRIFee) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRIFee) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRIFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZilliqaTransactionDetailsByTransactionIDRIFee(val *GetZilliqaTransactionDetailsByTransactionIDRIFee) *NullableGetZilliqaTransactionDetailsByTransactionIDRIFee {
	return &NullableGetZilliqaTransactionDetailsByTransactionIDRIFee{value: val, isSet: true}
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRIFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRIFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


