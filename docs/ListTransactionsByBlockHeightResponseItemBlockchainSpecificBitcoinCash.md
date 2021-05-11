# ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the total size of this transaction. | 
**Vin** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin**](ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout**](ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin, vout []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCashWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCashWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCashWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetVin() []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetVinOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) SetVin(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetVout() []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) GetVoutOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificBitcoinCash) SetVout(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


