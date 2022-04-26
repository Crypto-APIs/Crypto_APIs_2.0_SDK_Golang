# CreateFungibleTokensTransactionRequestFromAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**Recipients** | [**[]CreateFungibleTokensTransactionRequestFromAddressRIRecipients**](CreateFungibleTokensTransactionRequestFromAddressRIRecipients.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Senders** | [**CreateFungibleTokensTransactionRequestFromAddressRISenders**](CreateFungibleTokensTransactionRequestFromAddressRISenders.md) |  | 
**TokenTypeSpecificData** | [**CreateFungibleTokensTransactionRequestFromAddressRIS**](CreateFungibleTokensTransactionRequestFromAddressRIS.md) |  | 
**TransactionRequestId** | **string** | Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which &#x60;referenceId&#x60; concern that specific transaction request. | 

## Methods

### NewCreateFungibleTokensTransactionRequestFromAddressRI

`func NewCreateFungibleTokensTransactionRequestFromAddressRI(callbackSecretKey string, callbackUrl string, feePriority string, recipients []CreateFungibleTokensTransactionRequestFromAddressRIRecipients, senders CreateFungibleTokensTransactionRequestFromAddressRISenders, tokenTypeSpecificData CreateFungibleTokensTransactionRequestFromAddressRIS, transactionRequestId string, ) *CreateFungibleTokensTransactionRequestFromAddressRI`

NewCreateFungibleTokensTransactionRequestFromAddressRI instantiates a new CreateFungibleTokensTransactionRequestFromAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFungibleTokensTransactionRequestFromAddressRIWithDefaults

`func NewCreateFungibleTokensTransactionRequestFromAddressRIWithDefaults() *CreateFungibleTokensTransactionRequestFromAddressRI`

NewCreateFungibleTokensTransactionRequestFromAddressRIWithDefaults instantiates a new CreateFungibleTokensTransactionRequestFromAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetFeePriority

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetNote

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipients

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetRecipients() []CreateFungibleTokensTransactionRequestFromAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetRecipientsOk() (*[]CreateFungibleTokensTransactionRequestFromAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetRecipients(v []CreateFungibleTokensTransactionRequestFromAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetSenders() CreateFungibleTokensTransactionRequestFromAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetSendersOk() (*CreateFungibleTokensTransactionRequestFromAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetSenders(v CreateFungibleTokensTransactionRequestFromAddressRISenders)`

SetSenders sets Senders field to given value.


### GetTokenTypeSpecificData

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificData() CreateFungibleTokensTransactionRequestFromAddressRIS`

GetTokenTypeSpecificData returns the TokenTypeSpecificData field if non-nil, zero value otherwise.

### GetTokenTypeSpecificDataOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificDataOk() (*CreateFungibleTokensTransactionRequestFromAddressRIS, bool)`

GetTokenTypeSpecificDataOk returns a tuple with the TokenTypeSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTypeSpecificData

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetTokenTypeSpecificData(v CreateFungibleTokensTransactionRequestFromAddressRIS)`

SetTokenTypeSpecificData sets TokenTypeSpecificData field to given value.


### GetTransactionRequestId

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTransactionRequestId() string`

GetTransactionRequestId returns the TransactionRequestId field if non-nil, zero value otherwise.

### GetTransactionRequestIdOk

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) GetTransactionRequestIdOk() (*string, bool)`

GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestId

`func (o *CreateFungibleTokensTransactionRequestFromAddressRI) SetTransactionRequestId(v string)`

SetTransactionRequestId sets TransactionRequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


