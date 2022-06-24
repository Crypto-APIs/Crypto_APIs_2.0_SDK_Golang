# GetTransactionDetailsByTransactionIDRIBSZVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the transaction output has been spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSZVoutInnerScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSZVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the specific amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSZVoutInner

`func NewGetTransactionDetailsByTransactionIDRIBSZVoutInner(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSZVoutInnerScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDRIBSZVoutInner`

NewGetTransactionDetailsByTransactionIDRIBSZVoutInner instantiates a new GetTransactionDetailsByTransactionIDRIBSZVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSZVoutInnerWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSZVoutInnerWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVoutInner`

NewGetTransactionDetailsByTransactionIDRIBSZVoutInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSZVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSZVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSZVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSZVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


