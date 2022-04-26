# DecodeRawTransactionHexRISB2Vin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the address which send/receive the amount. | [optional] 
**InputHash** | Pointer to **string** | Represents the transaction inputs&#39; indentifier. | [optional] 
**OutputIndex** | Pointer to **string** | Defines the output index of a transaction. | [optional] 
**ScriptSig** | [**DecodeRawTransactionHexRISBScriptSig**](DecodeRawTransactionHexRISBScriptSig.md) |  | 
**Sequence** | Pointer to **string** | Represents the script sequence number. | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDecodeRawTransactionHexRISB2Vin

`func NewDecodeRawTransactionHexRISB2Vin(scriptSig DecodeRawTransactionHexRISBScriptSig, ) *DecodeRawTransactionHexRISB2Vin`

NewDecodeRawTransactionHexRISB2Vin instantiates a new DecodeRawTransactionHexRISB2Vin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeRawTransactionHexRISB2VinWithDefaults

`func NewDecodeRawTransactionHexRISB2VinWithDefaults() *DecodeRawTransactionHexRISB2Vin`

NewDecodeRawTransactionHexRISB2VinWithDefaults instantiates a new DecodeRawTransactionHexRISB2Vin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DecodeRawTransactionHexRISB2Vin) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DecodeRawTransactionHexRISB2Vin) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DecodeRawTransactionHexRISB2Vin) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DecodeRawTransactionHexRISB2Vin) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetInputHash

`func (o *DecodeRawTransactionHexRISB2Vin) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *DecodeRawTransactionHexRISB2Vin) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *DecodeRawTransactionHexRISB2Vin) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.

### HasInputHash

`func (o *DecodeRawTransactionHexRISB2Vin) HasInputHash() bool`

HasInputHash returns a boolean if a field has been set.

### GetOutputIndex

`func (o *DecodeRawTransactionHexRISB2Vin) GetOutputIndex() string`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *DecodeRawTransactionHexRISB2Vin) GetOutputIndexOk() (*string, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *DecodeRawTransactionHexRISB2Vin) SetOutputIndex(v string)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *DecodeRawTransactionHexRISB2Vin) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetScriptSig

`func (o *DecodeRawTransactionHexRISB2Vin) GetScriptSig() DecodeRawTransactionHexRISBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *DecodeRawTransactionHexRISB2Vin) GetScriptSigOk() (*DecodeRawTransactionHexRISBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *DecodeRawTransactionHexRISB2Vin) SetScriptSig(v DecodeRawTransactionHexRISBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *DecodeRawTransactionHexRISB2Vin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *DecodeRawTransactionHexRISB2Vin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *DecodeRawTransactionHexRISB2Vin) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *DecodeRawTransactionHexRISB2Vin) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTxinwitness

`func (o *DecodeRawTransactionHexRISB2Vin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *DecodeRawTransactionHexRISB2Vin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *DecodeRawTransactionHexRISB2Vin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *DecodeRawTransactionHexRISB2Vin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


