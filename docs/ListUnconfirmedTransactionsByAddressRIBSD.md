# ListUnconfirmedTransactionsByAddressRIBSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Numeric representation of the transaction version | 
**Vin** | [**[]ListUnconfirmedTransactionsByAddressRIBSDVin**](ListUnconfirmedTransactionsByAddressRIBSDVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSDVout**](GetTransactionDetailsByTransactionIDRIBSDVout.md) | Represents the transaction outputs. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSD

`func NewListUnconfirmedTransactionsByAddressRIBSD(locktime int64, size int32, version int32, vin []ListUnconfirmedTransactionsByAddressRIBSDVin, vout []GetTransactionDetailsByTransactionIDRIBSDVout, ) *ListUnconfirmedTransactionsByAddressRIBSD`

NewListUnconfirmedTransactionsByAddressRIBSD instantiates a new ListUnconfirmedTransactionsByAddressRIBSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSDWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSDWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSD`

NewListUnconfirmedTransactionsByAddressRIBSDWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVin() []ListUnconfirmedTransactionsByAddressRIBSDVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVinOk() (*[]ListUnconfirmedTransactionsByAddressRIBSDVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetVin(v []ListUnconfirmedTransactionsByAddressRIBSDVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVout() []GetTransactionDetailsByTransactionIDRIBSDVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetVout(v []GetTransactionDetailsByTransactionIDRIBSDVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


