# GetInternalTransactionByTransactionHashAndOperationIdRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the specific amount of the transaction. | 
**BlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**BlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**OperationType** | **string** | Defines the specific type of the operation. | 
**ParentHash** | **string** | Defines the specific hash of the parent transaction. | 
**Recipient** | **string** | Represents the recipient address with the respective amount. | 
**Sender** | **string** | Represents the sender address with the respective amount. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 

## Methods

### NewGetInternalTransactionByTransactionHashAndOperationIdRI

`func NewGetInternalTransactionByTransactionHashAndOperationIdRI(amount string, blockHash string, blockHeight int32, operationType string, parentHash string, recipient string, sender string, timestamp int32, ) *GetInternalTransactionByTransactionHashAndOperationIdRI`

NewGetInternalTransactionByTransactionHashAndOperationIdRI instantiates a new GetInternalTransactionByTransactionHashAndOperationIdRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInternalTransactionByTransactionHashAndOperationIdRIWithDefaults

`func NewGetInternalTransactionByTransactionHashAndOperationIdRIWithDefaults() *GetInternalTransactionByTransactionHashAndOperationIdRI`

NewGetInternalTransactionByTransactionHashAndOperationIdRIWithDefaults instantiates a new GetInternalTransactionByTransactionHashAndOperationIdRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBlockHash

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetOperationType

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetParentHash

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### GetRecipient

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetSender(v string)`

SetSender sets Sender field to given value.


### GetTimestamp

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetInternalTransactionByTransactionHashAndOperationIdRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


