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

// ListLatestMinedBlocksRIBSL Litecoin
type ListLatestMinedBlocksRIBSL struct {
	// Represents a specific sub-unit of Litecoin. Bits have two-decimal precision.
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	MerkleRoot string `json:"merkleRoot"`
	// Represents a random value that can be adjusted to satisfy the proof of work
	Nonce int64 `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// Defines the numeric representation of the block size excluding the witness data.
	StrippedSize int32 `json:"strippedSize"`
	// Represents the version of the specific block on the blockchain.
	Version int32 `json:"version"`
	// Is the hexadecimal string representation of the block's version.
	VersionHex string `json:"versionHex"`
	// Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit.
	Weight int32 `json:"weight"`
}

// NewListLatestMinedBlocksRIBSL instantiates a new ListLatestMinedBlocksRIBSL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLatestMinedBlocksRIBSL(bits string, chainwork string, difficulty string, merkleRoot string, nonce int64, size int32, strippedSize int32, version int32, versionHex string, weight int32) *ListLatestMinedBlocksRIBSL {
	this := ListLatestMinedBlocksRIBSL{}
	this.Bits = bits
	this.Chainwork = chainwork
	this.Difficulty = difficulty
	this.MerkleRoot = merkleRoot
	this.Nonce = nonce
	this.Size = size
	this.StrippedSize = strippedSize
	this.Version = version
	this.VersionHex = versionHex
	this.Weight = weight
	return &this
}

// NewListLatestMinedBlocksRIBSLWithDefaults instantiates a new ListLatestMinedBlocksRIBSL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLatestMinedBlocksRIBSLWithDefaults() *ListLatestMinedBlocksRIBSL {
	this := ListLatestMinedBlocksRIBSL{}
	return &this
}

// GetBits returns the Bits field value
func (o *ListLatestMinedBlocksRIBSL) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *ListLatestMinedBlocksRIBSL) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *ListLatestMinedBlocksRIBSL) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetChainworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *ListLatestMinedBlocksRIBSL) SetChainwork(v string) {
	o.Chainwork = v
}

// GetDifficulty returns the Difficulty field value
func (o *ListLatestMinedBlocksRIBSL) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *ListLatestMinedBlocksRIBSL) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *ListLatestMinedBlocksRIBSL) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetMerkleRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *ListLatestMinedBlocksRIBSL) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetNonce returns the Nonce field value
func (o *ListLatestMinedBlocksRIBSL) GetNonce() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetNonceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ListLatestMinedBlocksRIBSL) SetNonce(v int64) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *ListLatestMinedBlocksRIBSL) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListLatestMinedBlocksRIBSL) SetSize(v int32) {
	o.Size = v
}

// GetStrippedSize returns the StrippedSize field value
func (o *ListLatestMinedBlocksRIBSL) GetStrippedSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrippedSize
}

// GetStrippedSizeOk returns a tuple with the StrippedSize field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetStrippedSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrippedSize, true
}

// SetStrippedSize sets field value
func (o *ListLatestMinedBlocksRIBSL) SetStrippedSize(v int32) {
	o.StrippedSize = v
}

// GetVersion returns the Version field value
func (o *ListLatestMinedBlocksRIBSL) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListLatestMinedBlocksRIBSL) SetVersion(v int32) {
	o.Version = v
}

// GetVersionHex returns the VersionHex field value
func (o *ListLatestMinedBlocksRIBSL) GetVersionHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionHex
}

// GetVersionHexOk returns a tuple with the VersionHex field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetVersionHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionHex, true
}

// SetVersionHex sets field value
func (o *ListLatestMinedBlocksRIBSL) SetVersionHex(v string) {
	o.VersionHex = v
}

// GetWeight returns the Weight field value
func (o *ListLatestMinedBlocksRIBSL) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSL) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ListLatestMinedBlocksRIBSL) SetWeight(v int32) {
	o.Weight = v
}

func (o ListLatestMinedBlocksRIBSL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bits"] = o.Bits
	}
	if true {
		toSerialize["chainwork"] = o.Chainwork
	}
	if true {
		toSerialize["difficulty"] = o.Difficulty
	}
	if true {
		toSerialize["merkleRoot"] = o.MerkleRoot
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["strippedSize"] = o.StrippedSize
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["versionHex"] = o.VersionHex
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableListLatestMinedBlocksRIBSL struct {
	value *ListLatestMinedBlocksRIBSL
	isSet bool
}

func (v NullableListLatestMinedBlocksRIBSL) Get() *ListLatestMinedBlocksRIBSL {
	return v.value
}

func (v *NullableListLatestMinedBlocksRIBSL) Set(val *ListLatestMinedBlocksRIBSL) {
	v.value = val
	v.isSet = true
}

func (v NullableListLatestMinedBlocksRIBSL) IsSet() bool {
	return v.isSet
}

func (v *NullableListLatestMinedBlocksRIBSL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLatestMinedBlocksRIBSL(val *ListLatestMinedBlocksRIBSL) *NullableListLatestMinedBlocksRIBSL {
	return &NullableListLatestMinedBlocksRIBSL{value: val, isSet: true}
}

func (v NullableListLatestMinedBlocksRIBSL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLatestMinedBlocksRIBSL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


