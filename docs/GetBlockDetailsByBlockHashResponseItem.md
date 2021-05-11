# GetBlockDetailsByBlockHashResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**NextBlockHash** | **string** | Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**BlockchainSpecific** | [**GetBlockDetailsByBlockHashResponseItemBlockchainSpecific**](GetBlockDetailsByBlockHashResponseItemBlockchainSpecific.md) |  | 

## Methods

### NewGetBlockDetailsByBlockHashResponseItem

`func NewGetBlockDetailsByBlockHashResponseItem(hash string, height int32, nextBlockHash string, previousBlockHash string, timestamp int32, transactionsCount int32, blockchainSpecific GetBlockDetailsByBlockHashResponseItemBlockchainSpecific, ) *GetBlockDetailsByBlockHashResponseItem`

NewGetBlockDetailsByBlockHashResponseItem instantiates a new GetBlockDetailsByBlockHashResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHashResponseItemWithDefaults

`func NewGetBlockDetailsByBlockHashResponseItemWithDefaults() *GetBlockDetailsByBlockHashResponseItem`

NewGetBlockDetailsByBlockHashResponseItemWithDefaults instantiates a new GetBlockDetailsByBlockHashResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetBlockDetailsByBlockHashResponseItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetBlockDetailsByBlockHashResponseItem) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHeight

`func (o *GetBlockDetailsByBlockHashResponseItem) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetBlockDetailsByBlockHashResponseItem) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetNextBlockHash

`func (o *GetBlockDetailsByBlockHashResponseItem) GetNextBlockHash() string`

GetNextBlockHash returns the NextBlockHash field if non-nil, zero value otherwise.

### GetNextBlockHashOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetNextBlockHashOk() (*string, bool)`

GetNextBlockHashOk returns a tuple with the NextBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBlockHash

`func (o *GetBlockDetailsByBlockHashResponseItem) SetNextBlockHash(v string)`

SetNextBlockHash sets NextBlockHash field to given value.


### GetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHashResponseItem) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHashResponseItem) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetBlockDetailsByBlockHashResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBlockDetailsByBlockHashResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetBlockDetailsByBlockHashResponseItem) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetBlockDetailsByBlockHashResponseItem) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHashResponseItem) GetBlockchainSpecific() GetBlockDetailsByBlockHashResponseItemBlockchainSpecific`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetBlockDetailsByBlockHashResponseItem) GetBlockchainSpecificOk() (*GetBlockDetailsByBlockHashResponseItemBlockchainSpecific, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHashResponseItem) SetBlockchainSpecific(v GetBlockDetailsByBlockHashResponseItemBlockchainSpecific)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


