# ListTransactionsByBlockHashRIBSL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHashRIBSLVin**](ListTransactionsByBlockHashRIBSLVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashRIBSLVout**](ListTransactionsByBlockHashRIBSLVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHashRIBSL

`func NewListTransactionsByBlockHashRIBSL(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHashRIBSLVin, vout []ListTransactionsByBlockHashRIBSLVout, ) *ListTransactionsByBlockHashRIBSL`

NewListTransactionsByBlockHashRIBSL instantiates a new ListTransactionsByBlockHashRIBSL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSLWithDefaults

`func NewListTransactionsByBlockHashRIBSLWithDefaults() *ListTransactionsByBlockHashRIBSL`

NewListTransactionsByBlockHashRIBSLWithDefaults instantiates a new ListTransactionsByBlockHashRIBSL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHashRIBSL) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHashRIBSL) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHashRIBSL) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHashRIBSL) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHashRIBSL) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHashRIBSL) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByBlockHashRIBSL) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByBlockHashRIBSL) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByBlockHashRIBSL) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHashRIBSL) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHashRIBSL) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHashRIBSL) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHashRIBSL) GetVin() []ListTransactionsByBlockHashRIBSLVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHashRIBSL) GetVinOk() (*[]ListTransactionsByBlockHashRIBSLVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHashRIBSL) SetVin(v []ListTransactionsByBlockHashRIBSLVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashRIBSL) GetVout() []ListTransactionsByBlockHashRIBSLVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashRIBSL) GetVoutOk() (*[]ListTransactionsByBlockHashRIBSLVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashRIBSL) SetVout(v []ListTransactionsByBlockHashRIBSLVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


