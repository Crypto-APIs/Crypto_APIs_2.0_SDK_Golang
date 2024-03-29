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

// GetTransactionDetailsByTransactionIDRIBSBSC Binance Smart Chain
type GetTransactionDetailsByTransactionIDRIBSBSC struct {
	// Represents the specific transaction contract
	Contract string `json:"contract"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	GasPrice GetTransactionDetailsByTransactionIDRIBSBSCGasPrice `json:"gasPrice"`
	// Defines the unit of the gas price amount, e.g. BTC, ETH, XRP.
	GasUsed string `json:"gasUsed"`
	// Represents additional information that is required for the transaction.
	InputData string `json:"inputData"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
}

// NewGetTransactionDetailsByTransactionIDRIBSBSC instantiates a new GetTransactionDetailsByTransactionIDRIBSBSC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDRIBSBSC(contract string, gasLimit string, gasPrice GetTransactionDetailsByTransactionIDRIBSBSCGasPrice, gasUsed string, inputData string, nonce int32) *GetTransactionDetailsByTransactionIDRIBSBSC {
	this := GetTransactionDetailsByTransactionIDRIBSBSC{}
	this.Contract = contract
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.InputData = inputData
	this.Nonce = nonce
	return &this
}

// NewGetTransactionDetailsByTransactionIDRIBSBSCWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBSC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDRIBSBSCWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBSC {
	this := GetTransactionDetailsByTransactionIDRIBSBSC{}
	return &this
}

// GetContract returns the Contract field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetContract(v string) {
	o.Contract = v
}

// GetGasLimit returns the GasLimit field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasPrice() GetTransactionDetailsByTransactionIDRIBSBSCGasPrice {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDRIBSBSCGasPrice
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDRIBSBSCGasPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetGasPrice(v GetTransactionDetailsByTransactionIDRIBSBSCGasPrice) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetInputData returns the InputData field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetInputData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetInputDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputData, true
}

// SetInputData sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetInputData(v string) {
	o.InputData = v
}

// GetNonce returns the Nonce field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetNonce(v int32) {
	o.Nonce = v
}

func (o GetTransactionDetailsByTransactionIDRIBSBSC) MarshalJSON() ([]byte, error) {
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
		toSerialize["nonce"] = o.Nonce
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDRIBSBSC struct {
	value *GetTransactionDetailsByTransactionIDRIBSBSC
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBSC) Get() *GetTransactionDetailsByTransactionIDRIBSBSC {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBSC) Set(val *GetTransactionDetailsByTransactionIDRIBSBSC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBSC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBSC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDRIBSBSC(val *GetTransactionDetailsByTransactionIDRIBSBSC) *NullableGetTransactionDetailsByTransactionIDRIBSBSC {
	return &NullableGetTransactionDetailsByTransactionIDRIBSBSC{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDRIBSBSC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDRIBSBSC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


