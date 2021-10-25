# NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Defines the specific address of the internal transaction. | 
**AllowDuplicates** | **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [default to false]
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | **int32** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | 

## Methods

### NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem

`func NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem(address string, allowDuplicates bool, callbackSecretKey string, callbackUrl string, confirmationsCount int32, ) *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem`

NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItem instantiates a new NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItemWithDefaults

`func NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItemWithDefaults() *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem`

NewNewConfirmedInternalTransactionsAndEachConfirmationRBDataItemWithDefaults instantiates a new NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllowDuplicates

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.


### GetCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetConfirmationsCount() int32`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) GetConfirmationsCountOk() (*int32, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *NewConfirmedInternalTransactionsAndEachConfirmationRBDataItem) SetConfirmationsCount(v int32)`

SetConfirmationsCount sets ConfirmationsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


