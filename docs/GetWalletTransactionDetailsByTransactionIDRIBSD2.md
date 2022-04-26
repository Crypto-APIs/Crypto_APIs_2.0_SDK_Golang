# GetWalletTransactionDetailsByTransactionIDRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSD2Vin**](GetWalletTransactionDetailsByTransactionIDRIBSD2Vin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSD2Vout**](GetTransactionDetailsByTransactionIDRIBSD2Vout.md) | Object Array representation of transaction outputs | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSD2

`func NewGetWalletTransactionDetailsByTransactionIDRIBSD2(locktime int64, size int32, version int32, vin []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin, vout []GetTransactionDetailsByTransactionIDRIBSD2Vout, ) *GetWalletTransactionDetailsByTransactionIDRIBSD2`

NewGetWalletTransactionDetailsByTransactionIDRIBSD2 instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSD2WithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSD2WithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSD2`

NewGetWalletTransactionDetailsByTransactionIDRIBSD2WithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSD2Vin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSD2Vin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVout() []GetTransactionDetailsByTransactionIDRIBSD2Vout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSD2Vout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD2) SetVout(v []GetTransactionDetailsByTransactionIDRIBSD2Vout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


