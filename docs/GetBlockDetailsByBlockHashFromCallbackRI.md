# GetBlockDetailsByBlockHashFromCallbackRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in seconds. | 
**Nonce** | **string** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block. | 
**Uncles** | **[]string** |  | 
**BlockchainSpecific** | [**GetBlockDetailsByBlockHashFromCallbackRIBS**](GetBlockDetailsByBlockHashFromCallbackRIBS.md) |  | 

## Methods

### NewGetBlockDetailsByBlockHashFromCallbackRI

`func NewGetBlockDetailsByBlockHashFromCallbackRI(hash string, height int32, previousBlockHash string, timestamp int32, transactionsCount int32, difficulty string, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, nonce string, sha3Uncles string, size int32, totalDifficulty string, uncles []string, blockchainSpecific GetBlockDetailsByBlockHashFromCallbackRIBS, ) *GetBlockDetailsByBlockHashFromCallbackRI`

NewGetBlockDetailsByBlockHashFromCallbackRI instantiates a new GetBlockDetailsByBlockHashFromCallbackRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashFromCallbackRIWithDefaults

`func NewGetBlockDetailsByBlockHashFromCallbackRIWithDefaults() *GetBlockDetailsByBlockHashFromCallbackRI`

NewGetBlockDetailsByBlockHashFromCallbackRIWithDefaults instantiates a new GetBlockDetailsByBlockHashFromCallbackRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHeight

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetDifficulty

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetExtraData

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSha3Uncles

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotalDifficulty

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetUncles(v []string)`

SetUncles sets Uncles field to given value.


### GetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetBlockchainSpecific() GetBlockDetailsByBlockHashFromCallbackRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) GetBlockchainSpecificOk() (*GetBlockDetailsByBlockHashFromCallbackRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHashFromCallbackRI) SetBlockchainSpecific(v GetBlockDetailsByBlockHashFromCallbackRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


