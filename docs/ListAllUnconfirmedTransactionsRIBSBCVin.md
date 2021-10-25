# ListAllUnconfirmedTransactionsRIBSBCVin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**ScriptSig** | [**GetTransactionDetailsByTransactionIDRIBSBScriptSig**](GetTransactionDetailsByTransactionIDRIBSBScriptSig.md) |  | 
**Sequence** | **string** | Represents the script sequence number. | 
**Txid** | **string** | Represents the reference transaction identifier. | 
**Txinwitness** | **[]string** |  | 
**Value** | **string** | Represents the sent/received amount. | 
**Vout** | **int32** | Defines the vout of the transaction output, i.e. which output to spend. | 

## Methods

### NewListAllUnconfirmedTransactionsRIBSBCVin

`func NewListAllUnconfirmedTransactionsRIBSBCVin(addresses []string, scriptSig GetTransactionDetailsByTransactionIDRIBSBScriptSig, sequence string, txid string, txinwitness []string, value string, vout int32, ) *ListAllUnconfirmedTransactionsRIBSBCVin`

NewListAllUnconfirmedTransactionsRIBSBCVin instantiates a new ListAllUnconfirmedTransactionsRIBSBCVin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIBSBCVinWithDefaults

`func NewListAllUnconfirmedTransactionsRIBSBCVinWithDefaults() *ListAllUnconfirmedTransactionsRIBSBCVin`

NewListAllUnconfirmedTransactionsRIBSBCVinWithDefaults instantiates a new ListAllUnconfirmedTransactionsRIBSBCVin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetScriptSig() GetTransactionDetailsByTransactionIDRIBSBScriptSig`

GetScriptSig returns the ScriptSig field if non-nil, zero value otherwise.

### GetScriptSigOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetScriptSigOk() (*GetTransactionDetailsByTransactionIDRIBSBScriptSig, bool)`

GetScriptSigOk returns a tuple with the ScriptSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSig

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetScriptSig(v GetTransactionDetailsByTransactionIDRIBSBScriptSig)`

SetScriptSig sets ScriptSig field to given value.


### GetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetTxid(v string)`

SetTxid sets Txid field to given value.


### GetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetTxinwitness() []string`

GetTxinwitness returns the Txinwitness field if non-nil, zero value otherwise.

### GetTxinwitnessOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetTxinwitnessOk() (*[]string, bool)`

GetTxinwitnessOk returns a tuple with the Txinwitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxinwitness

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetTxinwitness(v []string)`

SetTxinwitness sets Txinwitness field to given value.


### GetValue

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetValue(v string)`

SetValue sets Value field to given value.


### GetVout

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetVout() int32`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) GetVoutOk() (*int32, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListAllUnconfirmedTransactionsRIBSBCVin) SetVout(v int32)`

SetVout sets Vout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


