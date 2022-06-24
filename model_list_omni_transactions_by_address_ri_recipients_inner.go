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

// ListOmniTransactionsByAddressRIRecipientsInner struct for ListOmniTransactionsByAddressRIRecipientsInner
type ListOmniTransactionsByAddressRIRecipientsInner struct {
	// Represents the hash of the address that receives the funds.
	Address string `json:"address"`
	// Defines the amount of the received funds as a string.
	Amount string `json:"amount"`
}

// NewListOmniTransactionsByAddressRIRecipientsInner instantiates a new ListOmniTransactionsByAddressRIRecipientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOmniTransactionsByAddressRIRecipientsInner(address string, amount string) *ListOmniTransactionsByAddressRIRecipientsInner {
	this := ListOmniTransactionsByAddressRIRecipientsInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListOmniTransactionsByAddressRIRecipientsInnerWithDefaults instantiates a new ListOmniTransactionsByAddressRIRecipientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOmniTransactionsByAddressRIRecipientsInnerWithDefaults() *ListOmniTransactionsByAddressRIRecipientsInner {
	this := ListOmniTransactionsByAddressRIRecipientsInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListOmniTransactionsByAddressRIRecipientsInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRIRecipientsInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListOmniTransactionsByAddressRIRecipientsInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListOmniTransactionsByAddressRIRecipientsInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRIRecipientsInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListOmniTransactionsByAddressRIRecipientsInner) SetAmount(v string) {
	o.Amount = v
}

func (o ListOmniTransactionsByAddressRIRecipientsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListOmniTransactionsByAddressRIRecipientsInner struct {
	value *ListOmniTransactionsByAddressRIRecipientsInner
	isSet bool
}

func (v NullableListOmniTransactionsByAddressRIRecipientsInner) Get() *ListOmniTransactionsByAddressRIRecipientsInner {
	return v.value
}

func (v *NullableListOmniTransactionsByAddressRIRecipientsInner) Set(val *ListOmniTransactionsByAddressRIRecipientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByAddressRIRecipientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByAddressRIRecipientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByAddressRIRecipientsInner(val *ListOmniTransactionsByAddressRIRecipientsInner) *NullableListOmniTransactionsByAddressRIRecipientsInner {
	return &NullableListOmniTransactionsByAddressRIRecipientsInner{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByAddressRIRecipientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByAddressRIRecipientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


