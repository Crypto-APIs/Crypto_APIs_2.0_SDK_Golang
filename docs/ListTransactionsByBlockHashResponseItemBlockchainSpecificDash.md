# ListTransactionsByBlockHashResponseItemBlockchainSpecificDash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificDash

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDash(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin, vout []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificDash instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificDashWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetVin() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetVinOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) SetVin(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetVout() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) GetVoutOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDash) SetVout(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


