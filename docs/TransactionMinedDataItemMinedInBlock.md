# TransactionMinedDataItemMinedInBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **int32** | Defines the number of blocks in the blockchain preceding this specific block. | 
**Hash** | **string** | Represents the hash of the block&#39;s header, i.e. an output that has a fixed length. | 
**Timestamp** | **int32** | Defines the exact date/time when this transaction was mined in seconds since Unix Epoch time. | 

## Methods

### NewTransactionMinedDataItemMinedInBlock

`func NewTransactionMinedDataItemMinedInBlock(height int32, hash string, timestamp int32, ) *TransactionMinedDataItemMinedInBlock`

NewTransactionMinedDataItemMinedInBlock instantiates a new TransactionMinedDataItemMinedInBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMinedDataItemMinedInBlockWithDefaults

`func NewTransactionMinedDataItemMinedInBlockWithDefaults() *TransactionMinedDataItemMinedInBlock`

NewTransactionMinedDataItemMinedInBlockWithDefaults instantiates a new TransactionMinedDataItemMinedInBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *TransactionMinedDataItemMinedInBlock) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TransactionMinedDataItemMinedInBlock) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TransactionMinedDataItemMinedInBlock) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetHash

`func (o *TransactionMinedDataItemMinedInBlock) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TransactionMinedDataItemMinedInBlock) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TransactionMinedDataItemMinedInBlock) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTimestamp

`func (o *TransactionMinedDataItemMinedInBlock) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TransactionMinedDataItemMinedInBlock) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TransactionMinedDataItemMinedInBlock) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


