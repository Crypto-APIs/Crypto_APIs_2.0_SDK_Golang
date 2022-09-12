# CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents the specific amount of the transaction&#39;s destination. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**RecipientAddress** | **string** | Defines the specific recipient/destination address. | 

## Methods

### NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem

`func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem(amount string, recipientAddress string, ) *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem`

NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItemWithDefaults

`func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItemWithDefaults() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem`

NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItemWithDefaults instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCallbackSecretKey

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetNote

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


