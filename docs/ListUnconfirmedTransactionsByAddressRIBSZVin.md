# ListUnconfirmedTransactionsByAddressRIBSZVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSZScriptSig**](ListConfirmedTransactionsByAddressRIBSZScriptSig.md) |  | 
**Sequence** | **int32** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Defines the specific amount. | 
**Vout** | **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSZVin

`func NewListUnconfirmedTransactionsByAddressRIBSZVin(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSZScriptSig, sequence int32, txid string, txinwitness []string, value string, vout int32, ) *ListUnconfirmedTransactionsByAddressRIBSZVin`

NewListUnconfirmedTransactionsByAddressRIBSZVin instantiates a new ListUnconfirmedTransactionsByAddressRIBSZVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSZVinWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSZVinWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSZVin`

NewListUnconfirmedTransactionsByAddressRIBSZVinWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSZVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSZScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSZScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSZScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSZVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


