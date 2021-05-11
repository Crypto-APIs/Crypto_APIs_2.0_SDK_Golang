# GetXRPRippleTransactionDetailsByTransactionIDResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | **string** | Represents additional data that may be needed. | 
**Index** | **string** | Defines the index of the transaction, i.e. the consecutive place it takes in the blockchain. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **string** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Offer** | [**GetXRPRippleTransactionDetailsByTransactionIDResponseItemOffer**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemOffer.md) |  | 
**Receive** | [**GetXRPRippleTransactionDetailsByTransactionIDResponseItemReceive**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemReceive.md) |  | 
**Recipients** | [**[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders.md) | Represents an object of addresses that provide the funds. | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | Pointer to **string** | Defines the status of the transaction. | [optional] 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**Type** | **string** | Defines the type of the transaction. | 
**Fee** | [**GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee.md) |  | 
**Value** | [**GetXRPRippleTransactionDetailsByTransactionIDResponseItemValue**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemValue.md) |  | 

## Methods

### NewGetXRPRippleTransactionDetailsByTransactionIDResponseItem

`func NewGetXRPRippleTransactionDetailsByTransactionIDResponseItem(additionalData string, index string, minedInBlockHash string, minedInBlockHeight string, offer GetXRPRippleTransactionDetailsByTransactionIDResponseItemOffer, receive GetXRPRippleTransactionDetailsByTransactionIDResponseItemReceive, recipients []GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients, senders []GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders, sequence int32, timestamp int32, transactionHash string, type_ string, fee GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee, value GetXRPRippleTransactionDetailsByTransactionIDResponseItemValue, ) *GetXRPRippleTransactionDetailsByTransactionIDResponseItem`

NewGetXRPRippleTransactionDetailsByTransactionIDResponseItem instantiates a new GetXRPRippleTransactionDetailsByTransactionIDResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemWithDefaults

`func NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemWithDefaults() *GetXRPRippleTransactionDetailsByTransactionIDResponseItem`

NewGetXRPRippleTransactionDetailsByTransactionIDResponseItemWithDefaults instantiates a new GetXRPRippleTransactionDetailsByTransactionIDResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.


### GetIndex

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetMinedInBlockHeight() string`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetMinedInBlockHeightOk() (*string, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetMinedInBlockHeight(v string)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetOffer

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetOffer() GetXRPRippleTransactionDetailsByTransactionIDResponseItemOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetOfferOk() (*GetXRPRippleTransactionDetailsByTransactionIDResponseItemOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetOffer(v GetXRPRippleTransactionDetailsByTransactionIDResponseItemOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetReceive() GetXRPRippleTransactionDetailsByTransactionIDResponseItemReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetReceiveOk() (*GetXRPRippleTransactionDetailsByTransactionIDResponseItemReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetReceive(v GetXRPRippleTransactionDetailsByTransactionIDResponseItemReceive)`

SetReceive sets Receive field to given value.


### GetRecipients

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetRecipients() []GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetRecipientsOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetRecipients(v []GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetSenders() []GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetSendersOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetSenders(v []GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetFee() GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetFeeOk() (*GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetFee(v GetXRPRippleTransactionDetailsByTransactionIDResponseItemFee)`

SetFee sets Fee field to given value.


### GetValue

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetValue() GetXRPRippleTransactionDetailsByTransactionIDResponseItemValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) GetValueOk() (*GetXRPRippleTransactionDetailsByTransactionIDResponseItemValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetXRPRippleTransactionDetailsByTransactionIDResponseItem) SetValue(v GetXRPRippleTransactionDetailsByTransactionIDResponseItemValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


