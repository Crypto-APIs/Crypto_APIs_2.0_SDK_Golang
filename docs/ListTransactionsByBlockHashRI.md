# ListTransactionsByBlockHashRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Represents the index position of the transaction in the specific block. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]ListTransactionsByBlockHashRIRecipientsInner**](ListTransactionsByBlockHashRIRecipientsInner.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]ListTransactionsByBlockHashRISendersInner**](ListTransactionsByBlockHashRISendersInner.md) | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Fee** | [**ListTransactionsByBlockHashRIFee**](ListTransactionsByBlockHashRIFee.md) |  | 
**BlockchainSpecific** | [**ListTransactionsByBlockHashRIBS**](ListTransactionsByBlockHashRIBS.md) |  | 

## Methods

### NewListTransactionsByBlockHashRI

`func NewListTransactionsByBlockHashRI(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []ListTransactionsByBlockHashRIRecipientsInner, senders []ListTransactionsByBlockHashRISendersInner, timestamp int32, transactionHash string, transactionId string, fee ListTransactionsByBlockHashRIFee, blockchainSpecific ListTransactionsByBlockHashRIBS, ) *ListTransactionsByBlockHashRI`

NewListTransactionsByBlockHashRI instantiates a new ListTransactionsByBlockHashRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIWithDefaults

`func NewListTransactionsByBlockHashRIWithDefaults() *ListTransactionsByBlockHashRI`

NewListTransactionsByBlockHashRIWithDefaults instantiates a new ListTransactionsByBlockHashRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ListTransactionsByBlockHashRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListTransactionsByBlockHashRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListTransactionsByBlockHashRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListTransactionsByBlockHashRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListTransactionsByBlockHashRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListTransactionsByBlockHashRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *ListTransactionsByBlockHashRI) GetRecipients() []ListTransactionsByBlockHashRIRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListTransactionsByBlockHashRI) GetRecipientsOk() (*[]ListTransactionsByBlockHashRIRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListTransactionsByBlockHashRI) SetRecipients(v []ListTransactionsByBlockHashRIRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListTransactionsByBlockHashRI) GetSenders() []ListTransactionsByBlockHashRISendersInner`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListTransactionsByBlockHashRI) GetSendersOk() (*[]ListTransactionsByBlockHashRISendersInner, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListTransactionsByBlockHashRI) SetSenders(v []ListTransactionsByBlockHashRISendersInner)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *ListTransactionsByBlockHashRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListTransactionsByBlockHashRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListTransactionsByBlockHashRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListTransactionsByBlockHashRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListTransactionsByBlockHashRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListTransactionsByBlockHashRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionId

`func (o *ListTransactionsByBlockHashRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListTransactionsByBlockHashRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListTransactionsByBlockHashRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetFee

`func (o *ListTransactionsByBlockHashRI) GetFee() ListTransactionsByBlockHashRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListTransactionsByBlockHashRI) GetFeeOk() (*ListTransactionsByBlockHashRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListTransactionsByBlockHashRI) SetFee(v ListTransactionsByBlockHashRIFee)`

SetFee sets Fee field to given value.


### GetBlockchainSpecific

`func (o *ListTransactionsByBlockHashRI) GetBlockchainSpecific() ListTransactionsByBlockHashRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *ListTransactionsByBlockHashRI) GetBlockchainSpecificOk() (*ListTransactionsByBlockHashRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *ListTransactionsByBlockHashRI) SetBlockchainSpecific(v ListTransactionsByBlockHashRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


