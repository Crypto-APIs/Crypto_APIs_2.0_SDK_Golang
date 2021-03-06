# ListTransactionsByBlockHeightRIBSDVinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | **string** | Represents the coinbase hex. | 
**ScriptSig** | [**ListTransactionsByBlockHeightRIBSDVinInnerScriptSig**](ListTransactionsByBlockHeightRIBSDVinInnerScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListTransactionsByBlockHeightRIBSDVinInner

`func NewListTransactionsByBlockHeightRIBSDVinInner(addresses []string, coinbase string, scriptSig ListTransactionsByBlockHeightRIBSDVinInnerScriptSig, sequence string, txinwitness []string, vout int32, ) *ListTransactionsByBlockHeightRIBSDVinInner`

NewListTransactionsByBlockHeightRIBSDVinInner instantiates a new ListTransactionsByBlockHeightRIBSDVinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSDVinInnerWithDefaults

`func NewListTransactionsByBlockHeightRIBSDVinInnerWithDefaults() *ListTransactionsByBlockHeightRIBSDVinInner`

NewListTransactionsByBlockHeightRIBSDVinInnerWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSDVinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.


### GetScriptSig

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetScriptSig() ListTransactionsByBlockHeightRIBSDVinInnerScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetScriptSigOk() (*ListTransactionsByBlockHeightRIBSDVinInnerScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetScriptSig(v ListTransactionsByBlockHeightRIBSDVinInnerScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightRIBSDVinInner) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


