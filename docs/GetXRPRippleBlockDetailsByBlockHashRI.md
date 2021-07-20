# GetXRPRippleBlockDetailsByBlockHashRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**BlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**NextBlockHash** | **string** | Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TotalCoins** | [**GetXRPRippleBlockDetailsByBlockHashRITotalCoins**](GetXRPRippleBlockDetailsByBlockHashRITotalCoins.md) |  | 
**TotalFees** | [**GetXRPRippleBlockDetailsByBlockHeightRITotalFees**](GetXRPRippleBlockDetailsByBlockHeightRITotalFees.md) |  | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 

## Methods

### NewGetXRPRippleBlockDetailsByBlockHashRI

`func NewGetXRPRippleBlockDetailsByBlockHashRI(blockHash string, blockHeight int32, nextBlockHash string, previousBlockHash string, timestamp int32, totalCoins GetXRPRippleBlockDetailsByBlockHashRITotalCoins, totalFees GetXRPRippleBlockDetailsByBlockHeightRITotalFees, transactionsCount int32, ) *GetXRPRippleBlockDetailsByBlockHashRI`

NewGetXRPRippleBlockDetailsByBlockHashRI instantiates a new GetXRPRippleBlockDetailsByBlockHashRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXRPRippleBlockDetailsByBlockHashRIWithDefaults

`func NewGetXRPRippleBlockDetailsByBlockHashRIWithDefaults() *GetXRPRippleBlockDetailsByBlockHashRI`

NewGetXRPRippleBlockDetailsByBlockHashRIWithDefaults instantiates a new GetXRPRippleBlockDetailsByBlockHashRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetNextBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetNextBlockHash() string`

GetNextBlockHash returns the NextBlockHash field if non-nil, zero value otherwise.

### GetNextBlockHashOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetNextBlockHashOk() (*string, bool)`

GetNextBlockHashOk returns a tuple with the NextBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetNextBlockHash(v string)`

SetNextBlockHash sets NextBlockHash field to given value.


### GetPreviousBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTotalCoins

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTotalCoins() GetXRPRippleBlockDetailsByBlockHashRITotalCoins`

GetTotalCoins returns the TotalCoins field if non-nil, zero value otherwise.

### GetTotalCoinsOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTotalCoinsOk() (*GetXRPRippleBlockDetailsByBlockHashRITotalCoins, bool)`

GetTotalCoinsOk returns a tuple with the TotalCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoins

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetTotalCoins(v GetXRPRippleBlockDetailsByBlockHashRITotalCoins)`

SetTotalCoins sets TotalCoins field to given value.


### GetTotalFees

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTotalFees() GetXRPRippleBlockDetailsByBlockHeightRITotalFees`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTotalFeesOk() (*GetXRPRippleBlockDetailsByBlockHeightRITotalFees, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetTotalFees(v GetXRPRippleBlockDetailsByBlockHeightRITotalFees)`

SetTotalFees sets TotalFees field to given value.


### GetTransactionsCount

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetXRPRippleBlockDetailsByBlockHashRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


