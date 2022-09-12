# PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** | Representation of the additional data. | [optional] 
**Fee** | [**PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee**](PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee.md) |  | 
**Locktime** | Pointer to **int32** | Represents the time at which a particular transaction can be added to the blockchain. | [optional] 
**PrepareStrategy** | Pointer to **string** | Representation of the transaction&#39;s strategy type | [optional] 
**Recipients** | [**[]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner**](PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Replaceable** | Pointer to **bool** | Representation of whether the transaction is replaceable | [optional] 
**Xpub** | **string** | Defines the account extended publicly known key which is used to derive all child public keys. | 

## Methods

### NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem

`func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem(fee PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee, recipients []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner, xpub string, ) *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem`

NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemWithDefaults

`func NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemWithDefaults() *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem`

NewPrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemWithDefaults instantiates a new PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetFee

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetFee() PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetFeeOk() (*PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetFee(v PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemFee)`

SetFee sets Fee field to given value.


### GetLocktime

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.

### HasLocktime

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasLocktime() bool`

HasLocktime returns a boolean if a field has been set.

### GetPrepareStrategy

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetPrepareStrategy() string`

GetPrepareStrategy returns the PrepareStrategy field if non-nil, zero value otherwise.

### GetPrepareStrategyOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetPrepareStrategyOk() (*string, bool)`

GetPrepareStrategyOk returns a tuple with the PrepareStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareStrategy

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetPrepareStrategy(v string)`

SetPrepareStrategy sets PrepareStrategy field to given value.

### HasPrepareStrategy

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasPrepareStrategy() bool`

HasPrepareStrategy returns a boolean if a field has been set.

### GetRecipients

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetRecipients() []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetRecipientsOk() (*[]PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetRecipients(v []PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItemRecipientsInner)`

SetRecipients sets Recipients field to given value.


### GetReplaceable

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetReplaceable() bool`

GetReplaceable returns the Replaceable field if non-nil, zero value otherwise.

### GetReplaceableOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetReplaceableOk() (*bool, bool)`

GetReplaceableOk returns a tuple with the Replaceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceable

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetReplaceable(v bool)`

SetReplaceable sets Replaceable field to given value.

### HasReplaceable

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) HasReplaceable() bool`

HasReplaceable returns a boolean if a field has been set.

### GetXpub

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *PrepareAUTXOBasedTransactionFromHDWalletXPubYPubZPubRBDataItem) SetXpub(v string)`

SetXpub sets Xpub field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


