# ListUnspentTransactionOutputsByAddressRIVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDRIBSBScriptSig**](GetTransactionDetailsByTransactionIDRIBSBScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | Pointer to **[]string** |  | [optional] 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | Pointer to **int32** | Defines the vout of the transaction output, i.e. which output to spend. | [optional] 

## Methods

### NewListUnspentTransactionOutputsByAddressRIVin

`func NewListUnspentTransactionOutputsByAddressRIVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSBScriptSig, sequence string, txid string, value string, ) *ListUnspentTransactionOutputsByAddressRIVin`

NewListUnspentTransactionOutputsByAddressRIVin instantiates a new ListUnspentTransactionOutputsByAddressRIVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnspentTransactionOutputsByAddressRIVinWithDefaults

`func NewListUnspentTransactionOutputsByAddressRIVinWithDefaults() *ListUnspentTransactionOutputsByAddressRIVin`

NewListUnspentTransactionOutputsByAddressRIVinWithDefaults instantiates a new ListUnspentTransactionOutputsByAddressRIVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListUnspentTransactionOutputsByAddressRIVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *ListUnspentTransactionOutputsByAddressRIVin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.

### GetValue

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnspentTransactionOutputsByAddressRIVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnspentTransactionOutputsByAddressRIVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListUnspentTransactionOutputsByAddressRIVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


