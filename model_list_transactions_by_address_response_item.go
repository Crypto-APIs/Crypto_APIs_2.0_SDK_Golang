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

// ListTransactionsByAddressResponseItem struct for ListTransactionsByAddressResponseItem
type ListTransactionsByAddressResponseItem struct {
	// Represents the index position of the transaction in the block.
	Index int32 `json:"index"`
	// Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash *string `json:"minedInBlockHash,omitempty"`
	// Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block.
	MinedInBlockHeight *int32 `json:"minedInBlockHeight,omitempty"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []GetTransactionDetailsByTransactionIDResponseItemRecipients `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []GetTransactionDetailsByTransactionIDResponseItemSenders `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the same as `transactionId` for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols `hash` is different from `transactionId` for SegWit transactions.
	TransactionHash string `json:"transactionHash"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	Fee GetTransactionDetailsByTransactionIDResponseItemFee `json:"fee"`
	BlockchainSpecific ListTransactionsByAddressResponseItemBlockchainSpecific `json:"blockchainSpecific"`
}

// NewListTransactionsByAddressResponseItem instantiates a new ListTransactionsByAddressResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByAddressResponseItem(index int32, recipients []GetTransactionDetailsByTransactionIDResponseItemRecipients, senders []GetTransactionDetailsByTransactionIDResponseItemSenders, timestamp int32, transactionHash string, transactionId string, fee GetTransactionDetailsByTransactionIDResponseItemFee, blockchainSpecific ListTransactionsByAddressResponseItemBlockchainSpecific) *ListTransactionsByAddressResponseItem {
	this := ListTransactionsByAddressResponseItem{}
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

// NewListTransactionsByAddressResponseItemWithDefaults instantiates a new ListTransactionsByAddressResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByAddressResponseItemWithDefaults() *ListTransactionsByAddressResponseItem {
	this := ListTransactionsByAddressResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *ListTransactionsByAddressResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListTransactionsByAddressResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItem) GetMinedInBlockHash() string {
	if o == nil || o.MinedInBlockHash == nil {
		var ret string
		return ret
	}
	return *o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil || o.MinedInBlockHash == nil {
		return nil, false
	}
	return o.MinedInBlockHash, true
}

// HasMinedInBlockHash returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItem) HasMinedInBlockHash() bool {
	if o != nil && o.MinedInBlockHash != nil {
		return true
	}

	return false
}

// SetMinedInBlockHash gets a reference to the given string and assigns it to the MinedInBlockHash field.
func (o *ListTransactionsByAddressResponseItem) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = &v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value if set, zero value otherwise.
func (o *ListTransactionsByAddressResponseItem) GetMinedInBlockHeight() int32 {
	if o == nil || o.MinedInBlockHeight == nil {
		var ret int32
		return ret
	}
	return *o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil || o.MinedInBlockHeight == nil {
		return nil, false
	}
	return o.MinedInBlockHeight, true
}

// HasMinedInBlockHeight returns a boolean if a field has been set.
func (o *ListTransactionsByAddressResponseItem) HasMinedInBlockHeight() bool {
	if o != nil && o.MinedInBlockHeight != nil {
		return true
	}

	return false
}

// SetMinedInBlockHeight gets a reference to the given int32 and assigns it to the MinedInBlockHeight field.
func (o *ListTransactionsByAddressResponseItem) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = &v
}

// GetRecipients returns the Recipients field value
func (o *ListTransactionsByAddressResponseItem) GetRecipients() []GetTransactionDetailsByTransactionIDResponseItemRecipients {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDResponseItemRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetRecipientsOk() (*[]GetTransactionDetailsByTransactionIDResponseItemRecipients, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *ListTransactionsByAddressResponseItem) SetRecipients(v []GetTransactionDetailsByTransactionIDResponseItemRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListTransactionsByAddressResponseItem) GetSenders() []GetTransactionDetailsByTransactionIDResponseItemSenders {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDResponseItemSenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetSendersOk() (*[]GetTransactionDetailsByTransactionIDResponseItemSenders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *ListTransactionsByAddressResponseItem) SetSenders(v []GetTransactionDetailsByTransactionIDResponseItemSenders) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListTransactionsByAddressResponseItem) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetTimestampOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListTransactionsByAddressResponseItem) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListTransactionsByAddressResponseItem) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetTransactionHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListTransactionsByAddressResponseItem) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListTransactionsByAddressResponseItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListTransactionsByAddressResponseItem) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetFee returns the Fee field value
func (o *ListTransactionsByAddressResponseItem) GetFee() GetTransactionDetailsByTransactionIDResponseItemFee {
	if o == nil {
		var ret GetTransactionDetailsByTransactionIDResponseItemFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetFeeOk() (*GetTransactionDetailsByTransactionIDResponseItemFee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListTransactionsByAddressResponseItem) SetFee(v GetTransactionDetailsByTransactionIDResponseItemFee) {
	o.Fee = v
}

// GetBlockchainSpecific returns the BlockchainSpecific field value
func (o *ListTransactionsByAddressResponseItem) GetBlockchainSpecific() ListTransactionsByAddressResponseItemBlockchainSpecific {
	if o == nil {
		var ret ListTransactionsByAddressResponseItemBlockchainSpecific
		return ret
	}

	return o.BlockchainSpecific
}

// GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItem) GetBlockchainSpecificOk() (*ListTransactionsByAddressResponseItemBlockchainSpecific, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BlockchainSpecific, true
}

// SetBlockchainSpecific sets field value
func (o *ListTransactionsByAddressResponseItem) SetBlockchainSpecific(v ListTransactionsByAddressResponseItemBlockchainSpecific) {
	o.BlockchainSpecific = v
}

func (o ListTransactionsByAddressResponseItem) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByAddressResponseItem struct {
	value *ListTransactionsByAddressResponseItem
	isSet bool
}

func (v NullableListTransactionsByAddressResponseItem) Get() *ListTransactionsByAddressResponseItem {
	return v.value
}

func (v *NullableListTransactionsByAddressResponseItem) Set(val *ListTransactionsByAddressResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByAddressResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByAddressResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByAddressResponseItem(val *ListTransactionsByAddressResponseItem) *NullableListTransactionsByAddressResponseItem {
	return &NullableListTransactionsByAddressResponseItem{value: val, isSet: true}
}

func (v NullableListTransactionsByAddressResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByAddressResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


