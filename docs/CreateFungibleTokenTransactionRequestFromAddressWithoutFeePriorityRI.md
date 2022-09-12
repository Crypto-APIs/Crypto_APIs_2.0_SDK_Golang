# CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**Recipient** | [**[]CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner**](CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Sender** | [**CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender**](CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender.md) |  | 
**TokenTypeSpecificData** | [**CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIS**](CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIS.md) |  | 
**TransactionRequestId** | **string** | Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which &#x60;referenceId&#x60; concern that specific transaction request. | 

## Methods

### NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI

`func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI(recipient []CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner, sender CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender, tokenTypeSpecificData CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIS, transactionRequestId string, ) *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI`

NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIWithDefaults

`func NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIWithDefaults() *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI`

NewCreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIWithDefaults instantiates a new CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetNote

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipient

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetRecipient() []CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetRecipientOk() (*[]CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetRecipient(v []CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetSender() CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetSenderOk() (*CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetSender(v CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender)`

SetSender sets Sender field to given value.


### GetTokenTypeSpecificData

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetTokenTypeSpecificData() CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIS`

GetTokenTypeSpecificData returns the TokenTypeSpecificData field if non-nil, zero value otherwise.

### GetTokenTypeSpecificDataOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetTokenTypeSpecificDataOk() (*CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIS, bool)`

GetTokenTypeSpecificDataOk returns a tuple with the TokenTypeSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTypeSpecificData

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetTokenTypeSpecificData(v CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRIS)`

SetTokenTypeSpecificData sets TokenTypeSpecificData field to given value.


### GetTransactionRequestId

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetTransactionRequestId() string`

GetTransactionRequestId returns the TransactionRequestId field if non-nil, zero value otherwise.

### GetTransactionRequestIdOk

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) GetTransactionRequestIdOk() (*string, bool)`

GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestId

`func (o *CreateFungibleTokenTransactionRequestFromAddressWithoutFeePriorityRI) SetTransactionRequestId(v string)`

SetTransactionRequestId sets TransactionRequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


