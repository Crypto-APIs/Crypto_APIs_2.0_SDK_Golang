# ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ

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
**VJoinSplit** | [**[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit**](ListConfirmedTransactionsByAddressRIBSZVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | 
**VShieldedOutput** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput**](GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | 
**VShieldedSpend** | [**[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend**](GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**Version** | **int32** | Defines the version of the transaction. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 
**Vin** | [**[]ListConfirmedTransactionsByAddressRIBSZVin**](ListConfirmedTransactionsByAddressRIBSZVin.md) | Object Array representation of transaction inputs | 
**Vout** | [**[]GetTransactionDetailsByTransactionIDRIBSZVout**](GetTransactionDetailsByTransactionIDRIBSZVout.md) | Object Array representation of transaction outputs | 

## Methods

### NewListConfirmedTransactionsByAddressAndTimeRangeRIBSZ

`func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSZ(bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, locktime int64, overwintered bool, size int32, vJoinSplit []ListConfirmedTransactionsByAddressRIBSZVJoinSplit, vShieldedOutput []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, vShieldedSpend []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, valueBalance string, version int32, versionGroupId string, vin []ListConfirmedTransactionsByAddressRIBSZVin, vout []GetTransactionDetailsByTransactionIDRIBSZVout, ) *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ`

NewListConfirmedTransactionsByAddressAndTimeRangeRIBSZ instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressAndTimeRangeRIBSZWithDefaults

`func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSZWithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ`

NewListConfirmedTransactionsByAddressAndTimeRangeRIBSZWithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSig

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetLocktime

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetLocktime() int64`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetLocktimeOk() (*int64, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetLocktime(v int64)`

SetLocktime sets Locktime field to given value.


### GetOverwintered

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetSize

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVJoinSplit

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVJoinSplit() []ListConfirmedTransactionsByAddressRIBSZVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVJoinSplitOk() (*[]ListConfirmedTransactionsByAddressRIBSZVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVJoinSplit(v []ListConfirmedTransactionsByAddressRIBSZVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.


### GetVShieldedOutput

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVShieldedOutput() []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVShieldedOutputOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVShieldedOutput(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.


### GetVShieldedSpend

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVShieldedSpend() []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVShieldedSpendOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVShieldedSpend(v []GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersion

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersionGroupId

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.


### GetVin

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVin() []ListConfirmedTransactionsByAddressRIBSZVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVinOk() (*[]ListConfirmedTransactionsByAddressRIBSZVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVin(v []ListConfirmedTransactionsByAddressRIBSZVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVout() []GetTransactionDetailsByTransactionIDRIBSZVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSZVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSZ) SetVout(v []GetTransactionDetailsByTransactionIDRIBSZVout)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


