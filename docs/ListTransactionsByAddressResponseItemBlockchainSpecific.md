# ListTransactionsByAddressResponseItemBlockchainSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction&#39;s version number. | 
**Vin** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVin**](ListTransactionsByAddressResponseItemBlockchainSpecificDashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVout**](ListTransactionsByAddressResponseItemBlockchainSpecificDashVout.md) | Represents the transaction outputs. | 
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListTransactionsByAddressResponseItemBlockchainSpecificEthereumGasPrice**](ListTransactionsByAddressResponseItemBlockchainSpecificEthereumGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | String representation of the transaction status | 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecific

`func NewListTransactionsByAddressResponseItemBlockchainSpecific(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin, vout []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout, contract string, gasLimit string, gasPrice ListTransactionsByAddressResponseItemBlockchainSpecificEthereumGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string, ) *ListTransactionsByAddressResponseItemBlockchainSpecific`

NewListTransactionsByAddressResponseItemBlockchainSpecific instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecific`

NewListTransactionsByAddressResponseItemBlockchainSpecificWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVin() []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVinOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetVin(v []ListTransactionsByAddressResponseItemBlockchainSpecificDashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVout() []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetVoutOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificDashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetVout(v []ListTransactionsByAddressResponseItemBlockchainSpecificDashVout)`

SetVout sets Vout field to given value.


### GetContract

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetGasPrice() ListTransactionsByAddressResponseItemBlockchainSpecificEthereumGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetGasPriceOk() (*ListTransactionsByAddressResponseItemBlockchainSpecificEthereumGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetGasPrice(v ListTransactionsByAddressResponseItemBlockchainSpecificEthereumGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecific) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


