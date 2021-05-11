# ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | String representation of the transaction Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice**](ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum(contract string, gasLimit string, gasPrice ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetGasPrice() ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetGasPriceOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetGasPrice(v ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereumGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificEthereum) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


