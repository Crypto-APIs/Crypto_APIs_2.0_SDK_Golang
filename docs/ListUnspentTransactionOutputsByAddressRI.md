# ListUnspentTransactionOutputsByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address that has unspend funds per which the result is returned. | 
**Amount** | **string** | Represents the sent/received amount. | 
**Index** | **int32** | Represents the index position of the transaction in the block. | 
**IsConfirmed** | **bool** | Represents the state of the transaction whether it is confirmed or not confirmed. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 

## Methods

### NewListUnspentTransactionOutputsByAddressRI

`func NewListUnspentTransactionOutputsByAddressRI(address string, amount string, index int32, isConfirmed bool, timestamp int32, transactionId string, ) *ListUnspentTransactionOutputsByAddressRI`

NewListUnspentTransactionOutputsByAddressRI instantiates a new ListUnspentTransactionOutputsByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnspentTransactionOutputsByAddressRIWithDefaults

`func NewListUnspentTransactionOutputsByAddressRIWithDefaults() *ListUnspentTransactionOutputsByAddressRI`

NewListUnspentTransactionOutputsByAddressRIWithDefaults instantiates a new ListUnspentTransactionOutputsByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListUnspentTransactionOutputsByAddressRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListUnspentTransactionOutputsByAddressRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListUnspentTransactionOutputsByAddressRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListUnspentTransactionOutputsByAddressRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIndex

`func (o *ListUnspentTransactionOutputsByAddressRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListUnspentTransactionOutputsByAddressRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetIsConfirmed

`func (o *ListUnspentTransactionOutputsByAddressRI) GetIsConfirmed() bool`

GetIsConfirmed returns the IsConfirmed field if non-nil, zero value otherwise.

### GetIsConfirmedOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetIsConfirmedOk() (*bool, bool)`

GetIsConfirmedOk returns a tuple with the IsConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmed

`func (o *ListUnspentTransactionOutputsByAddressRI) SetIsConfirmed(v bool)`

SetIsConfirmed sets IsConfirmed field to given value.


### GetTimestamp

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListUnspentTransactionOutputsByAddressRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListUnspentTransactionOutputsByAddressRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


