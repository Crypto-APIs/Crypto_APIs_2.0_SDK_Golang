# ListTransactionsByBlockHeightRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHeightRIBSZVin**](ListTransactionsByBlockHeightRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListTransactionsByBlockHeightRIBSZVout**](ListTransactionsByBlockHeightRIBSZVout.md) | Object Array representation of transaction outputs | 
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListTransactionsByBlockHeightRIBSBSCGasPrice**](ListTransactionsByBlockHeightRIBSBSCGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Represents the status of this transaction | 
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**VJoinSplit** | [**[]ListTransactionsByBlockHeightRIBSZVJoinSplit**](ListTransactionsByBlockHeightRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]ListTransactionsByBlockHeightRIBSZVShieldedOutput**](ListTransactionsByBlockHeightRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend**](ListConfirmedTransactionsByAddressRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 

## Methods

### NewListTransactionsByBlockHeightRIBS

`func NewListTransactionsByBlockHeightRIBS(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByBlockHeightRIBSZVin, vout []ListTransactionsByBlockHeightRIBSZVout, contract string, gasLimit string, gasPrice ListTransactionsByBlockHeightRIBSBSCGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string, bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, overwintered bool, vJoinSplit []ListTransactionsByBlockHeightRIBSZVJoinSplit, vShieldedOutput []ListTransactionsByBlockHeightRIBSZVShieldedOutput, vShieldedSpend []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, valueBalance string, versionGroupId string, ) *ListTransactionsByBlockHeightRIBS`

NewListTransactionsByBlockHeightRIBS instantiates a new ListTransactionsByBlockHeightRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSWithDefaults

`func NewListTransactionsByBlockHeightRIBSWithDefaults() *ListTransactionsByBlockHeightRIBS`

NewListTransactionsByBlockHeightRIBSWithDefaults instantiates a new ListTransactionsByBlockHeightRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightRIBS) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightRIBS) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightRIBS) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListTransactionsByBlockHeightRIBS) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListTransactionsByBlockHeightRIBS) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightRIBS) GetVin() []ListTransactionsByBlockHeightRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVinOk() (*[]ListTransactionsByBlockHeightRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightRIBS) SetVin(v []ListTransactionsByBlockHeightRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightRIBS) GetVout() []ListTransactionsByBlockHeightRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightRIBS) SetVout(v []ListTransactionsByBlockHeightRIBSZVout)`

SetVout sets Vout field to given value.


### GetContract

`func (o *ListTransactionsByBlockHeightRIBS) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListTransactionsByBlockHeightRIBS) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListTransactionsByBlockHeightRIBS) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListTransactionsByBlockHeightRIBS) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListTransactionsByBlockHeightRIBS) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListTransactionsByBlockHeightRIBS) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListTransactionsByBlockHeightRIBS) GetGasPrice() ListTransactionsByBlockHeightRIBSBSCGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListTransactionsByBlockHeightRIBS) GetGasPriceOk() (*ListTransactionsByBlockHeightRIBSBSCGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListTransactionsByBlockHeightRIBS) SetGasPrice(v ListTransactionsByBlockHeightRIBSBSCGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListTransactionsByBlockHeightRIBS) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListTransactionsByBlockHeightRIBS) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListTransactionsByBlockHeightRIBS) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListTransactionsByBlockHeightRIBS) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListTransactionsByBlockHeightRIBS) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListTransactionsByBlockHeightRIBS) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListTransactionsByBlockHeightRIBS) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListTransactionsByBlockHeightRIBS) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListTransactionsByBlockHeightRIBS) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListTransactionsByBlockHeightRIBS) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListTransactionsByBlockHeightRIBS) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListTransactionsByBlockHeightRIBS) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetBindingSig

`func (o *ListTransactionsByBlockHeightRIBS) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListTransactionsByBlockHeightRIBS) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListTransactionsByBlockHeightRIBS) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListTransactionsByBlockHeightRIBS) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListTransactionsByBlockHeightRIBS) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListTransactionsByBlockHeightRIBS) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListTransactionsByBlockHeightRIBS) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBS) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListTransactionsByBlockHeightRIBS) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListTransactionsByBlockHeightRIBS) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListTransactionsByBlockHeightRIBS) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListTransactionsByBlockHeightRIBS) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetOverwintered

`func (o *ListTransactionsByBlockHeightRIBS) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListTransactionsByBlockHeightRIBS) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListTransactionsByBlockHeightRIBS) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetVJoinSplit

`func (o *ListTransactionsByBlockHeightRIBS) GetVJoinSplit() []ListTransactionsByBlockHeightRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVJoinSplitOk() (*[]ListTransactionsByBlockHeightRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListTransactionsByBlockHeightRIBS) SetVJoinSplit(v []ListTransactionsByBlockHeightRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListTransactionsByBlockHeightRIBS) GetVShieldedOutput() []ListTransactionsByBlockHeightRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVShieldedOutputOk() (*[]ListTransactionsByBlockHeightRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListTransactionsByBlockHeightRIBS) SetVShieldedOutput(v []ListTransactionsByBlockHeightRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListTransactionsByBlockHeightRIBS) GetVShieldedSpend() []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVShieldedSpendOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListTransactionsByBlockHeightRIBS) SetVShieldedSpend(v []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListTransactionsByBlockHeightRIBS) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListTransactionsByBlockHeightRIBS) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListTransactionsByBlockHeightRIBS) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *ListTransactionsByBlockHeightRIBS) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListTransactionsByBlockHeightRIBS) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListTransactionsByBlockHeightRIBS) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


