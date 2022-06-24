# DecodeRawTransactionHexRISD2VinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the addresses which send/receive the amount. | [optional] 
**InputHash** | Pointer to **string** | Represents the transaction inputs&#39; indentifier. | [optional] 
**OutputIndex** | Pointer to **string** | Defines the output index of a transaction. | [optional] 
**ScriptSig** | [**DecodeRawTransactionHexRISD2VinInnerScriptSig**](DecodeRawTransactionHexRISD2VinInnerScriptSig.md) |  | 
**Sequence** | Pointer to **string** | Represents the script sequence number | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDecodeRawTransactionHexRISD2VinInner

`func NewDecodeRawTransactionHexRISD2VinInner(scriptSig DecodeRawTransactionHexRISD2VinInnerScriptSig, ) *DecodeRawTransactionHexRISD2VinInner`

NewDecodeRawTransactionHexRISD2VinInner instantiates a new DecodeRawTransactionHexRISD2VinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISD2VinInnerWithDefaults

`func NewDecodeRawTransactionHexRISD2VinInnerWithDefaults() *DecodeRawTransactionHexRISD2VinInner`

NewDecodeRawTransactionHexRISD2VinInnerWithDefaults instantiates a new DecodeRawTransactionHexRISD2VinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DecodeRawTransactionHexRISD2VinInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DecodeRawTransactionHexRISD2VinInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DecodeRawTransactionHexRISD2VinInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DecodeRawTransactionHexRISD2VinInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInputHash

`func (o *DecodeRawTransactionHexRISD2VinInner) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *DecodeRawTransactionHexRISD2VinInner) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *DecodeRawTransactionHexRISD2VinInner) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *DecodeRawTransactionHexRISD2VinInner) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### GetOutputIndex

`func (o *DecodeRawTransactionHexRISD2VinInner) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *DecodeRawTransactionHexRISD2VinInner) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *DecodeRawTransactionHexRISD2VinInner) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *DecodeRawTransactionHexRISD2VinInner) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetScriptSig

`func (o *DecodeRawTransactionHexRISD2VinInner) GetScriptSig() DecodeRawTransactionHexRISD2VinInnerScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *DecodeRawTransactionHexRISD2VinInner) GetScriptSigOk() (*DecodeRawTransactionHexRISD2VinInnerScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *DecodeRawTransactionHexRISD2VinInner) SetScriptSig(v DecodeRawTransactionHexRISD2VinInnerScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *DecodeRawTransactionHexRISD2VinInner) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DecodeRawTransactionHexRISD2VinInner) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DecodeRawTransactionHexRISD2VinInner) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DecodeRawTransactionHexRISD2VinInner) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTxinwitness

`func (o *DecodeRawTransactionHexRISD2VinInner) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *DecodeRawTransactionHexRISD2VinInner) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *DecodeRawTransactionHexRISD2VinInner) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *DecodeRawTransactionHexRISD2VinInner) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


