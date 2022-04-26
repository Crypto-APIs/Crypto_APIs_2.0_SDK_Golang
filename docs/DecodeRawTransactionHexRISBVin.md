# DecodeRawTransactionHexRISBVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the address which send/receive the amount. | [optional] 
**InputHash** | Pointer to **string** | Represents the transaction inputs&#39; indentifier. | [optional] 
**OutputIndex** | Pointer to **int32** | Represents the output of a transaction. | [optional] 
**ScriptSig** | [**DecodeRawTransactionHexRISBScriptSig**](DecodeRawTransactionHexRISBScriptSig.md) |  | 
**Sequence** | Pointer to **string** | Represents the script sequence number. | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDecodeRawTransactionHexRISBVin

`func NewDecodeRawTransactionHexRISBVin(scriptSig DecodeRawTransactionHexRISBScriptSig, ) *DecodeRawTransactionHexRISBVin`

NewDecodeRawTransactionHexRISBVin instantiates a new DecodeRawTransactionHexRISBVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISBVinWithDefaults

`func NewDecodeRawTransactionHexRISBVinWithDefaults() *DecodeRawTransactionHexRISBVin`

NewDecodeRawTransactionHexRISBVinWithDefaults instantiates a new DecodeRawTransactionHexRISBVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DecodeRawTransactionHexRISBVin) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DecodeRawTransactionHexRISBVin) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DecodeRawTransactionHexRISBVin) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DecodeRawTransactionHexRISBVin) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInputHash

`func (o *DecodeRawTransactionHexRISBVin) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *DecodeRawTransactionHexRISBVin) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *DecodeRawTransactionHexRISBVin) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *DecodeRawTransactionHexRISBVin) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### GetOutputIndex

`func (o *DecodeRawTransactionHexRISBVin) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *DecodeRawTransactionHexRISBVin) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *DecodeRawTransactionHexRISBVin) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *DecodeRawTransactionHexRISBVin) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetScriptSig

`func (o *DecodeRawTransactionHexRISBVin) GetScriptSig() DecodeRawTransactionHexRISBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *DecodeRawTransactionHexRISBVin) GetScriptSigOk() (*DecodeRawTransactionHexRISBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *DecodeRawTransactionHexRISBVin) SetScriptSig(v DecodeRawTransactionHexRISBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *DecodeRawTransactionHexRISBVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DecodeRawTransactionHexRISBVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DecodeRawTransactionHexRISBVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DecodeRawTransactionHexRISBVin) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTxinwitness

`func (o *DecodeRawTransactionHexRISBVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *DecodeRawTransactionHexRISBVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *DecodeRawTransactionHexRISBVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *DecodeRawTransactionHexRISBVin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


