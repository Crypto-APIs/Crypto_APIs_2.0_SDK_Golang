# ListUnconfirmedTransactionsByAddressRIBSBVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListUnconfirmedTransactionsByAddressRIBSBScriptPubKey**](ListUnconfirmedTransactionsByAddressRIBSBScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSBVout

`func NewListUnconfirmedTransactionsByAddressRIBSBVout(isSpent bool, scriptPubKey ListUnconfirmedTransactionsByAddressRIBSBScriptPubKey, value string, ) *ListUnconfirmedTransactionsByAddressRIBSBVout`

NewListUnconfirmedTransactionsByAddressRIBSBVout instantiates a new ListUnconfirmedTransactionsByAddressRIBSBVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSBVoutWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSBVoutWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSBVout`

NewListUnconfirmedTransactionsByAddressRIBSBVoutWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSBVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) GetScriptPubKey() ListUnconfirmedTransactionsByAddressRIBSBScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) GetScriptPubKeyOk() (*ListUnconfirmedTransactionsByAddressRIBSBScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) SetScriptPubKey(v ListUnconfirmedTransactionsByAddressRIBSBScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSBVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


