# ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVin**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVout**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDash

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDash(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVin, vout []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVout, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDash instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDashWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDashWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDashWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetVin() []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetVinOk() (*[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) SetVin(v []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetVout() []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) GetVoutOk() (*[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDash) SetVout(v []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDashVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


