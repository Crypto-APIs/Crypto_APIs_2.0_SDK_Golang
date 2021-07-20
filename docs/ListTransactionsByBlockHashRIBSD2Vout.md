# ListTransactionsByBlockHashRIBSD2Vout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHashRIBSD2ScriptPubKey**](ListTransactionsByBlockHashRIBSD2ScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByBlockHashRIBSD2Vout

`func NewListTransactionsByBlockHashRIBSD2Vout(isSpent bool, scriptPubKey ListTransactionsByBlockHashRIBSD2ScriptPubKey, value string, ) *ListTransactionsByBlockHashRIBSD2Vout`

NewListTransactionsByBlockHashRIBSD2Vout instantiates a new ListTransactionsByBlockHashRIBSD2Vout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashRIBSD2VoutWithDefaults

`func NewListTransactionsByBlockHashRIBSD2VoutWithDefaults() *ListTransactionsByBlockHashRIBSD2Vout`

NewListTransactionsByBlockHashRIBSD2VoutWithDefaults instantiates a new ListTransactionsByBlockHashRIBSD2Vout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHashRIBSD2Vout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHashRIBSD2Vout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHashRIBSD2Vout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHashRIBSD2Vout) GetScriptPubKey() ListTransactionsByBlockHashRIBSD2ScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHashRIBSD2Vout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashRIBSD2ScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHashRIBSD2Vout) SetScriptPubKey(v ListTransactionsByBlockHashRIBSD2ScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHashRIBSD2Vout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHashRIBSD2Vout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHashRIBSD2Vout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


