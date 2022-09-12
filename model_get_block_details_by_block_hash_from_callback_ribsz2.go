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

// GetBlockDetailsByBlockHashFromCallbackRIBSZ2 Zcash
type GetBlockDetailsByBlockHashFromCallbackRIBSZ2 struct {
	// Represents a specific sub-unit of Zcash. Bits have two-decimal precision
	Bits string `json:"bits"`
	// Represents a hexadecimal number of all the hashes necessary to produce the current chain. E.g., when converting 0000000000000000000000000000000000000000000086859f7a841475b236fd to a decimal you get 635262017308958427068157 hashes, or 635262 exahashes.
	Chainwork string `json:"chainwork"`
	// Defines the single and final (root) node of a Merkle tree. It is the combined hash of all transactions' hashes that are part of a blockchain block.
	MerkleRoot string `json:"merkleRoot"`
	// Represents a random value that can be adjusted to satisfy the proof of work
	Nonce string `json:"nonce"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// Represents the transaction version number.
	Version int32 `json:"version"`
}

// NewGetBlockDetailsByBlockHashFromCallbackRIBSZ2 instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSZ2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashFromCallbackRIBSZ2(bits string, chainwork string, merkleRoot string, nonce string, size int32, version int32) *GetBlockDetailsByBlockHashFromCallbackRIBSZ2 {
	this := GetBlockDetailsByBlockHashFromCallbackRIBSZ2{}
	this.Bits = bits
	this.Chainwork = chainwork
	this.MerkleRoot = merkleRoot
	this.Nonce = nonce
	this.Size = size
	this.Version = version
	return &this
}

// NewGetBlockDetailsByBlockHashFromCallbackRIBSZ2WithDefaults instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSZ2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashFromCallbackRIBSZ2WithDefaults() *GetBlockDetailsByBlockHashFromCallbackRIBSZ2 {
	this := GetBlockDetailsByBlockHashFromCallbackRIBSZ2{}
	return &this
}

// GetBits returns the Bits field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetBits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bits
}

// GetBitsOk returns a tuple with the Bits field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetBitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bits, true
}

// SetBits sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) SetBits(v string) {
	o.Bits = v
}

// GetChainwork returns the Chainwork field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetChainwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chainwork
}

// GetChainworkOk returns a tuple with the Chainwork field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetChainworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chainwork, true
}

// SetChainwork sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) SetChainwork(v string) {
	o.Chainwork = v
}

// GetMerkleRoot returns the MerkleRoot field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetMerkleRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetMerkleRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerkleRoot, true
}

// SetMerkleRoot sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) SetMerkleRoot(v string) {
	o.MerkleRoot = v
}

// GetNonce returns the Nonce field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) SetNonce(v string) {
	o.Nonce = v
}

// GetSize returns the Size field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) SetVersion(v int32) {
	o.Version = v
}

func (o GetBlockDetailsByBlockHashFromCallbackRIBSZ2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2 struct {
	value *GetBlockDetailsByBlockHashFromCallbackRIBSZ2
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2) Get() *GetBlockDetailsByBlockHashFromCallbackRIBSZ2 {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2) Set(val *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2(val *GetBlockDetailsByBlockHashFromCallbackRIBSZ2) *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2 {
	return &NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


