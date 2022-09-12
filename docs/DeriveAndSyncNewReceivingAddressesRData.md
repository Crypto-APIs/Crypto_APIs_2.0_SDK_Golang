# DeriveAndSyncNewReceivingAddressesRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]DeriveAndSyncNewReceivingAddressesRI**](DeriveAndSyncNewReceivingAddressesRI.md) |  | 

## Methods

### NewDeriveAndSyncNewReceivingAddressesRData

`func NewDeriveAndSyncNewReceivingAddressesRData(limit int32, offset int32, total int32, items []DeriveAndSyncNewReceivingAddressesRI, ) *DeriveAndSyncNewReceivingAddressesRData`

NewDeriveAndSyncNewReceivingAddressesRData instantiates a new DeriveAndSyncNewReceivingAddressesRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeriveAndSyncNewReceivingAddressesRDataWithDefaults

`func NewDeriveAndSyncNewReceivingAddressesRDataWithDefaults() *DeriveAndSyncNewReceivingAddressesRData`

NewDeriveAndSyncNewReceivingAddressesRDataWithDefaults instantiates a new DeriveAndSyncNewReceivingAddressesRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DeriveAndSyncNewReceivingAddressesRData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DeriveAndSyncNewReceivingAddressesRData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DeriveAndSyncNewReceivingAddressesRData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetItems() []DeriveAndSyncNewReceivingAddressesRI`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeriveAndSyncNewReceivingAddressesRData) GetItemsOk() (*[]DeriveAndSyncNewReceivingAddressesRI, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeriveAndSyncNewReceivingAddressesRData) SetItems(v []DeriveAndSyncNewReceivingAddressesRI)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


