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

// GetBlockDetailsByBlockHashRIBSD Dogecoin
type GetBlockDetailsByBlockHashRIBSD struct {
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Represents a random value that can be adjusted to satisfy the Proof of Work.
	Nonce string `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// Represents a specific sub-unit of Doge. Bits have two-decimal precision.
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	MerkleRoot string `json:"merkleRoot"`
	// Defines the numeric representation of the block size excluding the witness data.
	StrippedSize int32 `json:"strippedSize"`
	// Represents the version of the specific block on the blockchain.
	Version int32 `json:"version"`
	// Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit.
	Weight int32 `json:"weight"`
}

// NewGetBlockDetailsByBlockHashRIBSD instantiates a new GetBlockDetailsByBlockHashRIBSD object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashRIBSD(difficulty string, nonce string, size int32, bits string, chainwork string, merkleRoot string, strippedSize int32, version int32, weight int32) *GetBlockDetailsByBlockHashRIBSD {
	this := GetBlockDetailsByBlockHashRIBSD{}
	this.Difficulty = difficulty
	this.Nonce = nonce
	this.Size = size
	this.Bits = bits
	this.Chainwork = chainwork
	this.MerkleRoot = merkleRoot
	this.StrippedSize = strippedSize
	this.Version = version
	this.Weight = weight
	return &this
}

// NewGetBlockDetailsByBlockHashRIBSDWithDefaults instantiates a new GetBlockDetailsByBlockHashRIBSD object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashRIBSDWithDefaults() *GetBlockDetailsByBlockHashRIBSD {
	this := GetBlockDetailsByBlockHashRIBSD{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetNonce returns the Nonce field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetNonce(v string) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetSize(v int32) {
	o.Size = v
}

// GetBits returns the Bits field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetChainworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetChainwork(v string) {
	o.Chainwork = v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetMerkleRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetStrippedSize returns the StrippedSize field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetStrippedSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrippedSize
}

// GetStrippedSizeOk returns a tuple with the StrippedSize field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetStrippedSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrippedSize, true
}

// SetStrippedSize sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetStrippedSize(v int32) {
	o.StrippedSize = v
}

// GetVersion returns the Version field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetVersion(v int32) {
	o.Version = v
}

// GetWeight returns the Weight field value
func (o *GetBlockDetailsByBlockHashRIBSD) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashRIBSD) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *GetBlockDetailsByBlockHashRIBSD) SetWeight(v int32) {
	o.Weight = v
}

func (o GetBlockDetailsByBlockHashRIBSD) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["difficulty"] = o.Difficulty
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["bits"] = o.Bits
	}
	if true {
		toSerialize["chainwork"] = o.Chainwork
	}
	if true {
		toSerialize["merkleRoot"] = o.MerkleRoot
	}
	if true {
		toSerialize["strippedSize"] = o.StrippedSize
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashRIBSD struct {
	value *GetBlockDetailsByBlockHashRIBSD
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashRIBSD) Get() *GetBlockDetailsByBlockHashRIBSD {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashRIBSD) Set(val *GetBlockDetailsByBlockHashRIBSD) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashRIBSD) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashRIBSD) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashRIBSD(val *GetBlockDetailsByBlockHashRIBSD) *NullableGetBlockDetailsByBlockHashRIBSD {
	return &NullableGetBlockDetailsByBlockHashRIBSD{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashRIBSD) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashRIBSD) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


