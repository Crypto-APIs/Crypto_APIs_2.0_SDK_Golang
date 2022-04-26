# ListInternalTransactionsByAddressAndTimeRangeRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]ListInternalTransactionsByAddressAndTimeRangeRI**](ListInternalTransactionsByAddressAndTimeRangeRI.md) |  | 

## Methods

### NewListInternalTransactionsByAddressAndTimeRangeRData

`func NewListInternalTransactionsByAddressAndTimeRangeRData(limit int32, offset int32, total int32, items []ListInternalTransactionsByAddressAndTimeRangeRI, ) *ListInternalTransactionsByAddressAndTimeRangeRData`

NewListInternalTransactionsByAddressAndTimeRangeRData instantiates a new ListInternalTransactionsByAddressAndTimeRangeRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInternalTransactionsByAddressAndTimeRangeRDataWithDefaults

`func NewListInternalTransactionsByAddressAndTimeRangeRDataWithDefaults() *ListInternalTransactionsByAddressAndTimeRangeRData`

NewListInternalTransactionsByAddressAndTimeRangeRDataWithDefaults instantiates a new ListInternalTransactionsByAddressAndTimeRangeRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetItems() []ListInternalTransactionsByAddressAndTimeRangeRI`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) GetItemsOk() (*[]ListInternalTransactionsByAddressAndTimeRangeRI, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListInternalTransactionsByAddressAndTimeRangeRData) SetItems(v []ListInternalTransactionsByAddressAndTimeRangeRI)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


