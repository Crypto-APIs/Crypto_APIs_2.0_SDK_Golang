# ListUnconfirmedTransactionsByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | [**[]ListUnconfirmedTransactionsByAddressRIRecipientsInner**](ListUnconfirmedTransactionsByAddressRIRecipientsInner.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]ListUnconfirmedTransactionsByAddressRISendersInner**](ListUnconfirmedTransactionsByAddressRISendersInner.md) | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**BlockchainSpecific** | [**ListUnconfirmedTransactionsByAddressRIBS**](ListUnconfirmedTransactionsByAddressRIBS.md) |  | 

## Methods

### NewListUnconfirmedTransactionsByAddressRI

`func NewListUnconfirmedTransactionsByAddressRI(recipients []ListUnconfirmedTransactionsByAddressRIRecipientsInner, senders []ListUnconfirmedTransactionsByAddressRISendersInner, timestamp int32, transactionHash string, transactionId string, blockchainSpecific ListUnconfirmedTransactionsByAddressRIBS, ) *ListUnconfirmedTransactionsByAddressRI`

NewListUnconfirmedTransactionsByAddressRI instantiates a new ListUnconfirmedTransactionsByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIWithDefaults() *ListUnconfirmedTransactionsByAddressRI`

NewListUnconfirmedTransactionsByAddressRIWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *ListUnconfirmedTransactionsByAddressRI) GetRecipients() []ListUnconfirmedTransactionsByAddressRIRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListUnconfirmedTransactionsByAddressRI) GetRecipientsOk() (*[]ListUnconfirmedTransactionsByAddressRIRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListUnconfirmedTransactionsByAddressRI) SetRecipients(v []ListUnconfirmedTransactionsByAddressRIRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListUnconfirmedTransactionsByAddressRI) GetSenders() []ListUnconfirmedTransactionsByAddressRISendersInner`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListUnconfirmedTransactionsByAddressRI) GetSendersOk() (*[]ListUnconfirmedTransactionsByAddressRISendersInner, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListUnconfirmedTransactionsByAddressRI) SetSenders(v []ListUnconfirmedTransactionsByAddressRISendersInner)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *ListUnconfirmedTransactionsByAddressRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListUnconfirmedTransactionsByAddressRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListUnconfirmedTransactionsByAddressRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListUnconfirmedTransactionsByAddressRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListUnconfirmedTransactionsByAddressRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListUnconfirmedTransactionsByAddressRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionId

`func (o *ListUnconfirmedTransactionsByAddressRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListUnconfirmedTransactionsByAddressRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListUnconfirmedTransactionsByAddressRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetBlockchainSpecific

`func (o *ListUnconfirmedTransactionsByAddressRI) GetBlockchainSpecific() ListUnconfirmedTransactionsByAddressRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *ListUnconfirmedTransactionsByAddressRI) GetBlockchainSpecificOk() (*ListUnconfirmedTransactionsByAddressRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *ListUnconfirmedTransactionsByAddressRI) SetBlockchainSpecific(v ListUnconfirmedTransactionsByAddressRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


