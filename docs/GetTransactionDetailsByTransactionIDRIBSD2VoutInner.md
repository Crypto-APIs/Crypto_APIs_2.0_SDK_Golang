# GetTransactionDetailsByTransactionIDRIBSD2VoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSD2VoutInnerScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSD2VoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSD2VoutInner

`func NewGetTransactionDetailsByTransactionIDRIBSD2VoutInner(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSD2VoutInnerScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDRIBSD2VoutInner`

NewGetTransactionDetailsByTransactionIDRIBSD2VoutInner instantiates a new GetTransactionDetailsByTransactionIDRIBSD2VoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSD2VoutInnerWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSD2VoutInnerWithDefaults() *GetTransactionDetailsByTransactionIDRIBSD2VoutInner`

NewGetTransactionDetailsByTransactionIDRIBSD2VoutInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSD2VoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSD2VoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSD2VoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSD2VoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSD2VoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


