# ListLatestMinedBlocksRIBSZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DsBlock** | **int32** | Represents the Directory Service block which contains metadata about the miners who participate in the consensus protocol. | 
**DsDifficulty** | **string** | Defines how difficult it is to mine the dsBlocks. | 
**DsLeader** | **string** | Represents a part of the DS Committee which leads the consensus protocol for the epoch. | 
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**MicroBlocks** | **[]string** |  | 

## Methods

### NewListLatestMinedBlocksRIBSZ

`func NewListLatestMinedBlocksRIBSZ(dsBlock int32, dsDifficulty string, dsLeader string, gasLimit int32, gasUsed int32, microBlocks []string, ) *ListLatestMinedBlocksRIBSZ`

NewListLatestMinedBlocksRIBSZ instantiates a new ListLatestMinedBlocksRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIBSZWithDefaults

`func NewListLatestMinedBlocksRIBSZWithDefaults() *ListLatestMinedBlocksRIBSZ`

NewListLatestMinedBlocksRIBSZWithDefaults instantiates a new ListLatestMinedBlocksRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDsBlock

`func (o *ListLatestMinedBlocksRIBSZ) GetDsBlock() int32`

GetDsBlock returns the DsBlock field if non-nil, zero value otherwise.

### GetDsBlockOk

`func (o *ListLatestMinedBlocksRIBSZ) GetDsBlockOk() (*int32, bool)`

GetDsBlockOk returns a tuple with the DsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsBlock

`func (o *ListLatestMinedBlocksRIBSZ) SetDsBlock(v int32)`

SetDsBlock sets DsBlock field to given value.


### GetDsDifficulty

`func (o *ListLatestMinedBlocksRIBSZ) GetDsDifficulty() string`

GetDsDifficulty returns the DsDifficulty field if non-nil, zero value otherwise.

### GetDsDifficultyOk

`func (o *ListLatestMinedBlocksRIBSZ) GetDsDifficultyOk() (*string, bool)`

GetDsDifficultyOk returns a tuple with the DsDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsDifficulty

`func (o *ListLatestMinedBlocksRIBSZ) SetDsDifficulty(v string)`

SetDsDifficulty sets DsDifficulty field to given value.


### GetDsLeader

`func (o *ListLatestMinedBlocksRIBSZ) GetDsLeader() string`

GetDsLeader returns the DsLeader field if non-nil, zero value otherwise.

### GetDsLeaderOk

`func (o *ListLatestMinedBlocksRIBSZ) GetDsLeaderOk() (*string, bool)`

GetDsLeaderOk returns a tuple with the DsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsLeader

`func (o *ListLatestMinedBlocksRIBSZ) SetDsLeader(v string)`

SetDsLeader sets DsLeader field to given value.


### GetGasLimit

`func (o *ListLatestMinedBlocksRIBSZ) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListLatestMinedBlocksRIBSZ) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListLatestMinedBlocksRIBSZ) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *ListLatestMinedBlocksRIBSZ) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListLatestMinedBlocksRIBSZ) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListLatestMinedBlocksRIBSZ) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetMicroBlocks

`func (o *ListLatestMinedBlocksRIBSZ) GetMicroBlocks() []string`

GetMicroBlocks returns the MicroBlocks field if non-nil, zero value otherwise.

### GetMicroBlocksOk

`func (o *ListLatestMinedBlocksRIBSZ) GetMicroBlocksOk() (*[]string, bool)`

GetMicroBlocksOk returns a tuple with the MicroBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroBlocks

`func (o *ListLatestMinedBlocksRIBSZ) SetMicroBlocks(v []string)`

SetMicroBlocks sets MicroBlocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


