# ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin, vout []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetVin() []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetVinOk() (*[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) SetVin(v []ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoinVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetVout() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) GetVoutOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificDogecoin) SetVout(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


