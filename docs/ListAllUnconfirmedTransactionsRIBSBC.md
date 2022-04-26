# ListAllUnconfirmedTransactionsRIBSBC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListAllUnconfirmedTransactionsRIBSBCVin**](ListAllUnconfirmedTransactionsRIBSBCVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListAllUnconfirmedTransactionsRIBSBCVout**](ListAllUnconfirmedTransactionsRIBSBCVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSBC

`func NewListAllUnconfirmedTransactionsRIBSBC(locktime int64, size int32, version int32, vin []ListAllUnconfirmedTransactionsRIBSBCVin, vout []ListAllUnconfirmedTransactionsRIBSBCVout, ) *ListAllUnconfirmedTransactionsRIBSBC`

NewListAllUnconfirmedTransactionsRIBSBC instantiates a new ListAllUnconfirmedTransactionsRIBSBC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSBCWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSBCWithDefaults() *ListAllUnconfirmedTransactionsRIBSBC`

NewListAllUnconfirmedTransactionsRIBSBCWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSBC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBSBC) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListAllUnconfirmedTransactionsRIBSBC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListAllUnconfirmedTransactionsRIBSBC) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetVin() []ListAllUnconfirmedTransactionsRIBSBCVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetVinOk() (*[]ListAllUnconfirmedTransactionsRIBSBCVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListAllUnconfirmedTransactionsRIBSBC) SetVin(v []ListAllUnconfirmedTransactionsRIBSBCVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetVout() []ListAllUnconfirmedTransactionsRIBSBCVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSBC) GetVoutOk() (*[]ListAllUnconfirmedTransactionsRIBSBCVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSBC) SetVout(v []ListAllUnconfirmedTransactionsRIBSBCVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


