# TransactionRequestApproveData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** | Represents the Crypto APIs 2.0 product which sends the callback. | 
**Event** | **string** | Defines the specific event, for which a callback subscription is set. | 
**Item** | [**TransactionRequestApproveDataItem**](TransactionRequestApproveDataItem.md) |  | 

## Methods

### NewTransactionRequestApproveData

`func NewTransactionRequestApproveData(product string, event string, item TransactionRequestApproveDataItem, ) *TransactionRequestApproveData`

NewTransactionRequestApproveData instantiates a new TransactionRequestApproveData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestApproveDataWithDefaults

`func NewTransactionRequestApproveDataWithDefaults() *TransactionRequestApproveData`

NewTransactionRequestApproveDataWithDefaults instantiates a new TransactionRequestApproveData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *TransactionRequestApproveData) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *TransactionRequestApproveData) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *TransactionRequestApproveData) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetEvent

`func (o *TransactionRequestApproveData) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TransactionRequestApproveData) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TransactionRequestApproveData) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetItem

`func (o *TransactionRequestApproveData) GetItem() TransactionRequestApproveDataItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TransactionRequestApproveData) GetItemOk() (*TransactionRequestApproveDataItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TransactionRequestApproveData) SetItem(v TransactionRequestApproveDataItem)`

SetItem sets Item field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


