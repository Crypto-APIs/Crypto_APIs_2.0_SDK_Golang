# ListConfirmedTransactionsByAddressRIBSBC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListConfirmedTransactionsByAddressRIBSBCVinInner**](ListConfirmedTransactionsByAddressRIBSBCVinInner.md) | Represents the transaction inputs. | 
**Vout** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSBCVoutInner**](GetTransactionDetailsByTransactionIDRIBSBCVoutInner.md) | Represents the transaction outputs. | [optional] 

## Methods

### NewListConfirmedTransactionsByAddressRIBSBC

`func NewListConfirmedTransactionsByAddressRIBSBC(locktime int64, size int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSBCVinInner, ) *ListConfirmedTransactionsByAddressRIBSBC`

NewListConfirmedTransactionsByAddressRIBSBC instantiates a new ListConfirmedTransactionsByAddressRIBSBC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSBCWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSBCWithDefaults() *ListConfirmedTransactionsByAddressRIBSBC`

NewListConfirmedTransactionsByAddressRIBSBCWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSBC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListConfirmedTransactionsByAddressRIBSBC) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListConfirmedTransactionsByAddressRIBSBC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConfirmedTransactionsByAddressRIBSBC) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetVin() []ListConfirmedTransactionsByAddressRIBSBCVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetVinOk() (*[]ListConfirmedTransactionsByAddressRIBSBCVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListConfirmedTransactionsByAddressRIBSBC) SetVin(v []ListConfirmedTransactionsByAddressRIBSBCVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetVout() []GetTransactionDetailsByTransactionIDRIBSBCVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBSBC) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSBCVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBSBC) SetVout(v []GetTransactionDetailsByTransactionIDRIBSBCVoutInner)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListConfirmedTransactionsByAddressRIBSBC) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


