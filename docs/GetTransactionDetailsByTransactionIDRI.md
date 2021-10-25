# GetTransactionDetailsByTransactionIDRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Represents the index position of the transaction in the specific block. | 
**IsConfirmed** | **string** | Represents the state of the transaction whether it is confirmed or not confirmed. | 
**MinedInBlockHash** | Pointer to **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | [optional] 
**MinedInBlockHeight** | Pointer to **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | [optional] 
**Recipients** | [**[]GetTransactionDetailsByTransactionIDRIRecipients**](GetTransactionDetailsByTransactionIDRIRecipients.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]GetTransactionDetailsByTransactionIDRISenders**](GetTransactionDetailsByTransactionIDRISenders.md) | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Fee** | [**GetTransactionDetailsByTransactionIDRIFee**](GetTransactionDetailsByTransactionIDRIFee.md) |  | 
**BlockchainSpecific** | [**GetTransactionDetailsByTransactionIDRIBS**](GetTransactionDetailsByTransactionIDRIBS.md) |  | 

## Methods

### NewGetTransactionDetailsByTransactionIDRI

`func NewGetTransactionDetailsByTransactionIDRI(index int32, isConfirmed string, recipients []GetTransactionDetailsByTransactionIDRIRecipients, senders []GetTransactionDetailsByTransactionIDRISenders, timestamp int32, transactionHash string, transactionId string, fee GetTransactionDetailsByTransactionIDRIFee, blockchainSpecific GetTransactionDetailsByTransactionIDRIBS, ) *GetTransactionDetailsByTransactionIDRI`

NewGetTransactionDetailsByTransactionIDRI instantiates a new GetTransactionDetailsByTransactionIDRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIWithDefaults() *GetTransactionDetailsByTransactionIDRI`

NewGetTransactionDetailsByTransactionIDRIWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *GetTransactionDetailsByTransactionIDRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetTransactionDetailsByTransactionIDRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetIsConfirmed

`func (o *GetTransactionDetailsByTransactionIDRI) GetIsConfirmed() string`

GetIsConfirmed returns the IsConfirmed field if non-nil, zero value otherwise.

### GetIsConfirmedOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetIsConfirmedOk() (*string, bool)`

GetIsConfirmedOk returns a tuple with the IsConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmed

`func (o *GetTransactionDetailsByTransactionIDRI) SetIsConfirmed(v string)`

SetIsConfirmed sets IsConfirmed field to given value.


### GetMinedInBlockHash

`func (o *GetTransactionDetailsByTransactionIDRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetTransactionDetailsByTransactionIDRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.

### HasMinedInBlockHash

`func (o *GetTransactionDetailsByTransactionIDRI) HasMinedInBlockHash() bool`

HasMinedInBlockHash returns a boolean if a field has been set.

### GetMinedInBlockHeight

`func (o *GetTransactionDetailsByTransactionIDRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetTransactionDetailsByTransactionIDRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.

### HasMinedInBlockHeight

`func (o *GetTransactionDetailsByTransactionIDRI) HasMinedInBlockHeight() bool`

HasMinedInBlockHeight returns a boolean if a field has been set.

### GetRecipients

`func (o *GetTransactionDetailsByTransactionIDRI) GetRecipients() []GetTransactionDetailsByTransactionIDRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetRecipientsOk() (*[]GetTransactionDetailsByTransactionIDRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetTransactionDetailsByTransactionIDRI) SetRecipients(v []GetTransactionDetailsByTransactionIDRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetTransactionDetailsByTransactionIDRI) GetSenders() []GetTransactionDetailsByTransactionIDRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetSendersOk() (*[]GetTransactionDetailsByTransactionIDRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetTransactionDetailsByTransactionIDRI) SetSenders(v []GetTransactionDetailsByTransactionIDRISenders)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *GetTransactionDetailsByTransactionIDRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetTransactionDetailsByTransactionIDRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *GetTransactionDetailsByTransactionIDRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetTransactionDetailsByTransactionIDRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionId

`func (o *GetTransactionDetailsByTransactionIDRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetTransactionDetailsByTransactionIDRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetFee

`func (o *GetTransactionDetailsByTransactionIDRI) GetFee() GetTransactionDetailsByTransactionIDRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetFeeOk() (*GetTransactionDetailsByTransactionIDRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTransactionDetailsByTransactionIDRI) SetFee(v GetTransactionDetailsByTransactionIDRIFee)`

SetFee sets Fee field to given value.


### GetBlockchainSpecific

`func (o *GetTransactionDetailsByTransactionIDRI) GetBlockchainSpecific() GetTransactionDetailsByTransactionIDRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetTransactionDetailsByTransactionIDRI) GetBlockchainSpecificOk() (*GetTransactionDetailsByTransactionIDRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetTransactionDetailsByTransactionIDRI) SetBlockchainSpecific(v GetTransactionDetailsByTransactionIDRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


