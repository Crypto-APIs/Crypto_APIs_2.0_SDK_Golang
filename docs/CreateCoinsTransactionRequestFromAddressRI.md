# CreateCoinsTransactionRequestFromAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressTag** | Pointer to **int32** | Defines a specific Tag that is an additional XRP address feature. It helps identify a transaction recipient beyond a wallet address. The tag that was encoded into the x-Address along with the Source Classic Address. | [optional] 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**ClassicAddress** | Pointer to **string** | Represents the public address, which is a compressed and shortened form of a public key. The classic address is shown when the source address is an x-Address. | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**Recipients** | [**[]CreateCoinsTransactionRequestFromAddressRIRecipientsInner**](CreateCoinsTransactionRequestFromAddressRIRecipientsInner.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Senders** | [**CreateCoinsTransactionRequestFromAddressRISenders**](CreateCoinsTransactionRequestFromAddressRISenders.md) |  | 
**TransactionRequestId** | **string** | Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which &#x60;referenceId&#x60; concern that specific transaction request. | 
**TransactionRequestStatus** | **string** | Defines the status of the transaction request, e.g. \&quot;created, \&quot;await_approval\&quot;, \&quot;pending\&quot;, \&quot;prepared\&quot;, \&quot;signed\&quot;, \&quot;broadcasted\&quot;, \&quot;success\&quot;, \&quot;failed\&quot;, \&quot;rejected\&quot;, mined\&quot;. | 

## Methods

### NewCreateCoinsTransactionRequestFromAddressRI

`func NewCreateCoinsTransactionRequestFromAddressRI(feePriority string, recipients []CreateCoinsTransactionRequestFromAddressRIRecipientsInner, senders CreateCoinsTransactionRequestFromAddressRISenders, transactionRequestId string, transactionRequestStatus string, ) *CreateCoinsTransactionRequestFromAddressRI`

NewCreateCoinsTransactionRequestFromAddressRI instantiates a new CreateCoinsTransactionRequestFromAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromAddressRIWithDefaults

`func NewCreateCoinsTransactionRequestFromAddressRIWithDefaults() *CreateCoinsTransactionRequestFromAddressRI`

NewCreateCoinsTransactionRequestFromAddressRIWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressTag

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetAddressTag() int32`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetAddressTagOk() (*int32, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetAddressTag(v int32)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *CreateCoinsTransactionRequestFromAddressRI) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

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

### GetClassicAddress

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetClassicAddress() string`

GetClassicAddress returns the ClassicAddress field if non-nil, zero value otherwise.

### GetClassicAddressOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetClassicAddressOk() (*string, bool)`

GetClassicAddressOk returns a tuple with the ClassicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicAddress

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetClassicAddress(v string)`

SetClassicAddress sets ClassicAddress field to given value.

### HasClassicAddress

`func (o *CreateCoinsTransactionRequestFromAddressRI) HasClassicAddress() bool`

HasClassicAddress returns a boolean if a field has been set.

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


### GetNote

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateCoinsTransactionRequestFromAddressRI) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipients

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetRecipients() []CreateCoinsTransactionRequestFromAddressRIRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetRecipientsOk() (*[]CreateCoinsTransactionRequestFromAddressRIRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetRecipients(v []CreateCoinsTransactionRequestFromAddressRIRecipientsInner)`

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


### GetTransactionRequestId

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestId() string`

GetTransactionRequestId returns the TransactionRequestId field if non-nil, zero value otherwise.

### GetTransactionRequestIdOk

`func (o *CreateCoinsTransactionRequestFromAddressRI) GetTransactionRequestIdOk() (*string, bool)`

GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestId

`func (o *CreateCoinsTransactionRequestFromAddressRI) SetTransactionRequestId(v string)`

SetTransactionRequestId sets TransactionRequestId field to given value.


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


