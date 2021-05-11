# ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **string** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic(contract string, gasLimit string, gasPrice ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice, gasUsed string, inputData string, nonce string, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetGasPrice() ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetGasPriceOk() (*ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) SetGasPrice(v ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassicGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificEthereumClassic) SetNonce(v string)`

SetNonce sets Nonce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


