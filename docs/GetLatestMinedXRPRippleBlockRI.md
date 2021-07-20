# GetLatestMinedXRPRippleBlockRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**BlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**TotalCoins** | [**GetLatestMinedXRPRippleBlockRITotalCoins**](GetLatestMinedXRPRippleBlockRITotalCoins.md) |  | 
**TotalFees** | [**GetLatestMinedXRPRippleBlockRITotalFees**](GetLatestMinedXRPRippleBlockRITotalFees.md) |  | 

## Methods

### NewGetLatestMinedXRPRippleBlockRI

`func NewGetLatestMinedXRPRippleBlockRI(blockHash string, blockHeight int32, previousBlockHash string, timestamp int32, transactionsCount int32, totalCoins GetLatestMinedXRPRippleBlockRITotalCoins, totalFees GetLatestMinedXRPRippleBlockRITotalFees, ) *GetLatestMinedXRPRippleBlockRI`

NewGetLatestMinedXRPRippleBlockRI instantiates a new GetLatestMinedXRPRippleBlockRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLatestMinedXRPRippleBlockRIWithDefaults

`func NewGetLatestMinedXRPRippleBlockRIWithDefaults() *GetLatestMinedXRPRippleBlockRI`

NewGetLatestMinedXRPRippleBlockRIWithDefaults instantiates a new GetLatestMinedXRPRippleBlockRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *GetLatestMinedXRPRippleBlockRI) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GetLatestMinedXRPRippleBlockRI) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *GetLatestMinedXRPRippleBlockRI) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetLatestMinedXRPRippleBlockRI) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetPreviousBlockHash

`func (o *GetLatestMinedXRPRippleBlockRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetLatestMinedXRPRippleBlockRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetLatestMinedXRPRippleBlockRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetLatestMinedXRPRippleBlockRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetLatestMinedXRPRippleBlockRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetLatestMinedXRPRippleBlockRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetTotalCoins

`func (o *GetLatestMinedXRPRippleBlockRI) GetTotalCoins() GetLatestMinedXRPRippleBlockRITotalCoins`

GetTotalCoins returns the TotalCoins field if non-nil, zero value otherwise.

### GetTotalCoinsOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetTotalCoinsOk() (*GetLatestMinedXRPRippleBlockRITotalCoins, bool)`

GetTotalCoinsOk returns a tuple with the TotalCoins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoins

`func (o *GetLatestMinedXRPRippleBlockRI) SetTotalCoins(v GetLatestMinedXRPRippleBlockRITotalCoins)`

SetTotalCoins sets TotalCoins field to given value.


### GetTotalFees

`func (o *GetLatestMinedXRPRippleBlockRI) GetTotalFees() GetLatestMinedXRPRippleBlockRITotalFees`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *GetLatestMinedXRPRippleBlockRI) GetTotalFeesOk() (*GetLatestMinedXRPRippleBlockRITotalFees, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *GetLatestMinedXRPRippleBlockRI) SetTotalFees(v GetLatestMinedXRPRippleBlockRITotalFees)`

SetTotalFees sets TotalFees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


