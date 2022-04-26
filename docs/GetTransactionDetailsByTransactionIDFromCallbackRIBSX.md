# GetTransactionDetailsByTransactionIDFromCallbackRIBSX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | **string** | Represents additional data that may be needed. | 
**DestinationTag** | Pointer to **int64** | Defines the destination tag value. | [optional] 
**Offer** | [**GetXRPRippleTransactionDetailsByTransactionIDRIOffer**](GetXRPRippleTransactionDetailsByTransactionIDRIOffer.md) |  | 
**Receive** | [**GetXRPRippleTransactionDetailsByTransactionIDRIReceive**](GetXRPRippleTransactionDetailsByTransactionIDRIReceive.md) |  | 
**Sequence** | **int64** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | **string** | Defines the status of the transaction. | 
**Type** | **string** | Defines the type of the transaction. | 
**Value** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue**](GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue.md) |  | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSX

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSX(additionalData string, offer GetXRPRippleTransactionDetailsByTransactionIDRIOffer, receive GetXRPRippleTransactionDetailsByTransactionIDRIReceive, sequence int64, status string, type_ string, value GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSX`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSX instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSXWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSXWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSX`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSXWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.


### GetDestinationTag

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetDestinationTag() int64`

GetDestinationTag returns the DestinationTag field if non-nil, zero value otherwise.

### GetDestinationTagOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetDestinationTagOk() (*int64, bool)`

GetDestinationTagOk returns a tuple with the DestinationTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTag

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetDestinationTag(v int64)`

SetDestinationTag sets DestinationTag field to given value.

### HasDestinationTag

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) HasDestinationTag() bool`

HasDestinationTag returns a boolean if a field has been set.

### GetOffer

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetOffer() GetXRPRippleTransactionDetailsByTransactionIDRIOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetOfferOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetOffer(v GetXRPRippleTransactionDetailsByTransactionIDRIOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetReceive() GetXRPRippleTransactionDetailsByTransactionIDRIReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetReceiveOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetReceive(v GetXRPRippleTransactionDetailsByTransactionIDRIReceive)`

SetReceive sets Receive field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetValue() GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) GetValueOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSX) SetValue(v GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


