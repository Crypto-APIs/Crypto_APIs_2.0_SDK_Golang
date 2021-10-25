# ListUnconfirmedTransactionsByAddressRIBSZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Locktime** | **int32** | Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**VJoinSplit** | [**[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit**](ListConfirmedTransactionsByAddressRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput**](ListConfirmedTransactionsByAddressRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend**](ListConfirmedTransactionsByAddressRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**Version** | **int32** | Defines the version of the transaction. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 
**Vin** | [**[]ListUnconfirmedTransactionsByAddressRIBSZVin**](ListUnconfirmedTransactionsByAddressRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]ListConfirmedTransactionsByAddressRIBSZVout**](ListConfirmedTransactionsByAddressRIBSZVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSZ

`func NewListUnconfirmedTransactionsByAddressRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int32, overwintered bool, size int32, vJoinSplit []ListConfirmedTransactionsByAddressRIBSZVJoinSplit, vShieldedOutput []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, vShieldedSpend []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, valueBalance string, version int32, versionGroupId string, vin []ListUnconfirmedTransactionsByAddressRIBSZVin, vout []ListConfirmedTransactionsByAddressRIBSZVout, ) *ListUnconfirmedTransactionsByAddressRIBSZ`

NewListUnconfirmedTransactionsByAddressRIBSZ instantiates a new ListUnconfirmedTransactionsByAddressRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSZWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSZWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSZ`

NewListUnconfirmedTransactionsByAddressRIBSZWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVJoinSplit

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVJoinSplit() []ListConfirmedTransactionsByAddressRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVJoinSplitOk() (*[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVJoinSplit(v []ListConfirmedTransactionsByAddressRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVShieldedOutput() []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVShieldedOutputOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVShieldedOutput(v []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVShieldedSpend() []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVShieldedSpendOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVShieldedSpend(v []ListConfirmedTransactionsByAddressRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVin() []ListUnconfirmedTransactionsByAddressRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVinOk() (*[]ListUnconfirmedTransactionsByAddressRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVin(v []ListUnconfirmedTransactionsByAddressRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVout() []ListConfirmedTransactionsByAddressRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) GetVoutOk() (*[]ListConfirmedTransactionsByAddressRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSZ) SetVout(v []ListConfirmedTransactionsByAddressRIBSZVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


