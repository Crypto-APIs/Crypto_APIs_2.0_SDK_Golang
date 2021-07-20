# GetTransactionDetailsByTransactionIDRIBSBCVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDRIBSBCScriptPubKey**](GetTransactionDetailsByTransactionIDRIBSBCScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSBCVout

`func NewGetTransactionDetailsByTransactionIDRIBSBCVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDRIBSBCScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDRIBSBCVout`

NewGetTransactionDetailsByTransactionIDRIBSBCVout instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSBCVoutWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSBCVoutWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBCVout`

NewGetTransactionDetailsByTransactionIDRIBSBCVoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDRIBSBCScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDRIBSBCScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDRIBSBCScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


