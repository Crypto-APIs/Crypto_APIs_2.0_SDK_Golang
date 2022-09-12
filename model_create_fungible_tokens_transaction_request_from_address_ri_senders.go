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

// CreateFungibleTokensTransactionRequestFromAddressRISenders Defines details about the source, i.e. the sender.
type CreateFungibleTokensTransactionRequestFromAddressRISenders struct {
	// Defines the sender's public address.
	Address string `json:"address"`
}

// NewCreateFungibleTokensTransactionRequestFromAddressRISenders instantiates a new CreateFungibleTokensTransactionRequestFromAddressRISenders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFungibleTokensTransactionRequestFromAddressRISenders(address string) *CreateFungibleTokensTransactionRequestFromAddressRISenders {
	this := CreateFungibleTokensTransactionRequestFromAddressRISenders{}
	this.Address = address
	return &this
}

// NewCreateFungibleTokensTransactionRequestFromAddressRISendersWithDefaults instantiates a new CreateFungibleTokensTransactionRequestFromAddressRISenders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFungibleTokensTransactionRequestFromAddressRISendersWithDefaults() *CreateFungibleTokensTransactionRequestFromAddressRISenders {
	this := CreateFungibleTokensTransactionRequestFromAddressRISenders{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRISenders) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRISenders) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRISenders) SetAddress(v string) {
	o.Address = v
}

func (o CreateFungibleTokensTransactionRequestFromAddressRISenders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFungibleTokensTransactionRequestFromAddressRISenders struct {
	value *CreateFungibleTokensTransactionRequestFromAddressRISenders
	isSet bool
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRISenders) Get() *CreateFungibleTokensTransactionRequestFromAddressRISenders {
	return v.value
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRISenders) Set(val *CreateFungibleTokensTransactionRequestFromAddressRISenders) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRISenders) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRISenders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokensTransactionRequestFromAddressRISenders(val *CreateFungibleTokensTransactionRequestFromAddressRISenders) *NullableCreateFungibleTokensTransactionRequestFromAddressRISenders {
	return &NullableCreateFungibleTokensTransactionRequestFromAddressRISenders{value: val, isSet: true}
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRISenders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRISenders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


