# ListXRPRippleTransactionsByAddressAndTimeRangeRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationTag** | Pointer to **int64** | A destination tag is a value used to discern the holder of the Ripple (XRP) being deposited or withdrawn. | [optional] 
**Index** | **int32** | Represents the index position of the transaction in the block. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner**](GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetXRPRippleTransactionDetailsByTransactionIDRISendersInner**](GetXRPRippleTransactionDetailsByTransactionIDRISendersInner.md) | Represents an object of addresses that provide the funds. | 
**Sequence** | **int64** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | **string** | Defines the status of the transaction. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the hash of the XRP transaction. | 
**Type** | **string** | Specifies the type of the transaction. | 
**Fee** | [**ListXRPRippleTransactionsByAddressRIFee**](ListXRPRippleTransactionsByAddressRIFee.md) |  | 
**Offer** | [**ListXRPRippleTransactionsByAddressRIOffer**](ListXRPRippleTransactionsByAddressRIOffer.md) |  | 
**Receive** | [**ListXRPRippleTransactionsByAddressRIReceive**](ListXRPRippleTransactionsByAddressRIReceive.md) |  | 
**Value** | [**ListXRPRippleTransactionsByAddressRIValue**](ListXRPRippleTransactionsByAddressRIValue.md) |  | 

## Methods

### NewListXRPRippleTransactionsByAddressAndTimeRangeRI

`func NewListXRPRippleTransactionsByAddressAndTimeRangeRI(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner, senders []GetXRPRippleTransactionDetailsByTransactionIDRISendersInner, sequence int64, status string, timestamp int32, transactionHash string, type_ string, fee ListXRPRippleTransactionsByAddressRIFee, offer ListXRPRippleTransactionsByAddressRIOffer, receive ListXRPRippleTransactionsByAddressRIReceive, value ListXRPRippleTransactionsByAddressRIValue, ) *ListXRPRippleTransactionsByAddressAndTimeRangeRI`

NewListXRPRippleTransactionsByAddressAndTimeRangeRI instantiates a new ListXRPRippleTransactionsByAddressAndTimeRangeRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListXRPRippleTransactionsByAddressAndTimeRangeRIWithDefaults

`func NewListXRPRippleTransactionsByAddressAndTimeRangeRIWithDefaults() *ListXRPRippleTransactionsByAddressAndTimeRangeRI`

NewListXRPRippleTransactionsByAddressAndTimeRangeRIWithDefaults instantiates a new ListXRPRippleTransactionsByAddressAndTimeRangeRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationTag

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetDestinationTag() int64`

GetDestinationTag returns the DestinationTag field if non-nil, zero value otherwise.

### GetDestinationTagOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetDestinationTagOk() (*int64, bool)`

GetDestinationTagOk returns a tuple with the DestinationTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTag

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetDestinationTag(v int64)`

SetDestinationTag sets DestinationTag field to given value.

### HasDestinationTag

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) HasDestinationTag() bool`

HasDestinationTag returns a boolean if a field has been set.

### GetIndex

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetRecipients() []GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetRecipientsOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetRecipients(v []GetXRPRippleTransactionDetailsByTransactionIDRIRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetSenders() []GetXRPRippleTransactionDetailsByTransactionIDRISendersInner`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetSendersOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDRISendersInner, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetSenders(v []GetXRPRippleTransactionDetailsByTransactionIDRISendersInner)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetFee() ListXRPRippleTransactionsByAddressRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetFeeOk() (*ListXRPRippleTransactionsByAddressRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetFee(v ListXRPRippleTransactionsByAddressRIFee)`

SetFee sets Fee field to given value.


### GetOffer

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetOffer() ListXRPRippleTransactionsByAddressRIOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetOfferOk() (*ListXRPRippleTransactionsByAddressRIOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetOffer(v ListXRPRippleTransactionsByAddressRIOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetReceive() ListXRPRippleTransactionsByAddressRIReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetReceiveOk() (*ListXRPRippleTransactionsByAddressRIReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetReceive(v ListXRPRippleTransactionsByAddressRIReceive)`

SetReceive sets Receive field to given value.


### GetValue

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetValue() ListXRPRippleTransactionsByAddressRIValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) GetValueOk() (*ListXRPRippleTransactionsByAddressRIValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListXRPRippleTransactionsByAddressAndTimeRangeRI) SetValue(v ListXRPRippleTransactionsByAddressRIValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


