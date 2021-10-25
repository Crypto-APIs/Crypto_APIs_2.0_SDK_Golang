# ListWalletTransactionsRIRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**Amount** | **string** | Represents the amount received to this address. | 
**Label** | Pointer to **string** | Defines a plain text string as a label for the recipient. | [optional] 

## Methods

### NewListWalletTransactionsRIRecipients

`func NewListWalletTransactionsRIRecipients(address string, amount string, ) *ListWalletTransactionsRIRecipients`

NewListWalletTransactionsRIRecipients instantiates a new ListWalletTransactionsRIRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIRecipientsWithDefaults

`func NewListWalletTransactionsRIRecipientsWithDefaults() *ListWalletTransactionsRIRecipients`

NewListWalletTransactionsRIRecipientsWithDefaults instantiates a new ListWalletTransactionsRIRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListWalletTransactionsRIRecipients) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListWalletTransactionsRIRecipients) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListWalletTransactionsRIRecipients) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListWalletTransactionsRIRecipients) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIRecipients) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIRecipients) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetLabel

`func (o *ListWalletTransactionsRIRecipients) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListWalletTransactionsRIRecipients) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListWalletTransactionsRIRecipients) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListWalletTransactionsRIRecipients) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


