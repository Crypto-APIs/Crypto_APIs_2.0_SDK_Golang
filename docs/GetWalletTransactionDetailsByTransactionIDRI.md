# GetWalletTransactionDetailsByTransactionIDRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Represents the index position of the transaction in the specific block. | 
**IsConfirmed** | **bool** | Represents the state of the transaction whether it is confirmed or not confirmed. | 
**MinedInBlockHash** | Pointer to **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | [optional] 
**MinedInBlockHeight** | Pointer to **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | [optional] 
**Recipients** | **string** | String representation of the transaction to address | 
**Senders** | **string** | String representation of the transaction from address | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Fee** | [**GetWalletTransactionDetailsByTransactionIDRIFee**](GetWalletTransactionDetailsByTransactionIDRIFee.md) |  | 
**BlockchainSpecific** | [**GetWalletTransactionDetailsByTransactionIDRIBS**](GetWalletTransactionDetailsByTransactionIDRIBS.md) |  | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRI

`func NewGetWalletTransactionDetailsByTransactionIDRI(index int32, isConfirmed bool, recipients string, senders string, timestamp int32, transactionHash string, transactionId string, fee GetWalletTransactionDetailsByTransactionIDRIFee, blockchainSpecific GetWalletTransactionDetailsByTransactionIDRIBS, ) *GetWalletTransactionDetailsByTransactionIDRI`

NewGetWalletTransactionDetailsByTransactionIDRI instantiates a new GetWalletTransactionDetailsByTransactionIDRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIWithDefaults() *GetWalletTransactionDetailsByTransactionIDRI`

NewGetWalletTransactionDetailsByTransactionIDRIWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetIsConfirmed

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetIsConfirmed() bool`

GetIsConfirmed returns the IsConfirmed field if non-nil, zero value otherwise.

### GetIsConfirmedOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetIsConfirmedOk() (*bool, bool)`

GetIsConfirmedOk returns a tuple with the IsConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmed

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetIsConfirmed(v bool)`

SetIsConfirmed sets IsConfirmed field to given value.


### GetMinedInBlockHash

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.

### HasMinedInBlockHash

`func (o *GetWalletTransactionDetailsByTransactionIDRI) HasMinedInBlockHash() bool`

HasMinedInBlockHash returns a boolean if a field has been set.

### GetMinedInBlockHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.

### HasMinedInBlockHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRI) HasMinedInBlockHeight() bool`

HasMinedInBlockHeight returns a boolean if a field has been set.

### GetRecipients

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetRecipients() string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetRecipientsOk() (*string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetRecipients(v string)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetSenders() string`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetSendersOk() (*string, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetSenders(v string)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionId

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetFee

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetFee() GetWalletTransactionDetailsByTransactionIDRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetFeeOk() (*GetWalletTransactionDetailsByTransactionIDRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetFee(v GetWalletTransactionDetailsByTransactionIDRIFee)`

SetFee sets Fee field to given value.


### GetBlockchainSpecific

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetBlockchainSpecific() GetWalletTransactionDetailsByTransactionIDRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetWalletTransactionDetailsByTransactionIDRI) GetBlockchainSpecificOk() (*GetWalletTransactionDetailsByTransactionIDRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetWalletTransactionDetailsByTransactionIDRI) SetBlockchainSpecific(v GetWalletTransactionDetailsByTransactionIDRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


