# ListConfirmedTransactionsByAddressRIBSD2VinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig**](ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSD2VinInner

`func NewListConfirmedTransactionsByAddressRIBSD2VinInner(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig, sequence string, txinwitness []string, vout int32, ) *ListConfirmedTransactionsByAddressRIBSD2VinInner`

NewListConfirmedTransactionsByAddressRIBSD2VinInner instantiates a new ListConfirmedTransactionsByAddressRIBSD2VinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSD2VinInnerWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSD2VinInnerWithDefaults() *ListConfirmedTransactionsByAddressRIBSD2VinInner`

NewListConfirmedTransactionsByAddressRIBSD2VinInnerWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSD2VinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetScriptSig() ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSD2VinInnerScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBSD2VinInner) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


