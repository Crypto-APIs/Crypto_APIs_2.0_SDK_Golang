# GetWalletTransactionDetailsByTransactionIDRIBSZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Locktime** | **int64** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VJoinSplit** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplit**](GetTransactionDetailsByTransactionIDRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | [optional] 
**VShieldedOutput** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput**](GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | [optional] 
**VShieldedSpend** | Pointer to [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend**](GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | [optional] 
**ValueBalance** | **string** | String representation of the transaction value balance | 
**Version** | **int32** | Represents the transaction version number. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 
**Vin** | [**[]GetWalletTransactionDetailsByTransactionIDRIBSZVin**](GetWalletTransactionDetailsByTransactionIDRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListTransactionsByBlockHeightRIBSZVout**](ListTransactionsByBlockHeightRIBSZVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSZ

`func NewGetWalletTransactionDetailsByTransactionIDRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int64, overwintered bool, size int32, valueBalance string, version int32, versionGroupId string, vin []GetWalletTransactionDetailsByTransactionIDRIBSZVin, vout []ListTransactionsByBlockHeightRIBSZVout, ) *GetWalletTransactionDetailsByTransactionIDRIBSZ`

NewGetWalletTransactionDetailsByTransactionIDRIBSZ instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSZWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSZWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSZ`

NewGetWalletTransactionDetailsByTransactionIDRIBSZWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVJoinSplit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVJoinSplit() []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVJoinSplitOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVJoinSplit(v []GetTransactionDetailsByTransactionIDRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.

### HasVJoinSplit

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) HasVJoinSplit() bool`

HasVJoinSplit returns a boolean if a field has been set.

### GetVShieldedOutput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedOutputOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.

### HasVShieldedOutput

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) HasVShieldedOutput() bool`

HasVShieldedOutput returns a boolean if a field has been set.

### GetVShieldedSpend

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVShieldedSpendOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.

### HasVShieldedSpend

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) HasVShieldedSpend() bool`

HasVShieldedSpend returns a boolean if a field has been set.

### GetValueBalance

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVin() []GetWalletTransactionDetailsByTransactionIDRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVinOk() (*[]GetWalletTransactionDetailsByTransactionIDRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVin(v []GetWalletTransactionDetailsByTransactionIDRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVout() []ListTransactionsByBlockHeightRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSZ) SetVout(v []ListTransactionsByBlockHeightRIBSZVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


