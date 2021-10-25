# ListTransactionsByBlockHeightRIBSZVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the transaction output has been spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHeightRIBSZScriptPubKey**](ListTransactionsByBlockHeightRIBSZScriptPubKey.md) |  | 
**Value** | **string** | Represents the specific amount. | 

## Methods

### NewListTransactionsByBlockHeightRIBSZVout

`func NewListTransactionsByBlockHeightRIBSZVout(isSpent bool, scriptPubKey ListTransactionsByBlockHeightRIBSZScriptPubKey, value string, ) *ListTransactionsByBlockHeightRIBSZVout`

NewListTransactionsByBlockHeightRIBSZVout instantiates a new ListTransactionsByBlockHeightRIBSZVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSZVoutWithDefaults

`func NewListTransactionsByBlockHeightRIBSZVoutWithDefaults() *ListTransactionsByBlockHeightRIBSZVout`

NewListTransactionsByBlockHeightRIBSZVoutWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSZVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSZVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHeightRIBSZVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSZVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSZVout) GetScriptPubKey() ListTransactionsByBlockHeightRIBSZScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBSZVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHeightRIBSZScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSZVout) SetScriptPubKey(v ListTransactionsByBlockHeightRIBSZScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSZVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSZVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSZVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


