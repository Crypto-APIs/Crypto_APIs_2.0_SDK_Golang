# ListConfirmedTransactionsByAddressRIBSDVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | **string** | Represents the coinbase hex. | 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSDScriptSig**](ListConfirmedTransactionsByAddressRIBSDScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewListConfirmedTransactionsByAddressRIBSDVin

`func NewListConfirmedTransactionsByAddressRIBSDVin(addresses []string, coinbase string, scriptSig ListConfirmedTransactionsByAddressRIBSDScriptSig, sequence string, txinwitness []string, value string, ) *ListConfirmedTransactionsByAddressRIBSDVin`

NewListConfirmedTransactionsByAddressRIBSDVin instantiates a new ListConfirmedTransactionsByAddressRIBSDVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSDVinWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSDVinWithDefaults() *ListConfirmedTransactionsByAddressRIBSDVin`

NewListConfirmedTransactionsByAddressRIBSDVinWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSDVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.


### GetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSDScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSDScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSDScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListConfirmedTransactionsByAddressRIBSDVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


