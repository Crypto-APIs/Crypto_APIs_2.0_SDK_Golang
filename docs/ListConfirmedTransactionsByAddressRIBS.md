# ListConfirmedTransactionsByAddressRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Defines the version of the transaction. | 
**Vin** | [**[]ListConfirmedTransactionsByAddressRIBSZVin**](ListConfirmedTransactionsByAddressRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSZVout**](GetTransactionDetailsByTransactionIDRIBSZVout.md) | Object Array representation of transaction outputs | 
**Contract** | **string** | Represents the specific transaction contract. | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**ListConfirmedTransactionsByAddressRIBSBSCGasPrice**](ListConfirmedTransactionsByAddressRIBSBSCGasPrice.md) |  | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | String representation of the transaction status | 
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**VJoinSplit** | [**[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit**](ListConfirmedTransactionsByAddressRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput**](GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend**](GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBS

`func NewListConfirmedTransactionsByAddressRIBS(locktime int64, size int32, vSize int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSZVin, vout []GetTransactionDetailsByTransactionIDRIBSZVout, contract string, gasLimit string, gasPrice ListConfirmedTransactionsByAddressRIBSBSCGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string, bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, overwintered bool, vJoinSplit []ListConfirmedTransactionsByAddressRIBSZVJoinSplit, vShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, vShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, valueBalance string, versionGroupId string, ) *ListConfirmedTransactionsByAddressRIBS`

NewListConfirmedTransactionsByAddressRIBS instantiates a new ListConfirmedTransactionsByAddressRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSWithDefaults() *ListConfirmedTransactionsByAddressRIBS`

NewListConfirmedTransactionsByAddressRIBSWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListConfirmedTransactionsByAddressRIBS) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListConfirmedTransactionsByAddressRIBS) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListConfirmedTransactionsByAddressRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListConfirmedTransactionsByAddressRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVin() []ListConfirmedTransactionsByAddressRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVinOk() (*[]ListConfirmedTransactionsByAddressRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVin(v []ListConfirmedTransactionsByAddressRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVout() []GetTransactionDetailsByTransactionIDRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVout(v []GetTransactionDetailsByTransactionIDRIBSZVout)`

SetVout sets Vout field to given value.


### GetContract

`func (o *ListConfirmedTransactionsByAddressRIBS) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ListConfirmedTransactionsByAddressRIBS) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *ListConfirmedTransactionsByAddressRIBS) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ListConfirmedTransactionsByAddressRIBS) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *ListConfirmedTransactionsByAddressRIBS) GetGasPrice() ListConfirmedTransactionsByAddressRIBSBSCGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetGasPriceOk() (*ListConfirmedTransactionsByAddressRIBSBSCGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ListConfirmedTransactionsByAddressRIBS) SetGasPrice(v ListConfirmedTransactionsByAddressRIBSBSCGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *ListConfirmedTransactionsByAddressRIBS) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ListConfirmedTransactionsByAddressRIBS) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *ListConfirmedTransactionsByAddressRIBS) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ListConfirmedTransactionsByAddressRIBS) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *ListConfirmedTransactionsByAddressRIBS) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ListConfirmedTransactionsByAddressRIBS) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *ListConfirmedTransactionsByAddressRIBS) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *ListConfirmedTransactionsByAddressRIBS) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetBindingSig

`func (o *ListConfirmedTransactionsByAddressRIBS) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListConfirmedTransactionsByAddressRIBS) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListConfirmedTransactionsByAddressRIBS) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListConfirmedTransactionsByAddressRIBS) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListConfirmedTransactionsByAddressRIBS) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListConfirmedTransactionsByAddressRIBS) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListConfirmedTransactionsByAddressRIBS) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListConfirmedTransactionsByAddressRIBS) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetOverwintered

`func (o *ListConfirmedTransactionsByAddressRIBS) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListConfirmedTransactionsByAddressRIBS) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetVJoinSplit

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVJoinSplit() []ListConfirmedTransactionsByAddressRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVJoinSplitOk() (*[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVJoinSplit(v []ListConfirmedTransactionsByAddressRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVShieldedOutputOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVShieldedSpendOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListConfirmedTransactionsByAddressRIBS) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListConfirmedTransactionsByAddressRIBS) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListConfirmedTransactionsByAddressRIBS) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListConfirmedTransactionsByAddressRIBS) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


