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

// ListOmniTransactionsByAddressResponseItemRecipients struct for ListOmniTransactionsByAddressResponseItemRecipients
type ListOmniTransactionsByAddressResponseItemRecipients struct {
	// Represents the hash of the address that receives the funds.
	Address string `json:"address"`
	// Defines the amount of the received funds as a string.
	Amount string `json:"amount"`
}

// NewListOmniTransactionsByAddressResponseItemRecipients instantiates a new ListOmniTransactionsByAddressResponseItemRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOmniTransactionsByAddressResponseItemRecipients(address string, amount string) *ListOmniTransactionsByAddressResponseItemRecipients {
	this := ListOmniTransactionsByAddressResponseItemRecipients{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListOmniTransactionsByAddressResponseItemRecipientsWithDefaults instantiates a new ListOmniTransactionsByAddressResponseItemRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOmniTransactionsByAddressResponseItemRecipientsWithDefaults() *ListOmniTransactionsByAddressResponseItemRecipients {
	this := ListOmniTransactionsByAddressResponseItemRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListOmniTransactionsByAddressResponseItemRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressResponseItemRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListOmniTransactionsByAddressResponseItemRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListOmniTransactionsByAddressResponseItemRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressResponseItemRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListOmniTransactionsByAddressResponseItemRecipients) SetAmount(v string) {
	o.Amount = v
}

func (o ListOmniTransactionsByAddressResponseItemRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListOmniTransactionsByAddressResponseItemRecipients struct {
	value *ListOmniTransactionsByAddressResponseItemRecipients
	isSet bool
}

func (v NullableListOmniTransactionsByAddressResponseItemRecipients) Get() *ListOmniTransactionsByAddressResponseItemRecipients {
	return v.value
}

func (v *NullableListOmniTransactionsByAddressResponseItemRecipients) Set(val *ListOmniTransactionsByAddressResponseItemRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByAddressResponseItemRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByAddressResponseItemRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByAddressResponseItemRecipients(val *ListOmniTransactionsByAddressResponseItemRecipients) *NullableListOmniTransactionsByAddressResponseItemRecipients {
	return &NullableListOmniTransactionsByAddressResponseItemRecipients{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByAddressResponseItemRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByAddressResponseItemRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


