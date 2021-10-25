# ListUnconfirmedTransactionsByAddressRIBSEC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | [**ListUnconfirmedTransactionsByAddressRIBSECFee**](ListUnconfirmedTransactionsByAddressRIBSECFee.md) |  | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListUnconfirmedTransactionsByAddressRIBSECGasPrice**](ListUnconfirmedTransactionsByAddressRIBSECGasPrice.md) |  | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | String representation of the transaction status | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSEC

`func NewListUnconfirmedTransactionsByAddressRIBSEC(fee ListUnconfirmedTransactionsByAddressRIBSECFee, gasLimit string, gasPrice ListUnconfirmedTransactionsByAddressRIBSECGasPrice, nonce int32, transactionStatus string, ) *ListUnconfirmedTransactionsByAddressRIBSEC`

NewListUnconfirmedTransactionsByAddressRIBSEC instantiates a new ListUnconfirmedTransactionsByAddressRIBSEC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSECWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSECWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSEC`

NewListUnconfirmedTransactionsByAddressRIBSECWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSEC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetFee() ListUnconfirmedTransactionsByAddressRIBSECFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetFeeOk() (*ListUnconfirmedTransactionsByAddressRIBSECFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) SetFee(v ListUnconfirmedTransactionsByAddressRIBSECFee)`

SetFee sets Fee field to given value.


### GetGasLimit

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetGasPrice() ListUnconfirmedTransactionsByAddressRIBSECGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetGasPriceOk() (*ListUnconfirmedTransactionsByAddressRIBSECGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) SetGasPrice(v ListUnconfirmedTransactionsByAddressRIBSECGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetNonce

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListUnconfirmedTransactionsByAddressRIBSEC) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


