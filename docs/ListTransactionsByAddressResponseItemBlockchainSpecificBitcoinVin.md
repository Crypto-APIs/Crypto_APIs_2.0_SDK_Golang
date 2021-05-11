# ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | Defines the vout of the transaction output, i.e. which output to spend. | [optional] 

## Methods

### NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin

`func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig, sequence string, txinwitness []string, ) *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin`

NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVinWithDefaults

`func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin`

NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


