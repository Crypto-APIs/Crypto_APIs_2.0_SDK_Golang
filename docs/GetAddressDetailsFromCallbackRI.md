# GetAddressDetailsFromCallbackRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingTransactionsCount** | **int32** | Defines the count of the incoming transactions. | 
**OutgoingTransactionsCount** | **int32** | Defines the count of the outgoing transactions. | 
**TransactionsCount** | **int32** | Represents the total number of confirmed coins transactions for this address, both incoming and outgoing. Applies for coins only **and not** tokens transfers e.g. for Ethereum. &#x60;transactionsCount&#x60; could result as less than incoming and outgoing transactions put together (e.g. in Bitcoin), due to the fact that one and the same address could be in senders and receivers addresses. | 
**ConfirmedBalance** | [**GetAddressDetailsRIConfirmedBalance**](GetAddressDetailsRIConfirmedBalance.md) |  | 
**TotalReceived** | Pointer to [**GetAddressDetailsFromCallbackRITotalReceived**](GetAddressDetailsFromCallbackRITotalReceived.md) |  | [optional] 
**TotalSpent** | Pointer to [**GetAddressDetailsFromCallbackRITotalSpent**](GetAddressDetailsFromCallbackRITotalSpent.md) |  | [optional] 
**Sequence** | Pointer to **int32** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | [optional] 

## Methods

### NewGetAddressDetailsFromCallbackRI

`func NewGetAddressDetailsFromCallbackRI(incomingTransactionsCount int32, outgoingTransactionsCount int32, transactionsCount int32, confirmedBalance GetAddressDetailsRIConfirmedBalance, ) *GetAddressDetailsFromCallbackRI`

NewGetAddressDetailsFromCallbackRI instantiates a new GetAddressDetailsFromCallbackRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddressDetailsFromCallbackRIWithDefaults

`func NewGetAddressDetailsFromCallbackRIWithDefaults() *GetAddressDetailsFromCallbackRI`

NewGetAddressDetailsFromCallbackRIWithDefaults instantiates a new GetAddressDetailsFromCallbackRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingTransactionsCount

`func (o *GetAddressDetailsFromCallbackRI) GetIncomingTransactionsCount() int32`

GetIncomingTransactionsCount returns the IncomingTransactionsCount field if non-nil, zero value otherwise.

### GetIncomingTransactionsCountOk

`func (o *GetAddressDetailsFromCallbackRI) GetIncomingTransactionsCountOk() (*int32, bool)`

GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTransactionsCount

`func (o *GetAddressDetailsFromCallbackRI) SetIncomingTransactionsCount(v int32)`

SetIncomingTransactionsCount sets IncomingTransactionsCount field to given value.


### GetOutgoingTransactionsCount

`func (o *GetAddressDetailsFromCallbackRI) GetOutgoingTransactionsCount() int32`

GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field if non-nil, zero value otherwise.

### GetOutgoingTransactionsCountOk

`func (o *GetAddressDetailsFromCallbackRI) GetOutgoingTransactionsCountOk() (*int32, bool)`

GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTransactionsCount

`func (o *GetAddressDetailsFromCallbackRI) SetOutgoingTransactionsCount(v int32)`

SetOutgoingTransactionsCount sets OutgoingTransactionsCount field to given value.


### GetTransactionsCount

`func (o *GetAddressDetailsFromCallbackRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetAddressDetailsFromCallbackRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetAddressDetailsFromCallbackRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetConfirmedBalance

`func (o *GetAddressDetailsFromCallbackRI) GetConfirmedBalance() GetAddressDetailsRIConfirmedBalance`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *GetAddressDetailsFromCallbackRI) GetConfirmedBalanceOk() (*GetAddressDetailsRIConfirmedBalance, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *GetAddressDetailsFromCallbackRI) SetConfirmedBalance(v GetAddressDetailsRIConfirmedBalance)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetTotalReceived

`func (o *GetAddressDetailsFromCallbackRI) GetTotalReceived() GetAddressDetailsFromCallbackRITotalReceived`

GetTotalReceived returns the TotalReceived field if non-nil, zero value otherwise.

### GetTotalReceivedOk

`func (o *GetAddressDetailsFromCallbackRI) GetTotalReceivedOk() (*GetAddressDetailsFromCallbackRITotalReceived, bool)`

GetTotalReceivedOk returns a tuple with the TotalReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceived

`func (o *GetAddressDetailsFromCallbackRI) SetTotalReceived(v GetAddressDetailsFromCallbackRITotalReceived)`

SetTotalReceived sets TotalReceived field to given value.

### HasTotalReceived

`func (o *GetAddressDetailsFromCallbackRI) HasTotalReceived() bool`

HasTotalReceived returns a boolean if a field has been set.

### GetTotalSpent

`func (o *GetAddressDetailsFromCallbackRI) GetTotalSpent() GetAddressDetailsFromCallbackRITotalSpent`

GetTotalSpent returns the TotalSpent field if non-nil, zero value otherwise.

### GetTotalSpentOk

`func (o *GetAddressDetailsFromCallbackRI) GetTotalSpentOk() (*GetAddressDetailsFromCallbackRITotalSpent, bool)`

GetTotalSpentOk returns a tuple with the TotalSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpent

`func (o *GetAddressDetailsFromCallbackRI) SetTotalSpent(v GetAddressDetailsFromCallbackRITotalSpent)`

SetTotalSpent sets TotalSpent field to given value.

### HasTotalSpent

`func (o *GetAddressDetailsFromCallbackRI) HasTotalSpent() bool`

HasTotalSpent returns a boolean if a field has been set.

### GetSequence

`func (o *GetAddressDetailsFromCallbackRI) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetAddressDetailsFromCallbackRI) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetAddressDetailsFromCallbackRI) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *GetAddressDetailsFromCallbackRI) HasSequence() bool`

HasSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


