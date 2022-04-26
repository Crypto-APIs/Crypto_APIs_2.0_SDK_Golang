# GetTransactionDetailsByTransactionIDFromCallbackRIBS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VSize** | **int32** | Represents the virtual size of this transaction. | 
**Version** | **int32** | Defines the version of the transaction. | 
**Vin** | [**[]GetTransactionDetailsByTransactionIDRIBSZVin**](GetTransactionDetailsByTransactionIDRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSZVout**](GetTransactionDetailsByTransactionIDRIBSZVout.md) | Object Array representation of transaction outputs | 
**Contract** | **string** | Represents the specific transaction contract | 
**GasLimit** | **int32** | Represents the maximum amount of gas allowed in the block in order to determine how many transactions it can fit. | 
**GasPrice** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice**](GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice.md) |  | 
**GasUsed** | **int32** | Defines how much of the gas for the block has been used. | 
**InputData** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**VJoinSplit** | [**[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplit**](GetTransactionDetailsByTransactionIDRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput**](GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend**](GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | String representation of the transaction value balance | 
**VersionGroupId** | **string** | Represents the transaction version group ID | 
**AdditionalData** | **string** | Represents additional data that may be needed. | 
**DestinationTag** | Pointer to **int64** | Defines the destination tag value. | [optional] 
**Offer** | [**GetXRPRippleTransactionDetailsByTransactionIDRIOffer**](GetXRPRippleTransactionDetailsByTransactionIDRIOffer.md) |  | 
**Receive** | [**GetXRPRippleTransactionDetailsByTransactionIDRIReceive**](GetXRPRippleTransactionDetailsByTransactionIDRIReceive.md) |  | 
**Sequence** | **int64** | Defines the transaction input&#39;s sequence as an integer, which is is used when transactions are replaced with newer versions before LockTime. | 
**Status** | **string** | Defines the status of the transaction. | 
**Type** | **string** | Defines the type of the transaction. | 
**Value** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue**](GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue.md) |  | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBS

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBS(locktime int64, size int32, vSize int32, version int32, vin []GetTransactionDetailsByTransactionIDRIBSZVin, vout []GetTransactionDetailsByTransactionIDRIBSZVout, contract string, gasLimit int32, gasPrice GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice, gasUsed int32, inputData string, nonce int32, transactionStatus string, bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, overwintered bool, vJoinSplit []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit, vShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, vShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, valueBalance string, versionGroupId string, additionalData string, offer GetXRPRippleTransactionDetailsByTransactionIDRIOffer, receive GetXRPRippleTransactionDetailsByTransactionIDRIReceive, sequence int64, status string, type_ string, value GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBS`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBS instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBS`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetVersion

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVin() []GetTransactionDetailsByTransactionIDRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVinOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVin(v []GetTransactionDetailsByTransactionIDRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVout() []GetTransactionDetailsByTransactionIDRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVout(v []GetTransactionDetailsByTransactionIDRIBSZVout)`

SetVout sets Vout field to given value.


### GetContract

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGasLimit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.


### GetGasPrice

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetGasPrice() GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetGasPriceOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetGasPrice(v GetTransactionDetailsByTransactionIDFromCallbackRIBSZ2GasPrice)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetInputData

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetInputData(v string)`

SetInputData sets InputData field to given value.


### GetNonce

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetBindingSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetOverwintered

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetVJoinSplit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVJoinSplit() []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVJoinSplitOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVJoinSplit(v []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVShieldedOutputOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVShieldedSpendOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetAdditionalData

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.


### GetDestinationTag

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetDestinationTag() int64`

GetDestinationTag returns the DestinationTag field if non-nil, zero value otherwise.

### GetDestinationTagOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetDestinationTagOk() (*int64, bool)`

GetDestinationTagOk returns a tuple with the DestinationTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTag

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetDestinationTag(v int64)`

SetDestinationTag sets DestinationTag field to given value.

### HasDestinationTag

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) HasDestinationTag() bool`

HasDestinationTag returns a boolean if a field has been set.

### GetOffer

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetOffer() GetXRPRippleTransactionDetailsByTransactionIDRIOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetOfferOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetOffer(v GetXRPRippleTransactionDetailsByTransactionIDRIOffer)`

SetOffer sets Offer field to given value.


### GetReceive

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetReceive() GetXRPRippleTransactionDetailsByTransactionIDRIReceive`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetReceiveOk() (*GetXRPRippleTransactionDetailsByTransactionIDRIReceive, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetReceive(v GetXRPRippleTransactionDetailsByTransactionIDRIReceive)`

SetReceive sets Receive field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetValue() GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) GetValueOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBS) SetValue(v GetTransactionDetailsByTransactionIDFromCallbackRIBSXValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


