# ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinScriptSig**](ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | Defines the vout of the transaction output, i.e. which output to spend. | [optional] 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin

`func NewListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin(addresses []string, scriptSig ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinScriptSig, sequence string, txid string, txinwitness []string, ) *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin`

NewListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVinWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin`

NewListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetScriptSig() ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetScriptSigOk() (*ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetScriptSig(v ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificLitecoinVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


