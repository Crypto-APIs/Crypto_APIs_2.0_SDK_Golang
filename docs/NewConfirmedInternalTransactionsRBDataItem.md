# NewConfirmedInternalTransactionsRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Defines the specific address of the internal transaction. | 
**AllowDuplicates** | **bool** | Flag that permits or denies creation of duplicates | [default to false]
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 

## Methods

### NewNewConfirmedInternalTransactionsRBDataItem

`func NewNewConfirmedInternalTransactionsRBDataItem(address string, allowDuplicates bool, callbackSecretKey string, callbackUrl string, ) *NewConfirmedInternalTransactionsRBDataItem`

NewNewConfirmedInternalTransactionsRBDataItem instantiates a new NewConfirmedInternalTransactionsRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfirmedInternalTransactionsRBDataItemWithDefaults

`func NewNewConfirmedInternalTransactionsRBDataItemWithDefaults() *NewConfirmedInternalTransactionsRBDataItem`

NewNewConfirmedInternalTransactionsRBDataItemWithDefaults instantiates a new NewConfirmedInternalTransactionsRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewConfirmedInternalTransactionsRBDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllowDuplicates

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *NewConfirmedInternalTransactionsRBDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.


### GetCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewConfirmedInternalTransactionsRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewConfirmedInternalTransactionsRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewConfirmedInternalTransactionsRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


