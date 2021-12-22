# GetTransactionDetailsByTransactionIDFromCallbackRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Represents the index position of the transaction in the specific block. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]GetTransactionDetailsByTransactionIDFromCallbackRIRecipients**](GetTransactionDetailsByTransactionIDFromCallbackRIRecipients.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]GetTransactionDetailsByTransactionIDFromCallbackRISenders**](GetTransactionDetailsByTransactionIDFromCallbackRISenders.md) | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as transactionId for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols hash is different from transactionId for SegWit transactions. | 
**Fee** | [**GetTransactionDetailsByTransactionIDFromCallbackRIFee**](GetTransactionDetailsByTransactionIDFromCallbackRIFee.md) |  | 
**IsConfirmed** | **bool** |  | 
**BlockchainSpecific** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBS**](GetTransactionDetailsByTransactionIDFromCallbackRIBS.md) |  | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRI

`func NewGetTransactionDetailsByTransactionIDFromCallbackRI(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []GetTransactionDetailsByTransactionIDFromCallbackRIRecipients, senders []GetTransactionDetailsByTransactionIDFromCallbackRISenders, timestamp int32, transactionHash string, fee GetTransactionDetailsByTransactionIDFromCallbackRIFee, isConfirmed bool, blockchainSpecific GetTransactionDetailsByTransactionIDFromCallbackRIBS, ) *GetTransactionDetailsByTransactionIDFromCallbackRI`

NewGetTransactionDetailsByTransactionIDFromCallbackRI instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRI`

NewGetTransactionDetailsByTransactionIDFromCallbackRIWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetRecipients() []GetTransactionDetailsByTransactionIDFromCallbackRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetRecipientsOk() (*[]GetTransactionDetailsByTransactionIDFromCallbackRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetRecipients(v []GetTransactionDetailsByTransactionIDFromCallbackRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetSenders() []GetTransactionDetailsByTransactionIDFromCallbackRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetSendersOk() (*[]GetTransactionDetailsByTransactionIDFromCallbackRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetSenders(v []GetTransactionDetailsByTransactionIDFromCallbackRISenders)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetFee

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetFee() GetTransactionDetailsByTransactionIDFromCallbackRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetFeeOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetFee(v GetTransactionDetailsByTransactionIDFromCallbackRIFee)`

SetFee sets Fee field to given value.


### GetIsConfirmed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetIsConfirmed() bool`

GetIsConfirmed returns the IsConfirmed field if non-nil, zero value otherwise.

### GetIsConfirmedOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetIsConfirmedOk() (*bool, bool)`

GetIsConfirmedOk returns a tuple with the IsConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetIsConfirmed(v bool)`

SetIsConfirmed sets IsConfirmed field to given value.


### GetBlockchainSpecific

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetBlockchainSpecific() GetTransactionDetailsByTransactionIDFromCallbackRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) GetBlockchainSpecificOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRI) SetBlockchainSpecific(v GetTransactionDetailsByTransactionIDFromCallbackRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


