# ListOmniTransactionsByAddressResponseItemSenders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the hash of the address that provides the funds. | 
**Amount** | **string** | Defines the amount of the sent funds as a string. | 

## Methods

### NewListOmniTransactionsByAddressResponseItemSenders

`func NewListOmniTransactionsByAddressResponseItemSenders(address string, amount string, ) *ListOmniTransactionsByAddressResponseItemSenders`

NewListOmniTransactionsByAddressResponseItemSenders instantiates a new ListOmniTransactionsByAddressResponseItemSenders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOmniTransactionsByAddressResponseItemSendersWithDefaults

`func NewListOmniTransactionsByAddressResponseItemSendersWithDefaults() *ListOmniTransactionsByAddressResponseItemSenders`

NewListOmniTransactionsByAddressResponseItemSendersWithDefaults instantiates a new ListOmniTransactionsByAddressResponseItemSenders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListOmniTransactionsByAddressResponseItemSenders) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListOmniTransactionsByAddressResponseItemSenders) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListOmniTransactionsByAddressResponseItemSenders) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListOmniTransactionsByAddressResponseItemSenders) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListOmniTransactionsByAddressResponseItemSenders) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListOmniTransactionsByAddressResponseItemSenders) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


