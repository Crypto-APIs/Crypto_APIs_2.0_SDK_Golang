# ListUnconfirmedTransactionsByAddressRIBSLVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSLScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSLScriptPubKey.md) |  | 
**Value** | **string** | String representation of the amount | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSLVout

`func NewListUnconfirmedTransactionsByAddressRIBSLVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSLScriptPubKey, value string, ) *ListUnconfirmedTransactionsByAddressRIBSLVout`

NewListUnconfirmedTransactionsByAddressRIBSLVout instantiates a new ListUnconfirmedTransactionsByAddressRIBSLVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSLVoutWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSLVoutWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSLVout`

NewListUnconfirmedTransactionsByAddressRIBSLVoutWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSLVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSLScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSLScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSLScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


