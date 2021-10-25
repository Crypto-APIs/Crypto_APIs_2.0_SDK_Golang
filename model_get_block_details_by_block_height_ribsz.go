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

// GetBlockDetailsByBlockHeightRIBSZ Zcash
type GetBlockDetailsByBlockHeightRIBSZ struct {
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Represents a random value that can be adjusted to satisfy the Proof of Work.
	Nonce string `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// Represents a specific sub-unit of Zcash. Bits have two-decimal precision
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	Merkleroot string `json:"merkleroot"`
	// Represents the block version number.
	Version int32 `json:"version"`
}

// NewGetBlockDetailsByBlockHeightRIBSZ instantiates a new GetBlockDetailsByBlockHeightRIBSZ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHeightRIBSZ(difficulty string, nonce string, size int32, bits string, chainwork string, merkleroot string, version int32) *GetBlockDetailsByBlockHeightRIBSZ {
	this := GetBlockDetailsByBlockHeightRIBSZ{}
	this.Difficulty = difficulty
	this.Nonce = nonce
	this.Size = size
	this.Bits = bits
	this.Chainwork = chainwork
	this.Merkleroot = merkleroot
	this.Version = version
	return &this
}

// NewGetBlockDetailsByBlockHeightRIBSZWithDefaults instantiates a new GetBlockDetailsByBlockHeightRIBSZ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHeightRIBSZWithDefaults() *GetBlockDetailsByBlockHeightRIBSZ {
	this := GetBlockDetailsByBlockHeightRIBSZ{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetDifficultyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetNonce returns the Nonce field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetNonceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetNonce(v string) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetSize(v int32) {
	o.Size = v
}

// GetBits returns the Bits field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetBitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetChainworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetChainwork(v string) {
	o.Chainwork = v
}

// GetMerkleroot returns the Merkleroot field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetMerkleroot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Merkleroot
}

// GetMerklerootOk returns a tuple with the Merkleroot field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetMerklerootOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Merkleroot, true
}

// SetMerkleroot sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetMerkleroot(v string) {
	o.Merkleroot = v
}

// GetVersion returns the Version field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHeightRIBSZ) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetBlockDetailsByBlockHeightRIBSZ) SetVersion(v int32) {
	o.Version = v
}

func (o GetBlockDetailsByBlockHeightRIBSZ) MarshalJSON() ([]byte, error) {
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
		toSerialize["merkleroot"] = o.Merkleroot
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHeightRIBSZ struct {
	value *GetBlockDetailsByBlockHeightRIBSZ
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHeightRIBSZ) Get() *GetBlockDetailsByBlockHeightRIBSZ {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHeightRIBSZ) Set(val *GetBlockDetailsByBlockHeightRIBSZ) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHeightRIBSZ) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHeightRIBSZ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHeightRIBSZ(val *GetBlockDetailsByBlockHeightRIBSZ) *NullableGetBlockDetailsByBlockHeightRIBSZ {
	return &NullableGetBlockDetailsByBlockHeightRIBSZ{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHeightRIBSZ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHeightRIBSZ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

