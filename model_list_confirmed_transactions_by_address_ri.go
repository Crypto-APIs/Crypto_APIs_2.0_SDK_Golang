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

// ListConfirmedTransactionsByAddressRI struct for ListConfirmedTransactionsByAddressRI
type ListConfirmedTransactionsByAddressRI struct {
	// Represents the index position of the transaction in the block.
	Index int32 `json:"index"`
	// Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash *string `json:"minedInBlockHash,omitempty"`
	// Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block.
	MinedInBlockHeight *int32 `json:"minedInBlockHeight,omitempty"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []GetTransactionDetailsByTransactionIDRIRecipientsInner `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []GetTransactionDetailsByTransactionIDRISendersInner `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the same as `transactionId` for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols `hash` is different from `transactionId` for SegWit transactions.
	TransactionHash string `json:"transactionHash"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	Fee ListConfirmedTransactionsByAddressRIFee `json:"fee"`
	BlockchainSpecific ListConfirmedTransactionsByAddressRIBS `json:"blockchainSpecific"`
}

// NewListConfirmedTransactionsByAddressRI instantiates a new ListConfirmedTransactionsByAddressRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressRI(index int32, recipients []GetTransactionDetailsByTransactionIDRIRecipientsInner, senders []GetTransactionDetailsByTransactionIDRISendersInner, timestamp int32, transactionHash string, transactionId string, fee ListConfirmedTransactionsByAddressRIFee, blockchainSpecific ListConfirmedTransactionsByAddressRIBS) *ListConfirmedTransactionsByAddressRI {
	this := ListConfirmedTransactionsByAddressRI{}
	this.Index = index
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.TransactionId = transactionId
	this.Fee = fee
	this.BlockchainSpecific = blockchainSpecific
	return &this
}

// NewListConfirmedTransactionsByAddressRIWithDefaults instantiates a new ListConfirmedTransactionsByAddressRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressRIWithDefaults() *ListConfirmedTransactionsByAddressRI {
	this := ListConfirmedTransactionsByAddressRI{}
	return &this
}

// GetIndex returns the Index field value
func (o *ListConfirmedTransactionsByAddressRI) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetIndex(v int32) {
	o.Index = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value if set, zero value otherwise.
func (o *ListConfirmedTransactionsByAddressRI) GetMinedInBlockHash() string {
	if o == nil || o.MinedInBlockHash == nil {
		var ret string
		return ret
	}
	return *o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil || o.MinedInBlockHash == nil {
		return nil, false
	}
	return o.MinedInBlockHash, true
}

// HasMinedInBlockHash returns a boolean if a field has been set.
func (o *ListConfirmedTransactionsByAddressRI) HasMinedInBlockHash() bool {
	if o != nil && o.MinedInBlockHash != nil {
		return true
	}

	return false
}

// SetMinedInBlockHash gets a reference to the given string and assigns it to the MinedInBlockHash field.
func (o *ListConfirmedTransactionsByAddressRI) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = &v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value if set, zero value otherwise.
func (o *ListConfirmedTransactionsByAddressRI) GetMinedInBlockHeight() int32 {
	if o == nil || o.MinedInBlockHeight == nil {
		var ret int32
		return ret
	}
	return *o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil || o.MinedInBlockHeight == nil {
		return nil, false
	}
	return o.MinedInBlockHeight, true
}

// HasMinedInBlockHeight returns a boolean if a field has been set.
func (o *ListConfirmedTransactionsByAddressRI) HasMinedInBlockHeight() bool {
	if o != nil && o.MinedInBlockHeight != nil {
		return true
	}

	return false
}

// SetMinedInBlockHeight gets a reference to the given int32 and assigns it to the MinedInBlockHeight field.
func (o *ListConfirmedTransactionsByAddressRI) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = &v
}

// GetRecipients returns the Recipients field value
func (o *ListConfirmedTransactionsByAddressRI) GetRecipients() []GetTransactionDetailsByTransactionIDRIRecipientsInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIRecipientsInner
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetRecipientsOk() ([]GetTransactionDetailsByTransactionIDRIRecipientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetRecipients(v []GetTransactionDetailsByTransactionIDRIRecipientsInner) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListConfirmedTransactionsByAddressRI) GetSenders() []GetTransactionDetailsByTransactionIDRISendersInner {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRISendersInner
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetSendersOk() ([]GetTransactionDetailsByTransactionIDRISendersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Senders, true
}

// SetSenders sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetSenders(v []GetTransactionDetailsByTransactionIDRISendersInner) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListConfirmedTransactionsByAddressRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListConfirmedTransactionsByAddressRI) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListConfirmedTransactionsByAddressRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetFee returns the Fee field value
func (o *ListConfirmedTransactionsByAddressRI) GetFee() ListConfirmedTransactionsByAddressRIFee {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetFeeOk() (*ListConfirmedTransactionsByAddressRIFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetFee(v ListConfirmedTransactionsByAddressRIFee) {
	o.Fee = v
}

// GetBlockchainSpecific returns the BlockchainSpecific field value
func (o *ListConfirmedTransactionsByAddressRI) GetBlockchainSpecific() ListConfirmedTransactionsByAddressRIBS {
	if o == nil {
		var ret ListConfirmedTransactionsByAddressRIBS
		return ret
	}

	return o.BlockchainSpecific
}

// GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressRI) GetBlockchainSpecificOk() (*ListConfirmedTransactionsByAddressRIBS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockchainSpecific, true
}

// SetBlockchainSpecific sets field value
func (o *ListConfirmedTransactionsByAddressRI) SetBlockchainSpecific(v ListConfirmedTransactionsByAddressRIBS) {
	o.BlockchainSpecific = v
}

func (o ListConfirmedTransactionsByAddressRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if o.MinedInBlockHash != nil {
		toSerialize["minedInBlockHash"] = o.MinedInBlockHash
	}
	if o.MinedInBlockHeight != nil {
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

type NullableListConfirmedTransactionsByAddressRI struct {
	value *ListConfirmedTransactionsByAddressRI
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressRI) Get() *ListConfirmedTransactionsByAddressRI {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressRI) Set(val *ListConfirmedTransactionsByAddressRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressRI(val *ListConfirmedTransactionsByAddressRI) *NullableListConfirmedTransactionsByAddressRI {
	return &NullableListConfirmedTransactionsByAddressRI{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


