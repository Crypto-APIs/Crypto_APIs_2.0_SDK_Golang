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

// CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST Tron Trc20 Token
type CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST struct {
	// Defines the contract address in the blockchain for an ERC20 token.
	ContractAddress string `json:"contractAddress"`
	// Defines the fee limit value.
	FeeLimit string `json:"feeLimit"`
	// Defines the Token symbol.
	Symbol string `json:"symbol"`
}

// NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST(contractAddress string, feeLimit string, symbol string) *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST {
	this := CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST{}
	this.ContractAddress = contractAddress
	this.FeeLimit = feeLimit
	this.Symbol = symbol
	return &this
}

// NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRISTWithDefaults instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRISTWithDefaults() *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST {
	this := CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetFeeLimit returns the FeeLimit field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) GetFeeLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeLimit
}

// GetFeeLimitOk returns a tuple with the FeeLimit field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) GetFeeLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeLimit, true
}

// SetFeeLimit sets field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) SetFeeLimit(v string) {
	o.FeeLimit = v
}

// GetSymbol returns the Symbol field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) SetSymbol(v string) {
	o.Symbol = v
}

func (o CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if true {
		toSerialize["feeLimit"] = o.FeeLimit
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST struct {
	value *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST
	isSet bool
}

func (v NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) Get() *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST {
	return v.value
}

func (v *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) Set(val *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST(val *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST {
	return &NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST{value: val, isSet: true}
}

func (v NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIST) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


