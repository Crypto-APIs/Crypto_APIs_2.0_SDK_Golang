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

// GetBlockDetailsByBlockHashFromCallbackRIBSL Litecoin
type GetBlockDetailsByBlockHashFromCallbackRIBSL struct {
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
	Strippedsize int32 `json:"strippedsize"`
	// Represents the version of the specific block on the blockchain.
	Version int32 `json:"version"`
	// Is the hexadecimal string representation of the block's version.
	VersionHex string `json:"versionHex"`
	// Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit.
	Weight int32 `json:"weight"`
}

// NewGetBlockDetailsByBlockHashFromCallbackRIBSL instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashFromCallbackRIBSL(bits string, chainwork string, difficulty string, merkleRoot string, nonce int64, size int32, strippedsize int32, version int32, versionHex string, weight int32) *GetBlockDetailsByBlockHashFromCallbackRIBSL {
	this := GetBlockDetailsByBlockHashFromCallbackRIBSL{}
	this.Bits = bits
	this.Chainwork = chainwork
	this.Difficulty = difficulty
	this.MerkleRoot = merkleRoot
	this.Nonce = nonce
	this.Size = size
	this.Strippedsize = strippedsize
	this.Version = version
	this.VersionHex = versionHex
	this.Weight = weight
	return &this
}

// NewGetBlockDetailsByBlockHashFromCallbackRIBSLWithDefaults instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashFromCallbackRIBSLWithDefaults() *GetBlockDetailsByBlockHashFromCallbackRIBSL {
	this := GetBlockDetailsByBlockHashFromCallbackRIBSL{}
	return &this
}

// GetBits returns the Bits field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetChainworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetChainwork(v string) {
	o.Chainwork = v
}

// GetDifficulty returns the Difficulty field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetMerkleRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetNonce returns the Nonce field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetNonce() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetNonceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetNonce(v int64) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetSize(v int32) {
	o.Size = v
}

// GetStrippedsize returns the Strippedsize field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetStrippedsize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Strippedsize
}

// GetStrippedsizeOk returns a tuple with the Strippedsize field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetStrippedsizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strippedsize, true
}

// SetStrippedsize sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetStrippedsize(v int32) {
	o.Strippedsize = v
}

// GetVersion returns the Version field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetVersion(v int32) {
	o.Version = v
}

// GetVersionHex returns the VersionHex field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersionHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionHex
}

// GetVersionHexOk returns a tuple with the VersionHex field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetVersionHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionHex, true
}

// SetVersionHex sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetVersionHex(v string) {
	o.VersionHex = v
}

// GetWeight returns the Weight field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSL) SetWeight(v int32) {
	o.Weight = v
}

func (o GetBlockDetailsByBlockHashFromCallbackRIBSL) MarshalJSON() ([]byte, error) {
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
		toSerialize["strippedsize"] = o.Strippedsize
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

type NullableGetBlockDetailsByBlockHashFromCallbackRIBSL struct {
	value *GetBlockDetailsByBlockHashFromCallbackRIBSL
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSL) Get() *GetBlockDetailsByBlockHashFromCallbackRIBSL {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSL) Set(val *GetBlockDetailsByBlockHashFromCallbackRIBSL) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSL) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashFromCallbackRIBSL(val *GetBlockDetailsByBlockHashFromCallbackRIBSL) *NullableGetBlockDetailsByBlockHashFromCallbackRIBSL {
	return &NullableGetBlockDetailsByBlockHashFromCallbackRIBSL{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


