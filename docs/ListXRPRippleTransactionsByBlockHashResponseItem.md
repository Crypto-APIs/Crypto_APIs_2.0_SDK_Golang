# ListXRPRippleTransactionsByBlockHashResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** | Represents any additional data that may be needed. | [optional] 
**Index** | **int32** | Represents the index position of the transaction in the specific block. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]ListXRPRippleTransactionsByBlockHashResponseItemRecipients**](ListXRPRippleTransactionsByBlockHashResponseItemRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]ListXRPRippleTransactionsByBlockHashResponseItemSenders**](ListXRPRippleTransactionsByBlockHashResponseItemSenders.md) | Represents an object of addresses that provide the funds. | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | **string** | Defines the status of the transaction. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**Type** | **string** | Defines the type of the transaction. | 
**Fee** | [**ListXRPRippleTransactionsByBlockHashResponseItemFee**](ListXRPRippleTransactionsByBlockHashResponseItemFee.md) |  | 
**Offer** | [**ListXRPRippleTransactionsByBlockHashResponseItemOffer**](ListXRPRippleTransactionsByBlockHashResponseItemOffer.md) |  | 
**Receive** | [**ListXRPRippleTransactionsByBlockHashResponseItemReceive**](ListXRPRippleTransactionsByBlockHashResponseItemReceive.md) |  | 
**Value** | [**ListXRPRippleTransactionsByBlockHashResponseItemValue**](ListXRPRippleTransactionsByBlockHashResponseItemValue.md) |  | 

## Methods

### NewListXRPRippleTransactionsByBlockHashResponseItem

`func NewListXRPRippleTransactionsByBlockHashResponseItem(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []ListXRPRippleTransactionsByBlockHashResponseItemRecipients, senders []ListXRPRippleTransactionsByBlockHashResponseItemSenders, sequence int32, status string, timestamp int32, transactionHash string, type_ string, fee ListXRPRippleTransactionsByBlockHashResponseItemFee, offer ListXRPRippleTransactionsByBlockHashResponseItemOffer, receive ListXRPRippleTransactionsByBlockHashResponseItemReceive, value ListXRPRippleTransactionsByBlockHashResponseItemValue, ) *ListXRPRippleTransactionsByBlockHashResponseItem`

NewListXRPRippleTransactionsByBlockHashResponseItem instantiates a new ListXRPRippleTransactionsByBlockHashResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListXRPRippleTransactionsByBlockHashResponseItemWithDefaults

`func NewListXRPRippleTransactionsByBlockHashResponseItemWithDefaults() *ListXRPRippleTransactionsByBlockHashResponseItem`

NewListXRPRippleTransactionsByBlockHashResponseItemWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHashResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetIndex

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetRecipients() []ListXRPRippleTransactionsByBlockHashResponseItemRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetRecipientsOk() (*[]ListXRPRippleTransactionsByBlockHashResponseItemRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetRecipients(v []ListXRPRippleTransactionsByBlockHashResponseItemRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetSenders() []ListXRPRippleTransactionsByBlockHashResponseItemSenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetSendersOk() (*[]ListXRPRippleTransactionsByBlockHashResponseItemSenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetSenders(v []ListXRPRippleTransactionsByBlockHashResponseItemSenders)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetFee() ListXRPRippleTransactionsByBlockHashResponseItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetFeeOk() (*ListXRPRippleTransactionsByBlockHashResponseItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetFee(v ListXRPRippleTransactionsByBlockHashResponseItemFee)`

SetFee sets Fee field to given value.


### GetOffer

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetOffer() ListXRPRippleTransactionsByBlockHashResponseItemOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetOfferOk() (*ListXRPRippleTransactionsByBlockHashResponseItemOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetOffer(v ListXRPRippleTransactionsByBlockHashResponseItemOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetReceive() ListXRPRippleTransactionsByBlockHashResponseItemReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetReceiveOk() (*ListXRPRippleTransactionsByBlockHashResponseItemReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetReceive(v ListXRPRippleTransactionsByBlockHashResponseItemReceive)`

SetReceive sets Receive field to given value.


### GetValue

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetValue() ListXRPRippleTransactionsByBlockHashResponseItemValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) GetValueOk() (*ListXRPRippleTransactionsByBlockHashResponseItemValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListXRPRippleTransactionsByBlockHashResponseItem) SetValue(v ListXRPRippleTransactionsByBlockHashResponseItemValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


