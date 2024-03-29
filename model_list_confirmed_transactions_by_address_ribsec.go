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

// ListConfirmedTransactionsByAddressRIBSEC Ethereum Classic
type ListConfirmedTransactionsByAddressRIBSEC struct {
	// Represents the specific transaction contract.
	Contract string `json:"contract"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	GasPrice ListConfirmedTransactionsByAddressRIBSECGasPrice `json:"gasPrice"`
	// Represents the exact unit of gas that was used for the transaction.
	GasUsed string `json:"gasUsed"`
	// Represents additional information that is required for the transaction.
	InputData string `json:"inputData"`
	// Represents the total internal transactions count.
	InternalTransactionsCount int32 `json:"internalTransactionsCount"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Represents the total token transfers count.
	TokenTransfersCount int32 `json:"tokenTransfersCount"`
	// String representation of the transaction status
	TransactionStatus string `json:"transactionStatus"`
}

// NewListConfirmedTransactionsByAddressRIBSEC instantiates a new ListConfirmedTransactionsByAddressRIBSEC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRIBSEC(contract string, gasLimit string, gasPrice ListConfirmedTransactionsByAddressRIBSECGasPrice, gasUsed string, inputData string, internalTransactionsCount int32, nonce int32, tokenTransfersCount int32, transactionStatus string) *ListConfirmedTransactionsByAddressRIBSEC {
	this := ListConfirmedTransactionsByAddressRIBSEC{}
	this.Contract = contract
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.InputData = inputData
	this.InternalTransactionsCount = internalTransactionsCount
	this.Nonce = nonce
	this.TokenTransfersCount = tokenTransfersCount
	this.TransactionStatus = transactionStatus
	return &this
}

// NewListConfirmedTransactionsByAddressRIBSECWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSEC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIBSECWithDefaults() *ListConfirmedTransactionsByAddressRIBSEC {
	this := ListConfirmedTransactionsByAddressRIBSEC{}
	return &this
}

// GetContract returns the Contract field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetContract(v string) {
	o.Contract = v
}

// GetGasLimit returns the GasLimit field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasPrice() ListConfirmedTransactionsByAddressRIBSECGasPrice {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBSECGasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasPriceOk() (*ListConfirmedTransactionsByAddressRIBSECGasPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetGasPrice(v ListConfirmedTransactionsByAddressRIBSECGasPrice) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetInputData returns the InputData field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInputData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInputDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputData, true
}

// SetInputData sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetInputData(v string) {
	o.InputData = v
}

// GetInternalTransactionsCount returns the InternalTransactionsCount field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInternalTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalTransactionsCount
}

// GetInternalTransactionsCountOk returns a tuple with the InternalTransactionsCount field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInternalTransactionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalTransactionsCount, true
}

// SetInternalTransactionsCount sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetInternalTransactionsCount(v int32) {
	o.InternalTransactionsCount = v
}

// GetNonce returns the Nonce field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetNonce(v int32) {
	o.Nonce = v
}

// GetTokenTransfersCount returns the TokenTransfersCount field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTokenTransfersCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenTransfersCount
}

// GetTokenTransfersCountOk returns a tuple with the TokenTransfersCount field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTokenTransfersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenTransfersCount, true
}

// SetTokenTransfersCount sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetTokenTransfersCount(v int32) {
	o.TokenTransfersCount = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *ListConfirmedTransactionsByAddressRIBSEC) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o ListConfirmedTransactionsByAddressRIBSEC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contract"] = o.Contract
	}
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
		toSerialize["inputData"] = o.InputData
	}
	if true {
		toSerialize["internalTransactionsCount"] = o.InternalTransactionsCount
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["tokenTransfersCount"] = o.TokenTransfersCount
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressRIBSEC struct {
	value *ListConfirmedTransactionsByAddressRIBSEC
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRIBSEC) Get() *ListConfirmedTransactionsByAddressRIBSEC {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRIBSEC) Set(val *ListConfirmedTransactionsByAddressRIBSEC) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRIBSEC) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRIBSEC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRIBSEC(val *ListConfirmedTransactionsByAddressRIBSEC) *NullableListConfirmedTransactionsByAddressRIBSEC {
	return &NullableListConfirmedTransactionsByAddressRIBSEC{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRIBSEC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRIBSEC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


