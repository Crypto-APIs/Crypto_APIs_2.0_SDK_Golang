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

// AddressCoinsTransactionConfirmedDataItemMinedInBlock Defines the block height in which this transaction was mined and confirmed in.
type AddressCoinsTransactionConfirmedDataItemMinedInBlock struct {
	// Defines the number of blocks in the blockchain preceding this specific block.
	Height int32 `json:"height"`
	// Represents the hash of the block's header, i.e. an output that has a fixed length.
	Hash string `json:"hash"`
	// Defines the exact date/time when this transaction was mined in seconds since Unix Epoch time.
	Timestamp int32 `json:"timestamp"`
}

// NewAddressCoinsTransactionConfirmedDataItemMinedInBlock instantiates a new AddressCoinsTransactionConfirmedDataItemMinedInBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCoinsTransactionConfirmedDataItemMinedInBlock(height int32, hash string, timestamp int32) *AddressCoinsTransactionConfirmedDataItemMinedInBlock {
	this := AddressCoinsTransactionConfirmedDataItemMinedInBlock{}
	this.Height = height
	this.Hash = hash
	this.Timestamp = timestamp
	return &this
}

// NewAddressCoinsTransactionConfirmedDataItemMinedInBlockWithDefaults instantiates a new AddressCoinsTransactionConfirmedDataItemMinedInBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCoinsTransactionConfirmedDataItemMinedInBlockWithDefaults() *AddressCoinsTransactionConfirmedDataItemMinedInBlock {
	this := AddressCoinsTransactionConfirmedDataItemMinedInBlock{}
	return &this
}

// GetHeight returns the Height field value
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) SetHeight(v int32) {
	o.Height = v
}

// GetHash returns the Hash field value
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) SetHash(v string) {
	o.Hash = v
}

// GetTimestamp returns the Timestamp field value
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) SetTimestamp(v int32) {
	o.Timestamp = v
}

func (o AddressCoinsTransactionConfirmedDataItemMinedInBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock struct {
	value *AddressCoinsTransactionConfirmedDataItemMinedInBlock
	isSet bool
}

func (v NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock) Get() *AddressCoinsTransactionConfirmedDataItemMinedInBlock {
	return v.value
}

func (v *NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock) Set(val *AddressCoinsTransactionConfirmedDataItemMinedInBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCoinsTransactionConfirmedDataItemMinedInBlock(val *AddressCoinsTransactionConfirmedDataItemMinedInBlock) *NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock {
	return &NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock{value: val, isSet: true}
}

func (v NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCoinsTransactionConfirmedDataItemMinedInBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


