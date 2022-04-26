# GetTransactionDetailsByTransactionIDRIBSBCVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDRIBSBScriptSig**](GetTransactionDetailsByTransactionIDRIBSBScriptSig.md) |  | 
**Sequence** | **int64** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSBCVin

`func NewGetTransactionDetailsByTransactionIDRIBSBCVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSBScriptSig, sequence int64, txinwitness []string, ) *GetTransactionDetailsByTransactionIDRIBSBCVin`

NewGetTransactionDetailsByTransactionIDRIBSBCVin instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSBCVinWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSBCVinWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBCVin`

NewGetTransactionDetailsByTransactionIDRIBSBCVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBCVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *GetTransactionDetailsByTransactionIDRIBSBCVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


