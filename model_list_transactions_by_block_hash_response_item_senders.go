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

// ListTransactionsByBlockHashResponseItemSenders struct for ListTransactionsByBlockHashResponseItemSenders
type ListTransactionsByBlockHashResponseItemSenders struct {
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Address string `json:"address"`
	// Represents the amount sent by this address.
	Amount string `json:"amount"`
}

// NewListTransactionsByBlockHashResponseItemSenders instantiates a new ListTransactionsByBlockHashResponseItemSenders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashResponseItemSenders(address string, amount string) *ListTransactionsByBlockHashResponseItemSenders {
	this := ListTransactionsByBlockHashResponseItemSenders{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListTransactionsByBlockHashResponseItemSendersWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemSenders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashResponseItemSendersWithDefaults() *ListTransactionsByBlockHashResponseItemSenders {
	this := ListTransactionsByBlockHashResponseItemSenders{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListTransactionsByBlockHashResponseItemSenders) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemSenders) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListTransactionsByBlockHashResponseItemSenders) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListTransactionsByBlockHashResponseItemSenders) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItemSenders) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListTransactionsByBlockHashResponseItemSenders) SetAmount(v string) {
	o.Amount = v
}

func (o ListTransactionsByBlockHashResponseItemSenders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHashResponseItemSenders struct {
	value *ListTransactionsByBlockHashResponseItemSenders
	isSet bool
}

func (v NullableListTransactionsByBlockHashResponseItemSenders) Get() *ListTransactionsByBlockHashResponseItemSenders {
	return v.value
}

func (v *NullableListTransactionsByBlockHashResponseItemSenders) Set(val *ListTransactionsByBlockHashResponseItemSenders) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashResponseItemSenders) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashResponseItemSenders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashResponseItemSenders(val *ListTransactionsByBlockHashResponseItemSenders) *NullableListTransactionsByBlockHashResponseItemSenders {
	return &NullableListTransactionsByBlockHashResponseItemSenders{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashResponseItemSenders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashResponseItemSenders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


