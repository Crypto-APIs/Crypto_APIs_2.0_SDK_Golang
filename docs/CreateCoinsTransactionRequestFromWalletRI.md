# CreateCoinsTransactionRequestFromWalletRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs.  For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | Pointer to **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. &#x60;We support ONLY httpS type of protocol&#x60;. | [optional] 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Note** | Pointer to **string** | Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request. | [optional] 
**Recipients** | [**[]CreateCoinsTransactionRequestFromWalletRIRecipients**](CreateCoinsTransactionRequestFromWalletRIRecipients.md) | Defines the destination of the transaction, whether it is incoming or outgoing. | 
**TotalTransactionAmount** | **string** | Represents the specific amount of the transaction. | 
**TransactionRequestId** | **string** | Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which &#x60;referenceId&#x60; concern that specific transaction request. | 
**TransactionRequestStatus** | **string** | Defines the status of the transaction, e.g. \&quot;created, \&quot;await_approval\&quot;, \&quot;pending\&quot;, \&quot;prepared\&quot;, \&quot;signed\&quot;, \&quot;broadcasted\&quot;, \&quot;success\&quot;, \&quot;failed\&quot;, \&quot;rejected\&quot;, mined\&quot;. | 

## Methods

### NewCreateCoinsTransactionRequestFromWalletRI

`func NewCreateCoinsTransactionRequestFromWalletRI(feePriority string, recipients []CreateCoinsTransactionRequestFromWalletRIRecipients, totalTransactionAmount string, transactionRequestId string, transactionRequestStatus string, ) *CreateCoinsTransactionRequestFromWalletRI`

NewCreateCoinsTransactionRequestFromWalletRI instantiates a new CreateCoinsTransactionRequestFromWalletRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromWalletRIWithDefaults

`func NewCreateCoinsTransactionRequestFromWalletRIWithDefaults() *CreateCoinsTransactionRequestFromWalletRI`

NewCreateCoinsTransactionRequestFromWalletRIWithDefaults instantiates a new CreateCoinsTransactionRequestFromWalletRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *CreateCoinsTransactionRequestFromWalletRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateCoinsTransactionRequestFromWalletRI) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetFeePriority

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetNote

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateCoinsTransactionRequestFromWalletRI) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRecipients

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetRecipients() []CreateCoinsTransactionRequestFromWalletRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetRecipientsOk() (*[]CreateCoinsTransactionRequestFromWalletRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetRecipients(v []CreateCoinsTransactionRequestFromWalletRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetTotalTransactionAmount

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetTotalTransactionAmount() string`

GetTotalTransactionAmount returns the TotalTransactionAmount field if non-nil, zero value otherwise.

### GetTotalTransactionAmountOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetTotalTransactionAmountOk() (*string, bool)`

GetTotalTransactionAmountOk returns a tuple with the TotalTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactionAmount

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetTotalTransactionAmount(v string)`

SetTotalTransactionAmount sets TotalTransactionAmount field to given value.


### GetTransactionRequestId

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetTransactionRequestId() string`

GetTransactionRequestId returns the TransactionRequestId field if non-nil, zero value otherwise.

### GetTransactionRequestIdOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetTransactionRequestIdOk() (*string, bool)`

GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestId

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetTransactionRequestId(v string)`

SetTransactionRequestId sets TransactionRequestId field to given value.


### GetTransactionRequestStatus

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetTransactionRequestStatus() string`

GetTransactionRequestStatus returns the TransactionRequestStatus field if non-nil, zero value otherwise.

### GetTransactionRequestStatusOk

`func (o *CreateCoinsTransactionRequestFromWalletRI) GetTransactionRequestStatusOk() (*string, bool)`

GetTransactionRequestStatusOk returns a tuple with the TransactionRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestStatus

`func (o *CreateCoinsTransactionRequestFromWalletRI) SetTransactionRequestStatus(v string)`

SetTransactionRequestStatus sets TransactionRequestStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


