# ListXRPRippleTransactionsByAddressResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | **string** | Represents any additional data that may be needed. | 
**Index** | **int32** | Represents the index position of the transaction in the block. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders**](GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders.md) | Represents an object of addresses that provide the funds. | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | **string** | Defines the status of the transaction. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the hash of the XRP transaction. | 
**Type** | **string** | Specifies the type of the transaction. | 
**Fee** | [**ListXRPRippleTransactionsByAddressResponseItemFee**](ListXRPRippleTransactionsByAddressResponseItemFee.md) |  | 
**Offer** | [**ListXRPRippleTransactionsByAddressResponseItemOffer**](ListXRPRippleTransactionsByAddressResponseItemOffer.md) |  | 
**Receive** | [**ListXRPRippleTransactionsByAddressResponseItemReceive**](ListXRPRippleTransactionsByAddressResponseItemReceive.md) |  | 
**Value** | [**ListXRPRippleTransactionsByAddressResponseItemValue**](ListXRPRippleTransactionsByAddressResponseItemValue.md) |  | 

## Methods

### NewListXRPRippleTransactionsByAddressResponseItem

`func NewListXRPRippleTransactionsByAddressResponseItem(additionalData string, index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients, senders []GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders, sequence int32, status string, timestamp int32, transactionHash string, type_ string, fee ListXRPRippleTransactionsByAddressResponseItemFee, offer ListXRPRippleTransactionsByAddressResponseItemOffer, receive ListXRPRippleTransactionsByAddressResponseItemReceive, value ListXRPRippleTransactionsByAddressResponseItemValue, ) *ListXRPRippleTransactionsByAddressResponseItem`

NewListXRPRippleTransactionsByAddressResponseItem instantiates a new ListXRPRippleTransactionsByAddressResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListXRPRippleTransactionsByAddressResponseItemWithDefaults

`func NewListXRPRippleTransactionsByAddressResponseItemWithDefaults() *ListXRPRippleTransactionsByAddressResponseItem`

NewListXRPRippleTransactionsByAddressResponseItemWithDefaults instantiates a new ListXRPRippleTransactionsByAddressResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.


### GetIndex

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetRecipients() []GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetRecipientsOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetRecipients(v []GetXRPRippleTransactionDetailsByTransactionIDResponseItemRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetSenders() []GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetSendersOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetSenders(v []GetXRPRippleTransactionDetailsByTransactionIDResponseItemSenders)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetFee() ListXRPRippleTransactionsByAddressResponseItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetFeeOk() (*ListXRPRippleTransactionsByAddressResponseItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetFee(v ListXRPRippleTransactionsByAddressResponseItemFee)`

SetFee sets Fee field to given value.


### GetOffer

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetOffer() ListXRPRippleTransactionsByAddressResponseItemOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetOfferOk() (*ListXRPRippleTransactionsByAddressResponseItemOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetOffer(v ListXRPRippleTransactionsByAddressResponseItemOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetReceive() ListXRPRippleTransactionsByAddressResponseItemReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetReceiveOk() (*ListXRPRippleTransactionsByAddressResponseItemReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetReceive(v ListXRPRippleTransactionsByAddressResponseItemReceive)`

SetReceive sets Receive field to given value.


### GetValue

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetValue() ListXRPRippleTransactionsByAddressResponseItemValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListXRPRippleTransactionsByAddressResponseItem) GetValueOk() (*ListXRPRippleTransactionsByAddressResponseItemValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListXRPRippleTransactionsByAddressResponseItem) SetValue(v ListXRPRippleTransactionsByAddressResponseItemValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


