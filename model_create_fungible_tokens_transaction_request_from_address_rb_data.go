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

// CreateFungibleTokensTransactionRequestFromAddressRBData struct for CreateFungibleTokensTransactionRequestFromAddressRBData
type CreateFungibleTokensTransactionRequestFromAddressRBData struct {
	Item CreateFungibleTokensTransactionRequestFromAddressRBDataItem `json:"item"`
}

// NewCreateFungibleTokensTransactionRequestFromAddressRBData instantiates a new CreateFungibleTokensTransactionRequestFromAddressRBData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFungibleTokensTransactionRequestFromAddressRBData(item CreateFungibleTokensTransactionRequestFromAddressRBDataItem) *CreateFungibleTokensTransactionRequestFromAddressRBData {
	this := CreateFungibleTokensTransactionRequestFromAddressRBData{}
	this.Item = item
	return &this
}

// NewCreateFungibleTokensTransactionRequestFromAddressRBDataWithDefaults instantiates a new CreateFungibleTokensTransactionRequestFromAddressRBData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFungibleTokensTransactionRequestFromAddressRBDataWithDefaults() *CreateFungibleTokensTransactionRequestFromAddressRBData {
	this := CreateFungibleTokensTransactionRequestFromAddressRBData{}
	return &this
}

// GetItem returns the Item field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRBData) GetItem() CreateFungibleTokensTransactionRequestFromAddressRBDataItem {
	if o == nil {
		var ret CreateFungibleTokensTransactionRequestFromAddressRBDataItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRBData) GetItemOk() (*CreateFungibleTokensTransactionRequestFromAddressRBDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRBData) SetItem(v CreateFungibleTokensTransactionRequestFromAddressRBDataItem) {
	o.Item = v
}

func (o CreateFungibleTokensTransactionRequestFromAddressRBData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFungibleTokensTransactionRequestFromAddressRBData struct {
	value *CreateFungibleTokensTransactionRequestFromAddressRBData
	isSet bool
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRBData) Get() *CreateFungibleTokensTransactionRequestFromAddressRBData {
	return v.value
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRBData) Set(val *CreateFungibleTokensTransactionRequestFromAddressRBData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRBData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRBData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokensTransactionRequestFromAddressRBData(val *CreateFungibleTokensTransactionRequestFromAddressRBData) *NullableCreateFungibleTokensTransactionRequestFromAddressRBData {
	return &NullableCreateFungibleTokensTransactionRequestFromAddressRBData{value: val, isSet: true}
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRBData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRBData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


