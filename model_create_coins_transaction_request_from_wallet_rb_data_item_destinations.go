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

// CreateCoinsTransactionRequestFromWalletRBDataItemDestinations struct for CreateCoinsTransactionRequestFromWalletRBDataItemDestinations
type CreateCoinsTransactionRequestFromWalletRBDataItemDestinations struct {
	// Defines the specific destination address.
	Address string `json:"address"`
	// Represents the specific amount of the transaction's destination.
	Amount string `json:"amount"`
}

// NewCreateCoinsTransactionRequestFromWalletRBDataItemDestinations instantiates a new CreateCoinsTransactionRequestFromWalletRBDataItemDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionRequestFromWalletRBDataItemDestinations(address string, amount string) *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations {
	this := CreateCoinsTransactionRequestFromWalletRBDataItemDestinations{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewCreateCoinsTransactionRequestFromWalletRBDataItemDestinationsWithDefaults instantiates a new CreateCoinsTransactionRequestFromWalletRBDataItemDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionRequestFromWalletRBDataItemDestinationsWithDefaults() *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations {
	this := CreateCoinsTransactionRequestFromWalletRBDataItemDestinations{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) SetAmount(v string) {
	o.Amount = v
}

func (o CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations struct {
	value *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations) Get() *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations) Set(val *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations(val *CreateCoinsTransactionRequestFromWalletRBDataItemDestinations) *NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations {
	return &NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromWalletRBDataItemDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

