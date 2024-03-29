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

// GetZilliqaTransactionDetailsByTransactionIDRISendersInner struct for GetZilliqaTransactionDetailsByTransactionIDRISendersInner
type GetZilliqaTransactionDetailsByTransactionIDRISendersInner struct {
	// Represents the hash of the address that provides the funds.
	Address string `json:"address"`
	// Represents the total amount sent by this address including the fee.
	Amount string `json:"amount"`
}

// NewGetZilliqaTransactionDetailsByTransactionIDRISendersInner instantiates a new GetZilliqaTransactionDetailsByTransactionIDRISendersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZilliqaTransactionDetailsByTransactionIDRISendersInner(address string, amount string) *GetZilliqaTransactionDetailsByTransactionIDRISendersInner {
	this := GetZilliqaTransactionDetailsByTransactionIDRISendersInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewGetZilliqaTransactionDetailsByTransactionIDRISendersInnerWithDefaults instantiates a new GetZilliqaTransactionDetailsByTransactionIDRISendersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZilliqaTransactionDetailsByTransactionIDRISendersInnerWithDefaults() *GetZilliqaTransactionDetailsByTransactionIDRISendersInner {
	this := GetZilliqaTransactionDetailsByTransactionIDRISendersInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) SetAmount(v string) {
	o.Amount = v
}

func (o GetZilliqaTransactionDetailsByTransactionIDRISendersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner struct {
	value *GetZilliqaTransactionDetailsByTransactionIDRISendersInner
	isSet bool
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner) Get() *GetZilliqaTransactionDetailsByTransactionIDRISendersInner {
	return v.value
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner) Set(val *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner(val *GetZilliqaTransactionDetailsByTransactionIDRISendersInner) *NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner {
	return &NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner{value: val, isSet: true}
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRISendersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


