# PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** | Representation of the additional data. | [optional] 
**Amount** | **string** | Representation of the amount of the transaction | 
**Fee** | [**PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee**](PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee.md) |  | 
**Nonce** | Pointer to **string** | Representation of the nonce value | [optional] 
**Recipient** | **string** | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Sender** | **string** | Represents a  sender address with the respective amount. In account-based protocols like Ethereum there is only one address in this list. | 
**TransactionType** | Pointer to **string** | Representation of the transaction type | [optional] 
**Xpub** | **string** | Defines the account extended publicly known key which is used to derive all child public keys. | 

## Methods

### NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem

`func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem(amount string, fee PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee, recipient string, sender string, xpub string, ) *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem`

NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemWithDefaults

`func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemWithDefaults() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem`

NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemWithDefaults instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetFee() PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetFeeOk() (*PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetFee(v PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee)`

SetFee sets Fee field to given value.


### GetNonce

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetRecipient

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetSender(v string)`

SetSender sets Sender field to given value.


### GetTransactionType

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetXpub

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetXpub(v string)`

SetXpub sets Xpub field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


