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

// GetXRPRippleBlockDetailsByBlockHashResponseItem struct for GetXRPRippleBlockDetailsByBlockHashResponseItem
type GetXRPRippleBlockDetailsByBlockHashResponseItem struct {
	// Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	BlockHash string `json:"blockHash"`
	// Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
	BlockHeight int32 `json:"blockHeight"`
	// Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string.
	NextBlockHash string `json:"nextBlockHash"`
	// Represents the hash of the previous block, also known as the parent block.
	PreviousBlockHash string `json:"previousBlockHash"`
	// Defines the exact date/time when this block was mined in Unix Timestamp.
	Timestamp int32 `json:"timestamp"`
	TotalCoins GetXRPRippleBlockDetailsByBlockHashResponseItemTotalCoins `json:"totalCoins"`
	TotalFees GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees `json:"totalFees"`
	// Represents the total number of all transactions as part of this block.
	TransactionsCount int32 `json:"transactionsCount"`
}

// NewGetXRPRippleBlockDetailsByBlockHashResponseItem instantiates a new GetXRPRippleBlockDetailsByBlockHashResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetXRPRippleBlockDetailsByBlockHashResponseItem(blockHash string, blockHeight int32, nextBlockHash string, previousBlockHash string, timestamp int32, totalCoins GetXRPRippleBlockDetailsByBlockHashResponseItemTotalCoins, totalFees GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees, transactionsCount int32) *GetXRPRippleBlockDetailsByBlockHashResponseItem {
	this := GetXRPRippleBlockDetailsByBlockHashResponseItem{}
	this.BlockHash = blockHash
	this.BlockHeight = blockHeight
	this.NextBlockHash = nextBlockHash
	this.PreviousBlockHash = previousBlockHash
	this.Timestamp = timestamp
	this.TotalCoins = totalCoins
	this.TotalFees = totalFees
	this.TransactionsCount = transactionsCount
	return &this
}

// NewGetXRPRippleBlockDetailsByBlockHashResponseItemWithDefaults instantiates a new GetXRPRippleBlockDetailsByBlockHashResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetXRPRippleBlockDetailsByBlockHashResponseItemWithDefaults() *GetXRPRippleBlockDetailsByBlockHashResponseItem {
	this := GetXRPRippleBlockDetailsByBlockHashResponseItem{}
	return &this
}

// GetBlockHash returns the BlockHash field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetBlockHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockHeight returns the BlockHeight field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetBlockHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BlockHeight, true
}

// SetBlockHeight sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetBlockHeight(v int32) {
	o.BlockHeight = v
}

// GetNextBlockHash returns the NextBlockHash field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetNextBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextBlockHash
}

// GetNextBlockHashOk returns a tuple with the NextBlockHash field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetNextBlockHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextBlockHash, true
}

// SetNextBlockHash sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetNextBlockHash(v string) {
	o.NextBlockHash = v
}

// GetPreviousBlockHash returns the PreviousBlockHash field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetPreviousBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousBlockHash
}

// GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetPreviousBlockHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PreviousBlockHash, true
}

// SetPreviousBlockHash sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetPreviousBlockHash(v string) {
	o.PreviousBlockHash = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTimestampOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTotalCoins returns the TotalCoins field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTotalCoins() GetXRPRippleBlockDetailsByBlockHashResponseItemTotalCoins {
	if o == nil {
		var ret GetXRPRippleBlockDetailsByBlockHashResponseItemTotalCoins
		return ret
	}

	return o.TotalCoins
}

// GetTotalCoinsOk returns a tuple with the TotalCoins field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTotalCoinsOk() (*GetXRPRippleBlockDetailsByBlockHashResponseItemTotalCoins, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCoins, true
}

// SetTotalCoins sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetTotalCoins(v GetXRPRippleBlockDetailsByBlockHashResponseItemTotalCoins) {
	o.TotalCoins = v
}

// GetTotalFees returns the TotalFees field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTotalFees() GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees {
	if o == nil {
		var ret GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees
		return ret
	}

	return o.TotalFees
}

// GetTotalFeesOk returns a tuple with the TotalFees field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTotalFeesOk() (*GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalFees, true
}

// SetTotalFees sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetTotalFees(v GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees) {
	o.TotalFees = v
}

// GetTransactionsCount returns the TransactionsCount field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionsCount
}

// GetTransactionsCountOk returns a tuple with the TransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) GetTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionsCount, true
}

// SetTransactionsCount sets field value
func (o *GetXRPRippleBlockDetailsByBlockHashResponseItem) SetTransactionsCount(v int32) {
	o.TransactionsCount = v
}

func (o GetXRPRippleBlockDetailsByBlockHashResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockHash"] = o.BlockHash
	}
	if true {
		toSerialize["blockHeight"] = o.BlockHeight
	}
	if true {
		toSerialize["nextBlockHash"] = o.NextBlockHash
	}
	if true {
		toSerialize["previousBlockHash"] = o.PreviousBlockHash
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["totalCoins"] = o.TotalCoins
	}
	if true {
		toSerialize["totalFees"] = o.TotalFees
	}
	if true {
		toSerialize["transactionsCount"] = o.TransactionsCount
	}
	return json.Marshal(toSerialize)
}

type NullableGetXRPRippleBlockDetailsByBlockHashResponseItem struct {
	value *GetXRPRippleBlockDetailsByBlockHashResponseItem
	isSet bool
}

func (v NullableGetXRPRippleBlockDetailsByBlockHashResponseItem) Get() *GetXRPRippleBlockDetailsByBlockHashResponseItem {
	return v.value
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHashResponseItem) Set(val *GetXRPRippleBlockDetailsByBlockHashResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleBlockDetailsByBlockHashResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHashResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleBlockDetailsByBlockHashResponseItem(val *GetXRPRippleBlockDetailsByBlockHashResponseItem) *NullableGetXRPRippleBlockDetailsByBlockHashResponseItem {
	return &NullableGetXRPRippleBlockDetailsByBlockHashResponseItem{value: val, isSet: true}
}

func (v NullableGetXRPRippleBlockDetailsByBlockHashResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleBlockDetailsByBlockHashResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

