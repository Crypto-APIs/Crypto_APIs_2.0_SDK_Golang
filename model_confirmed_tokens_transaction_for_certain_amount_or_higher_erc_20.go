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

// ConfirmedTokensTransactionForCertainAmountOrHigherErc20 ERC-20
type ConfirmedTokensTransactionForCertainAmountOrHigherErc20 struct {
	// Specifies the name of the token.
	Name string `json:"name"`
	// Specifies an identifier of the token, where up to five alphanumeric characters can be used for it.
	Symbol string `json:"symbol"`
	// Defines how many decimals can be used to break the token.
	Decimals *string `json:"decimals,omitempty"`
	// Defines the amount of tokens sent with the confirmed transaction.
	Amount string `json:"amount"`
	// Defines the address of the contract.
	ContractAddress string `json:"contractAddress"`
}

// NewConfirmedTokensTransactionForCertainAmountOrHigherErc20 instantiates a new ConfirmedTokensTransactionForCertainAmountOrHigherErc20 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmedTokensTransactionForCertainAmountOrHigherErc20(name string, symbol string, amount string, contractAddress string) *ConfirmedTokensTransactionForCertainAmountOrHigherErc20 {
	this := ConfirmedTokensTransactionForCertainAmountOrHigherErc20{}
	this.Name = name
	this.Symbol = symbol
	this.Amount = amount
	this.ContractAddress = contractAddress
	return &this
}

// NewConfirmedTokensTransactionForCertainAmountOrHigherErc20WithDefaults instantiates a new ConfirmedTokensTransactionForCertainAmountOrHigherErc20 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmedTokensTransactionForCertainAmountOrHigherErc20WithDefaults() *ConfirmedTokensTransactionForCertainAmountOrHigherErc20 {
	this := ConfirmedTokensTransactionForCertainAmountOrHigherErc20{}
	return &this
}

// GetName returns the Name field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) SetSymbol(v string) {
	o.Symbol = v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetDecimals() string {
	if o == nil || o.Decimals == nil {
		var ret string
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetDecimalsOk() (*string, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given string and assigns it to the Decimals field.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) SetDecimals(v string) {
	o.Decimals = &v
}

// GetAmount returns the Amount field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) SetAmount(v string) {
	o.Amount = v
}

// GetContractAddress returns the ContractAddress field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) SetContractAddress(v string) {
	o.ContractAddress = v
}

func (o ConfirmedTokensTransactionForCertainAmountOrHigherErc20) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	return json.Marshal(toSerialize)
}

type NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20 struct {
	value *ConfirmedTokensTransactionForCertainAmountOrHigherErc20
	isSet bool
}

func (v NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20) Get() *ConfirmedTokensTransactionForCertainAmountOrHigherErc20 {
	return v.value
}

func (v *NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20) Set(val *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmedTokensTransactionForCertainAmountOrHigherErc20(val *ConfirmedTokensTransactionForCertainAmountOrHigherErc20) *NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20 {
	return &NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20{value: val, isSet: true}
}

func (v NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmedTokensTransactionForCertainAmountOrHigherErc20) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


