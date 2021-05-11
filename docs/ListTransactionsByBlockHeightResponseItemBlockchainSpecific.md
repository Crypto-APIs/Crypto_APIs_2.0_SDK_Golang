# ListTransactionsByBlockHeightResponseItemBlockchainSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the total size of this transaction. | 
**Vin** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin**](ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout**](ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout.md) | Represents the transaction outputs. | 
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **string** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 
**Vsize** | **int32** | Represents the virtual size of this transaction. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecific

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecific(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin, vout []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout, contract string, gasLimit string, gasPrice ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice, gasUsed string, inputData string, nonce string, transactionStatus string, vsize int32, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecific`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecific instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecific`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVin() []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVinOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetVin(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVout() []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVoutOk() (*[]ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetVout(v []ListTransactionsByBlockHashResponseItemBlockchainSpecificBitcoinCashVout)`

SetVout sets Vout field to given value.


### GetContract

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetGasPrice() ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetGasPriceOk() (*ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetGasPrice(v ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetVsize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVsize() int32`

GetVsize returns the Vsize field if non-nil, zero value otherwise.

### GetVsizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) GetVsizeOk() (*int32, bool)`

GetVsizeOk returns a tuple with the Vsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecific) SetVsize(v int32)`

SetVsize sets Vsize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


