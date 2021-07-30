# CreateTokensTransactionRequestFromAddressRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents the specific amount of the transaction. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | [optional] 
**CallbackUrl** | Pointer to **string** | Verified URL for sending callbacks | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**RecipientAddress** | **string** | Defines the specific recipient address for the transaction. | 
**TokenIdentifier** | **string** | Defines the specific token identifier. For Bitcoin-based transactions it should be the &#x60;propertyId&#x60; and for Ethereum-based transactions - the &#x60;contract&#x60;. | 

## Methods

### NewCreateTokensTransactionRequestFromAddressRBDataItem

`func NewCreateTokensTransactionRequestFromAddressRBDataItem(amount string, feePriority string, recipientAddress string, tokenIdentifier string, ) *CreateTokensTransactionRequestFromAddressRBDataItem`

NewCreateTokensTransactionRequestFromAddressRBDataItem instantiates a new CreateTokensTransactionRequestFromAddressRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokensTransactionRequestFromAddressRBDataItemWithDefaults

`func NewCreateTokensTransactionRequestFromAddressRBDataItemWithDefaults() *CreateTokensTransactionRequestFromAddressRBDataItem`

NewCreateTokensTransactionRequestFromAddressRBDataItemWithDefaults instantiates a new CreateTokensTransactionRequestFromAddressRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCallbackSecretKey

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeePriority

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetRecipientAddress

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetTokenIdentifier

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetTokenIdentifier() string`

GetTokenIdentifier returns the TokenIdentifier field if non-nil, zero value otherwise.

### GetTokenIdentifierOk

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) GetTokenIdentifierOk() (*string, bool)`

GetTokenIdentifierOk returns a tuple with the TokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdentifier

`func (o *CreateTokensTransactionRequestFromAddressRBDataItem) SetTokenIdentifier(v string)`

SetTokenIdentifier sets TokenIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


