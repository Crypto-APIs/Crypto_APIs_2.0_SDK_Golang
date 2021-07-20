/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListHDWalletXPubYPubZPubTransactionsRI struct for ListHDWalletXPubYPubZPubTransactionsRI
type ListHDWalletXPubYPubZPubTransactionsRI struct {
	// Represents the index position of the transaction in the block.
	Index int32 `json:"index"`
	// Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash string `json:"minedInBlockHash"`
	// Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block.
	MinedInBlockHeight int32 `json:"minedInBlockHeight"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []ListHDWalletXPubYPubZPubTransactionsRIRecipients `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []ListHDWalletXPubYPubZPubTransactionsRISenders `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the same as `transactionId` for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols `hash` is different from `transactionId` for SegWit transactions.
	TransactionHash string `json:"transactionHash"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	Fee ListHDWalletXPubYPubZPubTransactionsRIFee `json:"fee"`
}

// NewListHDWalletXPubYPubZPubTransactionsRI instantiates a new ListHDWalletXPubYPubZPubTransactionsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHDWalletXPubYPubZPubTransactionsRI(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []ListHDWalletXPubYPubZPubTransactionsRIRecipients, senders []ListHDWalletXPubYPubZPubTransactionsRISenders, timestamp int32, transactionHash string, transactionId string, fee ListHDWalletXPubYPubZPubTransactionsRIFee) *ListHDWalletXPubYPubZPubTransactionsRI {
	this := ListHDWalletXPubYPubZPubTransactionsRI{}
	this.Index = index
	this.MinedInBlockHash = minedInBlockHash
	this.MinedInBlockHeight = minedInBlockHeight
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.TransactionId = transactionId
	this.Fee = fee
	return &this
}

// NewListHDWalletXPubYPubZPubTransactionsRIWithDefaults instantiates a new ListHDWalletXPubYPubZPubTransactionsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHDWalletXPubYPubZPubTransactionsRIWithDefaults() *ListHDWalletXPubYPubZPubTransactionsRI {
	this := ListHDWalletXPubYPubZPubTransactionsRI{}
	return &this
}

// GetIndex returns the Index field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetIndex(v int32) {
	o.Index = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetMinedInBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinedInBlockHash, true
}

// SetMinedInBlockHash sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetMinedInBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinedInBlockHeight, true
}

// SetMinedInBlockHeight sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = v
}

// GetRecipients returns the Recipients field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetRecipients() []ListHDWalletXPubYPubZPubTransactionsRIRecipients {
	if o == nil {
		var ret []ListHDWalletXPubYPubZPubTransactionsRIRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetRecipientsOk() (*[]ListHDWalletXPubYPubZPubTransactionsRIRecipients, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetRecipients(v []ListHDWalletXPubYPubZPubTransactionsRIRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetSenders() []ListHDWalletXPubYPubZPubTransactionsRISenders {
	if o == nil {
		var ret []ListHDWalletXPubYPubZPubTransactionsRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetSendersOk() (*[]ListHDWalletXPubYPubZPubTransactionsRISenders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetSenders(v []ListHDWalletXPubYPubZPubTransactionsRISenders) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetTimestampOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetTransactionHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetFee returns the Fee field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetFee() ListHDWalletXPubYPubZPubTransactionsRIFee {
	if o == nil {
		var ret ListHDWalletXPubYPubZPubTransactionsRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListHDWalletXPubYPubZPubTransactionsRI) GetFeeOk() (*ListHDWalletXPubYPubZPubTransactionsRIFee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListHDWalletXPubYPubZPubTransactionsRI) SetFee(v ListHDWalletXPubYPubZPubTransactionsRIFee) {
	o.Fee = v
}

func (o ListHDWalletXPubYPubZPubTransactionsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["minedInBlockHash"] = o.MinedInBlockHash
	}
	if true {
		toSerialize["minedInBlockHeight"] = o.MinedInBlockHeight
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
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	return json.Marshal(toSerialize)
}

type NullableListHDWalletXPubYPubZPubTransactionsRI struct {
	value *ListHDWalletXPubYPubZPubTransactionsRI
	isSet bool
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRI) Get() *ListHDWalletXPubYPubZPubTransactionsRI {
	return v.value
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRI) Set(val *ListHDWalletXPubYPubZPubTransactionsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHDWalletXPubYPubZPubTransactionsRI(val *ListHDWalletXPubYPubZPubTransactionsRI) *NullableListHDWalletXPubYPubZPubTransactionsRI {
	return &NullableListHDWalletXPubYPubZPubTransactionsRI{value: val, isSet: true}
}

func (v NullableListHDWalletXPubYPubZPubTransactionsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHDWalletXPubYPubZPubTransactionsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

