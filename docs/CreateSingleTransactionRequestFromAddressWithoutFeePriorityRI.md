# CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**ClassicAddress** | Pointer to **string** | Represents the public address, which is a compressed and shortened form of a public key. The classic address is shown when the source address is an x-Address. | [optional] 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**Recipient** | [**[]CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner**](CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner.md) | Defines the destination for the transaction, i.e. the recipient(s). | 
**Sender** | [**CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender**](CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender.md) |  | 
**TransactionRequestId** | **string** | Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which &#x60;referenceId&#x60; concern that specific transaction request. | 
**TransactionRequestStatus** | **string** | Defines the status of the transaction, e.g. \&quot;created, \&quot;await_approval\&quot;, \&quot;pending\&quot;, \&quot;prepared\&quot;, \&quot;signed\&quot;, \&quot;broadcasted\&quot;, \&quot;success\&quot;, \&quot;failed\&quot;, \&quot;rejected\&quot;, mined\&quot;. | 
**TotalAmount** | Pointer to [**CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount**](CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount.md) |  | [optional] 

## Methods

### NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRI

`func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRI(recipient []CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner, sender CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender, transactionRequestId string, transactionRequestStatus string, ) *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI`

NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRI instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIWithDefaults

`func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIWithDefaults() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI`

NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIWithDefaults instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetClassicAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetClassicAddress() string`

GetClassicAddress returns the ClassicAddress field if non-nil, zero value otherwise.

### GetClassicAddressOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetClassicAddressOk() (*string, bool)`

GetClassicAddressOk returns a tuple with the ClassicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetClassicAddress(v string)`

SetClassicAddress sets ClassicAddress field to given value.

### HasClassicAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) HasClassicAddress() bool`

HasClassicAddress returns a boolean if a field has been set.

### GetNote

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipient

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetRecipient() []CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetRecipientOk() (*[]CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetRecipient(v []CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetSender() CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetSenderOk() (*CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetSender(v CreateSingleTransactionRequestFromAddressWithoutFeePriorityRISender)`

SetSender sets Sender field to given value.


### GetTransactionRequestId

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetTransactionRequestId() string`

GetTransactionRequestId returns the TransactionRequestId field if non-nil, zero value otherwise.

### GetTransactionRequestIdOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetTransactionRequestIdOk() (*string, bool)`

GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestId

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetTransactionRequestId(v string)`

SetTransactionRequestId sets TransactionRequestId field to given value.


### GetTransactionRequestStatus

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetTransactionRequestStatus() string`

GetTransactionRequestStatus returns the TransactionRequestStatus field if non-nil, zero value otherwise.

### GetTransactionRequestStatusOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetTransactionRequestStatusOk() (*string, bool)`

GetTransactionRequestStatusOk returns a tuple with the TransactionRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestStatus

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetTransactionRequestStatus(v string)`

SetTransactionRequestStatus sets TransactionRequestStatus field to given value.


### GetTotalAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetTotalAmount() CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) GetTotalAmountOk() (*CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) SetTotalAmount(v CreateSingleTransactionRequestFromAddressWithoutFeePriorityRITotalAmount)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRI) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


