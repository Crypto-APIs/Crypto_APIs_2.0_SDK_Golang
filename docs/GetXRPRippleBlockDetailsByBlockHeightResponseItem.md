# GetXRPRippleBlockDetailsByBlockHeightResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**BlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**NextBlockHash** | **string** | Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TotalCoins** | [**GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalCoins**](GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalCoins.md) |  | 
**TotalFees** | [**GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees**](GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees.md) |  | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 

## Methods

### NewGetXRPRippleBlockDetailsByBlockHeightResponseItem

`func NewGetXRPRippleBlockDetailsByBlockHeightResponseItem(blockHash string, blockHeight int32, nextBlockHash string, previousBlockHash string, timestamp int32, totalCoins GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalCoins, totalFees GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees, transactionsCount int32, ) *GetXRPRippleBlockDetailsByBlockHeightResponseItem`

NewGetXRPRippleBlockDetailsByBlockHeightResponseItem instantiates a new GetXRPRippleBlockDetailsByBlockHeightResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXRPRippleBlockDetailsByBlockHeightResponseItemWithDefaults

`func NewGetXRPRippleBlockDetailsByBlockHeightResponseItemWithDefaults() *GetXRPRippleBlockDetailsByBlockHeightResponseItem`

NewGetXRPRippleBlockDetailsByBlockHeightResponseItemWithDefaults instantiates a new GetXRPRippleBlockDetailsByBlockHeightResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetNextBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetNextBlockHash() string`

GetNextBlockHash returns the NextBlockHash field if non-nil, zero value otherwise.

### GetNextBlockHashOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetNextBlockHashOk() (*string, bool)`

GetNextBlockHashOk returns a tuple with the NextBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetNextBlockHash(v string)`

SetNextBlockHash sets NextBlockHash field to given value.


### GetPreviousBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTotalCoins

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTotalCoins() GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalCoins`

GetTotalCoins returns the TotalCoins field if non-nil, zero value otherwise.

### GetTotalCoinsOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTotalCoinsOk() (*GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalCoins, bool)`

GetTotalCoinsOk returns a tuple with the TotalCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoins

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetTotalCoins(v GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalCoins)`

SetTotalCoins sets TotalCoins field to given value.


### GetTotalFees

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTotalFees() GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTotalFeesOk() (*GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetTotalFees(v GetXRPRippleBlockDetailsByBlockHeightResponseItemTotalFees)`

SetTotalFees sets TotalFees field to given value.


### GetTransactionsCount

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetXRPRippleBlockDetailsByBlockHeightResponseItem) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


