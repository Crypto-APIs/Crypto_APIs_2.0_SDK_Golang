# CreateCoinsTransactionRequestFromWalletRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Recipients** | [**[]CreateCoinsTransactionRequestFromWalletRIRecipients**](CreateCoinsTransactionRequestFromWalletRIRecipients.md) | Defines the destination of the transaction, whether it is incoming or outgoing. | 
**TotalTransactionAmount** | **string** | Represents the specific amount of the transaction. | 
**TransactionRequestStatus** | **string** | Defines the status of the transaction, e.g. \&quot;created, \&quot;await_approval\&quot;, \&quot;pending\&quot;, \&quot;prepared\&quot;, \&quot;signed\&quot;, \&quot;broadcasted\&quot;, \&quot;success\&quot;, \&quot;failed\&quot;, \&quot;rejected\&quot;, mined\&quot;. | 

## Methods

### NewCreateCoinsTransactionRequestFromWalletRI

`func NewCreateCoinsTransactionRequestFromWalletRI(feePriority string, recipients []CreateCoinsTransactionRequestFromWalletRIRecipients, totalTransactionAmount string, transactionRequestStatus string, ) *CreateCoinsTransactionRequestFromWalletRI`

NewCreateCoinsTransactionRequestFromWalletRI instantiates a new CreateCoinsTransactionRequestFromWalletRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromWalletRIWithDefaults

`func NewCreateCoinsTransactionRequestFromWalletRIWithDefaults() *CreateCoinsTransactionRequestFromWalletRI`

NewCreateCoinsTransactionRequestFromWalletRIWithDefaults instantiates a new CreateCoinsTransactionRequestFromWalletRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


