# GetTransactionDetailsByTransactionIDRIBSD2Vout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSD2Vout

`func NewGetTransactionDetailsByTransactionIDRIBSD2Vout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDRIBSD2Vout`

NewGetTransactionDetailsByTransactionIDRIBSD2Vout instantiates a new GetTransactionDetailsByTransactionIDRIBSD2Vout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSD2VoutWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSD2VoutWithDefaults() *GetTransactionDetailsByTransactionIDRIBSD2Vout`

NewGetTransactionDetailsByTransactionIDRIBSD2VoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSD2Vout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSD2ScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSD2Vout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


