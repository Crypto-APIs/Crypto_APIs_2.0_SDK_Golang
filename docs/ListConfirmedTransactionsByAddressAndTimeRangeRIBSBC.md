# ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListConfirmedTransactionsByAddressRIBSBCVin**](ListConfirmedTransactionsByAddressRIBSBCVin.md) | Represents the transaction inputs. | 
**Vout** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSBCVout**](GetTransactionDetailsByTransactionIDRIBSBCVout.md) | Represents the transaction outputs. | [optional] 

## Methods

### NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBC

`func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBC(locktime int64, size int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSBCVin, ) *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC`

NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBC instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBCWithDefaults

`func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBCWithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC`

NewListConfirmedTransactionsByAddressAndTimeRangeRIBSBCWithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetVin() []ListConfirmedTransactionsByAddressRIBSBCVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetVinOk() (*[]ListConfirmedTransactionsByAddressRIBSBCVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) SetVin(v []ListConfirmedTransactionsByAddressRIBSBCVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetVout() []GetTransactionDetailsByTransactionIDRIBSBCVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSBCVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) SetVout(v []GetTransactionDetailsByTransactionIDRIBSBCVout)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSBC) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


