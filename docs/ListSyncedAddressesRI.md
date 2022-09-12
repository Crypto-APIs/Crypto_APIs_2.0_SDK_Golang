# ListSyncedAddressesRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the address. | 
**Index** | **int64** | Represents the index position of the transaction in the specific block. | 

## Methods

### NewListSyncedAddressesRI

`func NewListSyncedAddressesRI(address string, index int64, ) *ListSyncedAddressesRI`

NewListSyncedAddressesRI instantiates a new ListSyncedAddressesRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncedAddressesRIWithDefaults

`func NewListSyncedAddressesRIWithDefaults() *ListSyncedAddressesRI`

NewListSyncedAddressesRIWithDefaults instantiates a new ListSyncedAddressesRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListSyncedAddressesRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListSyncedAddressesRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListSyncedAddressesRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIndex

`func (o *ListSyncedAddressesRI) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListSyncedAddressesRI) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListSyncedAddressesRI) SetIndex(v int64)`

SetIndex sets Index field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


