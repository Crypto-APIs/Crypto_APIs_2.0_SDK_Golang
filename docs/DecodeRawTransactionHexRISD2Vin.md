# DecodeRawTransactionHexRISD2Vin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the addresses which send/receive the amount. | [optional] 
**InputHash** | Pointer to **string** | Represents the transaction inputs&#39; indentifier. | [optional] 
**OutputIndex** | Pointer to **string** | Defines the output index of a transaction. | [optional] 
**ScriptSig** | [**DecodeRawTransactionHexRISD2ScriptSig**](DecodeRawTransactionHexRISD2ScriptSig.md) |  | 
**Sequence** | Pointer to **string** | Represents the script sequence number | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDecodeRawTransactionHexRISD2Vin

`func NewDecodeRawTransactionHexRISD2Vin(scriptSig DecodeRawTransactionHexRISD2ScriptSig, ) *DecodeRawTransactionHexRISD2Vin`

NewDecodeRawTransactionHexRISD2Vin instantiates a new DecodeRawTransactionHexRISD2Vin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISD2VinWithDefaults

`func NewDecodeRawTransactionHexRISD2VinWithDefaults() *DecodeRawTransactionHexRISD2Vin`

NewDecodeRawTransactionHexRISD2VinWithDefaults instantiates a new DecodeRawTransactionHexRISD2Vin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DecodeRawTransactionHexRISD2Vin) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DecodeRawTransactionHexRISD2Vin) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DecodeRawTransactionHexRISD2Vin) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DecodeRawTransactionHexRISD2Vin) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInputHash

`func (o *DecodeRawTransactionHexRISD2Vin) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *DecodeRawTransactionHexRISD2Vin) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *DecodeRawTransactionHexRISD2Vin) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *DecodeRawTransactionHexRISD2Vin) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### GetOutputIndex

`func (o *DecodeRawTransactionHexRISD2Vin) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *DecodeRawTransactionHexRISD2Vin) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *DecodeRawTransactionHexRISD2Vin) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *DecodeRawTransactionHexRISD2Vin) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetScriptSig

`func (o *DecodeRawTransactionHexRISD2Vin) GetScriptSig() DecodeRawTransactionHexRISD2ScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *DecodeRawTransactionHexRISD2Vin) GetScriptSigOk() (*DecodeRawTransactionHexRISD2ScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *DecodeRawTransactionHexRISD2Vin) SetScriptSig(v DecodeRawTransactionHexRISD2ScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *DecodeRawTransactionHexRISD2Vin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DecodeRawTransactionHexRISD2Vin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DecodeRawTransactionHexRISD2Vin) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DecodeRawTransactionHexRISD2Vin) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTxinwitness

`func (o *DecodeRawTransactionHexRISD2Vin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *DecodeRawTransactionHexRISD2Vin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *DecodeRawTransactionHexRISD2Vin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *DecodeRawTransactionHexRISD2Vin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


