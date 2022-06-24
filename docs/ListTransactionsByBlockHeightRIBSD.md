# ListTransactionsByBlockHeightRIBSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHeightRIBSDVinInner**](ListTransactionsByBlockHeightRIBSDVinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHeightRIBSDVoutInner**](ListTransactionsByBlockHeightRIBSDVoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHeightRIBSD

`func NewListTransactionsByBlockHeightRIBSD(locktime int64, size int32, version int32, vin []ListTransactionsByBlockHeightRIBSDVinInner, vout []ListTransactionsByBlockHeightRIBSDVoutInner, ) *ListTransactionsByBlockHeightRIBSD`

NewListTransactionsByBlockHeightRIBSD instantiates a new ListTransactionsByBlockHeightRIBSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSDWithDefaults

`func NewListTransactionsByBlockHeightRIBSDWithDefaults() *ListTransactionsByBlockHeightRIBSD`

NewListTransactionsByBlockHeightRIBSDWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightRIBSD) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightRIBSD) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightRIBSD) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightRIBSD) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightRIBSD) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightRIBSD) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightRIBSD) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightRIBSD) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightRIBSD) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightRIBSD) GetVin() []ListTransactionsByBlockHeightRIBSDVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightRIBSD) GetVinOk() (*[]ListTransactionsByBlockHeightRIBSDVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightRIBSD) SetVin(v []ListTransactionsByBlockHeightRIBSDVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightRIBSD) GetVout() []ListTransactionsByBlockHeightRIBSDVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightRIBSD) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSDVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightRIBSD) SetVout(v []ListTransactionsByBlockHeightRIBSDVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


