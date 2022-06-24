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

// ListZilliqaTransactionsByBlockHeightRI struct for ListZilliqaTransactionsByBlockHeightRI
type ListZilliqaTransactionsByBlockHeightRI struct {
	Fee GetZilliqaTransactionDetailsByTransactionIDRIFee `json:"fee"`
	// Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit.
	GasLimit int32 `json:"gasLimit"`
	// Defines the price of the gas.
	GasPrice int32 `json:"gasPrice"`
	// Defines how much of the gas for the block has been used.
	GasUsed int32 `json:"gasUsed"`
	// Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash string `json:"minedInBlockHash"`
	// Represents a random value that can be adjusted to satisfy the Proof of Work.
	Nonce int32 `json:"nonce"`
	// Defines an object array of the transaction recipients.
	Recipients []ListZilliqaTransactionsByAddressRIRecipientsInner `json:"recipients"`
	// Represents an object of addresses that provide the funds.
	Senders []ListZilliqaTransactionsByAddressRISendersInner `json:"senders"`
	// Defines the exact date/time when this block was mined in Unix Timestamp.
	Timestamp int32 `json:"timestamp"`
	// Represents the hash of the transaction, which is its unique identifier.
	TransactionHash string `json:"transactionHash"`
	// Defines the numeric representation of the transaction index.
	TransactionIndex int32 `json:"transactionIndex"`
	// Defines the status of the transaction, whether it is e.g. pending or complete.
	TransactionStatus string `json:"transactionStatus"`
}

// NewListZilliqaTransactionsByBlockHeightRI instantiates a new ListZilliqaTransactionsByBlockHeightRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListZilliqaTransactionsByBlockHeightRI(fee GetZilliqaTransactionDetailsByTransactionIDRIFee, gasLimit int32, gasPrice int32, gasUsed int32, minedInBlockHash string, nonce int32, recipients []ListZilliqaTransactionsByAddressRIRecipientsInner, senders []ListZilliqaTransactionsByAddressRISendersInner, timestamp int32, transactionHash string, transactionIndex int32, transactionStatus string) *ListZilliqaTransactionsByBlockHeightRI {
	this := ListZilliqaTransactionsByBlockHeightRI{}
	this.Fee = fee
	this.GasLimit = gasLimit
	this.GasPrice = gasPrice
	this.GasUsed = gasUsed
	this.MinedInBlockHash = minedInBlockHash
	this.Nonce = nonce
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.TransactionIndex = transactionIndex
	this.TransactionStatus = transactionStatus
	return &this
}

// NewListZilliqaTransactionsByBlockHeightRIWithDefaults instantiates a new ListZilliqaTransactionsByBlockHeightRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListZilliqaTransactionsByBlockHeightRIWithDefaults() *ListZilliqaTransactionsByBlockHeightRI {
	this := ListZilliqaTransactionsByBlockHeightRI{}
	return &this
}

// GetFee returns the Fee field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetFee() GetZilliqaTransactionDetailsByTransactionIDRIFee {
	if o == nil {
		var ret GetZilliqaTransactionDetailsByTransactionIDRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetFeeOk() (*GetZilliqaTransactionDetailsByTransactionIDRIFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetFee(v GetZilliqaTransactionDetailsByTransactionIDRIFee) {
	o.Fee = v
}

// GetGasLimit returns the GasLimit field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetGasLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetGasLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetGasLimit(v int32) {
	o.GasLimit = v
}

// GetGasPrice returns the GasPrice field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetGasPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetGasPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetGasPrice(v int32) {
	o.GasPrice = v
}

// GetGasUsed returns the GasUsed field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetGasUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetGasUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetGasUsed(v int32) {
	o.GasUsed = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetMinedInBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHash, true
}

// SetMinedInBlockHash sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = v
}

// GetNonce returns the Nonce field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetNonce(v int32) {
	o.Nonce = v
}

// GetRecipients returns the Recipients field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetRecipients() []ListZilliqaTransactionsByAddressRIRecipientsInner {
	if o == nil {
		var ret []ListZilliqaTransactionsByAddressRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetRecipientsOk() ([]ListZilliqaTransactionsByAddressRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetRecipients(v []ListZilliqaTransactionsByAddressRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetSenders() []ListZilliqaTransactionsByAddressRISendersInner {
	if o == nil {
		var ret []ListZilliqaTransactionsByAddressRISendersInner
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetSendersOk() ([]ListZilliqaTransactionsByAddressRISendersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetSenders(v []ListZilliqaTransactionsByAddressRISendersInner) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionIndex returns the TransactionIndex field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTransactionIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTransactionIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIndex, true
}

// SetTransactionIndex sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetTransactionIndex(v int32) {
	o.TransactionIndex = v
}

// GetTransactionStatus returns the TransactionStatus field value
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTransactionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value
// and a boolean to check if the value has been set.
func (o *ListZilliqaTransactionsByBlockHeightRI) GetTransactionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStatus, true
}

// SetTransactionStatus sets field value
func (o *ListZilliqaTransactionsByBlockHeightRI) SetTransactionStatus(v string) {
	o.TransactionStatus = v
}

func (o ListZilliqaTransactionsByBlockHeightRI) MarshalJSON() ([]byte, error) {
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
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["minedInBlockHash"] = o.MinedInBlockHash
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
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if true {
		toSerialize["transactionIndex"] = o.TransactionIndex
	}
	if true {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableListZilliqaTransactionsByBlockHeightRI struct {
	value *ListZilliqaTransactionsByBlockHeightRI
	isSet bool
}

func (v NullableListZilliqaTransactionsByBlockHeightRI) Get() *ListZilliqaTransactionsByBlockHeightRI {
	return v.value
}

func (v *NullableListZilliqaTransactionsByBlockHeightRI) Set(val *ListZilliqaTransactionsByBlockHeightRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListZilliqaTransactionsByBlockHeightRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListZilliqaTransactionsByBlockHeightRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListZilliqaTransactionsByBlockHeightRI(val *ListZilliqaTransactionsByBlockHeightRI) *NullableListZilliqaTransactionsByBlockHeightRI {
	return &NullableListZilliqaTransactionsByBlockHeightRI{value: val, isSet: true}
}

func (v NullableListZilliqaTransactionsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListZilliqaTransactionsByBlockHeightRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


