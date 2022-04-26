# DecodeRawTransactionHexRISE2

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

### NewDecodeRawTransactionHexRISE2

`func NewDecodeRawTransactionHexRISE2(gasLimit string, nonce int32, recipient string, sender string, type_ int32, ) *DecodeRawTransactionHexRISE2`

NewDecodeRawTransactionHexRISE2 instantiates a new DecodeRawTransactionHexRISE2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISE2WithDefaults

`func NewDecodeRawTransactionHexRISE2WithDefaults() *DecodeRawTransactionHexRISE2`

NewDecodeRawTransactionHexRISE2WithDefaults instantiates a new DecodeRawTransactionHexRISE2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproximateFee

`func (o *DecodeRawTransactionHexRISE2) GetApproximateFee() string`

GetApproximateFee returns the ApproximateFee field if non-nil, zero value otherwise.

### GetApproximateFeeOk

`func (o *DecodeRawTransactionHexRISE2) GetApproximateFeeOk() (*string, bool)`

GetApproximateFeeOk returns a tuple with the ApproximateFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateFee

`func (o *DecodeRawTransactionHexRISE2) SetApproximateFee(v string)`

SetApproximateFee sets ApproximateFee field to given value.

### HasApproximateFee

`func (o *DecodeRawTransactionHexRISE2) HasApproximateFee() bool`

HasApproximateFee returns a boolean if a field has been set.

### GetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISE2) GetApproximateMinimumRequiredFee() string`

GetApproximateMinimumRequiredFee returns the ApproximateMinimumRequiredFee field if non-nil, zero value otherwise.

### GetApproximateMinimumRequiredFeeOk

`func (o *DecodeRawTransactionHexRISE2) GetApproximateMinimumRequiredFeeOk() (*string, bool)`

GetApproximateMinimumRequiredFeeOk returns a tuple with the ApproximateMinimumRequiredFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISE2) SetApproximateMinimumRequiredFee(v string)`

SetApproximateMinimumRequiredFee sets ApproximateMinimumRequiredFee field to given value.

### HasApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRISE2) HasApproximateMinimumRequiredFee() bool`

HasApproximateMinimumRequiredFee returns a boolean if a field has been set.

### GetGasLimit

`func (o *DecodeRawTransactionHexRISE2) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *DecodeRawTransactionHexRISE2) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *DecodeRawTransactionHexRISE2) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPaidForData

`func (o *DecodeRawTransactionHexRISE2) GetGasPaidForData() string`

GetGasPaidForData returns the GasPaidForData field if non-nil, zero value otherwise.

### GetGasPaidForDataOk

`func (o *DecodeRawTransactionHexRISE2) GetGasPaidForDataOk() (*string, bool)`

GetGasPaidForDataOk returns a tuple with the GasPaidForData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPaidForData

`func (o *DecodeRawTransactionHexRISE2) SetGasPaidForData(v string)`

SetGasPaidForData sets GasPaidForData field to given value.

### HasGasPaidForData

`func (o *DecodeRawTransactionHexRISE2) HasGasPaidForData() bool`

HasGasPaidForData returns a boolean if a field has been set.

### GetGasPrice

`func (o *DecodeRawTransactionHexRISE2) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *DecodeRawTransactionHexRISE2) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *DecodeRawTransactionHexRISE2) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *DecodeRawTransactionHexRISE2) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetInputData

`func (o *DecodeRawTransactionHexRISE2) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *DecodeRawTransactionHexRISE2) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *DecodeRawTransactionHexRISE2) SetInputData(v string)`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *DecodeRawTransactionHexRISE2) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetNonce

`func (o *DecodeRawTransactionHexRISE2) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DecodeRawTransactionHexRISE2) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DecodeRawTransactionHexRISE2) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetR

`func (o *DecodeRawTransactionHexRISE2) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *DecodeRawTransactionHexRISE2) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *DecodeRawTransactionHexRISE2) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *DecodeRawTransactionHexRISE2) HasR() bool`

HasR returns a boolean if a field has been set.

### GetRecipient

`func (o *DecodeRawTransactionHexRISE2) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *DecodeRawTransactionHexRISE2) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *DecodeRawTransactionHexRISE2) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetS

`func (o *DecodeRawTransactionHexRISE2) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *DecodeRawTransactionHexRISE2) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *DecodeRawTransactionHexRISE2) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *DecodeRawTransactionHexRISE2) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSender

`func (o *DecodeRawTransactionHexRISE2) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DecodeRawTransactionHexRISE2) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DecodeRawTransactionHexRISE2) SetSender(v string)`

SetSender sets Sender field to given value.


### GetType

`func (o *DecodeRawTransactionHexRISE2) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecodeRawTransactionHexRISE2) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecodeRawTransactionHexRISE2) SetType(v int32)`

SetType sets Type field to given value.


### GetV

`func (o *DecodeRawTransactionHexRISE2) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *DecodeRawTransactionHexRISE2) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *DecodeRawTransactionHexRISE2) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *DecodeRawTransactionHexRISE2) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *DecodeRawTransactionHexRISE2) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DecodeRawTransactionHexRISE2) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DecodeRawTransactionHexRISE2) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DecodeRawTransactionHexRISE2) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


