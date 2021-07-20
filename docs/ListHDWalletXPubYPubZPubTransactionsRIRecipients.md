# ListHDWalletXPubYPubZPubTransactionsRIRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**Amount** | **string** | Represents the amount received to this address. | 
**IsMember** | **bool** | Defines whether an address is a child address derived from the HD wallet (xPub, yPub, zPub) as boolean. | 

## Methods

### NewListHDWalletXPubYPubZPubTransactionsRIRecipients

`func NewListHDWalletXPubYPubZPubTransactionsRIRecipients(address string, amount string, isMember bool, ) *ListHDWalletXPubYPubZPubTransactionsRIRecipients`

NewListHDWalletXPubYPubZPubTransactionsRIRecipients instantiates a new ListHDWalletXPubYPubZPubTransactionsRIRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHDWalletXPubYPubZPubTransactionsRIRecipientsWithDefaults

`func NewListHDWalletXPubYPubZPubTransactionsRIRecipientsWithDefaults() *ListHDWalletXPubYPubZPubTransactionsRIRecipients`

NewListHDWalletXPubYPubZPubTransactionsRIRecipientsWithDefaults instantiates a new ListHDWalletXPubYPubZPubTransactionsRIRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsMember

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *ListHDWalletXPubYPubZPubTransactionsRIRecipients) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


