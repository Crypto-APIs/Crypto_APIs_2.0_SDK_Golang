# ListUnconfirmedTransactionsByAddressRIBSBC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListUnconfirmedTransactionsByAddressRIBSBCVinInner**](ListUnconfirmedTransactionsByAddressRIBSBCVinInner.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListUnconfirmedTransactionsByAddressRIBSBCVoutInner**](ListUnconfirmedTransactionsByAddressRIBSBCVoutInner.md) | Represents the transaction outputs. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSBC

`func NewListUnconfirmedTransactionsByAddressRIBSBC(locktime int64, size int32, version int32, vin []ListUnconfirmedTransactionsByAddressRIBSBCVinInner, vout []ListUnconfirmedTransactionsByAddressRIBSBCVoutInner, ) *ListUnconfirmedTransactionsByAddressRIBSBC`

NewListUnconfirmedTransactionsByAddressRIBSBC instantiates a new ListUnconfirmedTransactionsByAddressRIBSBC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSBCWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSBCWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSBC`

NewListUnconfirmedTransactionsByAddressRIBSBCWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSBC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetVin() []ListUnconfirmedTransactionsByAddressRIBSBCVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetVinOk() (*[]ListUnconfirmedTransactionsByAddressRIBSBCVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) SetVin(v []ListUnconfirmedTransactionsByAddressRIBSBCVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetVout() []ListUnconfirmedTransactionsByAddressRIBSBCVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) GetVoutOk() (*[]ListUnconfirmedTransactionsByAddressRIBSBCVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSBC) SetVout(v []ListUnconfirmedTransactionsByAddressRIBSBCVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


