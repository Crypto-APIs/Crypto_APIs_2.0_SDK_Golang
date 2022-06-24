# DecodeRawTransactionHexRIS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**TransactionHash** | **string** | Represents the same as transactionId for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols hash is different from transactionId for SegWit transactions. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]DecodeRawTransactionHexRISZVinInner**](DecodeRawTransactionHexRISZVinInner.md) | Represents the Inputs of the transaction | 
**Vout** | [**[]DecodeRawTransactionHexRISZVoutInner**](DecodeRawTransactionHexRISZVoutInner.md) | Represents the Inputs of the transaction | 
**Weight** | Pointer to **int32** | Represents the size of a block, measured in weight units and including the segwit discount. | [optional] 
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
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**Saplinged** | **bool** | Defines if the transaction includes sapling or not. | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**VersionGroupId** | **string** | Represents the transaction version group ID | 

## Methods

### NewDecodeRawTransactionHexRIS

`func NewDecodeRawTransactionHexRIS(locktime int32, transactionHash string, vSize int32, version int32, vin []DecodeRawTransactionHexRISZVinInner, vout []DecodeRawTransactionHexRISZVoutInner, gasLimit string, nonce int32, recipient string, sender string, type_ int32, expiryHeight int32, overwintered bool, saplinged bool, valueBalance string, versionGroupId string, ) *DecodeRawTransactionHexRIS`

NewDecodeRawTransactionHexRIS instantiates a new DecodeRawTransactionHexRIS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISWithDefaults

`func NewDecodeRawTransactionHexRISWithDefaults() *DecodeRawTransactionHexRIS`

NewDecodeRawTransactionHexRISWithDefaults instantiates a new DecodeRawTransactionHexRIS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *DecodeRawTransactionHexRIS) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *DecodeRawTransactionHexRIS) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *DecodeRawTransactionHexRIS) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetTransactionHash

`func (o *DecodeRawTransactionHexRIS) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *DecodeRawTransactionHexRIS) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *DecodeRawTransactionHexRIS) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetVSize

`func (o *DecodeRawTransactionHexRIS) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *DecodeRawTransactionHexRIS) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *DecodeRawTransactionHexRIS) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *DecodeRawTransactionHexRIS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecodeRawTransactionHexRIS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecodeRawTransactionHexRIS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *DecodeRawTransactionHexRIS) GetVin() []DecodeRawTransactionHexRISZVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *DecodeRawTransactionHexRIS) GetVinOk() (*[]DecodeRawTransactionHexRISZVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *DecodeRawTransactionHexRIS) SetVin(v []DecodeRawTransactionHexRISZVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *DecodeRawTransactionHexRIS) GetVout() []DecodeRawTransactionHexRISZVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *DecodeRawTransactionHexRIS) GetVoutOk() (*[]DecodeRawTransactionHexRISZVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *DecodeRawTransactionHexRIS) SetVout(v []DecodeRawTransactionHexRISZVoutInner)`

SetVout sets Vout field to given value.


### GetWeight

`func (o *DecodeRawTransactionHexRIS) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DecodeRawTransactionHexRIS) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DecodeRawTransactionHexRIS) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DecodeRawTransactionHexRIS) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetApproximateFee

`func (o *DecodeRawTransactionHexRIS) GetApproximateFee() string`

GetApproximateFee returns the ApproximateFee field if non-nil, zero value otherwise.

### GetApproximateFeeOk

`func (o *DecodeRawTransactionHexRIS) GetApproximateFeeOk() (*string, bool)`

GetApproximateFeeOk returns a tuple with the ApproximateFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateFee

`func (o *DecodeRawTransactionHexRIS) SetApproximateFee(v string)`

SetApproximateFee sets ApproximateFee field to given value.

### HasApproximateFee

`func (o *DecodeRawTransactionHexRIS) HasApproximateFee() bool`

HasApproximateFee returns a boolean if a field has been set.

### GetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRIS) GetApproximateMinimumRequiredFee() string`

GetApproximateMinimumRequiredFee returns the ApproximateMinimumRequiredFee field if non-nil, zero value otherwise.

### GetApproximateMinimumRequiredFeeOk

`func (o *DecodeRawTransactionHexRIS) GetApproximateMinimumRequiredFeeOk() (*string, bool)`

GetApproximateMinimumRequiredFeeOk returns a tuple with the ApproximateMinimumRequiredFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRIS) SetApproximateMinimumRequiredFee(v string)`

SetApproximateMinimumRequiredFee sets ApproximateMinimumRequiredFee field to given value.

### HasApproximateMinimumRequiredFee

`func (o *DecodeRawTransactionHexRIS) HasApproximateMinimumRequiredFee() bool`

HasApproximateMinimumRequiredFee returns a boolean if a field has been set.

### GetGasLimit

`func (o *DecodeRawTransactionHexRIS) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *DecodeRawTransactionHexRIS) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *DecodeRawTransactionHexRIS) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPaidForData

`func (o *DecodeRawTransactionHexRIS) GetGasPaidForData() string`

GetGasPaidForData returns the GasPaidForData field if non-nil, zero value otherwise.

### GetGasPaidForDataOk

`func (o *DecodeRawTransactionHexRIS) GetGasPaidForDataOk() (*string, bool)`

GetGasPaidForDataOk returns a tuple with the GasPaidForData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPaidForData

`func (o *DecodeRawTransactionHexRIS) SetGasPaidForData(v string)`

SetGasPaidForData sets GasPaidForData field to given value.

### HasGasPaidForData

`func (o *DecodeRawTransactionHexRIS) HasGasPaidForData() bool`

HasGasPaidForData returns a boolean if a field has been set.

### GetGasPrice

`func (o *DecodeRawTransactionHexRIS) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *DecodeRawTransactionHexRIS) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *DecodeRawTransactionHexRIS) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *DecodeRawTransactionHexRIS) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetInputData

`func (o *DecodeRawTransactionHexRIS) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *DecodeRawTransactionHexRIS) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *DecodeRawTransactionHexRIS) SetInputData(v string)`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *DecodeRawTransactionHexRIS) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *DecodeRawTransactionHexRIS) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *DecodeRawTransactionHexRIS) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *DecodeRawTransactionHexRIS) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *DecodeRawTransactionHexRIS) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxFeePriorityPerGas

`func (o *DecodeRawTransactionHexRIS) GetMaxFeePriorityPerGas() string`

GetMaxFeePriorityPerGas returns the MaxFeePriorityPerGas field if non-nil, zero value otherwise.

### GetMaxFeePriorityPerGasOk

`func (o *DecodeRawTransactionHexRIS) GetMaxFeePriorityPerGasOk() (*string, bool)`

GetMaxFeePriorityPerGasOk returns a tuple with the MaxFeePriorityPerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePriorityPerGas

`func (o *DecodeRawTransactionHexRIS) SetMaxFeePriorityPerGas(v string)`

SetMaxFeePriorityPerGas sets MaxFeePriorityPerGas field to given value.

### HasMaxFeePriorityPerGas

`func (o *DecodeRawTransactionHexRIS) HasMaxFeePriorityPerGas() bool`

HasMaxFeePriorityPerGas returns a boolean if a field has been set.

### GetNonce

`func (o *DecodeRawTransactionHexRIS) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DecodeRawTransactionHexRIS) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DecodeRawTransactionHexRIS) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetR

`func (o *DecodeRawTransactionHexRIS) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *DecodeRawTransactionHexRIS) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *DecodeRawTransactionHexRIS) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *DecodeRawTransactionHexRIS) HasR() bool`

HasR returns a boolean if a field has been set.

### GetRecipient

`func (o *DecodeRawTransactionHexRIS) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *DecodeRawTransactionHexRIS) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *DecodeRawTransactionHexRIS) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetS

`func (o *DecodeRawTransactionHexRIS) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *DecodeRawTransactionHexRIS) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *DecodeRawTransactionHexRIS) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *DecodeRawTransactionHexRIS) HasS() bool`

HasS returns a boolean if a field has been set.

### GetSender

`func (o *DecodeRawTransactionHexRIS) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DecodeRawTransactionHexRIS) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DecodeRawTransactionHexRIS) SetSender(v string)`

SetSender sets Sender field to given value.


### GetType

`func (o *DecodeRawTransactionHexRIS) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecodeRawTransactionHexRIS) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecodeRawTransactionHexRIS) SetType(v int32)`

SetType sets Type field to given value.


### GetV

`func (o *DecodeRawTransactionHexRIS) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *DecodeRawTransactionHexRIS) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *DecodeRawTransactionHexRIS) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *DecodeRawTransactionHexRIS) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *DecodeRawTransactionHexRIS) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DecodeRawTransactionHexRIS) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DecodeRawTransactionHexRIS) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DecodeRawTransactionHexRIS) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExpiryHeight

`func (o *DecodeRawTransactionHexRIS) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *DecodeRawTransactionHexRIS) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *DecodeRawTransactionHexRIS) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetOverwintered

`func (o *DecodeRawTransactionHexRIS) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *DecodeRawTransactionHexRIS) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *DecodeRawTransactionHexRIS) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSaplinged

`func (o *DecodeRawTransactionHexRIS) GetSaplinged() bool`

GetSaplinged returns the Saplinged field if non-nil, zero value otherwise.

### GetSaplingedOk

`func (o *DecodeRawTransactionHexRIS) GetSaplingedOk() (*bool, bool)`

GetSaplingedOk returns a tuple with the Saplinged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaplinged

`func (o *DecodeRawTransactionHexRIS) SetSaplinged(v bool)`

SetSaplinged sets Saplinged field to given value.


### GetValueBalance

`func (o *DecodeRawTransactionHexRIS) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *DecodeRawTransactionHexRIS) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *DecodeRawTransactionHexRIS) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *DecodeRawTransactionHexRIS) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *DecodeRawTransactionHexRIS) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *DecodeRawTransactionHexRIS) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


