# GetXRPRippleAddressDetailsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | [**GetXRPRippleAddressDetailsResponseItemBalance**](GetXRPRippleAddressDetailsResponseItemBalance.md) |  | 
**IncomingTransactionsCount** | **int32** | Defines the count of all confirmed incoming transactions from the address for coins. This applies to coins only, not to tokens transfers | 
**OutgoingTransactionsCount** | **int32** | Defines the count of all confirmed outgoing transactions for coins. This applies to coins only, not to tokens transfers | 
**Sequence** | **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**TransactionsCount** | **int32** | Represents the total number of all transactions as part of this block. | 

## Methods

### NewGetXRPRippleAddressDetailsResponseItem

`func NewGetXRPRippleAddressDetailsResponseItem(balance GetXRPRippleAddressDetailsResponseItemBalance, incomingTransactionsCount int32, outgoingTransactionsCount int32, sequence int32, transactionsCount int32, ) *GetXRPRippleAddressDetailsResponseItem`

NewGetXRPRippleAddressDetailsResponseItem instantiates a new GetXRPRippleAddressDetailsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXRPRippleAddressDetailsResponseItemWithDefaults

`func NewGetXRPRippleAddressDetailsResponseItemWithDefaults() *GetXRPRippleAddressDetailsResponseItem`

NewGetXRPRippleAddressDetailsResponseItemWithDefaults instantiates a new GetXRPRippleAddressDetailsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *GetXRPRippleAddressDetailsResponseItem) GetBalance() GetXRPRippleAddressDetailsResponseItemBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetXRPRippleAddressDetailsResponseItem) GetBalanceOk() (*GetXRPRippleAddressDetailsResponseItemBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetXRPRippleAddressDetailsResponseItem) SetBalance(v GetXRPRippleAddressDetailsResponseItemBalance)`

SetBalance sets Balance field to given value.


### GetIncomingTransactionsCount

`func (o *GetXRPRippleAddressDetailsResponseItem) GetIncomingTransactionsCount() int32`

GetIncomingTransactionsCount returns the IncomingTransactionsCount field if non-nil, zero value otherwise.

### GetIncomingTransactionsCountOk

`func (o *GetXRPRippleAddressDetailsResponseItem) GetIncomingTransactionsCountOk() (*int32, bool)`

GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTransactionsCount

`func (o *GetXRPRippleAddressDetailsResponseItem) SetIncomingTransactionsCount(v int32)`

SetIncomingTransactionsCount sets IncomingTransactionsCount field to given value.


### GetOutgoingTransactionsCount

`func (o *GetXRPRippleAddressDetailsResponseItem) GetOutgoingTransactionsCount() int32`

GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field if non-nil, zero value otherwise.

### GetOutgoingTransactionsCountOk

`func (o *GetXRPRippleAddressDetailsResponseItem) GetOutgoingTransactionsCountOk() (*int32, bool)`

GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTransactionsCount

`func (o *GetXRPRippleAddressDetailsResponseItem) SetOutgoingTransactionsCount(v int32)`

SetOutgoingTransactionsCount sets OutgoingTransactionsCount field to given value.


### GetSequence

`func (o *GetXRPRippleAddressDetailsResponseItem) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetXRPRippleAddressDetailsResponseItem) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetXRPRippleAddressDetailsResponseItem) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTransactionsCount

`func (o *GetXRPRippleAddressDetailsResponseItem) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetXRPRippleAddressDetailsResponseItem) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetXRPRippleAddressDetailsResponseItem) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


