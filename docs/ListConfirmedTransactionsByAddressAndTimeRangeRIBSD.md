# ListConfirmedTransactionsByAddressAndTimeRangeRIBSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListConfirmedTransactionsByAddressRIBSDVin**](ListConfirmedTransactionsByAddressRIBSDVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSDVout**](GetTransactionDetailsByTransactionIDRIBSDVout.md) | Represents the transaction outputs. | 

## Methods

### NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD

`func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD(locktime int64, size int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSDVin, vout []GetTransactionDetailsByTransactionIDRIBSDVout, ) *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD`

NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressAndTimeRangeRIBSDWithDefaults

`func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSDWithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD`

NewListConfirmedTransactionsByAddressAndTimeRangeRIBSDWithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetVin() []ListConfirmedTransactionsByAddressRIBSDVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetVinOk() (*[]ListConfirmedTransactionsByAddressRIBSDVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) SetVin(v []ListConfirmedTransactionsByAddressRIBSDVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetVout() []GetTransactionDetailsByTransactionIDRIBSDVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD) SetVout(v []GetTransactionDetailsByTransactionIDRIBSDVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


