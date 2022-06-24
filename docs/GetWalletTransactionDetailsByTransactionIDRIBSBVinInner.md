# GetWalletTransactionDetailsByTransactionIDRIBSBVinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetWalletTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig**](GetWalletTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig.md) |  | 
**Sequence** | **int64** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSBVinInner

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBVinInner(addresses []string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig, sequence int64, txid string, vout int32, ) *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner`

NewGetWalletTransactionDetailsByTransactionIDRIBSBVinInner instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSBVinInnerWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSBVinInnerWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner`

NewGetWalletTransactionDetailsByTransactionIDRIBSBVinInnerWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSBVinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSBVinInnerScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.

### GetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSBVinInner) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


