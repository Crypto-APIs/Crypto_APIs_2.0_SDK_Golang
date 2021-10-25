# ListTransactionsByBlockHashRIBSZ

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
**VJoinSplit** | [**[]ListTransactionsByBlockHashRIBSZVJoinSplit**](ListTransactionsByBlockHashRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput**](ListConfirmedTransactionsByAddressRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend**](ListConfirmedTransactionsByAddressRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**Version** | **int32** | Numeric representation of the transaction Represents the transaction version number. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 
**Vin** | [**[]ListTransactionsByBlockHashRIBSZVin**](ListTransactionsByBlockHashRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListConfirmedTransactionsByAddressRIBSZVout**](ListConfirmedTransactionsByAddressRIBSZVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewListTransactionsByBlockHashRIBSZ

`func NewListTransactionsByBlockHashRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int32, overwintered bool, size int32, vJoinSplit []ListTransactionsByBlockHashRIBSZVJoinSplit, vShieldedOutput []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, vShieldedSpend []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, valueBalance string, version int32, versionGroupId string, vin []ListTransactionsByBlockHashRIBSZVin, vout []ListConfirmedTransactionsByAddressRIBSZVout, ) *ListTransactionsByBlockHashRIBSZ`

NewListTransactionsByBlockHashRIBSZ instantiates a new ListTransactionsByBlockHashRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSZWithDefaults

`func NewListTransactionsByBlockHashRIBSZWithDefaults() *ListTransactionsByBlockHashRIBSZ`

NewListTransactionsByBlockHashRIBSZWithDefaults instantiates a new ListTransactionsByBlockHashRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSig

`func (o *ListTransactionsByBlockHashRIBSZ) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListTransactionsByBlockHashRIBSZ) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListTransactionsByBlockHashRIBSZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListTransactionsByBlockHashRIBSZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListTransactionsByBlockHashRIBSZ) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListTransactionsByBlockHashRIBSZ) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListTransactionsByBlockHashRIBSZ) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListTransactionsByBlockHashRIBSZ) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetLocktime

`func (o *ListTransactionsByBlockHashRIBSZ) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHashRIBSZ) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *ListTransactionsByBlockHashRIBSZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListTransactionsByBlockHashRIBSZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSize

`func (o *ListTransactionsByBlockHashRIBSZ) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHashRIBSZ) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVJoinSplit

`func (o *ListTransactionsByBlockHashRIBSZ) GetVJoinSplit() []ListTransactionsByBlockHashRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVJoinSplitOk() (*[]ListTransactionsByBlockHashRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListTransactionsByBlockHashRIBSZ) SetVJoinSplit(v []ListTransactionsByBlockHashRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListTransactionsByBlockHashRIBSZ) GetVShieldedOutput() []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVShieldedOutputOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListTransactionsByBlockHashRIBSZ) SetVShieldedOutput(v []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListTransactionsByBlockHashRIBSZ) GetVShieldedSpend() []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVShieldedSpendOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListTransactionsByBlockHashRIBSZ) SetVShieldedSpend(v []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListTransactionsByBlockHashRIBSZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListTransactionsByBlockHashRIBSZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHashRIBSZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHashRIBSZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *ListTransactionsByBlockHashRIBSZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListTransactionsByBlockHashRIBSZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *ListTransactionsByBlockHashRIBSZ) GetVin() []ListTransactionsByBlockHashRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVinOk() (*[]ListTransactionsByBlockHashRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHashRIBSZ) SetVin(v []ListTransactionsByBlockHashRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashRIBSZ) GetVout() []ListConfirmedTransactionsByAddressRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashRIBSZ) GetVoutOk() (*[]ListConfirmedTransactionsByAddressRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashRIBSZ) SetVout(v []ListConfirmedTransactionsByAddressRIBSZVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


