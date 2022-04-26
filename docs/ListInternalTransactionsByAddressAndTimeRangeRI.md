# ListInternalTransactionsByAddressAndTimeRangeRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the specific amount of the transaction. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**OperationID** | **string** | Represents the unique internal transaction ID in regards to the parent transaction (type trace address). | 
**OperationType** | **string** | Defines the call type of the internal transaction. | 
**ParentHash** | **string** | Defines the specific hash of the parent transaction. | 
**Recipient** | **string** | Represents the recipient address with the respective amount. | 
**Sender** | **string** | Represents the sender address with the respective amount. | 
**Timestamp** | **int32** |  | 

## Methods

### NewListInternalTransactionsByAddressAndTimeRangeRI

`func NewListInternalTransactionsByAddressAndTimeRangeRI(amount string, minedInBlockHash string, minedInBlockHeight int32, operationID string, operationType string, parentHash string, recipient string, sender string, timestamp int32, ) *ListInternalTransactionsByAddressAndTimeRangeRI`

NewListInternalTransactionsByAddressAndTimeRangeRI instantiates a new ListInternalTransactionsByAddressAndTimeRangeRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInternalTransactionsByAddressAndTimeRangeRIWithDefaults

`func NewListInternalTransactionsByAddressAndTimeRangeRIWithDefaults() *ListInternalTransactionsByAddressAndTimeRangeRI`

NewListInternalTransactionsByAddressAndTimeRangeRIWithDefaults instantiates a new ListInternalTransactionsByAddressAndTimeRangeRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMinedInBlockHash

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetOperationID

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetOperationID() string`

GetOperationID returns the OperationID field if non-nil, zero value otherwise.

### GetOperationIDOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetOperationIDOk() (*string, bool)`

GetOperationIDOk returns a tuple with the OperationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationID

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetOperationID(v string)`

SetOperationID sets OperationID field to given value.


### GetOperationType

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetParentHash

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### GetRecipient

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetSender(v string)`

SetSender sets Sender field to given value.


### GetTimestamp

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListInternalTransactionsByAddressAndTimeRangeRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


