# GetWalletTransactionDetailsByTransactionIDRIBSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSDVin**](GetWalletTransactionDetailsByTransactionIDRIBSDVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSDVout**](GetTransactionDetailsByTransactionIDRIBSDVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSD

`func NewGetWalletTransactionDetailsByTransactionIDRIBSD(locktime int32, size int32, version int32, vin []GetWalletTransactionDetailsByTransactionIDRIBSDVin, vout []GetTransactionDetailsByTransactionIDRIBSDVout, ) *GetWalletTransactionDetailsByTransactionIDRIBSD`

NewGetWalletTransactionDetailsByTransactionIDRIBSD instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSDWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSDWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSD`

NewGetWalletTransactionDetailsByTransactionIDRIBSDWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSDVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSDVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSDVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetVout() []GetTransactionDetailsByTransactionIDRIBSDVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSD) SetVout(v []GetTransactionDetailsByTransactionIDRIBSDVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


