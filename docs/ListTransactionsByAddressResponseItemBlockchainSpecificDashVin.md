# ListTransactionsByAddressResponseItemBlockchainSpecificDashVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**ListTransactionsByAddressResponseItemBlockchainSpecificDashScriptSig**](ListTransactionsByAddressResponseItemBlockchainSpecificDashScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecificDashVin

`func NewListTransactionsByAddressResponseItemBlockchainSpecificDashVin(addresses []string, scriptSig ListTransactionsByAddressResponseItemBlockchainSpecificDashScriptSig, sequence string, txinwitness []string, vout int32, ) *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin`

NewListTransactionsByAddressResponseItemBlockchainSpecificDashVin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDashVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificDashVinWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificDashVinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin`

NewListTransactionsByAddressResponseItemBlockchainSpecificDashVinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificDashVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetScriptSig() ListTransactionsByAddressResponseItemBlockchainSpecificDashScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetScriptSigOk() (*ListTransactionsByAddressResponseItemBlockchainSpecificDashScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetScriptSig(v ListTransactionsByAddressResponseItemBlockchainSpecificDashScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificDashVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


