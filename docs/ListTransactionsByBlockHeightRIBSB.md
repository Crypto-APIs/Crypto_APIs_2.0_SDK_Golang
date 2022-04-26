# ListTransactionsByBlockHeightRIBSB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHashRIBSBVin**](ListTransactionsByBlockHashRIBSBVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHeightRIBSBVout**](ListTransactionsByBlockHeightRIBSBVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHeightRIBSB

`func NewListTransactionsByBlockHeightRIBSB(locktime int64, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHashRIBSBVin, vout []ListTransactionsByBlockHeightRIBSBVout, ) *ListTransactionsByBlockHeightRIBSB`

NewListTransactionsByBlockHeightRIBSB instantiates a new ListTransactionsByBlockHeightRIBSB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSBWithDefaults

`func NewListTransactionsByBlockHeightRIBSBWithDefaults() *ListTransactionsByBlockHeightRIBSB`

NewListTransactionsByBlockHeightRIBSBWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightRIBSB) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightRIBSB) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightRIBSB) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightRIBSB) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightRIBSB) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightRIBSB) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByBlockHeightRIBSB) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByBlockHeightRIBSB) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByBlockHeightRIBSB) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightRIBSB) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightRIBSB) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightRIBSB) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightRIBSB) GetVin() []ListTransactionsByBlockHashRIBSBVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightRIBSB) GetVinOk() (*[]ListTransactionsByBlockHashRIBSBVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightRIBSB) SetVin(v []ListTransactionsByBlockHashRIBSBVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightRIBSB) GetVout() []ListTransactionsByBlockHeightRIBSBVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightRIBSB) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSBVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightRIBSB) SetVout(v []ListTransactionsByBlockHeightRIBSBVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


