# ListUnconfirmedTransactionsByAddressRIBSD2Vin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSD2ScriptSig**](ListConfirmedTransactionsByAddressRIBSD2ScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | Pointer to **string** | String representation of the amount | [optional] 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSD2Vin

`func NewListUnconfirmedTransactionsByAddressRIBSD2Vin(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSD2ScriptSig, sequence string, txid string, txinwitness []string, vout int32, ) *ListUnconfirmedTransactionsByAddressRIBSD2Vin`

NewListUnconfirmedTransactionsByAddressRIBSD2Vin instantiates a new ListUnconfirmedTransactionsByAddressRIBSD2Vin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSD2VinWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSD2VinWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSD2Vin`

NewListUnconfirmedTransactionsByAddressRIBSD2VinWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSD2Vin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSD2ScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSD2ScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSD2ScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSD2Vin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


