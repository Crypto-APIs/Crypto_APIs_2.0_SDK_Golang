# DecodeRawTransactionHexRISE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproximateFee** | Pointer to **string** | Defines the approximate fee value. When isConfirmed is True - Defines the amount of the transaction fee When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value. | [optional] 
**ApproximateMinimumRequiredFee** | Pointer to **string** | Defines the approximate minimum fee that is required for the transaction. | [optional] 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPaidForData** | Pointer to **string** | Represents the amount of gas paid for the data in the transaction. | [optional] 
**GasPrice** | Pointer to **string** | Represents the price offered to the miner to purchase this amount of gas. | [optional] 
**InputData** | Pointer to **string** | Represents additional information that is required for the transaction. | [optional] 
**MaxFeePerGas** | Pointer to **string** | Defines the maximum amount that customer is willing to pay per unit of gas to get his transaction included in a block. | [optional] 
**MaxFeePriorityPerGas** | Pointer to **string** | Represents determined by the user value that is paid directly to miners. | [optional] 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**R** | Pointer to **string** | Represents output of an ECDSA signature. | [optional] 
**Recipient** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**S** | Pointer to **string** | Represents output of an ECDSA signature. | [optional] 
**Sender** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Type** | **int32** | Specifies the transaction type as one from three options: if response returns a &#x60;\&quot;0\&quot;&#x60; it means the raw transaction includes legacy transaction data, if it is &#x60;\&quot;1\&quot;&#x60; - includes access lists for EIP2930, and if it is &#x60;\&quot;2\&quot;&#x60; - EIP1559 data. | 
**V** | Pointer to **string** | Defines the the recovery id. | [optional] 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 

## Methods

### NewDecodeRawTransactionHexRISE

`func NewDecodeRawTransactionHexRISE(gasLimit string, nonce int32, recipient string, sender string, type_ int32, ) *DecodeRawTransactionHexRISE`

NewDecodeRawTransactionHexRISE instantiates a new DecodeRawTransactionHexRISE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISEWithDefaults

`func NewDecodeRawTransactionHexRISEWithDefaults() *DecodeRawTransactionHexRISE`

NewDecodeRawTransactionHexRISEWithDefaults instantiates a new DecodeRawTransactionHexRISE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproximateFee

`func (o *DecodeRawTransactionHexRISE) GetApproximateFee() string`

GetApproximateFee returns the ApproximateFee field if non-nil, zero value otherwise.

### GetApproximateFeeOk

`func (o *DecodeRawTransactionHexRISE) GetApproximateFeeOk() (*string, bool)`

GetApproximateFeeOk returns a tuple with the ApproximateFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateFee

`func (o *DecodeRawTransactionHexRISE) SetApproximateFee(v string)`

SetApproximateFee sets ApproximateFee field to given value.

### HasApproximateFee

`func (o *DecodeRawTransactionHexRISE) HasApproximateFee() bool`

HasApproximateFee returns a boolean if a field has been set.

### GetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISE) GetApproximateMinimumRequiredFee() string`

GetApproximateMinimumRequiredFee returns the ApproximateMinimumRequiredFee field if non-nil, zero value otherwise.

### GetApproximateMinimumRequiredFeeOk

`func (o *DecodeRawTransactionHexRISE) GetApproximateMinimumRequiredFeeOk() (*string, bool)`

GetApproximateMinimumRequiredFeeOk returns a tuple with the ApproximateMinimumRequiredFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISE) SetApproximateMinimumRequiredFee(v string)`

SetApproximateMinimumRequiredFee sets ApproximateMinimumRequiredFee field to given value.

### HasApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISE) HasApproximateMinimumRequiredFee() bool`

HasApproximateMinimumRequiredFee returns a boolean if a field has been set.

### GetGasLimit

`func (o *DecodeRawTransactionHexRISE) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *DecodeRawTransactionHexRISE) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *DecodeRawTransactionHexRISE) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPaidForData

`func (o *DecodeRawTransactionHexRISE) GetGasPaidForData() string`

GetGasPaidForData returns the GasPaidForData field if non-nil, zero value otherwise.

### GetGasPaidForDataOk

`func (o *DecodeRawTransactionHexRISE) GetGasPaidForDataOk() (*string, bool)`

GetGasPaidForDataOk returns a tuple with the GasPaidForData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPaidForData

`func (o *DecodeRawTransactionHexRISE) SetGasPaidForData(v string)`

SetGasPaidForData sets GasPaidForData field to given value.

### HasGasPaidForData

`func (o *DecodeRawTransactionHexRISE) HasGasPaidForData() bool`

HasGasPaidForData returns a boolean if a field has been set.

### GetGasPrice

`func (o *DecodeRawTransactionHexRISE) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *DecodeRawTransactionHexRISE) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *DecodeRawTransactionHexRISE) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *DecodeRawTransactionHexRISE) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetInputData

`func (o *DecodeRawTransactionHexRISE) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *DecodeRawTransactionHexRISE) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *DecodeRawTransactionHexRISE) SetInputData(v string)`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *DecodeRawTransactionHexRISE) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *DecodeRawTransactionHexRISE) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *DecodeRawTransactionHexRISE) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *DecodeRawTransactionHexRISE) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *DecodeRawTransactionHexRISE) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxFeePriorityPerGas

`func (o *DecodeRawTransactionHexRISE) GetMaxFeePriorityPerGas() string`

GetMaxFeePriorityPerGas returns the MaxFeePriorityPerGas field if non-nil, zero value otherwise.

### GetMaxFeePriorityPerGasOk

`func (o *DecodeRawTransactionHexRISE) GetMaxFeePriorityPerGasOk() (*string, bool)`

GetMaxFeePriorityPerGasOk returns a tuple with the MaxFeePriorityPerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePriorityPerGas

`func (o *DecodeRawTransactionHexRISE) SetMaxFeePriorityPerGas(v string)`

SetMaxFeePriorityPerGas sets MaxFeePriorityPerGas field to given value.

### HasMaxFeePriorityPerGas

`func (o *DecodeRawTransactionHexRISE) HasMaxFeePriorityPerGas() bool`

HasMaxFeePriorityPerGas returns a boolean if a field has been set.

### GetNonce

`func (o *DecodeRawTransactionHexRISE) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DecodeRawTransactionHexRISE) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DecodeRawTransactionHexRISE) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetR

`func (o *DecodeRawTransactionHexRISE) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *DecodeRawTransactionHexRISE) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *DecodeRawTransactionHexRISE) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *DecodeRawTransactionHexRISE) HasR() bool`

HasR returns a boolean if a field has been set.

### GetRecipient

`func (o *DecodeRawTransactionHexRISE) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *DecodeRawTransactionHexRISE) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *DecodeRawTransactionHexRISE) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetS

`func (o *DecodeRawTransactionHexRISE) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *DecodeRawTransactionHexRISE) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *DecodeRawTransactionHexRISE) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *DecodeRawTransactionHexRISE) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSender

`func (o *DecodeRawTransactionHexRISE) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DecodeRawTransactionHexRISE) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DecodeRawTransactionHexRISE) SetSender(v string)`

SetSender sets Sender field to given value.


### GetType

`func (o *DecodeRawTransactionHexRISE) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecodeRawTransactionHexRISE) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecodeRawTransactionHexRISE) SetType(v int32)`

SetType sets Type field to given value.


### GetV

`func (o *DecodeRawTransactionHexRISE) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *DecodeRawTransactionHexRISE) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *DecodeRawTransactionHexRISE) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *DecodeRawTransactionHexRISE) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *DecodeRawTransactionHexRISE) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DecodeRawTransactionHexRISE) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DecodeRawTransactionHexRISE) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DecodeRawTransactionHexRISE) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


