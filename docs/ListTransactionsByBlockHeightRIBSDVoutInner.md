# ListTransactionsByBlockHeightRIBSDVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHeightRIBSDVoutInnerScriptPubKey**](ListTransactionsByBlockHeightRIBSDVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByBlockHeightRIBSDVoutInner

`func NewListTransactionsByBlockHeightRIBSDVoutInner(isSpent bool, scriptPubKey ListTransactionsByBlockHeightRIBSDVoutInnerScriptPubKey, value string, ) *ListTransactionsByBlockHeightRIBSDVoutInner`

NewListTransactionsByBlockHeightRIBSDVoutInner instantiates a new ListTransactionsByBlockHeightRIBSDVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSDVoutInnerWithDefaults

`func NewListTransactionsByBlockHeightRIBSDVoutInnerWithDefaults() *ListTransactionsByBlockHeightRIBSDVoutInner`

NewListTransactionsByBlockHeightRIBSDVoutInnerWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSDVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) GetScriptPubKey() ListTransactionsByBlockHeightRIBSDVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) GetScriptPubKeyOk() (*ListTransactionsByBlockHeightRIBSDVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) SetScriptPubKey(v ListTransactionsByBlockHeightRIBSDVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSDVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


