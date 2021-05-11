# GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig, sequence string, txinwitness []string, ) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVinWithDefaults

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVinWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


