# GetZilliqaAddressDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | [**GetZilliqaAddressDetailsRIBalance**](GetZilliqaAddressDetailsRIBalance.md) |  | 
**IncomingTransactionsCount** | **int32** | Defines the received transaction count to the address. | 
**OutgoingTransactionsCount** | **int32** | Defines the sent transaction count from the address. | 
**TransactionsCount** | **int32** | Defines the entire count of the transactions. | 

## Methods

### NewGetZilliqaAddressDetailsRI

`func NewGetZilliqaAddressDetailsRI(balance GetZilliqaAddressDetailsRIBalance, incomingTransactionsCount int32, outgoingTransactionsCount int32, transactionsCount int32, ) *GetZilliqaAddressDetailsRI`

NewGetZilliqaAddressDetailsRI instantiates a new GetZilliqaAddressDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZilliqaAddressDetailsRIWithDefaults

`func NewGetZilliqaAddressDetailsRIWithDefaults() *GetZilliqaAddressDetailsRI`

NewGetZilliqaAddressDetailsRIWithDefaults instantiates a new GetZilliqaAddressDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *GetZilliqaAddressDetailsRI) GetBalance() GetZilliqaAddressDetailsRIBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetZilliqaAddressDetailsRI) GetBalanceOk() (*GetZilliqaAddressDetailsRIBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetZilliqaAddressDetailsRI) SetBalance(v GetZilliqaAddressDetailsRIBalance)`

SetBalance sets Balance field to given value.


### GetIncomingTransactionsCount

`func (o *GetZilliqaAddressDetailsRI) GetIncomingTransactionsCount() int32`

GetIncomingTransactionsCount returns the IncomingTransactionsCount field if non-nil, zero value otherwise.

### GetIncomingTransactionsCountOk

`func (o *GetZilliqaAddressDetailsRI) GetIncomingTransactionsCountOk() (*int32, bool)`

GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTransactionsCount

`func (o *GetZilliqaAddressDetailsRI) SetIncomingTransactionsCount(v int32)`

SetIncomingTransactionsCount sets IncomingTransactionsCount field to given value.


### GetOutgoingTransactionsCount

`func (o *GetZilliqaAddressDetailsRI) GetOutgoingTransactionsCount() int32`

GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field if non-nil, zero value otherwise.

### GetOutgoingTransactionsCountOk

`func (o *GetZilliqaAddressDetailsRI) GetOutgoingTransactionsCountOk() (*int32, bool)`

GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTransactionsCount

`func (o *GetZilliqaAddressDetailsRI) SetOutgoingTransactionsCount(v int32)`

SetOutgoingTransactionsCount sets OutgoingTransactionsCount field to given value.


### GetTransactionsCount

`func (o *GetZilliqaAddressDetailsRI) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *GetZilliqaAddressDetailsRI) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *GetZilliqaAddressDetailsRI) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


