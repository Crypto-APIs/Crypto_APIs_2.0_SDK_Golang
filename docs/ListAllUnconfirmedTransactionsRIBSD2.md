# ListAllUnconfirmedTransactionsRIBSD2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListAllUnconfirmedTransactionsRIBSD2Vin**](ListAllUnconfirmedTransactionsRIBSD2Vin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListConfirmedTransactionsByAddressRIBSD2Vout**](ListConfirmedTransactionsByAddressRIBSD2Vout.md) | Represents the transaction outputs. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSD2

`func NewListAllUnconfirmedTransactionsRIBSD2(locktime int64, size int32, version int32, vin []ListAllUnconfirmedTransactionsRIBSD2Vin, vout []ListConfirmedTransactionsByAddressRIBSD2Vout, ) *ListAllUnconfirmedTransactionsRIBSD2`

NewListAllUnconfirmedTransactionsRIBSD2 instantiates a new ListAllUnconfirmedTransactionsRIBSD2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSD2WithDefaults

`func NewListAllUnconfirmedTransactionsRIBSD2WithDefaults() *ListAllUnconfirmedTransactionsRIBSD2`

NewListAllUnconfirmedTransactionsRIBSD2WithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSD2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBSD2) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListAllUnconfirmedTransactionsRIBSD2) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListAllUnconfirmedTransactionsRIBSD2) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetVin() []ListAllUnconfirmedTransactionsRIBSD2Vin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetVinOk() (*[]ListAllUnconfirmedTransactionsRIBSD2Vin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListAllUnconfirmedTransactionsRIBSD2) SetVin(v []ListAllUnconfirmedTransactionsRIBSD2Vin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetVout() []ListConfirmedTransactionsByAddressRIBSD2Vout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSD2) GetVoutOk() (*[]ListConfirmedTransactionsByAddressRIBSD2Vout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSD2) SetVout(v []ListConfirmedTransactionsByAddressRIBSD2Vout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


