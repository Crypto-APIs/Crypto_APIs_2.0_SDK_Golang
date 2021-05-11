# NewUnconfirmedTokensTransactionsRequestBodyDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address of the transaction, per which the result is returned. | 
**AllowDuplicates** | Pointer to **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [optional] [default to false]
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 

## Methods

### NewNewUnconfirmedTokensTransactionsRequestBodyDataItem

`func NewNewUnconfirmedTokensTransactionsRequestBodyDataItem(address string, callbackUrl string, ) *NewUnconfirmedTokensTransactionsRequestBodyDataItem`

NewNewUnconfirmedTokensTransactionsRequestBodyDataItem instantiates a new NewUnconfirmedTokensTransactionsRequestBodyDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUnconfirmedTokensTransactionsRequestBodyDataItemWithDefaults

`func NewNewUnconfirmedTokensTransactionsRequestBodyDataItemWithDefaults() *NewUnconfirmedTokensTransactionsRequestBodyDataItem`

NewNewUnconfirmedTokensTransactionsRequestBodyDataItemWithDefaults instantiates a new NewUnconfirmedTokensTransactionsRequestBodyDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllowDuplicates

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetCallbackSecretKey

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewUnconfirmedTokensTransactionsRequestBodyDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


