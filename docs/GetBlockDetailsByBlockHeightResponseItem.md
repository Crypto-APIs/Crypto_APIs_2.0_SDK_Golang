# GetBlockDetailsByBlockHeightResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**NextBlockHash** | **string** | Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**BlockchainSpecific** | [**GetBlockDetailsByBlockHeightResponseItemBlockchainSpecific**](GetBlockDetailsByBlockHeightResponseItemBlockchainSpecific.md) |  | 

## Methods

### NewGetBlockDetailsByBlockHeightResponseItem

`func NewGetBlockDetailsByBlockHeightResponseItem(hash string, height int32, nextBlockHash string, previousBlockHash string, timestamp int32, transactionsCount int32, blockchainSpecific GetBlockDetailsByBlockHeightResponseItemBlockchainSpecific, ) *GetBlockDetailsByBlockHeightResponseItem`

NewGetBlockDetailsByBlockHeightResponseItem instantiates a new GetBlockDetailsByBlockHeightResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHeightResponseItemWithDefaults

`func NewGetBlockDetailsByBlockHeightResponseItemWithDefaults() *GetBlockDetailsByBlockHeightResponseItem`

NewGetBlockDetailsByBlockHeightResponseItemWithDefaults instantiates a new GetBlockDetailsByBlockHeightResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHeight

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetNextBlockHash

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetNextBlockHash() string`

GetNextBlockHash returns the NextBlockHash field if non-nil, zero value otherwise.

### GetNextBlockHashOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetNextBlockHashOk() (*string, bool)`

GetNextBlockHashOk returns a tuple with the NextBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBlockHash

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetNextBlockHash(v string)`

SetNextBlockHash sets NextBlockHash field to given value.


### GetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetBlockchainSpecific() GetBlockDetailsByBlockHeightResponseItemBlockchainSpecific`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetBlockDetailsByBlockHeightResponseItem) GetBlockchainSpecificOk() (*GetBlockDetailsByBlockHeightResponseItemBlockchainSpecific, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHeightResponseItem) SetBlockchainSpecific(v GetBlockDetailsByBlockHeightResponseItemBlockchainSpecific)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


