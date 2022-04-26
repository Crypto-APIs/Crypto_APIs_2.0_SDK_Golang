# ListAssetsDetailsRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]ListAssetsDetailsRI**](ListAssetsDetailsRI.md) |  | 

## Methods

### NewListAssetsDetailsRData

`func NewListAssetsDetailsRData(limit int32, offset int32, total int32, items []ListAssetsDetailsRI, ) *ListAssetsDetailsRData`

NewListAssetsDetailsRData instantiates a new ListAssetsDetailsRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetsDetailsRDataWithDefaults

`func NewListAssetsDetailsRDataWithDefaults() *ListAssetsDetailsRData`

NewListAssetsDetailsRDataWithDefaults instantiates a new ListAssetsDetailsRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListAssetsDetailsRData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListAssetsDetailsRData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListAssetsDetailsRData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListAssetsDetailsRData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListAssetsDetailsRData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListAssetsDetailsRData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListAssetsDetailsRData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListAssetsDetailsRData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListAssetsDetailsRData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListAssetsDetailsRData) GetItems() []ListAssetsDetailsRI`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAssetsDetailsRData) GetItemsOk() (*[]ListAssetsDetailsRI, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAssetsDetailsRData) SetItems(v []ListAssetsDetailsRI)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


