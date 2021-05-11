# ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinScriptSig**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin(addresses []string, scriptSig ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVinWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVinWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVinWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetScriptSig() ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetScriptSigOk() (*ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetScriptSig(v ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


