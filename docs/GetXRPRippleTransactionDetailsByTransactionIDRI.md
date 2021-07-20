# GetXRPRippleTransactionDetailsByTransactionIDRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | **string** | Represents additional data that may be needed. | 
**Index** | **string** | Defines the index of the transaction, i.e. the consecutive place it takes in the blockchain. | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **string** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Offer** | [**GetXRPRippleTransactionDetailsByTransactionIDRIOffer**](GetXRPRippleTransactionDetailsByTransactionIDRIOffer.md) |  | 
**Receive** | [**GetXRPRippleTransactionDetailsByTransactionIDRIReceive**](GetXRPRippleTransactionDetailsByTransactionIDRIReceive.md) |  | 
**Recipients** | [**[]GetXRPRippleTransactionDetailsByTransactionIDRIRecipients**](GetXRPRippleTransactionDetailsByTransactionIDRIRecipients.md) | Represents an object of addresses that receive the transactions. | 
**Senders** | [**[]GetXRPRippleTransactionDetailsByTransactionIDRISenders**](GetXRPRippleTransactionDetailsByTransactionIDRISenders.md) | Represents an object of addresses that provide the funds. | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | Pointer to **string** | Defines the status of the transaction. | [optional] 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**Type** | **string** | Defines the type of the transaction. | 
**Fee** | [**GetXRPRippleTransactionDetailsByTransactionIDRIFee**](GetXRPRippleTransactionDetailsByTransactionIDRIFee.md) |  | 
**Value** | [**GetXRPRippleTransactionDetailsByTransactionIDRIValue**](GetXRPRippleTransactionDetailsByTransactionIDRIValue.md) |  | 

## Methods

### NewGetXRPRippleTransactionDetailsByTransactionIDRI

`func NewGetXRPRippleTransactionDetailsByTransactionIDRI(additionalData string, index string, minedInBlockHash string, minedInBlockHeight string, offer GetXRPRippleTransactionDetailsByTransactionIDRIOffer, receive GetXRPRippleTransactionDetailsByTransactionIDRIReceive, recipients []GetXRPRippleTransactionDetailsByTransactionIDRIRecipients, senders []GetXRPRippleTransactionDetailsByTransactionIDRISenders, sequence int32, timestamp int32, transactionHash string, type_ string, fee GetXRPRippleTransactionDetailsByTransactionIDRIFee, value GetXRPRippleTransactionDetailsByTransactionIDRIValue, ) *GetXRPRippleTransactionDetailsByTransactionIDRI`

NewGetXRPRippleTransactionDetailsByTransactionIDRI instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXRPRippleTransactionDetailsByTransactionIDRIWithDefaults

`func NewGetXRPRippleTransactionDetailsByTransactionIDRIWithDefaults() *GetXRPRippleTransactionDetailsByTransactionIDRI`

NewGetXRPRippleTransactionDetailsByTransactionIDRIWithDefaults instantiates a new GetXRPRippleTransactionDetailsByTransactionIDRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.


### GetIndex

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetMinedInBlockHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetMinedInBlockHeight() string`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetMinedInBlockHeightOk() (*string, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetMinedInBlockHeight(v string)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetOffer

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetOffer() GetXRPRippleTransactionDetailsByTransactionIDRIOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetOfferOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetOffer(v GetXRPRippleTransactionDetailsByTransactionIDRIOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetReceive() GetXRPRippleTransactionDetailsByTransactionIDRIReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetReceiveOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetReceive(v GetXRPRippleTransactionDetailsByTransactionIDRIReceive)`

SetReceive sets Receive field to given value.


### GetRecipients

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetRecipients() []GetXRPRippleTransactionDetailsByTransactionIDRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetRecipientsOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetRecipients(v []GetXRPRippleTransactionDetailsByTransactionIDRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetSenders() []GetXRPRippleTransactionDetailsByTransactionIDRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetSendersOk() (*[]GetXRPRippleTransactionDetailsByTransactionIDRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetSenders(v []GetXRPRippleTransactionDetailsByTransactionIDRISenders)`

SetSenders sets Senders field to given value.


### GetSequence

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetType

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetType(v string)`

SetType sets Type field to given value.


### GetFee

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetFee() GetXRPRippleTransactionDetailsByTransactionIDRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetFeeOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetFee(v GetXRPRippleTransactionDetailsByTransactionIDRIFee)`

SetFee sets Fee field to given value.


### GetValue

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetValue() GetXRPRippleTransactionDetailsByTransactionIDRIValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) GetValueOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetXRPRippleTransactionDetailsByTransactionIDRI) SetValue(v GetXRPRippleTransactionDetailsByTransactionIDRIValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


