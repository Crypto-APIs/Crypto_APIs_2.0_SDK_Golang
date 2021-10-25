# ListXRPRippleTransactionsByBlockHeightRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** |  | [optional] 
**DestinationTag** | Pointer to **int32** |  | [optional] 
**Index** | **int32** |  | 
**MinedInBlockHash** | **string** |  | 
**Recipients** | [**[]ListXRPRippleTransactionsByBlockHeightRIRecipients**](ListXRPRippleTransactionsByBlockHeightRIRecipients.md) | Object Array representation of transaction receivers | 
**Senders** | [**[]ListXRPRippleTransactionsByBlockHeightRISenders**](ListXRPRippleTransactionsByBlockHeightRISenders.md) | Object Array representation of transaction senders | 
**Sequence** | **int32** |  | 
**Status** | **string** |  | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** |  | 
**Type** | **string** |  | 
**Fee** | [**ListXRPRippleTransactionsByBlockHeightRIFee**](ListXRPRippleTransactionsByBlockHeightRIFee.md) |  | 
**Offer** | [**ListXRPRippleTransactionsByBlockHeightRIOffer**](ListXRPRippleTransactionsByBlockHeightRIOffer.md) |  | 
**Receive** | [**ListXRPRippleTransactionsByBlockHeightRIReceive**](ListXRPRippleTransactionsByBlockHeightRIReceive.md) |  | 
**Value** | [**ListXRPRippleTransactionsByBlockHeightRIValue**](ListXRPRippleTransactionsByBlockHeightRIValue.md) |  | 

## Methods

### NewListXRPRippleTransactionsByBlockHeightRI

`func NewListXRPRippleTransactionsByBlockHeightRI(index int32, minedInBlockHash string, recipients []ListXRPRippleTransactionsByBlockHeightRIRecipients, senders []ListXRPRippleTransactionsByBlockHeightRISenders, sequence int32, status string, timestamp int32, transactionHash string, type_ string, fee ListXRPRippleTransactionsByBlockHeightRIFee, offer ListXRPRippleTransactionsByBlockHeightRIOffer, receive ListXRPRippleTransactionsByBlockHeightRIReceive, value ListXRPRippleTransactionsByBlockHeightRIValue, ) *ListXRPRippleTransactionsByBlockHeightRI`

NewListXRPRippleTransactionsByBlockHeightRI instantiates a new ListXRPRippleTransactionsByBlockHeightRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListXRPRippleTransactionsByBlockHeightRIWithDefaults

`func NewListXRPRippleTransactionsByBlockHeightRIWithDefaults() *ListXRPRippleTransactionsByBlockHeightRI`

NewListXRPRippleTransactionsByBlockHeightRIWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHeightRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ListXRPRippleTransactionsByBlockHeightRI) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetDestinationTag

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetDestinationTag() int32`

GetDestinationTag returns the DestinationTag field if non-nil, zero value otherwise.

### GetDestinationTagOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetDestinationTagOk() (*int32, bool)`

GetDestinationTagOk returns a tuple with the DestinationTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTag

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetDestinationTag(v int32)`

SetDestinationTag sets DestinationTag field to given value.

### HasDestinationTag

`func (o *ListXRPRippleTransactionsByBlockHeightRI) HasDestinationTag() bool`

HasDestinationTag returns a boolean if a field has been set.

### GetIndex

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetRecipients

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetRecipients() []ListXRPRippleTransactionsByBlockHeightRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetRecipientsOk() (*[]ListXRPRippleTransactionsByBlockHeightRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetRecipients(v []ListXRPRippleTransactionsByBlockHeightRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSenders() []ListXRPRippleTransactionsByBlockHeightRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSendersOk() (*[]ListXRPRippleTransactionsByBlockHeightRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetSenders(v []ListXRPRippleTransactionsByBlockHeightRISenders)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetFee() ListXRPRippleTransactionsByBlockHeightRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetFeeOk() (*ListXRPRippleTransactionsByBlockHeightRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetFee(v ListXRPRippleTransactionsByBlockHeightRIFee)`

SetFee sets Fee field to given value.


### GetOffer

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetOffer() ListXRPRippleTransactionsByBlockHeightRIOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetOfferOk() (*ListXRPRippleTransactionsByBlockHeightRIOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetOffer(v ListXRPRippleTransactionsByBlockHeightRIOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetReceive() ListXRPRippleTransactionsByBlockHeightRIReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetReceiveOk() (*ListXRPRippleTransactionsByBlockHeightRIReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetReceive(v ListXRPRippleTransactionsByBlockHeightRIReceive)`

SetReceive sets Receive field to given value.


### GetValue

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetValue() ListXRPRippleTransactionsByBlockHeightRIValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListXRPRippleTransactionsByBlockHeightRI) GetValueOk() (*ListXRPRippleTransactionsByBlockHeightRIValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListXRPRippleTransactionsByBlockHeightRI) SetValue(v ListXRPRippleTransactionsByBlockHeightRIValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


