# GetTransactionDetailsByTransactionIDRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]GetTransactionDetailsByTransactionIDRIBSD2Vin**](GetTransactionDetailsByTransactionIDRIBSD2Vin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSD2Vout**](GetTransactionDetailsByTransactionIDRIBSD2Vout.md) | Represents the transaction outputs. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSD2

`func NewGetTransactionDetailsByTransactionIDRIBSD2(locktime int32, size int32, version int32, vin []GetTransactionDetailsByTransactionIDRIBSD2Vin, vout []GetTransactionDetailsByTransactionIDRIBSD2Vout, ) *GetTransactionDetailsByTransactionIDRIBSD2`

NewGetTransactionDetailsByTransactionIDRIBSD2 instantiates a new GetTransactionDetailsByTransactionIDRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSD2WithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSD2WithDefaults() *GetTransactionDetailsByTransactionIDRIBSD2`

NewGetTransactionDetailsByTransactionIDRIBSD2WithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetVin() []GetTransactionDetailsByTransactionIDRIBSD2Vin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetVinOk() (*[]GetTransactionDetailsByTransactionIDRIBSD2Vin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) SetVin(v []GetTransactionDetailsByTransactionIDRIBSD2Vin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetVout() []GetTransactionDetailsByTransactionIDRIBSD2Vout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSD2Vout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSD2) SetVout(v []GetTransactionDetailsByTransactionIDRIBSD2Vout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


