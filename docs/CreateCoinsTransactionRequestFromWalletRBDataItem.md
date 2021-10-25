# CreateCoinsTransactionRequestFromWalletRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**PrepareStrategy** | Pointer to **string** | Refers to a model of a UTXO spending strategy, where customers can choose how to spend their transaction outputs from multiple Bitcoin addresses. Two options available - \&quot;minimize-dust\&quot; (select lower amounts from multiple addresses) or \&quot;optimize-size\&quot; (select higher amounts from less addresses). | [optional] [default to "minimize-dust"]
**Recipients** | [**[]CreateCoinsTransactionRequestFromWalletRBDataItemRecipients**](CreateCoinsTransactionRequestFromWalletRBDataItemRecipients.md) | Defines the destination of the transaction, whether it is incoming or outgoing. | 

## Methods

### NewCreateCoinsTransactionRequestFromWalletRBDataItem

`func NewCreateCoinsTransactionRequestFromWalletRBDataItem(feePriority string, recipients []CreateCoinsTransactionRequestFromWalletRBDataItemRecipients, ) *CreateCoinsTransactionRequestFromWalletRBDataItem`

NewCreateCoinsTransactionRequestFromWalletRBDataItem instantiates a new CreateCoinsTransactionRequestFromWalletRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromWalletRBDataItemWithDefaults

`func NewCreateCoinsTransactionRequestFromWalletRBDataItemWithDefaults() *CreateCoinsTransactionRequestFromWalletRBDataItem`

NewCreateCoinsTransactionRequestFromWalletRBDataItemWithDefaults instantiates a new CreateCoinsTransactionRequestFromWalletRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeePriority

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetNote

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPrepareStrategy

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetPrepareStrategy() string`

GetPrepareStrategy returns the PrepareStrategy field if non-nil, zero value otherwise.

### GetPrepareStrategyOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetPrepareStrategyOk() (*string, bool)`

GetPrepareStrategyOk returns a tuple with the PrepareStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareStrategy

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetPrepareStrategy(v string)`

SetPrepareStrategy sets PrepareStrategy field to given value.

### HasPrepareStrategy

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) HasPrepareStrategy() bool`

HasPrepareStrategy returns a boolean if a field has been set.

### GetRecipients

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetRecipients() []CreateCoinsTransactionRequestFromWalletRBDataItemRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetRecipientsOk() (*[]CreateCoinsTransactionRequestFromWalletRBDataItemRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetRecipients(v []CreateCoinsTransactionRequestFromWalletRBDataItemRecipients)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


