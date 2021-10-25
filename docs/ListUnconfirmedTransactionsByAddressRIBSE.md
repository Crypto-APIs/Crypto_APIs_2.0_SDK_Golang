# ListUnconfirmedTransactionsByAddressRIBSE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | [**ListUnconfirmedTransactionsByAddressRIBSEFee**](ListUnconfirmedTransactionsByAddressRIBSEFee.md) |  | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListConfirmedTransactionsByAddressRIBSEGasPrice**](ListConfirmedTransactionsByAddressRIBSEGasPrice.md) |  | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | String representation of the transaction status | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSE

`func NewListUnconfirmedTransactionsByAddressRIBSE(fee ListUnconfirmedTransactionsByAddressRIBSEFee, gasLimit string, gasPrice ListConfirmedTransactionsByAddressRIBSEGasPrice, inputData string, nonce int32, transactionStatus string, ) *ListUnconfirmedTransactionsByAddressRIBSE`

NewListUnconfirmedTransactionsByAddressRIBSE instantiates a new ListUnconfirmedTransactionsByAddressRIBSE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSEWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSEWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSE`

NewListUnconfirmedTransactionsByAddressRIBSEWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetFee() ListUnconfirmedTransactionsByAddressRIBSEFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetFeeOk() (*ListUnconfirmedTransactionsByAddressRIBSEFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) SetFee(v ListUnconfirmedTransactionsByAddressRIBSEFee)`

SetFee sets Fee field to given value.


### GetGasLimit

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetGasPrice() ListConfirmedTransactionsByAddressRIBSEGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetGasPriceOk() (*ListConfirmedTransactionsByAddressRIBSEGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) SetGasPrice(v ListConfirmedTransactionsByAddressRIBSEGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetInputData

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListUnconfirmedTransactionsByAddressRIBSE) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


