# ListAllUnconfirmedTransactionsRIBSBSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListConfirmedTransactionsByAddressRIBSBSCGasPrice**](ListConfirmedTransactionsByAddressRIBSBSCGasPrice.md) |  | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Defines the transaction status. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSBSC

`func NewListAllUnconfirmedTransactionsRIBSBSC(gasLimit string, gasPrice ListConfirmedTransactionsByAddressRIBSBSCGasPrice, inputData string, nonce int32, transactionStatus string, ) *ListAllUnconfirmedTransactionsRIBSBSC`

NewListAllUnconfirmedTransactionsRIBSBSC instantiates a new ListAllUnconfirmedTransactionsRIBSBSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSBSCWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSBSCWithDefaults() *ListAllUnconfirmedTransactionsRIBSBSC`

NewListAllUnconfirmedTransactionsRIBSBSCWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSBSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasLimit

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetGasPrice() ListConfirmedTransactionsByAddressRIBSBSCGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetGasPriceOk() (*ListConfirmedTransactionsByAddressRIBSBSCGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) SetGasPrice(v ListConfirmedTransactionsByAddressRIBSBSCGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetInputData

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListAllUnconfirmedTransactionsRIBSBSC) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


