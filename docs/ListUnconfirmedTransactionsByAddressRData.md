# ListUnconfirmedTransactionsByAddressRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]ListUnconfirmedTransactionsByAddressRI**](ListUnconfirmedTransactionsByAddressRI.md) |  | 

## Methods

### NewListUnconfirmedTransactionsByAddressRData

`func NewListUnconfirmedTransactionsByAddressRData(offset int32, limit int32, total int32, items []ListUnconfirmedTransactionsByAddressRI, ) *ListUnconfirmedTransactionsByAddressRData`

NewListUnconfirmedTransactionsByAddressRData instantiates a new ListUnconfirmedTransactionsByAddressRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRDataWithDefaults

`func NewListUnconfirmedTransactionsByAddressRDataWithDefaults() *ListUnconfirmedTransactionsByAddressRData`

NewListUnconfirmedTransactionsByAddressRDataWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *ListUnconfirmedTransactionsByAddressRData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListUnconfirmedTransactionsByAddressRData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListUnconfirmedTransactionsByAddressRData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ListUnconfirmedTransactionsByAddressRData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListUnconfirmedTransactionsByAddressRData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListUnconfirmedTransactionsByAddressRData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetTotal

`func (o *ListUnconfirmedTransactionsByAddressRData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListUnconfirmedTransactionsByAddressRData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListUnconfirmedTransactionsByAddressRData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListUnconfirmedTransactionsByAddressRData) GetItems() []ListUnconfirmedTransactionsByAddressRI`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListUnconfirmedTransactionsByAddressRData) GetItemsOk() (*[]ListUnconfirmedTransactionsByAddressRI, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListUnconfirmedTransactionsByAddressRData) SetItems(v []ListUnconfirmedTransactionsByAddressRI)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


