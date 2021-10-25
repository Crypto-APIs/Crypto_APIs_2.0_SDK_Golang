# ListAllUnconfirmedTransactionsRIBSLVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**ListConfirmedTransactionsByAddressRIBSLScriptSig**](ListConfirmedTransactionsByAddressRIBSLScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | Defines the vout of the transaction output, i.e. which output to spend. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSLVin

`func NewListAllUnconfirmedTransactionsRIBSLVin(addresses []string, scriptSig ListConfirmedTransactionsByAddressRIBSLScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32, ) *ListAllUnconfirmedTransactionsRIBSLVin`

NewListAllUnconfirmedTransactionsRIBSLVin instantiates a new ListAllUnconfirmedTransactionsRIBSLVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSLVinWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSLVinWithDefaults() *ListAllUnconfirmedTransactionsRIBSLVin`

NewListAllUnconfirmedTransactionsRIBSLVinWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSLVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetScriptSig() ListConfirmedTransactionsByAddressRIBSLScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetScriptSigOk() (*ListConfirmedTransactionsByAddressRIBSLScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetScriptSig(v ListConfirmedTransactionsByAddressRIBSLScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSLVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


