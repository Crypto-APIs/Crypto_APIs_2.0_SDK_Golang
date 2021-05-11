# ListTransactionsByBlockHashResponseItemSenders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Amount** | **string** | Represents the amount sent by this address. | 

## Methods

### NewListTransactionsByBlockHashResponseItemSenders

`func NewListTransactionsByBlockHashResponseItemSenders(address string, amount string, ) *ListTransactionsByBlockHashResponseItemSenders`

NewListTransactionsByBlockHashResponseItemSenders instantiates a new ListTransactionsByBlockHashResponseItemSenders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemSendersWithDefaults

`func NewListTransactionsByBlockHashResponseItemSendersWithDefaults() *ListTransactionsByBlockHashResponseItemSenders`

NewListTransactionsByBlockHashResponseItemSendersWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemSenders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListTransactionsByBlockHashResponseItemSenders) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListTransactionsByBlockHashResponseItemSenders) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListTransactionsByBlockHashResponseItemSenders) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListTransactionsByBlockHashResponseItemSenders) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListTransactionsByBlockHashResponseItemSenders) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListTransactionsByBlockHashResponseItemSenders) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


