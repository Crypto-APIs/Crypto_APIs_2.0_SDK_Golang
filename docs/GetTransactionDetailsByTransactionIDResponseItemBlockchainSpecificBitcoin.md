# GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVin**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout.md) | Represents the transaction outputs. | 

## Methods

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin(locktime int32, size int32, vSize int32, version int32, vin []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVin, vout []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout, ) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinWithDefaults

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVin() []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVinOk() (*[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) SetVin(v []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVout() []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoin) SetVout(v []GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


