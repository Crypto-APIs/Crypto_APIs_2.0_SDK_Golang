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

// ListLatestMinedBlocksRIBSB Bitcoin
type ListLatestMinedBlocksRIBSB struct {
	// A sub-unit of BCH equal to 0.000001 BCH, or 100 Satoshi, and is the same as microbitcoincash (μBCH). Bits have two-decimal precision.
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty *string `json:"difficulty,omitempty"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	MerkleRoot string `json:"merkleRoot"`
	// Represents a random value that can be adjusted to satisfy the proof of work
	Nonce *int64 `json:"nonce,omitempty"`
	// Represents the block size
	Size *int32 `json:"size,omitempty"`
	// Defines the numeric representation of the block size excluding the witness data.
	StrippedSize int32 `json:"strippedSize"`
	// Represents the version of the specific block on the blockchain.
	Version int32 `json:"version"`
	// Is the hexadecimal string representation of the block's version.
	VersionHex string `json:"versionHex"`
	// Represents a measurement to compare the size of different transactions to each other in proportion to the block size limit.
	Weight int32 `json:"weight"`
}

// NewListLatestMinedBlocksRIBSB instantiates a new ListLatestMinedBlocksRIBSB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLatestMinedBlocksRIBSB(bits string, chainwork string, merkleRoot string, strippedSize int32, version int32, versionHex string, weight int32) *ListLatestMinedBlocksRIBSB {
	this := ListLatestMinedBlocksRIBSB{}
	this.Bits = bits
	this.Chainwork = chainwork
	this.MerkleRoot = merkleRoot
	this.StrippedSize = strippedSize
	this.Version = version
	this.VersionHex = versionHex
	this.Weight = weight
	return &this
}

// NewListLatestMinedBlocksRIBSBWithDefaults instantiates a new ListLatestMinedBlocksRIBSB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLatestMinedBlocksRIBSBWithDefaults() *ListLatestMinedBlocksRIBSB {
	this := ListLatestMinedBlocksRIBSB{}
	return &this
}

// GetBits returns the Bits field value
func (o *ListLatestMinedBlocksRIBSB) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *ListLatestMinedBlocksRIBSB) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *ListLatestMinedBlocksRIBSB) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetChainworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *ListLatestMinedBlocksRIBSB) SetChainwork(v string) {
	o.Chainwork = v
}

// GetDifficulty returns the Difficulty field value if set, zero value otherwise.
func (o *ListLatestMinedBlocksRIBSB) GetDifficulty() string {
	if o == nil || o.Difficulty == nil {
		var ret string
		return ret
	}
	return *o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetDifficultyOk() (*string, bool) {
	if o == nil || o.Difficulty == nil {
		return nil, false
	}
	return o.Difficulty, true
}

// HasDifficulty returns a boolean if a field has been set.
func (o *ListLatestMinedBlocksRIBSB) HasDifficulty() bool {
	if o != nil && o.Difficulty != nil {
		return true
	}

	return false
}

// SetDifficulty gets a reference to the given string and assigns it to the Difficulty field.
func (o *ListLatestMinedBlocksRIBSB) SetDifficulty(v string) {
	o.Difficulty = &v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *ListLatestMinedBlocksRIBSB) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetMerkleRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *ListLatestMinedBlocksRIBSB) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *ListLatestMinedBlocksRIBSB) GetNonce() int64 {
	if o == nil || o.Nonce == nil {
		var ret int64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetNonceOk() (*int64, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *ListLatestMinedBlocksRIBSB) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int64 and assigns it to the Nonce field.
func (o *ListLatestMinedBlocksRIBSB) SetNonce(v int64) {
	o.Nonce = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ListLatestMinedBlocksRIBSB) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ListLatestMinedBlocksRIBSB) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ListLatestMinedBlocksRIBSB) SetSize(v int32) {
	o.Size = &v
}

// GetStrippedSize returns the StrippedSize field value
func (o *ListLatestMinedBlocksRIBSB) GetStrippedSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrippedSize
}

// GetStrippedSizeOk returns a tuple with the StrippedSize field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetStrippedSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrippedSize, true
}

// SetStrippedSize sets field value
func (o *ListLatestMinedBlocksRIBSB) SetStrippedSize(v int32) {
	o.StrippedSize = v
}

// GetVersion returns the Version field value
func (o *ListLatestMinedBlocksRIBSB) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListLatestMinedBlocksRIBSB) SetVersion(v int32) {
	o.Version = v
}

// GetVersionHex returns the VersionHex field value
func (o *ListLatestMinedBlocksRIBSB) GetVersionHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionHex
}

// GetVersionHexOk returns a tuple with the VersionHex field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetVersionHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionHex, true
}

// SetVersionHex sets field value
func (o *ListLatestMinedBlocksRIBSB) SetVersionHex(v string) {
	o.VersionHex = v
}

// GetWeight returns the Weight field value
func (o *ListLatestMinedBlocksRIBSB) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSB) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ListLatestMinedBlocksRIBSB) SetWeight(v int32) {
	o.Weight = v
}

func (o ListLatestMinedBlocksRIBSB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bits"] = o.Bits
	}
	if true {
		toSerialize["chainwork"] = o.Chainwork
	}
	if o.Difficulty != nil {
		toSerialize["difficulty"] = o.Difficulty
	}
	if true {
		toSerialize["merkleRoot"] = o.MerkleRoot
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Size != nil {
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

type NullableListLatestMinedBlocksRIBSB struct {
	value *ListLatestMinedBlocksRIBSB
	isSet bool
}

func (v NullableListLatestMinedBlocksRIBSB) Get() *ListLatestMinedBlocksRIBSB {
	return v.value
}

func (v *NullableListLatestMinedBlocksRIBSB) Set(val *ListLatestMinedBlocksRIBSB) {
	v.value = val
	v.isSet = true
}

func (v NullableListLatestMinedBlocksRIBSB) IsSet() bool {
	return v.isSet
}

func (v *NullableListLatestMinedBlocksRIBSB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLatestMinedBlocksRIBSB(val *ListLatestMinedBlocksRIBSB) *NullableListLatestMinedBlocksRIBSB {
	return &NullableListLatestMinedBlocksRIBSB{value: val, isSet: true}
}

func (v NullableListLatestMinedBlocksRIBSB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLatestMinedBlocksRIBSB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


