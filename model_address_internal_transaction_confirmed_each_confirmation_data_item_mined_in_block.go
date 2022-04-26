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

// AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock Refers to the specific block the transaction was mined in.
type AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock struct {
	// Defines the number of blocks in the blockchain preceding this specific block.
	Height int32 `json:"height"`
	// Represents the hash of the block's header, i.e. an output that has a fixed length.
	Hash string `json:"hash"`
	// Defines the exact date/time when this transaction was mined in seconds since Unix Epoch time.
	Timestamp int32 `json:"timestamp"`
}

// NewAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock instantiates a new AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock(height int32, hash string, timestamp int32) *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock {
	this := AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock{}
	this.Height = height
	this.Hash = hash
	this.Timestamp = timestamp
	return &this
}

// NewAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlockWithDefaults instantiates a new AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlockWithDefaults() *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock {
	this := AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock{}
	return &this
}

// GetHeight returns the Height field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) SetHeight(v int32) {
	o.Height = v
}

// GetHash returns the Hash field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) SetHash(v string) {
	o.Hash = v
}

// GetTimestamp returns the Timestamp field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) SetTimestamp(v int32) {
	o.Timestamp = v
}

func (o AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) MarshalJSON() ([]byte, error) {
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

type NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock struct {
	value *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock
	isSet bool
}

func (v NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) Get() *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock {
	return v.value
}

func (v *NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) Set(val *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock(val *AddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) *NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock {
	return &NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock{value: val, isSet: true}
}

func (v NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInternalTransactionConfirmedEachConfirmationDataItemMinedInBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


