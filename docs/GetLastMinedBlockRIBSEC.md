# GetLastMinedBlockRIBSEC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the proof of work | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Defines the total gas limit of all transactions in the block. | 
**GasUsed** | **string** | Represents the total amount of gas used by all transactions in this block. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in seconds. | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block. | 
**Uncles** | **[]string** |  | 

## Methods

### NewGetLastMinedBlockRIBSEC

`func NewGetLastMinedBlockRIBSEC(difficulty string, nonce string, size int32, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, sha3Uncles string, totalDifficulty string, uncles []string, ) *GetLastMinedBlockRIBSEC`

NewGetLastMinedBlockRIBSEC instantiates a new GetLastMinedBlockRIBSEC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLastMinedBlockRIBSECWithDefaults

`func NewGetLastMinedBlockRIBSECWithDefaults() *GetLastMinedBlockRIBSEC`

NewGetLastMinedBlockRIBSECWithDefaults instantiates a new GetLastMinedBlockRIBSEC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *GetLastMinedBlockRIBSEC) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetLastMinedBlockRIBSEC) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetLastMinedBlockRIBSEC) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetNonce

`func (o *GetLastMinedBlockRIBSEC) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetLastMinedBlockRIBSEC) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetLastMinedBlockRIBSEC) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetLastMinedBlockRIBSEC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetLastMinedBlockRIBSEC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetLastMinedBlockRIBSEC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetExtraData

`func (o *GetLastMinedBlockRIBSEC) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *GetLastMinedBlockRIBSEC) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *GetLastMinedBlockRIBSEC) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *GetLastMinedBlockRIBSEC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetLastMinedBlockRIBSEC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetLastMinedBlockRIBSEC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *GetLastMinedBlockRIBSEC) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetLastMinedBlockRIBSEC) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetLastMinedBlockRIBSEC) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *GetLastMinedBlockRIBSEC) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *GetLastMinedBlockRIBSEC) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *GetLastMinedBlockRIBSEC) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetSha3Uncles

`func (o *GetLastMinedBlockRIBSEC) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *GetLastMinedBlockRIBSEC) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *GetLastMinedBlockRIBSEC) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetTotalDifficulty

`func (o *GetLastMinedBlockRIBSEC) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *GetLastMinedBlockRIBSEC) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *GetLastMinedBlockRIBSEC) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *GetLastMinedBlockRIBSEC) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *GetLastMinedBlockRIBSEC) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *GetLastMinedBlockRIBSEC) SetUncles(v []string)`

SetUncles sets Uncles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


