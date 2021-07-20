# CreateCoinsTransactionRequestFromAddressRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents the specific amount of the transaction. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**ToAddress** | **string** | Defines the specific recipient address for the transaction. | 

## Methods

### NewCreateCoinsTransactionRequestFromAddressRBDataItem

`func NewCreateCoinsTransactionRequestFromAddressRBDataItem(amount string, feePriority string, toAddress string, ) *CreateCoinsTransactionRequestFromAddressRBDataItem`

NewCreateCoinsTransactionRequestFromAddressRBDataItem instantiates a new CreateCoinsTransactionRequestFromAddressRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults

`func NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults() *CreateCoinsTransactionRequestFromAddressRBDataItem`

NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeePriority

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetToAddress

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


