# ListUnspentTransactionOutputsByAddressRIBlockchainSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VSize** | **int32** | Represents the virtual size of this transaction | 
**BindingSig** | **string** | It is used to enforce balance of Spend and Output transfers, in order to prevent their replay across transactions. | 
**ExpiryHeight** | **int32** | Represents a block height after which the transaction will expire. | 
**JoinSplitPubKey** | **string** | Represents an encoding of a JoinSplitSig public validating key. | 
**JoinSplitSig** | **string** | Is used to sign transactions that contain at least one JoinSplit description. | 
**Overwintered** | **bool** | \&quot;Overwinter\&quot; is the network upgrade for the Zcash blockchain. | 
**VJoinSplit** | Pointer to [**[]ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVJoinSplit**](ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVJoinSplit.md) | Represents a sequence of JoinSplit descriptions using BCTV14 proofs. | [optional] 
**VShieldedOutput** | Pointer to [**[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput**](ListConfirmedTransactionsByAddressRIBSZVShieldedOutput.md) | Object Array representation of transaction output descriptions | [optional] 
**VShieldedSpend** | [**[]ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVShieldedSpend**](ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVShieldedSpend.md) | Object Array representation of transaction spend descriptions | 
**ValueBalance** | **string** | Defines the transaction value balance. | 
**VersionGroupId** | **string** | Represents the transaction version group ID. | 

## Methods

### NewListUnspentTransactionOutputsByAddressRIBlockchainSpecific

`func NewListUnspentTransactionOutputsByAddressRIBlockchainSpecific(vSize int32, bindingSig string, expiryHeight int32, joinSplitPubKey string, joinSplitSig string, overwintered bool, vShieldedSpend []ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVShieldedSpend, valueBalance string, versionGroupId string, ) *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific`

NewListUnspentTransactionOutputsByAddressRIBlockchainSpecific instantiates a new ListUnspentTransactionOutputsByAddressRIBlockchainSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnspentTransactionOutputsByAddressRIBlockchainSpecificWithDefaults

`func NewListUnspentTransactionOutputsByAddressRIBlockchainSpecificWithDefaults() *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific`

NewListUnspentTransactionOutputsByAddressRIBlockchainSpecificWithDefaults instantiates a new ListUnspentTransactionOutputsByAddressRIBlockchainSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVSize

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVSize() int32`

GetVSize returns the VSize field if non-nil, zero value otherwise.

### GetVSizeOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVSizeOk() (*int32, bool)`

GetVSizeOk returns a tuple with the VSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSize

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetVSize(v int32)`

SetVSize sets VSize field to given value.


### GetBindingSig

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetBindingSig() string`

GetBindingSig returns the BindingSig field if non-nil, zero value otherwise.

### GetBindingSigOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetBindingSigOk() (*string, bool)`

GetBindingSigOk returns a tuple with the BindingSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSig

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetBindingSig(v string)`

SetBindingSig sets BindingSig field to given value.


### GetExpiryHeight

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetExpiryHeight() int32`

GetExpiryHeight returns the ExpiryHeight field if non-nil, zero value otherwise.

### GetExpiryHeightOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetExpiryHeightOk() (*int32, bool)`

GetExpiryHeightOk returns a tuple with the ExpiryHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryHeight

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetExpiryHeight(v int32)`

SetExpiryHeight sets ExpiryHeight field to given value.


### GetJoinSplitPubKey

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetJoinSplitPubKey() string`

GetJoinSplitPubKey returns the JoinSplitPubKey field if non-nil, zero value otherwise.

### GetJoinSplitPubKeyOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetJoinSplitPubKeyOk() (*string, bool)`

GetJoinSplitPubKeyOk returns a tuple with the JoinSplitPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitPubKey

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetJoinSplitPubKey(v string)`

SetJoinSplitPubKey sets JoinSplitPubKey field to given value.


### GetJoinSplitSig

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetJoinSplitSig() string`

GetJoinSplitSig returns the JoinSplitSig field if non-nil, zero value otherwise.

### GetJoinSplitSigOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetJoinSplitSigOk() (*string, bool)`

GetJoinSplitSigOk returns a tuple with the JoinSplitSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSplitSig

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetJoinSplitSig(v string)`

SetJoinSplitSig sets JoinSplitSig field to given value.


### GetOverwintered

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetOverwintered() bool`

GetOverwintered returns the Overwintered field if non-nil, zero value otherwise.

### GetOverwinteredOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetOverwinteredOk() (*bool, bool)`

GetOverwinteredOk returns a tuple with the Overwintered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwintered

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetOverwintered(v bool)`

SetOverwintered sets Overwintered field to given value.


### GetVJoinSplit

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVJoinSplit() []ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVJoinSplit`

GetVJoinSplit returns the VJoinSplit field if non-nil, zero value otherwise.

### GetVJoinSplitOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVJoinSplitOk() (*[]ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVJoinSplit, bool)`

GetVJoinSplitOk returns a tuple with the VJoinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVJoinSplit

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetVJoinSplit(v []ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVJoinSplit)`

SetVJoinSplit sets VJoinSplit field to given value.

### HasVJoinSplit

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) HasVJoinSplit() bool`

HasVJoinSplit returns a boolean if a field has been set.

### GetVShieldedOutput

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVShieldedOutput() []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput`

GetVShieldedOutput returns the VShieldedOutput field if non-nil, zero value otherwise.

### GetVShieldedOutputOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVShieldedOutputOk() (*[]ListConfirmedTransactionsByAddressRIBSZVShieldedOutput, bool)`

GetVShieldedOutputOk returns a tuple with the VShieldedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedOutput

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetVShieldedOutput(v []ListConfirmedTransactionsByAddressRIBSZVShieldedOutput)`

SetVShieldedOutput sets VShieldedOutput field to given value.

### HasVShieldedOutput

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) HasVShieldedOutput() bool`

HasVShieldedOutput returns a boolean if a field has been set.

### GetVShieldedSpend

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVShieldedSpend() []ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVShieldedSpend`

GetVShieldedSpend returns the VShieldedSpend field if non-nil, zero value otherwise.

### GetVShieldedSpendOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVShieldedSpendOk() (*[]ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVShieldedSpend, bool)`

GetVShieldedSpendOk returns a tuple with the VShieldedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVShieldedSpend

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetVShieldedSpend(v []ListUnspentTransactionOutputsByAddressRIBlockchainSpecificVShieldedSpend)`

SetVShieldedSpend sets VShieldedSpend field to given value.


### GetValueBalance

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetValueBalance() string`

GetValueBalance returns the ValueBalance field if non-nil, zero value otherwise.

### GetValueBalanceOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetValueBalanceOk() (*string, bool)`

GetValueBalanceOk returns a tuple with the ValueBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBalance

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetValueBalance(v string)`

SetValueBalance sets ValueBalance field to given value.


### GetVersionGroupId

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVersionGroupId() string`

GetVersionGroupId returns the VersionGroupId field if non-nil, zero value otherwise.

### GetVersionGroupIdOk

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) GetVersionGroupIdOk() (*string, bool)`

GetVersionGroupIdOk returns a tuple with the VersionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroupId

`func (o *ListUnspentTransactionOutputsByAddressRIBlockchainSpecific) SetVersionGroupId(v string)`

SetVersionGroupId sets VersionGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


