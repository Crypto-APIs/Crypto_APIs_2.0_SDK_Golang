# ListTransactionsByBlockHeightRIBSZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VJoinSplit** | [**[]ListTransactionsByBlockHeightRIBSZVJoinSplit**](ListTransactionsByBlockHeightRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]ListTransactionsByBlockHeightRIBSZVShieldedOutput**](ListTransactionsByBlockHeightRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend**](ListConfirmedTransactionsByAddressRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**Version** | **int32** | Represents the transaction version number. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 
**Vin** | [**[]ListTransactionsByBlockHeightRIBSZVin**](ListTransactionsByBlockHeightRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListTransactionsByBlockHeightRIBSZVout**](ListTransactionsByBlockHeightRIBSZVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewListTransactionsByBlockHeightRIBSZ

`func NewListTransactionsByBlockHeightRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int32, overwintered bool, size int32, vJoinSplit []ListTransactionsByBlockHeightRIBSZVJoinSplit, vShieldedOutput []ListTransactionsByBlockHeightRIBSZVShieldedOutput, vShieldedSpend []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, valueBalance string, version int32, versionGroupId string, vin []ListTransactionsByBlockHeightRIBSZVin, vout []ListTransactionsByBlockHeightRIBSZVout, ) *ListTransactionsByBlockHeightRIBSZ`

NewListTransactionsByBlockHeightRIBSZ instantiates a new ListTransactionsByBlockHeightRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSZWithDefaults

`func NewListTransactionsByBlockHeightRIBSZWithDefaults() *ListTransactionsByBlockHeightRIBSZ`

NewListTransactionsByBlockHeightRIBSZWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSig

`func (o *ListTransactionsByBlockHeightRIBSZ) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListTransactionsByBlockHeightRIBSZ) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListTransactionsByBlockHeightRIBSZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListTransactionsByBlockHeightRIBSZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListTransactionsByBlockHeightRIBSZ) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListTransactionsByBlockHeightRIBSZ) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListTransactionsByBlockHeightRIBSZ) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListTransactionsByBlockHeightRIBSZ) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetLocktime

`func (o *ListTransactionsByBlockHeightRIBSZ) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightRIBSZ) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *ListTransactionsByBlockHeightRIBSZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListTransactionsByBlockHeightRIBSZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightRIBSZ) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightRIBSZ) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVJoinSplit

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVJoinSplit() []ListTransactionsByBlockHeightRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVJoinSplitOk() (*[]ListTransactionsByBlockHeightRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVJoinSplit(v []ListTransactionsByBlockHeightRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVShieldedOutput() []ListTransactionsByBlockHeightRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVShieldedOutputOk() (*[]ListTransactionsByBlockHeightRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVShieldedOutput(v []ListTransactionsByBlockHeightRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVShieldedSpend() []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVShieldedSpendOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVShieldedSpend(v []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListTransactionsByBlockHeightRIBSZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListTransactionsByBlockHeightRIBSZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVin() []ListTransactionsByBlockHeightRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVinOk() (*[]ListTransactionsByBlockHeightRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVin(v []ListTransactionsByBlockHeightRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVout() []ListTransactionsByBlockHeightRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightRIBSZ) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightRIBSZ) SetVout(v []ListTransactionsByBlockHeightRIBSZVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


