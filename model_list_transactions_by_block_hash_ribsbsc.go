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

// ListTransactionsByBlockHashRIBSBSC Binance Smart Chain
type ListTransactionsByBlockHashRIBSBSC struct {
	// Represents the specific transaction contract.
	Contract *string `json:"contract,omitempty"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	GasPrice ListTransactionsByBlockHashRIBSBSCGasPrice `json:"gasPrice"`
	// Represents the exact unit of gas that was used for the transaction.
	GasUsed string `json:"gasUsed"`
	// Represents additional information that is required for the transaction.
	InputData string `json:"inputData"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Represents the status of this transaction
	TransactionStatus string `json:"transactionStatus"`
}

// NewListTransactionsByBlockHashRIBSBSC instantiates a new ListTransactionsByBlockHashRIBSBSC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSBSC(gasLimit string, gasPrice ListTransactionsByBlockHashRIBSBSCGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string) *ListTransactionsByBlockHashRIBSBSC {
	this := ListTransactionsByBlockHashRIBSBSC{}
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.InputData = inputData
	this.Nonce = nonce
	this.TransactionStatus = transactionStatus
	return &this
}

// NewListTransactionsByBlockHashRIBSBSCWithDefaults instantiates a new ListTransactionsByBlockHashRIBSBSC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSBSCWithDefaults() *ListTransactionsByBlockHashRIBSBSC {
	this := ListTransactionsByBlockHashRIBSBSC{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *ListTransactionsByBlockHashRIBSBSC) GetContract() string {
	if o == nil || o.Contract == nil {
		var ret string
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetContractOk() (*string, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given string and assigns it to the Contract field.
func (o *ListTransactionsByBlockHashRIBSBSC) SetContract(v string) {
	o.Contract = &v
}

// GetGasLimit returns the GasLimit field value
func (o *ListTransactionsByBlockHashRIBSBSC) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *ListTransactionsByBlockHashRIBSBSC) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *ListTransactionsByBlockHashRIBSBSC) GetGasPrice() ListTransactionsByBlockHashRIBSBSCGasPrice {
	if o == nil {
		var ret ListTransactionsByBlockHashRIBSBSCGasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetGasPriceOk() (*ListTransactionsByBlockHashRIBSBSCGasPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *ListTransactionsByBlockHashRIBSBSC) SetGasPrice(v ListTransactionsByBlockHashRIBSBSCGasPrice) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *ListTransactionsByBlockHashRIBSBSC) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *ListTransactionsByBlockHashRIBSBSC) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetInputData returns the InputData field value
func (o *ListTransactionsByBlockHashRIBSBSC) GetInputData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetInputDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputData, true
}

// SetInputData sets field value
func (o *ListTransactionsByBlockHashRIBSBSC) SetInputData(v string) {
	o.InputData = v
}

// GetNonce returns the Nonce field value
func (o *ListTransactionsByBlockHashRIBSBSC) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ListTransactionsByBlockHashRIBSBSC) SetNonce(v int32) {
	o.Nonce = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *ListTransactionsByBlockHashRIBSBSC) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSBSC) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *ListTransactionsByBlockHashRIBSBSC) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o ListTransactionsByBlockHashRIBSBSC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contract != nil {
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
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHashRIBSBSC struct {
	value *ListTransactionsByBlockHashRIBSBSC
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSBSC) Get() *ListTransactionsByBlockHashRIBSBSC {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSBSC) Set(val *ListTransactionsByBlockHashRIBSBSC) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSBSC) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSBSC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSBSC(val *ListTransactionsByBlockHashRIBSBSC) *NullableListTransactionsByBlockHashRIBSBSC {
	return &NullableListTransactionsByBlockHashRIBSBSC{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSBSC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSBSC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


