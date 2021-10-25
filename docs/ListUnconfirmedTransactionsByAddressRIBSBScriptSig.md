# ListUnconfirmedTransactionsByAddressRIBSBScriptSig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asm** | **string** | The asm strands for assembly, which is the symbolic representation of the Bitcoin&#39;s Script language op-codes. | 
**Hex** | Pointer to **string** | Represents the hex of the public key of the address. | [optional] 
**Type** | **string** | Represents the script type of the reference transaction identifier. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIBSBScriptSig

`func NewListUnconfirmedTransactionsByAddressRIBSBScriptSig(asm string, type_ string, ) *ListUnconfirmedTransactionsByAddressRIBSBScriptSig`

NewListUnconfirmedTransactionsByAddressRIBSBScriptSig instantiates a new ListUnconfirmedTransactionsByAddressRIBSBScriptSig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIBSBScriptSigWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIBSBScriptSigWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSBScriptSig`

NewListUnconfirmedTransactionsByAddressRIBSBScriptSigWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSBScriptSig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsm

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetAsm() string`

GetAsm returns the Asm field if non-nil, zero value otherwise.

### GetAsmOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetAsmOk() (*string, bool)`

GetAsmOk returns a tuple with the Asm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsm

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) SetAsm(v string)`

SetAsm sets Asm field to given value.


### GetHex

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) SetHex(v string)`

SetHex sets Hex field to given value.

### HasHex

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) HasHex() bool`

HasHex returns a boolean if a field has been set.

### GetType

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListUnconfirmedTransactionsByAddressRIBSBScriptSig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


