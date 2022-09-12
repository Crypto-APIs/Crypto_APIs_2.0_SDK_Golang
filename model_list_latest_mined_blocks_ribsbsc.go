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

// ListLatestMinedBlocksRIBSBSC Binance Smart Chain
type ListLatestMinedBlocksRIBSBSC struct {
	// Represents a mathematical value of how hard it is to find a valid hash for this block.
	Difficulty string `json:"difficulty"`
	// Represents any data that can be included by the miner in the block.
	ExtraData string `json:"extraData"`
	// Defines the total gas limit of all transactions in the block.
	GasLimit string `json:"gasLimit"`
	// Represents the total amount of gas used by all transactions in this block.
	GasUsed string `json:"gasUsed"`
	// Specifies the amount of time required for the block to be mined in second
	MinedInSeconds int32 `json:"minedInSeconds"`
	// Represents a random value that can be adjusted to satisfy the proof of work
	Nonce string `json:"nonce"`
	// Defines the combined hash of all uncles for a given parent.
	Sha3Uncles string `json:"sha3Uncles"`
	// Represents the total size of the block in Bytes.
	Size int32 `json:"size"`
	// Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block
	TotalDifficulty string `json:"totalDifficulty"`
	Uncles []string `json:"uncles"`
}

// NewListLatestMinedBlocksRIBSBSC instantiates a new ListLatestMinedBlocksRIBSBSC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLatestMinedBlocksRIBSBSC(difficulty string, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, nonce string, sha3Uncles string, size int32, totalDifficulty string, uncles []string) *ListLatestMinedBlocksRIBSBSC {
	this := ListLatestMinedBlocksRIBSBSC{}
	this.Difficulty = difficulty
	this.ExtraData = extraData
	this.GasLimit = gasLimit
	this.GasUsed = gasUsed
	this.MinedInSeconds = minedInSeconds
	this.Nonce = nonce
	this.Sha3Uncles = sha3Uncles
	this.Size = size
	this.TotalDifficulty = totalDifficulty
	this.Uncles = uncles
	return &this
}

// NewListLatestMinedBlocksRIBSBSCWithDefaults instantiates a new ListLatestMinedBlocksRIBSBSC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLatestMinedBlocksRIBSBSCWithDefaults() *ListLatestMinedBlocksRIBSBSC {
	this := ListLatestMinedBlocksRIBSBSC{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *ListLatestMinedBlocksRIBSBSC) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetExtraData returns the ExtraData field value
func (o *ListLatestMinedBlocksRIBSBSC) GetExtraData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetExtraDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraData, true
}

// SetExtraData sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetExtraData(v string) {
	o.ExtraData = v
}

// GetGasLimit returns the GasLimit field value
func (o *ListLatestMinedBlocksRIBSBSC) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasUsed returns the GasUsed field value
func (o *ListLatestMinedBlocksRIBSBSC) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetMinedInSeconds returns the MinedInSeconds field value
func (o *ListLatestMinedBlocksRIBSBSC) GetMinedInSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInSeconds
}

// GetMinedInSecondsOk returns a tuple with the MinedInSeconds field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetMinedInSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinedInSeconds, true
}

// SetMinedInSeconds sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetMinedInSeconds(v int32) {
	o.MinedInSeconds = v
}

// GetNonce returns the Nonce field value
func (o *ListLatestMinedBlocksRIBSBSC) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetNonce(v string) {
	o.Nonce = v
}

// GetSha3Uncles returns the Sha3Uncles field value
func (o *ListLatestMinedBlocksRIBSBSC) GetSha3Uncles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha3Uncles
}

// GetSha3UnclesOk returns a tuple with the Sha3Uncles field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetSha3UnclesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha3Uncles, true
}

// SetSha3Uncles sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetSha3Uncles(v string) {
	o.Sha3Uncles = v
}

// GetSize returns the Size field value
func (o *ListLatestMinedBlocksRIBSBSC) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetSize(v int32) {
	o.Size = v
}

// GetTotalDifficulty returns the TotalDifficulty field value
func (o *ListLatestMinedBlocksRIBSBSC) GetTotalDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalDifficulty
}

// GetTotalDifficultyOk returns a tuple with the TotalDifficulty field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetTotalDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDifficulty, true
}

// SetTotalDifficulty sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetTotalDifficulty(v string) {
	o.TotalDifficulty = v
}

// GetUncles returns the Uncles field value
func (o *ListLatestMinedBlocksRIBSBSC) GetUncles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Uncles
}

// GetUnclesOk returns a tuple with the Uncles field value
// and a boolean to check if the value has been set.
func (o *ListLatestMinedBlocksRIBSBSC) GetUnclesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uncles, true
}

// SetUncles sets field value
func (o *ListLatestMinedBlocksRIBSBSC) SetUncles(v []string) {
	o.Uncles = v
}

func (o ListLatestMinedBlocksRIBSBSC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["difficulty"] = o.Difficulty
	}
	if true {
		toSerialize["extraData"] = o.ExtraData
	}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if true {
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["minedInSeconds"] = o.MinedInSeconds
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["sha3Uncles"] = o.Sha3Uncles
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["totalDifficulty"] = o.TotalDifficulty
	}
	if true {
		toSerialize["uncles"] = o.Uncles
	}
	return json.Marshal(toSerialize)
}

type NullableListLatestMinedBlocksRIBSBSC struct {
	value *ListLatestMinedBlocksRIBSBSC
	isSet bool
}

func (v NullableListLatestMinedBlocksRIBSBSC) Get() *ListLatestMinedBlocksRIBSBSC {
	return v.value
}

func (v *NullableListLatestMinedBlocksRIBSBSC) Set(val *ListLatestMinedBlocksRIBSBSC) {
	v.value = val
	v.isSet = true
}

func (v NullableListLatestMinedBlocksRIBSBSC) IsSet() bool {
	return v.isSet
}

func (v *NullableListLatestMinedBlocksRIBSBSC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLatestMinedBlocksRIBSBSC(val *ListLatestMinedBlocksRIBSBSC) *NullableListLatestMinedBlocksRIBSBSC {
	return &NullableListLatestMinedBlocksRIBSBSC{value: val, isSet: true}
}

func (v NullableListLatestMinedBlocksRIBSBSC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLatestMinedBlocksRIBSBSC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


