# ListConfirmedTransactionsByAddressRIBSD2Vout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey**](ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSD2Vout

`func NewListConfirmedTransactionsByAddressRIBSD2Vout(isSpent bool, scriptPubKey ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey, value string, ) *ListConfirmedTransactionsByAddressRIBSD2Vout`

NewListConfirmedTransactionsByAddressRIBSD2Vout instantiates a new ListConfirmedTransactionsByAddressRIBSD2Vout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSD2VoutWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSD2VoutWithDefaults() *ListConfirmedTransactionsByAddressRIBSD2Vout`

NewListConfirmedTransactionsByAddressRIBSD2VoutWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSD2Vout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetScriptPubKey() ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetScriptPubKeyOk() (*ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) SetScriptPubKey(v ListConfirmedTransactionsByAddressRIBSD2ScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2Vout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


