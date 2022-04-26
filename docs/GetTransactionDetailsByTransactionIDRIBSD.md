# GetTransactionDetailsByTransactionIDRIBSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]GetTransactionDetailsByTransactionIDRIBSDVin**](GetTransactionDetailsByTransactionIDRIBSDVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSDVout**](GetTransactionDetailsByTransactionIDRIBSDVout.md) | Represents the transaction outputs. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSD

`func NewGetTransactionDetailsByTransactionIDRIBSD(locktime int64, size int32, version int32, vin []GetTransactionDetailsByTransactionIDRIBSDVin, vout []GetTransactionDetailsByTransactionIDRIBSDVout, ) *GetTransactionDetailsByTransactionIDRIBSD`

NewGetTransactionDetailsByTransactionIDRIBSD instantiates a new GetTransactionDetailsByTransactionIDRIBSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSDWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSDWithDefaults() *GetTransactionDetailsByTransactionIDRIBSD`

NewGetTransactionDetailsByTransactionIDRIBSDWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetTransactionDetailsByTransactionIDRIBSD) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTransactionDetailsByTransactionIDRIBSD) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTransactionDetailsByTransactionIDRIBSD) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetVin() []GetTransactionDetailsByTransactionIDRIBSDVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetVinOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetTransactionDetailsByTransactionIDRIBSD) SetVin(v []GetTransactionDetailsByTransactionIDRIBSDVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetVout() []GetTransactionDetailsByTransactionIDRIBSDVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSD) SetVout(v []GetTransactionDetailsByTransactionIDRIBSDVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


