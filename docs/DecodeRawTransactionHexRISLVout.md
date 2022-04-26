# DecodeRawTransactionHexRISLVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScriptPubKey** | [**DecodeRawTransactionHexRISLScriptPubKey**](DecodeRawTransactionHexRISLScriptPubKey.md) |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 

## Methods

### NewDecodeRawTransactionHexRISLVout

`func NewDecodeRawTransactionHexRISLVout(scriptPubKey DecodeRawTransactionHexRISLScriptPubKey, ) *DecodeRawTransactionHexRISLVout`

NewDecodeRawTransactionHexRISLVout instantiates a new DecodeRawTransactionHexRISLVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISLVoutWithDefaults

`func NewDecodeRawTransactionHexRISLVoutWithDefaults() *DecodeRawTransactionHexRISLVout`

NewDecodeRawTransactionHexRISLVoutWithDefaults instantiates a new DecodeRawTransactionHexRISLVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScriptPubKey

`func (o *DecodeRawTransactionHexRISLVout) GetScriptPubKey() DecodeRawTransactionHexRISLScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *DecodeRawTransactionHexRISLVout) GetScriptPubKeyOk() (*DecodeRawTransactionHexRISLScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *DecodeRawTransactionHexRISLVout) SetScriptPubKey(v DecodeRawTransactionHexRISLScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *DecodeRawTransactionHexRISLVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DecodeRawTransactionHexRISLVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DecodeRawTransactionHexRISLVout) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DecodeRawTransactionHexRISLVout) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


