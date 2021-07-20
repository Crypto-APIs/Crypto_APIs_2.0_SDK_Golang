# GetAddressDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionsCount** | **int32** | Represents the total number of confirmed coins transactions for this address, both incoming and outgoing. Applies for coins only **and not** tokens transfers e.g. for Ethereum. &#x60;transactionsCount&#x60; could result as less than incoming and outgoing transactions put together (e.g. in Bitcoin), due to the fact that one and the same address could be in senders and receivers addresses. | 
**ConfirmedBalance** | [**GetAddressDetailsRIConfirmedBalance**](GetAddressDetailsRIConfirmedBalance.md) |  | 
**TotalReceived** | [**GetAddressDetailsRITotalReceived**](GetAddressDetailsRITotalReceived.md) |  | 
**TotalSpent** | [**GetAddressDetailsRITotalSpent**](GetAddressDetailsRITotalSpent.md) |  | 
**IncomingTransactionsCount** | **int32** | Defines the count of all confirmed incoming transactions from the address for coins. This applies to **coins** only, **not** to tokens transfers e.g. for Ethereum. | 
**OutgoingTransactionsCount** | **int32** | Defines the count of all confirmed outgoing transactions from the address for coins. This applies to **coins** only, **not** to tokens transfers e.g. for Ethereum. | 

## Methods

### NewGetAddressDetailsRI

`func NewGetAddressDetailsRI(transactionsCount int32, confirmedBalance GetAddressDetailsRIConfirmedBalance, totalReceived GetAddressDetailsRITotalReceived, totalSpent GetAddressDetailsRITotalSpent, incomingTransactionsCount int32, outgoingTransactionsCount int32, ) *GetAddressDetailsRI`

NewGetAddressDetailsRI instantiates a new GetAddressDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddressDetailsRIWithDefaults

`func NewGetAddressDetailsRIWithDefaults() *GetAddressDetailsRI`

NewGetAddressDetailsRIWithDefaults instantiates a new GetAddressDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionsCount

`func (o *GetAddressDetailsRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetAddressDetailsRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetAddressDetailsRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.


### GetConfirmedBalance

`func (o *GetAddressDetailsRI) GetConfirmedBalance() GetAddressDetailsRIConfirmedBalance`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *GetAddressDetailsRI) GetConfirmedBalanceOk() (*GetAddressDetailsRIConfirmedBalance, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *GetAddressDetailsRI) SetConfirmedBalance(v GetAddressDetailsRIConfirmedBalance)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetTotalReceived

`func (o *GetAddressDetailsRI) GetTotalReceived() GetAddressDetailsRITotalReceived`

GetTotalReceived returns the TotalReceived field if non-nil, zero value otherwise.

### GetTotalReceivedOk

`func (o *GetAddressDetailsRI) GetTotalReceivedOk() (*GetAddressDetailsRITotalReceived, bool)`

GetTotalReceivedOk returns a tuple with the TotalReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceived

`func (o *GetAddressDetailsRI) SetTotalReceived(v GetAddressDetailsRITotalReceived)`

SetTotalReceived sets TotalReceived field to given value.


### GetTotalSpent

`func (o *GetAddressDetailsRI) GetTotalSpent() GetAddressDetailsRITotalSpent`

GetTotalSpent returns the TotalSpent field if non-nil, zero value otherwise.

### GetTotalSpentOk

`func (o *GetAddressDetailsRI) GetTotalSpentOk() (*GetAddressDetailsRITotalSpent, bool)`

GetTotalSpentOk returns a tuple with the TotalSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpent

`func (o *GetAddressDetailsRI) SetTotalSpent(v GetAddressDetailsRITotalSpent)`

SetTotalSpent sets TotalSpent field to given value.


### GetIncomingTransactionsCount

`func (o *GetAddressDetailsRI) GetIncomingTransactionsCount() int32`

GetIncomingTransactionsCount returns the IncomingTransactionsCount field if non-nil, zero value otherwise.

### GetIncomingTransactionsCountOk

`func (o *GetAddressDetailsRI) GetIncomingTransactionsCountOk() (*int32, bool)`

GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTransactionsCount

`func (o *GetAddressDetailsRI) SetIncomingTransactionsCount(v int32)`

SetIncomingTransactionsCount sets IncomingTransactionsCount field to given value.


### GetOutgoingTransactionsCount

`func (o *GetAddressDetailsRI) GetOutgoingTransactionsCount() int32`

GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field if non-nil, zero value otherwise.

### GetOutgoingTransactionsCountOk

`func (o *GetAddressDetailsRI) GetOutgoingTransactionsCountOk() (*int32, bool)`

GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTransactionsCount

`func (o *GetAddressDetailsRI) SetOutgoingTransactionsCount(v int32)`

SetOutgoingTransactionsCount sets OutgoingTransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


