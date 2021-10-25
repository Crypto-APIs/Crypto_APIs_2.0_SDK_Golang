# ListTransactionsByBlockHeightRIBSZVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | **string** | Represents the coinbase hex. | 
**ScriptSig** | [**ListTransactionsByBlockHeightRIBSZScriptSig**](ListTransactionsByBlockHeightRIBSZScriptSig.md) |  | 
**Sequence** | **int32** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Defines the specific amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListTransactionsByBlockHeightRIBSZVin

`func NewListTransactionsByBlockHeightRIBSZVin(addresses []string, coinbase string, scriptSig ListTransactionsByBlockHeightRIBSZScriptSig, sequence int32, txid string, txinwitness []string, value string, vout int32, ) *ListTransactionsByBlockHeightRIBSZVin`

NewListTransactionsByBlockHeightRIBSZVin instantiates a new ListTransactionsByBlockHeightRIBSZVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightRIBSZVinWithDefaults

`func NewListTransactionsByBlockHeightRIBSZVinWithDefaults() *ListTransactionsByBlockHeightRIBSZVin`

NewListTransactionsByBlockHeightRIBSZVinWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSZVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.


### GetScriptSig

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetScriptSig() ListTransactionsByBlockHeightRIBSZScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetScriptSigOk() (*ListTransactionsByBlockHeightRIBSZScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetScriptSig(v ListTransactionsByBlockHeightRIBSZScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightRIBSZVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightRIBSZVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


