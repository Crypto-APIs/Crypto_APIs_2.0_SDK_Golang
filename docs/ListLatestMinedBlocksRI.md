# ListLatestMinedBlocksRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**TotalCoins** | [**GetLatestMinedXRPRippleBlockRITotalCoins**](GetLatestMinedXRPRippleBlockRITotalCoins.md) |  | 
**TotalFees** | [**GetLatestMinedXRPRippleBlockRITotalFees**](GetLatestMinedXRPRippleBlockRITotalFees.md) |  | 
**BlockchainSpecific** | [**ListLatestMinedBlocksRIBS**](ListLatestMinedBlocksRIBS.md) |  | 

## Methods

### NewListLatestMinedBlocksRI

`func NewListLatestMinedBlocksRI(hash string, height int32, previousBlockHash string, timestamp int32, transactionsCount int32, totalCoins GetLatestMinedXRPRippleBlockRITotalCoins, totalFees GetLatestMinedXRPRippleBlockRITotalFees, blockchainSpecific ListLatestMinedBlocksRIBS, ) *ListLatestMinedBlocksRI`

NewListLatestMinedBlocksRI instantiates a new ListLatestMinedBlocksRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLatestMinedBlocksRIWithDefaults

`func NewListLatestMinedBlocksRIWithDefaults() *ListLatestMinedBlocksRI`

NewListLatestMinedBlocksRIWithDefaults instantiates a new ListLatestMinedBlocksRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ListLatestMinedBlocksRI) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ListLatestMinedBlocksRI) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ListLatestMinedBlocksRI) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHeight

`func (o *ListLatestMinedBlocksRI) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ListLatestMinedBlocksRI) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ListLatestMinedBlocksRI) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetPreviousBlockHash

`func (o *ListLatestMinedBlocksRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *ListLatestMinedBlocksRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *ListLatestMinedBlocksRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *ListLatestMinedBlocksRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListLatestMinedBlocksRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListLatestMinedBlocksRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *ListLatestMinedBlocksRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *ListLatestMinedBlocksRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *ListLatestMinedBlocksRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetTotalCoins

`func (o *ListLatestMinedBlocksRI) GetTotalCoins() GetLatestMinedXRPRippleBlockRITotalCoins`

GetTotalCoins returns the TotalCoins field if non-nil, zero value otherwise.

### GetTotalCoinsOk

`func (o *ListLatestMinedBlocksRI) GetTotalCoinsOk() (*GetLatestMinedXRPRippleBlockRITotalCoins, bool)`

GetTotalCoinsOk returns a tuple with the TotalCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoins

`func (o *ListLatestMinedBlocksRI) SetTotalCoins(v GetLatestMinedXRPRippleBlockRITotalCoins)`

SetTotalCoins sets TotalCoins field to given value.


### GetTotalFees

`func (o *ListLatestMinedBlocksRI) GetTotalFees() GetLatestMinedXRPRippleBlockRITotalFees`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *ListLatestMinedBlocksRI) GetTotalFeesOk() (*GetLatestMinedXRPRippleBlockRITotalFees, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *ListLatestMinedBlocksRI) SetTotalFees(v GetLatestMinedXRPRippleBlockRITotalFees)`

SetTotalFees sets TotalFees field to given value.


### GetBlockchainSpecific

`func (o *ListLatestMinedBlocksRI) GetBlockchainSpecific() ListLatestMinedBlocksRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *ListLatestMinedBlocksRI) GetBlockchainSpecificOk() (*ListLatestMinedBlocksRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *ListLatestMinedBlocksRI) SetBlockchainSpecific(v ListLatestMinedBlocksRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


