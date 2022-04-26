# GetWalletTransactionDetailsByTransactionIDRIBSBVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig**](GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig.md) |  | 
**Sequence** | **int64** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSBVin

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBVin(addresses []string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig, sequence int64, txid string, vout int32, ) *GetWalletTransactionDetailsByTransactionIDRIBSBVin`

NewGetWalletTransactionDetailsByTransactionIDRIBSBVin instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSBVinWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBVinWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSBVin`

NewGetWalletTransactionDetailsByTransactionIDRIBSBVinWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.

### GetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


