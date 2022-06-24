# GetWalletTransactionDetailsByTransactionIDRIBSBC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSBCVinInner**](GetWalletTransactionDetailsByTransactionIDRIBSBCVinInner.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSBCVoutInner**](GetWalletTransactionDetailsByTransactionIDRIBSBCVoutInner.md) | Object Array representation of transaction outputs | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSBC

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBC(locktime int64, size int32, version int32, vin []GetWalletTransactionDetailsByTransactionIDRIBSBCVinInner, vout []GetWalletTransactionDetailsByTransactionIDRIBSBCVoutInner, ) *GetWalletTransactionDetailsByTransactionIDRIBSBC`

NewGetWalletTransactionDetailsByTransactionIDRIBSBC instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSBCWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBCWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSBC`

NewGetWalletTransactionDetailsByTransactionIDRIBSBCWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSBCVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSBCVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSBCVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetVout() []GetWalletTransactionDetailsByTransactionIDRIBSBCVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) GetVoutOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSBCVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBC) SetVout(v []GetWalletTransactionDetailsByTransactionIDRIBSBCVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


