# BroadcastLocallySignedTransactionRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 

## Methods

### NewBroadcastLocallySignedTransactionRI

`func NewBroadcastLocallySignedTransactionRI(transactionId string, ) *BroadcastLocallySignedTransactionRI`

NewBroadcastLocallySignedTransactionRI instantiates a new BroadcastLocallySignedTransactionRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastLocallySignedTransactionRIWithDefaults

`func NewBroadcastLocallySignedTransactionRIWithDefaults() *BroadcastLocallySignedTransactionRI`

NewBroadcastLocallySignedTransactionRIWithDefaults instantiates a new BroadcastLocallySignedTransactionRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *BroadcastLocallySignedTransactionRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BroadcastLocallySignedTransactionRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BroadcastLocallySignedTransactionRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


