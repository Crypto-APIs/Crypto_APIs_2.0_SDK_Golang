# GetXRPRippleAddressDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | [**GetXRPRippleAddressDetailsRIBalance**](GetXRPRippleAddressDetailsRIBalance.md) |  | 
**IncomingTransactionsCount** | **int32** | Defines the count of all confirmed incoming transactions from the address for coins. This applies to coins only, not to tokens transfers | 
**OutgoingTransactionsCount** | **int32** | Defines the count of all confirmed outgoing transactions for coins. This applies to coins only, not to tokens transfers | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 

## Methods

### NewGetXRPRippleAddressDetailsRI

`func NewGetXRPRippleAddressDetailsRI(balance GetXRPRippleAddressDetailsRIBalance, incomingTransactionsCount int32, outgoingTransactionsCount int32, sequence int32, transactionsCount int32, ) *GetXRPRippleAddressDetailsRI`

NewGetXRPRippleAddressDetailsRI instantiates a new GetXRPRippleAddressDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXRPRippleAddressDetailsRIWithDefaults

`func NewGetXRPRippleAddressDetailsRIWithDefaults() *GetXRPRippleAddressDetailsRI`

NewGetXRPRippleAddressDetailsRIWithDefaults instantiates a new GetXRPRippleAddressDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *GetXRPRippleAddressDetailsRI) GetBalance() GetXRPRippleAddressDetailsRIBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetXRPRippleAddressDetailsRI) GetBalanceOk() (*GetXRPRippleAddressDetailsRIBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetXRPRippleAddressDetailsRI) SetBalance(v GetXRPRippleAddressDetailsRIBalance)`

SetBalance sets Balance field to given value.


### GetIncomingTransactionsCount

`func (o *GetXRPRippleAddressDetailsRI) GetIncomingTransactionsCount() int32`

GetIncomingTransactionsCount returns the IncomingTransactionsCount field if non-nil, zero value otherwise.

### GetIncomingTransactionsCountOk

`func (o *GetXRPRippleAddressDetailsRI) GetIncomingTransactionsCountOk() (*int32, bool)`

GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTransactionsCount

`func (o *GetXRPRippleAddressDetailsRI) SetIncomingTransactionsCount(v int32)`

SetIncomingTransactionsCount sets IncomingTransactionsCount field to given value.


### GetOutgoingTransactionsCount

`func (o *GetXRPRippleAddressDetailsRI) GetOutgoingTransactionsCount() int32`

GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field if non-nil, zero value otherwise.

### GetOutgoingTransactionsCountOk

`func (o *GetXRPRippleAddressDetailsRI) GetOutgoingTransactionsCountOk() (*int32, bool)`

GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTransactionsCount

`func (o *GetXRPRippleAddressDetailsRI) SetOutgoingTransactionsCount(v int32)`

SetOutgoingTransactionsCount sets OutgoingTransactionsCount field to given value.


### GetSequence

`func (o *GetXRPRippleAddressDetailsRI) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetXRPRippleAddressDetailsRI) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetXRPRippleAddressDetailsRI) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTransactionsCount

`func (o *GetXRPRippleAddressDetailsRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetXRPRippleAddressDetailsRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetXRPRippleAddressDetailsRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


