# CreateCoinsTransactionRequestFromWalletRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | [**[]CreateCoinsTransactionRequestFromWalletRBDataItemDestinations**](CreateCoinsTransactionRequestFromWalletRBDataItemDestinations.md) | Defines the destination of the transaction, whether it is incoming or outgoing. | 
**FeePriority** | **string** | Represents the fee priority of the automation, whether it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 

## Methods

### NewCreateCoinsTransactionRequestFromWalletRBDataItem

`func NewCreateCoinsTransactionRequestFromWalletRBDataItem(destinations []CreateCoinsTransactionRequestFromWalletRBDataItemDestinations, feePriority string, ) *CreateCoinsTransactionRequestFromWalletRBDataItem`

NewCreateCoinsTransactionRequestFromWalletRBDataItem instantiates a new CreateCoinsTransactionRequestFromWalletRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromWalletRBDataItemWithDefaults

`func NewCreateCoinsTransactionRequestFromWalletRBDataItemWithDefaults() *CreateCoinsTransactionRequestFromWalletRBDataItem`

NewCreateCoinsTransactionRequestFromWalletRBDataItemWithDefaults instantiates a new CreateCoinsTransactionRequestFromWalletRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetDestinations() []CreateCoinsTransactionRequestFromWalletRBDataItemDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetDestinationsOk() (*[]CreateCoinsTransactionRequestFromWalletRBDataItemDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetDestinations(v []CreateCoinsTransactionRequestFromWalletRBDataItemDestinations)`

SetDestinations sets Destinations field to given value.


### GetFeePriority

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *CreateCoinsTransactionRequestFromWalletRBDataItem) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


