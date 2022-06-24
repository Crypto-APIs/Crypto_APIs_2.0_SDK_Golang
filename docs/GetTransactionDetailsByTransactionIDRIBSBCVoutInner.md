# GetTransactionDetailsByTransactionIDRIBSBCVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSBCVoutInner

`func NewGetTransactionDetailsByTransactionIDRIBSBCVoutInner(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDRIBSBCVoutInner`

NewGetTransactionDetailsByTransactionIDRIBSBCVoutInner instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBCVoutInner`

NewGetTransactionDetailsByTransactionIDRIBSBCVoutInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSBCVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


