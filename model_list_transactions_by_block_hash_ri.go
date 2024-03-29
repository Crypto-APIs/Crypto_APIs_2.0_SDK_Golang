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

// ListTransactionsByBlockHashRI struct for ListTransactionsByBlockHashRI
type ListTransactionsByBlockHashRI struct {
	// Represents the index position of the transaction in the specific block.
	Index int32 `json:"index"`
	// Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash string `json:"minedInBlockHash"`
	// Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block.
	MinedInBlockHeight int32 `json:"minedInBlockHeight"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []ListTransactionsByBlockHashRIRecipientsInner `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []ListTransactionsByBlockHashRISendersInner `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the same as `transactionId` for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols `hash` is different from `transactionId` for SegWit transactions.
	TransactionHash string `json:"transactionHash"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	Fee ListTransactionsByBlockHashRIFee `json:"fee"`
	BlockchainSpecific ListTransactionsByBlockHashRIBS `json:"blockchainSpecific"`
}

// NewListTransactionsByBlockHashRI instantiates a new ListTransactionsByBlockHashRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRI(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []ListTransactionsByBlockHashRIRecipientsInner, senders []ListTransactionsByBlockHashRISendersInner, timestamp int32, transactionHash string, transactionId string, fee ListTransactionsByBlockHashRIFee, blockchainSpecific ListTransactionsByBlockHashRIBS) *ListTransactionsByBlockHashRI {
	this := ListTransactionsByBlockHashRI{}
	this.Index = index
	this.MinedInBlockHash = minedInBlockHash
	this.MinedInBlockHeight = minedInBlockHeight
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.TransactionId = transactionId
	this.Fee = fee
	this.BlockchainSpecific = blockchainSpecific
	return &this
}

// NewListTransactionsByBlockHashRIWithDefaults instantiates a new ListTransactionsByBlockHashRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIWithDefaults() *ListTransactionsByBlockHashRI {
	this := ListTransactionsByBlockHashRI{}
	return &this
}

// GetIndex returns the Index field value
func (o *ListTransactionsByBlockHashRI) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListTransactionsByBlockHashRI) SetIndex(v int32) {
	o.Index = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value
func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHash, true
}

// SetMinedInBlockHash sets field value
func (o *ListTransactionsByBlockHashRI) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value
func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlockHeight, true
}

// SetMinedInBlockHeight sets field value
func (o *ListTransactionsByBlockHashRI) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = v
}

// GetRecipients returns the Recipients field value
func (o *ListTransactionsByBlockHashRI) GetRecipients() []ListTransactionsByBlockHashRIRecipientsInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetRecipientsOk() ([]ListTransactionsByBlockHashRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ListTransactionsByBlockHashRI) SetRecipients(v []ListTransactionsByBlockHashRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListTransactionsByBlockHashRI) GetSenders() []ListTransactionsByBlockHashRISendersInner {
	if o == nil {
		var ret []ListTransactionsByBlockHashRISendersInner
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetSendersOk() ([]ListTransactionsByBlockHashRISendersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ListTransactionsByBlockHashRI) SetSenders(v []ListTransactionsByBlockHashRISendersInner) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListTransactionsByBlockHashRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListTransactionsByBlockHashRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListTransactionsByBlockHashRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListTransactionsByBlockHashRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListTransactionsByBlockHashRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListTransactionsByBlockHashRI) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetFee returns the Fee field value
func (o *ListTransactionsByBlockHashRI) GetFee() ListTransactionsByBlockHashRIFee {
	if o == nil {
		var ret ListTransactionsByBlockHashRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetFeeOk() (*ListTransactionsByBlockHashRIFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListTransactionsByBlockHashRI) SetFee(v ListTransactionsByBlockHashRIFee) {
	o.Fee = v
}

// GetBlockchainSpecific returns the BlockchainSpecific field value
func (o *ListTransactionsByBlockHashRI) GetBlockchainSpecific() ListTransactionsByBlockHashRIBS {
	if o == nil {
		var ret ListTransactionsByBlockHashRIBS
		return ret
	}

	return o.BlockchainSpecific
}

// GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRI) GetBlockchainSpecificOk() (*ListTransactionsByBlockHashRIBS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockchainSpecific, true
}

// SetBlockchainSpecific sets field value
func (o *ListTransactionsByBlockHashRI) SetBlockchainSpecific(v ListTransactionsByBlockHashRIBS) {
	o.BlockchainSpecific = v
}

func (o ListTransactionsByBlockHashRI) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["blockchainSpecific"] = o.BlockchainSpecific
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHashRI struct {
	value *ListTransactionsByBlockHashRI
	isSet bool
}

func (v NullableListTransactionsByBlockHashRI) Get() *ListTransactionsByBlockHashRI {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRI) Set(val *ListTransactionsByBlockHashRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRI(val *ListTransactionsByBlockHashRI) *NullableListTransactionsByBlockHashRI {
	return &NullableListTransactionsByBlockHashRI{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


