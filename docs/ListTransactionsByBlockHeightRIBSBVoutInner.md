# ListTransactionsByBlockHeightRIBSBVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHeightRIBSBVoutInnerScriptPubKey**](ListTransactionsByBlockHeightRIBSBVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByBlockHeightRIBSBVoutInner

`func NewListTransactionsByBlockHeightRIBSBVoutInner(isSpent bool, scriptPubKey ListTransactionsByBlockHeightRIBSBVoutInnerScriptPubKey, value string, ) *ListTransactionsByBlockHeightRIBSBVoutInner`

NewListTransactionsByBlockHeightRIBSBVoutInner instantiates a new ListTransactionsByBlockHeightRIBSBVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSBVoutInnerWithDefaults

`func NewListTransactionsByBlockHeightRIBSBVoutInnerWithDefaults() *ListTransactionsByBlockHeightRIBSBVoutInner`

NewListTransactionsByBlockHeightRIBSBVoutInnerWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSBVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) GetScriptPubKey() ListTransactionsByBlockHeightRIBSBVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) GetScriptPubKeyOk() (*ListTransactionsByBlockHeightRIBSBVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) SetScriptPubKey(v ListTransactionsByBlockHeightRIBSBVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSBVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


