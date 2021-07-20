# TransactionRequestApproveDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** |  | 
**Network** | **string** |  | 
**RequiredApproves** | **int32** |  | 
**RequiredRejects** | **int32** |  | 
**CurrentApproves** | **int32** |  | 
**CurrentRejects** | **int32** |  | 

## Methods

### NewTransactionRequestApproveDataItem

`func NewTransactionRequestApproveDataItem(blockchain string, network string, requiredApproves int32, requiredRejects int32, currentApproves int32, currentRejects int32, ) *TransactionRequestApproveDataItem`

NewTransactionRequestApproveDataItem instantiates a new TransactionRequestApproveDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestApproveDataItemWithDefaults

`func NewTransactionRequestApproveDataItemWithDefaults() *TransactionRequestApproveDataItem`

NewTransactionRequestApproveDataItemWithDefaults instantiates a new TransactionRequestApproveDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *TransactionRequestApproveDataItem) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *TransactionRequestApproveDataItem) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *TransactionRequestApproveDataItem) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetNetwork

`func (o *TransactionRequestApproveDataItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionRequestApproveDataItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionRequestApproveDataItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRequiredApproves

`func (o *TransactionRequestApproveDataItem) GetRequiredApproves() int32`

GetRequiredApproves returns the RequiredApproves field if non-nil, zero value otherwise.

### GetRequiredApprovesOk

`func (o *TransactionRequestApproveDataItem) GetRequiredApprovesOk() (*int32, bool)`

GetRequiredApprovesOk returns a tuple with the RequiredApproves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApproves

`func (o *TransactionRequestApproveDataItem) SetRequiredApproves(v int32)`

SetRequiredApproves sets RequiredApproves field to given value.


### GetRequiredRejects

`func (o *TransactionRequestApproveDataItem) GetRequiredRejects() int32`

GetRequiredRejects returns the RequiredRejects field if non-nil, zero value otherwise.

### GetRequiredRejectsOk

`func (o *TransactionRequestApproveDataItem) GetRequiredRejectsOk() (*int32, bool)`

GetRequiredRejectsOk returns a tuple with the RequiredRejects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRejects

`func (o *TransactionRequestApproveDataItem) SetRequiredRejects(v int32)`

SetRequiredRejects sets RequiredRejects field to given value.


### GetCurrentApproves

`func (o *TransactionRequestApproveDataItem) GetCurrentApproves() int32`

GetCurrentApproves returns the CurrentApproves field if non-nil, zero value otherwise.

### GetCurrentApprovesOk

`func (o *TransactionRequestApproveDataItem) GetCurrentApprovesOk() (*int32, bool)`

GetCurrentApprovesOk returns a tuple with the CurrentApproves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentApproves

`func (o *TransactionRequestApproveDataItem) SetCurrentApproves(v int32)`

SetCurrentApproves sets CurrentApproves field to given value.


### GetCurrentRejects

`func (o *TransactionRequestApproveDataItem) GetCurrentRejects() int32`

GetCurrentRejects returns the CurrentRejects field if non-nil, zero value otherwise.

### GetCurrentRejectsOk

`func (o *TransactionRequestApproveDataItem) GetCurrentRejectsOk() (*int32, bool)`

GetCurrentRejectsOk returns a tuple with the CurrentRejects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRejects

`func (o *TransactionRequestApproveDataItem) SetCurrentRejects(v int32)`

SetCurrentRejects sets CurrentRejects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


