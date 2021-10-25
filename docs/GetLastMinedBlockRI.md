# GetLastMinedBlockRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**Height** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**PreviousBlockHash** | **string** | Represents the hash of the previous block, also known as the parent block. | 
**Timestamp** | **int32** | Defines the exact date/time when this block was mined in Unix Timestamp. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 
**BlockchainSpecific** | [**GetLastMinedBlockRIBS**](GetLastMinedBlockRIBS.md) |  | 

## Methods

### NewGetLastMinedBlockRI

`func NewGetLastMinedBlockRI(hash string, height int32, previousBlockHash string, timestamp int32, transactionsCount int32, blockchainSpecific GetLastMinedBlockRIBS, ) *GetLastMinedBlockRI`

NewGetLastMinedBlockRI instantiates a new GetLastMinedBlockRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLastMinedBlockRIWithDefaults

`func NewGetLastMinedBlockRIWithDefaults() *GetLastMinedBlockRI`

NewGetLastMinedBlockRIWithDefaults instantiates a new GetLastMinedBlockRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetLastMinedBlockRI) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetLastMinedBlockRI) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetLastMinedBlockRI) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHeight

`func (o *GetLastMinedBlockRI) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *GetLastMinedBlockRI) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *GetLastMinedBlockRI) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetPreviousBlockHash

`func (o *GetLastMinedBlockRI) GetPreviousBlockHash() string`

GetPreviousBlockHash returns the PreviousBlockHash field if non-nil, zero value otherwise.

### GetPreviousBlockHashOk

`func (o *GetLastMinedBlockRI) GetPreviousBlockHashOk() (*string, bool)`

GetPreviousBlockHashOk returns a tuple with the PreviousBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBlockHash

`func (o *GetLastMinedBlockRI) SetPreviousBlockHash(v string)`

SetPreviousBlockHash sets PreviousBlockHash field to given value.


### GetTimestamp

`func (o *GetLastMinedBlockRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetLastMinedBlockRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetLastMinedBlockRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionsCount

`func (o *GetLastMinedBlockRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetLastMinedBlockRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetLastMinedBlockRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetBlockchainSpecific

`func (o *GetLastMinedBlockRI) GetBlockchainSpecific() GetLastMinedBlockRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetLastMinedBlockRI) GetBlockchainSpecificOk() (*GetLastMinedBlockRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetLastMinedBlockRI) SetBlockchainSpecific(v GetLastMinedBlockRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


