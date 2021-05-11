# AddressCoinsTransactionConfirmedDataItemMinedInBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **int32** | Defines the number of blocks in the blockchain preceding this specific block. | 
**Hash** | **string** | Represents the hash of the block&#39;s header, i.e. an output that has a fixed length. | 
**Timestamp** | **int32** | Defines the exact date/time when this transaction was mined in seconds since Unix Epoch time. | 

## Methods

### NewAddressCoinsTransactionConfirmedDataItemMinedInBlock

`func NewAddressCoinsTransactionConfirmedDataItemMinedInBlock(height int32, hash string, timestamp int32, ) *AddressCoinsTransactionConfirmedDataItemMinedInBlock`

NewAddressCoinsTransactionConfirmedDataItemMinedInBlock instantiates a new AddressCoinsTransactionConfirmedDataItemMinedInBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressCoinsTransactionConfirmedDataItemMinedInBlockWithDefaults

`func NewAddressCoinsTransactionConfirmedDataItemMinedInBlockWithDefaults() *AddressCoinsTransactionConfirmedDataItemMinedInBlock`

NewAddressCoinsTransactionConfirmedDataItemMinedInBlockWithDefaults instantiates a new AddressCoinsTransactionConfirmedDataItemMinedInBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetHash

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTimestamp

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AddressCoinsTransactionConfirmedDataItemMinedInBlock) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


