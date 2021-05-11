# ListHDWalletxPubYPubZPubTransactionsResponseItemSenders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Amount** | **string** | Represents the amount sent by this address. | 
**IsMember** | **bool** | Defines whether an address is a child address derived from the HD wallet (xPub, yPub, zPub) as boolean. | 

## Methods

### NewListHDWalletxPubYPubZPubTransactionsResponseItemSenders

`func NewListHDWalletxPubYPubZPubTransactionsResponseItemSenders(address string, amount string, isMember bool, ) *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders`

NewListHDWalletxPubYPubZPubTransactionsResponseItemSenders instantiates a new ListHDWalletxPubYPubZPubTransactionsResponseItemSenders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHDWalletxPubYPubZPubTransactionsResponseItemSendersWithDefaults

`func NewListHDWalletxPubYPubZPubTransactionsResponseItemSendersWithDefaults() *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders`

NewListHDWalletxPubYPubZPubTransactionsResponseItemSendersWithDefaults instantiates a new ListHDWalletxPubYPubZPubTransactionsResponseItemSenders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsMember

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *ListHDWalletxPubYPubZPubTransactionsResponseItemSenders) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


