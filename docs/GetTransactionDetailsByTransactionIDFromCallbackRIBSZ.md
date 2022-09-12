# GetTransactionDetailsByTransactionIDFromCallbackRIBSZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Locktime** | **int64** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VJoinSplit** | [**[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner**](GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner**](GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner**](GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | String representation of the transaction value balance | 
**Version** | **int32** | Defines the version of the transaction. | 
**VersionGroupId** | **string** | Represents the transaction version group ID | 
**Vin** | [**[]GetTransactionDetailsByTransactionIDRIBSZVinInner**](GetTransactionDetailsByTransactionIDRIBSZVinInner.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDFromCallbackRIBSZVoutInner**](GetTransactionDetailsByTransactionIDFromCallbackRIBSZVoutInner.md) | Object Array representation of transaction outputs | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int64, overwintered bool, size int32, vJoinSplit []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner, vShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner, vShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner, valueBalance string, version int32, versionGroupId string, vin []GetTransactionDetailsByTransactionIDRIBSZVinInner, vout []GetTransactionDetailsByTransactionIDFromCallbackRIBSZVoutInner, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZ instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSZWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetLocktime

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVJoinSplit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVJoinSplit() []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVJoinSplitOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVJoinSplit(v []GetTransactionDetailsByTransactionIDRIBSZVJoinSplitInner)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedOutputOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVShieldedSpendOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpendInner)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVin() []GetTransactionDetailsByTransactionIDRIBSZVinInner`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVinOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVinInner, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVin(v []GetTransactionDetailsByTransactionIDRIBSZVinInner)`

SetVin sets Vin field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVout() []GetTransactionDetailsByTransactionIDFromCallbackRIBSZVoutInner`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDFromCallbackRIBSZVoutInner, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSZ) SetVout(v []GetTransactionDetailsByTransactionIDFromCallbackRIBSZVoutInner)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


