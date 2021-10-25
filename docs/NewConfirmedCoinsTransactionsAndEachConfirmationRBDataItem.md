# NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address of the transaction, per which the result is returned. | 
**AllowDuplicates** | Pointer to **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [optional] [default to false]
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | **int32** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 

## Methods

### NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem

`func NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem(address string, callbackUrl string, confirmationsCount int32, ) *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem`

NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItemWithDefaults

`func NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItemWithDefaults() *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem`

NewNewConfirmedCoinsTransactionsAndEachConfirmationRBDataItemWithDefaults instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllowDuplicates

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetCallbackSecretKey

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetConfirmationsCount() int32`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) GetConfirmationsCountOk() (*int32, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRBDataItem) SetConfirmationsCount(v int32)`

SetConfirmationsCount sets ConfirmationsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


