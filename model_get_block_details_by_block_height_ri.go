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

// GetBlockDetailsByBlockHeightRI struct for GetBlockDetailsByBlockHeightRI
type GetBlockDetailsByBlockHeightRI struct {
	// Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	Hash string `json:"hash"`
	// Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \"Genesis block\".
	Height int32 `json:"height"`
	// Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string.
	NextBlockHash string `json:"nextBlockHash"`
	// Represents the hash of the previous block, also known as the parent block.
	PreviousBlockHash string `json:"previousBlockHash"`
	// Defines the exact date/time when this block was mined in Unix
	Timestamp int32 `json:"timestamp"`
	// Represents the total number of all transactions as part of this block.
	TransactionsCount int32 `json:"transactionsCount"`
	BlockchainSpecific GetBlockDetailsByBlockHeightRIBS `json:"blockchainSpecific"`
}

// NewGetBlockDetailsByBlockHeightRI instantiates a new GetBlockDetailsByBlockHeightRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHeightRI(hash string, height int32, nextBlockHash string, previousBlockHash string, timestamp int32, transactionsCount int32, blockchainSpecific GetBlockDetailsByBlockHeightRIBS) *GetBlockDetailsByBlockHeightRI {
	this := GetBlockDetailsByBlockHeightRI{}
	this.Hash = hash
	this.Height = height
	this.NextBlockHash = nextBlockHash
	this.PreviousBlockHash = previousBlockHash
	this.Timestamp = timestamp
	this.TransactionsCount = transactionsCount
	this.BlockchainSpecific = blockchainSpecific
	return &this
}

// NewGetBlockDetailsByBlockHeightRIWithDefaults instantiates a new GetBlockDetailsByBlockHeightRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHeightRIWithDefaults() *GetBlockDetailsByBlockHeightRI {
	this := GetBlockDetailsByBlockHeightRI{}
	return &this
}

// GetHash returns the Hash field value
func (o *GetBlockDetailsByBlockHeightRI) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetHash(v string) {
	o.Hash = v
}

// GetHeight returns the Height field value
func (o *GetBlockDetailsByBlockHeightRI) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetHeight(v int32) {
	o.Height = v
}

// GetNextBlockHash returns the NextBlockHash field value
func (o *GetBlockDetailsByBlockHeightRI) GetNextBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextBlockHash
}

// GetNextBlockHashOk returns a tuple with the NextBlockHash field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetNextBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextBlockHash, true
}

// SetNextBlockHash sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetNextBlockHash(v string) {
	o.NextBlockHash = v
}

// GetPreviousBlockHash returns the PreviousBlockHash field value
func (o *GetBlockDetailsByBlockHeightRI) GetPreviousBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousBlockHash
}

// GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetPreviousBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousBlockHash, true
}

// SetPreviousBlockHash sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetPreviousBlockHash(v string) {
	o.PreviousBlockHash = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetBlockDetailsByBlockHeightRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionsCount returns the TransactionsCount field value
func (o *GetBlockDetailsByBlockHeightRI) GetTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionsCount
}

// GetTransactionsCountOk returns a tuple with the TransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetTransactionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionsCount, true
}

// SetTransactionsCount sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetTransactionsCount(v int32) {
	o.TransactionsCount = v
}

// GetBlockchainSpecific returns the BlockchainSpecific field value
func (o *GetBlockDetailsByBlockHeightRI) GetBlockchainSpecific() GetBlockDetailsByBlockHeightRIBS {
	if o == nil {
		var ret GetBlockDetailsByBlockHeightRIBS
		return ret
	}

	return o.BlockchainSpecific
}

// GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRI) GetBlockchainSpecificOk() (*GetBlockDetailsByBlockHeightRIBS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockchainSpecific, true
}

// SetBlockchainSpecific sets field value
func (o *GetBlockDetailsByBlockHeightRI) SetBlockchainSpecific(v GetBlockDetailsByBlockHeightRIBS) {
	o.BlockchainSpecific = v
}

func (o GetBlockDetailsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["height"] = o.Height
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
		toSerialize["transactionsCount"] = o.TransactionsCount
	}
	if true {
		toSerialize["blockchainSpecific"] = o.BlockchainSpecific
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHeightRI struct {
	value *GetBlockDetailsByBlockHeightRI
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHeightRI) Get() *GetBlockDetailsByBlockHeightRI {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHeightRI) Set(val *GetBlockDetailsByBlockHeightRI) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHeightRI) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHeightRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHeightRI(val *GetBlockDetailsByBlockHeightRI) *NullableGetBlockDetailsByBlockHeightRI {
	return &NullableGetBlockDetailsByBlockHeightRI{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHeightRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHeightRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


