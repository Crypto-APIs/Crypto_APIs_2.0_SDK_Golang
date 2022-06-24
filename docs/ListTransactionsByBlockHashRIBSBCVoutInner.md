# ListTransactionsByBlockHashRIBSBCVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHashRIBSBCVoutInnerScriptPubKey**](ListTransactionsByBlockHashRIBSBCVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByBlockHashRIBSBCVoutInner

`func NewListTransactionsByBlockHashRIBSBCVoutInner(isSpent bool, scriptPubKey ListTransactionsByBlockHashRIBSBCVoutInnerScriptPubKey, value string, ) *ListTransactionsByBlockHashRIBSBCVoutInner`

NewListTransactionsByBlockHashRIBSBCVoutInner instantiates a new ListTransactionsByBlockHashRIBSBCVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSBCVoutInnerWithDefaults

`func NewListTransactionsByBlockHashRIBSBCVoutInnerWithDefaults() *ListTransactionsByBlockHashRIBSBCVoutInner`

NewListTransactionsByBlockHashRIBSBCVoutInnerWithDefaults instantiates a new ListTransactionsByBlockHashRIBSBCVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) GetScriptPubKey() ListTransactionsByBlockHashRIBSBCVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) GetScriptPubKeyOk() (*ListTransactionsByBlockHashRIBSBCVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) SetScriptPubKey(v ListTransactionsByBlockHashRIBSBCVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHashRIBSBCVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


