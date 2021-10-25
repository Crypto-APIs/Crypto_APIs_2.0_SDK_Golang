# ListWalletTransactionsRISenders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Amount** | **string** | Represents the total amount sent by this address including the fee. | 
**Label** | Pointer to **string** | Defines a plain text string as a label for the sender. | [optional] 

## Methods

### NewListWalletTransactionsRISenders

`func NewListWalletTransactionsRISenders(address string, amount string, ) *ListWalletTransactionsRISenders`

NewListWalletTransactionsRISenders instantiates a new ListWalletTransactionsRISenders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRISendersWithDefaults

`func NewListWalletTransactionsRISendersWithDefaults() *ListWalletTransactionsRISenders`

NewListWalletTransactionsRISendersWithDefaults instantiates a new ListWalletTransactionsRISenders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListWalletTransactionsRISenders) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListWalletTransactionsRISenders) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListWalletTransactionsRISenders) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListWalletTransactionsRISenders) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRISenders) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRISenders) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetLabel

`func (o *ListWalletTransactionsRISenders) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListWalletTransactionsRISenders) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListWalletTransactionsRISenders) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListWalletTransactionsRISenders) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


