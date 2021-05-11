# GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumGasPrice**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 

## Methods

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum(contract string, gasLimit string, gasPrice GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string, ) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumWithDefaults

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetGasPrice() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetGasPrice(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereumGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificEthereum) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


