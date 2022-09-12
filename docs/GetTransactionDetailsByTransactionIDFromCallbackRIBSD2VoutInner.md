# GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerScriptPubKey**](GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerScriptPubKey.md) |  | 
**Value** | **string** | String representation of the amount | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) GetScriptPubKey() GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) SetScriptPubKey(v GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


