# GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difficulty** | **string** | Represents a mathematical value of how hard it is to find a valid hash for this block. | 
**Nonce** | **string** | Represents a random value that can be adjusted to satisfy the Proof of Work. | 
**Size** | **int32** | Represents the total size of the block in Bytes. | 
**ExtraData** | **string** | Represents any data that can be included by the miner in the block. | 
**GasLimit** | **string** | Defines the total gas limit of all transactions in the block. | 
**GasUsed** | **string** | Represents the total amount of gas used by all transactions in this block. | 
**MinedInSeconds** | **int32** | Specifies the amount of time required for the block to be mined in seconds. | 
**Sha3Uncles** | **string** | Defines the combined hash of all uncles for a given parent. | 
**TotalDifficulty** | **string** | Defines the total difficulty of the chain until this block, i.e. how difficult it is for a specific miner to mine a new block. | 
**Uncles** | **[]string** |  | 

## Methods

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic(difficulty string, nonce string, size int32, extraData string, gasLimit string, gasUsed string, minedInSeconds int32, sha3Uncles string, totalDifficulty string, uncles []string, ) *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassicWithDefaults

`func NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassicWithDefaults() *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic`

NewGetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassicWithDefaults instantiates a new GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetSize(v int32)`

SetSize sets Size field to given value.


### GetExtraData

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetMinedInSeconds

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetMinedInSeconds() int32`

GetMinedInSeconds returns the MinedInSeconds field if non-nil, zero value otherwise.

### GetMinedInSecondsOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetMinedInSecondsOk() (*int32, bool)`

GetMinedInSecondsOk returns a tuple with the MinedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInSeconds

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetMinedInSeconds(v int32)`

SetMinedInSeconds sets MinedInSeconds field to given value.


### GetSha3Uncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetTotalDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetUncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *GetBlockDetailsByBlockHashResponseItemBlockchainSpecificEthereumClassic) SetUncles(v []string)`

SetUncles sets Uncles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


