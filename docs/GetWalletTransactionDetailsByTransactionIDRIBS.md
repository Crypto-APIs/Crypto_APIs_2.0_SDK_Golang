# GetWalletTransactionDetailsByTransactionIDRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSZVinInner**](GetWalletTransactionDetailsByTransactionIDRIBSZVinInner.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListTransactionsByBlockHeightRIBSZVoutInner**](ListTransactionsByBlockHeightRIBSZVoutInner.md) | Object Array representation of transaction outputs | 
**Contract** | **string** | Represents the specific transaction contract | 
**GasLimit** | **string** | Represents the amount of gas used by this specific transaction alone. | 
**GasPrice** | [**GetWalletTransactionDetailsByTransactionIDRIBSPGasPrice**](GetWalletTransactionDetailsByTransactionIDRIBSPGasPrice.md) |  | 
**GasUsed** | **string** | Defines the unit of the gas price amount, e.g. BTC, ETH, XRP. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | String representation of the transaction status | 
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**VJoinSplit** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner**](GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | [optional] 
**VShieldedOutput** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner**](GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner.md) | Object Array representation of transaction output descriptions | [optional] 
**VShieldedSpend** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner**](GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner.md) | Object Array representation of transaction spend descriptions | [optional] 
**ValueBalance** | **string** | String representation of the transaction value balance | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 
**Amount** | **string** | String representation of the amount value | 
**BandwidthUsed** | **string** | Numeric representation of the transaction used bandwidth | 
**EnergyUsed** | **string** | String representation of the transaction used energy | 
**HasInternalTransactions** | **bool** |  | 
**HasTokenTransfers** | **bool** |  | 
**Input** | **string** | Numeric representation of the transaction input | 
**Status** | **string** | String representation of the transaction status | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBS

`func NewGetWalletTransactionDetailsByTransactionIDRIBS(locktime int64, size int32, vSize int32, version int32, vin []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner, vout []ListTransactionsByBlockHeightRIBSZVoutInner, contract string, gasLimit string, gasPrice GetWalletTransactionDetailsByTransactionIDRIBSPGasPrice, gasUsed string, inputData string, nonce int32, transactionStatus string, bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, overwintered bool, valueBalance string, versionGroupId string, amount string, bandwidthUsed string, energyUsed string, hasInternalTransactions bool, hasTokenTransfers bool, input string, status string, ) *GetWalletTransactionDetailsByTransactionIDRIBS`

NewGetWalletTransactionDetailsByTransactionIDRIBS instantiates a new GetWalletTransactionDetailsByTransactionIDRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBS`

NewGetWalletTransactionDetailsByTransactionIDRIBSWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSZVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSZVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVout() []ListTransactionsByBlockHeightRIBSZVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSZVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVout(v []ListTransactionsByBlockHeightRIBSZVoutInner)`

SetVout sets Vout field to given value.


### GetContract

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetGasPrice() GetWalletTransactionDetailsByTransactionIDRIBSPGasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetGasPriceOk() (*GetWalletTransactionDetailsByTransactionIDRIBSPGasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetGasPrice(v GetWalletTransactionDetailsByTransactionIDRIBSPGasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetBindingSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetOverwintered

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetVJoinSplit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVJoinSplit() []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVJoinSplitOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVJoinSplit(v []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner)`

SetVJoinSplit sets VJoinSplit field to given value.

### HasVJoinSplit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) HasVJoinSplit() bool`

HasVJoinSplit returns a boolean if a field has been set.

### GetVShieldedOutput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVShieldedOutputOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner)`

SetVShieldedOutput sets VShieldedOutput field to given value.

### HasVShieldedOutput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) HasVShieldedOutput() bool`

HasVShieldedOutput returns a boolean if a field has been set.

### GetVShieldedSpend

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVShieldedSpendOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner)`

SetVShieldedSpend sets VShieldedSpend field to given value.

### HasVShieldedSpend

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) HasVShieldedSpend() bool`

HasVShieldedSpend returns a boolean if a field has been set.

### GetValueBalance

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetAmount

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBandwidthUsed

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetBandwidthUsed() string`

GetBandwidthUsed returns the BandwidthUsed field if non-nil, zero value otherwise.

### GetBandwidthUsedOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetBandwidthUsedOk() (*string, bool)`

GetBandwidthUsedOk returns a tuple with the BandwidthUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthUsed

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetBandwidthUsed(v string)`

SetBandwidthUsed sets BandwidthUsed field to given value.


### GetEnergyUsed

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetEnergyUsed() string`

GetEnergyUsed returns the EnergyUsed field if non-nil, zero value otherwise.

### GetEnergyUsedOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetEnergyUsedOk() (*string, bool)`

GetEnergyUsedOk returns a tuple with the EnergyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUsed

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetEnergyUsed(v string)`

SetEnergyUsed sets EnergyUsed field to given value.


### GetHasInternalTransactions

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetHasInternalTransactions() bool`

GetHasInternalTransactions returns the HasInternalTransactions field if non-nil, zero value otherwise.

### GetHasInternalTransactionsOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetHasInternalTransactionsOk() (*bool, bool)`

GetHasInternalTransactionsOk returns a tuple with the HasInternalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInternalTransactions

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetHasInternalTransactions(v bool)`

SetHasInternalTransactions sets HasInternalTransactions field to given value.


### GetHasTokenTransfers

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetHasTokenTransfers() bool`

GetHasTokenTransfers returns the HasTokenTransfers field if non-nil, zero value otherwise.

### GetHasTokenTransfersOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetHasTokenTransfersOk() (*bool, bool)`

GetHasTokenTransfersOk returns a tuple with the HasTokenTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTokenTransfers

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetHasTokenTransfers(v bool)`

SetHasTokenTransfers sets HasTokenTransfers field to given value.


### GetInput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetInput(v string)`

SetInput sets Input field to given value.


### GetStatus

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWalletTransactionDetailsByTransactionIDRIBS) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


