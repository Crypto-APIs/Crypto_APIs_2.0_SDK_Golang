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

// CreateCoinsTransactionRequestFromAddressRISenders Defines details about the source, i.e. the sender.
type CreateCoinsTransactionRequestFromAddressRISenders struct {
	// Defines the sender's public address.
	Address string `json:"address"`
}

// NewCreateCoinsTransactionRequestFromAddressRISenders instantiates a new CreateCoinsTransactionRequestFromAddressRISenders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionRequestFromAddressRISenders(address string) *CreateCoinsTransactionRequestFromAddressRISenders {
	this := CreateCoinsTransactionRequestFromAddressRISenders{}
	this.Address = address
	return &this
}

// NewCreateCoinsTransactionRequestFromAddressRISendersWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRISenders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionRequestFromAddressRISendersWithDefaults() *CreateCoinsTransactionRequestFromAddressRISenders {
	this := CreateCoinsTransactionRequestFromAddressRISenders{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateCoinsTransactionRequestFromAddressRISenders) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRISenders) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateCoinsTransactionRequestFromAddressRISenders) SetAddress(v string) {
	o.Address = v
}

func (o CreateCoinsTransactionRequestFromAddressRISenders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionRequestFromAddressRISenders struct {
	value *CreateCoinsTransactionRequestFromAddressRISenders
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromAddressRISenders) Get() *CreateCoinsTransactionRequestFromAddressRISenders {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRISenders) Set(val *CreateCoinsTransactionRequestFromAddressRISenders) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromAddressRISenders) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRISenders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromAddressRISenders(val *CreateCoinsTransactionRequestFromAddressRISenders) *NullableCreateCoinsTransactionRequestFromAddressRISenders {
	return &NullableCreateCoinsTransactionRequestFromAddressRISenders{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromAddressRISenders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRISenders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


