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

// CreateCoinsTransactionRequestFromAddressRIRecipients struct for CreateCoinsTransactionRequestFromAddressRIRecipients
type CreateCoinsTransactionRequestFromAddressRIRecipients struct {
	// Defines the destination address.
	Address string `json:"address"`
	// Defines the amount sent to the destination address.
	Amount string `json:"amount"`
}

// NewCreateCoinsTransactionRequestFromAddressRIRecipients instantiates a new CreateCoinsTransactionRequestFromAddressRIRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionRequestFromAddressRIRecipients(address string, amount string) *CreateCoinsTransactionRequestFromAddressRIRecipients {
	this := CreateCoinsTransactionRequestFromAddressRIRecipients{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewCreateCoinsTransactionRequestFromAddressRIRecipientsWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRIRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionRequestFromAddressRIRecipientsWithDefaults() *CreateCoinsTransactionRequestFromAddressRIRecipients {
	this := CreateCoinsTransactionRequestFromAddressRIRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateCoinsTransactionRequestFromAddressRIRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRIRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateCoinsTransactionRequestFromAddressRIRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *CreateCoinsTransactionRequestFromAddressRIRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRIRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateCoinsTransactionRequestFromAddressRIRecipients) SetAmount(v string) {
	o.Amount = v
}

func (o CreateCoinsTransactionRequestFromAddressRIRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionRequestFromAddressRIRecipients struct {
	value *CreateCoinsTransactionRequestFromAddressRIRecipients
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromAddressRIRecipients) Get() *CreateCoinsTransactionRequestFromAddressRIRecipients {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRIRecipients) Set(val *CreateCoinsTransactionRequestFromAddressRIRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromAddressRIRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRIRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromAddressRIRecipients(val *CreateCoinsTransactionRequestFromAddressRIRecipients) *NullableCreateCoinsTransactionRequestFromAddressRIRecipients {
	return &NullableCreateCoinsTransactionRequestFromAddressRIRecipients{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromAddressRIRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRIRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

