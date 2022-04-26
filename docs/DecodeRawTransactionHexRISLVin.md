# DecodeRawTransactionHexRISLVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the addresses which send/receive the amount. | [optional] 
**InputHash** | Pointer to **string** | Represents the transaction inputs&#39; indentifier. | [optional] 
**OutputIndex** | Pointer to **string** | Defines the output index of a transaction. | [optional] 
**ScriptSig** | [**DecodeRawTransactionHexRISLScriptSig**](DecodeRawTransactionHexRISLScriptSig.md) |  | 
**Sequence** | Pointer to **string** | Represents the script sequence number. | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDecodeRawTransactionHexRISLVin

`func NewDecodeRawTransactionHexRISLVin(scriptSig DecodeRawTransactionHexRISLScriptSig, ) *DecodeRawTransactionHexRISLVin`

NewDecodeRawTransactionHexRISLVin instantiates a new DecodeRawTransactionHexRISLVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISLVinWithDefaults

`func NewDecodeRawTransactionHexRISLVinWithDefaults() *DecodeRawTransactionHexRISLVin`

NewDecodeRawTransactionHexRISLVinWithDefaults instantiates a new DecodeRawTransactionHexRISLVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DecodeRawTransactionHexRISLVin) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DecodeRawTransactionHexRISLVin) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DecodeRawTransactionHexRISLVin) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DecodeRawTransactionHexRISLVin) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInputHash

`func (o *DecodeRawTransactionHexRISLVin) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *DecodeRawTransactionHexRISLVin) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *DecodeRawTransactionHexRISLVin) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *DecodeRawTransactionHexRISLVin) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### GetOutputIndex

`func (o *DecodeRawTransactionHexRISLVin) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *DecodeRawTransactionHexRISLVin) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *DecodeRawTransactionHexRISLVin) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *DecodeRawTransactionHexRISLVin) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetScriptSig

`func (o *DecodeRawTransactionHexRISLVin) GetScriptSig() DecodeRawTransactionHexRISLScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *DecodeRawTransactionHexRISLVin) GetScriptSigOk() (*DecodeRawTransactionHexRISLScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *DecodeRawTransactionHexRISLVin) SetScriptSig(v DecodeRawTransactionHexRISLScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *DecodeRawTransactionHexRISLVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DecodeRawTransactionHexRISLVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DecodeRawTransactionHexRISLVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DecodeRawTransactionHexRISLVin) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTxinwitness

`func (o *DecodeRawTransactionHexRISLVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *DecodeRawTransactionHexRISLVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *DecodeRawTransactionHexRISLVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *DecodeRawTransactionHexRISLVin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


