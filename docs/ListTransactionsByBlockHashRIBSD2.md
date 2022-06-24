# ListTransactionsByBlockHashRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHashRIBSD2VinInner**](ListTransactionsByBlockHashRIBSD2VinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashRIBSD2VoutInner**](ListTransactionsByBlockHashRIBSD2VoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHashRIBSD2

`func NewListTransactionsByBlockHashRIBSD2(locktime int64, size int32, version int32, vin []ListTransactionsByBlockHashRIBSD2VinInner, vout []ListTransactionsByBlockHashRIBSD2VoutInner, ) *ListTransactionsByBlockHashRIBSD2`

NewListTransactionsByBlockHashRIBSD2 instantiates a new ListTransactionsByBlockHashRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSD2WithDefaults

`func NewListTransactionsByBlockHashRIBSD2WithDefaults() *ListTransactionsByBlockHashRIBSD2`

NewListTransactionsByBlockHashRIBSD2WithDefaults instantiates a new ListTransactionsByBlockHashRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHashRIBSD2) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHashRIBSD2) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHashRIBSD2) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHashRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHashRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHashRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHashRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHashRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHashRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHashRIBSD2) GetVin() []ListTransactionsByBlockHashRIBSD2VinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHashRIBSD2) GetVinOk() (*[]ListTransactionsByBlockHashRIBSD2VinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHashRIBSD2) SetVin(v []ListTransactionsByBlockHashRIBSD2VinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashRIBSD2) GetVout() []ListTransactionsByBlockHashRIBSD2VoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashRIBSD2) GetVoutOk() (*[]ListTransactionsByBlockHashRIBSD2VoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashRIBSD2) SetVout(v []ListTransactionsByBlockHashRIBSD2VoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


