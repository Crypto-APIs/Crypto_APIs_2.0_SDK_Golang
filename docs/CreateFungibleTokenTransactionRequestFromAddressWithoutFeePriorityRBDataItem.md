# CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents the specific amount of the transaction. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**FeeLimit** | **string** | Fee limit of the smart contract | 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**RecipientAddress** | **string** | Defines the specific recipient address for the transaction. | 
**TokenIdentifier** | **string** | Token identifier - for BITCOIN BASED should be property id e.g 31 for ETHEREUM BASED shoud be contract e.g 0xdac17f958d2ee523a2206206994597c13d831ec7 | 

## Methods

### NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem

`func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem(amount string, feeLimit string, recipientAddress string, tokenIdentifier string, ) *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem`

NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItemWithDefaults

`func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItemWithDefaults() *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem`

NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItemWithDefaults instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCallbackSecretKey

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeeLimit

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetFeeLimit() string`

GetFeeLimit returns the FeeLimit field if non-nil, zero value otherwise.

### GetFeeLimitOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetFeeLimitOk() (*string, bool)`

GetFeeLimitOk returns a tuple with the FeeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeLimit

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetFeeLimit(v string)`

SetFeeLimit sets FeeLimit field to given value.


### GetNote

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetTokenIdentifier

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetTokenIdentifier() string`

GetTokenIdentifier returns the TokenIdentifier field if non-nil, zero value otherwise.

### GetTokenIdentifierOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) GetTokenIdentifierOk() (*string, bool)`

GetTokenIdentifierOk returns a tuple with the TokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdentifier

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRBDataItem) SetTokenIdentifier(v string)`

SetTokenIdentifier sets TokenIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


