# ListConfirmedTransactionsByAddressRIBSEC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListConfirmedTransactionsByAddressRIBSECGasPrice**](ListConfirmedTransactionsByAddressRIBSECGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**InternalTransactionsCount** | **int32** | Represents the total internal transactions count. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TokenTransfersCount** | **int32** | Represents the total token transfers count. | 
**TransactionStatus** | **string** | String representation of the transaction status | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSEC

`func NewListConfirmedTransactionsByAddressRIBSEC(contract string, gasLimit string, gasPrice ListConfirmedTransactionsByAddressRIBSECGasPrice, gasUsed string, inputData string, internalTransactionsCount int32, nonce int32, tokenTransfersCount int32, transactionStatus string, ) *ListConfirmedTransactionsByAddressRIBSEC`

NewListConfirmedTransactionsByAddressRIBSEC instantiates a new ListConfirmedTransactionsByAddressRIBSEC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSECWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSECWithDefaults() *ListConfirmedTransactionsByAddressRIBSEC`

NewListConfirmedTransactionsByAddressRIBSECWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSEC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasPrice() ListConfirmedTransactionsByAddressRIBSECGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasPriceOk() (*ListConfirmedTransactionsByAddressRIBSECGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetGasPrice(v ListConfirmedTransactionsByAddressRIBSECGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetInternalTransactionsCount

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInternalTransactionsCount() int32`

GetInternalTransactionsCount returns the InternalTransactionsCount field if non-nil, zero value otherwise.

### GetInternalTransactionsCountOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetInternalTransactionsCountOk() (*int32, bool)`

GetInternalTransactionsCountOk returns a tuple with the InternalTransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransactionsCount

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetInternalTransactionsCount(v int32)`

SetInternalTransactionsCount sets InternalTransactionsCount field to given value.


### GetNonce

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTokenTransfersCount

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTokenTransfersCount() int32`

GetTokenTransfersCount returns the TokenTransfersCount field if non-nil, zero value otherwise.

### GetTokenTransfersCountOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTokenTransfersCountOk() (*int32, bool)`

GetTokenTransfersCountOk returns a tuple with the TokenTransfersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfersCount

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetTokenTransfersCount(v int32)`

SetTokenTransfersCount sets TokenTransfersCount field to given value.


### GetTransactionStatus

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListConfirmedTransactionsByAddressRIBSEC) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListConfirmedTransactionsByAddressRIBSEC) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


