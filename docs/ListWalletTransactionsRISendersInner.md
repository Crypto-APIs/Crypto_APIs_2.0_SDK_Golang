# ListWalletTransactionsRISendersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Amount** | **string** | Represents the total amount sent by this address including the fee. | 
**Label** | Pointer to **string** | Defines a plain text string as a label for the sender. | [optional] 

## Methods

### NewListWalletTransactionsRISendersInner

`func NewListWalletTransactionsRISendersInner(address string, amount string, ) *ListWalletTransactionsRISendersInner`

NewListWalletTransactionsRISendersInner instantiates a new ListWalletTransactionsRISendersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRISendersInnerWithDefaults

`func NewListWalletTransactionsRISendersInnerWithDefaults() *ListWalletTransactionsRISendersInner`

NewListWalletTransactionsRISendersInnerWithDefaults instantiates a new ListWalletTransactionsRISendersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListWalletTransactionsRISendersInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListWalletTransactionsRISendersInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListWalletTransactionsRISendersInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListWalletTransactionsRISendersInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRISendersInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRISendersInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetLabel

`func (o *ListWalletTransactionsRISendersInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListWalletTransactionsRISendersInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListWalletTransactionsRISendersInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListWalletTransactionsRISendersInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


