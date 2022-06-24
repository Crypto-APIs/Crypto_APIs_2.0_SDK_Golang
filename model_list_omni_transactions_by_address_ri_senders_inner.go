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

// ListOmniTransactionsByAddressRISendersInner struct for ListOmniTransactionsByAddressRISendersInner
type ListOmniTransactionsByAddressRISendersInner struct {
	// Represents the hash of the address that provides the funds.
	Address string `json:"address"`
	// Represents the total amount sent by this address including the fee.
	Amount string `json:"amount"`
}

// NewListOmniTransactionsByAddressRISendersInner instantiates a new ListOmniTransactionsByAddressRISendersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOmniTransactionsByAddressRISendersInner(address string, amount string) *ListOmniTransactionsByAddressRISendersInner {
	this := ListOmniTransactionsByAddressRISendersInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListOmniTransactionsByAddressRISendersInnerWithDefaults instantiates a new ListOmniTransactionsByAddressRISendersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOmniTransactionsByAddressRISendersInnerWithDefaults() *ListOmniTransactionsByAddressRISendersInner {
	this := ListOmniTransactionsByAddressRISendersInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListOmniTransactionsByAddressRISendersInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRISendersInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListOmniTransactionsByAddressRISendersInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListOmniTransactionsByAddressRISendersInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRISendersInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListOmniTransactionsByAddressRISendersInner) SetAmount(v string) {
	o.Amount = v
}

func (o ListOmniTransactionsByAddressRISendersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListOmniTransactionsByAddressRISendersInner struct {
	value *ListOmniTransactionsByAddressRISendersInner
	isSet bool
}

func (v NullableListOmniTransactionsByAddressRISendersInner) Get() *ListOmniTransactionsByAddressRISendersInner {
	return v.value
}

func (v *NullableListOmniTransactionsByAddressRISendersInner) Set(val *ListOmniTransactionsByAddressRISendersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByAddressRISendersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByAddressRISendersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByAddressRISendersInner(val *ListOmniTransactionsByAddressRISendersInner) *NullableListOmniTransactionsByAddressRISendersInner {
	return &NullableListOmniTransactionsByAddressRISendersInner{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByAddressRISendersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByAddressRISendersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


