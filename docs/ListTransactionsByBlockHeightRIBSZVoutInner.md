# ListTransactionsByBlockHeightRIBSZVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the transaction output has been spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHeightRIBSZVoutInnerScriptPubKey**](ListTransactionsByBlockHeightRIBSZVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the specific amount. | 

## Methods

### NewListTransactionsByBlockHeightRIBSZVoutInner

`func NewListTransactionsByBlockHeightRIBSZVoutInner(isSpent bool, scriptPubKey ListTransactionsByBlockHeightRIBSZVoutInnerScriptPubKey, value string, ) *ListTransactionsByBlockHeightRIBSZVoutInner`

NewListTransactionsByBlockHeightRIBSZVoutInner instantiates a new ListTransactionsByBlockHeightRIBSZVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSZVoutInnerWithDefaults

`func NewListTransactionsByBlockHeightRIBSZVoutInnerWithDefaults() *ListTransactionsByBlockHeightRIBSZVoutInner`

NewListTransactionsByBlockHeightRIBSZVoutInnerWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSZVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) GetScriptPubKey() ListTransactionsByBlockHeightRIBSZVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) GetScriptPubKeyOk() (*ListTransactionsByBlockHeightRIBSZVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) SetScriptPubKey(v ListTransactionsByBlockHeightRIBSZVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSZVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


