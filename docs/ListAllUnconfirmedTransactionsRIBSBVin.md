# ListAllUnconfirmedTransactionsRIBSBVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListAllUnconfirmedTransactionsRIBSBScriptSig**](ListAllUnconfirmedTransactionsRIBSBScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | Pointer to **string** | Represents the reference transaction identifier. | [optional] 
**Txinwitness** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **string** | Represents the sent/received amount. | [optional] 
**Vout** | **int32** | Defines the vout of the transaction output, i.e. which output to spend. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSBVin

`func NewListAllUnconfirmedTransactionsRIBSBVin(addresses []string, scriptSig ListAllUnconfirmedTransactionsRIBSBScriptSig, sequence string, vout int32, ) *ListAllUnconfirmedTransactionsRIBSBVin`

NewListAllUnconfirmedTransactionsRIBSBVin instantiates a new ListAllUnconfirmedTransactionsRIBSBVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSBVinWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSBVinWithDefaults() *ListAllUnconfirmedTransactionsRIBSBVin`

NewListAllUnconfirmedTransactionsRIBSBVinWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSBVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetScriptSig() ListAllUnconfirmedTransactionsRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetScriptSigOk() (*ListAllUnconfirmedTransactionsRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetScriptSig(v ListAllUnconfirmedTransactionsRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.

### HasTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) HasTxinwitness() bool`

HasTxinwitness returns a boolean if a field has been set.

### GetValue

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSBVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


