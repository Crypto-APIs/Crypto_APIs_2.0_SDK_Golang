# ListTransactionsByBlockHashResponseItemBlockchainSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout**](ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout.md) | Represents the transaction outputs. | 
**Vsize** | **int32** | Represents the virtual size of this transaction. | 
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice**](ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **string** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | String representation of the transaction status | 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecific

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecific(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin, vout []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout, vsize int32, contract string, gasLimit string, gasPrice ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice, gasUsed string, inputData string, nonce string, transactionStatus string, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecific`

NewListTransactionsByBlockHashResponseItemBlockchainSpecific instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecific`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVin() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVinOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetVin(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVout() []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVoutOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetVout(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificDashVout)`

SetVout sets Vout field to given value.


### GetVsize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVsize() int32`

GetVsize returns the Vsize field if non-nil, zero value otherwise.

### GetVsizeOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetVsizeOk() (*int32, bool)`

GetVsizeOk returns a tuple with the Vsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsize

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetVsize(v int32)`

SetVsize sets Vsize field to given value.


### GetContract

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetGasPrice() ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetGasPriceOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetGasPrice(v ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecific) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


