# MinedTransactionRequestBodyDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicates** | Pointer to **bool** | Specifies a flag that permits or denies the creation of duplicate addresses. | [optional] 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**TransactionId** | **string** | Represents the unique identification string that defines the transaction. | 

## Methods

### NewMinedTransactionRequestBodyDataItem

`func NewMinedTransactionRequestBodyDataItem(callbackUrl string, transactionId string, ) *MinedTransactionRequestBodyDataItem`

NewMinedTransactionRequestBodyDataItem instantiates a new MinedTransactionRequestBodyDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinedTransactionRequestBodyDataItemWithDefaults

`func NewMinedTransactionRequestBodyDataItemWithDefaults() *MinedTransactionRequestBodyDataItem`

NewMinedTransactionRequestBodyDataItemWithDefaults instantiates a new MinedTransactionRequestBodyDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDuplicates

`func (o *MinedTransactionRequestBodyDataItem) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *MinedTransactionRequestBodyDataItem) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *MinedTransactionRequestBodyDataItem) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *MinedTransactionRequestBodyDataItem) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.

### GetCallbackSecretKey

`func (o *MinedTransactionRequestBodyDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *MinedTransactionRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *MinedTransactionRequestBodyDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *MinedTransactionRequestBodyDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *MinedTransactionRequestBodyDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *MinedTransactionRequestBodyDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *MinedTransactionRequestBodyDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetTransactionId

`func (o *MinedTransactionRequestBodyDataItem) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *MinedTransactionRequestBodyDataItem) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *MinedTransactionRequestBodyDataItem) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


