# CreateTokensTransactionRequestFromAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | 
**CallbackUrl** | **string** | Verified URL for sending callbacks | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Recipients** | [**[]CreateTokensTransactionRequestFromAddressRIRecipients**](CreateTokensTransactionRequestFromAddressRIRecipients.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Senders** | [**CreateTokensTransactionRequestFromAddressRISenders**](CreateTokensTransactionRequestFromAddressRISenders.md) |  | 
**TokenTypeSpecificData** | [**CreateTokensTransactionRequestFromAddressRIS**](CreateTokensTransactionRequestFromAddressRIS.md) |  | 

## Methods

### NewCreateTokensTransactionRequestFromAddressRI

`func NewCreateTokensTransactionRequestFromAddressRI(callbackSecretKey string, callbackUrl string, feePriority string, recipients []CreateTokensTransactionRequestFromAddressRIRecipients, senders CreateTokensTransactionRequestFromAddressRISenders, tokenTypeSpecificData CreateTokensTransactionRequestFromAddressRIS, ) *CreateTokensTransactionRequestFromAddressRI`

NewCreateTokensTransactionRequestFromAddressRI instantiates a new CreateTokensTransactionRequestFromAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokensTransactionRequestFromAddressRIWithDefaults

`func NewCreateTokensTransactionRequestFromAddressRIWithDefaults() *CreateTokensTransactionRequestFromAddressRI`

NewCreateTokensTransactionRequestFromAddressRIWithDefaults instantiates a new CreateTokensTransactionRequestFromAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateTokensTransactionRequestFromAddressRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.


### GetCallbackUrl

`func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateTokensTransactionRequestFromAddressRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetFeePriority

`func (o *CreateTokensTransactionRequestFromAddressRI) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateTokensTransactionRequestFromAddressRI) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetRecipients

`func (o *CreateTokensTransactionRequestFromAddressRI) GetRecipients() []CreateTokensTransactionRequestFromAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetRecipientsOk() (*[]CreateTokensTransactionRequestFromAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateTokensTransactionRequestFromAddressRI) SetRecipients(v []CreateTokensTransactionRequestFromAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *CreateTokensTransactionRequestFromAddressRI) GetSenders() CreateTokensTransactionRequestFromAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetSendersOk() (*CreateTokensTransactionRequestFromAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *CreateTokensTransactionRequestFromAddressRI) SetSenders(v CreateTokensTransactionRequestFromAddressRISenders)`

SetSenders sets Senders field to given value.


### GetTokenTypeSpecificData

`func (o *CreateTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificData() CreateTokensTransactionRequestFromAddressRIS`

GetTokenTypeSpecificData returns the TokenTypeSpecificData field if non-nil, zero value otherwise.

### GetTokenTypeSpecificDataOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificDataOk() (*CreateTokensTransactionRequestFromAddressRIS, bool)`

GetTokenTypeSpecificDataOk returns a tuple with the TokenTypeSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTypeSpecificData

`func (o *CreateTokensTransactionRequestFromAddressRI) SetTokenTypeSpecificData(v CreateTokensTransactionRequestFromAddressRIS)`

SetTokenTypeSpecificData sets TokenTypeSpecificData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


