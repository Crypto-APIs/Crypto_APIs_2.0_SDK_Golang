# GetZilliqaBlockDetailsByBlockHeightRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Difficulty** | **string** | Defines how difficult it is for a specific miner to mine the block. | 
**DsBlock** | **int32** | Represents the Directory Service block which contains metadata about the miners who participate in the consensus protocol. | 
**DsDifficulty** | **string** | Defines how difficult it is to mine the dsBlocks. | 
**DsLeader** | **string** | Represents a part of the DS Committee which leads the consensus protocol for the epoch. | 
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**MicroBlocks** | **[]string** |  | 
**NextBlockHash** | **string** | Defines the hash of the next block from the specific blockchain. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 

## Methods

### NewGetZilliqaBlockDetailsByBlockHeightRI

`func NewGetZilliqaBlockDetailsByBlockHeightRI(blockHash string, difficulty string, dsBlock int32, dsDifficulty string, dsLeader string, gasLimit int32, gasUsed int32, microBlocks []string, nextBlockHash string, previousBlockHash string, timestamp int32, transactionsCount int32, ) *GetZilliqaBlockDetailsByBlockHeightRI`

NewGetZilliqaBlockDetailsByBlockHeightRI instantiates a new GetZilliqaBlockDetailsByBlockHeightRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZilliqaBlockDetailsByBlockHeightRIWithDefaults

`func NewGetZilliqaBlockDetailsByBlockHeightRIWithDefaults() *GetZilliqaBlockDetailsByBlockHeightRI`

NewGetZilliqaBlockDetailsByBlockHeightRIWithDefaults instantiates a new GetZilliqaBlockDetailsByBlockHeightRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetDifficulty

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetDsBlock

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDsBlock() int32`

GetDsBlock returns the DsBlock field if non-nil, zero value otherwise.

### GetDsBlockOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDsBlockOk() (*int32, bool)`

GetDsBlockOk returns a tuple with the DsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsBlock

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetDsBlock(v int32)`

SetDsBlock sets DsBlock field to given value.


### GetDsDifficulty

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDsDifficulty() string`

GetDsDifficulty returns the DsDifficulty field if non-nil, zero value otherwise.

### GetDsDifficultyOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDsDifficultyOk() (*string, bool)`

GetDsDifficultyOk returns a tuple with the DsDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsDifficulty

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetDsDifficulty(v string)`

SetDsDifficulty sets DsDifficulty field to given value.


### GetDsLeader

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDsLeader() string`

GetDsLeader returns the DsLeader field if non-nil, zero value otherwise.

### GetDsLeaderOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetDsLeaderOk() (*string, bool)`

GetDsLeaderOk returns a tuple with the DsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsLeader

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetDsLeader(v string)`

SetDsLeader sets DsLeader field to given value.


### GetGasLimit

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetMicroBlocks

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetMicroBlocks() []string`

GetMicroBlocks returns the MicroBlocks field if non-nil, zero value otherwise.

### GetMicroBlocksOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetMicroBlocksOk() (*[]string, bool)`

GetMicroBlocksOk returns a tuple with the MicroBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroBlocks

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetMicroBlocks(v []string)`

SetMicroBlocks sets MicroBlocks field to given value.


### GetNextBlockHash

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetNextBlockHash() string`

GetNextBlockHash returns the NextBlockHash field if non-nil, zero value otherwise.

### GetNextBlockHashOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetNextBlockHashOk() (*string, bool)`

GetNextBlockHashOk returns a tuple with the NextBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBlockHash

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetNextBlockHash(v string)`

SetNextBlockHash sets NextBlockHash field to given value.


### GetPreviousBlockHash

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetZilliqaBlockDetailsByBlockHeightRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


