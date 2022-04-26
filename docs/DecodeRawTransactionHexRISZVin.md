# DecodeRawTransactionHexRISZVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the addresses which send/receive the amount. | [optional] 
**InputHash** | Pointer to **string** | Represents the transaction inputs&#39; indentifier. | [optional] 
**OutputIndex** | Pointer to **string** | Defines the output index of a transaction. | [optional] 
**ScriptSig** | [**DecodeRawTransactionHexRISZScriptSig**](DecodeRawTransactionHexRISZScriptSig.md) |  | 
**Sequence** | Pointer to **string** | Represents the script sequence number. | [optional] 

## Methods

### NewDecodeRawTransactionHexRISZVin

`func NewDecodeRawTransactionHexRISZVin(scriptSig DecodeRawTransactionHexRISZScriptSig, ) *DecodeRawTransactionHexRISZVin`

NewDecodeRawTransactionHexRISZVin instantiates a new DecodeRawTransactionHexRISZVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISZVinWithDefaults

`func NewDecodeRawTransactionHexRISZVinWithDefaults() *DecodeRawTransactionHexRISZVin`

NewDecodeRawTransactionHexRISZVinWithDefaults instantiates a new DecodeRawTransactionHexRISZVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DecodeRawTransactionHexRISZVin) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DecodeRawTransactionHexRISZVin) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DecodeRawTransactionHexRISZVin) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DecodeRawTransactionHexRISZVin) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInputHash

`func (o *DecodeRawTransactionHexRISZVin) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *DecodeRawTransactionHexRISZVin) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *DecodeRawTransactionHexRISZVin) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *DecodeRawTransactionHexRISZVin) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### GetOutputIndex

`func (o *DecodeRawTransactionHexRISZVin) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *DecodeRawTransactionHexRISZVin) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *DecodeRawTransactionHexRISZVin) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *DecodeRawTransactionHexRISZVin) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetScriptSig

`func (o *DecodeRawTransactionHexRISZVin) GetScriptSig() DecodeRawTransactionHexRISZScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *DecodeRawTransactionHexRISZVin) GetScriptSigOk() (*DecodeRawTransactionHexRISZScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *DecodeRawTransactionHexRISZVin) SetScriptSig(v DecodeRawTransactionHexRISZScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *DecodeRawTransactionHexRISZVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DecodeRawTransactionHexRISZVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DecodeRawTransactionHexRISZVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DecodeRawTransactionHexRISZVin) HasSequence() bool`

HasSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


