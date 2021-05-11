# ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin, vout []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetVin() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetVinOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) SetVin(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetVout() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) GetVoutOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoin) SetVout(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


