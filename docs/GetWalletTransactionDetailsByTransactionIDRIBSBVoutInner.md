# GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey**](GetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSBVoutInner

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBVoutInner(isSpent bool, scriptPubKey GetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey, value string, ) *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner`

NewGetWalletTransactionDetailsByTransactionIDRIBSBVoutInner instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner`

NewGetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) GetScriptPubKey() GetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) GetScriptPubKeyOk() (*GetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) SetScriptPubKey(v GetWalletTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVoutInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


