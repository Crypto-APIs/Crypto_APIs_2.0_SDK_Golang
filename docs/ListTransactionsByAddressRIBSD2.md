# ListTransactionsByAddressRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListTransactionsByAddressRIBSD2Vin**](ListTransactionsByAddressRIBSD2Vin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByAddressRIBSD2Vout**](ListTransactionsByAddressRIBSD2Vout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByAddressRIBSD2

`func NewListTransactionsByAddressRIBSD2(locktime int32, size int32, version int32, vin []ListTransactionsByAddressRIBSD2Vin, vout []ListTransactionsByAddressRIBSD2Vout, ) *ListTransactionsByAddressRIBSD2`

NewListTransactionsByAddressRIBSD2 instantiates a new ListTransactionsByAddressRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressRIBSD2WithDefaults

`func NewListTransactionsByAddressRIBSD2WithDefaults() *ListTransactionsByAddressRIBSD2`

NewListTransactionsByAddressRIBSD2WithDefaults instantiates a new ListTransactionsByAddressRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByAddressRIBSD2) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByAddressRIBSD2) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByAddressRIBSD2) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByAddressRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByAddressRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByAddressRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByAddressRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByAddressRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByAddressRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByAddressRIBSD2) GetVin() []ListTransactionsByAddressRIBSD2Vin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByAddressRIBSD2) GetVinOk() (*[]ListTransactionsByAddressRIBSD2Vin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByAddressRIBSD2) SetVin(v []ListTransactionsByAddressRIBSD2Vin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByAddressRIBSD2) GetVout() []ListTransactionsByAddressRIBSD2Vout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressRIBSD2) GetVoutOk() (*[]ListTransactionsByAddressRIBSD2Vout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressRIBSD2) SetVout(v []ListTransactionsByAddressRIBSD2Vout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


