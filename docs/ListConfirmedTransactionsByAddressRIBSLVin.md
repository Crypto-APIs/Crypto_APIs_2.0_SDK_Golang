# ListConfirmedTransactionsByAddressRIBSLVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSLScriptSig**](ListConfirmedTransactionsByAddressRIBSLScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | Defines the vout of the transaction output, i.e. which output to spend. | [optional] 

## Methods

### NewListConfirmedTransactionsByAddressRIBSLVin

`func NewListConfirmedTransactionsByAddressRIBSLVin(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSLScriptSig, sequence string, txid string, txinwitness []string, ) *ListConfirmedTransactionsByAddressRIBSLVin`

NewListConfirmedTransactionsByAddressRIBSLVin instantiates a new ListConfirmedTransactionsByAddressRIBSLVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSLVinWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSLVinWithDefaults() *ListConfirmedTransactionsByAddressRIBSLVin`

NewListConfirmedTransactionsByAddressRIBSLVinWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSLVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSLScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSLScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSLScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListConfirmedTransactionsByAddressRIBSLVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


