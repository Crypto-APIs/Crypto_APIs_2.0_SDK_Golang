# GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasPrice** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice**](GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice.md) |  | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2(gasLimit int32, gasPrice GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice, gasUsed int32, nonce int32, transactionStatus string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2WithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2WithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ2WithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasLimit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasPrice() GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetGasPrice(v GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetNonce

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


