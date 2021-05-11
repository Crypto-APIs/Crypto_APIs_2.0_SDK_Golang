# NewUnconfirmedCoinsTransactionsRequestBodyDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address of the transaction, per which the result is returned. | 
**AllowDuplicates** | Pointer to **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [optional] [default to false]
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs 2.0. | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 

## Methods

### NewNewUnconfirmedCoinsTransactionsRequestBodyDataItem

`func NewNewUnconfirmedCoinsTransactionsRequestBodyDataItem(address string, callbackUrl string, ) *NewUnconfirmedCoinsTransactionsRequestBodyDataItem`

NewNewUnconfirmedCoinsTransactionsRequestBodyDataItem instantiates a new NewUnconfirmedCoinsTransactionsRequestBodyDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUnconfirmedCoinsTransactionsRequestBodyDataItemWithDefaults

`func NewNewUnconfirmedCoinsTransactionsRequestBodyDataItemWithDefaults() *NewUnconfirmedCoinsTransactionsRequestBodyDataItem`

NewNewUnconfirmedCoinsTransactionsRequestBodyDataItemWithDefaults instantiates a new NewUnconfirmedCoinsTransactionsRequestBodyDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllowDuplicates

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetCallbackSecretKey

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *NewUnconfirmedCoinsTransactionsRequestBodyDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


