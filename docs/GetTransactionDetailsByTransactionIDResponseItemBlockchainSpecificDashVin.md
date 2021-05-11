# GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashScriptSig**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashScriptSig, sequence string, txinwitness []string, ) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVinWithDefaults

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVinWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetScriptSig() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetScriptSig(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificDashVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


