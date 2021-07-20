# ListTransactionsByBlockHeightRIBSLVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHeightRIBSLScriptPubKey**](ListTransactionsByBlockHeightRIBSLScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByBlockHeightRIBSLVout

`func NewListTransactionsByBlockHeightRIBSLVout(isSpent bool, scriptPubKey ListTransactionsByBlockHeightRIBSLScriptPubKey, value string, ) *ListTransactionsByBlockHeightRIBSLVout`

NewListTransactionsByBlockHeightRIBSLVout instantiates a new ListTransactionsByBlockHeightRIBSLVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSLVoutWithDefaults

`func NewListTransactionsByBlockHeightRIBSLVoutWithDefaults() *ListTransactionsByBlockHeightRIBSLVout`

NewListTransactionsByBlockHeightRIBSLVoutWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSLVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSLVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHeightRIBSLVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSLVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSLVout) GetScriptPubKey() ListTransactionsByBlockHeightRIBSLScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBSLVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHeightRIBSLScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSLVout) SetScriptPubKey(v ListTransactionsByBlockHeightRIBSLScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSLVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSLVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSLVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


