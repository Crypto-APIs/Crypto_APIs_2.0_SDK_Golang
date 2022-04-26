# GetWalletTransactionDetailsByTransactionIDRIBSDVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Coinbase** | Pointer to **string** | Represents the coinbase hex. | [optional] 
**ScriptSig** | [**GetWalletTransactionDetailsByTransactionIDRIBSDScriptSig**](GetWalletTransactionDetailsByTransactionIDRIBSDScriptSig.md) |  | 
**Sequence** | **int64** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewGetWalletTransactionDetailsByTransactionIDRIBSDVin

`func NewGetWalletTransactionDetailsByTransactionIDRIBSDVin(addresses []string, scriptSig GetWalletTransactionDetailsByTransactionIDRIBSDScriptSig, sequence int64, value string, ) *GetWalletTransactionDetailsByTransactionIDRIBSDVin`

NewGetWalletTransactionDetailsByTransactionIDRIBSDVin instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSDVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletTransactionDetailsByTransactionIDRIBSDVinWithDefaults

`func NewGetWalletTransactionDetailsByTransactionIDRIBSDVinWithDefaults() *GetWalletTransactionDetailsByTransactionIDRIBSDVin`

NewGetWalletTransactionDetailsByTransactionIDRIBSDVinWithDefaults instantiates a new GetWalletTransactionDetailsByTransactionIDRIBSDVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.

### GetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetScriptSig() GetWalletTransactionDetailsByTransactionIDRIBSDScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetScriptSigOk() (*GetWalletTransactionDetailsByTransactionIDRIBSDScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetScriptSig(v GetWalletTransactionDetailsByTransactionIDRIBSDScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.

### GetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *GetWalletTransactionDetailsByTransactionIDRIBSDVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


