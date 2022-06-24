# GetTransactionDetailsByTransactionIDFromCallbackRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VinInner**](GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner**](GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2(locktime int64, size int32, version int32, vin []GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VinInner, vout []GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2 instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2WithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2WithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2WithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetVin() []GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetVinOk() (*[]GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) SetVin(v []GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetVout() []GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2) SetVout(v []GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


