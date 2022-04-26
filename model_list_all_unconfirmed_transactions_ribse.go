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

// ListAllUnconfirmedTransactionsRIBSE Ethereum
type ListAllUnconfirmedTransactionsRIBSE struct {
	Fee ListAllUnconfirmedTransactionsRIBSEFee `json:"fee"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	GasPrice ListAllUnconfirmedTransactionsRIBSEGasPrice `json:"gasPrice"`
	// Represents additional information that is required for the transaction.
	InputData string `json:"inputData"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Defines the transaction status.
	TransactionStatus string `json:"transactionStatus"`
}

// NewListAllUnconfirmedTransactionsRIBSE instantiates a new ListAllUnconfirmedTransactionsRIBSE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllUnconfirmedTransactionsRIBSE(fee ListAllUnconfirmedTransactionsRIBSEFee, gasLimit string, gasPrice ListAllUnconfirmedTransactionsRIBSEGasPrice, inputData string, nonce int32, transactionStatus string) *ListAllUnconfirmedTransactionsRIBSE {
	this := ListAllUnconfirmedTransactionsRIBSE{}
	this.Fee = fee
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.InputData = inputData
	this.Nonce = nonce
	this.TransactionStatus = transactionStatus
	return &this
}

// NewListAllUnconfirmedTransactionsRIBSEWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllUnconfirmedTransactionsRIBSEWithDefaults() *ListAllUnconfirmedTransactionsRIBSE {
	this := ListAllUnconfirmedTransactionsRIBSE{}
	return &this
}

// GetFee returns the Fee field value
func (o *ListAllUnconfirmedTransactionsRIBSE) GetFee() ListAllUnconfirmedTransactionsRIBSEFee {
	if o == nil {
		var ret ListAllUnconfirmedTransactionsRIBSEFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSE) GetFeeOk() (*ListAllUnconfirmedTransactionsRIBSEFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListAllUnconfirmedTransactionsRIBSE) SetFee(v ListAllUnconfirmedTransactionsRIBSEFee) {
	o.Fee = v
}

// GetGasLimit returns the GasLimit field value
func (o *ListAllUnconfirmedTransactionsRIBSE) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSE) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *ListAllUnconfirmedTransactionsRIBSE) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *ListAllUnconfirmedTransactionsRIBSE) GetGasPrice() ListAllUnconfirmedTransactionsRIBSEGasPrice {
	if o == nil {
		var ret ListAllUnconfirmedTransactionsRIBSEGasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSE) GetGasPriceOk() (*ListAllUnconfirmedTransactionsRIBSEGasPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *ListAllUnconfirmedTransactionsRIBSE) SetGasPrice(v ListAllUnconfirmedTransactionsRIBSEGasPrice) {
	o.GasPrice = v
}

// GetInputData returns the InputData field value
func (o *ListAllUnconfirmedTransactionsRIBSE) GetInputData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSE) GetInputDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputData, true
}

// SetInputData sets field value
func (o *ListAllUnconfirmedTransactionsRIBSE) SetInputData(v string) {
	o.InputData = v
}

// GetNonce returns the Nonce field value
func (o *ListAllUnconfirmedTransactionsRIBSE) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSE) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ListAllUnconfirmedTransactionsRIBSE) SetNonce(v int32) {
	o.Nonce = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *ListAllUnconfirmedTransactionsRIBSE) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRIBSE) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *ListAllUnconfirmedTransactionsRIBSE) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o ListAllUnconfirmedTransactionsRIBSE) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if true {
		toSerialize["inputData"] = o.InputData
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableListAllUnconfirmedTransactionsRIBSE struct {
	value *ListAllUnconfirmedTransactionsRIBSE
	isSet bool
}

func (v NullableListAllUnconfirmedTransactionsRIBSE) Get() *ListAllUnconfirmedTransactionsRIBSE {
	return v.value
}

func (v *NullableListAllUnconfirmedTransactionsRIBSE) Set(val *ListAllUnconfirmedTransactionsRIBSE) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllUnconfirmedTransactionsRIBSE) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllUnconfirmedTransactionsRIBSE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllUnconfirmedTransactionsRIBSE(val *ListAllUnconfirmedTransactionsRIBSE) *NullableListAllUnconfirmedTransactionsRIBSE {
	return &NullableListAllUnconfirmedTransactionsRIBSE{value: val, isSet: true}
}

func (v NullableListAllUnconfirmedTransactionsRIBSE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllUnconfirmedTransactionsRIBSE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


