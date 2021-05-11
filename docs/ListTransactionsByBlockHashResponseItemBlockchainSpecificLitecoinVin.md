# ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVinWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVinWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVinWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


