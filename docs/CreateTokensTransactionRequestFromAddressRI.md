# CreateTokensTransactionRequestFromAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Recipients** | [**[]CreateTokensTransactionRequestFromAddressRIRecipients**](CreateTokensTransactionRequestFromAddressRIRecipients.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Senders** | [**CreateTokensTransactionRequestFromAddressRISenders**](CreateTokensTransactionRequestFromAddressRISenders.md) |  | 
**TokenTypeSpecificData** | [**CreateTokensTransactionRequestFromAddressRITokenTypeSpecificData**](CreateTokensTransactionRequestFromAddressRITokenTypeSpecificData.md) |  | 

## Methods

### NewCreateTokensTransactionRequestFromAddressRI

`func NewCreateTokensTransactionRequestFromAddressRI(feePriority string, recipients []CreateTokensTransactionRequestFromAddressRIRecipients, senders CreateTokensTransactionRequestFromAddressRISenders, tokenTypeSpecificData CreateTokensTransactionRequestFromAddressRITokenTypeSpecificData, ) *CreateTokensTransactionRequestFromAddressRI`

NewCreateTokensTransactionRequestFromAddressRI instantiates a new CreateTokensTransactionRequestFromAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokensTransactionRequestFromAddressRIWithDefaults

`func NewCreateTokensTransactionRequestFromAddressRIWithDefaults() *CreateTokensTransactionRequestFromAddressRI`

NewCreateTokensTransactionRequestFromAddressRIWithDefaults instantiates a new CreateTokensTransactionRequestFromAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *CreateTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificData() CreateTokensTransactionRequestFromAddressRITokenTypeSpecificData`

GetTokenTypeSpecificData returns the TokenTypeSpecificData field if non-nil, zero value otherwise.

### GetTokenTypeSpecificDataOk

`func (o *CreateTokensTransactionRequestFromAddressRI) GetTokenTypeSpecificDataOk() (*CreateTokensTransactionRequestFromAddressRITokenTypeSpecificData, bool)`

GetTokenTypeSpecificDataOk returns a tuple with the TokenTypeSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTypeSpecificData

`func (o *CreateTokensTransactionRequestFromAddressRI) SetTokenTypeSpecificData(v CreateTokensTransactionRequestFromAddressRITokenTypeSpecificData)`

SetTokenTypeSpecificData sets TokenTypeSpecificData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


