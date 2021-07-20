# ListTransactionsByAddressRIBSD2Vout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByAddressRIBSD2ScriptPubKey**](ListTransactionsByAddressRIBSD2ScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByAddressRIBSD2Vout

`func NewListTransactionsByAddressRIBSD2Vout(isSpent bool, scriptPubKey ListTransactionsByAddressRIBSD2ScriptPubKey, value string, ) *ListTransactionsByAddressRIBSD2Vout`

NewListTransactionsByAddressRIBSD2Vout instantiates a new ListTransactionsByAddressRIBSD2Vout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressRIBSD2VoutWithDefaults

`func NewListTransactionsByAddressRIBSD2VoutWithDefaults() *ListTransactionsByAddressRIBSD2Vout`

NewListTransactionsByAddressRIBSD2VoutWithDefaults instantiates a new ListTransactionsByAddressRIBSD2Vout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByAddressRIBSD2Vout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByAddressRIBSD2Vout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByAddressRIBSD2Vout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByAddressRIBSD2Vout) GetScriptPubKey() ListTransactionsByAddressRIBSD2ScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByAddressRIBSD2Vout) GetScriptPubKeyOk() (*ListTransactionsByAddressRIBSD2ScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByAddressRIBSD2Vout) SetScriptPubKey(v ListTransactionsByAddressRIBSD2ScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByAddressRIBSD2Vout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByAddressRIBSD2Vout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByAddressRIBSD2Vout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


