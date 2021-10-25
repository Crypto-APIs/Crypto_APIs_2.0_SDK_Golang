# ListConfirmedTransactionsByAddressRIBSZVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the transaction output has been spent or not. | 
**ScriptPubKey** | [**ListConfirmedTransactionsByAddressRIBSZScriptPubKey**](ListConfirmedTransactionsByAddressRIBSZScriptPubKey.md) |  | 
**Value** | **string** | Represents the specific amount. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSZVout

`func NewListConfirmedTransactionsByAddressRIBSZVout(isSpent bool, scriptPubKey ListConfirmedTransactionsByAddressRIBSZScriptPubKey, value string, ) *ListConfirmedTransactionsByAddressRIBSZVout`

NewListConfirmedTransactionsByAddressRIBSZVout instantiates a new ListConfirmedTransactionsByAddressRIBSZVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSZVoutWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSZVoutWithDefaults() *ListConfirmedTransactionsByAddressRIBSZVout`

NewListConfirmedTransactionsByAddressRIBSZVoutWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSZVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) GetScriptPubKey() ListConfirmedTransactionsByAddressRIBSZScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) GetScriptPubKeyOk() (*ListConfirmedTransactionsByAddressRIBSZScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) SetScriptPubKey(v ListConfirmedTransactionsByAddressRIBSZScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSZVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


