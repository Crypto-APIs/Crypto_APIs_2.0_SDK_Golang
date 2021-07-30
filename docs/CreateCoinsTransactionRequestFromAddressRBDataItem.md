# CreateCoinsTransactionRequestFromAddressRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents the specific amount of the transaction. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | [optional] 
**CallbackUrl** | Pointer to **string** | Verified URL for sending callbacks | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**RecipientAddress** | **string** | Defines the specific recipient address for the transaction. | 

## Methods

### NewCreateCoinsTransactionRequestFromAddressRBDataItem

`func NewCreateCoinsTransactionRequestFromAddressRBDataItem(amount string, feePriority string, recipientAddress string, ) *CreateCoinsTransactionRequestFromAddressRBDataItem`

NewCreateCoinsTransactionRequestFromAddressRBDataItem instantiates a new CreateCoinsTransactionRequestFromAddressRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults

`func NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults() *CreateCoinsTransactionRequestFromAddressRBDataItem`

NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeePriority

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetRecipientAddress

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


