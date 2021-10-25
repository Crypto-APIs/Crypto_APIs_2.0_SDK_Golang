# GetTransactionDetailsByTransactionIDRIBSBSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | Pointer to **string** | Represents the specific transaction contract | [optional] 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**GetTransactionDetailsByTransactionIDRIBSBSCGasPrice**](GetTransactionDetailsByTransactionIDRIBSBSCGasPrice.md) |  | 
**GasUsed** | **string** | Defines the unit of the gas price amount, e.g. BTC, ETH, XRP. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSBSC

`func NewGetTransactionDetailsByTransactionIDRIBSBSC(gasLimit string, gasPrice GetTransactionDetailsByTransactionIDRIBSBSCGasPrice, gasUsed string, inputData string, nonce int32, ) *GetTransactionDetailsByTransactionIDRIBSBSC`

NewGetTransactionDetailsByTransactionIDRIBSBSC instantiates a new GetTransactionDetailsByTransactionIDRIBSBSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSBSCWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSBSCWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBSC`

NewGetTransactionDetailsByTransactionIDRIBSBSCWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetGasLimit

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasPrice() GetTransactionDetailsByTransactionIDRIBSBSCGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDRIBSBSCGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetGasPrice(v GetTransactionDetailsByTransactionIDRIBSBSCGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetTransactionDetailsByTransactionIDRIBSBSC) SetNonce(v int32)`

SetNonce sets Nonce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


