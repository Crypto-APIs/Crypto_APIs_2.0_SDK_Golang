# ListAllUnconfirmedTransactionsRIBSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Defines the version of the transaction. | 
**Vin** | [**[]ListAllUnconfirmedTransactionsRIBSDVin**](ListAllUnconfirmedTransactionsRIBSDVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSDVout**](GetTransactionDetailsByTransactionIDRIBSDVout.md) | Represents the transaction outputs. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSD

`func NewListAllUnconfirmedTransactionsRIBSD(locktime int64, size int32, version int32, vin []ListAllUnconfirmedTransactionsRIBSDVin, vout []GetTransactionDetailsByTransactionIDRIBSDVout, ) *ListAllUnconfirmedTransactionsRIBSD`

NewListAllUnconfirmedTransactionsRIBSD instantiates a new ListAllUnconfirmedTransactionsRIBSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSDWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSDWithDefaults() *ListAllUnconfirmedTransactionsRIBSD`

NewListAllUnconfirmedTransactionsRIBSDWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBSD) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListAllUnconfirmedTransactionsRIBSD) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListAllUnconfirmedTransactionsRIBSD) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetVin() []ListAllUnconfirmedTransactionsRIBSDVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetVinOk() (*[]ListAllUnconfirmedTransactionsRIBSDVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListAllUnconfirmedTransactionsRIBSD) SetVin(v []ListAllUnconfirmedTransactionsRIBSDVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetVout() []GetTransactionDetailsByTransactionIDRIBSDVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSD) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSD) SetVout(v []GetTransactionDetailsByTransactionIDRIBSDVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


