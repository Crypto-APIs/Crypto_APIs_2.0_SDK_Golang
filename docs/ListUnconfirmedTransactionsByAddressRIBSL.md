# ListUnconfirmedTransactionsByAddressRIBSL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListUnconfirmedTransactionsByAddressRIBSLVinInner**](ListUnconfirmedTransactionsByAddressRIBSLVinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListUnconfirmedTransactionsByAddressRIBSLVoutInner**](ListUnconfirmedTransactionsByAddressRIBSLVoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSL

`func NewListUnconfirmedTransactionsByAddressRIBSL(locktime int64, size int32, vSize int32, version int32, vin []ListUnconfirmedTransactionsByAddressRIBSLVinInner, vout []ListUnconfirmedTransactionsByAddressRIBSLVoutInner, ) *ListUnconfirmedTransactionsByAddressRIBSL`

NewListUnconfirmedTransactionsByAddressRIBSL instantiates a new ListUnconfirmedTransactionsByAddressRIBSL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSLWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSLWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSL`

NewListUnconfirmedTransactionsByAddressRIBSLWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVin() []ListUnconfirmedTransactionsByAddressRIBSLVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVinOk() (*[]ListUnconfirmedTransactionsByAddressRIBSLVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) SetVin(v []ListUnconfirmedTransactionsByAddressRIBSLVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVout() []ListUnconfirmedTransactionsByAddressRIBSLVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) GetVoutOk() (*[]ListUnconfirmedTransactionsByAddressRIBSLVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSL) SetVout(v []ListUnconfirmedTransactionsByAddressRIBSLVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


