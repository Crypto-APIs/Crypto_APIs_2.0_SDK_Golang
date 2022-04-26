# DecodeRawTransactionHexRISB22

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproximateFee** | Pointer to **string** | Defines the approximate fee value. When isConfirmed is True - Defines the amount of the transaction fee When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value. | [optional] 
**ApproximateMinimumRequiredFee** | Pointer to **string** | Defines the approximate minimum fee that is required for the transaction. | [optional] 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPaidForData** | Pointer to **string** | Represents the amount of gas paid for the data in the transaction. | [optional] 
**GasPrice** | Pointer to **string** | Represents the price offered to the miner to purchase this amount of gas. | [optional] 
**InputData** | Pointer to **string** | Represents additional information that is required for the transaction. | [optional] 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**R** | Pointer to **string** | Represents output of an ECDSA signature. | [optional] 
**Recipient** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**S** | Pointer to **string** | Represents output of an ECDSA signature. | [optional] 
**Sender** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Type** | **int32** | Specifies the transaction type as one from three options: if response returns a &#x60;\&quot;0\&quot;&#x60; it means the raw transaction includes legacy transaction data, if it is &#x60;\&quot;1\&quot;&#x60; - includes access lists for EIP2930, and if it is &#x60;\&quot;2\&quot;&#x60; - EIP1559 data. | 
**V** | Pointer to **string** | Defines the the recovery id. | [optional] 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 

## Methods

### NewDecodeRawTransactionHexRISB22

`func NewDecodeRawTransactionHexRISB22(gasLimit string, nonce int32, recipient string, sender string, type_ int32, ) *DecodeRawTransactionHexRISB22`

NewDecodeRawTransactionHexRISB22 instantiates a new DecodeRawTransactionHexRISB22 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISB22WithDefaults

`func NewDecodeRawTransactionHexRISB22WithDefaults() *DecodeRawTransactionHexRISB22`

NewDecodeRawTransactionHexRISB22WithDefaults instantiates a new DecodeRawTransactionHexRISB22 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproximateFee

`func (o *DecodeRawTransactionHexRISB22) GetApproximateFee() string`

GetApproximateFee returns the ApproximateFee field if non-nil, zero value otherwise.

### GetApproximateFeeOk

`func (o *DecodeRawTransactionHexRISB22) GetApproximateFeeOk() (*string, bool)`

GetApproximateFeeOk returns a tuple with the ApproximateFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateFee

`func (o *DecodeRawTransactionHexRISB22) SetApproximateFee(v string)`

SetApproximateFee sets ApproximateFee field to given value.

### HasApproximateFee

`func (o *DecodeRawTransactionHexRISB22) HasApproximateFee() bool`

HasApproximateFee returns a boolean if a field has been set.

### GetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISB22) GetApproximateMinimumRequiredFee() string`

GetApproximateMinimumRequiredFee returns the ApproximateMinimumRequiredFee field if non-nil, zero value otherwise.

### GetApproximateMinimumRequiredFeeOk

`func (o *DecodeRawTransactionHexRISB22) GetApproximateMinimumRequiredFeeOk() (*string, bool)`

GetApproximateMinimumRequiredFeeOk returns a tuple with the ApproximateMinimumRequiredFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISB22) SetApproximateMinimumRequiredFee(v string)`

SetApproximateMinimumRequiredFee sets ApproximateMinimumRequiredFee field to given value.

### HasApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISB22) HasApproximateMinimumRequiredFee() bool`

HasApproximateMinimumRequiredFee returns a boolean if a field has been set.

### GetGasLimit

`func (o *DecodeRawTransactionHexRISB22) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *DecodeRawTransactionHexRISB22) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *DecodeRawTransactionHexRISB22) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPaidForData

`func (o *DecodeRawTransactionHexRISB22) GetGasPaidForData() string`

GetGasPaidForData returns the GasPaidForData field if non-nil, zero value otherwise.

### GetGasPaidForDataOk

`func (o *DecodeRawTransactionHexRISB22) GetGasPaidForDataOk() (*string, bool)`

GetGasPaidForDataOk returns a tuple with the GasPaidForData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPaidForData

`func (o *DecodeRawTransactionHexRISB22) SetGasPaidForData(v string)`

SetGasPaidForData sets GasPaidForData field to given value.

### HasGasPaidForData

`func (o *DecodeRawTransactionHexRISB22) HasGasPaidForData() bool`

HasGasPaidForData returns a boolean if a field has been set.

### GetGasPrice

`func (o *DecodeRawTransactionHexRISB22) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *DecodeRawTransactionHexRISB22) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *DecodeRawTransactionHexRISB22) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *DecodeRawTransactionHexRISB22) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetInputData

`func (o *DecodeRawTransactionHexRISB22) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *DecodeRawTransactionHexRISB22) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *DecodeRawTransactionHexRISB22) SetInputData(v string)`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *DecodeRawTransactionHexRISB22) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetNonce

`func (o *DecodeRawTransactionHexRISB22) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DecodeRawTransactionHexRISB22) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DecodeRawTransactionHexRISB22) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetR

`func (o *DecodeRawTransactionHexRISB22) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *DecodeRawTransactionHexRISB22) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *DecodeRawTransactionHexRISB22) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *DecodeRawTransactionHexRISB22) HasR() bool`

HasR returns a boolean if a field has been set.

### GetRecipient

`func (o *DecodeRawTransactionHexRISB22) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *DecodeRawTransactionHexRISB22) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *DecodeRawTransactionHexRISB22) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetS

`func (o *DecodeRawTransactionHexRISB22) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *DecodeRawTransactionHexRISB22) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *DecodeRawTransactionHexRISB22) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *DecodeRawTransactionHexRISB22) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSender

`func (o *DecodeRawTransactionHexRISB22) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DecodeRawTransactionHexRISB22) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DecodeRawTransactionHexRISB22) SetSender(v string)`

SetSender sets Sender field to given value.


### GetType

`func (o *DecodeRawTransactionHexRISB22) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecodeRawTransactionHexRISB22) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecodeRawTransactionHexRISB22) SetType(v int32)`

SetType sets Type field to given value.


### GetV

`func (o *DecodeRawTransactionHexRISB22) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *DecodeRawTransactionHexRISB22) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *DecodeRawTransactionHexRISB22) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *DecodeRawTransactionHexRISB22) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *DecodeRawTransactionHexRISB22) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DecodeRawTransactionHexRISB22) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DecodeRawTransactionHexRISB22) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DecodeRawTransactionHexRISB22) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


