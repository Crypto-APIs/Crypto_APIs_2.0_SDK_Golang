# ListConfirmedTransactionsByAddressRIBSD2VoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey**](ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSD2VoutInner

`func NewListConfirmedTransactionsByAddressRIBSD2VoutInner(isSpent bool, scriptPubKey ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey, value string, ) *ListConfirmedTransactionsByAddressRIBSD2VoutInner`

NewListConfirmedTransactionsByAddressRIBSD2VoutInner instantiates a new ListConfirmedTransactionsByAddressRIBSD2VoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSD2VoutInnerWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSD2VoutInnerWithDefaults() *ListConfirmedTransactionsByAddressRIBSD2VoutInner`

NewListConfirmedTransactionsByAddressRIBSD2VoutInnerWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSD2VoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) GetScriptPubKey() ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) GetScriptPubKeyOk() (*ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) SetScriptPubKey(v ListConfirmedTransactionsByAddressRIBSD2VoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2VoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


