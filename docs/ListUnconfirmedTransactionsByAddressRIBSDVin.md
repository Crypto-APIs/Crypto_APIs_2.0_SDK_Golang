# ListUnconfirmedTransactionsByAddressRIBSDVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSDScriptSig**](ListConfirmedTransactionsByAddressRIBSDScriptSig.md) |  | 
**Sequence** | **int32** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | Pointer to **int32** | It refers to the index of the output address of this transaction. The index starts from 0. | [optional] 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSDVin

`func NewListUnconfirmedTransactionsByAddressRIBSDVin(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSDScriptSig, sequence int32, txinwitness []string, value string, ) *ListUnconfirmedTransactionsByAddressRIBSDVin`

NewListUnconfirmedTransactionsByAddressRIBSDVin instantiates a new ListUnconfirmedTransactionsByAddressRIBSDVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSDVinWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSDVinWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSDVin`

NewListUnconfirmedTransactionsByAddressRIBSDVinWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSDVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSDScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSDScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSDScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSDVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


