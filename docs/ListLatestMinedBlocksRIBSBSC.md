# ListLatestMinedBlocksRIBSBSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Defines the total gas limit of all transactions in the block. | 
**GasUsed** | **string** | Represents the total amount of gas used by all transactions in this block. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in second | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block | 
**Uncles** | **[]string** |  | 

## Methods

### NewListLatestMinedBlocksRIBSBSC

`func NewListLatestMinedBlocksRIBSBSC(extraData string, gasLimit string, gasUsed string, minedInSeconds int32, sha3Uncles string, totalDifficulty string, uncles []string, ) *ListLatestMinedBlocksRIBSBSC`

NewListLatestMinedBlocksRIBSBSC instantiates a new ListLatestMinedBlocksRIBSBSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSBSCWithDefaults

`func NewListLatestMinedBlocksRIBSBSCWithDefaults() *ListLatestMinedBlocksRIBSBSC`

NewListLatestMinedBlocksRIBSBSCWithDefaults instantiates a new ListLatestMinedBlocksRIBSBSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraData

`func (o *ListLatestMinedBlocksRIBSBSC) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ListLatestMinedBlocksRIBSBSC) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *ListLatestMinedBlocksRIBSBSC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListLatestMinedBlocksRIBSBSC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *ListLatestMinedBlocksRIBSBSC) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListLatestMinedBlocksRIBSBSC) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *ListLatestMinedBlocksRIBSBSC) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *ListLatestMinedBlocksRIBSBSC) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetSha3Uncles

`func (o *ListLatestMinedBlocksRIBSBSC) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *ListLatestMinedBlocksRIBSBSC) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetTotalDifficulty

`func (o *ListLatestMinedBlocksRIBSBSC) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *ListLatestMinedBlocksRIBSBSC) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *ListLatestMinedBlocksRIBSBSC) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *ListLatestMinedBlocksRIBSBSC) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *ListLatestMinedBlocksRIBSBSC) SetUncles(v []string)`

SetUncles sets Uncles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

