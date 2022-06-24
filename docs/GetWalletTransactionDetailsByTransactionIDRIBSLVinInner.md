# GetWalletTransactionDetailsByTransactionIDRIBSLVinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | **string** | Represents the coinbase hex. | 
**ScriptSig** | [**GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig**](GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig.md) |  | 
**Sequence** | **int64** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | Pointer to **[]string** |  | [optional] 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInner

`func NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInner(addresses []string, coinbase string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig, sequence int64, txid string, value string, vout int32, ) *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner`

NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInner instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSLVinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInnerWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInnerWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner`

NewGetWalletTransactionDetailsByTransactionIDRIBSLVinInnerWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSLVinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.


### GetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSLVinInnerScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.

### GetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSLVinInner) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


