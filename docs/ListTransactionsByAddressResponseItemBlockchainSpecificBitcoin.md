# ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Defines the transaction&#39;s virtual size. | 
**Version** | **int32** | Defines the version of the transaction. | 
**Vin** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin**](ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout**](ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoin

`func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoin(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin, vout []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout, ) *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin`

NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin`

NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVin() []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVinOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVin(v []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVout() []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVoutOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVout(v []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


