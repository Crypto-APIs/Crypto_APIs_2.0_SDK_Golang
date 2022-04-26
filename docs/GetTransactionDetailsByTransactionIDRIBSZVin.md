# GetTransactionDetailsByTransactionIDRIBSZVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDRIBSZScriptSig**](GetTransactionDetailsByTransactionIDRIBSZScriptSig.md) |  | 
**Sequence** | **int64** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Defines the specific amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSZVin

`func NewGetTransactionDetailsByTransactionIDRIBSZVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSZScriptSig, sequence int64, txid string, txinwitness []string, value string, vout int32, ) *GetTransactionDetailsByTransactionIDRIBSZVin`

NewGetTransactionDetailsByTransactionIDRIBSZVin instantiates a new GetTransactionDetailsByTransactionIDRIBSZVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSZVinWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSZVinWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVin`

NewGetTransactionDetailsByTransactionIDRIBSZVinWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSZScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSZScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSZScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetSequence(v int64)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *GetTransactionDetailsByTransactionIDRIBSZVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


