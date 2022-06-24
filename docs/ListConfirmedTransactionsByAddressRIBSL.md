# ListConfirmedTransactionsByAddressRIBSL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListConfirmedTransactionsByAddressRIBSLVinInner**](ListConfirmedTransactionsByAddressRIBSLVinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSLVoutInner**](GetTransactionDetailsByTransactionIDRIBSLVoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSL

`func NewListConfirmedTransactionsByAddressRIBSL(locktime int64, size int32, vSize int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSLVinInner, vout []GetTransactionDetailsByTransactionIDRIBSLVoutInner, ) *ListConfirmedTransactionsByAddressRIBSL`

NewListConfirmedTransactionsByAddressRIBSL instantiates a new ListConfirmedTransactionsByAddressRIBSL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSLWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSLWithDefaults() *ListConfirmedTransactionsByAddressRIBSL`

NewListConfirmedTransactionsByAddressRIBSLWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListConfirmedTransactionsByAddressRIBSL) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListConfirmedTransactionsByAddressRIBSL) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListConfirmedTransactionsByAddressRIBSL) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConfirmedTransactionsByAddressRIBSL) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVin() []ListConfirmedTransactionsByAddressRIBSLVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVinOk() (*[]ListConfirmedTransactionsByAddressRIBSLVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListConfirmedTransactionsByAddressRIBSL) SetVin(v []ListConfirmedTransactionsByAddressRIBSLVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVout() []GetTransactionDetailsByTransactionIDRIBSLVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBSL) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSLVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBSL) SetVout(v []GetTransactionDetailsByTransactionIDRIBSLVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


