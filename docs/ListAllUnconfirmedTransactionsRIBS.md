# ListAllUnconfirmedTransactionsRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Defines the version of the transaction. | 
**Vin** | [**[]ListUnconfirmedTransactionsByAddressRIBSZVin**](ListUnconfirmedTransactionsByAddressRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListConfirmedTransactionsByAddressRIBSZVout**](ListConfirmedTransactionsByAddressRIBSZVout.md) | Object Array representation of transaction outputs | 
**Fee** | [**ListAllUnconfirmedTransactionsRIBSECFee**](ListAllUnconfirmedTransactionsRIBSECFee.md) |  | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListConfirmedTransactionsByAddressRIBSBSCGasPrice**](ListConfirmedTransactionsByAddressRIBSBSCGasPrice.md) |  | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Defines the transaction status. | 
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**VJoinSplit** | [**[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit**](ListConfirmedTransactionsByAddressRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput**](ListConfirmedTransactionsByAddressRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend**](ListConfirmedTransactionsByAddressRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBS

`func NewListAllUnconfirmedTransactionsRIBS(locktime int32, size int32, vSize int32, version int32, vin []ListUnconfirmedTransactionsByAddressRIBSZVin, vout []ListConfirmedTransactionsByAddressRIBSZVout, fee ListAllUnconfirmedTransactionsRIBSECFee, gasLimit string, gasPrice ListConfirmedTransactionsByAddressRIBSBSCGasPrice, inputData string, nonce int32, transactionStatus string, bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, overwintered bool, vJoinSplit []ListConfirmedTransactionsByAddressRIBSZVJoinSplit, vShieldedOutput []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, vShieldedSpend []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, valueBalance string, versionGroupId string, ) *ListAllUnconfirmedTransactionsRIBS`

NewListAllUnconfirmedTransactionsRIBS instantiates a new ListAllUnconfirmedTransactionsRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSWithDefaults() *ListAllUnconfirmedTransactionsRIBS`

NewListAllUnconfirmedTransactionsRIBSWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBS) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListAllUnconfirmedTransactionsRIBS) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListAllUnconfirmedTransactionsRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListAllUnconfirmedTransactionsRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVin() []ListUnconfirmedTransactionsByAddressRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVinOk() (*[]ListUnconfirmedTransactionsByAddressRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVin(v []ListUnconfirmedTransactionsByAddressRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVout() []ListConfirmedTransactionsByAddressRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVoutOk() (*[]ListConfirmedTransactionsByAddressRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVout(v []ListConfirmedTransactionsByAddressRIBSZVout)`

SetVout sets Vout field to given value.


### GetFee

`func (o *ListAllUnconfirmedTransactionsRIBS) GetFee() ListAllUnconfirmedTransactionsRIBSECFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetFeeOk() (*ListAllUnconfirmedTransactionsRIBSECFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListAllUnconfirmedTransactionsRIBS) SetFee(v ListAllUnconfirmedTransactionsRIBSECFee)`

SetFee sets Fee field to given value.


### GetGasLimit

`func (o *ListAllUnconfirmedTransactionsRIBS) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListAllUnconfirmedTransactionsRIBS) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListAllUnconfirmedTransactionsRIBS) GetGasPrice() ListConfirmedTransactionsByAddressRIBSBSCGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetGasPriceOk() (*ListConfirmedTransactionsByAddressRIBSBSCGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListAllUnconfirmedTransactionsRIBS) SetGasPrice(v ListConfirmedTransactionsByAddressRIBSBSCGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetInputData

`func (o *ListAllUnconfirmedTransactionsRIBS) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListAllUnconfirmedTransactionsRIBS) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListAllUnconfirmedTransactionsRIBS) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListAllUnconfirmedTransactionsRIBS) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListAllUnconfirmedTransactionsRIBS) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListAllUnconfirmedTransactionsRIBS) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetBindingSig

`func (o *ListAllUnconfirmedTransactionsRIBS) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListAllUnconfirmedTransactionsRIBS) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListAllUnconfirmedTransactionsRIBS) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListAllUnconfirmedTransactionsRIBS) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListAllUnconfirmedTransactionsRIBS) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListAllUnconfirmedTransactionsRIBS) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListAllUnconfirmedTransactionsRIBS) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListAllUnconfirmedTransactionsRIBS) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetOverwintered

`func (o *ListAllUnconfirmedTransactionsRIBS) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListAllUnconfirmedTransactionsRIBS) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetVJoinSplit

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVJoinSplit() []ListConfirmedTransactionsByAddressRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVJoinSplitOk() (*[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVJoinSplit(v []ListConfirmedTransactionsByAddressRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVShieldedOutput() []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVShieldedOutputOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVShieldedOutput(v []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVShieldedSpend() []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVShieldedSpendOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVShieldedSpend(v []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListAllUnconfirmedTransactionsRIBS) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListAllUnconfirmedTransactionsRIBS) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListAllUnconfirmedTransactionsRIBS) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListAllUnconfirmedTransactionsRIBS) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


