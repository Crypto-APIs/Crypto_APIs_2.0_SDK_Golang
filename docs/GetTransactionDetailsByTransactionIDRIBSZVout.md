# GetTransactionDetailsByTransactionIDRIBSZVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the transaction output has been spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSZScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSZScriptPubKey.md) |  | 
**Value** | **string** | Represents the specific amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSZVout

`func NewGetTransactionDetailsByTransactionIDRIBSZVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSZScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDRIBSZVout`

NewGetTransactionDetailsByTransactionIDRIBSZVout instantiates a new GetTransactionDetailsByTransactionIDRIBSZVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSZVoutWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSZVoutWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVout`

NewGetTransactionDetailsByTransactionIDRIBSZVoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSZScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSZScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSZScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSZVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


