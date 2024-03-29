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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 Zilliqa
type GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 struct {
	// Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit.
	GasLimit int32 `json:"gasLimit"`
	GasPrice GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice `json:"gasPrice"`
	// Defines how much of the gas for the block has been used.
	GasUsed int32 `json:"gasUsed"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Represents the status of this transaction.
	TransactionStatus string `json:"transactionStatus"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2(gasLimit int32, gasPrice GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice, gasUsed int32, nonce int32, transactionStatus string) *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2{}
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.Nonce = nonce
	this.TransactionStatus = transactionStatus
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2WithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2WithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2{}
	return &this
}

// GetGasLimit returns the GasLimit field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetGasLimit(v int32) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasPrice() GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetGasPrice(v GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetGasUsed(v int32) {
	o.GasUsed = v
}

// GetNonce returns the Nonce field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetNonce(v int32) {
	o.Nonce = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if true {
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


