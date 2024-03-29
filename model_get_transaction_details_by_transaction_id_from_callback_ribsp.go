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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSP Polygon
type GetTransactionDetailsByTransactionIDFromCallbackRIBSP struct {
	// Representation of the amount value.
	Amount string `json:"amount"`
	// Represents the specific transaction contract.
	Contract string `json:"contract"`
	// Represents the price offered to the miner to purchase this amount of gas.
	Gas string `json:"gas"`
	// Represents the price offered to the miner to purchase this amount of gas.
	GasPrice string `json:"gasPrice"`
	// Represents the exact unit of gas that was used for the transaction.
	GasUsed string `json:"gasUsed"`
	// Represents additional information that is required for the transaction.
	Input string `json:"input"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients string `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders string `json:"senders"`
	// Represents the status of this transaction.
	TransactionStatus string `json:"transactionStatus"`
	// Represents the unique identifier of a transaction, i.e. it could be transactionId in UTXO-based protocols like Bitcoin, and transaction hash in Ethereum blockchain.
	Txid string `json:"txid"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSP instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSP(amount string, contract string, gas string, gasPrice string, gasUsed string, input string, nonce int32, recipients string, senders string, transactionStatus string, txid string) *GetTransactionDetailsByTransactionIDFromCallbackRIBSP {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSP{}
	this.Amount = amount
	this.Contract = contract
	this.Gas = gas
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.Input = input
	this.Nonce = nonce
	this.Recipients = recipients
	this.Senders = senders
	this.TransactionStatus = transactionStatus
	this.Txid = txid
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSPWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSPWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSP {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSP{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetAmount(v string) {
	o.Amount = v
}

// GetContract returns the Contract field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetContract() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetContract(v string) {
	o.Contract = v
}

// GetGas returns the Gas field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gas
}

// GetGasOk returns a tuple with the Gas field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gas, true
}

// SetGas sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetGas(v string) {
	o.Gas = v
}

// GetGasPrice returns the GasPrice field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetGasPrice(v string) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetInput returns the Input field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetInput(v string) {
	o.Input = v
}

// GetNonce returns the Nonce field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetNonce(v int32) {
	o.Nonce = v
}

// GetRecipients returns the Recipients field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetRecipients() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetRecipientsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetRecipients(v string) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetSenders() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetSendersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetSenders(v string) {
	o.Senders = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

// GetTxid returns the Txid field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetTxid(v string) {
	o.Txid = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["contract"] = o.Contract
	}
	if true {
		toSerialize["gas"] = o.Gas
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if true {
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["input"] = o.Input
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	if true {
		toSerialize["txid"] = o.Txid
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSP
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSP {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


