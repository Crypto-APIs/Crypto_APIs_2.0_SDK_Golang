# CreateCoinsTransactionRequestFromAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. | [optional] 
**CallbackUrl** | Pointer to **string** | Verified URL for sending callbacks | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Recipients** | [**[]CreateCoinsTransactionRequestFromAddressRIRecipients**](CreateCoinsTransactionRequestFromAddressRIRecipients.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Senders** | [**CreateCoinsTransactionRequestFromAddressRISenders**](CreateCoinsTransactionRequestFromAddressRISenders.md) |  | 
**TransactionRequestStatus** | **string** | Defines the status of the transaction request, e.g. \&quot;created, \&quot;await_approval\&quot;, \&quot;pending\&quot;, \&quot;prepared\&quot;, \&quot;signed\&quot;, \&quot;broadcasted\&quot;, \&quot;success\&quot;, \&quot;failed\&quot;, \&quot;rejected\&quot;, mined\&quot;. | 

## Methods

### NewCreateCoinsTransactionRequestFromAddressRI

`func NewCreateCoinsTransactionRequestFromAddressRI(feePriority string, recipients []CreateCoinsTransactionRequestFromAddressRIRecipients, senders CreateCoinsTransactionRequestFromAddressRISenders, transactionRequestStatus string, ) *CreateCoinsTransactionRequestFromAddressRI`

NewCreateCoinsTransactionRequestFromAddressRI instantiates a new CreateCoinsTransactionRequestFromAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromAddressRIWithDefaults

`func NewCreateCoinsTransactionRequestFromAddressRIWithDefaults() *CreateCoinsTransactionRequestFromAddressRI`

NewCreateCoinsTransactionRequestFromAddressRIWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromAddressRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateCoinsTransactionRequestFromAddressRI) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeePriority

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetRecipients

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetRecipients() []CreateCoinsTransactionRequestFromAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetRecipientsOk() (*[]CreateCoinsTransactionRequestFromAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetRecipients(v []CreateCoinsTransactionRequestFromAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetSenders() CreateCoinsTransactionRequestFromAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetSendersOk() (*CreateCoinsTransactionRequestFromAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetSenders(v CreateCoinsTransactionRequestFromAddressRISenders)`

SetSenders sets Senders field to given value.


### GetTransactionRequestStatus

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestStatus() string`

GetTransactionRequestStatus returns the TransactionRequestStatus field if non-nil, zero value otherwise.

### GetTransactionRequestStatusOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestStatusOk() (*string, bool)`

GetTransactionRequestStatusOk returns a tuple with the TransactionRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestStatus

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetTransactionRequestStatus(v string)`

SetTransactionRequestStatus sets TransactionRequestStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


