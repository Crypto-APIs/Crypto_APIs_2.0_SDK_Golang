# ListUnconfirmedTransactionsByAddressRIBSLVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListUnconfirmedTransactionsByAddressRIBSLScriptSig**](ListUnconfirmedTransactionsByAddressRIBSLScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | Pointer to **int32** | Defines the vout of the transaction output, i.e. which output to spend. | [optional] 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSLVin

`func NewListUnconfirmedTransactionsByAddressRIBSLVin(addresses []string, scriptSig ListUnconfirmedTransactionsByAddressRIBSLScriptSig, sequence string, txid string, txinwitness []string, value string, ) *ListUnconfirmedTransactionsByAddressRIBSLVin`

NewListUnconfirmedTransactionsByAddressRIBSLVin instantiates a new ListUnconfirmedTransactionsByAddressRIBSLVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSLVinWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSLVinWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSLVin`

NewListUnconfirmedTransactionsByAddressRIBSLVinWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSLVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetScriptSig() ListUnconfirmedTransactionsByAddressRIBSLScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetScriptSigOk() (*ListUnconfirmedTransactionsByAddressRIBSLScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetScriptSig(v ListUnconfirmedTransactionsByAddressRIBSLScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) SetVout(v int32)`

SetVout sets Vout field to given value.

### HasVout

`func (o *ListUnconfirmedTransactionsByAddressRIBSLVin) HasVout() bool`

HasVout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


