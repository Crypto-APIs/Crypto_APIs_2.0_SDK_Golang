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

// TransactionMinedDataItem Defines an `item` as one result.
type TransactionMinedDataItem struct {
	// Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
	Blockchain string `json:"blockchain"`
	// Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
	Network string `json:"network"`
	// Defines the unique ID of the specific transaction, i.e. its identification number.
	TransactionId string `json:"transactionId"`
	MinedInBlock TransactionMinedDataItemMinedInBlock `json:"minedInBlock"`
}

// NewTransactionMinedDataItem instantiates a new TransactionMinedDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMinedDataItem(blockchain string, network string, transactionId string, minedInBlock TransactionMinedDataItemMinedInBlock) *TransactionMinedDataItem {
	this := TransactionMinedDataItem{}
	this.Blockchain = blockchain
	this.Network = network
	this.TransactionId = transactionId
	this.MinedInBlock = minedInBlock
	return &this
}

// NewTransactionMinedDataItemWithDefaults instantiates a new TransactionMinedDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionMinedDataItemWithDefaults() *TransactionMinedDataItem {
	this := TransactionMinedDataItem{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *TransactionMinedDataItem) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedDataItem) GetBlockchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *TransactionMinedDataItem) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *TransactionMinedDataItem) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedDataItem) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransactionMinedDataItem) SetNetwork(v string) {
	o.Network = v
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionMinedDataItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedDataItem) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionMinedDataItem) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetMinedInBlock returns the MinedInBlock field value
func (o *TransactionMinedDataItem) GetMinedInBlock() TransactionMinedDataItemMinedInBlock {
	if o == nil {
		var ret TransactionMinedDataItemMinedInBlock
		return ret
	}

	return o.MinedInBlock
}

// GetMinedInBlockOk returns a tuple with the MinedInBlock field value
// and a boolean to check if the value has been set.
func (o *TransactionMinedDataItem) GetMinedInBlockOk() (*TransactionMinedDataItemMinedInBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInBlock, true
}

// SetMinedInBlock sets field value
func (o *TransactionMinedDataItem) SetMinedInBlock(v TransactionMinedDataItemMinedInBlock) {
	o.MinedInBlock = v
}

func (o TransactionMinedDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockchain"] = o.Blockchain
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["minedInBlock"] = o.MinedInBlock
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionMinedDataItem struct {
	value *TransactionMinedDataItem
	isSet bool
}

func (v NullableTransactionMinedDataItem) Get() *TransactionMinedDataItem {
	return v.value
}

func (v *NullableTransactionMinedDataItem) Set(val *TransactionMinedDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMinedDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMinedDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMinedDataItem(val *TransactionMinedDataItem) *NullableTransactionMinedDataItem {
	return &NullableTransactionMinedDataItem{value: val, isSet: true}
}

func (v NullableTransactionMinedDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMinedDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


