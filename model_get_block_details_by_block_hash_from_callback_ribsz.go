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

// GetBlockDetailsByBlockHashFromCallbackRIBSZ Zilliqa
type GetBlockDetailsByBlockHashFromCallbackRIBSZ struct {
	// Defines how difficult it is for a specific miner to mine the block.
	Difficulty string `json:"difficulty"`
	// Represents the Directory Service block which contains metadata about the miners who participate in the consensus protocol.
	DsBlock int32 `json:"dsBlock"`
	// Defines how difficult it is to mine the dsBlocks.
	DsDifficulty string `json:"dsDifficulty"`
	// Represents a part of the DS Committee which leads the consensus protocol for the epoch.
	DsLeader string `json:"dsLeader"`
	// Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit.
	GasLimit int32 `json:"gasLimit"`
	// Defines how much of the gas for the block has been used.
	GasUsed int32 `json:"gasUsed"`
	MicroBlocks []string `json:"microBlocks"`
}

// NewGetBlockDetailsByBlockHashFromCallbackRIBSZ instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSZ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockDetailsByBlockHashFromCallbackRIBSZ(difficulty string, dsBlock int32, dsDifficulty string, dsLeader string, gasLimit int32, gasUsed int32, microBlocks []string) *GetBlockDetailsByBlockHashFromCallbackRIBSZ {
	this := GetBlockDetailsByBlockHashFromCallbackRIBSZ{}
	this.Difficulty = difficulty
	this.DsBlock = dsBlock
	this.DsDifficulty = dsDifficulty
	this.DsLeader = dsLeader
	this.GasLimit = gasLimit
	this.GasUsed = gasUsed
	this.MicroBlocks = microBlocks
	return &this
}

// NewGetBlockDetailsByBlockHashFromCallbackRIBSZWithDefaults instantiates a new GetBlockDetailsByBlockHashFromCallbackRIBSZ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockDetailsByBlockHashFromCallbackRIBSZWithDefaults() *GetBlockDetailsByBlockHashFromCallbackRIBSZ {
	this := GetBlockDetailsByBlockHashFromCallbackRIBSZ{}
	return &this
}

// GetDifficulty returns the Difficulty field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetDsBlock returns the DsBlock field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDsBlock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DsBlock
}

// GetDsBlockOk returns a tuple with the DsBlock field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDsBlockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DsBlock, true
}

// SetDsBlock sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetDsBlock(v int32) {
	o.DsBlock = v
}

// GetDsDifficulty returns the DsDifficulty field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDsDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DsDifficulty
}

// GetDsDifficultyOk returns a tuple with the DsDifficulty field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDsDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DsDifficulty, true
}

// SetDsDifficulty sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetDsDifficulty(v string) {
	o.DsDifficulty = v
}

// GetDsLeader returns the DsLeader field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDsLeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DsLeader
}

// GetDsLeaderOk returns a tuple with the DsLeader field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetDsLeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DsLeader, true
}

// SetDsLeader sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetDsLeader(v string) {
	o.DsLeader = v
}

// GetGasLimit returns the GasLimit field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetGasLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetGasLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetGasLimit(v int32) {
	o.GasLimit = v
}

// GetGasUsed returns the GasUsed field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetGasUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetGasUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetGasUsed(v int32) {
	o.GasUsed = v
}

// GetMicroBlocks returns the MicroBlocks field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetMicroBlocks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MicroBlocks
}

// GetMicroBlocksOk returns a tuple with the MicroBlocks field value
// and a boolean to check if the value has been set.
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) GetMicroBlocksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicroBlocks, true
}

// SetMicroBlocks sets field value
func (o *GetBlockDetailsByBlockHashFromCallbackRIBSZ) SetMicroBlocks(v []string) {
	o.MicroBlocks = v
}

func (o GetBlockDetailsByBlockHashFromCallbackRIBSZ) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["difficulty"] = o.Difficulty
	}
	if true {
		toSerialize["dsBlock"] = o.DsBlock
	}
	if true {
		toSerialize["dsDifficulty"] = o.DsDifficulty
	}
	if true {
		toSerialize["dsLeader"] = o.DsLeader
	}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if true {
		toSerialize["gasUsed"] = o.GasUsed
	}
	if true {
		toSerialize["microBlocks"] = o.MicroBlocks
	}
	return json.Marshal(toSerialize)
}

type NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ struct {
	value *GetBlockDetailsByBlockHashFromCallbackRIBSZ
	isSet bool
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ) Get() *GetBlockDetailsByBlockHashFromCallbackRIBSZ {
	return v.value
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ) Set(val *GetBlockDetailsByBlockHashFromCallbackRIBSZ) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockDetailsByBlockHashFromCallbackRIBSZ(val *GetBlockDetailsByBlockHashFromCallbackRIBSZ) *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ {
	return &NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ{value: val, isSet: true}
}

func (v NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockDetailsByBlockHashFromCallbackRIBSZ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


