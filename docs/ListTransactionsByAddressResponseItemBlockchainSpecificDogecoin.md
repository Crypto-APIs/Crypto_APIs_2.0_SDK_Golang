# ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin**](ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinVout**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinVout.md) | Represents the transaction outputs. | 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoin

`func NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoin(locktime int32, size int32, version int32, vin []ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin, vout []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinVout, ) *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin`

NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin`

NewListTransactionsByAddressResponseItemBlockchainSpecificDogecoinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetVin() []ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetVinOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) SetVin(v []ListTransactionsByAddressResponseItemBlockchainSpecificDogecoinVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetVout() []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDogecoin) SetVout(v []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


