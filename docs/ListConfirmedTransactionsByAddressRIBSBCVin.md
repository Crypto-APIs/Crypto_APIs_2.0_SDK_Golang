# ListConfirmedTransactionsByAddressRIBSBCVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDRIBSBScriptSig**](GetTransactionDetailsByTransactionIDRIBSBScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | Defines the vout of the transaction output, i.e. which output to spend. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSBCVin

`func NewListConfirmedTransactionsByAddressRIBSBCVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSBScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32, ) *ListConfirmedTransactionsByAddressRIBSBCVin`

NewListConfirmedTransactionsByAddressRIBSBCVin instantiates a new ListConfirmedTransactionsByAddressRIBSBCVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSBCVinWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSBCVinWithDefaults() *ListConfirmedTransactionsByAddressRIBSBCVin`

NewListConfirmedTransactionsByAddressRIBSBCVinWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSBCVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListConfirmedTransactionsByAddressRIBSBCVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


