# GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSDScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSDScriptPubKey.md) |  | 
**Value** | **string** | String representation of the amount | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSDScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSD2VoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSDScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSDScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSDScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSD2Vout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


