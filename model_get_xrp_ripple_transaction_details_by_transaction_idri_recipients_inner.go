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

// GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner struct for GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner
type GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner struct {
	// Represents the hash of the address that receives the funds.
	Address string `json:"address"`
	// Defines the amount of the received funds as a string.
	Amount string `json:"amount"`
}

// NewGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner(address string, amount string) *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner {
	this := GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInnerWithDefaults instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInnerWithDefaults() *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner {
	this := GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) SetAmount(v string) {
	o.Amount = v
}

func (o GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner struct {
	value *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner
	isSet bool
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) Get() *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner {
	return v.value
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) Set(val *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner(val *GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) *NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner {
	return &NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner{value: val, isSet: true}
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


