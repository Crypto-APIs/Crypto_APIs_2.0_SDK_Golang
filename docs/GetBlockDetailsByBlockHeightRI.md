# GetBlockDetailsByBlockHeightRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**NextBlockHash** | **string** | Represents the hash of the next block. When this is the last block of the blockchain this value will be an empty string. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**BlockchainSpecific** | [**GetBlockDetailsByBlockHeightRIBS**](GetBlockDetailsByBlockHeightRIBS.md) |  | 

## Methods

### NewGetBlockDetailsByBlockHeightRI

`func NewGetBlockDetailsByBlockHeightRI(hash string, height int32, nextBlockHash string, previousBlockHash string, timestamp int32, transactionsCount int32, blockchainSpecific GetBlockDetailsByBlockHeightRIBS, ) *GetBlockDetailsByBlockHeightRI`

NewGetBlockDetailsByBlockHeightRI instantiates a new GetBlockDetailsByBlockHeightRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockDetailsByBlockHeightRIWithDefaults

`func NewGetBlockDetailsByBlockHeightRIWithDefaults() *GetBlockDetailsByBlockHeightRI`

NewGetBlockDetailsByBlockHeightRIWithDefaults instantiates a new GetBlockDetailsByBlockHeightRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetBlockDetailsByBlockHeightRI) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetBlockDetailsByBlockHeightRI) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetBlockDetailsByBlockHeightRI) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHeight

`func (o *GetBlockDetailsByBlockHeightRI) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetBlockDetailsByBlockHeightRI) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetBlockDetailsByBlockHeightRI) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetNextBlockHash

`func (o *GetBlockDetailsByBlockHeightRI) GetNextBlockHash() string`

GetNextBlockHash returns the NextBlockHash field if non-nil, zero value otherwise.

### GetNextBlockHashOk

`func (o *GetBlockDetailsByBlockHeightRI) GetNextBlockHashOk() (*string, bool)`

GetNextBlockHashOk returns a tuple with the NextBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBlockHash

`func (o *GetBlockDetailsByBlockHeightRI) SetNextBlockHash(v string)`

SetNextBlockHash sets NextBlockHash field to given value.


### GetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHeightRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetBlockDetailsByBlockHeightRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetBlockDetailsByBlockHeightRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetBlockDetailsByBlockHeightRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetBlockDetailsByBlockHeightRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetBlockDetailsByBlockHeightRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetBlockDetailsByBlockHeightRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetBlockDetailsByBlockHeightRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetBlockDetailsByBlockHeightRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHeightRI) GetBlockchainSpecific() GetBlockDetailsByBlockHeightRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetBlockDetailsByBlockHeightRI) GetBlockchainSpecificOk() (*GetBlockDetailsByBlockHeightRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetBlockDetailsByBlockHeightRI) SetBlockchainSpecific(v GetBlockDetailsByBlockHeightRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


