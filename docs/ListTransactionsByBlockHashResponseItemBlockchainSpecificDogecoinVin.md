# ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig, sequence string, txinwitness []string, value string, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVinWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVinWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVinWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDogecoinScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificDogecoinVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


