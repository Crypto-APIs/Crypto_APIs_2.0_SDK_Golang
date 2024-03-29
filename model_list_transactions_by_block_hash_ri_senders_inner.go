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

// ListTransactionsByBlockHashRISendersInner struct for ListTransactionsByBlockHashRISendersInner
type ListTransactionsByBlockHashRISendersInner struct {
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Address string `json:"address"`
	// Represents the total amount sent by this address including the fee.
	Amount string `json:"amount"`
}

// NewListTransactionsByBlockHashRISendersInner instantiates a new ListTransactionsByBlockHashRISendersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRISendersInner(address string, amount string) *ListTransactionsByBlockHashRISendersInner {
	this := ListTransactionsByBlockHashRISendersInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListTransactionsByBlockHashRISendersInnerWithDefaults instantiates a new ListTransactionsByBlockHashRISendersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRISendersInnerWithDefaults() *ListTransactionsByBlockHashRISendersInner {
	this := ListTransactionsByBlockHashRISendersInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListTransactionsByBlockHashRISendersInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRISendersInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListTransactionsByBlockHashRISendersInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListTransactionsByBlockHashRISendersInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRISendersInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListTransactionsByBlockHashRISendersInner) SetAmount(v string) {
	o.Amount = v
}

func (o ListTransactionsByBlockHashRISendersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHashRISendersInner struct {
	value *ListTransactionsByBlockHashRISendersInner
	isSet bool
}

func (v NullableListTransactionsByBlockHashRISendersInner) Get() *ListTransactionsByBlockHashRISendersInner {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRISendersInner) Set(val *ListTransactionsByBlockHashRISendersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRISendersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRISendersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRISendersInner(val *ListTransactionsByBlockHashRISendersInner) *NullableListTransactionsByBlockHashRISendersInner {
	return &NullableListTransactionsByBlockHashRISendersInner{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRISendersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRISendersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


