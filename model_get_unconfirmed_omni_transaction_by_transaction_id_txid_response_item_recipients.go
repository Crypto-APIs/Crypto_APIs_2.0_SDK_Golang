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

// GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients struct for GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients
type GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients struct {
	// Represents the hash of the address that receives the funds.
	Address string `json:"address"`
	// Defines the amount of the received funds as a string.
	Amount string `json:"amount"`
}

// NewGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients instantiates a new GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients(address string, amount string) *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients {
	this := GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipientsWithDefaults instantiates a new GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipientsWithDefaults() *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients {
	this := GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) SetAmount(v string) {
	o.Amount = v
}

func (o GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients struct {
	value *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients
	isSet bool
}

func (v NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) Get() *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients {
	return v.value
}

func (v *NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) Set(val *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients(val *GetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) *NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients {
	return &NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients{value: val, isSet: true}
}

func (v NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUnconfirmedOmniTransactionByTransactionIDTxidResponseItemRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

