# CreateCoinsTransactionFromAddressForWholeAmountRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**RecipientAddress** | **string** | Defines the specific recipient address for the transaction. | 

## Methods

### NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItem

`func NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItem(feePriority string, recipientAddress string, ) *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem`

NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItem instantiates a new CreateCoinsTransactionFromAddressForWholeAmountRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItemWithDefaults

`func NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItemWithDefaults() *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem`

NewCreateCoinsTransactionFromAddressForWholeAmountRBDataItemWithDefaults instantiates a new CreateCoinsTransactionFromAddressForWholeAmountRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeePriority

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetNote

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *CreateCoinsTransactionFromAddressForWholeAmountRBDataItem) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


