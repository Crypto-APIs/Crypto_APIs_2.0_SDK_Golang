# GetWalletTransactionDetailsByTransactionIDRIBSB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSBVinInner**](GetWalletTransactionDetailsByTransactionIDRIBSBVinInner.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner**](GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSB

`func NewGetWalletTransactionDetailsByTransactionIDRIBSB(locktime int64, size int32, vSize int32, version int32, vin []GetWalletTransactionDetailsByTransactionIDRIBSBVinInner, vout []GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner, ) *GetWalletTransactionDetailsByTransactionIDRIBSB`

NewGetWalletTransactionDetailsByTransactionIDRIBSB instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSBWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSB`

NewGetWalletTransactionDetailsByTransactionIDRIBSBWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSBVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSBVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSBVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVout() []GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) GetVoutOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSB) SetVout(v []GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


