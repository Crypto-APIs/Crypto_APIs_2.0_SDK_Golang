# ListInternalTransactionDetailsByTransactionHashRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the specific amount of the transaction. | 
**BlockHash** | **string** | Represents the hash of the block, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**BlockHeight** | **int32** | Represents the number of blocks in the blockchain preceding this specific block. Block numbers have no gaps. A blockchain usually starts with block 0 called the \&quot;Genesis block\&quot;. | 
**OperationID** | **string** | Represents the unique internal transaction ID in regards to the parent transaction (type trace address). | 
**OperationType** | **string** | Defines the call type of the internal transaction. | 
**ParentHash** | **string** | Defines the specific hash of the parent transaction. | 
**Recipient** | **string** | Represents the recipient address with the respective amount. | 
**Sender** | **string** | Represents the sender address with the respective amount. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 

## Methods

### NewListInternalTransactionDetailsByTransactionHashRI

`func NewListInternalTransactionDetailsByTransactionHashRI(amount string, blockHash string, blockHeight int32, operationID string, operationType string, parentHash string, recipient string, sender string, timestamp int32, ) *ListInternalTransactionDetailsByTransactionHashRI`

NewListInternalTransactionDetailsByTransactionHashRI instantiates a new ListInternalTransactionDetailsByTransactionHashRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInternalTransactionDetailsByTransactionHashRIWithDefaults

`func NewListInternalTransactionDetailsByTransactionHashRIWithDefaults() *ListInternalTransactionDetailsByTransactionHashRI`

NewListInternalTransactionDetailsByTransactionHashRIWithDefaults instantiates a new ListInternalTransactionDetailsByTransactionHashRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBlockHash

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockHeight

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.


### GetOperationID

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetOperationID() string`

GetOperationID returns the OperationID field if non-nil, zero value otherwise.

### GetOperationIDOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetOperationIDOk() (*string, bool)`

GetOperationIDOk returns a tuple with the OperationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationID

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetOperationID(v string)`

SetOperationID sets OperationID field to given value.


### GetOperationType

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetParentHash

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### GetRecipient

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetSender(v string)`

SetSender sets Sender field to given value.


### GetTimestamp

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListInternalTransactionDetailsByTransactionHashRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListInternalTransactionDetailsByTransactionHashRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


