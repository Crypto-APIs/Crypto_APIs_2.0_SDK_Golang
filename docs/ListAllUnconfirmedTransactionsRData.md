# ListAllUnconfirmedTransactionsRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]ListAllUnconfirmedTransactionsRI**](ListAllUnconfirmedTransactionsRI.md) |  | 

## Methods

### NewListAllUnconfirmedTransactionsRData

`func NewListAllUnconfirmedTransactionsRData(limit int32, offset int32, total int32, items []ListAllUnconfirmedTransactionsRI, ) *ListAllUnconfirmedTransactionsRData`

NewListAllUnconfirmedTransactionsRData instantiates a new ListAllUnconfirmedTransactionsRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRDataWithDefaults

`func NewListAllUnconfirmedTransactionsRDataWithDefaults() *ListAllUnconfirmedTransactionsRData`

NewListAllUnconfirmedTransactionsRDataWithDefaults instantiates a new ListAllUnconfirmedTransactionsRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListAllUnconfirmedTransactionsRData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListAllUnconfirmedTransactionsRData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListAllUnconfirmedTransactionsRData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListAllUnconfirmedTransactionsRData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListAllUnconfirmedTransactionsRData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListAllUnconfirmedTransactionsRData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListAllUnconfirmedTransactionsRData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListAllUnconfirmedTransactionsRData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListAllUnconfirmedTransactionsRData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListAllUnconfirmedTransactionsRData) GetItems() []ListAllUnconfirmedTransactionsRI`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAllUnconfirmedTransactionsRData) GetItemsOk() (*[]ListAllUnconfirmedTransactionsRI, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAllUnconfirmedTransactionsRData) SetItems(v []ListAllUnconfirmedTransactionsRI)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


